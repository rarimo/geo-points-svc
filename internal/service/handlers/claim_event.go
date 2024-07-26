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

const langHeader = "Accept-Language"

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
	if balance == nil || !balance.IsVerified() {
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

	ape.Render(w, newClaimEventResponse(*event, evType.Resource(r.Header.Get(langHeader)), *balance))
}

// claimEvent requires event to exist
// call in transaction to prevent unexpected changes
func claimEvent(r *http.Request, event *data.Event, balance *data.Balance) (*data.Event, error) {
	evType := EventTypes(r).Get(event.Type, evtypes.FilterInactive)
	if evType == nil {
		return event, nil
	}

	claimed, err := EventsQ(r).FilterByID(event.ID).Update(data.EventClaimed, nil, &evType.Reward)
	if err != nil {
		return nil, fmt.Errorf("update event status: %w", err)
	}

	if len(claimed) == 0 {
		return nil, fmt.Errorf("event wasn't updated")
	}

	err = DoClaimEventUpdates(
		Levels(r),
		ReferralsQ(r),
		BalancesQ(r),
		balance,
		evType.Reward)
	if err != nil {
		return nil, fmt.Errorf("failed to do claim event updates: %w", err)
	}

	return &claimed[0], nil
}

// DoClaimEventUpdates do updates which link to claim event:
// update reserved amount in country;
// lvlup and update referrals count;
// accruing points;
//
// Balance must be active and with verified passport
func DoClaimEventUpdates(
	levels *config.Levels,
	referralsQ data.ReferralsQ,
	balancesQ data.BalancesQ,
	balance *data.Balance,
	reward int64) (err error) {

	level, err := doLevelRefUpgrade(levels, referralsQ, balance, reward)
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

// doLevelRefUpgrade calculates new level by provided reward: if level is up,
// referrals are added
func doLevelRefUpgrade(levels *config.Levels, refQ data.ReferralsQ, balance *data.Balance, reward int64) (level int, err error) {
	refsCount, level := levels.LvlChange(balance.Level, reward+balance.Amount)
	referrals := []data.Referral{}
	// count used to calculate ref code
	dbReferrals, err := refQ.New().FilterByNullifier(balance.Nullifier).Select()
	if err != nil {
		return 0, fmt.Errorf("failed to get referral count: %w", err)
	}

	toDelete, infinityPresent := getRefIDToDelete(dbReferrals, negAbs(refsCount))

	switch {
	case refsCount == nil:
		if infinityPresent {
			break
		}
		referrals = append(referrals, data.Referral{
			ID:        referralid.New(balance.Nullifier, uint64(len(dbReferrals))),
			Nullifier: balance.Nullifier,
			Infinity:  true,
		})
	case *refsCount > 0:
		referrals = append(referrals, prepareReferralsToAdd(balance.Nullifier, uint64(*refsCount), uint64(len(dbReferrals)))...)
	case *refsCount < 0:
		if len(toDelete) == 0 {
			break
		}
		err = refQ.New().FilterByIDs(toDelete...).Delete()
		if err != nil {
			return 0, fmt.Errorf("failed to delete referrals: %w", err)
		}
		return level, nil
	}

	if err = refQ.New().Insert(referrals...); err != nil {
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

func getRefIDToDelete(referrals []data.Referral, count int) ([]string, bool) {
	var infinityRef *data.Referral
	// extra capacity for infinity ref
	ids := make([]string, 0, count+1)
	for i, ref := range referrals {
		if ref.Infinity {
			infinityRef = &referrals[i]
		}

		if ref.UsageLeft > 0 && len(ids) < count {
			ids = append(ids, ref.ID)
		}
	}

	if infinityRef != nil {
		ids = append(ids, infinityRef.ID)
	}

	return ids, infinityRef != nil
}

func negAbs(i *int) int {
	if i == nil || *i >= 0 {
		return 0
	}
	return -(*i)
}
