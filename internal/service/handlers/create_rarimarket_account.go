package handlers

import (
	"bytes"
	"errors"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	zkp "github.com/rarimo/zkverifier-kit"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateRarimarketAccount(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewCreateRarimarketAccount(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var (
		proof  = req.Data.Attributes.Proof
		getter = zkp.PubSignalGetter{
			ProofType: zkp.GeorgianPassport,
			Signals:   proof.PubSignals,
		}
	)
	nullifierString := getter.Get(zkp.Nullifier)
	nullifierBig, ok := new(big.Int).SetString(nullifierString, 10)
	if !ok {
		Log(r).Errorf("Failed to extract nullifier from proof: %s", nullifierString)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	var nullifierBytes [32]byte
	var nullifier = hexutil.Encode(nullifierBig.FillBytes(nullifierBytes[:]))

	log := Log(r).WithField("nullifier", nullifier)

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	addr, err := Rarimarket(r).GetAccount(nullifierBytes)
	if err != nil {
		log.WithError(err).Error("Failed to get account")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if !bytes.Equal(addr[:], config.ZeroAddress[:]) {
		log.Debug("Rarimarket account already exist")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	err = Verifiers(r).Passport.VerifyProof(proof)
	if err != nil {
		var vErr validation.Errors
		if !errors.As(err, &vErr) {
			log.WithError(err).Error("Failed to verify proof")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	addr, err = Rarimarket(r).CreateAccount(r.Context(), nullifierBytes)
	if err != nil {
		log.WithError(err).Error("Failed to create rarimarket account")
		ape.RenderErr(w, problems.InternalError())
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
