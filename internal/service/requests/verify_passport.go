package requests

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi"
	val "github.com/go-ozzo/ozzo-validation/v4"
	zkptypes "github.com/iden3/go-rapidsnark/types"
	"github.com/rarimo/geo-points-svc/resources"
)

var (
	nullifierRegexp = regexp.MustCompile("^0x[0-9a-fA-F]{64}$")
	hex32bRegexp    = regexp.MustCompile("^[0-9a-f]{64}$")
	// endpoint is hardcoded to reuse handlers.VerifyPassport
	joinProgramPathRegexp = regexp.MustCompile("^/integrations/geo-points-svc/v1/public/balances/0x[0-9a-fA-F]{64}/join_program$")
)

func NewVerifyPassport(r *http.Request) (req resources.VerifyPassportRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, newDecodeError("body", err)
	}

	req.Data.ID = strings.ToLower(req.Data.ID)
	var (
		attr  = req.Data.Attributes
		proof zkptypes.ZKProof // safe dereference
	)

	if attr.Proof != nil {
		proof = *attr.Proof
	}

	return req, val.Errors{
		"data/id": val.Validate(req.Data.ID,
			val.Required,
			val.In(strings.ToLower(chi.URLParam(r, "nullifier"))),
			val.Match(nullifierRegexp)),
		"data/type": val.Validate(req.Data.Type,
			val.Required,
			val.In(resources.VERIFY_PASSPORT)),
		"data/attributes/anonymous_id":      val.Validate(attr.AnonymousId, val.Required, val.Match(hex32bRegexp)),
		"data/attributes/proof":             val.Validate(attr.Proof, val.When(joinProgramPathRegexp.MatchString(r.URL.Path), val.Nil)),
		"data/attributes/proof/proof":       val.Validate(proof.Proof, val.When(attr.Proof != nil, val.Required)),
		"data/attributes/proof/pub_signals": val.Validate(proof.PubSignals, val.When(attr.Proof != nil, val.Required, val.Length(24, 24))),
	}.Filter()
}
