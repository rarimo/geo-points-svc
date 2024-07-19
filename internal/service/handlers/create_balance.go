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

	var (
		refCode      = req.Data.Attributes.ReferredBy
		isGenesisRef = false
		log          = Log(r).WithField("nullifier", nullifier)
	)

	if refCode != nil {
		log.Debug("Balance will be activated with referral code")
		referral, err := ReferralsQ(r).FilterInactive().Get(*refCode)
		if err != nil {
			log.WithError(err).Error("Failed to get referral by ID")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if referral == nil {
			ape.RenderErr(w, problems.NotFound())
			return
		}

		refBalance, err := BalancesQ(r).FilterByNullifier(referral.Nullifier).Get()
		if err != nil || refBalance == nil { // must exist due to FK constraint
			log.WithError(err).Error("Failed to get referrer balance by nullifier")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		isGenesisRef = refBalance.ReferredBy == nil
	}

	events := prepareEventsWithRef(nullifier, refCode, isGenesisRef, r)
	if refCode == nil {
		balance, err = createBalanceWithEvents(nullifier, events, r)
		if err != nil {
			log.WithError(err).Error("Failed to create disabled balance with events")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		log.Debug("Created disabled balance with events")
		ape.Render(w, newBalanceResponse(*balance, nil, 0, 0))
		return
	}

	err = createBalanceWithEventsAndReferrals(nullifier, *refCode, events, r)
	if err != nil {
		log.WithError(err).Error("Failed to create balance with events and referrals")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// We can't return inserted balance in a single query, because we can't calculate
	// rank in transaction: RANK() is a window function allowed on a set of rows,
	// while with RETURNING we operate a single one.
	// Balance will exist cause of previous logic.
	balance, err = BalancesQ(r).GetWithRank(nullifier)
	if err != nil {
		log.WithError(err).Error("Failed to get created balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	referrals, err := ReferralsQ(r).
		FilterByNullifier(nullifier).
		WithStatus().
		Select()
	if err != nil {
		log.WithError(err).Error("Failed to get referrals by nullifier with rewarding field")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newBalanceResponse(*balance, referrals, 0, 0))
}

func prepareEventsWithRef(nullifier string, refBy *string, isGenesisRef bool, r *http.Request) []data.Event {
	events := EventTypes(r).PrepareEvents(nullifier, evtypes.FilterNotOpenable)
	refType := EventTypes(r).Get(models.TypeBeReferred, evtypes.FilterInactive)

	if refBy == nil || isGenesisRef || refType == nil {
		return events
	}

	Log(r).WithFields(map[string]any{"nullifier": nullifier, "referred_by": *refBy}).
		Debug("`Be referred` event will be added for referee user")

	return append(events, data.Event{
		Nullifier: nullifier,
		Type:      models.TypeBeReferred,
		Status:    data.EventFulfilled,
	})
}

// createBalanceWithEvents should be called in transaction to avoid database corruption
func createBalanceWithEvents(nullifier string, events []data.Event, r *http.Request) (*data.Balance, error) {
	balance := data.Balance{
		Nullifier: nullifier,
		Level:     0,
	}

	if err := BalancesQ(r).Insert(balance); err != nil {
		return nil, fmt.Errorf("add balance: %w", err)
	}

	Log(r).Debugf("%d events will be added for nullifier=%s", len(events), nullifier)
	if err := EventsQ(r).Insert(events...); err != nil {
		return nil, fmt.Errorf("add open events: %w", err)
	}

	return &balance, nil
}

func createBalanceWithEventsAndReferrals(nullifier, refBy string, events []data.Event, r *http.Request) error {
	return EventsQ(r).Transaction(func() error {
		balance, err := createBalanceWithEvents(nullifier, events, r)
		if err != nil {
			return fmt.Errorf("create balance with events: %w", err)
		}
		balance.ReferredBy = &refBy

		level, err := doLevelRefUpgrade(Levels(r), ReferralsQ(r), *balance, 0)
		if err != nil {
			return fmt.Errorf("failed to do lvlup and referrals update: %w", err)
		}

		err = BalancesQ(r).FilterByNullifier(balance.Nullifier).Update(map[string]any{
			data.ColLevel:      level,
			data.ColReferredBy: refBy,
		})
		if err != nil {
			return fmt.Errorf("update balance amount and level: %w", err)
		}

		if err = ReferralsQ(r).Consume(refBy); err != nil {
			return fmt.Errorf("failed to consume referral: %w", err)
		}

		return nil
	})
}
