package handlers

import (
	"math"
	"net/http"
	"time"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
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

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(*req.Nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	dailyQuestionEvent, err := EventsQ(r).
		FilterByNullifier(*req.Nullifier).
		FilterTodayEvents(req.Timezone).
		FilterByType(models.TypeDailyQuestion).
		Get()

	if err != nil {
		Log(r).WithError(err).Error("error getting event daily_question")
		ape.RenderErr(w, problems.InternalError())
	}

	location := GetLocationFromTimezone(req.Timezone)

	if dailyQuestionEvent != nil {
		alreadyDoneForUser = true
		timeToNext, err = TimeToNextQuestion(r, location)
		ape.Render(w, NewDailyQuestionsStatus(alreadyDoneForUser, timeToNext))
		return
	}
	timeToNext = 0
	ape.Render(w, NewDailyQuestionsStatus(alreadyDoneForUser, timeToNext))
	return
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
