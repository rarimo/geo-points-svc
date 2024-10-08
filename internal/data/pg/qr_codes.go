package pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/rarimo/geo-points-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const qrCodesTable = "qr_codes"

type qrCodesQ struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
	counter  squirrel.SelectBuilder
}

func NewQRCodesQ(db *pgdb.DB) data.QRCodesQ {
	return &qrCodesQ{
		db:       db,
		selector: squirrel.Select("id", qrCodesTable+".nullifier AS nullifier", "usage_count", "infinity", "reward").From(qrCodesTable),
		updater:  squirrel.Update(qrCodesTable),
		counter:  squirrel.Select("COUNT(*) as count").From(qrCodesTable),
	}
}

func (q *qrCodesQ) New() data.QRCodesQ {
	return NewQRCodesQ(q.db)
}

func (q *qrCodesQ) Insert(qrCodes ...data.QRCode) error {
	if len(qrCodes) == 0 {
		return nil
	}

	stmt := squirrel.Insert(qrCodesTable).Columns("id", "nullifier", "reward", "usage_count", "infinity")
	for _, qr := range qrCodes {
		stmt = stmt.Values(qr.ID, qr.Nullifier, qr.Reward, qr.UsageCount, qr.Infinity)
	}

	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("insert qr codes: %w", err)
	}

	return nil
}

func (q *qrCodesQ) Update(values map[string]any) error {

	if err := q.db.Exec(q.updater.SetMap(values)); err != nil {
		return fmt.Errorf("update qrCode: %w", err)
	}

	return nil
}

func (q *qrCodesQ) Select() ([]data.QRCode, error) {
	var res []data.QRCode

	if err := q.db.Select(&res, q.selector); err != nil {
		return nil, fmt.Errorf("select qrCodes: %w", err)
	}

	return res, nil
}

func (q *qrCodesQ) Get() (*data.QRCode, error) {
	var res data.QRCode

	if err := q.db.Get(&res, q.selector); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get qrCode: %w", err)
	}

	return &res, nil
}

func (q *qrCodesQ) Page(page *pgdb.OffsetPageParams) data.QRCodesQ {
	q.selector = page.ApplyTo(q.selector, "updated_at")
	return q
}

func (q *qrCodesQ) Count() (uint64, error) {
	var res struct {
		Count uint64 `db:"count"`
	}

	if err := q.db.Get(&res, q.counter); err != nil {
		return 0, fmt.Errorf("count qrCodes: %w", err)
	}

	return res.Count, nil
}

func (q *qrCodesQ) FilterByNullifier(nullifiers ...string) data.QRCodesQ {
	return q.applyCondition(squirrel.Eq{fmt.Sprintf("%s.nullifier", qrCodesTable): nullifiers})
}

func (q *qrCodesQ) FilterByID(ids ...string) data.QRCodesQ {
	return q.applyCondition(squirrel.Eq{fmt.Sprintf("%s.id", qrCodesTable): ids})
}

func (q *qrCodesQ) applyCondition(cond squirrel.Sqlizer) data.QRCodesQ {
	q.selector = q.selector.Where(cond)
	q.updater = q.updater.Where(cond)
	q.counter = q.counter.Where(cond)
	return q
}
