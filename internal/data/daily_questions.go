package data

type DailyQuestions struct {
	ID            int    `db:"id"`
	Title         string `db:"title"`
	TimeForAnswer int    `db:"time_for_answer"`
	Bounty        int    `db:"bounty"`
	AnswerOptions string `db:"answer_options"`
	Active        bool   `db:"active"`
	StartsAt      int    `db:"starts_at"`
}

type DailyQuestionsQ interface {
	New() DailyQuestionsQ
	Insert(DailyQuestions) error
	Update(map[string]any) error

	Count() (int64, error)
	Select() ([]DailyQuestions, error)

	FilteredActive(status bool) DailyQuestionsQ
	FilteredStartAt(date int) DailyQuestionsQ
}
