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
		"time_for_answer": validation.Validate(&req.Data.Attributes.TimeForAnswer),
		"correct_answer":  validation.Validate(&req.Data.Attributes.CorrectAnswer),
		"starts_at":       validation.Validate(&req.Data.Attributes.StartsAt),
		"options":         validation.Validate(&req.Data.Attributes.Options),
		"reward":          validation.Validate(&req.Data.Attributes.Reward),
		"title":           validation.Validate(&req.Data.Attributes.Title),
	}.Filter()
}
