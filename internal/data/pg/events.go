package pg

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/rarimo/geo-points-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const eventsTable = "events"

type events struct {
	db         *pgdb.DB
	selector   squirrel.SelectBuilder
	updater    squirrel.UpdateBuilder
	deleter    squirrel.DeleteBuilder
	counter    squirrel.SelectBuilder
	reopenable squirrel.SelectBuilder
	last       squirrel.SelectBuilder
}

func NewEvents(db *pgdb.DB) data.EventsQ {
	return &events{
		db:         db,
		selector:   squirrel.Select("*").From(eventsTable),
		updater:    squirrel.Update(eventsTable),
		deleter:    squirrel.Delete(eventsTable),
		counter:    squirrel.Select("COUNT(*) AS count").From(eventsTable),
		reopenable: squirrel.Select("nullifier", "type").Distinct().From(eventsTable + " e1"),
		last:       squirrel.Select("*").From(eventsTable).OrderBy("created_at DESC"),
	}
}

func (q *events) New() data.EventsQ {
	return NewEvents(q.db)
}

func (q *events) Insert(events ...data.Event) error {
	if len(events) == 0 {
		return nil
	}

	stmt := squirrel.Insert(eventsTable).
		Columns("nullifier", "type", "status", "meta", "points_amount", "external_id")
	for _, event := range events {
		var meta any
		if len(event.Meta) != 0 {
			meta = event.Meta
		}
		stmt = stmt.Values(event.Nullifier, event.Type, event.Status, meta, event.PointsAmount, event.ExternalID)
	}

	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("insert events [%+v]: %w", events, err)
	}

	return nil
}

func (q *events) Update(status data.EventStatus, meta json.RawMessage, points *int64) ([]data.Event, error) {
	umap := map[string]any{
		"status": status,
	}
	if len(meta) != 0 {
		umap["meta"] = meta
	}
	if points != nil {
		umap["points_amount"] = points
	}

	var res []data.Event
	stmt := q.updater.SetMap(umap).Suffix("RETURNING *")

	if err := q.db.Select(&res, stmt); err != nil {
		return nil, fmt.Errorf("update event with map %+v: %w", umap, err)
	}

	return res, nil
}

func (q *events) Delete() (int64, error) {
	res, err := q.db.ExecWithResult(q.deleter)
	if err != nil {
		return 0, fmt.Errorf("delete events: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("count rows affected: %w", err)
	}

	return rows, nil
}

func (q *events) Transaction(f func() error) error {
	return q.db.Transaction(f)
}

func (q *events) Page(page *pgdb.OffsetPageParams) data.EventsQ {
	ord := fmt.Sprintf("case when status = '%s' then 1 when status = '%s' then 2 when status = '%s' then 3 end", data.EventFulfilled, data.EventOpen, data.EventClaimed)
	q.selector = page.ApplyTo(q.selector.OrderBy(ord), "updated_at")
	return q
}

func (q *events) Select() ([]data.Event, error) {
	var res []data.Event

	if err := q.db.Select(&res, q.selector); err != nil {
		return nil, fmt.Errorf("select events: %w", err)
	}

	return res, nil
}

func (q *events) Get() (*data.Event, error) {
	var res data.Event

	if err := q.db.Get(&res, q.selector); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get event: %w", err)
	}

	return &res, nil
}

func (q *events) GetLast() (*data.Event, error) {
	var res data.Event

	if err := q.db.Get(&res, q.last); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get last event: %w", err)
	}

	return &res, nil
}

func (q *events) Count() (int, error) {
	var res struct {
		Count int `db:"count"`
	}

	if err := q.db.Get(&res, q.counter); err != nil {
		return 0, fmt.Errorf("count events: %w", err)
	}

	return res.Count, nil
}

func (q *events) SelectReopenable() ([]data.ReopenableEvent, error) {
	// including disabled balances to open events intentionally
	// join balances and filter by referred_by if you need the opposite
	subq := fmt.Sprintf(`NOT EXISTS (
	SELECT 1 FROM %s e2
    WHERE e2.nullifier = e1.nullifier
    AND e2.type = e1.type
    AND e2.status IN (?, ?))`, eventsTable)
	stmt := q.reopenable.Where(subq, data.EventOpen, data.EventFulfilled)

	var res []data.ReopenableEvent
	if err := q.db.Select(&res, stmt); err != nil {
		return nil, fmt.Errorf("select reopenable events: %w", err)
	}

	return res, nil
}

func (q *events) SelectAbsentTypes(allTypes ...string) ([]data.ReopenableEvent, error) {
	interfaceTypes := make([]any, 0, len(allTypes))
	patternValues := make([]string, 0, len(allTypes))
	for _, types := range allTypes {
		interfaceTypes = append(interfaceTypes, types)
		patternValues = append(patternValues, "(?)")
	}

	// including disabled balances to open events intentionally
	stmt := squirrel.Select("b.nullifier, t.type").
		From(fmt.Sprintf("%s b", balancesTable)).
		Prefix(fmt.Sprintf("WITH types(type) AS (VALUES %s)", strings.Join(patternValues, ", ")), interfaceTypes...).
		CrossJoin("types t").
		LeftJoin(fmt.Sprintf("%s e ON e.nullifier = b.nullifier AND e.type = t.type", eventsTable)).
		Where(squirrel.Eq{"e.type": nil})

	var res []data.ReopenableEvent
	if err := q.db.Select(&res, stmt); err != nil {
		return nil, fmt.Errorf("select absent types for each nullifier: %w", err)
	}

	return res, nil
}

func (q *events) FilterByID(ids ...string) data.EventsQ {
	if len(ids) == 0 {
		return q
	}
	return q.applyCondition(squirrel.Eq{"id": ids})
}

func (q *events) FilterByNullifier(nullifier ...string) data.EventsQ {
	return q.applyCondition(squirrel.Eq{"nullifier": nullifier})
}

func (q *events) FilterByStatus(statuses ...data.EventStatus) data.EventsQ {
	if len(statuses) == 0 {
		return q
	}
	return q.applyCondition(squirrel.Eq{"status": statuses})
}

func (q *events) FilterByType(types ...string) data.EventsQ {
	if len(types) == 0 {
		return q
	}
	return q.applyCondition(squirrel.Eq{"type": types})
}

func (q *events) FilterByNotType(types ...string) data.EventsQ {
	if len(types) == 0 {
		return q
	}
	return q.applyCondition(squirrel.NotEq{"type": types})
}

func (q *events) FilterByExternalID(id string) data.EventsQ {
	return q.applyCondition(squirrel.Eq{"external_id": id})
}

func (q *events) FilterByUpdatedAtBefore(unix int64) data.EventsQ {
	return q.applyCondition(squirrel.Lt{"updated_at": unix})
}

func (q *events) FilterTodayEvents(offset int) data.EventsQ {
	location := time.FixedZone(fmt.Sprintf("GMT%+d", offset), offset*3600)
	now := time.Now().In(location)

	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	todayEnd := todayStart.Add(24 * time.Hour).Add(-time.Nanosecond)

	utcStart := todayStart.UTC().Unix()
	utcEnd := todayEnd.UTC().Unix()

	res := q.applyCondition(squirrel.And{
		squirrel.GtOrEq{"created_at": utcStart},
		squirrel.LtOrEq{"created_at": utcEnd},
	})

	return res
}

func (q *events) FilterByQuestionID(id int) data.EventsQ {
	return q.applyCondition(squirrel.Eq{"meta->>'question_id'": id})
}

func (q *events) FilterByQRCode(qrCode string) data.EventsQ {
	return q.applyCondition(squirrel.Eq{"meta->>'qr_code'": qrCode})
}

func (q *events) FilterInactiveNotClaimed(types ...string) data.EventsQ {
	if len(types) == 0 {
		return q
	}
	return q.applyCondition(squirrel.Or{
		squirrel.NotEq{"type": types},
		squirrel.Eq{"status": data.EventClaimed},
	})
}

func (q *events) applyCondition(cond squirrel.Sqlizer) data.EventsQ {
	q.selector = q.selector.Where(cond)
	q.updater = q.updater.Where(cond)
	q.deleter = q.deleter.Where(cond)
	q.counter = q.counter.Where(cond)
	q.reopenable = q.reopenable.Where(cond)
	q.last = q.last.Where(cond)
	return q
}
