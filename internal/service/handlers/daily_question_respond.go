package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DailyQuestionRespond(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetDailyQuestionsRequest(r)
	currentTime := time.Now().UTC()

	//if !auth.Authenticates(UserClaims(r), auth.UserGrant(req.Nullifier)) {
	//	ape.RenderErr(w, problems.Unauthorized())
	//	return
	//}

	activeQuestions, err := DailyQuestionsQ(r).FilterActive().Select()
	if err != nil {
		Log(r).Errorf("error getting active questions: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if len(activeQuestions) == 0 {
		Log(r).Info("no active questions found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	index := rand.Intn(len(activeQuestions))
	question := activeQuestions[index]

	answers, err := DailyQuestionsResponsesQ(r).
		FilterByNullifier(req.Nullifier).
		FilterByCreatedAfter(question.StartsAt.Unix()).
		FilterByCreatedBefore(question.StartsAt.Unix() + question.TimeForAnswer).
		Select()

	Log(r).Infof("answers: %+v", answers)

	if err != nil {
		Log(r).Errorf("error selecting answer: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if len(answers) == 1 {
		index := answers[0].DailyQuestionId
		for _, question := range activeQuestions {
			if question.ID == index {
				ape.Render(w, NewDailyQuestion(question))
				return
			}
		}
		Log(r).Errorf("Not found answer for daily_question_answers nullifier: %v", req.Nullifier)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if len(answers) > 1 {
		Log(r).Errorf("multiple answers found for daily_question_answers nullifier: %v", req.Nullifier)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = DailyQuestionsResponsesQ(r).Transaction(func() error {
		evType := EventTypes(r).Get(models.TypeDailyQuestion, evtypes.FilterInactive)
		if evType == nil {
			Log(r).Infof("Event type %s is inactive", models.TypeDailyQuestion)
			return problems.InternalError()
		}

		err = DailyQuestionsResponsesQ(r).Insert(data.DailyQuestionsResponses{
			DailyQuestionId: question.ID,
			Nullifier:       req.Nullifier,
			CreatedAt:       currentTime,
		})

		if err != nil {
			Log(r).Errorf("error inserting daily_question_answer: %v", err)
			return fmt.Errorf("error insert event: %w", err)
		}

		return nil
	})

	ape.Render(w, NewDailyQuestion(question))
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
