package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteDailyQuestion(w http.ResponseWriter, r *http.Request) {
	if !auth.Authenticates(UserClaims(r), auth.AdminGrant) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	IDStr := strings.ToLower(chi.URLParam(r, "question_id"))
	ID, err := strconv.ParseInt(IDStr, 10, 64)
	Log(r).Infof("ID: %v", ID)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse ID")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	question, err := DailyQuestionsQ(r).FilterByID(ID).Get()
	if err != nil {
		Log(r).WithError(err).Error("Error getting question")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if question == nil {
		Log(r).Errorf("Question with ID %d not found", ID)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	location := DailyQuestions(r).Location
	timeReq, err := time.ParseInLocation("2006-01-02 15:04:05 -0700 MST", question.StartsAt.String(), location)
	if err != nil {
		Log(r).WithError(err).Error("Failed to parse start time")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"starts_at": fmt.Errorf("failed to parse start time %s err: %s", question.StartsAt.String(), err),
		})...)
		return
	}

	nowTime := time.Now().In(location).UTC()

	if timeReq.UTC().Before(nowTime.AddDate(0, 0, 1)) {
		Log(r).Errorf("Error %s", timeReq.UTC().String())
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"starts_at": fmt.Errorf("argument start_at must be more or equal tomorow mid night now: %s", timeReq.UTC().String()),
		})...)
		return
	}

	_, err = DailyQuestionsQ(r).New().FilterByID(ID).Delete()
	if err != nil {
		Log(r).WithError(err).Error("Error deleting daily question")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response, err := NewDailyQuestionDelete(ID, *question, location)
	if err != nil {
		Log(r).WithError(err).Error("Error deleting daily question")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, response)
}

func NewDailyQuestionDelete(ID int64, q data.DailyQuestion, loc *time.Location) (resources.DailyQuestionDetailsResponse, error) {
	var options []resources.DailyQuestionOptions
	err := json.Unmarshal(q.AnswerOptions, &options)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal AnswerOptions: %v", err)
		return resources.DailyQuestionDetailsResponse{}, err
	}

	return resources.DailyQuestionDetailsResponse{
		Data: resources.DailyQuestionDetails{
			Key: resources.NewKeyInt64(ID, resources.DAILY_QUESTIONS),
			Attributes: resources.DailyQuestionDetailsAttributes{
				Title:         q.Title,
				Options:       options,
				CorrectAnswer: q.CorrectAnswer,
				Reward:        q.Reward,
				TimeForAnswer: q.TimeForAnswer,
				StartsAt:      q.StartsAt.In(loc).String(),
			},
		},
	}, nil
}
