package evtypes

import (
	"fmt"

	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type EventTypeser interface {
	EventTypes() *Types
}

type config struct {
	once   comfig.Once
	getter kv.Getter
}

func NewConfig(getter kv.Getter) EventTypeser {
	return &config{getter: getter}
}

func (c *config) EventTypes() *Types {
	return c.once.Do(func() interface{} {
		var raw struct {
			Types []models.EventType `fig:"types,required"`
		}

		err := figure.Out(&raw).
			From(kv.MustGetStringMap(c.getter, "event_types")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out event_types: %s", err))
		}

		m := make(map[string]models.EventType, len(raw.Types))
		for _, t := range raw.Types {
			if !checkFreqValue(t.Frequency) {
				panic(fmt.Errorf("invalid frequency: %s", t.Frequency))
			}

			if t.ExpiresAt != nil && t.StartsAt != nil && !t.StartsAt.Before(*t.ExpiresAt) {
				panic(fmt.Errorf("starts_at must be before expires_at: %s > %s",
					t.StartsAt, t.ExpiresAt))
			}

			m[t.Name] = t
		}

		return &Types{m: m, list: raw.Types}
	}).(*Types)
}

func checkFreqValue(f models.Frequency) bool {
	switch f {
	case models.OneTime, models.Daily, models.Weekly, models.Unlimited:
		return true
	}
	return false
}
