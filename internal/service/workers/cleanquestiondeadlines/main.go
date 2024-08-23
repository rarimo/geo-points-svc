package cleanquestiondeadlines

import (
	"context"
	"fmt"

	"github.com/go-co-op/gocron/v2"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/service/workers/cron"
)

func Run(ctx context.Context, cfg config.Config, sig chan struct{}) {
	offset := cfg.DailyQuestions().Timezone
	if offset < 0 {
		offset = 12 + offset
	}
	cron.Init(cfg.Log())

	atDayStart := gocron.NewAtTimes(gocron.NewAtTime(uint(offset), 0, 0))

	_, err := cron.NewJob(
		gocron.DailyJob(1, atDayStart),
		gocron.NewTask(job, cfg, ctx, sig),
		gocron.WithName("Daily Questions leaner"),
	)
	if err != nil {
		panic(fmt.Errorf("cleaner daily questions: failed to initialize daily job: %w", err))
	}

}

func job(cfg config.Config, ctx context.Context, sig chan struct{}) {
	select {
	default:
		res := cfg.DailyQuestions().ClearDeadlines()
		cfg.Log().Infof("Сleared daily questions quantity: %v", res)

	case <-sig:
		cfg.Log().Info("Daily Question cleaning stop")
		return

	case <-ctx.Done():
		cfg.Log().Info("Daily Question cleaning stop")
		return
	}
}
