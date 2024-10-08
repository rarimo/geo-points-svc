package handlers

import (
	"net/http"

	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func ListEventTypes(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewListEventTypes(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	types := EventTypes(r).List(
		evtypes.FilterByNames(req.FilterName...),
		evtypes.FilterByFlags(req.FilterFlag...),
		func(ev models.EventType) bool {
			return len(req.FilterNotName) > 0 && !evtypes.FilterByNames(req.FilterNotName...)(ev)
		},
	)

	resTypes := make([]resources.EventType, len(types))
	for i, t := range types {
		resTypes[i] = resources.EventType{
			Key: resources.Key{
				ID:   t.Name,
				Type: resources.EVENT_TYPE,
			},
			Attributes: t.Resource(r.Header.Get(langHeader)),
		}
	}

	ape.Render(w, resources.EventTypeListResponse{Data: resTypes})
}
