package handlers

import (
	"fmt"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateEventType(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewUpdateEventType(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if !auth.Authenticates(UserClaims(r), auth.AdminGrant) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	evType := EventTypes(r).Get(req.Data.Attributes.Name)
	if evType == nil {
		evType, err = EventTypesQ(r).Get(req.Data.Attributes.Name)
		if err != nil {
			Log(r).WithError(err).Error("Failed to get event type by name")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if evType == nil {
			Log(r).Debugf("Event type %s not found", req.Data.Attributes.Name)
			ape.RenderErr(w, problems.NotFound())
			return
		}
	}

	typeModel := models.ResourceToModel(req.Data.Attributes)

	var updated []models.EventType
	err = EventsQ(r).Transaction(func() error {
		updated, err = EventTypesQ(r).FilterByNames(typeModel.Name).Update(typeModel.ForUpdate())
		if err != nil {
			return fmt.Errorf("update event type: %w", err)
		}
		if len(updated) != 1 {
			return fmt.Errorf("critical: count of updated event types is %d, expected 1", len(updated))
		}
		// Currently, event cannot be 'not openable' in other ways,
		// add extra checks more fields are supported.
		if evType.Disabled == typeModel.Disabled {
			return nil
		}
		// Open events if we have enabled the type, otherwise clean them up.
		if !typeModel.Disabled {
			return openQREvents(r, typeModel)
		}

		deleted, err := EventsQ(r).
			FilterByType(typeModel.Name).
			FilterByStatus(data.EventOpen, data.EventFulfilled).
			Delete()
		if err != nil {
			return fmt.Errorf("delete disabled events: %w", err)
		}

		Log(r).Infof("Deleted %d events on disabling event type %s", deleted, typeModel.Name)
		return nil
	})

	if err != nil {
		Log(r).WithError(err).Error("Failed to update event type")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	EventTypes(r).Push(typeModel)
	resp := newEventTypeResponse(updated[0], r.Header.Get(langHeader))
	resp.Data.Attributes.QrCodeValue = typeModel.QRCodeValue
	ape.Render(w, resp)
}

func newEventTypeResponse(evType models.EventType, locale string) resources.EventTypeResponse {
	return resources.EventTypeResponse{
		Data: resources.EventType{
			Key: resources.Key{
				ID:   evType.Name,
				Type: resources.EVENT_TYPE,
			},
			Attributes: evType.Resource(locale),
		},
	}
}
