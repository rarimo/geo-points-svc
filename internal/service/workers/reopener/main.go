package reopener

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/workers/cron"
)

const retryPeriod = 5 * time.Minute
const maxRetries = 12

func Run(ctx context.Context, cfg config.Config, sig chan struct{}) {
	if err := initialRun(cfg); err != nil {
		panic(fmt.Errorf("reopener: initial run failed: %w", err))
	}
	sig <- struct{}{}

	cron.Init(cfg.Log())

	if err := runStartingWatchers(ctx, cfg); err != nil {
		panic(fmt.Errorf("reopener: failed to initialize opener: %w", err))
	}

	atDayStart := gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0))

	daily := newWorker(cfg, models.Daily)
	_, err := cron.NewJob(
		gocron.DailyJob(1, atDayStart),
		gocron.NewTask(daily.job, ctx),
		gocron.WithName(daily.name),
	)
	if err != nil {
		panic(fmt.Errorf("reopener: failed to initialize daily job: %w", err))
	}

	weekly := newWorker(cfg, models.Weekly)
	_, err = cron.NewJob(
		gocron.WeeklyJob(1, gocron.NewWeekdays(time.Monday), atDayStart),
		gocron.NewTask(weekly.job, ctx),
		gocron.WithName(weekly.name),
	)
	if err != nil {
		panic(fmt.Errorf("reopener: failed to initialize weekly job: %w", err))
	}

	cron.Start(ctx)
}

func prepareForReopening(events []data.ReopenableEvent) []data.Event {
	res := make([]data.Event, len(events))

	for i, ev := range events {
		res[i] = data.Event{
			Nullifier: ev.Nullifier,
			Type:      ev.Type,
			Status:    data.EventOpen,
		}

		if ev.Type == models.TypeFreeWeekly {
			res[i].Status = data.EventFulfilled
		}
	}

	return res
}
