package expirywatch

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/workers/cron"
)

const retryPeriod = 1 * time.Minute
const maxRetries = 12

func Run(ctx context.Context, cfg config.Config, sig chan struct{}) {
	w := newWatcher(cfg)
	if err := w.initialRun(); err != nil {
		panic(fmt.Errorf("expiry-watcher: initial run failed: %w", err))
	}
	sig <- struct{}{}

	cron.Init(cfg.Log())
	expirable := w.types.List(func(ev models.EventType) bool {
		return ev.Disabled || ev.ExpiresAt == nil || evtypes.FilterExpired(ev)
	})

	for _, ev := range expirable {
		if ev.ExpiresAt.Before(time.Now().UTC()) {
			continue // although we filtered expired, ensure extra safety due to possible delay
		}

		_, err := cron.NewJob(
			gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(*ev.ExpiresAt)),
			gocron.NewTask(w.job, ctx, ev.Name),
			gocron.WithName(fmt.Sprintf("expiry-watch[%s]", ev.Name)),
		)
		if err != nil {
			panic(fmt.Errorf("failed to initialize job [event_type=%+v]: %w", ev, err))
		}
	}

	cron.Start(ctx)
}
