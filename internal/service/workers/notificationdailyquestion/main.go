package notificationdailyquestion

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/workers/cleanquestiondeadlines"
	"github.com/rarimo/geo-points-svc/internal/service/workers/cron"
	"time"
)

func Run(ctx context.Context, cfg config.Config, sig chan struct{}) {
	cron.Init(cfg.Log())
	log := cfg.Log().WithField("who", "daily-questions-cleaner")

	questionsQ := pg.NewDailyQuestionsQ(cfg.DB().Clone())

	offset := cfg.DailyQuestions().LocalTime(cleanquestiondeadlines.AtDayStart(time.Now().UTC())).Hour()
	_, err := cron.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(uint(offset), 0, 0))),
		gocron.NewTask(func() {
			curQuestion, err := questionsQ.FilterDayQuestions(time.Now().UTC()).Get()
			if err != nil {
				err = fmt.Errorf("error get daily question: %s", err)
				return
			}
			if curQuestion == nil {
				log.Infof("There's no daily question today")
				return
			}
			//TODO
		}),
		gocron.WithName("daily-questions-notification"),
	)
	if err != nil {
		panic(fmt.Errorf("failed to initialize daily job: %w", err))
	}

	sig <- struct{}{}
	cron.Start(ctx)
}
