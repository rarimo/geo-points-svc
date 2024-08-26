package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func EditDailyQuestion(w http.ResponseWriter, r *http.Request) {
	IDStr := strings.ToLower(chi.URLParam(r, "ID"))
	ID, err := strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	req, err := requests.NewDailyQuestionEdit(r)
	if err != nil {
		Log(r).WithError(err).Error("Error creating daily question edit request")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	location := DailyQuestions(r).Location
	question, err := DailyQuestionsQ(r).FilterDayQuestions(location, req.StartsAt).Get()
	if question != nil && ID != question.ID {
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

	err = DailyQuestionsQ(r).FilterByID(ID).Update(map[string]any{
		data.ColDailyQuestionTitle:  req.Title,
		data.ColTimeForAnswer:       req.TimeForAnswer,
		data.ColDailyQuestionReward: req.Reward,
		data.ColAnswerOption:        answerOptions,
		data.ColCorrectAnswerId:     req.CorrectAnswer,
	})

	if err != nil {
		Log(r).WithError(err).Error("Error editing daily question")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, NewDailyQuestionEdite(ID, req))
}

func NewDailyQuestionEdite(ID int64, req resources.DailyQuestionEditAttributes) resources.DailyQuestionEditResponse {
	return resources.DailyQuestionEditResponse{
		Data: resources.DailyQuestionEdit{
			Key: resources.Key{
				ID:   strconv.Itoa(int(ID)),
				Type: resources.DAILY_QUESTION_EDIT,
			},
			Attributes: resources.DailyQuestionEditAttributes{
				CorrectAnswer: req.CorrectAnswer,
				Options:       req.Options,
				Reward:        req.Reward,
				StartsAt:      req.StartsAt,
				TimeForAnswer: req.TimeForAnswer,
				Title:         req.Title,
			},
		},
	}

}
