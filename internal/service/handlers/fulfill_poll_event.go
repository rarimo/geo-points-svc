package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"math/big"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func FulfillPollEvent(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewFulfillPollEvent(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	proof := req.Data.Attributes.Proof
	nullifier := UserClaims(r)[0].Nullifier
	proposalID, _ := new(big.Int).SetString(req.Data.Attributes.ProposalId, 10)
	proposalEventID, _ := new(big.Int).SetString(proof.PubSignals[config.PollParticipationEventID], 10)

	log := Log(r).WithFields(map[string]any{
		"nullifier":         nullifier,
		"proof":             proof,
		"proposal_id":       proposalID,
		"proposal_event_id": proposalEventID,
	})

	if !auth.Authenticates(UserClaims(r), auth.VerifiedGrant(nullifier)) ||
		new(big.Int).SetBytes(hexutil.MustDecode(nullifier)).String() != proof.PubSignals[config.PollChallengedNullifier] {
		log.Debug("failed to authenticate user")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	balance, err := BalancesQ(r).FilterByNullifier(nullifier).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if balance == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if !balance.IsVerified() || balance.IsDisabled() {
		log.Infof("Balance is forbidden to fulfill or claim: is_verified=%t, referred_by=%v",
			balance.IsVerified(), balance.ReferredBy)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	evType := EventTypes(r).Get(models.TypePollParticipation, evtypes.FilterInactive)
	if evType == nil {
		log.Infof("Event poll participation type is inactive")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	pollUserEvents, err := EventsQ(r).FilterByNullifier(nullifier).FilterByType(models.TypePollParticipation).Select()
	if err != nil {
		log.WithError(err).Error("Failed to get user poll events")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if len(pollUserEvents) != 0 {
		pollID := struct {
			PollID string `json:"poll_id"`
		}{}

		for _, event := range pollUserEvents {
			err := json.Unmarshal(event.Meta, &pollID)
			if err != nil {
				log.WithError(err).Errorf("Failed to parse event meta with eventID=%s", event.ID)
				ape.RenderErr(w, problems.InternalError())
				return
			}

			if pollID.PollID == proof.PubSignals[config.PollParticipationEventID] {
				log.Debugf("Poll event already fulfilled")
				ape.RenderErr(w, problems.Conflict())
				return
			}
		}
	}

	err = PollVerifier(r).VerifyProof(proof, proposalID, proposalEventID)
	if err != nil {
		log.WithError(err).Debug("Failed to verify poll participation proof")
		if errors.Is(err, config.ErrInvalidProposalEventID) ||
			errors.Is(err, config.ErrInvalidRoot) ||
			errors.Is(err, config.ErrInvalidChallengedEventID) {
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"proof": err,
			})...)
			return
		}

		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = EventsQ(r).Insert(data.Event{
		Nullifier: nullifier,
		Type:      models.TypePollParticipation,
		Status:    data.EventFulfilled,
		Meta:      data.Jsonb(fmt.Sprintf(`{"poll_id": "%s"}`, proposalEventID.String())),
	})
	if err != nil {
		log.WithError(err).Error("Failed to insert poll event")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if !evType.AutoClaim {
		log.Debug("Event fulfilled due to disabled auto-claim")
		ape.Render(w, newEventClaimingStateResponse(balance.Nullifier, false))
		return
	}

	err = EventsQ(r).Transaction(func() error {
		return autoClaimEventsForBalance(r, balance)
	})
	if err != nil {
		log.WithError(err).Error("Failed to autoclaim events for user")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newEventClaimingStateResponse(balance.Nullifier, true))
}
