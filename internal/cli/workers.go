package cli

import (
	"context"
	"sync"

	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/service"
	"github.com/rarimo/geo-points-svc/internal/service/workers/expirywatch"
	"github.com/rarimo/geo-points-svc/internal/service/workers/nooneisforgotten"
	"github.com/rarimo/geo-points-svc/internal/service/workers/reopener"
)

// runServices manages service's dependencies and runs them in the correct order
func runServices(ctx context.Context, cfg config.Config, wg *sync.WaitGroup) {
	// signals indicate the finished initialization of each worker
	var (
		reopenerSig         = make(chan struct{})
		expiryWatchSig      = make(chan struct{})
		noOneIsForgottenSig = make(chan struct{})
	)

	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	// these services can safely run in parallel and don't have dependencies
	run(func() { reopener.Run(ctx, cfg, reopenerSig) })
	run(func() { expirywatch.Run(ctx, cfg, expiryWatchSig) })

	// this one depends on reopener, because events must be opened before they can be fulfilled
	<-reopenerSig
	run(func() { nooneisforgotten.Run(cfg, noOneIsForgottenSig) })

	// service depends on all the workers for good UX
	<-expiryWatchSig
	<-noOneIsForgottenSig
	run(func() { service.Run(ctx, cfg) })
}
