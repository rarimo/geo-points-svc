package handlers

import (
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

func ListQREventTypes(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewListEventTypes(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if !auth.Authenticates(UserClaims(r), auth.AdminGrant) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	types := EventTypes(r).List(
		func(ev models.EventType) bool {
			return ev.QRCodeValue == nil
		},
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
			Attributes: t.Resource(true),
		}
		if req.Count {
			evCount, err := EventsQ(r).FilterByType(t.Name).FilterByStatus(data.EventFulfilled, data.EventClaimed).Count()
			if err != nil {
				Log(r).WithError(err).Errorf("failed to get %s event usage count", t.Name)
				ape.RenderErr(w, problems.InternalError())
				return
			}
			resTypes[i].Attributes.UsageCount = &evCount
		}
	}

	ape.Render(w, resources.EventTypeListResponse{Data: resTypes})
}
