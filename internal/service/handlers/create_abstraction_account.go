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

func CreateAbstractionAccount(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewCreateAbstractionAccount(r)
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

	addr, err := Abstraction(r).GetAccount(nullifierBytes)
	if err != nil {
		log.WithError(err).Error("Failed to get account")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if !bytes.Equal(addr[:], config.ZeroAddress[:]) {
		log.Debug("Abstraction account already exist")
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

	addr, err = Abstraction(r).CreateAccount(r.Context(), nullifierBytes)
	if err != nil {
		log.WithError(err).Error("Failed to create abstraction account")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.AbstractionAccountResponse{
		Data: resources.AbstractionAccount{
			Key: resources.Key{
				ID:   hexutil.Encode(addr[:]),
				Type: resources.ABSTRACTION_ACCOUNT,
			},
			Attributes: resources.AbstractionAccountAttributes{
				Address: hexutil.Encode(addr[:]),
			},
		},
	})
}
