package data

import "time"

type DailyQuestionsResponses struct {
	ID              int       `db:"id"`
	DailyQuestionId int       `db:"daily_question_id"`
	Nullifier       string    `db:"nullifier"`
	CreatedAt       time.Time `db:"created_at"`
}

type DailyQuestionsResponsesQ interface {
	New() DailyQuestionsResponsesQ
	Insert(DailyQuestionsResponses) error
	Update(map[string]any) error
	Transaction(func() error) error

	Count() (int64, error)
	Select() ([]DailyQuestionsResponses, error)
	Get() (*DailyQuestionsResponses, error)

	FilterByQuestionID(ID int) DailyQuestionsResponsesQ
	FilterByNullifier(nullifier string) DailyQuestionsResponsesQ
	FilterByCreatedBefore(before int64) DailyQuestionsResponsesQ
	FilterByCreatedAfter(after int64) DailyQuestionsResponsesQ
	FilterByCreatedAt(date time.Time) DailyQuestionsResponsesQ
}
