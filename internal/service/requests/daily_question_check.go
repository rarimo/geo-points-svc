package requests

import (
	"encoding/json"
	"net/http"

	"github.com/rarimo/geo-points-svc/resources"
)

func NewDailyQuestionAnswer(r *http.Request) (req resources.DailyQuestionAnswersResponse, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	return req, nil
}
