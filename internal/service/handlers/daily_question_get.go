package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

var mu sync.Mutex

func GetDailyQuestion(w http.ResponseWriter, r *http.Request) {
	nullifier := strings.ToLower(chi.URLParam(r, "nullifier"))
	cfg := DailyQuestions(r)

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	balance, err := BalancesQ(r).FilterByNullifier(nullifier).Get()
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if balance == nil {
		Log(r).Errorf("error getting balance by nullifier")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	question, err := DailyQuestionsQ(r).
		FilterTodayQuestions(cfg.Timezone).
		Get()
	if question == nil {
		Log(r).Errorf("error getting daily question")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if err != nil {
		Log(r).WithError(err).Error("Failed to get active questions")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	questionEvent, err := EventsQ(r).
		FilterTodayEvents(cfg.Timezone).
		FilterByType(models.TypeDailyQuestion).
		FilterByNullifier(nullifier).
		Get()

	if err != nil {
		Log(r).WithError(err).Error("Failed to get active questions")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if questionEvent != nil {
		Log(r).Infof("User already answered %s", nullifier)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	deadline := cfg.GetFromQuestionsQueue(nullifier)
	if deadline != nil {
		Log(r).Errorf("The user's nullifier was not found in active requests, it does not exist, or the user has already answered: %s", nullifier)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	cfg.SetDailyQuestionTimeWithExpiration(questionEvent, balance.Nullifier, time.Now().UTC().Unix()+question.TimeForAnswer)

	ape.Render(w, NewDailyQuestion(*question))
	return
}

func NewDailyQuestion(question data.DailyQuestion) resources.DailyQuestion {
	var q map[string]interface{}
	_ = json.Unmarshal(question.AnswerOptions, &q)

	return resources.DailyQuestion{
		Key: resources.Key{
			ID:   strconv.Itoa(question.ID),
			Type: resources.DAILY_QUESTION,
		},
		Attributes: resources.DailyQuestionAttributes{
			Title:         question.Title,
			Reward:        question.Reward,
			AnswerOptions: q,
			TimeForAnswer: question.TimeForAnswer,
			StartsAt:      question.StartsAt.Unix(),
		},
	}
}
