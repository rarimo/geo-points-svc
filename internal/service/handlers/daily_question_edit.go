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
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func EditDailyQuestion(w http.ResponseWriter, r *http.Request) {
	if !auth.Authenticates(UserClaims(r), auth.AdminGrant) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	IDStr := strings.ToLower(chi.URLParam(r, "question_id"))
	ID, err := strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		Log(r).WithError(err).Error("Failed to parse ID")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	req, err := requests.NewDailyQuestionEdit(r)
	if err != nil {
		Log(r).WithError(err).Error("Error creating daily question edit request")
		ape.RenderErr(w, problems.InternalError())
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

	nowTime := time.Now().UTC()
	if !question.StartsAt.After(time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()+1, 0, 0, 0, 0, DailyQuestions(r).Location)) {
		Log(r).Errorf("Cannot change a question id: %v that is available today or in the past", ID)
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	requestBody := map[string]any{}

	if req.Title != nil {
		requestBody[data.ColDailyQuestionTitle] = *req.Title
	}

	if req.StartsAt != nil {
		timeReq, err := time.Parse("2006-01-02", *req.StartsAt)
		if err != nil {
			Log(r).WithError(err).Error("Failed to parse start time")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		nowTime := time.Now().UTC()
		if !timeReq.After(time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()+1, 0, 0, 0, 0, DailyQuestions(r).Location)) {
			Log(r).Errorf("Argument start_at must be more or equal tomorow midnoght now its: %s", timeReq.String())
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		location := DailyQuestions(r).Location
		question, err := DailyQuestionsQ(r).FilterDayQuestions(location, timeReq).Get()
		if err != nil {
			Log(r).WithError(err).Error("Error on this day")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if question != nil && ID != question.ID {
			Log(r).Errorf("Error on this day %v, the daily question already has %v", question.StartsAt, question)
			ape.RenderErr(w, problems.Conflict())
			return
		}
		requestBody[data.ColStartAt] = req.StartsAt
	}

	if req.CorrectAnswer != nil {
		l := len(question.AnswerOptions)
		if *req.CorrectAnswer < 0 || l <= int(*req.CorrectAnswer) {
			Log(r).Error("Invalid CorrectAnswer")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		requestBody[data.ColCorrectAnswerID] = *req.CorrectAnswer
	}

	if req.Options != nil {
		err = ValidateOptions(*req.Options)
		if err != nil {
			Log(r).WithError(err).Error("Error Answer Options")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		answerOptions, err := json.Marshal(req.Options)
		if err != nil {
			Log(r).Errorf("Error marshalling answer options: %v", err)
			ape.RenderErr(w, problems.InternalError())
			return
		}
		correctAnswerFound := false

		var localCorrectAnswer int64
		if req.CorrectAnswer != nil {
			localCorrectAnswer = *req.CorrectAnswer
		}

		for _, option := range *req.Options {
			if option.Id == int(localCorrectAnswer) {
				correctAnswerFound = true
				break
			}
		}
		if !correctAnswerFound {
			Log(r).Warnf("Correct answer option out of range: %v", question.CorrectAnswer)
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		requestBody[data.ColAnswerOption] = answerOptions
	}

	if req.Reward != nil {
		if *req.Reward <= 0 {
			Log(r).Error("Invalid Reward")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		requestBody[data.ColReward] = *req.Reward
	}

	if req.TimeForAnswer != nil {
		if *req.TimeForAnswer < 0 {
			Log(r).Error("Invalid Time for answer")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		requestBody[data.ColTimeForAnswer] = *req.TimeForAnswer
	}

	err = DailyQuestionsQ(r).FilterByID(ID).Update(requestBody)
	if err != nil {
		Log(r).WithError(err).Error("Error editing daily question")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	questionNew, err := DailyQuestionsQ(r).FilterByID(ID).Get()
	if err != nil {
		Log(r).WithError(err).Error("Error on this day")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if questionNew == nil {
		Log(r).Errorf("Error get qurstion for response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	resp, err := NewDailyQuestionEdite(ID, questionNew)
	if err != nil {
		Log(r).WithError(err).Error("Error editing daily question")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resp)
}

func NewDailyQuestionEdite(ID int64, q *data.DailyQuestion) (resources.DailyQuestionDetailsResponse, error) {
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
				CorrectAnswer: q.CorrectAnswer,
				Options:       options,
				Reward:        q.Reward,
				StartsAt:      q.StartsAt.String(),
				TimeForAnswer: q.TimeForAnswer,
				Title:         q.Title,
			},
		},
	}, nil
}
