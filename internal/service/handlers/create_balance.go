package handlers

import (
	"fmt"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateBalance(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewCreateBalance(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nullifier := req.Data.ID

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	balance, err := BalancesQ(r).FilterByNullifier(nullifier).Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if balance != nil {
		ape.RenderErr(w, problems.Conflict())
		return
	}

	referral, err := ReferralsQ(r).Get(req.Data.Attributes.ReferredBy) // infinite referrals allowed
	if err != nil {
		Log(r).WithError(err).Error("Failed to get referral by ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if referral == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	events := prepareEventsWithRef(nullifier, req.Data.Attributes.ReferredBy, r)
	if err = createBalanceWithEventsAndReferrals(nullifier, &req.Data.Attributes.ReferredBy, events, r); err != nil {
		Log(r).WithError(err).Error("Failed to create balance with events")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// We can't return inserted balance in a single query, because we can't calculate
	// rank in transaction: RANK() is a window function allowed on a set of rows,
	// while with RETURNING we operate a single one.
	// Balance will exist cause of previous logic.
	balance, err = BalancesQ(r).GetWithRank(nullifier)
	if err != nil {
		Log(r).WithError(err).Error("Failed to get created balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	referrals, err := ReferralsQ(r).FilterByNullifier(nullifier).Select()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get referral code by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if len(referrals) != 1 {
		Log(r).WithError(err).Error("There must be only 1 referral code")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newBalanceResponse(*balance, &referrals[0]))
}

func prepareEventsWithRef(nullifier, refBy string, r *http.Request) []data.Event {
	events := EventTypes(r).PrepareEvents(nullifier, evtypes.FilterNotOpenable)
	if refBy == "" {
		return events
	}

	refType := EventTypes(r).Get(models.TypeBeReferred, evtypes.FilterInactive)
	if refType == nil {
		Log(r).Debug("`Be referred` event is inactive, skipping it")
		return events
	}

	Log(r).WithFields(map[string]any{"nullifier": nullifier, "referred_by": refBy}).
		Debug("`Be referred` event will be added for referee user")

	return append(events, data.Event{
		Nullifier: nullifier,
		Type:      models.TypeBeReferred,
		Status:    data.EventFulfilled,
	})
}

// createBalanceWithEvents should be called in transaction to avoid database corruption
func createBalanceWithEvents(nullifier string, refBy *string, events []data.Event, r *http.Request) error {
	balance := data.Balance{
		Nullifier:  nullifier,
		ReferredBy: refBy,
		Level:      0,
	}

	err := BalancesQ(r).Insert(balance)
	if err != nil {
		return fmt.Errorf("add balance: %w", err)
	}

	Log(r).Debugf("%d events will be added for nullifier=%s", len(events), nullifier)
	if err = EventsQ(r).Insert(events...); err != nil {
		return fmt.Errorf("add open events: %w", err)
	}

	// not consuming referral code
	return nil
}

func createBalanceWithEventsAndReferrals(nullifier string, refBy *string, events []data.Event, r *http.Request) error {
	return EventsQ(r).Transaction(func() error {
		balance := data.Balance{
			Nullifier:  nullifier,
			ReferredBy: refBy,
			Level:      0,
		}

		err := BalancesQ(r).Insert(balance)
		if err != nil {
			return fmt.Errorf("add balance: %w", err)
		}

		Log(r).Debugf("%d events will be added for nullifier=%s", len(events), nullifier)
		if err = EventsQ(r).Insert(events...); err != nil {
			return fmt.Errorf("add open events: %w", err)
		}

		level, err := doLvlUpAndReferralsUpdate(Levels(r), ReferralsQ(r), balance, 0)
		if err != nil {
			return fmt.Errorf("failed to do lvlup and referrals update: %w", err)
		}

		err = BalancesQ(r).FilterByNullifier(balance.Nullifier).Update(map[string]any{
			data.ColLevel: level,
		})
		if err != nil {
			return fmt.Errorf("update balance amount and level: %w", err)
		}

		return nil
	})
}
