package pg

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/rarimo/geo-points-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const withdrawalsTable = "withdrawals"

type withdrawals struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewWithdrawals(db *pgdb.DB) data.WithdrawalsQ {
	return &withdrawals{
		db:       db,
		selector: squirrel.Select("*").From(withdrawalsTable),
	}
}

func (q *withdrawals) New() data.WithdrawalsQ {
	return NewWithdrawals(q.db)
}

func (q *withdrawals) Insert(w data.Withdrawal) (*data.Withdrawal, error) {
	var res data.Withdrawal
	stmt := squirrel.Insert(withdrawalsTable).SetMap(map[string]interface{}{
		"tx_hash":   w.TxHash,
		"nullifier": w.Nullifier,
		"amount":    w.Amount,
	}).Suffix("RETURNING *")

	if err := q.db.Get(&res, stmt); err != nil {
		return nil, fmt.Errorf("insert withdrawal [%+v]: %w", w, err)
	}

	return &res, nil
}

func (q *withdrawals) Page(page *pgdb.CursorPageParams) data.WithdrawalsQ {
	q.selector = page.ApplyTo(q.selector, "created_at")
	return q
}

func (q *withdrawals) Select() ([]data.Withdrawal, error) {
	var res []data.Withdrawal

	if err := q.db.Select(&res, q.selector); err != nil {
		return nil, fmt.Errorf("select withdrawals: %w", err)
	}

	return res, nil
}

func (q *withdrawals) FilterByNullifier(nullifier string) data.WithdrawalsQ {
	q.selector = q.selector.Where(squirrel.Eq{"nullifier": nullifier})
	return q
}
