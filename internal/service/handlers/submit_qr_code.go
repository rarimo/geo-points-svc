package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

var qrCodeRegexp = regexp.MustCompile("[0-9A-Za-z_]+")

func SubmitQRCode(w http.ResponseWriter, r *http.Request) {
	qrCode := chi.URLParam(r, "qr_code")
	if !qrCodeRegexp.MatchString(qrCode) {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"path": fmt.Errorf("invalid qr_code format"),
		})...)
		return
	}

	// never panic because of auth validation
	nullifier := UserClaims(r)[0].Nullifier

	log := Log(r).WithFields(map[string]any{
		"nullifier": nullifier,
		"qr_code":   qrCode,
	})

	balance, err := BalancesQ(r).FilterDisabled().FilterByNullifier(nullifier).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get balance")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if balance == nil {
		log.Debug("Balance not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if !balance.IsVerified() {
		log.Debug("Balance not verified")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	qr, err := QRCodesQ(r).FilterByID(qrCode).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get qr code")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if qr == nil || (!qr.Infinity && qr.UsageCount <= 0) {
		if qr == nil {
			log.Debug("QR code absent in db")
			ape.RenderErr(w, problems.NotFound())
			return
		}
		log.Debug("QR code usage count exceed")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ev, err := EventsQ(r).FilterByNullifier(nullifier).FilterByType(qrCode).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get event by type")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if ev != nil {
		log.Debug("User already scan old qr code")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	ev, err = EventsQ(r).FilterByNullifier(nullifier).FilterByType(models.TypeQRCode).FilterByQRCode(qrCode).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get event by qrcode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if ev != nil {
		log.Debug("User already scan qr code")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	evType := EventTypes(r).Get(models.TypeQRCode, evtypes.FilterInactive)
	if evType == nil {
		log.Debug("Event QR code absent or inactive")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	reward := int64(qr.Reward)
	err = EventsQ(r).Transaction(func() error {
		if !evType.AutoClaim {
			return EventsQ(r).Insert(data.Event{
				Nullifier:    nullifier,
				Type:         models.TypeQRCode,
				Status:       data.EventFulfilled,
				PointsAmount: &reward,
				Meta:         data.Jsonb(fmt.Sprintf(`{"qr_code": "%s"}`, qr.ID)),
			})
		}

		err = EventsQ(r).Insert(data.Event{
			Nullifier:    nullifier,
			Type:         models.TypeQRCode,
			Status:       data.EventClaimed,
			PointsAmount: &reward,
			Meta:         data.Jsonb(fmt.Sprintf(`{"qr_code": "%s"}`, qr.ID)),
		})
		if err != nil {
			return fmt.Errorf("failed to insert event: %w", err)
		}

		level, err := DoLevelRefUpgrade(Levels(r), ReferralsQ(r), balance, reward)
		if err != nil {
			return fmt.Errorf("failed to do lvlup and referrals updates: %w", err)
		}

		err = BalancesQ(r).FilterByNullifier(balance.Nullifier).Update(map[string]any{
			data.ColAmount: pg.AddToValue(data.ColAmount, reward),
			data.ColLevel:  level,
		})
		if err != nil {
			return fmt.Errorf("update balance amount and level: %w", err)
		}

		if !qr.Infinity {
			err = QRCodesQ(r).FilterByID(qrCode).Update(map[string]any{
				data.ColUsageCount: qr.UsageCount - 1,
			})
			if err != nil {
				return fmt.Errorf("failed to update qr code: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		log.WithError(err).Error("Failed to exec tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newEventClaimingStateResponse(balance.Nullifier, true))
}
