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

var bonusCodeRegexp = regexp.MustCompile("[0-9A-Za-z_]+")

func SubmitBonusCode(w http.ResponseWriter, r *http.Request) {
	bonusCode := chi.URLParam(r, "bonus_code")
	if !bonusCodeRegexp.MatchString(bonusCode) {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"path": fmt.Errorf("invalid bonus_code format"),
		})...)
		return
	}

	// never panic because of auth validation
	nullifier := UserClaims(r)[0].Nullifier

	log := Log(r).WithFields(map[string]any{
		"nullifier":  nullifier,
		"bonus_code": bonusCode,
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

	bonus, err := BonusCodesQ(r).FilterByID(bonusCode).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get bonus code")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if bonus == nil || (!bonus.Infinity && bonus.UsageCount <= 0) {
		if bonus == nil {
			log.Debug("Bonus code absent in db")
			ape.RenderErr(w, problems.NotFound())
			return
		}
		log.Debug("Bonus code usage count exceed")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ev, err := EventsQ(r).FilterByNullifier(nullifier).FilterByType(bonusCode).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get event by type")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if ev != nil {
		log.Debug("User already scan old bonus code")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	ev, err = EventsQ(r).FilterByNullifier(nullifier).FilterByType(models.TypeBonusCode).FilterByBonusCode(bonusCode).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get event by bonuscode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if ev != nil {
		log.Debug("User already scan bonus code")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	evType := EventTypes(r).Get(models.TypeBonusCode, evtypes.FilterInactive)
	if evType == nil {
		log.Debug("Event Bonus code absent or inactive")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	reward := int64(bonus.Reward)
	err = EventsQ(r).Transaction(func() error {
		if !evType.AutoClaim {
			return EventsQ(r).Insert(data.Event{
				Nullifier:    nullifier,
				Type:         models.TypeBonusCode,
				Status:       data.EventFulfilled,
				PointsAmount: &reward,
				Meta:         data.Jsonb(fmt.Sprintf(`{"bonus_code": "%s"}`, bonus.ID)),
			})
		}

		err = EventsQ(r).Insert(data.Event{
			Nullifier:    nullifier,
			Type:         models.TypeBonusCode,
			Status:       data.EventClaimed,
			PointsAmount: &reward,
			Meta:         data.Jsonb(fmt.Sprintf(`{"bonus_code": "%s"}`, bonus.ID)),
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

		if !bonus.Infinity {
			err = BonusCodesQ(r).FilterByID(bonusCode).Update(map[string]any{
				data.ColUsageCount: bonus.UsageCount - 1,
			})
			if err != nil {
				return fmt.Errorf("failed to update bonus code: %w", err)
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
