package requests

import (
	"encoding/json"
	"net/http"

	val "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewCreateEventType(r *http.Request) (req resources.EventTypeResponse, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	attr := req.Data.Attributes
	return req, val.Errors{
		// only QR code events can be currently created or updated
		"data/id":                           val.Validate(req.Data.ID, val.Required),
		"data/type":                         val.Validate(req.Data.Type, val.Required, val.In(resources.EVENT_TYPE)),
		"data/attributes/action_url":        val.Validate(attr.ActionUrl, is.URL),
		"data/attributes/description":       val.Validate(attr.Description, val.Required),
		"data/attributes/frequency":         val.Validate(attr.Frequency, val.Required, val.In(models.Unlimited)),
		"data/attributes/logo":              val.Validate(attr.Logo, is.URL),
		"data/attributes/name":              val.Validate(attr.Name, val.Required, val.In(req.Data.ID)),
		"data/attributes/flag":              val.Validate(attr.Flag, val.Empty),
		"data/attributes/qr_code_value":     val.Validate(attr.QrCodeValue, val.Required),
		"data/attributes/reward":            val.Validate(attr.Reward, val.Required, val.Min(1)),
		"data/attributes/short_description": val.Validate(attr.ShortDescription, val.Required),
		"data/attributes/title":             val.Validate(attr.Title, val.Required),
	}.Filter()
}
