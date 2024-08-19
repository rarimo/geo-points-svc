package requests

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewDailyQuestionUserAccess(r *http.Request) (req resources.UserTimezoneAttributes, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	nullifier := strings.ToLower(chi.URLParam(r, "nullifier"))

	return req, validation.Errors{
		"nullifier": validation.Validate(nullifier, validation.Required, validation.Match(nullifierRegexp)),
		"timezone":  validation.Validate(req.Timezone, validation.Required),
	}.Filter()
}
