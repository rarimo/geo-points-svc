package requests

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	val "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewVerifyPassportV2(r *http.Request) (req resources.VerifyPassportRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, newDecodeError("body", err)
	}

	req.Data.ID = strings.ToLower(req.Data.ID)

	return req, val.Errors{
		"data/id": val.Validate(req.Data.ID,
			val.Required,
			val.In(strings.ToLower(chi.URLParam(r, "nullifier"))),
			val.Match(nullifierRegexp)),
		"data/type": val.Validate(req.Data.Type,
			val.Required,
			val.In(resources.VERIFY_PASSPORT)),
		"data/attributes/anonymous_id": val.Validate(req.Data.Attributes.AnonymousId, val.Required, val.Match(hex32bRegexp)),
	}.Filter()
}
