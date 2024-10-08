package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func ListEvents(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewListEvents(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(*req.FilterNullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	bal, err := BalancesQ(r).FilterByNullifier(*req.FilterNullifier).FilterDisabled().Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get event by ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if bal == nil {
		Log(r).Debugf("Balance disabled or absent: %s", *req.FilterNullifier)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if req.FilterHasExpiration != nil {
		filter := func(ev models.EventType) bool { return ev.ExpiresAt != nil }
		// keep in mind that these filters eliminate values matching the condition, see evtypes/filters.go
		if *req.FilterHasExpiration {
			filter = func(ev models.EventType) bool { return ev.ExpiresAt == nil }
		}

		types := EventTypes(r).Names(filter)
		if len(types) == 0 {
			// filter won't be correctly applied if there are no types matching the condition
			ape.Render(w, newEventsResponse(nil, nil))
			return
		}
		req.FilterType = append(req.FilterType, types...)
	}

	inactiveTypes := EventTypes(r).Names(func(ev models.EventType) bool {
		return !evtypes.FilterInactive(ev)
	})

	events, err := EventsQ(r).
		FilterByNullifier(*req.FilterNullifier).
		FilterByStatus(req.FilterStatus...).
		FilterByType(req.FilterType...).
		FilterByNotType(req.FilterNotType...).
		FilterInactiveNotClaimed(inactiveTypes...).
		Page(&req.OffsetPageParams).
		Select()
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to get filtered paginated event list: nullifier=%s status=%v type=%v",
			*req.FilterNullifier, req.FilterStatus, req.FilterType)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var eventsCount int
	eventsCount, err = EventsQ(r).
		FilterByNullifier(*req.FilterNullifier).
		FilterByStatus(req.FilterStatus...).
		FilterByType(req.FilterType...).
		FilterInactiveNotClaimed(inactiveTypes...).
		Count()
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to count filtered events: nullifier=%s status=%v type=%v",
			*req.FilterNullifier, req.FilterStatus, req.FilterType)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	meta, err := getOrderedEventsMeta(events, r)
	if err != nil {
		Log(r).WithError(err).Error("Failed to get ordered events metadata")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	resp := newEventsResponse(events, meta)
	resp.Links = req.OffsetParams.GetLinks(r, uint64(eventsCount))
	if req.Count {
		_ = resp.PutMeta(struct {
			EventsCount int `json:"events_count"`
		}{eventsCount})
	}
	ape.Render(w, resp)
}

func getOrderedEventsMeta(events []data.Event, r *http.Request) ([]resources.EventStaticMeta, error) {
	res := make([]resources.EventStaticMeta, len(events))

	for i, event := range events {
		// even if event type was disabled, we should return it from history
		evType := EventTypes(r).Get(event.Type)
		if evType == nil {
			return nil, errors.New("wrong event type is stored in DB: might be bad event config")
		}
		res[i] = evType.Resource(r.Header.Get(langHeader))
	}

	return res, nil
}

func newEventModel(event data.Event, meta resources.EventStaticMeta) resources.Event {
	return resources.Event{
		Key: resources.Key{
			ID:   event.ID,
			Type: resources.EVENT,
		},
		Attributes: resources.EventAttributes{
			CreatedAt:     event.CreatedAt,
			UpdatedAt:     event.UpdatedAt,
			HasExpiration: meta.ExpiresAt != nil,
			Meta: resources.EventMeta{
				Static:  meta,
				Dynamic: (*json.RawMessage)(&event.Meta),
			},
			Status:       event.Status.String(),
			PointsAmount: event.PointsAmount,
		},
	}
}

func newEventsResponse(events []data.Event, meta []resources.EventStaticMeta) *resources.EventListResponse {
	list := make([]resources.Event, len(events))
	for i, event := range events {
		list[i] = newEventModel(event, meta[i])
	}

	return &resources.EventListResponse{Data: list}
}
