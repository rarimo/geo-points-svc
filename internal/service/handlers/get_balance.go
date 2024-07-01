package handlers

import (
	"net/http"

	"github.com/rarimo/decentralized-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetBalance(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(req.Nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	var balance *data.Balance
	if req.Rank {
		balance, err = BalancesQ(r).GetWithRank(req.Nullifier)
	} else {
		balance, err = BalancesQ(r).FilterByNullifier(req.Nullifier).Get()
	}

	if err != nil {
		Log(r).WithError(err).Error("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if balance == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	var referral *data.Referral
	if req.ReferralCodes {
		referral, err = ReferralsQ(r).Get(req.Nullifier)
		if err != nil || referral == nil {
			Log(r).WithError(err).Error("Failed to get referral by code nullifier")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	ape.Render(w, newBalanceResponse(*balance, referral))
}

// newBalanceModel forms a balance response without referral fields, which must
// only be accessed with authorization.
func newBalanceModel(balance data.Balance) resources.Balance {
	return resources.Balance{
		Key: resources.Key{
			ID:   balance.Nullifier,
			Type: resources.BALANCE,
		},
		Attributes: resources.BalanceAttributes{
			Amount:    balance.Amount,
			CreatedAt: balance.CreatedAt,
			UpdatedAt: balance.UpdatedAt,
			Rank:      balance.Rank,
			Level:     balance.Level,
		},
	}
}

func newBalanceResponse(balance data.Balance, referral *data.Referral) resources.BalanceResponse {
	resp := resources.BalanceResponse{Data: newBalanceModel(balance)}
	boolP := func(b bool) *bool { return &b }

	resp.Data.Attributes.IsDisabled = boolP(balance.ReferredBy == nil)
	resp.Data.Attributes.IsVerified = boolP(balance.IsVerified)

	if referral != nil {
		resp.Data.Attributes.ReferralCode = &referral.ID
	}
	return resp
}
