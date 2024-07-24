package handlers

import (
	"fmt"
	"math/big"
	"net/http"

	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/jsonapi"
	zkptypes "github.com/iden3/go-rapidsnark/types"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	zk "github.com/rarimo/zkverifier-kit"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

// VerifyPassport handler processes 3 different flows:
//   - Old passport verification with proof for the current release
//   - New passport verification with JWT for the future
//   - Legacy joining program logic when the client fails to generate query proof
func VerifyPassport(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVerifyPassport(r)
	if err != nil {
		Log(r).WithError(err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var (
		anonymousID = req.Data.Attributes.AnonymousId
		proof       = req.Data.Attributes.Proof
		log         = Log(r).WithFields(map[string]any{
			"balance.nullifier":    req.Data.ID,
			"balance.anonymous_id": anonymousID,
		})

		gotSig = r.Header.Get("Signature")
	)

	wantSig, err := SigCalculator(r).PassportVerificationSignature(req.Data.ID, anonymousID)
	if err != nil { // must never happen due to preceding validation
		Log(r).WithError(err).Error("Failed to calculate HMAC signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if gotSig != wantSig {
		log.Warnf("Passport verification unauthorized access: HMAC signature mismatch: got %s, want %s", gotSig, wantSig)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	balance, errs := getAndVerifyBalanceEligibility(r, req.Data.ID, proof)
	if len(errs) > 0 {
		ape.RenderErr(w, errs...)
		return
	}

	byAnonymousID, err := BalancesQ(r).FilterByAnonymousID(anonymousID).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get balance by anonymous ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if byAnonymousID != nil && byAnonymousID.Nullifier != balance.Nullifier {
		log.Warn("Balance with the same anonymous ID already exists")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	// UserClaims(r)[0] will not panic because of authorization validation
	sharedHash := UserClaims(r)[0].SharedHash
	if proof != nil {
		sig := zk.PubSignalGetter{
			ProofType: zk.GeorgianPassport,
			Signals:   proof.PubSignals,
		}
		h := sig.Get(zk.PersonalNumberHash)
		if h == "" {
			log.Errorf("Shared hash was not obtained for valid proof: %+v", proof)
			ape.RenderErr(w, problems.InternalError())
			return
		}
		sharedHash = &h
	}

	if balance.IsVerified {
		if balance.SharedHash != nil {
			log.Warnf("Balance %s already verified", balance.Nullifier)
			ape.RenderErr(w, problems.Conflict())
			return
		}
		if proof == nil {
			log.Warnf("Balance %s tried to re-join program", balance.Nullifier)
			ape.RenderErr(w, problems.Conflict())
			return
		}

		var balAID string
		if balance.AnonymousID != nil {
			balAID = *balance.AnonymousID
		}

		err = validation.Errors{
			"data/attributes/anonymous_id": validation.Validate(anonymousID, validation.Required, validation.In(balAID)),
		}.Filter()
		if err != nil {
			log.Warnf("Anonymous ID was changed, got %s, want %s", anonymousID, balAID)
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		log.Debug("Balance has joined program previously, updating shared hash")
		err = BalancesQ(r).FilterByNullifier(balance.Nullifier).Update(map[string]any{
			data.ColSharedHash: *sharedHash,
		})
		if err != nil {
			log.WithError(err).Error("Failed to update balance")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		ape.Render(w, newEventClaimingStateResponse(req.Data.ID, false))
		return
	}

	err = EventsQ(r).Transaction(func() error {
		if err = updateBalanceVerification(r, *balance, anonymousID, sharedHash); err != nil {
			return fmt.Errorf("update balance verification info: %w", err)
		}

		if balance.ReferredBy == nil {
			log.Debug("Balance is disabled, events will not be claimed")
			return nil
		}

		return doVerificationEventUpdates(r, *balance)
	})

	if err != nil {
		log.WithError(err).Error("Failed to do passport scan updates")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	event, err := EventsQ(r).FilterByNullifier(balance.Nullifier).
		FilterByType(models.TypePassportScan).
		FilterByStatus(data.EventClaimed).
		Get()
	if err != nil {
		log.WithError(err).Error("Failed to get claimed event")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newEventClaimingStateResponse(req.Data.ID, event != nil))
}

func newEventClaimingStateResponse(id string, isClaimed bool) resources.EventClaimingStateResponse {
	var res resources.EventClaimingStateResponse
	res.Data.ID = id
	res.Data.Type = resources.EVENT_CLAIMING_STATE
	res.Data.Attributes.Claimed = isClaimed
	return res
}

// getAndVerifyBalanceEligibility provides shared logic to verify that the user
// is eligible to verify passport or withdraw. Some extra checks still exist in
// the flows. You may provide nil proof to handle its verification outside.
func getAndVerifyBalanceEligibility(
	r *http.Request,
	nullifier string,
	proof *zkptypes.ZKProof,
) (balance *data.Balance, errs []*jsonapi.ErrorObject) {

	log := Log(r).WithField("balance.nullifier", nullifier)
	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
		return nil, append(errs, problems.Unauthorized())
	}

	balance, err := BalancesQ(r).FilterByNullifier(nullifier).Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get balance by nullifier")
		return nil, append(errs, problems.InternalError())
	}

	if balance == nil {
		log.Debug("Balance not found")
		return nil, append(errs, problems.NotFound())
	}
	if proof == nil {
		log.Debug("Proof is not provided and will not be verified")
		return balance, nil
	}

	// never panics because of request validation
	ni := zk.Indexes(zk.GeorgianPassport)[zk.Nullifier]
	proof.PubSignals[ni] = mustHexToInt(nullifier)
	err = Verifiers(r).Passport.VerifyProof(*proof)
	if err != nil {
		var vErr validation.Errors
		if !errors.As(err, &vErr) {
			log.WithError(err).Error("Failed to verify proof")
			return nil, append(errs, problems.InternalError())
		}
		return nil, problems.BadRequest(err)
	}

	log.Debug("Passport proof successfully verified")
	return balance, nil
}

func updateBalanceVerification(r *http.Request, balance data.Balance, anonymousID string, sharedHash *string) error {
	toUpd := map[string]any{
		data.ColIsVerified:  true,
		data.ColAnonymousID: anonymousID,
	}
	if sharedHash != nil {
		toUpd[data.ColSharedHash] = *sharedHash
	}

	err := BalancesQ(r).FilterByNullifier(balance.Nullifier).Update(toUpd)
	if err != nil {
		return fmt.Errorf("update balance: %w", err)
	}

	return nil
}

func doVerificationEventUpdates(r *http.Request, balance data.Balance) error {
	if err := fulfillOrClaimPassportScanEvent(r, balance); err != nil {
		return fmt.Errorf("fulfill passport scan event: %w", err)
	}

	if err := claimBeReferredEvent(r, balance); err != nil {
		return fmt.Errorf("failed to claim be referred event: %w", err)
	}
	if err := claimReferralSpecificEvents(r, balance.Nullifier); err != nil {
		return fmt.Errorf("failed to claim referral specific events: %w", err)
	}
	if err := addEventForReferrer(r, balance); err != nil {
		return fmt.Errorf("add event for referrer: %w", err)
	}

	Log(r).WithField("balance.nullifier", balance.Nullifier).
		Debug("All verification-related events successfully updated")
	return nil
}

// fulfillOrClaimPassportScanEvent Fulfill passport scan event for user if event
// active. Event can be automatically claimed if auto-claim is enabled.
func fulfillOrClaimPassportScanEvent(r *http.Request, balance data.Balance) error {
	evTypePassport := EventTypes(r).Get(models.TypePassportScan, evtypes.FilterInactive)
	if evTypePassport == nil {
		Log(r).Debug("Passport scan event type is inactive")
		return nil
	}

	// event could also be fulfilled when balance was activated
	event, err := EventsQ(r).
		FilterByNullifier(balance.Nullifier).
		FilterByType(models.TypePassportScan).
		FilterByStatus(data.EventOpen, data.EventFulfilled).
		Get()
	if err != nil {
		return fmt.Errorf("get open passport scan event: %w", err)
	}
	if event == nil {
		return errors.New("inconsistent state: balance is not verified, event type is active, but no open or fulfilled event was found")
	}

	if !evTypePassport.AutoClaim || balance.ReferredBy == nil {
		Log(r).Debug("Passport scan event is not auto-claimed or balance is disabled, event will be fulfilled")
		_, err = EventsQ(r).
			FilterByID(event.ID).
			Update(data.EventFulfilled, nil, nil)
		if err != nil {
			return fmt.Errorf("failed to update event: %w", err)
		}

		return nil
	}

	_, err = EventsQ(r).FilterByID(event.ID).Update(data.EventClaimed, nil, &evTypePassport.Reward)
	if err != nil {
		return fmt.Errorf("update event status: %w", err)
	}

	err = DoClaimEventUpdates(
		Levels(r),
		ReferralsQ(r),
		BalancesQ(r),
		balance,
		evTypePassport.Reward)
	if err != nil {
		return fmt.Errorf("failed to do claim event updates for passport scan: %w", err)
	}

	return nil
}

// claimReferralSpecificEvents Claim events for invited friends who scanned the
// passport. This is possible when the user registered in the referral program
// and invited friends, the friends scanned the passport, but since the user
// hadn't a supported passport, the event could not be claimed. And now that user
// has scanned the passport, it is necessary to claim events for user's friends
// if auto-claim is enabled.
func claimReferralSpecificEvents(r *http.Request, nullifier string) error {
	evTypeRef := EventTypes(r).Get(models.TypeReferralSpecific, evtypes.FilterInactive)
	if evTypeRef == nil || !evTypeRef.AutoClaim {
		Log(r).Debugf("Referral specific event is inactive or cannot be auto-claimed")
		return nil
	}

	balance, err := BalancesQ(r).FilterByNullifier(nullifier).FilterDisabled().Get()
	if err != nil || balance == nil { // must not be nil due to previous logic
		return fmt.Errorf("failed to get balance: %w", err)
	}

	events, err := EventsQ(r).
		FilterByNullifier(balance.Nullifier).
		FilterByType(models.TypeReferralSpecific).
		FilterByStatus(data.EventFulfilled).
		Select()
	if err != nil {
		return fmt.Errorf("get fulfilled referral specific events: %w", err)
	}

	eventsToClaimed := make([]string, len(events))
	for i := 0; i < len(events); i++ {
		eventsToClaimed[i] = events[i].ID
	}

	if len(eventsToClaimed) == 0 {
		return nil
	}

	_, err = EventsQ(r).FilterByID(eventsToClaimed...).Update(data.EventClaimed, nil, &evTypeRef.Reward)
	if err != nil {
		return fmt.Errorf("update event status: %w", err)
	}

	err = DoClaimEventUpdates(
		Levels(r),
		ReferralsQ(r),
		BalancesQ(r),
		*balance,
		int64(len(events))*evTypeRef.Reward)
	if err != nil {
		return fmt.Errorf("failed to do claim event updates for referral specific events: %w", err)
	}

	return nil
}

// claimBeReferredEvent get fulfilled be_referred event and claims it if user is
// invited with non-genesis referral code
func claimBeReferredEvent(r *http.Request, balance data.Balance) error {
	evTypeBeRef := EventTypes(r).Get(models.TypeBeReferred, evtypes.FilterInactive)
	if evTypeBeRef == nil || !evTypeBeRef.AutoClaim {
		return nil
	}

	event, err := EventsQ(r).FilterByNullifier(balance.Nullifier).
		FilterByType(models.TypeBeReferred).
		FilterByStatus(data.EventFulfilled).
		Get()
	if err != nil {
		return fmt.Errorf("get fulfilled be_referred event: %w", err)
	}
	if event == nil {
		Log(r).Debug("User is not eligible for be_referred event")
		return nil
	}

	_, err = EventsQ(r).FilterByID(event.ID).Update(data.EventClaimed, nil, &evTypeBeRef.Reward)
	if err != nil {
		return fmt.Errorf("update event status: %w", err)
	}

	err = DoClaimEventUpdates(
		Levels(r),
		ReferralsQ(r),
		BalancesQ(r),
		balance,
		evTypeBeRef.Reward)
	if err != nil {
		return fmt.Errorf("do claim event updates for be_referred: %w", err)
	}

	return nil
}

// addEventForReferrer adds a friend event for the referrer. If the event is
// inactive, then nothing happens. If active, the fulfilled event is added and,
// if possible, the event claimed.
func addEventForReferrer(r *http.Request, balance data.Balance) error {
	evTypeRef := EventTypes(r).Get(models.TypeReferralSpecific, evtypes.FilterInactive)
	if evTypeRef == nil {
		Log(r).Debug("Referral specific event type is inactive")
		return nil
	}

	if balance.ReferredBy == nil {
		Log(r).Debug("Balance is disabled, event for referrer will not be added")
		return nil
	}

	referral, err := ReferralsQ(r).Get(*balance.ReferredBy)
	if err != nil {
		return fmt.Errorf("get referral by ID: %w", err)
	}
	if referral == nil {
		return fmt.Errorf("critical: referred_by not null, but row in referrals absent")
	}

	refBalance, err := BalancesQ(r).FilterByNullifier(referral.Nullifier).Get()
	if err != nil {
		return fmt.Errorf("failed to get referrer balance: %w", err)
	}
	if refBalance == nil {
		return fmt.Errorf("critical: referrer balance not exist [%s], while referral code exist", referral.Nullifier)
	}

	if refBalance.ReferredBy == nil {
		Log(r).Debug("Referrer is genesis balance, referee will not be rewarded")
		return nil
	}

	if !evTypeRef.AutoClaim || !refBalance.IsVerified {
		if !refBalance.IsVerified {
			Log(r).Debug("Referrer has not scanned passport yet, adding fulfilled events")
		}
		err = EventsQ(r).Insert(data.Event{
			Nullifier: referral.Nullifier,
			Type:      evTypeRef.Name,
			Status:    data.EventFulfilled,
			Meta:      data.Jsonb(fmt.Sprintf(`{"nullifier": "%s"}`, balance.Nullifier)),
		})
		if err != nil {
			return fmt.Errorf("failed to insert fulfilled event for referrer: %w", err)
		}

		return nil
	}

	err = EventsQ(r).Insert(data.Event{
		Nullifier:    referral.Nullifier,
		Type:         evTypeRef.Name,
		Status:       data.EventClaimed,
		PointsAmount: &evTypeRef.Reward,
		Meta:         data.Jsonb(fmt.Sprintf(`{"nullifier": "%s"}`, balance.Nullifier)),
	})
	if err != nil {
		return fmt.Errorf("failed to insert claimed event for referrer: %w", err)
	}

	err = DoClaimEventUpdates(
		Levels(r),
		ReferralsQ(r),
		BalancesQ(r),
		*refBalance,
		evTypeRef.Reward)
	if err != nil {
		return fmt.Errorf("failed to do claim event updates for referrer referral specific events: %w", err)
	}

	return nil
}

func mustHexToInt(s string) string {
	return new(big.Int).SetBytes(hexutil.MustDecode(s)).String()
}
