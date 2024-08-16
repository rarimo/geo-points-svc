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

func DailyQuestionCheck(w http.ResponseWriter, r *http.Request) {
	answerIsTrue := true
	currentTime := time.Now().UTC()

	req, err := requests.NewAnswerDailyQuestionData(r)
	if err != nil {
		Log(r).Errorf("error getting active questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	//if !auth.Authenticates(UserClaims(r), auth.UserGrant(req.Nullifier)) {
	//	ape.RenderErr(w, problems.Unauthorized())
	//	return
	//}

	question, err := DailyQuestionsQ(r).FilterByID(req.QuestionID).FilterActive().Get()
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

	questionData, err := DailyQuestionsResponsesQ(r).FilterByQuestionID(question.ID).FilterByNullifier(req.Nullifier).Get()
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if questionData == nil {
		Log(r).Errorf("error getting question data")
		ape.RenderErr(w, problems.Forbidden())
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
			FilterByCreatedAtAfter(question.CreatedAt.Unix()).
			FilterByUpdatedAtBefore(currentTime.Unix()).
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

		if answersMap[req.Answer] != true {
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

	TimeToNext, err := TimeToNextQuestion(r, time.Now())
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

func autoClaimDailyQuestionEventsForBalance(r *http.Request, balance *data.Balance, question *data.DailyQuestion) error {
	if balance == nil {
		Log(r).Debug("Balance absent. Events not claimed.")
		return nil
	}

	//if balance.IsDisabled() || !balance.IsVerified() {
	//	Log(r).Debug("User not eligible for event claiming. Events not claimed.")
	//	return nil
	//}

	event, err := EventsQ(r).
		FilterByNullifier(balance.Nullifier).
		FilterByType(models.TypeDailyQuestion).
		FilterByStatus(data.EventFulfilled).
		FilterByCreatedAtAfter(question.StartsAt.Unix()).
		FilterByCreatedAtBefore(question.StartsAt.Unix() + question.TimeForAnswer).
		Get()

	if err != nil {
		return fmt.Errorf("failed to select events for user=%s: %w", balance.Nullifier, err)
	}
	if event != nil {
		return fmt.Errorf("user already claimed events for user=%s", balance.Nullifier)
	}

	totalPoints := balance.Amount + int64(question.Reward)

	level, err := DoLevelRefUpgrade(Levels(r), ReferralsQ(r), balance, totalPoints)
	if err != nil {
		return fmt.Errorf("failed to do lvlup and referrals updates: %w", err)
	}

	err = BalancesQ(r).FilterByNullifier(balance.Nullifier).Update(map[string]any{
		data.ColAmount: pg.AddToValue(data.ColAmount, totalPoints),
		data.ColLevel:  level,
	})
	if err != nil {
		return fmt.Errorf("update balance amount and level: %w", err)
	}

	return nil
}
