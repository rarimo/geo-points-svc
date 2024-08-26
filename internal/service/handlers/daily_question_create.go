package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateDailyQuestion(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewDailyQuestion(r)
	if err != nil {
		Log(r).Errorf("Error get request NewDailyQuestion: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	location := DailyQuestions(r).Location
	question, err := DailyQuestionsQ(r).FilterDayQuestions(location, req.StartsAt).Get()
	if question != nil {
		Log(r).Errorf("Error on this day %v, the daily question already has %v", question.StartsAt, question)
		ape.RenderErr(w, problems.Conflict())
		return
	}
	if err != nil {
		Log(r).Errorf("Error on this day %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	answerOptions, err := json.Marshal(req.Options)
	if err != nil {
		Log(r).Errorf("Error marshalling answer options: %v", err)
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
		Reward:        int64(req.Reward),
		AnswerOptions: answerOptions,
		CorrectAnswer: req.CorrectAnswer,
		StartsAt:      req.StartsAt,
	}

	err = DailyQuestionsQ(r).Insert(stmt)
	if err != nil {
		Log(r).Errorf("Error ger request NewDailyQuestion: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	question, err = DailyQuestionsQ(r).FilterDayQuestions(location, req.StartsAt).Get()
	if question == nil {
		Log(r).Errorf("Error on this day %v, create question '%v'", req.StartsAt, req.Title)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if err != nil {
		Log(r).Errorf("Error on this day %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, NewDailyQuestionCrate(&stmt, req.Options, question.ID))
}

func NewDailyQuestionCrate(q *data.DailyQuestion, options []resources.DailyQuestionOptions, ID int64) resources.DailyQuestionDetailsResponse {
	return resources.DailyQuestionDetailsResponse{
		Data: resources.DailyQuestionDetails{
			Key: resources.Key{
				ID:   strconv.Itoa(int(ID)),
				Type: resources.DAILY_QUESTIONS,
			},
			Attributes: resources.DailyQuestionDetailsAttributes{
				CorrectAnswer: q.CorrectAnswer,
				Options:       options,
				Reward:        int(q.Reward),
				StartsAt:      q.StartsAt,
				TimeForAnswer: q.TimeForAnswer,
				Title:         q.Title,
			},
		},
	}

}
