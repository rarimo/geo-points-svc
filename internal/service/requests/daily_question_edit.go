package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewDailyQuestionEdit(r *http.Request) (req resources.DailyQuestionEditAttributes, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	return req, validation.Errors{
		"time_for_answer": validation.Validate(&req.TimeForAnswer),
		"correct_answer":  validation.Validate(&req.CorrectAnswer),
		"starts_at":       validation.Validate(&req.StartsAt),
		"options":         validation.Validate(&req.Options),
		"reward":          validation.Validate(&req.Reward),
		"title":           validation.Validate(&req.Title),
	}.Filter()
}
