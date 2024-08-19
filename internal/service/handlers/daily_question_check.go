package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

	req, err := requests.NewDailyQuestionAnswer(r)
	if err != nil {
		Log(r).Errorf("error getting active questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	question, err := DailyQuestionsQ(r).FilterTodayQuestions(req.Timezone).Get()

	if question == nil {
		Log(r).Errorf("error getting question: %v", err)
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if err != nil {
		Log(r).Errorf("error getting question or quesrion inactive: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	balance, err := BalancesQ(r).FilterByNullifier(req.Nullifier).Get()
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	answersMap, err := JSONBToMap(question.AnswerOptions)
	if err != nil {
		Log(r).WithError(err).Errorf("error converting answer options to map")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = EventsQ(r).Transaction(func() error {
		evType := EventTypes(r).Get(models.TypeDailyQuestion, evtypes.FilterInactive)
		if evType == nil {
			return fmt.Errorf("event type %s is inactive", models.TypeDailyQuestion)
		}

		eventCheck, err := EventsQ(r).
			FilterByNullifier(req.Nullifier).
			FilterByType(models.TypeDailyQuestion).
			FilterTodayEvents(req.Timezone).
			Get()

		if eventCheck != nil {
			return problems.Forbidden()
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

		if answersMap[req.UserAnswer] != true {
			answerIsTrue = false
			Log(r).Infof("Anser wrong")
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

	location, err := time.LoadLocation(req.Timezone)
	if err != nil {
		location = time.UTC
	}

	TimeToNext, err := TimeToNextQuestion(r, location)
	ape.Render(w, NewDailyAnswer(answerIsTrue, int(TimeToNext)))
}

func JSONBToMap(data data.Jsonb) (map[string]interface{}, error) {
	var result map[string]interface{}

	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSONB: %w", err)
	}

	return result, nil
}

func NewDailyAnswer(answerStatus bool, TimeToNext int) resources.DailyQuestionResultAttributes {
	return resources.DailyQuestionResultAttributes{
		AnswerStatus: answerStatus,
		TimeToNext:   int64(TimeToNext),
	}
}
