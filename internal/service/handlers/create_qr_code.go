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

func CreateQRCode(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewCreateQRCode(r)
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

	qrValue := make([]byte, 10)
	_, err = rand.Read(qrValue[:])
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

	qr := data.QRCode{
		ID:         "one_time_" + hex.EncodeToString(qrValue),
		Nullifier:  sql.NullString{},
		UsageCount: usageCount,
		Reward:     reward,
	}

	if err = QRCodesQ(r).Insert(qr); err != nil {
		Log(r).WithError(err).Error("Failed to insert qr code")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.QrCodeRequest{
		Data: resources.QrCode{
			Key: resources.Key{
				ID:   qr.ID,
				Type: resources.QR_CODE,
			},
			Attributes: resources.QrCodeAttributes{
				Reward:     &reward,
				UsageCount: &usageCount,
			},
		},
	})

}
