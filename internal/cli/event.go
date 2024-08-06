package cli

import (
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/service/event"
)

func eventStart(cfg config.Config, data int) {
	err := event.Run(cfg, data)
	if err != nil {
		cfg.Log().Errorf("Error starting event: %s", err)
	}
}
