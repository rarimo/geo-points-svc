package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CheckDailyQuestion(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewDailyQuestionAnswer(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nullifier := strings.ToLower(chi.URLParam(r, "nullifier"))
	option := req.Data.Attributes.Answer
	dq := DailyQuestions(r)

	if !auth.Authenticates(UserClaims(r), auth.VerifiedGrant(nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	log := Log(r).WithField("nullifier", nullifier).WithField("option", option)

	balance, err := BalancesQ(r).FilterByNullifier(nullifier).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if balance == nil {
		log.Debug("Balance absent")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	deadline := dq.GetDeadline(nullifier)
	if deadline == nil || deadline.At.After(time.Now()) {
		log.Debugf("Question already recieved with deadline %s", deadline.At)
		ape.RenderErr(w, problems.Conflict())
		return
	}

	question, err := DailyQuestionsQ(r).FilterByID(int64(deadline.ID)).Get()
	if err != nil {
		log.WithError(err).WithField("questions_id", deadline.ID).Error("Failed to get question by id from deadline")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if question == nil {
		log.WithField("questions_id", deadline.ID).Error("Question absent, but deadline exists")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if question.CorrectAnswer != option {
		ape.Render(w, newDailyAnswer(question))
		return
	}

	err = EventsQ(r).Transaction(func() error {
		evType := EventTypes(r).Get(models.TypeDailyQuestion, evtypes.FilterInactive)
		if evType == nil {
			log.Info("Daily question event type inactive")
			return nil
		}

		if !evType.AutoClaim {
			return EventsQ(r).Insert(data.Event{
				Nullifier:    nullifier,
				Type:         models.TypeDailyQuestion,
				Status:       data.EventFulfilled,
				PointsAmount: &question.Reward,
				Meta:         data.Jsonb(fmt.Sprintf(`{"question_id": %d}`, question.ID)),
			})
		}

		err = EventsQ(r).Insert(data.Event{
			Nullifier:    nullifier,
			Type:         models.TypeDailyQuestion,
			Status:       data.EventClaimed,
			PointsAmount: &question.Reward,
			Meta:         data.Jsonb(fmt.Sprintf(`{"question_id": %d}`, question.ID)),
		})
		if err != nil {
			return fmt.Errorf("failed to insert event: %w", err)
		}

		level, err := DoLevelRefUpgrade(Levels(r), ReferralsQ(r), balance, question.Reward)
		if err != nil {
			return fmt.Errorf("failed to do lvlup and referrals updates: %w", err)
		}

		err = BalancesQ(r).FilterByNullifier(balance.Nullifier).Update(map[string]any{
			data.ColAmount: pg.AddToValue(data.ColAmount, question.Reward),
			data.ColLevel:  level,
		})
		if err != nil {
			return fmt.Errorf("update balance amount and level: %w", err)
		}

		return nil
	})
	if err != nil {
		Log(r).WithError(err).Error("Failed to insert execute transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newDailyAnswer(question))
}

func newDailyAnswer(question *data.DailyQuestion) resources.DailyQuestionAnswers {
	return resources.DailyQuestionAnswers{
		Key: resources.NewKeyInt64(question.ID, resources.DAILY_QUESTIONS),
		Attributes: resources.DailyQuestionAnswersAttributes{
			Answer: question.CorrectAnswer,
		},
	}
}
