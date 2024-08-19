package handlers

import (
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetDailyQuestionsStatus(w http.ResponseWriter, r *http.Request) {
	alreadyDoneForUser := false
	var timeToNext int64
	nullifier := strings.ToLower(chi.URLParam(r, "nullifier"))

	//if !auth.Authenticates(UserClaims(r), auth.UserGrant(*req.Nullifier)) {
	//	ape.RenderErr(w, problems.Unauthorized())
	//	return
	//}

	balance, err := BalancesQ(r).FilterByNullifier(nullifier).Get()
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to get balance by nullifier %v", nullifier)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if balance == nil {
		Log(r).Errorf("error getting balance by nullifier")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	dailyQuestionEvent, err := EventsQ(r).
		FilterByNullifier(balance.Nullifier).
		FilterTodayEvents().
		FilterByType(models.TypeDailyQuestion).
		Get()

	if err != nil {
		Log(r).WithError(err).Error("error getting event daily_question")
		ape.RenderErr(w, problems.InternalError())
	}
	if dailyQuestionEvent != nil {
		alreadyDoneForUser = true
		timeToNext, err = TimeToNextQuestion(r)
		ape.Render(w, NewDailyQuestionsStatus(alreadyDoneForUser, timeToNext))
		return
	}

	timeToNext = 0
	ape.Render(w, NewDailyQuestionsStatus(alreadyDoneForUser, timeToNext))
	return
}

func TimeToNextQuestion(r *http.Request) (int64, error) {
	now := time.Now()
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
