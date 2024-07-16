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

func ActivateBalance(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewActivateBalance(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nullifier := req.Data.ID
	referralCode := req.Data.Attributes.ReferredBy

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

	if balance == nil {
		Log(r).Debugf("Balance %s not exist", nullifier)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if balance.ReferredBy != nil {
		Log(r).Infof("Balance already activated with code '%s'", *balance.ReferredBy)
		ape.RenderErr(w, problems.Conflict())
		return
	}

	referral, err := ReferralsQ(r).FilterInactive().Get(req.Data.Attributes.ReferredBy)
	if err != nil {
		Log(r).WithError(err).Error("Failed to get referral by ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if referral == nil {
		Log(r).Debugf("Referral code '%s' not found", referralCode)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	refBalance, err := BalancesQ(r).FilterByNullifier(referral.Nullifier).Get()
	if err != nil || refBalance == nil { // must exist due to FK constraint
		Log(r).WithError(err).Error("Failed to get referrer balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = EventsQ(r).Transaction(func() error {
		level, err := doLvlUpAndReferralsUpdate(Levels(r), ReferralsQ(r), *balance, 0)
		if err != nil {
			return fmt.Errorf("failed to do lvlup and referrals update: %w", err)
		}

		err = BalancesQ(r).FilterByNullifier(balance.Nullifier).Update(map[string]any{
			data.ColReferredBy: referralCode,
			data.ColLevel:      level,
		})
		if err != nil {
			return fmt.Errorf("update balance amount and level: %w", err)
		}

		if refBalance.ReferredBy != nil {
			err = EventsQ(r).Insert(data.Event{
				Nullifier: nullifier,
				Type:      models.TypeBeReferred,
				Status:    data.EventFulfilled,
			})
			if err != nil {
				return fmt.Errorf("failed to insert `be_referred` event: %w", err)
			}
		}

		if err = ReferralsQ(r).Consume(referralCode); err != nil {
			return fmt.Errorf("failed to consume referral: %w", err)
		}

		evTypeRef := EventTypes(r).Get(models.TypeReferralSpecific, evtypes.FilterInactive)
		if evTypeRef == nil {
			Log(r).Debug("Referral specific event type is inactive")
			return nil
		}

		if balance.IsVerified {
			// Be referred event is a welcome bonus when you created balance with non-genesis referral code
			if err = claimBeReferredEvent(r, *balance); err != nil {
				return fmt.Errorf("failed to claim be referred event: %w", err)
			}
		}

		// Adds a friend event for the referrer. If the event
		// is inactive, then nothing happens. If active, the
		// fulfilled event is added and, if possible, the event claimed
		if err = addEventForReferrer(r, evTypeRef, *balance); err != nil {
			return fmt.Errorf("add event for referrer: %w", err)
		}

		return nil
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to insert events and consume referral for balance")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// We can't return inserted balance in a single query, because we can't calculate
	// rank in transaction: RANK() is a window function allowed on a set of rows,
	// while with RETURNING we operate a single one.
	// Balance will exist cause of previous logic.
	if balance, err = BalancesQ(r).GetWithRank(nullifier); err != nil {
		Log(r).WithError(err).Error("Failed to get created balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	referrals, err := ReferralsQ(r).
		FilterByNullifier(nullifier).
		WithStatus().
		Select()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get referrals by nullifier with rewarding field")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newBalanceResponse(*balance, referrals, 0, 0))
}
