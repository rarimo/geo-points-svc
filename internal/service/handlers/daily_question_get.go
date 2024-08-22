package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

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

	question, err := DailyQuestionsQ(r).
		FilterTodayQuestions(cfg.Timezone).
		Get()
	if question == nil {
		Log(r).Error("Error getting daily question")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if err != nil {
		Log(r).WithError(err).Error("Failed to get active questions")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	deadline := cfg.GetDeadline(nullifier)
	if deadline != nil {
		Log(r).Errorf("The user's nullifier was not found in active requests, it does not exist, or the user has already answered: %s", nullifier)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	nowTime := time.Now().UTC()
	cfg.SetDeadlineTimer(Log(r), DailyQuestionsQ(r), questionEvent, balance.Nullifier, nowTime.Unix()+question.TimeForAnswer)

	options, err := ConvertJsonbToDailyQuestionOptions(question.AnswerOptions)
	if err != nil {
		Log(r).WithError(err).Error("Failed to convert json options to daily question options")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, NewDailyQuestion(question, nowTime, options))
	return
}

func ConvertJsonbToDailyQuestionOptions(answerOptions []byte) ([]resources.DailyQuestionOptions, error) {
	var options []resources.DailyQuestionOptions
	err := json.Unmarshal(answerOptions, &options)
	if err != nil {
		return nil, err
	}
	return options, nil
}

func NewDailyQuestion(question *data.DailyQuestion, receiptTime time.Time, options []resources.DailyQuestionOptions) resources.DailyQuestions {
	var q map[string]interface{}
	_ = json.Unmarshal(question.AnswerOptions, &q)

	return resources.DailyQuestions{
		Key: resources.Key{
			ID:   strconv.Itoa(int(question.ID)),
			Type: resources.DAILY_QUESTIONS,
		},
		Attributes: resources.DailyQuestionsAttributes{
			Deadline: question.TimeForAnswer + receiptTime.Unix(),
			Options:  options,
			Title:    question.Title,
		},
	}
}
