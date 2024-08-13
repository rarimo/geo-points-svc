package handlers

import (
	"math"
	"net/http"
	"time"

	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetDailyQuestionsStatus(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetDailyQuestionsStatus(r)
	if err != nil {
		Log(r).Errorf("error getting status daily questions: %v", err)
	}

	//TODO come back authentication

	activeQuestions, err := DailyQuestionsQ(r).FilteredActive(true).Select()

	if err != nil {
		Log(r).Errorf("error getting active questions: %v", err)
	}

	if len(activeQuestions) == 0 {
		req.Availability = false
		req.AlreadyDoneForUser = false

		closes, err := TimeToNExtQuestion(r)

		if err != nil || closes == -2 {
			Log(r).Errorf("error getting time to next question: %v", err)
			ape.RenderErr(w, problems.InternalError())
			return
		}

		req.TimeToNext = int64(closes)
		ape.Render(w, NewDailyQuestionsStatus(req))
		return
	}

	req.Availability = true

	checkPassed, err := checkPassedUserQuestion(r) //TODO func for this
	if err != nil {
		Log(r).Errorf("error checkPassedUserQuestion: %v", err)
	}
	req.AlreadyDoneForUser = checkPassed

	if !checkPassed {
		closes, err := TimeToNExtQuestion(r)
		if err != nil || req.TimeToNext == -2 {
			Log(r).Errorf("error getting time to next question: %v", err)
		}
		req.TimeToNext = int64(closes)
	} else {
		req.TimeToNext = 0
	}

	ape.Render(w, NewDailyQuestionsStatus(req))
}

func checkPassedUserQuestion(r *http.Request) (bool, error) { //TODO
	return true, nil
}

func TimeToNExtQuestion(r *http.Request) (int, error) {
	location, err := time.LoadLocation("Asia/Tbilisi")
	if err != nil {
		Log(r).Errorf("Error load georgia location: %v", err)
		return -2, err
	}

	activeQuestions, err := DailyQuestionsQ(r).Select()
	if err != nil {
		Log(r).Errorf("Error getting all questions: %v", err)
		return -2, err
	}

	if len(activeQuestions) == 0 {
		Log(r).Error("No active questions found")
		return -1, nil
	}

	nowTime := int(time.Now().In(location).Unix())
	questions, err := DailyQuestionsQ(r).FilteredStartAt(nowTime).Select()

	if err != nil {
		Log(r).Errorf("Error getting all questions: %v", err)
		return -2, nil
	}
	if len(questions) == 0 {
		Log(r).Error("No questions found")
		return -1, nil
	}

	closes := math.MaxInt64

	for _, q := range questions {
		if q.StartsAt < closes {
			closes = q.StartsAt
		}
	}
	return closes - nowTime, nil
}

func NewDailyQuestionsStatus(req requests.GetDailyQuestionsStatus) resources.DailyQuestionsStatusAttributes {
	return resources.DailyQuestionsStatusAttributes{
		AlreadyDoneForUser: req.AlreadyDoneForUser,
		Availability:       req.Availability,
		TimeToNext:         req.TimeToNext,
	}

}
