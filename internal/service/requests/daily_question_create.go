package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewDailyQuestion(r *http.Request) (req resources.DailyQuestionCreateResponse, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	return req, validation.Errors{
		"data/attributes/time_for_answer": validation.Validate(req.Data.Attributes.TimeForAnswer, validation.Required),
		"data/attributes/correct_answer":  validation.Validate(req.Data.Attributes.CorrectAnswer, validation.Required),
		"data/attributes/starts_at":       validation.Validate(req.Data.Attributes.StartsAt, validation.Required),
		"data/attributes/options":         validation.Validate(req.Data.Attributes.Options, validation.Required),
		"data/attributes/reward":          validation.Validate(req.Data.Attributes.Reward, validation.Required),
		"data/attributes/title":           validation.Validate(req.Data.Attributes.Title, validation.Required),
	}.Filter()
}
