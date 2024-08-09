package evtypes

import (
	"context"
	"fmt"

	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type extConfig interface {
	comfig.Logger
	pgdb.Databaser
	EventTypeser
}

func Init(_ context.Context, cfg extConfig, sig chan struct{}) {
	var (
		log   = cfg.Log().WithField("who", "evtypes")
		q     = pg.NewEventTypes(cfg.DB().Clone())
		types = cfg.EventTypes()
	)

	dbTypes, err := q.New().Select()
	if err != nil {
		panic(fmt.Errorf("select all event types: %w", err))
	}

	defer func() {
		types.dbSynced = true
		sig <- struct{}{}
	}()

	if len(dbTypes) == 0 {
		log.Info("No event types in database")
		return
	}

	log.Debugf("Adding/overwriting event types from DB: %+v", dbTypes)
	types.Push(dbTypes...)
}

func InitForOneTimeEvent(cfg extConfig) {
	var (
		log   = cfg.Log().WithField("who", "evtypes")
		q     = pg.NewEventTypes(cfg.DB().Clone())
		types = cfg.EventTypes()
	)

	dbTypes, err := q.New().Select()
	if err != nil {
		panic(fmt.Errorf("select all event types: %w", err))
	}

	defer func() {
		types.dbSynced = true
	}()

	log.Debugf("Adding/overwriting event types from DB: %+v", dbTypes)
	types.Push(dbTypes...)
}
