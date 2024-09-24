package handlers

import (
	"fmt"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateEventType(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewCreateEventType(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if !auth.Authenticates(UserClaims(r), auth.AdminGrant) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	evType, err := EventTypesQ(r).Get(req.Data.Attributes.Name)
	if err != nil {
		Log(r).WithError(err).Error("Failed to get event type by name")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	memEvType := EventTypes(r).Get(req.Data.Attributes.Name)
	if evType != nil || memEvType != nil {
		Log(r).Debugf("Event type %s already exists: inMem: %v, inDb: %v", req.Data.Attributes.Name, memEvType, evType)
		ape.RenderErr(w, problems.Conflict())
		return
	}

	typeModel := models.ResourceToModel(req.Data.Attributes)
	err = EventsQ(r).Transaction(func() error {
		if err = EventTypesQ(r).Insert(typeModel); err != nil {
			return fmt.Errorf("insert event type: %w", err)
		}
		EventTypes(r).Push(typeModel)

		// TODO: add cron jobs for limited events and other special logic when updating other fields is supported
		if evtypes.FilterNotOpenable(typeModel) {
			return nil
		}
		return openEvents(r, typeModel)
	})

	if err != nil {
		Log(r).WithError(err).Error("Failed to add event type and open events")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func openEvents(r *http.Request, evType models.EventType) error {
	balances, err := BalancesQ(r).Select()
	if err != nil {
		return fmt.Errorf("select balances: %w", err)
	}

	balancesWithEvent := make(map[string]struct{})
	existedEvents, err := EventsQ(r).FilterByType(evType.Name).Select()
	if err != nil {
		return fmt.Errorf("failed to select events with type %s: %w", evType.Name, err)
	}

	for _, existedEvent := range existedEvents {
		balancesWithEvent[existedEvent.Nullifier] = struct{}{}
	}

	eventsToInsert := make([]data.Event, 0, len(balances))
	for _, b := range balances {
		if _, ok := balancesWithEvent[b.Nullifier]; ok {
			continue
		}
		eventsToInsert = append(eventsToInsert, data.Event{
			Nullifier: b.Nullifier,
			Status:    data.EventOpen,
			Type:      evType.Name,
		})
	}

	if err = EventsQ(r).Insert(eventsToInsert...); err != nil {
		return fmt.Errorf("insert events: %w", err)
	}

	return nil
}
