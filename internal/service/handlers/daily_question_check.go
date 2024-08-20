package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
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

	req, err := requests.NewDailyQuestionAnswer(r)
	if err != nil {
		Log(r).Errorf("error getting active questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	Log(r).Infof("A")
	if !auth.Authenticates(UserClaims(r), auth.UserGrant(req.Nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	cell := DailyQuestionTimeHash(r).GetDailyQuestionsTimeHash(req.Nullifier)
	if cell == nil {
		Log(r).Errorf("The user's nullifier was not found in active requests, it does not exist, or the user has already answered: %s", req.Nullifier)
		ape.RenderErr(w, problems.NotAllowed())
		return
	}
	if cell.MaxDateToAnswer < time.Now().UTC().Unix() {
		Log(r).Infof("Time is up :%s", req.Nullifier)
		ape.RenderErr(w, problems.Forbidden())
		return
	}
	if cell.Answered {
		Log(r).Infof("User has already answered: %s", req.Nullifier)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	location := Location(r).String()

	question, err := DailyQuestionsQ(r).FilterTodayQuestions(location).Get()
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

		eventCheck, err := EventsQ(r).
			FilterByNullifier(req.Nullifier).
			FilterByType(models.TypeDailyQuestion).
			FilterTodayEvents(location).
			Get()

		if err != nil {
			return err
		}
		if eventCheck != nil {
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

		cell.Answered = true

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

	TimeToNext, err := TimeToNextQuestionUnix(r)
	if err != nil {
		Log(r).WithError(err).Errorf("error updating daily questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, NewDailyAnswer(answerIsTrue, int(TimeToNext)))
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

func NewDailyAnswer(answerStatus bool, TimeToNext int) resources.DailyQuestionResultAttributes {
	return resources.DailyQuestionResultAttributes{
		AnswerStatus: answerStatus,
		TimeToNext:   FormatUnixTimeToDate(int64(TimeToNext)),
	}
}

func TimeToNextQuestionUnix(r *http.Request) (int64, error) {
	questions, err := DailyQuestionsQ(r).FilterByStartAtToday(Location(r).String()).Select()
	if err != nil {
		return -2, err
	}

	var futureTimes []int64
	now := time.Now().Unix()

	for _, q := range questions {
		timeToNext := q.StartsAt.Unix() - now
		if timeToNext > 0 { // Учитываем только будущее время
			futureTimes = append(futureTimes, timeToNext)
		}
	}

	if len(futureTimes) == 0 {
		return -1, nil // Нет предстоящих вопросов
	}

	// Сортируем по возрастанию
	sort.Slice(futureTimes, func(i, j int) bool {
		return futureTimes[i] < futureTimes[j]
	})

	if len(futureTimes) > 1 {
		return futureTimes[1], nil // Возвращаем второе минимальное значение
	}

	return -1, nil // Если второго элемента нет, возвращаем -1
}
