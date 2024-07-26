package handlers

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	zk "github.com/rarimo/zkverifier-kit"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func FulfillPollEvent(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewFulfillPollEvent(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	event, err := EventsQ(r).FilterByID(req.Data.ID).FilterByStatus(data.EventOpen).Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get event by ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if event == nil {
		Log(r).Debugf("Event not found for id=%s status=%s", req.Data.ID, data.EventOpen)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(event.Nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	log := Log(r).WithFields(map[string]any{
		"event.nullifier": req.Data.ID,
		"event.id":        event.ID,
		"event.type":      event.Type,
	})

	balance, err := BalancesQ(r).FilterByNullifier(event.Nullifier).Get()
	if err != nil || balance == nil { // must never be nil due to foreign key constraint
		log.WithError(err).Error("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if !BalanceIsVerified(balance) || balance.ReferredBy == nil {
		log.Infof("Balance is forbidden to fulfill or claim: is_verified=%t, referred_by=%v",
			BalanceIsVerified(balance), balance.ReferredBy)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	evType, err := EventTypesQ(r).Get(event.Type)
	if err != nil {
		log.WithError(err).Error("Failed to get event type from DB")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if evType == nil || evType.PollEventID == nil || evtypes.FilterInactive(*evType) {
		log.Infof("Event type %s is not poll kind or is inactive", event.Type)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	proof := req.Data.Attributes.Proof
	ni := zk.Indexes(zk.PollParticipation)[zk.Nullifier]
	proof.PubSignals[ni] = mustHexToInt(balance.Nullifier)

	err = Verifiers(r).Poll.VerifyProof(proof,
		zk.WithPollParticipationEventID(*evType.PollEventID),
		// contract must be not nil when event ID is present
		zk.WithPollRootVerifier(Verifiers(r).PollRoot.WithContract(*evType.PollContract)),
	)

	if err != nil {
		var vErr validation.Errors
		if !errors.As(err, &vErr) {
			log.WithError(err).Error("Failed to verify proof")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		log.WithError(err).Info("Invalid proof")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	log.Debug("Poll participation proof successfully verified")

	if !evType.AutoClaim {
		_, err = EventsQ(r).FilterByID(event.ID).Update(data.EventFulfilled, nil, nil)
		if err != nil {
			log.WithError(err).Error("Failed to update event status")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		log.Debug("Event fulfilled due to disabled auto-claim")
		ape.Render(w, newEventClaimingStateResponse(balance.Nullifier, false))
		return
	}

	err = EventsQ(r).Transaction(func() error {
		event, err = claimEvent(r, event, balance)
		return err
	})
	if err != nil {
		log.WithError(err).Errorf("Failed to claim event %s and accrue %d points to the balance %s",
			event.ID, evType.Reward, event.Nullifier)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newEventClaimingStateResponse(balance.Nullifier, true))
}
