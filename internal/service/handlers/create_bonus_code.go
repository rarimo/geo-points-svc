package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateBonusCode(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewCreateBonusCode(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if !auth.Authenticates(UserClaims(r), auth.AdminGrant) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	var (
		_              = req.Data.Attributes.Nullifier
		dataUsageCount = req.Data.Attributes.UsageCount
		dataReward     = req.Data.Attributes.Reward
		usageCount     = 1
		reward         = 10
	)

	bonusValue := make([]byte, 10)
	_, err = rand.Read(bonusValue[:])
	if err != nil {
		Log(r).WithError(err).Error("Failed to get rand bytes")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if dataUsageCount != nil {
		usageCount = *dataUsageCount
	}

	if dataReward != nil {
		reward = *dataReward
	}

	bonus := data.BonusCode{
		ID:         "one_time_" + hex.EncodeToString(bonusValue),
		Nullifier:  sql.NullString{},
		UsageCount: usageCount,
		Reward:     reward,
	}

	if err = BonusCodesQ(r).Insert(bonus); err != nil {
		Log(r).WithError(err).Error("Failed to insert bonus code")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.BonusCodeRequest{
		Data: resources.BonusCode{
			Key: resources.Key{
				ID:   bonus.ID,
				Type: resources.BONUS_CODE,
			},
			Attributes: resources.BonusCodeAttributes{
				Reward:     &reward,
				UsageCount: &usageCount,
			},
		},
	})

}
