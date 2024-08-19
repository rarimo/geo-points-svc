package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetDailyQuestion(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewDailyQuestionUserAccess(r)
	if err != nil {
		Log(r).WithError(err).Error("error getting daily question user details")
	}

	question, err := DailyQuestionsQ(r).
		FilterTodayQuestions(req.Timezone).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get active questions")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, NewDailyQuestion(*question))
}

func ConvertJsonbToDailyQuestionAnswers(jb data.Jsonb) map[string]interface{} {
	var res map[string]interface{}

	err := json.Unmarshal(jb, &res)
	if err != nil {
		return res
	}

	return res
}

func NewDailyQuestion(question data.DailyQuestion) resources.DailyQuestionAttributes {
	return resources.DailyQuestionAttributes{
		ID:            question.ID,
		Title:         question.Title,
		Reward:        question.Reward,
		AnswerOptions: ConvertJsonbToDailyQuestionAnswers(question.AnswerOptions),
		TimeForAnswer: question.TimeForAnswer,
		StartsAt:      question.StartsAt.Unix(),
	}
}
