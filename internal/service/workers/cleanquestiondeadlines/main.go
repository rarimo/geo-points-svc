package cleanquestiondeadlines

import (
	"context"
	"time"

	"github.com/rarimo/geo-points-svc/internal/config"
)

func Run(ctx context.Context, cfg config.Config, sig chan struct{}) {
	offset := cfg.DailyQuestions().Timezone

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
			cfg.Log().Infof("Сleared daily questions quantity: %v", res)

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
