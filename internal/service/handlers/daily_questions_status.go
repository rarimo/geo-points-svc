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
	alreadyDoneForUser := false
	var timeToNext int64
	req, err := requests.NewDailyQuestionUserAccess(r)
	if err != nil {
		Log(r).WithError(err).Error("error getting daily questions status")
		ape.RenderErr(w, problems.InternalError())
	}

	dailyQuestionEvent, err := EventsQ(r).
		FilterByNullifier(req.Nullifier).
		FilterTodayEvents(req.Timezone).
		Get()

	if err != nil {
		Log(r).WithError(err).Error("error getting daily questions status")
		ape.RenderErr(w, problems.InternalError())
	}

	location, err := time.LoadLocation(req.Timezone)
	if err != nil {
		location = time.UTC
	}

	if dailyQuestionEvent != nil {
		alreadyDoneForUser = true
		timeToNext, err = TimeToNextQuestion(r, location)
		ape.Render(w, NewDailyQuestionsStatus(alreadyDoneForUser, timeToNext))
	} else {
		timeToNext = 0
	}

}

func TimeToNextQuestion(r *http.Request, loc *time.Location) (int64, error) {
	now := time.Now().In(loc)
	questions, err := DailyQuestionsQ(r).FilterByStartAt(now).Select()
	if err != nil {
		return -2, err
	}

	if len(questions) == 0 {
		return -1, nil
	}

	closes := int64(math.MaxInt64)
	for _, q := range questions {
		timeToNext := q.StartsAt.Unix() - now.Unix()
		if timeToNext < closes {
			closes = timeToNext
		}
	}

	return closes, nil
}

func NewDailyQuestionsStatus(AlreadyDoneForUser bool, TimeToNext int64) resources.DailyQuestionStatusAttributes {
	return resources.DailyQuestionStatusAttributes{
		AlreadyDoneForUser: AlreadyDoneForUser,
		TimeToNext:         TimeToNext,
	}
}
