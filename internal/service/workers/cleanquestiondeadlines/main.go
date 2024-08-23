package cleanquestiondeadlines

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/service/workers/cron"
)

func Run(ctx context.Context, cfg config.Config, sig chan struct{}) {
	cron.Init(cfg.Log())

	offset := cfg.DailyQuestions().LocalTime(atDayStart(time.Now().UTC())).Hour()
	_, err := cron.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(uint(offset), 0, 0))),
		gocron.NewTask(func() {
			cfg.DailyQuestions()
		}),
		gocron.WithName("daily-questions-cleaner"),
	)
	if err != nil {
		panic(fmt.Errorf(": failed to initialize daily job: %w", err))
	}

	for {
		now := time.Now().UTC().Add(time.Duration(offset) * time.Hour)
		cfg.Log().Info("Daily Question cleaning start")
		nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC).
			Add(time.Duration(offset) * time.Hour)

		durationUntilNextTick := nextMidnight.Sub(now)

		timer := time.NewTimer(durationUntilNextTick)

		select {
		case <-timer.C:
			res := cfg.DailyQuestions().ClearDeadlines()
			cfg.Log().Infof("Cleared daily questions quantity: %v", res)

			timer.Stop()

		case <-sig:
			cfg.Log().Info("Daily Question cleaning stop")
			timer.Stop()
			return

		case <-ctx.Done():
			cfg.Log().Info("Daily Question cleaning stop")
			timer.Stop()
			return
		}
	}
}

func atDayStart(date time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, date.Location())
}
