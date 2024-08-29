package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewDailyQuestionEdit(r *http.Request) (req resources.DailyQuestionEditResponse, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	return req, validation.Errors{
		"data/id":         validation.Validate(&req.Data.ID),
		"data/type":       validation.Validate(&req.Data.Type, validation.Required),
		"data/attributes": validation.Validate(&req.Data.Attributes, validation.Required),
	}.Filter()
}
