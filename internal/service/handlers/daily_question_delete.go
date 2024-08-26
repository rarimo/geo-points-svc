package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteDailyQuestion(w http.ResponseWriter, r *http.Request) {
	IDStr := strings.ToLower(chi.URLParam(r, "ID"))
	ID, err := strconv.ParseInt(IDStr, 10, 64)
	Log(r).Infof("ID: %v", ID)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse ID")
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
		Log(r).Warnf("Question with ID %d not found", ID)
		ape.RenderErr(w, problems.NotFound())
		return
	}
	title := question.Title

	_, err = DailyQuestionsQ(r).New().FilterByID(ID).Delete()
	if err != nil {
		Log(r).WithError(err).Error("Error deleting daily question")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, NewDailyQuestionDelete(ID, title))
}

func NewDailyQuestionDelete(ID int64, title string) resources.DailyQuestionDelResponse {
	return resources.DailyQuestionDelResponse{
		Data: resources.DailyQuestionDel{
			Key: resources.Key{
				ID:   strconv.Itoa(int(ID)),
				Type: resources.DAILY_QUESTION_DEL,
			},
			Attributes: resources.DailyQuestionDelAttributes{
				Title: title,
			},
		},
	}
}
