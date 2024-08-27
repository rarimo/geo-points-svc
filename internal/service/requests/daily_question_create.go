package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewDailyQuestion(r *http.Request) (req resources.DailyQuestionCreateAttributes, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	return req, validation.Errors{
		"data/attributes/time_for_answer": validation.Validate(req.TimeForAnswer, validation.Required),
		"data/attributes/correct_answer":  validation.Validate(req.CorrectAnswer, validation.Required),
		"data/attributes/starts_at":       validation.Validate(req.StartsAt, validation.Required),
		"data/attributes/options":         validation.Validate(req.Options, validation.Required),
		"data/attributes/reward":          validation.Validate(req.Reward, validation.Required),
		"data/attributes/title":           validation.Validate(req.Title, validation.Required),
	}.Filter()
}
