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

// VerifyInternalPassport handler processes 3 different flows:
//   - Old passport verification with proof for the current release
//   - New passport verification with JWT for the future
//   - Legacy joining program logic when the client fails to generate query proof
func VerifyInternalPassport(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVerifyInternalPassport(r)
	if err != nil {
		Log(r).WithError(err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var (
		internalAID = req.Data.Attributes.AnonymousId
		proof       = req.Data.Attributes.Proof
		log         = Log(r).WithFields(map[string]any{
			"balance.nullifier":    req.Data.ID,
			"balance.internal_aid": internalAID,
		})
		gotSig = r.Header.Get("Signature")
	)

	wantSig, err := SigCalculator(r).PassportVerificationSignature(req.Data.ID, internalAID)
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

	byAnonymousID, err := BalancesQ(r).FilterByInternalAID(internalAID).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get balance by internal AID")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if byAnonymousID != nil && byAnonymousID.Nullifier != balance.Nullifier {
		log.Warn("Balance with the same internal AID already exists")
		ape.RenderErr(w, problems.Forbidden())
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

	if sharedHash == nil {
		empty := ""
		sharedHash = &empty
	}

	bySharedHash, err := BalancesQ(r).FilterBySharedHash(*sharedHash).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get balance by shared hash")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if bySharedHash != nil && bySharedHash.Nullifier != balance.Nullifier {
		if proof == nil {
			log.Warn("Balance with the same shared hash already exists")
			ape.RenderErr(w, problems.Forbidden())
			return
		}

		// transfer bySharedHash balance to current balance
		// because we have proof that user is registered on smart contract
		// it is main user account, so just remove shared hash and external id from unverified balance
		// it allows user to attach external passport to his verified balance
		err = BalancesQ(r).FilterByNullifier(bySharedHash.Nullifier).Update(map[string]any{
			data.ColSharedHash:  nil,
			data.ColExternalAID: nil,
		})
		if err != nil {
			log.WithError(err).Error("Failed to remove shared hash from unverified balance")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	if byAnonymousID != nil {
		if balance.SharedHash != nil {
			log.Warnf("Balance %s already verified", balance.Nullifier)
			ape.RenderErr(w, problems.Conflict())
			return
		}
		if proof == nil && *sharedHash == "" {
			log.Warnf("Balance %s tried to re-join program", balance.Nullifier)
			ape.RenderErr(w, problems.Conflict())
			return
		}

		var balAID string
		if balance.InternalAID != nil {
			balAID = *balance.InternalAID
		}

		err = validation.Errors{
			"data/attributes/anonymous_id": validation.Validate(internalAID, validation.Required, validation.In(balAID)),
		}.Filter()
		if err != nil {
			log.Warnf("Anonymous ID was changed, got %s, want %s", internalAID, balAID)
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

		ape.Render(w, newEventClaimingStateResponse(req.Data.ID, false, 0))
		return
	}

	err = EventsQ(r).Transaction(func() error {
		if err = updateBalanceVerification(r, balance, internalAID, data.VerifyInternalType, sharedHash); err != nil {
			return fmt.Errorf("update balance verification info: %w", err)
		}

		passportScanEvent := EventTypes(r).Get(models.TypePassportScan, evtypes.FilterInactive)
		if passportScanEvent != nil {
			_, err := EventsQ(r).FilterByNullifier(balance.Nullifier).
				FilterByStatus(data.EventOpen).
				FilterByType(models.TypePassportScan).
				Update(data.EventFulfilled, nil, nil)
			if err != nil {
				return fmt.Errorf("failed to fulfill external passport scan event: %w", err)
			}
		}

		if balance.IsDisabled() {
			log.Debug("Balance is disabled, events will not be claimed")
			return nil
		}

		balance.InternalAID = &internalAID

		if err = autoClaimEventsForBalance(r, balance); err != nil {
			return fmt.Errorf("failed to autoclaim events for user")
		}

		if err := addEventForReferrer(r, balance); err != nil {
			return fmt.Errorf("add event for referrer: %w", err)
		}
		return nil
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

	ape.Render(w, newEventClaimingStateResponse(req.Data.ID, event != nil, 0))
}

func newEventClaimingStateResponse(id string, isClaimed bool, reward int64) resources.EventClaimingStateResponse {
	var res resources.EventClaimingStateResponse
	res.Data.ID = id
	res.Data.Type = resources.EVENT_CLAIMING_STATE
	res.Data.Attributes.Claimed = isClaimed
	res.Data.Attributes.Reward = reward
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

// balance must be not nil
func updateBalanceVerification(r *http.Request, balance *data.Balance, anonymousID, verifyType string, sharedHash *string) error {
	var toUpd map[string]any
	switch verifyType {
	case data.VerifyInternalType:
		toUpd = map[string]any{
			data.ColInternalAID: anonymousID,
		}
	case data.VerifyExternalType:
		toUpd = map[string]any{
			data.ColExternalAID: anonymousID,
		}
	default:
		return fmt.Errorf("invalid verify type: want %s or %s, got %s", data.VerifyInternalType, data.VerifyExternalType, verifyType)
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

// addEventForReferrer adds a friend event for the referrer. If the event is
// inactive, then nothing happens. If active, the fulfilled event is added and,
// if possible, the event claimed.
// balance must be not nil
func addEventForReferrer(r *http.Request, balance *data.Balance) error {
	evTypeRef := EventTypes(r).Get(models.TypeReferralSpecific, evtypes.FilterInactive)
	if evTypeRef == nil {
		Log(r).Debug("Referral specific event type is inactive")
		return nil
	}

	if balance.IsDisabled() {
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

	if refBalance.IsDisabled() {
		Log(r).Debug("Referrer is genesis balance, referee will not be rewarded")
		return nil
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

	err = autoClaimEventsForBalance(r, refBalance)
	if err != nil {
		return fmt.Errorf("failed to autoclaim events for referrer: %w", err)
	}

	return nil
}

func mustHexToInt(s string) string {
	return new(big.Int).SetBytes(hexutil.MustDecode(s)).String()
}
