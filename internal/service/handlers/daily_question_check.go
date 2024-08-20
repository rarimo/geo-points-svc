package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

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
	answerIsTrue := true
	cfg := DailyQuestions(r)

	req, err := requests.NewDailyQuestionAnswer(r)
	if err != nil {
		Log(r).Errorf("error getting active questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(req.Nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	question, err := DailyQuestionsQ(r).FilterTodayQuestions(cfg.Timezone).Get()
	if err != nil {
		Log(r).Errorf("error getting question or quesrion inactive: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if question == nil {
		Log(r).Errorf("error getting question: %v", err)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	deadline := cfg.GetFromQuestionsQueue(req.Nullifier)
	if deadline == nil {
		Log(r).Errorf("The user's nullifier was not found in active requests, it does not exist, or the user has already answered: %s", req.Nullifier)
		ape.RenderErr(w, problems.NotAllowed())
		return
	}
	if *deadline < time.Now().UTC().Unix() {
		Log(r).Errorf("Time is up: %s", req.Nullifier)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	balance, err := BalancesQ(r).FilterByNullifier(req.Nullifier).Get()
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to get balance by nullifier %s", req.Nullifier)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if balance == nil {
		Log(r).Errorf("error getting balance by nullifier")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	answersMap, err := JSONBToMap(question.AnswerOptions)
	if err != nil {
		Log(r).WithError(err).Errorf("error converting answer options to map")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	errDuplicateEvent := fmt.Errorf("user already send answer %s", req.Nullifier)

	err = EventsQ(r).Transaction(func() error {
		evType := EventTypes(r).Get(models.TypeDailyQuestion, evtypes.FilterInactive)
		if evType == nil {
			return fmt.Errorf("event type %s is inactive", models.TypeDailyQuestion)
		}

		if answersMap[req.UserAnswer] != true {
			answerIsTrue = false
			Log(r).Infof("User %s anser wrong", req.UserAnswer)
			return nil
		}

		event, err := EventsQ(r).
			FilterByNullifier(req.Nullifier).
			FilterByType(models.TypeDailyQuestion).
			FilterTodayEvents(cfg.Timezone).
			Get()

		if err != nil {
			return err
		}
		if event != nil {
			return errDuplicateEvent
		}

		err = EventsQ(r).Insert(data.Event{
			Nullifier: req.Nullifier,
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
			Log(r).Infof("User already submitted an answer: %s", req.Nullifier)
			ape.RenderErr(w, problems.Conflict())
			return
		}

		Log(r).WithError(err).Errorf("error updating daily questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	TimeToNext, err := GetQuestionQueue(r)
	if err != nil {
		Log(r).WithError(err).Errorf("error updating daily questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if len(TimeToNext) < 2 {
		ape.Render(w, NewDailyAnswer(answerIsTrue, time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)))
		return
	}

	ape.Render(w, NewDailyAnswer(answerIsTrue, TimeToNext[1]))
	return
}

func JSONBToMap(data data.Jsonb) (map[string]interface{}, error) {
	var result map[string]interface{}

	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSONB: %w", err)
	}

	return result, nil
}

func NewDailyAnswer(answerStatus bool, TimeToNext time.Time) resources.DailyQuestionResultAttributes {
	return resources.DailyQuestionResultAttributes{
		AnswerStatus: answerStatus,
		TimeToNext:   TimeToNext.String(),
	}
}
