package requests

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewFulfillQREvent(r *http.Request) (req resources.FulfillQrEventRequest, err error) {
	id := chi.URLParam(r, "id")
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	return req, validation.Errors{
		"data/id":                 validation.Validate(req.Data.ID, validation.Required, validation.In(id), is.UUID),
		"data/type":               validation.Validate(req.Data.Type, validation.Required, validation.In(resources.FULFILL_QR_EVENT)),
		"data/attributes/qr_code": validation.Validate(req.Data.Attributes.QrCode, validation.Required, is.Base64),
	}.Filter()
}
