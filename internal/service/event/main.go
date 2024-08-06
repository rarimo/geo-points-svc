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
	filtEvents := events.FilterByType(models.TypeEarlyTest)

	if err != nil {
		log.WithError(err).Error("failed to filter by updated before")
		return err
	}

	if users == nil {
		log.Infof("no users found")
		return nil
	}

	for _, user := range users {
		userNull := user.Nullifier
		eve, err := filtEvents.FilterByNullifier(userNull).Get()

		if err != nil {
			log.WithError(err).Error("failed to filter by nullifier")
			return err
		}
		if eve != nil {
			break
		}

		err = events.Insert(data.Event{
			Nullifier: user.Nullifier,
			Type:      models.TypeEarlyTest,
			Status:    data.EventFulfilled,
		})
		if err != nil {
			return fmt.Errorf("failed to insert `early_test` event: %w", err)
		}
	}

	return nil
}
