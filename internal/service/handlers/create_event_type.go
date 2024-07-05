package handlers

import (
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
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
	if err = EventTypesQ(r).Insert(typeModel); err != nil {
		Log(r).WithError(err).Error("Failed to insert event type")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	EventTypes(r).Push(typeModel)

	w.WriteHeader(http.StatusNoContent)
}
