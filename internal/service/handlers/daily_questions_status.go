package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetDailyQuestionsStatus(w http.ResponseWriter, r *http.Request) {
	nullifier := strings.ToLower(chi.URLParam(r, "nullifier"))

	if !auth.Authenticates(UserClaims(r), auth.VerifiedGrant(nullifier)) {
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

	deadline := DailyQuestions(r).GetDeadline(nullifier)
	query := DailyQuestionsQ(r).FilterByStartsAtAfter(atDayStart(DailyQuestions(r).LocalTime(time.Now().UTC())))
	if deadline != nil {
		query = DailyQuestionsQ(r).FilterByStartsAtAfter(atDayStart(DailyQuestions(r).LocalTime(time.Now().UTC())).Add(time.Hour * 24))
	}

	question, err := query.Get()
	if err != nil {
		log.WithError(err).Error("Failed to get question")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if question == nil {
		log.Debug("Next question absent")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, newDailyQuestionsStatus(question))
}

func newDailyQuestionsStatus(question *data.DailyQuestion) resources.DailyQuestionsStatus {
	return resources.DailyQuestionsStatus{
		Key: resources.NewKeyInt64(question.ID, resources.DAILY_QUESTIONS),
		Attributes: resources.DailyQuestionsStatusAttributes{
			NextQuestionDate: question.StartsAt.Unix(),
			Reward:           int64(question.Reward),
			TimeForAnswer:    question.TimeForAnswer,
		},
	}
}

func atDayStart(date time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, date.Location())
}
