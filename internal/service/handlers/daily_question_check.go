package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
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
	cfg := DailyQuestions(r)
	nullifier := strings.ToLower(chi.URLParam(r, "nullifier"))

	req, err := requests.NewDailyQuestionAnswer(r)
	if err != nil {
		Log(r).Errorf("Error getting active questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	//if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
	//	ape.RenderErr(w, problems.Unauthorized())
	//	return
	//}

	question, err := DailyQuestionsQ(r).FilterTodayQuestions(cfg.Timezone).Get()
	if err != nil {
		Log(r).Errorf("Error getting question or quesrion inactive: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if question == nil {
		Log(r).Error("Error getting question, question is nil")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	balance, err := BalancesQ(r).FilterByNullifier(nullifier).Get()
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to get balance by nullifier %s", nullifier)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if balance == nil {
		Log(r).Errorf("Error getting balance by nullifier")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if cfg.ResponderExists(nullifier) {
		Log(r).Infof("User is already answered %s", nullifier)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	deadline := cfg.GetDeadline(nullifier)
	if deadline == nil {
		Log(r).Errorf("The user's nullifier was not found in active requests, it does not exist, or the user has already answered: %s", nullifier)
		ape.RenderErr(w, problems.Forbidden())
		return
	}
	if *deadline < time.Now().UTC().Unix() {
		Log(r).Errorf("Time is up: %s", nullifier)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	errDuplicateEvent := fmt.Errorf("user already send answer %s", nullifier)

	err = EventsQ(r).Transaction(func() error {
		evType := EventTypes(r).Get(models.TypeDailyQuestion, evtypes.FilterInactive)
		if evType == nil {
			return fmt.Errorf("event type %s is inactive", models.TypeDailyQuestion)
		}

		event, err := EventsQ(r).
			FilterByNullifier(nullifier).
			FilterByType(models.TypeDailyQuestion).
			FilterTodayEvents(cfg.Timezone).
			Get()

		if err != nil {
			return err
		}
		if event != nil {
			return errDuplicateEvent
		}

		if question.CorrectAnswer != req.Answer {
			Log(r).Infof("Wrong answer for daily question: %v", req.Answer)
			err = DailyQuestionsQ(r).FilterTodayQuestions(cfg.Timezone).IncrementIncorrectAnswer()
			if err != nil {
				Log(r).WithError(err).Errorf("Error incrementing question answers incorect answered")
			}
			return nil
		}

		question.NumCorrectAnswers++
		err = DailyQuestionsQ(r).FilterTodayQuestions(cfg.Timezone).IncrementCorrectAnswer()
		if err != nil {
			Log(r).WithError(err).Errorf("Error incrementing question answers corect answered")
		}

		err = EventsQ(r).Insert(data.Event{
			Nullifier: nullifier,
			Type:      models.TypeDailyQuestion,
			Status:    data.EventFulfilled,
		})

		if err != nil {
			return fmt.Errorf("error insert event: %w", err)
		}

		if evtypes.FilterByAutoClaim(true)(*evType) {
			return nil
		}

		level, err := DoLevelRefUpgrade(Levels(r), ReferralsQ(r), balance, evType.Reward)
		if err != nil {
			return fmt.Errorf("failed to do lvlup and referrals updates: %w", err)
		}

		err = BalancesQ(r).New().FilterByNullifier(balance.Nullifier).Update(map[string]any{
			data.ColAmount: pg.AddToValue(data.ColAmount, int64(question.Reward)),
			data.ColLevel:  level,
		})

		if err != nil {
			return fmt.Errorf("error update balance amount and level: %w", err)
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, errDuplicateEvent) {
			Log(r).Infof("User already submitted an answer: %s", nullifier)
			ape.RenderErr(w, problems.Conflict())
			return
		}

		Log(r).WithError(err).Errorf("Error updating daily questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	cfg.SetResponsesTimer(balance.Nullifier, time.Duration(*deadline-time.Now().UTC().Unix())*time.Second)
	ape.Render(w, NewDailyAnswer(question.ID, nullifier))
	return
}

func NewDailyAnswer(answerIndex int64, nullifier string) resources.DailyQuestionAnswers {
	return resources.DailyQuestionAnswers{
		Key: resources.Key{
			ID:   nullifier,
			Type: resources.DAILY_QUESTIONS,
		},
		Attributes: resources.DailyQuestionAnswersAttributes{
			Answer: answerIndex,
		},
	}
}
