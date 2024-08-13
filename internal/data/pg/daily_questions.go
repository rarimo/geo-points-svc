package pg

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/rarimo/geo-points-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const dailyQuestionsTable = "daily_questions"

type dailyQuestions struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
	counter  squirrel.SelectBuilder
}

func NewDailyQuestions(db *pgdb.DB) data.DailyQuestionsQ {
	return &dailyQuestions{
		db:       db,
		selector: squirrel.Select("*").From(dailyQuestionsTable),
		updater:  squirrel.Update(dailyQuestionsTable),
		counter:  squirrel.Select("COUNT(*) as count").From(dailyQuestionsTable),
	}
}

func (q *dailyQuestions) New() data.DailyQuestionsQ {
	return NewDailyQuestions(q.db)
}

func (q *dailyQuestions) Insert(quest data.DailyQuestions) error {
	stmt := squirrel.Insert(dailyQuestionsTable).SetMap(map[string]interface{}{
		"title":           quest.Title,
		"time_for_answer": quest.TimeForAnswer,
		"bounty":          quest.Bounty,
		"answer_options":  quest.AnswerOptions,
		"active":          quest.Active,
	})

	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("insert daily_questions %+v: %w", quest, err)
	}

	return nil
}

func (q *dailyQuestions) Update(fields map[string]any) error {
	if err := q.db.Exec(q.updater.SetMap(fields)); err != nil {
		return fmt.Errorf("update daily_questions: %w", err)
	}

	return nil
}

func (q *dailyQuestions) Count() (int64, error) {
	res := struct {
		Count int64 `db:"count"`
	}{}

	if err := q.db.Get(&res, q.counter); err != nil {
		return 0, fmt.Errorf("get daily_questions: %w", err)
	}

	return res.Count, nil
}

func (q *dailyQuestions) Select() ([]data.DailyQuestions, error) {
	var res []data.DailyQuestions
	if err := q.db.Select(&res, q.selector); err != nil {
		return res, fmt.Errorf("select daily_questions: %w", err)
	}
	return res, nil
}

func (q *dailyQuestions) FilteredActive(status bool) data.DailyQuestionsQ {
	return q.applyCondition(squirrel.Eq{"active": status})
}

func (q *dailyQuestions) FilteredStartAt(date int) data.DailyQuestionsQ {
	res := q.applyCondition(squirrel.Gt{"starts_at": date})
	return res
}

func (q *dailyQuestions) applyCondition(cond squirrel.Sqlizer) data.DailyQuestionsQ {
	q.selector = q.selector.Where(cond)
	q.updater = q.updater.Where(cond)
	q.counter = q.counter.Where(cond)
	return q
}
