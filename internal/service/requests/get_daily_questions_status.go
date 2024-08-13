package requests

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/urlval/v4"
)

type GetDailyQuestionsStatus struct {
	UserNullifier      string `json:"user_nullifier"`
	Availability       bool   `json:"availability"`
	AlreadyDoneForUser bool   `json:"already_done_for_user"`
	TimeToNext         int64  `json:"time_to_next"`
}

func NewGetDailyQuestionsStatus(r *http.Request) (req GetDailyQuestionsStatus, err error) {
	req.UserNullifier = strings.ToLower(chi.URLParam(r, "nullifier"))

	if err = urlval.Decode(r.URL.Query(), &req); err != nil {
		err = newDecodeError("query", err)
		return
	}

	return req, nil
}
