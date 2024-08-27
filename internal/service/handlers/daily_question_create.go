package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
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
		Log(r).WithError(err).Error("Error get request NewDailyQuestion: %v")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = ValidateOptions(req.Options)
	if err != nil {
		Log(r).WithError(err).Error("Error Answer Options")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	location := DailyQuestions(r).Location
	timeReq, err := time.Parse("2006-01-02", req.StartsAt)
	if err != nil {
		Log(r).WithError(err).Error("Failed to parse start time")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	nowTime := time.Now().UTC()
	if !timeReq.After(time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.UTC)) {
		Log(r).Warnf("Arg start_at must be more or equal tommorow midnoght noe: %s", timeReq.String())
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	question, err := DailyQuestionsQ(r).FilterDayQuestions(location, timeReq).Get()
	if err != nil {
		Log(r).WithError(err).Error("Error on this day")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if question != nil {
		Log(r).Infof("Question already exist for date %s, question: %+v", question.StartsAt, question)
		ape.RenderErr(w, problems.Conflict())
		return
	}

	answerOptions, err := json.Marshal(req.Options)
	if err != nil {
		Log(r).WithError(err).Error("Failed to get questions")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	correctAnswerFound := false
	for _, option := range req.Options {
		if option.Id == int(req.CorrectAnswer) {
			correctAnswerFound = true
			break
		}
	}
	if !correctAnswerFound {
		Log(r).Warnf("Correct answer option out of range: %v", req.CorrectAnswer)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	stmt := data.DailyQuestion{
		Title:         req.Title,
		TimeForAnswer: req.TimeForAnswer,
		Reward:        req.Reward,
		AnswerOptions: answerOptions,
		CorrectAnswer: req.CorrectAnswer,
		StartsAt:      timeReq,
	}

	err = DailyQuestionsQ(r).Insert(stmt)
	if err != nil {
		Log(r).Errorf("Error ger request NewDailyQuestion: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	question, _ = DailyQuestionsQ(r).FilterDayQuestions(location, timeReq).Get()

	ape.Render(w, NewDailyQuestionCrate(&stmt, req.Options, question.ID))
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

func NewDailyQuestionCrate(q *data.DailyQuestion, options []resources.DailyQuestionOptions, ID int64) resources.DailyQuestionDetailsResponse {
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
				CreatedAt:     time.Now().UTC().String(),
			},
		},
	}

}