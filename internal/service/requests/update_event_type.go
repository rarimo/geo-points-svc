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
		// not updatable, as QR code includes event type name
		"data/attributes/qr_code_value": val.Validate(attr.QrCodeValue, val.Empty),

		"data/attributes/poll_event_id": val.Validate(attr.PollEventId, val.Nil),
		"data/attributes/poll_contract": val.Validate(attr.PollContract, val.Nil),
		// these fields are not currently supported, because cron jobs implementation is required
		"data/attributes/starts_at":  val.Validate(attr.StartsAt, val.Empty),
		"data/attributes/expires_at": val.Validate(attr.ExpiresAt, val.Empty),
		// read-only fields due to reusing the same model
		"data/attributes/flag":        val.Validate(attr.Flag, val.Empty),
		"data/attributes/usage_count": val.Validate(attr.UsageCount, val.Empty),
	}.Filter()
}
