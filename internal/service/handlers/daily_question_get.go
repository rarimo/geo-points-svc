package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

var mu sync.Mutex

func GetDailyQuestion(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewDailyQuestionUserAccess(r)
	if err != nil {
		Log(r).WithError(err).Error("error getting daily question user details")
	}

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(*req.Nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	location := GetLocationFromTimezone(req.Timezone)

	question, err := DailyQuestionsQ(r).
		FilterTodayQuestions(req.Timezone).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get active questions")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	SetDailyQuestionTimeWithExpiration(r, *req.Nullifier, time.Now().In(location).Unix(), question.TimeForAnswer)

	ape.Render(w, NewDailyQuestion(*question))
}

func ConvertJsonbToDailyQuestionAnswers(jb data.Jsonb) map[string]interface{} {
	var res map[string]interface{}

	err := json.Unmarshal(jb, &res)
	if err != nil {
		return res
	}

	return res
}

func NewDailyQuestion(question data.DailyQuestion) resources.DailyQuestionAttributes {
	return resources.DailyQuestionAttributes{
		ID:            question.ID,
		Title:         question.Title,
		Reward:        question.Reward,
		AnswerOptions: ConvertJsonbToDailyQuestionAnswers(question.AnswerOptions),
		TimeForAnswer: question.TimeForAnswer,
		StartsAt:      question.StartsAt.Unix(),
	}
}

func GetLocationFromTimezone(timezone string) *time.Location {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		Log(nil).WithError(err).Errorf("error loading timezone, defaulting to UTC")
		return time.UTC
	}
	return location
}

func SetDailyQuestionTimeWithExpiration(r *http.Request, nullifier string, timestamp int64, duration int64) {
	mu.Lock()
	defer mu.Unlock()

	DailyQuestionTimeHash(r).SetDailyQuestionsTimeHash(nullifier, timestamp)

	go func() {
		time.Sleep(time.Duration(duration) * time.Second)

		mu.Lock()
		defer mu.Unlock()

		delete(DailyQuestionTimeHash(r).GetDailyQuestionsTimeHash(), nullifier)
		Log(r).Infof("Removed entry for nullifier: %s after expiration", nullifier)
	}()
}
