package handlers

import (
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetDailyQuestionsStatus(w http.ResponseWriter, r *http.Request) {
	var timeToNext time.Time
	cfg := DailyQuestions(r)
	nullifier := strings.ToLower(chi.URLParam(r, "nullifier"))

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

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
		FilterTodayEvents(cfg.Timezone).
		FilterByType(models.TypeDailyQuestion).
		Get()

	if err != nil {
		Log(r).WithError(err).Error("error getting event daily_question")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	quesDataArr, err := GetQuestionQueue(r)
	if err != nil || len(quesDataArr) == 0 {
		Log(r).WithError(err).Error("error getting time to next question")
		timeToNext = time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
	} else {
		timeToNext = quesDataArr[0]
	}

	if dailyQuestionEvent != nil {
		ape.Render(w, NewDailyQuestionsStatus(true, timeToNext))
		return
	}

	ape.Render(w, NewDailyQuestionsStatus(false, time.Now()))
	return
}

func GetQuestionQueue(r *http.Request) ([]time.Time, error) {
	questions, err := DailyQuestionsQ(r).
		FilterTodayQuestions(DailyQuestions(r).Timezone).
		Select()

	if err != nil {
		return nil, err
	}

	if len(questions) == 0 {
		return nil, nil
	}

	times := make([]time.Time, len(questions))
	for i, q := range questions {
		times[i] = q.StartsAt
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].Before(times[j])
	})

	return times, nil
}

func NewDailyQuestionsStatus(AlreadyDoneForUser bool, timeToNextStr time.Time) resources.DailyQuestionStatusAttributes {
	return resources.DailyQuestionStatusAttributes{
		AlreadyDoneForUser: AlreadyDoneForUser,
		TimeToNext:         timeToNextStr.String(),
	}
}
