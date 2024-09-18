package handlers

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-chi/chi"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetRarimarketAccount(w http.ResponseWriter, r *http.Request) {
	nullifier := strings.ToLower(chi.URLParam(r, "nullifier"))

	log := Log(r).WithField("nullifier", nullifier)

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	nullifierBytes, err := hexutil.Decode(nullifier)
	if err != nil {
		log.WithError(err).Error("Failed to parse nullifier")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	addr, err := Rarimarket(r).GetAccount([32]byte(nullifierBytes))
	if err != nil {
		log.WithError(err).Error("Failed to get account")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if bytes.Equal(addr[:], config.ZeroAddress[:]) {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, resources.RarimarketAccountResponse{
		Data: resources.RarimarketAccount{
			Key: resources.Key{
				ID:   hexutil.Encode(addr[:]),
				Type: resources.RARIMARKET_ACCOUNT,
			},
			Attributes: resources.RarimarketAccountAttributes{
				Address: hexutil.Encode(addr[:]),
			},
		},
	})
}
