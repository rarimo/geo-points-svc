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

	evType := evTypes.Get(models.TypeEarlyTest, evtypes.FilterInactive)
	if evType == nil {
		log.Infof("Event type %s is inactive", models.TypeEarlyTest)
		return nil
	}

	balances, err := balancesQ.FilterByCreatedBefore(date).FilterVerified().Select()

	if err != nil {
		log.WithError(err).Error("failed to filter by updated before")
		return err
	}
	if len(balances) == 0 {
		log.Infof("no balances found")
		return nil
	}

	nullifiers := make([]string, 0, len(balances))

	for _, balance := range balances {
		nullifiers = append(nullifiers, balance.Nullifier)
	}

	filteredEvents, err := eventsQ.
		FilterByStatus(data.EventFulfilled).
		FilterByNullifier(nullifiers...).
		Select()

	if err != nil {
		log.WithError(err).Errorf("Failed to select %s events", err)
		return err
	}

	EventsMap := make(map[string]data.Event)
	for _, event := range filteredEvents {
		EventsMap[event.Nullifier] = event
	}

	for _, balance := range balances {
		if _, exists := EventsMap[balance.Nullifier]; exists {
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

		var totalPoints int64

		_, err = eventsQ.FilterByNullifier(nullifiers...).Update(data.EventClaimed, nil, &evType.Reward)
		if err != nil {
			return fmt.Errorf("failedt to update %s events for user=%s: %w", models.TypeEarlyTest, balance.Nullifier, err)
		}

		totalPoints += evType.Reward

		level, err := handlers.DoLevelRefUpgrade(lvls, referralsQ, &balance, totalPoints)
		if err != nil {
			return fmt.Errorf("failed to do lvlup and referrals updates: %w", err)
		}

		err = balancesQ.FilterByNullifier(balance.Nullifier).Update(map[string]any{
			data.ColAmount: pg.AddToValue(data.ColAmount, totalPoints),
			data.ColLevel:  level,
		})

		if err != nil {
			return fmt.Errorf("error update balance amount and level: %w", err)
		}
	}

	return nil
}
