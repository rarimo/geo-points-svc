package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
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
	deletedQuestion := *question

	timeReq := question.StartsAt
	nowTime := time.Now().UTC()
	if !timeReq.After(time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()+1, 0, 0, 0, 0, DailyQuestions(r).Location)) {
		Log(r).Errorf("Only questions that start tomorrow or later can be delete: %s", timeReq.String())
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	_, err = DailyQuestionsQ(r).New().FilterByID(ID).Delete()
	if err != nil {
		Log(r).WithError(err).Error("Error deleting daily question")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	loc := DailyQuestions(r).Location
	response, err := NewDailyQuestionDelete(ID, deletedQuestion, loc)
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
