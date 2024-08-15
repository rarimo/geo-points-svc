package requests

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

type GetDailyQuestionRequest struct {
	UserNullifier string `json:"user_nullifier"`
}

func NewGetDailyQuestionsRequest(r *http.Request) (req GetDailyQuestionRequest, err error) {
	req.UserNullifier = strings.ToLower(chi.URLParam(r, "nullifier"))
	return req, nil
}
