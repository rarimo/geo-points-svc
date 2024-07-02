package nooneisforgotten

import (
	"errors"
	"fmt"

	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
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
		panic(fmt.Errorf("failed to update referral user events"))
	}

	if err := pg.NewEvents(db).Transaction(func() error {
		return claimReferralSpecificEvents(db, cfg.EventTypes(), cfg.Levels())
	}); err != nil {
		panic(fmt.Errorf("failed to claim referral specific events: %w", err))
	}

	sig <- struct{}{}
}

// updatePassportScanEvents is needed so that if the passport
// scan events were not fulfilled or claimed because the event was disabled,
// expired or no autoclaimed, fulfill and, if possible, claim them.
// First, there is an attempt to claim as many events as
// possible and to fulfill the rest of the events.
//
// Event will not be claimed if AutoClaim is disabled.
func updatePassportScanEvents(db *pgdb.DB, types evtypes.Types, levels config.Levels) error {
	evType := types.Get(evtypes.TypePassportScan, evtypes.FilterInactive)
	if evType == nil {
		return nil
	}

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
func updateReferralUserEvents(db *pgdb.DB, types evtypes.Types) error {
	evTypeRef := types.Get(evtypes.TypeReferralSpecific, evtypes.FilterInactive)
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
			Type:      evtypes.TypeReferralSpecific,
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

// claimReferralSpecificEvents claim fulfilled events for invited
// friends which have passport scanned, if it possible
func claimReferralSpecificEvents(db *pgdb.DB, types evtypes.Types, levels config.Levels) error {
	evType := types.Get(evtypes.TypeReferralSpecific, evtypes.FilterInactive)
	if evType == nil || !evType.AutoClaim {
		return nil
	}

	events, err := pg.NewEvents(db).
		FilterByType(evtypes.TypeReferralSpecific).
		FilterByStatus(data.EventFulfilled).
		Select()
	if err != nil {
		return fmt.Errorf("failed to select passport scan events: %w", err)
	}

	// we need to have maps which link nullifiers to events slice
	toClaim := make([]string, 0, len(events))
	nullifiers := make([]string, 0, len(events))
	for _, event := range events {
		toClaim = append(toClaim, event.ID)
		nullifiers = append(nullifiers, event.Nullifier)
	}
	if len(toClaim) == 0 {
		return nil
	}

	_, err = pg.NewEvents(db).FilterByID(toClaim...).Update(data.EventClaimed, nil, &evType.Reward)
	if err != nil {
		return fmt.Errorf("update event status: %w", err)
	}

	balances, err := pg.NewBalances(db).FilterByNullifier(nullifiers...).Select()
	if err != nil {
		return fmt.Errorf("failed to select balances for claim passport scan event: %w", err)
	}
	if len(balances) == 0 {
		return errors.New("critical: events present, but no balances with nullifier")
	}

	for _, balance := range balances {
		err = handlers.DoClaimEventUpdates(
			levels,
			pg.NewReferrals(db),
			pg.NewBalances(db),
			balance,
			evType.Reward)
		if err != nil {
			return fmt.Errorf("failed to do claim event updates for referral specific event: %w", err)
		}
	}

	return nil
}
