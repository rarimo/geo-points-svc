package handlers

import (
	"bytes"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/jsonapi"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/requests"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Withdraw(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewWithdraw(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	log := Log(r).WithFields(map[string]any{
		"nullifier":     req.Data.ID,
		"points_amount": req.Data.Attributes.Amount,
	})

	var (
		nullifier = req.Data.ID
		proof     = req.Data.Attributes.Proof
	)

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	balance, errs := getAndVerifyBalanceEligibility(r, nullifier, &proof)
	if len(errs) > 0 {
		ape.RenderErr(w, errs...)
		return
	}

	var nullifierBytes [32]byte
	// never panic because of request validation
	new(big.Int).SetBytes(hexutil.MustDecode(nullifier)).FillBytes(nullifierBytes[:])

	addr, err := Rarimarket(r).GetAccount(nullifierBytes)
	if err != nil {
		log.WithError(err).Error("Failed to get account")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if bytes.Equal(addr[:], config.ZeroAddress[:]) {
		log.Info("Rarimarket account absent. Creating new!")
		addr, err = Rarimarket(r).CreateAccount(r.Context(), nullifierBytes)
		if err != nil {
			log.WithError(err).Error("Failed to create account")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	errs = isEligibleToWithdraw(r, balance, req.Data.Attributes.Amount)
	if len(errs) > 0 {
		ape.RenderErr(w, errs...)
		return
	}

	err = EventsQ(r).Transaction(func() error {
		err = BalancesQ(r).FilterByNullifier(nullifier).Update(map[string]any{
			data.ColAmount: pg.AddToValue(data.ColAmount, -req.Data.Attributes.Amount),
		})
		if err != nil {
			return fmt.Errorf("decrease points amount: %w", err)
		}

		_, err = WithdrawalsQ(r).Insert(data.Withdrawal{
			Nullifier: nullifier,
			Amount:    req.Data.Attributes.Amount,
		})
		if err != nil {
			return fmt.Errorf("add withdrawal entry: %w", err)
		}

		err = Rarimarket(r).Mint(r.Context(), addr, new(big.Int).Mul(Rarimarket(r).PointPrice, big.NewInt(req.Data.Attributes.Amount)))

		if err != nil {
			return fmt.Errorf("failed to mint points: %w", err)
		}
		return nil
	})
	if err != nil {
		log.WithError(err).Error("Failed to perform withdrawal")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// balance should exist cause of previous logic
	balance, err = BalancesQ(r).GetWithRank(nullifier)
	if err != nil {
		log.WithError(err).Error("Failed to get balance by nullifier with rank")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newBalanceModel(*balance))
}

func isEligibleToWithdraw(
	r *http.Request,
	balance *data.Balance,
	amount int64,
) []*jsonapi.ErrorObject {

	mapValidationErr := func(field, format string, a ...any) []*jsonapi.ErrorObject {
		return problems.BadRequest(validation.Errors{
			field: fmt.Errorf(format, a...),
		})
	}

	switch {
	case !balance.IsVerified():
		return mapValidationErr("data/attributes/proof", "passport must be proven beforehand")
	case balance.Amount < amount:
		return mapValidationErr("data/attributes/amount", "insufficient balance: %d", balance.Amount)
	case !Levels(r).WithdrawalAllowed(balance.Level):
		return mapValidationErr("level", "must up level to have withdraw ability")
	}

	return nil
}
