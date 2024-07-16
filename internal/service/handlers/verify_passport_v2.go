package handlers

import (
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

// VerifyPassportV2
// TODO: Use this handler for passport verification
// when backwards compatibility becomes unnecessary
func VerifyPassportV2(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVerifyPassportV2(r)
	if err != nil {
		Log(r).WithError(err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var (
		nullifier   = req.Data.ID
		anonymousID = req.Data.Attributes.AnonymousId
		log         = Log(r).WithFields(map[string]any{
			"balance.nullifier":    nullifier,
			"balance.anonymous_id": anonymousID,
		})

		gotSig = r.Header.Get("Signature")
	)

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

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

	byNullifier, err := BalancesQ(r).FilterByNullifier(nullifier).FilterDisabled().Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if byNullifier == nil {
		Log(r).Debugf("Balance not found: nullifier=%s", nullifier)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if byNullifier.SharedHash != nil {
		Log(r).Debugf("Already verified: nullifier=%s, AID=%s", nullifier, anonymousID)
		ape.RenderErr(w, problems.Conflict())
		return
	}

	byAnonymousID, err := BalancesQ(r).FilterByAnonymousID(anonymousID).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get balance by anonymous ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if byAnonymousID != nil && byAnonymousID.Nullifier != byNullifier.Nullifier {
		Log(r).Debugf("AnonymousID already used: nullifier=%s, AID=%s, AIDBalance=%+v", nullifier, anonymousID, byAnonymousID)
		ape.RenderErr(w, problems.Conflict())
		return
	}

	// UserClaims(r)[0] will not panic because of authorization validation
	sharedHash := UserClaims(r)[0].SharedHash
	if sharedHash == nil {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"shared_hash": fmt.Errorf("not provided in JWT"),
		})...)
		return
	}

	err = EventsQ(r).Transaction(func() error {
		err = BalancesQ(r).FilterByNullifier(byNullifier.Nullifier).Update(map[string]any{
			data.ColSharedHash:  *sharedHash,
			data.ColAnonymousID: anonymousID,
			data.ColIsVerified:  true,
		})
		if err != nil {
			return fmt.Errorf("failed to update balance: %w", err)
		}

		return doPassportScanUpdates(r, *byNullifier, anonymousID, sharedHash)
	})
	if err != nil {
		log.WithError(err).Error("Failed to execute transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	event, err := EventsQ(r).FilterByNullifier(byNullifier.Nullifier).
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
