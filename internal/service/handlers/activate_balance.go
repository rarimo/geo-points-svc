package handlers

import (
	"fmt"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
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

	var (
		nullifier    = req.Data.ID
		referralCode = req.Data.Attributes.ReferredBy
		log          = Log(r).WithField("nullifier", nullifier)
	)

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
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
		log.Debug("Balance not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if !balance.IsDisabled() {
		log.Infof("Balance already activated with code %s", *balance.ReferredBy)
		ape.RenderErr(w, problems.Conflict())
		return
	}

	referral, err := ReferralsQ(r).FilterInactive().Get(referralCode)
	if err != nil {
		log.WithError(err).Error("Failed to get referral by ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if referral == nil {
		log.Debugf("Referral code %s not found", referralCode)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	refBalance, err := BalancesQ(r).FilterByNullifier(referral.Nullifier).Get()
	if err != nil || refBalance == nil { // must exist due to FK constraint
		log.WithError(err).Error("Failed to get referrer balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = EventsQ(r).Transaction(func() error {
		level, err := doLevelRefUpgrade(Levels(r), ReferralsQ(r), balance, 0)
		if err != nil {
			return fmt.Errorf("failed to do lvlup and referrals update: %w", err)
		}

		err = BalancesQ(r).FilterByNullifier(balance.Nullifier).Update(map[string]any{
			data.ColReferredBy: referralCode,
			data.ColLevel:      level,
		})
		if err != nil {
			return fmt.Errorf("update balance: %w", err)
		}

		if !refBalance.IsDisabled() {
			log.Debug("Be referred event will be fulfilled for referee")
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

		if !balance.IsVerified() {
			log.Debug("Balance is not verified, events will not be claimed")
			return nil
		}

		balance.ReferredBy = &referral.ID

		if err := addEventForReferrer(r, balance); err != nil {
			return fmt.Errorf("add event for referrer: %w", err)
		}

		return autoClaimEventsForBalance(r, balance)
	})
	if err != nil {
		log.WithError(err).Error("Failed to insert events and consume referral for balance")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// We can't return inserted balance in a single query, because we can't calculate
	// rank in transaction: RANK() is a window function allowed on a set of rows,
	// while with RETURNING we operate a single one.
	// Balance will exist cause of previous logic.
	if balance, err = BalancesQ(r).GetWithRank(nullifier); err != nil {
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
