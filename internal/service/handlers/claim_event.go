package handlers

import (
	"fmt"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/referralid"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func ClaimEvent(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewClaimEvent(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	event, err := EventsQ(r).FilterByID(req.Data.ID).FilterByStatus(data.EventFulfilled).Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get event by balance ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if event == nil {
		Log(r).Debugf("Event not found for id=%s status=%s", req.Data.ID, data.EventFulfilled)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(event.Nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	evType := EventTypes(r).Get(event.Type, evtypes.FilterInactive)
	if evType == nil {
		Log(r).Infof("Event type %s is inactive", event.Type)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	balance, err := BalancesQ(r).FilterByNullifier(event.Nullifier).FilterDisabled().Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if balance == nil || !balance.IsVerified {
		msg := "did not verify passport"
		if balance == nil {
			msg = "is disabled"
		}
		Log(r).Infof("Balance nullifier=%s %s", event.Nullifier, msg)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	err = EventsQ(r).Transaction(func() error {
		event, err = claimEvent(r, event, balance)
		return err
	})
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to claim event %s and accrue %d points to the balance %s",
			event.ID, evType.Reward, event.Nullifier)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// balance should exist cause of previous logic
	balance, err = BalancesQ(r).GetWithRank(event.Nullifier)
	if err != nil {
		Log(r).WithError(err).Error("Failed to get balance by nullifier with rank")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newClaimEventResponse(*event, evType.Resource(), *balance))
}

// claimEvent requires event to exist
// call in transaction to prevent unexpected changes
func claimEvent(r *http.Request, event *data.Event, balance *data.Balance) (claimed *data.Event, err error) {
	evType := EventTypes(r).Get(event.Type, evtypes.FilterInactive)
	if evType == nil {
		return event, nil
	}

	claimed, err = EventsQ(r).FilterByID(event.ID).Update(data.EventClaimed, nil, &evType.Reward)
	if err != nil {
		return nil, fmt.Errorf("update event status: %w", err)
	}

	err = DoClaimEventUpdates(
		Levels(r),
		ReferralsQ(r),
		BalancesQ(r),
		*balance,
		evType.Reward)
	if err != nil {
		return nil, fmt.Errorf("failed to do claim event updates: %w", err)
	}

	return claimed, nil
}

// DoClaimEventUpdates do updates which link to claim event:
// update reserved amount in country;
// lvlup and update referrals count;
// accruing points;
//
// Balance must be active and with verified passport
func DoClaimEventUpdates(
	levels config.Levels,
	referralsQ data.ReferralsQ,
	balancesQ data.BalancesQ,
	balance data.Balance,
	reward int64) (err error) {

	level, err := doLvlUpAndReferralsUpdate(levels, referralsQ, balance, reward)
	if err != nil {
		return fmt.Errorf("failed to do lvlup and referrals updates: %w", err)
	}

	err = balancesQ.FilterByNullifier(balance.Nullifier).Update(map[string]any{
		data.ColAmount: pg.AddToValue(data.ColAmount, reward),
		data.ColLevel:  level,
	})
	if err != nil {
		return fmt.Errorf("update balance amount and level: %w", err)
	}

	return nil
}

func doLvlUpAndReferralsUpdate(levels config.Levels, referralsQ data.ReferralsQ, balance data.Balance, reward int64) (level int, err error) {
	refsCount, level := levels.LvlUp(balance.Level, reward+balance.Amount)
	// we need +2 because refsCount can be -1
	referrals := make([]data.Referral, 0, refsCount+2)

	// count used to calculate ref code
	count, err := referralsQ.New().FilterByNullifier(balance.Nullifier).Count()
	if err != nil {
		return 0, fmt.Errorf("failed to get referral count: %w", err)
	}
	switch {
	case refsCount > 0:
		referrals = append(referrals, prepareReferralsToAdd(balance.Nullifier, uint64(refsCount), count)...)
	case refsCount == -1:
		referrals = append(referrals, data.Referral{
			ID:        referralid.New(balance.Nullifier, count),
			Nullifier: balance.Nullifier,
			Infinity:  true,
		})
	}
	if err = referralsQ.New().Insert(referrals...); err != nil {
		return 0, fmt.Errorf("failed to insert referrals: %w", err)
	}

	return level, nil
}

func newClaimEventResponse(
	event data.Event,
	meta resources.EventStaticMeta,
	balance data.Balance,
) resources.EventResponse {

	eventModel := newEventModel(event, meta)
	eventModel.Relationships = &resources.EventRelationships{
		Balance: resources.Relation{
			Data: &resources.Key{
				ID:   balance.Nullifier,
				Type: resources.BALANCE,
			},
		},
	}

	resp := resources.EventResponse{Data: eventModel}
	inc := newBalanceModel(balance)
	resp.Included.Add(&inc)

	return resp
}
