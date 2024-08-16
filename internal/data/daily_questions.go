package data

import "time"

type DailyQuestion struct {
	ID            int       `db:"id"`
	Title         string    `db:"title"`
	TimeForAnswer int64     `db:"time_for_answer"`
	Reward        int       `db:"reward"`
	AnswerOptions Jsonb     `db:"answer_options"`
	StartsAt      time.Time `db:"starts_at"`
	CreatedAt     time.Time `db:"created_at"`
}

type DailyQuestionQ interface {
	New() DailyQuestionQ
	Insert(DailyQuestion) error
	Update(map[string]any) error

	Count() (int64, error)
	Select() ([]DailyQuestion, error)
	Get() (*DailyQuestion, error)

	FilterActive() DailyQuestionQ
	FilterByStartAt(date time.Time) DailyQuestionQ
	FilterByCreatedAt(date time.Time) DailyQuestionQ
	FilterByID(ID int) DailyQuestionQ
}
