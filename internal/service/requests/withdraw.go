package requests

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewWithdraw(r *http.Request) (req resources.WithdrawRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	req.Data.ID = strings.ToLower(req.Data.ID)

	return req, validation.Errors{
		"data/id": validation.Validate(req.Data.ID,
			validation.Required,
			validation.In(strings.ToLower(chi.URLParam(r, "nullifier"))),
			validation.Match(nullifierRegexp)),
		"data/type":              validation.Validate(req.Data.Type, validation.Required, validation.In(resources.WITHDRAW)),
		"data/attributes/amount": validation.Validate(req.Data.Attributes.Amount, validation.Required, validation.Min(1)),
		"data/attributes/proof":  validation.Validate(req.Data.Attributes.Proof, validation.Required),
	}.Filter()
}
