package data

import (
	"time"
)

type DailyQuestion struct {
	ID            int64     `db:"id"`
	Title         string    `db:"title"`
	TimeForAnswer int64     `db:"time_for_answer"`
	Reward        int       `db:"reward"`
	AnswerOptions Jsonb     `db:"answer_options"`
	StartsAt      time.Time `db:"starts_at"`
	CreatedAt     time.Time `db:"created_at"`
	CorrectAnswer int64     `db:"correct_answer"`
}

type DailyQuestionsQ interface {
	New() DailyQuestionsQ
	Insert(DailyQuestion) error
	Update(map[string]any) error

	Count() (int64, error)
	Select() ([]DailyQuestion, error)
	Get() (*DailyQuestion, error)

	FilterTodayQuestions(offset int) DailyQuestionsQ
	FilterByCreatedAt(date time.Time) DailyQuestionsQ
	FilterByID(ID int64) DailyQuestionsQ
}
