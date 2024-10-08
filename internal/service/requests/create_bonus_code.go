package requests

import (
	"encoding/json"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
)

func NewCreateBonusCode(r *http.Request) (req resources.BonusCodeRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	req.Data.ID = strings.ToLower(req.Data.ID)

	errs := validation.Errors{
		"data/type":                   validation.Validate(req.Data.Type, validation.Required, validation.In(resources.BONUS_CODE)),
		"data/attributes/nullifier":   validation.Validate(req.Data.Attributes.Nullifier, validation.Match(nullifierRegexp)),
		"data/attributes/usage_count": validation.Validate(req.Data.Attributes.UsageCount, validation.Min(int(0))),
		"data/attributes/reward":      validation.Validate(req.Data.Attributes.Reward, validation.Min(int(1))),
	}

	return req, errs.Filter()
}
