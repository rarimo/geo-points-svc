package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetDailyQuestion(w http.ResponseWriter, r *http.Request) {
	//if !auth.Authenticates(UserClaims(r), auth.UserGrant(req.UserNullifier)) {
	//	ape.RenderErr(w, problems.Unauthorized())
	//	return
	//}

	req, err := requests.NewGetDailyQuestionsRequest(r)

	activeQuestions, err := DailyQuestionsQ(r).FilterByActive(true).Select()
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

	byNullifier, err := BalancesQ(r).FilterByNullifier(req.UserNullifier).Get()
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = EventsQ(r).Transaction(func() error {
		evType := EventTypes(r).Get(models.TypeDailyQuestion, evtypes.FilterInactive)
		if evType == nil {
			Log(r).Infof("Event type %s is inactive", models.TypeEarlyTest)
			return nil
		}

		eveCheck, err := EventsQ(r).FilterByNullifier(req.UserNullifier).FilterByType(models.TypeDailyQuestion).Get()
		if err != nil {
			Log(r).WithError(err).Errorf("Failed to get balance by nullifier %s", req.UserNullifier)
		}
		if eveCheck != nil {
			Log(r).Infof("User already Gget daily question %s", req.UserNullifier)
			return nil
		}

		err = EventsQ(r).Insert(data.Event{
			Nullifier: req.UserNullifier,
			Type:      models.TypeDailyQuestion,
			Status:    data.EventOpen,
		})

		if err != nil {
			return fmt.Errorf("failed to insert `%v` event: %w", models.TypeDailyQuestion, err)
		}

		//TODO
		//if byNullifier.IsDisabled() {
		//	Log(r).Debug("Balance is disabled, events will not be claimed")
		//	return nil
		//}

		if err = autoClaimEventsForBalance(r, byNullifier); err != nil {
			return fmt.Errorf("failed to autoclaim events for user")
		}

		if err := addEventForReferrer(r, byNullifier); err != nil {
			return fmt.Errorf("add event for referrer: %w", err)
		}

		return nil
	})

	if err != nil {
		Log(r).WithError(err).Error("Failed to transaction event getDailyQuestion")
		ape.RenderErr(w, problems.InternalError())
		return
	}

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

func NewDailyQuestion(req data.DailyQuestion) resources.DailyQuestionAttributes {
	return resources.DailyQuestionAttributes{
		ID:            req.ID,
		Title:         req.Title,
		Reward:        req.Reward,
		AnswerOptions: ConvertJsonbToDailyQuestionAnswers(req.AnswerOptions),
		TimeForAnswer: req.TimeForAnswer,
		Active:        &req.Active,
		StartsAt:      int(req.StartsAt.Unix()),
	}
}
