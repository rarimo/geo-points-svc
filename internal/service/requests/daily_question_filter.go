package requests

import (
	"net/http"

	"github.com/rarimo/geo-points-svc/internal/service/page"
	"gitlab.com/distributed_lab/urlval/v4"
)

type QuestionList struct {
	page.OffsetParams
	Count bool `url:"count"`
}

func NewFilterStartAtDailyQuestions(r *http.Request) (req QuestionList, err error) {
	if err = urlval.Decode(r.URL.Query(), &req); err != nil {
		err = newDecodeError("query", err)
		return
	}

	return req, req.Validate()
}
