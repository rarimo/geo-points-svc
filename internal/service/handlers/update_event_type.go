package handlers

import (
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
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
			ape.RenderErr(w, problems.Conflict())
			return
		}
	}

	typeModel := models.ResourceToModel(req.Data.Attributes)
	res, err := EventTypesQ(r).FilterByNames(typeModel.Name).Update(typeModel.ForUpdate())
	if err != nil {
		Log(r).WithError(err).Error("Failed to update event type")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if len(res) == 0 {
		Log(r).Error("Count of updated event_types = 0")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	EventTypes(r).Push(typeModel)
	ape.Render(w, newEventTypeResponse(res[0]))
}

func newEventTypeResponse(evType models.EventType) resources.EventTypeResponse {
	return resources.EventTypeResponse{
		Data: resources.EventType{
			Key: resources.Key{
				ID:   evType.Name,
				Type: resources.EVENT_TYPE,
			},
			Attributes: evType.Resource(),
		},
	}
}
