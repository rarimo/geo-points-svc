package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func FilterStartAtDailyQuestions(w http.ResponseWriter, r *http.Request) {
	if !auth.Authenticates(UserClaims(r), auth.AdminGrant) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	req, err := requests.NewFilterStartAtDailyQuestions(r)
	if err != nil {
		Log(r).WithError(err).Error("error creating filter start at daily questions request")
		ape.RenderErr(w, problems.BadRequest(err)...)
	}

	res, err := DailyQuestionsQ(r).Page(&req.OffsetPageParams).SelectByTime()
	if err != nil {
		Log(r).WithError(err).Error("Error filtering questions")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	resp, err := NewDailyQuestionsFilterDate(res)
	if err != nil {
		Log(r).WithError(err).Error("Error filtering questions")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	resp.Links = req.GetLinks(r)
	if req.Count {
		questionListCount, err := DailyQuestionsQ(r).Count()
		if err != nil {
			Log(r).WithError(err).Error("Failed to count balances")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		_ = resp.PutMeta(struct {
			QuestionCount int64 `json:"question_count"`
		}{questionListCount})
	}
	ape.Render(w, resp)
}

func NewDailyQuestionModel(question data.DailyQuestion) (resources.DailyQuestionDetails, error) {
	var options []resources.DailyQuestionOptions

	err := json.Unmarshal(question.AnswerOptions, &options)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal AnswerOptions: %v", err)
		return resources.DailyQuestionDetails{}, err
	}

	return resources.DailyQuestionDetails{
		Key: resources.NewKeyInt64(question.ID, resources.DAILY_QUESTIONS),
		Attributes: resources.DailyQuestionDetailsAttributes{
			CorrectAnswer:       question.CorrectAnswer,
			CreatedAt:           question.CreatedAt.String(),
			NumAllParticipants:  question.NumAllParticipants,
			NumCorrectAnswers:   question.NumCorrectAnswers,
			NumIncorrectAnswers: question.NumIncorrectAnswers,
			Options:             options,
			Reward:              question.Reward,
			StartsAt:            question.StartsAt.String(),
			TimeForAnswer:       question.TimeForAnswer,
			Title:               question.Title,
		},
	}, nil
}

func NewDailyQuestionsFilterDate(questions []data.DailyQuestion) (resources.DailyQuestionDetailsListResponse, error) {
	list := make([]resources.DailyQuestionDetails, len(questions))
	for i, q := range questions {
		qModel, err := NewDailyQuestionModel(q)
		if err != nil {
			return resources.DailyQuestionDetailsListResponse{}, fmt.Errorf("error make daily question model, %s", err)
		}
		list[i] = qModel
	}

	return resources.DailyQuestionDetailsListResponse{Data: list}, nil
}
