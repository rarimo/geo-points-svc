package leveljustice

import (
	"fmt"

	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/handlers"
)

func Run(cfg config.Config, sig chan struct{}) {
	db := cfg.DB().Clone()
	balances, err := pg.NewBalances(db).FilterDisabled().Select()
	if err != nil {
		panic(fmt.Errorf("failed to select balances: %w", err))
	}

	err = pg.NewEvents(db).Transaction(func() error {
		for _, balance := range balances {
			level, err := handlers.DoLevelRefUpgrade(cfg.Levels(), pg.NewReferrals(db), &balance, 0)
			if err != nil {
				return fmt.Errorf("failed to do lvlup and referrals updates: %w", err)
			}

			err = pg.NewBalances(db).FilterByNullifier(balance.Nullifier).Update(map[string]any{
				data.ColLevel: level,
			})
			if err != nil {
				return fmt.Errorf("update balance amount and level: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	sig <- struct{}{}
}
