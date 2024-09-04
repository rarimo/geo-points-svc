package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rarimo/geo-points-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	ColDailyQuestionTitle = "title"
	ColTimeForAnswer      = "time_for_answer"
	ColAnswerOption       = "answer_options"
	ColCorrectAnswerID    = "correct_answer"
	ColReward             = "reward"
	ColStartAt            = "starts_at"
	ColCorrectAnswers     = "num_correct_answers"
	ColIncorrectAnswers   = "num_incorrect_answers"
	ColAllParticipants    = "num_all_participants"
)

type DailyQuestion struct {
	ID                  int64     `db:"id"`
	Title               string    `db:"title"`
	TimeForAnswer       int64     `db:"time_for_answer"`
	Reward              int64     `db:"reward"`
	AnswerOptions       Jsonb     `db:"answer_options"`
	StartsAt            time.Time `db:"starts_at"`
	CreatedAt           time.Time `db:"created_at"`
	CorrectAnswer       int64     `db:"correct_answer"`
	NumCorrectAnswers   int64     `db:"num_correct_answers"`
	NumIncorrectAnswers int64     `db:"num_incorrect_answers"`
	NumAllParticipants  int64     `db:"num_all_participants"`
}

type DailyQuestionsQ interface {
	New() DailyQuestionsQ
	Insert(DailyQuestion) error
	Update(map[string]any) error
	Delete() (int64, error)
	Count() (int64, error)
	Select() ([]DailyQuestion, error)

	Get() (*DailyQuestion, error)
	Page(*pgdb.OffsetPageParams) DailyQuestionsQ
	FilterTodayQuestions(offset int) DailyQuestionsQ
	FilterByCreatedAtAfter(date time.Time) DailyQuestionsQ
	FilterByStartsAtAfter(date time.Time) DailyQuestionsQ
	FilterByID(ID int64) DailyQuestionsQ
	FilterDayQuestions(day time.Time) DailyQuestionsQ

	IncrementCorrectAnswer() error
	IncrementIncorrectAnswer() error
	IncrementAllParticipants() error
}

func (q *DailyQuestion) ExtractOptions() ([]resources.DailyQuestionOptions, error) {
	var options []resources.DailyQuestionOptions
	err := json.NewDecoder(bytes.NewReader(q.AnswerOptions)).Decode(&options)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal question options: %w", err)
	}

	return options, nil
}
