package evtypes

import (
	"time"

	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
)

// Filter functions work in the following way:
//
// 1. For FilterBy* functions, the config is only added when it matches the filter:
// FilterByName(name1, name2) will only return events with name1 or name2
//
// 2. For other Filter* functions, the configs matching the filter are excluded:
// FilterExpired eliminates all expired events (instead of including only them)

type filter func(models.EventType) bool

func FilterExpired(ev models.EventType) bool {
	return ev.ExpiresAt != nil && ev.ExpiresAt.Before(time.Now().UTC())
}

func FilterNotStarted(ev models.EventType) bool {
	return ev.StartsAt != nil && ev.StartsAt.After(time.Now().UTC())
}

func FilterInactive(ev models.EventType) bool {
	return ev.Disabled || FilterExpired(ev) || FilterNotStarted(ev)
}

func FilterNotOpenable(ev models.EventType) bool {
	return FilterInactive(ev) || ev.NoAutoOpen
}

func FilterByFrequency(f models.Frequency) func(models.EventType) bool {
	return func(ev models.EventType) bool {
		return ev.Frequency != f
	}
}

func FilterByAutoClaim(autoClaim bool) func(models.EventType) bool {
	return func(ev models.EventType) bool {
		return ev.AutoClaim != autoClaim
	}
}

func FilterByNames(names ...string) func(models.EventType) bool {
	return func(ev models.EventType) bool {
		if len(names) == 0 {
			return false
		}
		for _, name := range names {
			if ev.Name == name {
				return false
			}
		}
		return true
	}
}

func FilterByFlags(flags ...string) func(models.EventType) bool {
	return func(ev models.EventType) bool {
		if len(flags) == 0 {
			return false
		}
		for _, flag := range flags {
			if ev.Flag() == flag {
				return false
			}
		}
		return true
	}
}

func isFiltered(ev models.EventType, filters ...filter) bool {
	for _, f := range filters {
		if f(ev) {
			return true
		}
	}
	return false
}
