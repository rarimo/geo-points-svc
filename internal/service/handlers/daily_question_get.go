package handlers

import (
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
	dq := DailyQuestions(r)

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	log := Log(r).WithField("nullifier", nullifier)

	balance, err := BalancesQ(r).FilterByNullifier(nullifier).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if balance == nil {
		log.Debug("Balance absent")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	deadline := dq.GetDeadline(nullifier)
	if deadline != nil {
		log.Debugf("Question already recieved with deadline %s", deadline.At)
		ape.RenderErr(w, problems.Conflict())
		return
	}

	localDayStart := atDayStart(dq.LocalTime(time.Now().UTC()))

	ev, err := EventsQ(r).FilterByNullifier(nullifier).FilterByType(models.TypeDailyQuestion).GetLast()
	if err != nil {
		log.WithError(err).Error("Failed to get last daily_question event")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if ev != nil &&
		ev.CreatedAt > int32(localDayStart.Unix()) &&
		ev.CreatedAt < int32(localDayStart.Add(24*time.Hour).Unix()) {
		log.Debug("Points already accruing for daily question")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	question, err := DailyQuestionsQ(r).FilterByStartsAtAfter(localDayStart).Get()
	if err != nil {
		log.WithError(err).Error("Failed to get question")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if question == nil || !question.StartsAt.Before(localDayStart.Add(24*time.Hour)) {
		log.Debugf("Next question: %+v", question)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if !dq.SetDeadline(nullifier, int(question.ID), time.Duration(question.TimeForAnswer)*time.Second) {
		log.Debug("Worker clear deadlines before next question, now getting questions unavailable")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	options, err := question.ExtractOptions()
	if err != nil {
		log.WithError(err).Errorf("Failed to extract options from question: %s", string(question.AnswerOptions))
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newDailyQuestions(question, options))
}

func newDailyQuestions(question *data.DailyQuestion, options []resources.DailyQuestionOptions) resources.DailyQuestionsResponse {

	return resources.DailyQuestionsResponse{
		Data: resources.DailyQuestions{
			Key: resources.Key{
				ID:   strconv.Itoa(int(question.ID)),
				Type: resources.DAILY_QUESTIONS,
			},
			Attributes: resources.DailyQuestionsAttributes{
				Options: options,
				Title:   question.Title,
			},
		},
	}
}
