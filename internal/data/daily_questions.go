package data

import (
	"time"
)

type DailyQuestion struct {
	ID            int       `db:"id"`
	Title         string    `db:"title"`
	TimeForAnswer int64     `db:"time_for_answer"`
	Reward        int       `db:"reward"`
	AnswerOptions Jsonb     `db:"answer_options"`
	StartsAt      time.Time `db:"starts_at"`
	CreatedAt     time.Time `db:"created_at"`
}

type DailyQuestionsQ interface {
	New() DailyQuestionsQ
	Insert(DailyQuestion) error
	Update(map[string]any) error

	Count() (int64, error)
	Select() ([]DailyQuestion, error)
	Get() (*DailyQuestion, error)

	FilterTodayQuestions(location string) DailyQuestionsQ
	FilterByStartAtToday(location string) DailyQuestionsQ
	FilterByCreatedAt(date time.Time) DailyQuestionsQ
	FilterByID(ID int) DailyQuestionsQ
}
