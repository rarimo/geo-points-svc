package pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/rarimo/geo-points-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const referralsTable = "referrals"

type referrals struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
	consumer squirrel.UpdateBuilder
	counter  squirrel.SelectBuilder
}

func NewReferrals(db *pgdb.DB) data.ReferralsQ {
	return &referrals{
		db:       db,
		selector: squirrel.Select("id", referralsTable+".nullifier AS nullifier", "usage_left").From(referralsTable),
		updater:  squirrel.Update(referralsTable),
		consumer: squirrel.Update(referralsTable).Set("usage_left", squirrel.Expr("usage_left - 1")),
		counter:  squirrel.Select("COUNT(*) as count").From(referralsTable),
	}
}

func (q *referrals) New() data.ReferralsQ {
	return NewReferrals(q.db)
}

func (q *referrals) Insert(referrals ...data.Referral) error {
	if len(referrals) == 0 {
		return nil
	}

	stmt := squirrel.Insert(referralsTable).Columns("id", "nullifier", "usage_left")
	for _, ref := range referrals {
		stmt = stmt.Values(ref.ID, ref.Nullifier, ref.UsageLeft)
	}

	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("insert referrals [%+v]: %w", referrals, err)
	}

	return nil
}

func (q *referrals) Update(usageLeft int) (*data.Referral, error) {
	var res data.Referral

	if err := q.db.Get(&res, q.updater.Set("usage_left", usageLeft).Suffix("RETURNING *")); err != nil {
		return nil, fmt.Errorf("update referral: %w", err)
	}

	return &res, nil
}

func (q *referrals) Select() ([]data.Referral, error) {
	var res []data.Referral

	if err := q.db.Select(&res, q.selector); err != nil {
		return nil, fmt.Errorf("select referrals: %w", err)
	}

	return res, nil
}

func (q *referrals) Get(id string) (*data.Referral, error) {
	var res data.Referral

	if err := q.db.Get(&res, q.selector.Where(squirrel.Eq{"id": id})); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get referral by ID: %w", err)
	}

	return &res, nil
}

func (q *referrals) Count() (uint64, error) {
	var res struct {
		Count uint64 `db:"count"`
	}

	if err := q.db.Get(&res, q.counter); err != nil {
		return 0, fmt.Errorf("count referrals: %w", err)
	}

	return res.Count, nil
}

func (q *referrals) FilterByNullifier(nullifier string) data.ReferralsQ {
	return q.applyCondition(squirrel.Eq{fmt.Sprintf("%s.nullifier", referralsTable): nullifier})
}

func (q *referrals) applyCondition(cond squirrel.Sqlizer) data.ReferralsQ {
	q.selector = q.selector.Where(cond)
	q.consumer = q.consumer.Where(cond)
	q.updater = q.updater.Where(cond)
	q.counter = q.counter.Where(cond)
	return q
}
