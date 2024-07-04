package evtypes

import (
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
)

type Types struct {
	m        map[string]models.EventType
	list     []models.EventType
	dbSynced bool
}

func (t *Types) Get(name string, filters ...filter) *models.EventType {
	t.ensureInitialized()
	v, ok := t.m[name]
	if !ok || isFiltered(v, filters...) {
		return nil
	}

	return &v
}

func (t *Types) List(filters ...filter) []models.EventType {
	t.ensureInitialized()
	res := make([]models.EventType, 0, len(t.list))
	for _, v := range t.list {
		if isFiltered(v, filters...) {
			continue
		}
		res = append(res, v)
	}
	return res
}

func (t *Types) Names(filters ...filter) []string {
	t.ensureInitialized()
	res := make([]string, 0, len(t.list))
	for _, v := range t.list {
		if isFiltered(v, filters...) {
			continue
		}
		res = append(res, v.Name)
	}
	return res
}

func (t *Types) PrepareEvents(nullifier string, filters ...filter) []data.Event {
	t.ensureInitialized()
	const extraCap = 1 // in case we append to the resulting slice outside the function
	events := make([]data.Event, 0, len(t.list)+extraCap)

	for _, et := range t.list {
		if isFiltered(et, filters...) {
			continue
		}

		status := data.EventOpen
		if et.Name == models.TypeFreeWeekly {
			status = data.EventFulfilled
		}

		events = append(events, data.Event{
			Nullifier: nullifier,
			Type:      et.Name,
			Status:    status,
		})
	}

	return events
}

// Push adds new event type or overwrites existing one
func (t *Types) Push(types ...models.EventType) {
	for _, et := range types {
		_, ok := t.m[et.Name]
		t.m[et.Name] = et
		if !ok {
			t.list = append(t.list, et)
			continue
		}

		for i := range t.list {
			if t.list[i].Name == et.Name {
				t.list[i] = et
				break
			}
		}
	}
}

func (t *Types) ensureInitialized() {
	if t.m == nil || t.list == nil || !t.dbSynced {
		panic("event types are not correctly initialized")
	}
}
