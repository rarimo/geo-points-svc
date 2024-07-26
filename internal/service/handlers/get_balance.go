package handlers

import (
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
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
	if balance.ReferredBy == nil {
		balance.Rank = nil
	}

	var (
		referrals             []data.Referral
		referredUsers         int
		rewardedReferredUsers int
	)
	if req.ReferralCodes && balance.ReferredBy != nil {
		referrals, err = ReferralsQ(r).
			FilterByNullifier(req.Nullifier).
			WithStatus().
			Select()
		if err != nil {
			Log(r).WithError(err).Error("Failed to get referrals by nullifier with rewarding field")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		// Infinite referral codes initially have 0 uses and,
		// accordingly, after use, this value will decrease,
		// i.e. the number of invited users for this code will
		// be an absolute value
		//
		// A one-time code is considered used if it has 0 uses,
		// because the initial value is 1
		for _, ref := range referrals {
			if ref.Infinity {
				referredUsers += -int(ref.UsageLeft)
				continue
			}
			if ref.UsageLeft == 0 {
				referredUsers++
			}
		}

		rewardedReferredUsers, err = EventsQ(r).FilterByNullifier(req.Nullifier).FilterByType(models.TypeReferralSpecific).FilterByStatus(data.EventClaimed).Count()
		if err != nil {
			Log(r).WithError(err).Error("Failed to get count rewarded referral_specific events for user")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	ape.Render(w, newBalanceResponse(*balance, referrals, referredUsers, rewardedReferredUsers))
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

func newBalanceResponse(balance data.Balance, referrals []data.Referral, referredUsers, rewardedReferredUsers int) resources.BalanceResponse {
	resp := resources.BalanceResponse{Data: newBalanceModel(balance)}
	boolP := func(b bool) *bool { return &b }

	resp.Data.Attributes.IsDisabled = boolP(balance.ReferredBy == nil)
	resp.Data.Attributes.IsVerified = boolP(BalanceIsVerified(&balance))
	resp.Data.Attributes.ReferredUsersCount = &referredUsers
	resp.Data.Attributes.RewardedReferredUsersCount = &rewardedReferredUsers

	if len(referrals) == 0 {
		return resp
	}

	res := make([]resources.ReferralCode, len(referrals))
	for i, r := range referrals {
		res[i] = resources.ReferralCode{
			Id:     r.ID,
			Status: r.Status,
		}
	}

	resp.Data.Attributes.ReferralCodes = &res
	return resp
}
