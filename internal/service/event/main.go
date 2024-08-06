package event

import (
	"fmt"

	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/handlers"
)

func Run(cfg config.Config, date int) error {
	log := cfg.Log()
	db := cfg.DB()
	lvls := cfg.Levels()
	evTypes := cfg.EventTypes()

	balancesQ := pg.NewBalances(db)
	eventsQ := pg.NewEvents(db)
	referralsQ := pg.NewReferrals(db)

	balances, err := balancesQ.FilterByCreatedBefore(date).FilterVerified().Select()

	if err != nil {
		log.WithError(err).Error("failed to filter by updated before")
		return err
	}

	if len(balances) == 0 {
		log.Infof("no balances found")
		return nil
	}

	existingEvents, err := eventsQ.FilterByType(models.TypeEarlyTest).Select()
	if err != nil {
		log.WithError(err).Error("failed to filter eventsQ")
		return err
	}

	existingEventsMap := make(map[string]data.Event)
	for _, event := range existingEvents {
		existingEventsMap[event.Nullifier] = event
	}

	for _, balance := range balances {
		if _, exists := existingEventsMap[balance.Nullifier]; exists {
			continue
		}
		err = eventsQ.Insert(data.Event{
			Nullifier: balance.Nullifier,
			Type:      models.TypeEarlyTest,
			Status:    data.EventFulfilled,
		})
		if err != nil {
			return fmt.Errorf("failed to insert `early_test` event: %w", err)
		}
		err = autoClaimEventsForBalance(eventsQ, balancesQ, &balance, lvls, referralsQ, *evTypes)
		if err != nil {
			return fmt.Errorf("failed to auto-claim eventsQ: %w", err)
		}
	}

	return nil
}

func autoClaimEventsForBalance(
	eventsQ data.EventsQ, balanceQ data.BalancesQ, balance *data.Balance,
	lvls *config.Levels, referralsQ data.ReferralsQ, evTypes evtypes.Types) error {

	var totalPoints int64
	eventsToClaim, err := eventsQ.FilterByStatus(data.EventFulfilled).Select()
	if err != nil {
		return fmt.Errorf("failed to select events for user=%s: %w", balance.Nullifier, err)
	}

	eventsMap := map[string][]string{}
	for _, e := range eventsToClaim {
		eventsMap[e.Type] = append(eventsMap[e.Type], e.ID)
	}

	for evName, evIDs := range eventsMap {
		evType := evTypes.Get(evName, evtypes.FilterInactive, evtypes.FilterByAutoClaim(true))
		if evType == nil {
			continue
		}

		_, err = eventsQ.FilterByID(evIDs...).Update(data.EventClaimed, nil, &evType.Reward)
		if err != nil {
			return fmt.Errorf("failedt to update %s events for user=%s: %w", evName, balance.Nullifier, err)
		}

		totalPoints += evType.Reward * int64(len(evIDs))
	}

	level, err := handlers.DoLevelRefUpgrade(lvls, referralsQ, balance, totalPoints)
	if err != nil {
		return fmt.Errorf("failed to do lvlup and referrals updates: %w", err)
	}

	err = balanceQ.FilterByNullifier(balance.Nullifier).Update(map[string]any{
		data.ColAmount: pg.AddToValue(data.ColAmount, totalPoints),
		data.ColLevel:  level,
	})
	if err != nil {
		return fmt.Errorf("error update balance amount and level: %w", err)
	}

	return nil
}
