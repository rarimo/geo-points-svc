package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewCreateBalanceV2(r *http.Request) (req resources.Relation, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	if req.Data == nil {
		return req, fmt.Errorf("data: not provided")
	}

	req.Data.ID = strings.ToLower(req.Data.ID)

	errs := validation.Errors{
		"data/id":   validation.Validate(req.Data.ID, validation.Required, validation.Match(nullifierRegexp)),
		"data/type": validation.Validate(req.Data.Type, validation.Required, validation.In(resources.CREATE_BALANCE)),
	}

	return req, errs.Filter()
}
