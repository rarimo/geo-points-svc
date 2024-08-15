package requests

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AnswerDailyQuestionData struct {
	QuestionID int    `json:"question_id"`
	Nullifier  string `json:"nullifier"`
	Answer     string `json:"answer"`
}

func NewAnswerDailyQuestionData(r *http.Request) (req AnswerDailyQuestionData, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}
	req.Nullifier = strings.ToLower(chi.URLParam(r, "nullifier"))

	return req, validation.Errors{
		"question_id": validation.Validate(req.QuestionID, validation.Required),
		"nullifier":   validation.Validate(req.Nullifier, validation.Required, validation.Match(nullifierRegexp)),
		"answer":      validation.Validate(req.Answer, validation.Required),
	}.Filter()
}
