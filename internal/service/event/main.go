package event

import (
	"fmt"

	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
)

func Run(cfg config.Config, date int) error {
	log := cfg.Log()
	db := cfg.DB()

	balances := pg.NewBalances(db)
	events := pg.NewEvents(db)

	users, err := balances.FilterByCreatedBefore(date).FilterVerified().Select()

	if err != nil {
		log.WithError(err).Error("failed to filter by updated before")
		return err
	}

	if users == nil {
		log.Infof("no users found")
		return nil
	}

	var userNullifiers []string
	for _, user := range users {
		userNullifiers = append(userNullifiers, user.Nullifier)
	}

	existingEvents, err := events.FilterByType(models.TypeEarlyTest).FilterByNullifiers(userNullifiers).Select()
	if err != nil {
		log.WithError(err).Error("failed to filter events")
		return err
	}

	existingEventsMap := make(map[string]data.Event)
	for _, event := range existingEvents {
		existingEventsMap[event.Nullifier] = event
	}

	for _, user := range users {
		if _, exists := existingEventsMap[user.Nullifier]; !exists {
			err = events.Insert(data.Event{
				Nullifier: user.Nullifier,
				Type:      models.TypeEarlyTest,
				Status:    data.EventFulfilled,
			})
			if err != nil {
				return fmt.Errorf("failed to insert `early_test` event: %w", err)
			}
		}
	}

	return nil
}
