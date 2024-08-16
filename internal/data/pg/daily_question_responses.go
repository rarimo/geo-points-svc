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

const dailyQuestionAnswerTable = "daily_question_responses"

type DailyQuestionsAnswerQ struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
	counter  squirrel.SelectBuilder
}

func NewDailyQuestionsAnswerQ(db *pgdb.DB) data.DailyQuestionsResponsesQ {
	return &DailyQuestionsAnswerQ{
		db:       db,
		selector: squirrel.Select("*").From(dailyQuestionAnswerTable),
		updater:  squirrel.Update(dailyQuestionAnswerTable),
		counter:  squirrel.Select("COUNT(*) as count").From(dailyQuestionAnswerTable),
	}
}

func (q *DailyQuestionsAnswerQ) New() data.DailyQuestionsResponsesQ {
	return NewDailyQuestionsAnswerQ(q.db)
}

func (q *DailyQuestionsAnswerQ) Insert(answer data.DailyQuestionsResponses) error {
	stmt := squirrel.Insert(dailyQuestionAnswerTable).SetMap(map[string]interface{}{
		"daily_question_id": answer.DailyQuestionId,
		"nullifier":         answer.Nullifier,
		"created_at":        answer.CreatedAt,
	})

	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("insert daily_questions_answer %+v: %w", answer, err)
	}

	return nil
}

func (q *DailyQuestionsAnswerQ) Update(fields map[string]any) error {
	if err := q.db.Exec(q.updater.SetMap(fields)); err != nil {
		return fmt.Errorf("update daily_question_responses: %w", err)
	}

	return nil
}

func (q *DailyQuestionsAnswerQ) Count() (int64, error) {
	res := struct {
		Count int64 `db:"count"`
	}{}

	if err := q.db.Get(&res, q.counter); err != nil {
		return 0, fmt.Errorf("get daily_question_responses: %w", err)
	}

	return res.Count, nil
}

func (q *DailyQuestionsAnswerQ) Select() ([]data.DailyQuestionsResponses, error) {
	var res []data.DailyQuestionsResponses
	if err := q.db.Select(&res, q.selector); err != nil {
		return res, fmt.Errorf("select daily_question_responses: %w", err)
	}
	return res, nil
}

func (q *DailyQuestionsAnswerQ) Get() (*data.DailyQuestionsResponses, error) {
	var res data.DailyQuestionsResponses
	if err := q.db.Get(&res, q.selector); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get daily_question_responses: %w", err)
	}

	return &res, nil
}

func (q *DailyQuestionsAnswerQ) Transaction(f func() error) error {
	return q.db.Transaction(f)
}

func (q *DailyQuestionsAnswerQ) FilterByQuestionID(ID int) data.DailyQuestionsResponsesQ {
	return q.applyCondition(squirrel.Eq{"daily_question_id": ID})
}

func (q *DailyQuestionsAnswerQ) FilterByNullifier(nullifier string) data.DailyQuestionsResponsesQ {
	return q.applyCondition(squirrel.Eq{"nullifier": nullifier})
}

func (q *DailyQuestionsAnswerQ) FilterByCreatedAt(date time.Time) data.DailyQuestionsResponsesQ {
	return q.applyCondition(squirrel.Eq{"created_at": date})
}

func (q *DailyQuestionsAnswerQ) FilterByCreatedAfter(after int64) data.DailyQuestionsResponsesQ {
	return q.applyCondition(squirrel.Expr("created_at > TO_TIMESTAMP(?)", after))
}

func (q *DailyQuestionsAnswerQ) FilterByCreatedBefore(before int64) data.DailyQuestionsResponsesQ {
	return q.applyCondition(squirrel.Expr("created_at < TO_TIMESTAMP(?)", before))
}

func (q *DailyQuestionsAnswerQ) applyCondition(cond squirrel.Sqlizer) data.DailyQuestionsResponsesQ {
	q.selector = q.selector.Where(cond)
	q.updater = q.updater.Where(cond)
	q.counter = q.counter.Where(cond)
	return q
}
