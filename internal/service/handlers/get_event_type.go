package handlers

import (
	"net/http"

	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetEventType(w http.ResponseWriter, r *http.Request) {
	name, err := requests.NewGetEventType(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	evType := EventTypes(r).Get(name)
	if evType == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, newEventTypeResponse(*evType))
}
