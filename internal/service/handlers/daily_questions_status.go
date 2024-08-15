package handlers

import (
	"math"
	"net/http"

	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetDailyQuestionsStatus(w http.ResponseWriter, r *http.Request) {
	AlreadyDoneForUser, TimeToNext := false, 0

	// if !auth.Authenticates(UserClaims(r), auth.UserGrant(req.UserNullifier)) {
	//	ape.RenderErr(w, problems.Unauthorized())
	//	return
	//}

	req, err := requests.NewGetDailyQuestionsRequest(r)
	if err != nil {
		Log(r).Errorf("error getting status daily questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	activeQuestions, err := DailyQuestionsQ(r).FilterByActive(true).Select()
	Log(r).Infof("len: %v", len(activeQuestions))
	if err != nil {
		Log(r).Errorf("error getting active questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if len(activeQuestions) == 0 {
		AlreadyDoneForUser = false

		closes, err := TimeToNextQuestion(r)

		if err != nil || closes == -1 {
			Log(r).Errorf("error getting time to next question: %v", err)
			ape.RenderErr(w, problems.InternalError())
			return
		}

		TimeToNext = closes
		ape.Render(w, NewDailyQuestionsStatus(AlreadyDoneForUser, TimeToNext))
		return
	}

	checkPassed, err := checkPassedUserQuestion(r, req)
	if err != nil {
		Log(r).Errorf("error checkPassedUserQuestion: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	AlreadyDoneForUser = checkPassed

	if !checkPassed {
		closes, err := TimeToNextQuestion(r)
		if err != nil || closes != -2 {
			Log(r).Errorf("error getting time to next question: %v", err)
			ape.RenderErr(w, problems.InternalError())
			return
		}
		TimeToNext = closes
	} else {
		TimeToNext = 0
	}

	ape.Render(w, NewDailyQuestionsStatus(AlreadyDoneForUser, TimeToNext))
}

func TimeToNextQuestion(r *http.Request) (int, error) {
	activeQuestions, err := DailyQuestionsQ(r).Select()
	if err != nil {
		Log(r).Errorf("Error getting all questions: %v", err)
		return -2, err
	}

	if len(activeQuestions) == 0 {
		Log(r).Error("No active questions found")
		return -1, nil
	}
	nowTime := config.NowTime()
	questions, err := DailyQuestionsQ(r).FilterByStartAt(nowTime).Select()

	if err != nil {
		Log(r).Errorf("Error getting all questions: %v", err)
		return -2, err
	}
	if len(questions) == 0 {
		Log(r).Error("No questions found")
		return -1, nil
	}

	closes := math.MaxInt64

	for _, q := range questions {
		if int(q.StartsAt.Unix()) < closes {
			closes = int(q.StartsAt.Unix())
		}
	}
	return closes - int(nowTime.Unix()), nil
}

func checkPassedUserQuestion(r *http.Request, req requests.GetDailyQuestionRequest) (bool, error) {
	eve, err := EventsQ(r).FilterByNullifier(req.UserNullifier).FilterByType(models.TypeDailyQuestion).Get()
	if err != nil {
		Log(r).Errorf("Error getting events: %v", err)
		return false, err
	}

	if eve == nil {
		return true, nil
	}

	if eve.Status != data.EventOpen {
		return false, nil
	}

	return true, nil
}

func NewDailyQuestionsStatus(AlreadyDoneForUser bool, TimeToNext int) resources.DailyQuestionsStatusAttributes {
	return resources.DailyQuestionsStatusAttributes{
		AlreadyDoneForUser: AlreadyDoneForUser,
		TimeToNext:         int64(TimeToNext),
	}

}
