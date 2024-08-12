package page

import (
	"math"
	"net/http"
	"strconv"

	val "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const pageParamNumber = "page[number]"

// OffsetParams is a wrapper around pgdb.OffsetPageParams with useful validation and rendering methods
type OffsetParams struct {
	pgdb.OffsetPageParams
}

func (p *OffsetParams) Validate() error {
	return val.Errors{
		pageParamLimit:  val.Validate(p.Limit, val.Max(maxLimit)),
		pageParamNumber: val.Validate(p.PageNumber, val.Max(uint64(math.MaxInt32))),
		pageParamOrder:  val.Validate(p.Order, val.In(pgdb.OrderTypeAsc, pgdb.OrderTypeDesc)),
	}.Filter()
}

func (p *OffsetParams) GetLinks(r *http.Request, resourceCount uint64) *resources.Links {
	result := resources.Links{
		Self: p.getLink(r, p.PageNumber),
	}

	if p.PageNumber != 0 {
		result.Prev = p.getLink(r, p.PageNumber-1)
	}

	if p.Limit*p.PageNumber < resourceCount {
		result.Next = p.getLink(r, p.PageNumber+1)
	}
	return &result
}

func (p *OffsetParams) getLink(r *http.Request, number uint64) string {
	u := r.URL
	query := u.Query()
	query.Set(pageParamNumber, strconv.FormatUint(number, 10))
	query.Set(pageParamLimit, strconv.FormatUint(p.Limit, 10))
	query.Set(pageParamOrder, p.Order)
	u.RawQuery = query.Encode()
	return u.String()
}
