package handlers

import (
	"fmt"
	"net/http"

	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DailyQuestionAnswer(w http.ResponseWriter, r *http.Request) {
	//if !auth.Authenticates(UserClaims(r), auth.UserGrant(req.UserNullifier)) {
	//	ape.RenderErr(w, problems.Unauthorized())
	//	return
	//}

	answerIsTrue := true

	req, err := requests.NewAnswerDailyQuestionData(r)
	if err != nil {
		Log(r).Errorf("error getting active questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	question, err := DailyQuestionsQ(r).FilterByID(req.QuestionID).Get()
	if err != nil {
		Log(r).Errorf("error getting question: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	balance, err := BalancesQ(r).FilterByNullifier(req.Nullifier).Get()
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = EventsQ(r).Transaction(func() error {
		evType := EventTypes(r).Get(models.TypeDailyQuestion, evtypes.FilterInactive)
		if evType == nil {
			Log(r).Infof("Event type %s is inactive", models.TypeEarlyTest)
			return problems.InternalError()
		}

		eventCheck, err := EventsQ(r).FilterByNullifier(req.Nullifier).FilterByType(models.TypeDailyQuestion).Get()
		if err != nil {
			Log(r).WithError(err).Errorf("Failed to get balance by nullifier %s", req.Nullifier)
		}
		if eventCheck == nil {
			Log(r).Infof("User haven`t question %s", req.Nullifier)
			return problems.NotFound()
		}

		if eventCheck.Status != data.EventOpen {
			Log(r).Infof("User already answered the question nullifier: %s , event status: %s", req.Nullifier, eventCheck.Status)
			return problems.NotAllowed()
		}

		answersOptions := ConvertJsonbToDailyQuestionAnswers(question.AnswerOptions)
		if answersOptions == nil {
			Log(r).Infof("Question haven`t answers, or error JSON data answer id: %v", req.QuestionID)
			return problems.InternalError()
		}

		_, err = EventsQ(r).FilterByNullifier(req.Nullifier).
			FilterByType(models.TypeDailyQuestion).
			Update(data.EventFulfilled, nil, nil)

		if err != nil {
			return fmt.Errorf("failed to fulfill daily question event (answer post): %w", err)
		}

		if answersOptions[req.Answer] != true {
			answerIsTrue = false
			Log(r).Infof("Anser wrong")
			return nil
		}

		//if balance.IsDisabled() {
		//	Log(r).Debug("Balance is disabled, events will not be claimed")
		//	return nil
		//}

		if err := addEventForReferrer(r, balance); err != nil {
			return fmt.Errorf("add event for referrer: %w", err)
		}

		err = autoClaimDailyQuestionEventsForBalance(r, balance, question)
		if err != nil {
			Log(r).WithError(err).Errorf("Failed to claim daily question events for balance")
			return problems.InternalError()
		}

		return nil
	})

	if err != nil {
		Log(r).WithError(err).Error("Failed to transaction event getDailyQuestion")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	TimeToNext, err := TimeToNextQuestion(r)
	ape.Render(w, NewDailyAnswer(answerIsTrue, TimeToNext))
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

	if balance.IsDisabled() || !balance.IsVerified() {
		Log(r).Debug("User not eligible for event claiming. Events not claimed.")
		return nil
	}

	eventsToClaim, err := EventsQ(r).
		FilterByNullifier(balance.Nullifier).
		FilterByStatus(data.EventFulfilled).
		Select()
	if err != nil {
		return fmt.Errorf("failed to select events for user=%s: %w", balance.Nullifier, err)
	}

	var totalPoints int64
	eventsMap := map[string][]string{}
	for _, e := range eventsToClaim {
		if _, ok := eventsMap[e.Type]; !ok {
			eventsMap[e.Type] = []string{}
		}
		eventsMap[e.Type] = append(eventsMap[e.Type], e.ID)
	}

	for evName, evIDs := range eventsMap {
		evType := EventTypes(r).Get(evName, evtypes.FilterInactive, evtypes.FilterByAutoClaim(true))
		if evType == nil {
			continue
		}

		_, err = EventsQ(r).FilterByID(evIDs...).Update(data.EventClaimed, nil, &evType.Reward)
		if err != nil {
			return fmt.Errorf("failedt to update %s events for user=%s: %w", evName, balance.Nullifier, err)
		}

		totalPoints = balance.Amount + int64(question.Reward)
	}

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
