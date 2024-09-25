package notificationdailyquestion

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/workers/cron"
)

func Run(ctx context.Context, cfg config.Config, sig chan struct{}) {
	cron.Init(cfg.Log())
	log := cfg.Log().WithField("who", "daily-questions-notification")

	dqConfig := cfg.DailyQuestions()
	questionsQ := pg.NewDailyQuestionsQ(cfg.DB().Clone())

	_, err := cron.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(uint(dqConfig.RawLocation+dqConfig.Notifications.SendAt), 0, 0))),
		gocron.NewTask(func() {
			dailyQuestion, err := questionsQ.FilterDayQuestions(AtDayStart(dqConfig.LocalTime(time.Now())).UTC()).Get()
			if err != nil {
				log.WithError(err).Error("Failed to get daily question")
				return
			}
			if dailyQuestion == nil {
				log.Infof("There's no daily question today")
				return
			}

			if err = dqConfig.SendNotification(); err != nil {
				log.WithError(err).Error("Failed to send notification")
				return
			}

		}),
		gocron.WithName("daily-questions-notification"),
	)
	if err != nil {
		panic(fmt.Errorf("failed to initialize daily job: %w", err))
	}
	sig <- struct{}{}

	cron.Start(ctx)
}

func AtDayStart(date time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, date.Location())
}
