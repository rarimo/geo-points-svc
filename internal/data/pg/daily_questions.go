package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/rarimo/geo-points-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const dailyQuestionsTable = "daily_questions"

type dailyQuestionsQ struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
	counter  squirrel.SelectBuilder
}

func NewDailyQuestionsQ(db *pgdb.DB) data.DailyQuestionQ {
	return &dailyQuestionsQ{
		db:       db,
		selector: squirrel.Select("*").From(dailyQuestionsTable),
		updater:  squirrel.Update(dailyQuestionsTable),
		counter:  squirrel.Select("COUNT(*) as count").From(dailyQuestionsTable),
	}
}

func (q *dailyQuestionsQ) New() data.DailyQuestionQ {
	return NewDailyQuestionsQ(q.db)
}

func (q *dailyQuestionsQ) Insert(quest data.DailyQuestion) error {
	stmt := squirrel.Insert(dailyQuestionsTable).SetMap(map[string]interface{}{
		"title":           quest.Title,
		"time_for_answer": quest.TimeForAnswer,
		"reward":          quest.Reward,
		"answer_options":  quest.AnswerOptions,
		"starts_at":       quest.StartsAt,
	})

	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("insert daily_questions %+v: %w", quest, err)
	}

	return nil
}

func (q *dailyQuestionsQ) Update(fields map[string]any) error {
	if err := q.db.Exec(q.updater.SetMap(fields)); err != nil {
		return fmt.Errorf("update daily_questions: %w", err)
	}

	return nil
}

func (q *dailyQuestionsQ) Count() (int64, error) {
	res := struct {
		Count int64 `db:"count"`
	}{}

	if err := q.db.Get(&res, q.counter); err != nil {
		return 0, fmt.Errorf("get daily_questions: %w", err)
	}

	return res.Count, nil
}

func (q *dailyQuestionsQ) Select() ([]data.DailyQuestion, error) {
	var res []data.DailyQuestion
	if err := q.db.Select(&res, q.selector); err != nil {
		return res, fmt.Errorf("select daily_questions: %w", err)
	}
	return res, nil
}

func (q *dailyQuestionsQ) Get() (*data.DailyQuestion, error) {
	var res data.DailyQuestion

	if err := q.db.Get(&res, q.selector); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get balance: %w", err)
	}

	return &res, nil
}

func (q *dailyQuestionsQ) FilterActive() data.DailyQuestionQ {
	now := time.Now().Unix()
	return q.applyCondition(squirrel.Expr("EXTRACT(EPOCH FROM starts_at) + time_for_answer > ? AND EXTRACT(EPOCH FROM starts_at) < ?", now, now))
}

func (q *dailyQuestionsQ) FilterByStartAt(date time.Time) data.DailyQuestionQ {
	return q.applyCondition(squirrel.Gt{"starts_at": date})

}

func (q *dailyQuestionsQ) FilterByCreatedAt(date time.Time) data.DailyQuestionQ {
	return q.applyCondition(squirrel.Gt{"created_at": date})
}

func (q *dailyQuestionsQ) FilterByID(ID int) data.DailyQuestionQ {
	return q.applyCondition(squirrel.Eq{"id": ID})
}

func (q *dailyQuestionsQ) applyCondition(cond squirrel.Sqlizer) data.DailyQuestionQ {
	q.selector = q.selector.Where(cond)
	q.updater = q.updater.Where(cond)
	q.counter = q.counter.Where(cond)
	return q
}
