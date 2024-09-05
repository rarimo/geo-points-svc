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
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"query": fmt.Errorf("failed to parse ID: %v, err: %s", ID, err),
		})...)
		return
	}

	req, err := requests.NewDailyQuestionEdit(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	attributes := req.Data.Attributes

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
	nowTime := time.Now().UTC()

	if question.StartsAt.UTC().Before(nowTime) {
		Log(r).Errorf("Cannot change a question id: %v that is available today or in the past", ID)
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"starts_at": fmt.Errorf("cannot change a question id: %v that is available today or in the past", ID),
		})...)
		return
	}

	requestBody := map[string]any{}

	if attributes.Title != nil {
		requestBody[data.ColDailyQuestionTitle] = *attributes.Title
	}

	if attributes.StartsAt != nil {
		timeReq, err := time.ParseInLocation("2006-01-02", *attributes.StartsAt, location)
		if err != nil {
			Log(r).WithError(err).Error("Failed to parse start time")
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"starts_at": fmt.Errorf("failed to parse start time %s err: %s", *attributes.StartsAt, err),
			})...)
			return
		}

		if !timeReq.After(nowTime.UTC().AddDate(0, 0, -1)) {
			Log(r).Errorf("Argument start_at must be more or equal tomorow midnoght now its: %s", timeReq.UTC().String())
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"starts_at": fmt.Errorf("argument start_at must be more or equal tomorow midnoght now its: %s", timeReq.UTC().String()),
			})...)
			return
		}

		question, err := DailyQuestionsQ(r).FilterDayQuestions(timeReq.UTC()).Get()
		if err != nil {
			Log(r).WithError(err).Error("Error on this day")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if question != nil && ID != question.ID {
			Log(r).Errorf("Error on this day %s, the daily question already has %s", question.StartsAt.String(), question)
			ape.RenderErr(w, problems.Conflict())
			return
		}
		requestBody[data.ColStartAt] = timeReq.UTC()
	}

	if attributes.CorrectAnswer != nil {
		l := len(question.AnswerOptions)
		if *attributes.CorrectAnswer < 0 || l <= int(*attributes.CorrectAnswer) {
			Log(r).Error("Invalid CorrectAnswer")
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"correct_answer": fmt.Errorf("invalid value for correct_answer: %v", *attributes.CorrectAnswer),
			})...)
			return
		}
		requestBody[data.ColCorrectAnswerID] = *attributes.CorrectAnswer
	}

	if attributes.Options != nil {
		err = ValidateOptions(*attributes.Options)
		if err != nil {
			Log(r).WithError(err).Error("Error Answer Options")
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"options": fmt.Errorf("invalid options: %v, err: %s", *attributes.Options, err),
			})...)
			return
		}

		answerOptions, err := json.Marshal(attributes.Options)
		if err != nil {
			Log(r).Errorf("Error marshalling answer options: %v", err)
			ape.RenderErr(w, problems.InternalError())
			return
		}
		correctAnswerFound := false

		var localCorrectAnswer int64
		if attributes.CorrectAnswer != nil {
			localCorrectAnswer = *attributes.CorrectAnswer
		}

		for _, option := range *attributes.Options {
			if option.Id == int(localCorrectAnswer) {
				correctAnswerFound = true
				break
			}
		}
		if !correctAnswerFound {
			Log(r).Warnf("Correct answer option out of range: %v", question.CorrectAnswer)
			ape.RenderErr(w, problems.BadRequest(
				validation.Errors{
					"correct_answer": fmt.Errorf("correct answer option out of range %v", localCorrectAnswer),
				})...)
			return
		}
		requestBody[data.ColAnswerOption] = answerOptions
	}

	if attributes.Reward != nil {
		if *attributes.Reward <= 0 {
			Log(r).Error("Invalid Reward")
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"reward": fmt.Errorf("reward less than or equal to 0 reward: %v", attributes.Reward),
			})...)
			return
		}
		requestBody[data.ColReward] = *attributes.Reward
	}

	if attributes.TimeForAnswer != nil {
		if *attributes.TimeForAnswer < 0 {
			Log(r).Error("Invalid Time for answer")
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"time_for_answer": fmt.Errorf("invalid value for time_for_answer: %v", *attributes.TimeForAnswer),
			})...)
			return
		}
		requestBody[data.ColTimeForAnswer] = *attributes.TimeForAnswer
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
		Log(r).Errorf("Error get question for response")
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
