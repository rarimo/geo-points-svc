package pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/rarimo/geo-points-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const bonusCodesTable = "bonus_codes"

type bonusCodesQ struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
	counter  squirrel.SelectBuilder
}

func NewBonusCodesQ(db *pgdb.DB) data.BonusCodesQ {
	return &bonusCodesQ{
		db:       db,
		selector: squirrel.Select("id", bonusCodesTable+".nullifier AS nullifier", "usage_count", "infinity", "reward").From(bonusCodesTable),
		updater:  squirrel.Update(bonusCodesTable),
		counter:  squirrel.Select("COUNT(*) as count").From(bonusCodesTable),
	}
}

func (q *bonusCodesQ) New() data.BonusCodesQ {
	return NewBonusCodesQ(q.db)
}

func (q *bonusCodesQ) Insert(bonusCodes ...data.BonusCode) error {
	if len(bonusCodes) == 0 {
		return nil
	}

	stmt := squirrel.Insert(bonusCodesTable).Columns("id", "nullifier", "reward", "usage_count", "infinity")
	for _, bonusCode := range bonusCodes {
		stmt = stmt.Values(bonusCode.ID, bonusCode.Nullifier, bonusCode.Reward, bonusCode.UsageCount, bonusCode.Infinity)
	}

	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("insert bonus codes: %w", err)
	}

	return nil
}

func (q *bonusCodesQ) Update(values map[string]any) error {

	if err := q.db.Exec(q.updater.SetMap(values)); err != nil {
		return fmt.Errorf("update bonusCode: %w", err)
	}

	return nil
}

func (q *bonusCodesQ) Select() ([]data.BonusCode, error) {
	var res []data.BonusCode

	if err := q.db.Select(&res, q.selector); err != nil {
		return nil, fmt.Errorf("select bonusCodes: %w", err)
	}

	return res, nil
}

func (q *bonusCodesQ) Get() (*data.BonusCode, error) {
	var res data.BonusCode

	if err := q.db.Get(&res, q.selector); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get bonusCode: %w", err)
	}

	return &res, nil
}

func (q *bonusCodesQ) Page(page *pgdb.OffsetPageParams) data.BonusCodesQ {
	q.selector = page.ApplyTo(q.selector, "updated_at")
	return q
}

func (q *bonusCodesQ) Count() (uint64, error) {
	var res struct {
		Count uint64 `db:"count"`
	}

	if err := q.db.Get(&res, q.counter); err != nil {
		return 0, fmt.Errorf("count bonusCodes: %w", err)
	}

	return res.Count, nil
}

func (q *bonusCodesQ) FilterByNullifier(nullifiers ...string) data.BonusCodesQ {
	return q.applyCondition(squirrel.Eq{fmt.Sprintf("%s.nullifier", bonusCodesTable): nullifiers})
}

func (q *bonusCodesQ) FilterByID(ids ...string) data.BonusCodesQ {
	return q.applyCondition(squirrel.Eq{fmt.Sprintf("%s.id", bonusCodesTable): ids})
}

func (q *bonusCodesQ) applyCondition(cond squirrel.Sqlizer) data.BonusCodesQ {
	q.selector = q.selector.Where(cond)
	q.updater = q.updater.Where(cond)
	q.counter = q.counter.Where(cond)
	return q
}
