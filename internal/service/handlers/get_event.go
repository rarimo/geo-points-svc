package handlers

import (
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetEvent(w http.ResponseWriter, r *http.Request) {
	id, err := requests.NewGetEvent(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	event, err := EventsQ(r).FilterByID(id).Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get event by ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if event == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	evType := EventTypes(r).Get(event.Type, evtypes.FilterInactive)
	if evType == nil {
		Log(r).Debugf("Event type is not active at this moment: %s", event.Type)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(event.Nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	bal, err := BalancesQ(r).FilterByNullifier(event.Nullifier).FilterDisabled().Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get event by ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if bal == nil {
		Log(r).Debugf("Balance disabled or absent: %s", event.Nullifier)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, resources.EventResponse{Data: newEventModel(*event, evType.Resource(r.Header.Get(langHeader)))})
}
