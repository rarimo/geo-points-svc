package requests

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	val "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewUpdateEventType(r *http.Request) (req resources.EventTypeResponse, err error) {
	name := chi.URLParam(r, "name")
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	attr := req.Data.Attributes
	return req, val.Errors{
		// only QR code events can be currently created or updated
		// only Unlimited frequency is intended to be used for them
		"data/id":                    val.Validate(req.Data.ID, val.Required, val.In(name)),
		"data/type":                  val.Validate(req.Data.Type, val.Required, val.In(resources.EVENT_TYPE)),
		"data/attributes/action_url": val.Validate(attr.ActionUrl, is.URL),
		"data/attributes/frequency":  val.Validate(attr.Frequency, val.In(string(models.Unlimited))),
		"data/attributes/logo":       val.Validate(attr.Logo, is.URL),
		"data/attributes/reward":     val.Validate(attr.Reward, val.Min(1)),
	}.Filter()
}
