package cleanquestiondeadlines

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/workers/cron"
)

func Run(ctx context.Context, cfg config.Config, sig chan struct{}) {
	cron.Init(cfg.Log())
	log := cfg.Log().WithField("who", "daily-questions-cleaner")

	eventsQ := pg.NewEvents(cfg.DB().Clone())
	questionsQ := pg.NewDailyQuestionsQ(cfg.DB().Clone())

	offset := cfg.DailyQuestions().LocalTime(AtDayStart(time.Now().UTC())).Hour()
	_, err := cron.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(uint(offset), 0, 0))),
		gocron.NewTask(func() {
			counts := cfg.DailyQuestions().ClearDeadlines()
			if len(counts) == 0 {
				log.Infof("Questions absent")
				return
			}

			err := eventsQ.New().Transaction(func() error {
				for k := range counts {
					count, err := eventsQ.New().FilterByQuestionID(k).Count()
					if err != nil {
						return fmt.Errorf("failed to get count events by question id: %w", err)
					}

					err = questionsQ.FilterByID(int64(k)).Update(map[string]any{
						data.ColCorrectAnswers:  count,
						data.ColAllParticipants: counts[k],
					})
					if err != nil {
						return fmt.Errorf("failed to update daily question: %w", err)
					}
					log.WithField("question_id", k).Infof("Correct answers: %d; Total participants: %d", count, counts[k])
				}
				return nil
			})
			if err != nil {
				log.WithError(err).Error("Failed to correct update question statistic")
			}
		}),
		gocron.WithName("daily-questions-cleaner"),
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
