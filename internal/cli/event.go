package cli

import (
	"fmt"
	"sync"

	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/service/event"
)

func eventStart(cfg config.Config, data int) {
	var evTypesSig = make(chan struct{})
	wg := &sync.WaitGroup{}
	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	run(func() { evtypes.InitFoOneTimeEvent(cfg, evTypesSig) })
	<-evTypesSig

	fmt.Println("ini end")
	err := event.Run(cfg, data)
	if err != nil {
		cfg.Log().Errorf("Error starting event: %s", err)
	}
}
