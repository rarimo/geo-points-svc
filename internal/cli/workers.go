package cli

import (
	"context"
	"sync"

	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/service"
	"github.com/rarimo/geo-points-svc/internal/service/workers/cleanquestiondeadlines"
	"github.com/rarimo/geo-points-svc/internal/service/workers/expirywatch"
	"github.com/rarimo/geo-points-svc/internal/service/workers/leveljustice"
	"github.com/rarimo/geo-points-svc/internal/service/workers/nooneisforgotten"
	"github.com/rarimo/geo-points-svc/internal/service/workers/reopener"
)

// runServices manages service's dependencies and runs them in the correct order
func runServices(ctx context.Context, cfg config.Config, wg *sync.WaitGroup) {
	// signals indicate the finished initialization of each worker
	var (
		reopenerSig             = make(chan struct{})
		expiryWatchSig          = make(chan struct{})
		evTypesSig              = make(chan struct{})
		noOneIsForgottenSig     = make(chan struct{})
		levelJustice            = make(chan struct{})
		cleanDQuestionDeadlines = make(chan struct{})
	)

	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	// all services depend on event types
	run(func() { evtypes.Init(ctx, cfg, evTypesSig) })
	<-evTypesSig

	// these services can safely run in parallel and depend only on event types
	run(func() { reopener.Run(ctx, cfg, reopenerSig) })
	run(func() { expirywatch.Run(ctx, cfg, expiryWatchSig) })

	// this one depends on reopener, because events must be opened before they can be fulfilled
	<-reopenerSig
	run(func() { nooneisforgotten.Run(cfg, noOneIsForgottenSig) })

	// depends on noOneIsForgoten, because this worker can claim events and change balance
	<-noOneIsForgottenSig
	run(func() { leveljustice.Run(cfg, levelJustice) })

	// service depends on all the workers for good UX
	<-expiryWatchSig
	<-levelJustice
	run(func() { service.Run(ctx, cfg) })

	//service for cleaning daily question deadlines after day
	run(func() { cleanquestiondeadlines.Run(cfg, cleanDQuestionDeadlines) })
}
