package handlers

import (
	"fmt"
	"math"
	"net/http"
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
	var timeToNext int64
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
		FilterTodayEvents(Location(r).String()).
		FilterByType(models.TypeDailyQuestion).
		Get()

	if err != nil {
		Log(r).WithError(err).Error("error getting event daily_question")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	timeToNext, err = TimeToQuestionUnix(r)
	if err != nil {
		Log(r).WithError(err).Error("error getting time to next question")
		timeToNext = -1
	}

	if dailyQuestionEvent != nil {
		ape.Render(w, NewDailyQuestionsStatus(true, timeToNext))
		return
	}

	ape.Render(w, NewDailyQuestionsStatus(false, timeToNext))
	return
}

func TimeToQuestionUnix(r *http.Request) (int64, error) {
	questions, err := DailyQuestionsQ(r).FilterByStartAtToday(Location(r).String()).Select()
	if err != nil {
		return -2, err
	}

	if len(questions) == 0 {
		return -1, nil
	}

	closes := int64(math.MaxInt64)
	for _, q := range questions {
		timeToNext := q.StartsAt.Unix() - time.Now().Unix()
		if timeToNext < closes {
			closes = timeToNext
		}
	}

	return closes, nil
}

func FormatUnixTimeToDate(TimeToNext int64) string {
	if TimeToNext == -1 {
		return "soon"
	}

	if TimeToNext < 0 {
		return "0d:00h:00m:00s"
	}

	days := TimeToNext / (24 * 3600)
	TimeToNext %= 24 * 3600
	hours := TimeToNext / 3600
	TimeToNext %= 3600
	minutes := TimeToNext / 60
	seconds := TimeToNext % 60

	return fmt.Sprintf("%dd:%02dh:%02dm:%02ds", days, hours, minutes, seconds)
}

func NewDailyQuestionsStatus(AlreadyDoneForUser bool, TimeToNext int64) resources.DailyQuestionStatusAttributes {
	timeToNextStr := FormatUnixTimeToDate(TimeToNext)

	return resources.DailyQuestionStatusAttributes{
		AlreadyDoneForUser: AlreadyDoneForUser,
		TimeToNext:         timeToNextStr,
	}
}
