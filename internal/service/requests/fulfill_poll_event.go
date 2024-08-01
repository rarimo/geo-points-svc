package requests

import (
	"encoding/json"
	"net/http"

	val "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/rarimo/geo-points-svc/resources"
	zk "github.com/rarimo/zkverifier-kit"
)

func NewFulfillPollEvent(r *http.Request) (req resources.FulfillPollEventRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	var (
		proof = req.Data.Attributes.Proof
		count = zk.PubSignalsCount(zk.PollParticipation)
	)

	return req, val.Errors{
		"data/type":                         val.Validate(req.Data.Type, val.Required, val.In(resources.FULFILL_POLL_EVENT)),
		"data/attributes/proof/proof":       val.Validate(proof.Proof, val.Required),
		"data/attributes/proof/pub_signals": val.Validate(proof.PubSignals, val.Required, val.Length(count, count)),
		"data/attributes/proposal_id":       val.Validate(req.Data.Attributes.ProposalId, val.Required, is.Digit),
	}.Filter()
}
