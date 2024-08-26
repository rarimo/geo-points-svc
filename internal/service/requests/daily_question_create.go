package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewDailyQuestion(r *http.Request) (req resources.DailyQuestionDetailsAttributes, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	return req, validation.Errors{
		"time_for_answer": validation.Validate(&req.TimeForAnswer, validation.Required),
		"correct_answer":  validation.Validate(&req.CorrectAnswer, validation.Required),
		"starts_at":       validation.Validate(&req.StartsAt, validation.Required),
		"options":         validation.Validate(&req.Options, validation.Required),
		"reward":          validation.Validate(&req.Reward, validation.Required),
		"title":           validation.Validate(&req.Title, validation.Required),
	}.Filter()
}
