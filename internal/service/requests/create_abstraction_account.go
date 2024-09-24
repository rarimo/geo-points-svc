package requests

import (
	"encoding/json"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewCreateAbstractionAccount(r *http.Request) (req resources.CreateAbstractionAccountRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	req.Data.ID = strings.ToLower(req.Data.ID)

	return req, validation.Errors{
		"data/type":             validation.Validate(req.Data.Type, validation.Required, validation.In(resources.ABSTRACTION_ACCOUNT)),
		"data/attributes/proof": validation.Validate(req.Data.Attributes.Proof, validation.Required),
	}.Filter()
}
