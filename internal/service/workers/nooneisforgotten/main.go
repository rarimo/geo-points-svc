package nooneisforgotten

import (
	"errors"
	"fmt"

	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/kit/pgdb"
)

func Run(cfg config.Config, sig chan struct{}) {
	db := cfg.DB().Clone()
	if err := pg.NewEvents(db).Transaction(func() error {
		return updatePassportScanEvents(db, cfg.EventTypes(), cfg.Levels())
	}); err != nil {
		panic(fmt.Errorf("failed to update passport scan events: %w", err))
	}

	if err := pg.NewEvents(db).Transaction(func() error {
		return updateReferralUserEvents(db, cfg.EventTypes())
	}); err != nil {
		panic(fmt.Errorf("failed to update referral user events: %w", err))
	}

	if err := pg.NewEvents(db).Transaction(func() error {
		return autoClaimEvents(db, cfg.EventTypes(), cfg.Levels())
	}); err != nil {
		panic(fmt.Errorf("failed to claim referral specific events: %w", err))
	}

	sig <- struct{}{}
}

// updatePassportScanEvents is needed so that if the passport
// scan events were not fulfilled or claimed because the event was disabled,
// expired or no auto-claimed, fulfill and, if possible, claim them.
// First, there is an attempt to claim as many events as
// possible and to fulfill the rest of the events.
//
// Event will not be claimed if AutoClaim is disabled.
func updatePassportScanEvents(db *pgdb.DB, types *evtypes.Types, levels *config.Levels) error {
	evType := types.Get(models.TypePassportScan, evtypes.FilterInactive)
	if evType == nil {
		return nil
	}

	// ensured in query that all balances are verified and active
	balances, err := pg.NewBalances(db).WithoutPassportEvent()
	if err != nil {
		return fmt.Errorf("failed to select balances without points for passport scan: %w", err)
	}

	toFulfillOrClaim := make([]string, 0, len(balances))
	for _, b := range balances {
		if b.EventStatus == data.EventOpen {
			toFulfillOrClaim = append(toFulfillOrClaim, b.EventID)
		}
	}
	if len(toFulfillOrClaim) == 0 {
		return nil
	}

	status := data.EventFulfilled
	if evType.AutoClaim {
		status = data.EventClaimed
	}

	_, err = pg.NewEvents(db).
		FilterByID(toFulfillOrClaim...).
		Update(status, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to update passport scan events: %w", err)
	}
	if !evType.AutoClaim {
		return nil
	}

	for _, b := range balances {
		err = handlers.DoClaimEventUpdates(
			levels,
			pg.NewReferrals(db),
			pg.NewBalances(db),
			b.Balance,
			evType.Reward)
		if err != nil {
			return fmt.Errorf("failed to do claim event updates for passport scan: %w", err)
		}
	}

	return nil
}

// updateReferralUserEvents is used to add events for referrers
// for friends who have scanned the passport, if they have not been added.
//
// Events are not added if the event type is inactive
func updateReferralUserEvents(db *pgdb.DB, types *evtypes.Types) error {
	evTypeRef := types.Get(models.TypeReferralSpecific, evtypes.FilterInactive)
	if evTypeRef == nil {
		return nil
	}

	refPairs, err := pg.NewBalances(db).WithoutReferralEvent()
	if err != nil {
		return fmt.Errorf("failed to select balances without points for referred users: %w", err)
	}

	toInsert := make([]data.Event, 0, len(refPairs))
	for _, ref := range refPairs {
		toInsert = append(toInsert, data.Event{
			Nullifier: ref.Referrer,
			Type:      models.TypeReferralSpecific,
			Status:    data.EventFulfilled,
			Meta:      data.Jsonb(fmt.Sprintf(`{"nullifier": "%s"}`, ref.Referred)),
		})
	}

	if len(toInsert) == 0 {
		return nil
	}

	if err = pg.NewEvents(db).Insert(toInsert...); err != nil {
		return fmt.Errorf("failed to insert referred user events: %w", err)
	}

	return nil
}

// autoClaimEvents claim fulfilled events which have auto-claim enabled. This is
// useful if some events were inactive, then became active and must be claimed
// automatically.
func autoClaimEvents(db *pgdb.DB, types *evtypes.Types, levels *config.Levels) error {
	claimTypes := types.Names(evtypes.FilterByAutoClaim(true))
	if len(claimTypes) == 0 {
		return nil
	}

	events, err := pg.NewEvents(db).
		FilterByType(claimTypes...).
		FilterByStatus(data.EventFulfilled).
		Select()
	if err != nil {
		return fmt.Errorf("failed to select fulfilled events: %w", err)
	}

	// nullifiers var is used only for selection, so we don't care about duplicates
	nullifiers := make([]string, 0, len(events))
	for _, event := range events {
		nullifiers = append(nullifiers, event.Nullifier)
	}
	if len(nullifiers) == 0 {
		return nil
	}

	balances, err := pg.NewBalances(db).FilterByNullifier(nullifiers...).Select()
	if err != nil {
		return fmt.Errorf("failed to select balances for claim passport scan event: %w", err)
	}
	if len(balances) == 0 {
		return errors.New("critical: events present, but no balances with nullifier")
	}

	// select events to claim only for verified and active balances, group by type name
	claimByTypes := make(map[string][]data.Event, len(claimTypes))
	for _, event := range events {
		for _, balance := range balances {
			if event.Nullifier != balance.Nullifier || !balance.IsVerified || balance.ReferredBy == nil {
				continue
			}
			claimByTypes[event.Type] = append(claimByTypes[event.Type], event)
			break
		}
	}
	if len(claimByTypes) == 0 {
		return nil
	}

	rewardByNullifier := make(map[string]int64, len(balances))
	for _, evType := range types.List(evtypes.FilterByAutoClaim(true)) {
		byType := claimByTypes[evType.Name]
		if len(byType) == 0 {
			continue
		}

		ids := make([]string, 0, len(byType))
		for _, ev := range byType {
			ids = append(ids, ev.ID)
			rewardByNullifier[ev.Nullifier] += evType.Reward
		}

		_, err = pg.NewEvents(db).FilterByID(ids...).Update(data.EventClaimed, nil, &evType.Reward)
		if err != nil {
			return fmt.Errorf("update event status: %w", err)
		}
	}

	for _, balance := range balances {
		err = handlers.DoClaimEventUpdates(
			levels,
			pg.NewReferrals(db),
			pg.NewBalances(db),
			balance,
			rewardByNullifier[balance.Nullifier])
		if err != nil {
			return fmt.Errorf("failed to do claim event updates for referral specific event: %w", err)
		}
	}

	return nil
}
