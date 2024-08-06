package cli

import (
	"log"

	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/service/event"
)

func eventStart(cfg config.Config, data int) {
	err := event.Run(cfg, data)
	if err != nil {
		log.Fatalf("Error starting event: %v", err)
	}
}
