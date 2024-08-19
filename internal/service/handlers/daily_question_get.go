package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

var mu sync.Mutex

func GetDailyQuestion(w http.ResponseWriter, r *http.Request) {
	nullifier := strings.ToLower(chi.URLParam(r, "nullifier"))

	//if !auth.Authenticates(UserClaims(r), auth.UserGrant(*req.Nullifier)) {
	//	ape.RenderErr(w, problems.Unauthorized())
	//	return
	//}

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
		FilterTodayQuestions().
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
		FilterTodayEvents().
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

	SetDailyQuestionTimeWithExpiration(r, balance.Nullifier, time.Now().UTC().Unix(), question.TimeForAnswer, false)

	ape.Render(w, NewDailyQuestion(*question))
	return
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

func SetDailyQuestionTimeWithExpiration(r *http.Request, nullifier string, timestamp int64, duration int64, status bool) {
	mu.Lock()
	defer mu.Unlock()

	DailyQuestionTimeHash(r).SetDailyQuestionsTimeHash(nullifier, config.DailyQuestionTimeInfo{
		MaxDateToAnswer: timestamp + duration,
		Answered:        status,
	})
	Log(r).Infof("add %s %v, length q: %v, mapm %+v", nullifier, duration, len(DailyQuestionTimeHash(r)), DailyQuestionTimeHash(r))

	go func() {
		time.Sleep(time.Duration(duration) * time.Second)

		mu.Lock()
		defer mu.Unlock()

		info := DailyQuestionTimeHash(r).GetDailyQuestionsTimeHash(nullifier)
		if info.Answered {
			delete(DailyQuestionTimeHash(r), nullifier)
			Log(r).Infof("Removed entry for nullifier: %s after expiration", nullifier)
		}
	}()
}

func GetLocationFromTimezone(timezone string) *time.Location {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		Log(nil).WithError(err).Errorf("error loading timezone, defaulting to UTC")
		return time.UTC
	}
	return location
}
