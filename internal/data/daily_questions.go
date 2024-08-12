package data

type DailyQuestions struct {
	ID            int    `db:"id"`
	Title         string `db:"title"`
	TimeForAnswer int    `db:"time_for_answer"`
	Bounty        int    `db:"bounty"`
	AnswerOptions string `db:"answer_options"`
	Active        bool   `db:"active"`
}

type DailyQuestionsQ interface {
	New() DailyQuestionsQ
	Insert(DailyQuestions) error
	Update(map[string]any) error

	Count() (int64, error)
	CountActive() (int64, error)

	FilteredActive(status bool) DailyQuestionsQ
}
