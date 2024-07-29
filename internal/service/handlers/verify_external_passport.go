package handlers

import (
	"fmt"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func VerifyExternalPassport(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVerifyExternalPassport(r)
	if err != nil {
		Log(r).WithError(err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var (
		nullifier   = req.Data.ID
		externalAID = req.Data.Attributes.AnonymousId
		sharedHash  = *req.Data.Attributes.SharedHash

		gotSig = r.Header.Get("Signature")
	)
	log := Log(r).WithFields(map[string]any{
		"nullifier":   nullifier,
		"externalAID": externalAID,
		"sharedHash":  sharedHash,
	})

	wantSig, err := SigCalculator(r).PassportVerificationSignature(req.Data.ID, externalAID)
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

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	byNullifier, err := BalancesQ(r).FilterByNullifier(nullifier).Get()
	if err != nil {
		log.WithError(err).Errorf("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if byNullifier == nil {
		log.Error("Balance absent")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if byNullifier.ExternalAID != nil {
		log.Debug("Already verified")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	byExternalAID, err := BalancesQ(r).FilterByExternalAID(externalAID).Get()
	if err != nil {
		log.WithError(err).Errorf("Failed to get balance by external AID")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if byExternalAID != nil {
		log.Debug("External AID already used")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	if byNullifier.SharedHash == nil || *byNullifier.SharedHash != sharedHash {
		if byNullifier.SharedHash != nil {
			log.Debug("Shared hash already used")
			ape.RenderErr(w, problems.Conflict())
			return
		}

		bySharedHash, err := BalancesQ(r).FilterBySharedHash(sharedHash).Get()
		if err != nil {
			log.WithError(err).Errorf("Failed to get balance by shared hash")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if bySharedHash != nil {
			log.Debug("Shared hash already used")
			ape.RenderErr(w, problems.Conflict())
			return
		}
	}

	err = EventsQ(r).Transaction(func() error {
		if err = updateBalanceVerification(r, byNullifier, externalAID, data.VerifyExternalType, &sharedHash); err != nil {
			return fmt.Errorf("update balance verification info: %w", err)
		}

		externalPassportEvent := EventTypes(r).Get(models.TypeExternalPassportScan, evtypes.FilterInactive)
		if externalPassportEvent != nil {
			_, err := EventsQ(r).FilterByNullifier(nullifier).
				FilterByStatus(data.EventOpen).
				FilterByType(models.TypeExternalPassportScan).
				Update(data.EventFulfilled, nil, nil)
			if err != nil {
				return fmt.Errorf("failed to fulfill external passport scan event: %w", err)
			}
		}

		if byNullifier.IsDisabled() {
			log.Debug("Balance is disabled, events will not be claimed")
			return nil
		}

		byNullifier.ExternalAID = &externalAID

		if err = autoClaimEventsForBalance(r, byNullifier); err != nil {
			return fmt.Errorf("failed to autoclaim events for user")
		}
		if byNullifier.IsVerified() {
			return nil
		}

		if err := addEventForReferrer(r, byNullifier); err != nil {
			return fmt.Errorf("add event for referrer: %w", err)
		}
		return nil

	})
	if err != nil {
		log.WithError(err).Error("Failed to do passport scan updates")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	event, err := EventsQ(r).FilterByNullifier(nullifier).
		FilterByType(models.TypeExternalPassportScan).
		FilterByStatus(data.EventClaimed).
		Get()
	if err != nil {
		log.WithError(err).Error("Failed to get claimed event")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newEventClaimingStateResponse(req.Data.ID, event != nil))

}

func autoClaimEventsForBalance(r *http.Request, balance *data.Balance) error {
	if balance == nil {
		Log(r).Debug("Balance absent. Events not claimed.")
		return nil
	}

	if balance.IsDisabled() || !balance.IsVerified() {
		Log(r).Debug("User not eligible for event claiming. Events not claimed.")
		return nil
	}

	var totalPoints int64
	eventsToClaim, err := EventsQ(r).
		FilterByNullifier(balance.Nullifier).
		FilterByStatus(data.EventFulfilled).
		Select()
	if err != nil {
		return fmt.Errorf("failed to select events for user=%s: %w", balance.Nullifier, err)
	}

	eventsMap := map[string][]string{}
	for _, e := range eventsToClaim {
		if _, ok := eventsMap[e.Type]; !ok {
			eventsMap[e.Type] = []string{}
		}
		eventsMap[e.Type] = append(eventsMap[e.Type], e.ID)
	}

	for evName, evIDs := range eventsMap {
		evType := EventTypes(r).Get(evName, evtypes.FilterInactive, evtypes.FilterByAutoClaim(true))
		if evType == nil {
			continue
		}

		_, err = EventsQ(r).FilterByID(evIDs...).Update(data.EventClaimed, nil, &evType.Reward)
		if err != nil {
			return fmt.Errorf("failedt to update %s events for user=%s: %w", evName, balance.Nullifier, err)
		}

		totalPoints += evType.Reward * int64(len(evIDs))
	}

	level, err := doLevelRefUpgrade(Levels(r), ReferralsQ(r), balance, totalPoints)
	if err != nil {
		return fmt.Errorf("failed to do lvlup and referrals updates: %w", err)
	}

	err = BalancesQ(r).FilterByNullifier(balance.Nullifier).Update(map[string]any{
		data.ColAmount: pg.AddToValue(data.ColAmount, totalPoints),
		data.ColLevel:  level,
	})
	if err != nil {
		return fmt.Errorf("update balance amount and level: %w", err)
	}

	return nil
}
