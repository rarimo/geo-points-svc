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

func NewDailyQuestionsQ(db *pgdb.DB) data.DailyQuestionsQ {
	return &dailyQuestionsQ{
		db:       db,
		selector: squirrel.Select("*").From(dailyQuestionsTable),
		updater:  squirrel.Update(dailyQuestionsTable),
		counter:  squirrel.Select("COUNT(*) as count").From(dailyQuestionsTable),
	}
}

func (q *dailyQuestionsQ) New() data.DailyQuestionsQ {
	return NewDailyQuestionsQ(q.db)
}

func (q *dailyQuestionsQ) Insert(quest data.DailyQuestion) error {
	stmt := squirrel.Insert(dailyQuestionsTable).SetMap(map[string]interface{}{
		"title":                 quest.Title,
		"time_for_answer":       quest.TimeForAnswer,
		"reward":                quest.Reward,
		"answer_options":        quest.AnswerOptions,
		"starts_at":             quest.StartsAt,
		"correct_answer":        quest.CorrectAnswer,
		"num_correct_answers":   quest.NumCorrectAnswers,
		"num_incorrect_answers": quest.NumIncorrectAnswers,
		"num_all_participants":  quest.NumAllParticipants,
	})

	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("insert daily questions %+v: %w", quest, err)
	}

	return nil
}

func (q *dailyQuestionsQ) Update(fields map[string]any) error {
	if err := q.db.Exec(q.updater.SetMap(fields)); err != nil {
		return fmt.Errorf("update daily questions: %w", err)
	}

	return nil
}

func (q *dailyQuestionsQ) Count() (int64, error) {
	res := struct {
		Count int64 `db:"count"`
	}{}

	if err := q.db.Get(&res, q.counter); err != nil {
		return 0, fmt.Errorf("count daily questions: %w", err)
	}

	return res.Count, nil
}

func (q *dailyQuestionsQ) Select() ([]data.DailyQuestion, error) {
	var res []data.DailyQuestion
	if err := q.db.Select(&res, q.selector); err != nil {
		return res, fmt.Errorf("select daily questions: %w", err)
	}
	return res, nil
}

func (q *dailyQuestionsQ) Get() (*data.DailyQuestion, error) {
	var res data.DailyQuestion

	if err := q.db.Get(&res, q.selector); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get daily question: %w", err)
	}

	return &res, nil
}

func (q *dailyQuestionsQ) FilterByCreatedAt(date time.Time) data.DailyQuestionsQ {
	return q.applyCondition(squirrel.Gt{"created_at": date})
}

func (q *dailyQuestionsQ) FilterTodayQuestions(offset int) data.DailyQuestionsQ {
	location := time.FixedZone(fmt.Sprintf("GMT%+d", offset), offset*3600)

	now := time.Now().In(location)
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location).UTC()
	todayEnd := todayStart.Add(24 * time.Hour).Add(-time.Nanosecond).UTC()

	return q.applyCondition(squirrel.And{
		squirrel.GtOrEq{"starts_at": todayStart},
		squirrel.LtOrEq{"starts_at": todayEnd},
	})
}

func (q *dailyQuestionsQ) FilterByID(ID int64) data.DailyQuestionsQ {
	return q.applyCondition(squirrel.Eq{"id": ID})
}

func (q *dailyQuestionsQ) IncrementCorrectAnswer() error {
	stmt := q.updater.Set("num_correct_answers", squirrel.Expr("num_correct_answers + 1"))
	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("increment correct answer: %w", err)
	}
	return nil
}

func (q *dailyQuestionsQ) IncrementIncorrectAnswer() error {
	stmt := q.updater.
		Set("num_incorrect_answers", squirrel.Expr("num_incorrect_answers + 1"))
	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("increment incorrect answer: %w", err)
	}
	return nil
}

func (q *dailyQuestionsQ) IncrementAllParticipants() error {
	stmt := q.updater.
		Set("num_all_participants", squirrel.Expr("num_all_participants + 1"))
	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("increment no answer: %w", err)
	}
	return nil
}

func (q *dailyQuestionsQ) applyCondition(cond squirrel.Sqlizer) data.DailyQuestionsQ {
	q.selector = q.selector.Where(cond)
	q.updater = q.updater.Where(cond)
	q.counter = q.counter.Where(cond)
	return q
}