package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/service/referralid"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateDailyQuestion(w http.ResponseWriter, r *http.Request) {
	if !auth.Authenticates(UserClaims(r), auth.AdminGrant) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	req, err := requests.NewDailyQuestion(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	attributes := req.Data.Attributes

	err = ValidateOptions(attributes.Options)
	if err != nil {
		Log(r).WithError(err).Error("Error Answer Options")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"options": fmt.Errorf("invalid options: %v, err: %s", attributes.Options, err),
		})...)
		return
	}

	location := DailyQuestions(r).Location
	timeReq, err := time.ParseInLocation("2006-01-02", attributes.StartsAt, location)
	if err != nil {
		Log(r).WithError(err).Error("Failed to parse start time")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"starts_at": fmt.Errorf("failed to parse start time %s err: %s", attributes.StartsAt, err),
		})...)
		return
	}
	timeReq = timeReq.UTC()
	nowTime := time.Now().UTC()
	if !referralid.CheckOpportunityChange(nowTime, timeReq, location) {
		Log(r).Errorf("Arg start_at must be more or equal tomorow mid night noe: %s", timeReq.String())
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"starts_at": fmt.Errorf("argument start_at must be more or equal tomorow mid night now its: %s", timeReq.UTC().String()),
		})...)
		return
	}

	question, err := DailyQuestionsQ(r).FilterDayQuestions(timeReq).Get()
	if err != nil {
		Log(r).WithError(err).Error("Error on this day")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if question != nil {
		Log(r).Errorf("Question already exist for date %s, question: %+v", question.StartsAt, question)
		ape.RenderErr(w, problems.Conflict())
		return
	}

	answerOptions, err := json.Marshal(attributes.Options)
	if err != nil {
		Log(r).WithError(err).Error("Failed to get questions")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	correctAnswerFound := false
	for _, option := range attributes.Options {
		if option.Id == int(attributes.CorrectAnswer) {
			correctAnswerFound = true
			break
		}
	}
	if !correctAnswerFound {
		Log(r).Errorf("Correct answer option out of range: %v", attributes.CorrectAnswer)
		ape.RenderErr(w, problems.BadRequest(
			validation.Errors{
				"correct_answer": fmt.Errorf("correct answer option out of range %v", attributes.CorrectAnswer),
			})...)
		return
	}

	if attributes.Reward <= 0 {
		Log(r).Errorf("Reward option out of range: %v", attributes.Reward)
		ape.RenderErr(w, problems.BadRequest(
			validation.Errors{
				"reward": fmt.Errorf("reward less than or equal to 0 reward: %v", attributes.Reward),
			})...)
		return
	}

	dailyQuestion := data.DailyQuestion{
		Title:         attributes.Title,
		TimeForAnswer: attributes.TimeForAnswer,
		Reward:        attributes.Reward,
		AnswerOptions: answerOptions,
		CorrectAnswer: attributes.CorrectAnswer,
		StartsAt:      timeReq.UTC(),
	}

	err = DailyQuestionsQ(r).Insert(dailyQuestion)
	if err != nil {
		Log(r).WithError(err).Error("Error ger request NewDailyQuestion")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	question, err = DailyQuestionsQ(r).FilterDayQuestions(timeReq.UTC()).Get()
	if err != nil {
		Log(r).WithError(err).Error("Error on this day")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if question == nil {
		Log(r).Errorf("Error get question for response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, NewDailyQuestionCreate(&dailyQuestion, attributes.Options, question.ID))
}

func ValidateOptions(options []resources.DailyQuestionOptions) error {
	if len(options) < 2 || len(options) > 6 {
		return fmt.Errorf("the number of options must be between 2 and 6")
	}

	uniqueTitles := make(map[string]bool)

	for _, option := range options {
		if option.Title == "" {
			return fmt.Errorf("option titles must not be empty")
		}
		if _, exists := uniqueTitles[option.Title]; exists {
			return fmt.Errorf("option titles must be unique, found duplicate: %s", option.Title)
		}
		uniqueTitles[option.Title] = true
	}

	ids := make([]int, len(options))
	for i, option := range options {
		ids[i] = option.Id
	}
	sort.Ints(ids)
	for i := 0; i < len(ids); i++ {
		if ids[i] != i {
			return fmt.Errorf("option IDs must be sequential and start from 0")
		}
	}
	return nil
}

func NewDailyQuestionCreate(q *data.DailyQuestion, options []resources.DailyQuestionOptions, ID int64) resources.DailyQuestionDetailsResponse {
	return resources.DailyQuestionDetailsResponse{
		Data: resources.DailyQuestionDetails{
			Key: resources.NewKeyInt64(ID, resources.DAILY_QUESTIONS),
			Attributes: resources.DailyQuestionDetailsAttributes{
				Title:         q.Title,
				Options:       options,
				CorrectAnswer: q.CorrectAnswer,
				Reward:        q.Reward,
				TimeForAnswer: q.TimeForAnswer,
				StartsAt:      q.StartsAt.String(),
			},
		},
	}

}
