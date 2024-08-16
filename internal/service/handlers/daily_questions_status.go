package handlers

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetDailyQuestionsStatus(w http.ResponseWriter, r *http.Request) {
	var AlreadyDoneForUser bool
	var TimeToNext int64

	req, err := requests.NewGetDailyQuestionsRequest(r)
	if err != nil {
		Log(r).Errorf("error getting status daily questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// Закомментировано, т.к. аутентификация временно отключена
	// if !auth.Authenticates(UserClaims(r), auth.UserGrant(req.Nullifier)) {
	// 	ape.RenderErr(w, problems.Unauthorized())
	// 	return
	// }

	currentTime := time.Now().UTC()

	activeQuestions, err := DailyQuestionsQ(r).FilterActive().Select()
	if err != nil {
		Log(r).Errorf("error getting active questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	Log(r).Infof("activeQuestions: %v", len(activeQuestions))

	if len(activeQuestions) == 0 {
		AlreadyDoneForUser = false
		TimeToNext, err = TimeToNextQuestion(r, currentTime)

		if err != nil || TimeToNext == -2 {
			Log(r).Errorf("error getting time to next question: %v", err)
			ape.RenderErr(w, problems.InternalError())
			return
		}

		ape.Render(w, NewDailyQuestionsStatus(AlreadyDoneForUser, TimeToNext))
		return
	}

	TimeToNext, err = TimeToNextQuestion(r, currentTime)

	if err != nil || TimeToNext == -2 {
		Log(r).Errorf("error getting time to next question: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	for _, quest := range activeQuestions {
		AlreadyDoneForUser, err = checkPassedUserQuestion(r, req.Nullifier, quest)
		if err != nil {
			Log(r).Errorf("error checkPassedUserQuestion: %v", err)
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if AlreadyDoneForUser {
			break
		}
	}

	if !AlreadyDoneForUser {
		TimeToNext = 0
	}

	ape.Render(w, NewDailyQuestionsStatus(AlreadyDoneForUser, TimeToNext))
}

func TimeToNextQuestion(r *http.Request, currentTime time.Time) (int64, error) {
	questions, err := DailyQuestionsQ(r).FilterByStartAt(currentTime).Select()
	if err != nil {
		return -2, err
	}

	if len(questions) == 0 {
		return -1, nil
	}

	closes := int64(math.MaxInt64)
	for _, q := range questions {
		timeToNext := q.StartsAt.Unix() - currentTime.Unix()
		if timeToNext < closes {
			closes = timeToNext
		}
	}

	return closes, nil
}

func checkPassedUserQuestion(r *http.Request, nullifier string, question data.DailyQuestion) (bool, error) {
	event, err := EventsQ(r).
		FilterByNullifier(nullifier).
		FilterByType(models.TypeDailyQuestion).
		FilterByStatus(data.EventFulfilled).
		FilterByCreatedAtAfter(question.StartsAt.Unix()).
		FilterByCreatedAtBefore(question.StartsAt.Unix() + question.TimeForAnswer).
		Get()

	Log(r).Infof("%+v", event)

	if err != nil {
		Log(r).Errorf("Error getting events: %v", err)
		return false, fmt.Errorf("failed to select events for user=%s: %w", nullifier, err)
	}
	if event == nil {
		return false, nil
	}

	return true, nil
}

func NewDailyQuestionsStatus(AlreadyDoneForUser bool, TimeToNext int64) resources.DailyQuestionsStatusAttributes {
	return resources.DailyQuestionsStatusAttributes{
		AlreadyDoneForUser: AlreadyDoneForUser,
		TimeToNext:         TimeToNext,
	}
}
