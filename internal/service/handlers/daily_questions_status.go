package handlers

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetDailyQuestionsStatus(w http.ResponseWriter, r *http.Request) {
	var timeToNext int64
	cfg := DailyQuestions(r)
	nullifier := strings.ToLower(chi.URLParam(r, "nullifier"))

	//if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
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
		Log(r).Errorf("Error getting balance by nullifier %v", nullifier)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	timeToNext, err = TimeToNextQuestion(r)
	if err != nil {
		Log(r).WithError(err).Error("Error getting time to next question")
		timeToNext = -1
	}

	question, err := DailyQuestionsQ(r).
		FilterTodayQuestions(cfg.Timezone).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("Error getting daily question")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if question == nil {
		Log(r).Infof("Not found question today %v", question)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, NewDailyQuestionsStatus(timeToNext, question))
	return
}

func TimeToNextQuestion(r *http.Request) (int64, error) {
	questions, err := DailyQuestionsQ(r).
		FilterTodayQuestions(DailyQuestions(r).Timezone).
		Select()

	if err != nil {
		return -1, err
	}

	if len(questions) == 0 {
		return -1, nil
	}

	times := make([]time.Time, len(questions))
	for i, q := range questions {
		times[i] = q.StartsAt
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].Before(times[j])
	})

	now := time.Now().UTC()

	for _, t := range times {
		if t.After(now) {
			return t.Unix(), nil
		}
	}

	return -1, nil
}

func NewDailyQuestionsStatus(timeToNextStr int64, question *data.DailyQuestion) resources.DailyQuestionsStatus {
	return resources.DailyQuestionsStatus{
		Key: resources.Key{
			ID:   strconv.Itoa(int(question.ID)),
			Type: resources.DAILY_QUESTIONS,
		},
		Attributes: resources.DailyQuestionsStatusAttributes{
			NextQuestionDate: timeToNextStr,
			Reward:           int64(question.Reward),
			TimeForAnswer:    question.TimeForAnswer,
		},
	}
}
