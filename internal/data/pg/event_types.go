package pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const eventTypesTable = "event_types"

type eventTypes struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
}

func NewEventTypes(db *pgdb.DB) data.EventTypesQ {
	return &eventTypes{
		db:       db,
		selector: squirrel.Select("*").From(eventTypesTable),
		updater:  squirrel.Update(eventTypesTable),
	}
}

func (q *eventTypes) New() data.EventTypesQ {
	return NewEventTypes(q.db)
}

func (q *eventTypes) Insert(eventTypes ...models.EventType) error {
	if len(eventTypes) == 0 {
		return nil
	}

	stmt := squirrel.Insert(eventTypesTable).Columns(
		"name",
		"description",
		"short_description",
		"reward",
		"title",
		"frequency",
		"starts_at",
		"expires_at",
		"no_auto_open",
		"auto_claim",
		"disabled",
		"action_url",
		"logo",
		"qr_code_value",
	)
	for _, eventType := range eventTypes {
		stmt = stmt.Values(
			eventType.Name,
			eventType.Description,
			eventType.ShortDescription,
			eventType.Reward,
			eventType.Title,
			eventType.Frequency,
			eventType.StartsAt,
			eventType.ExpiresAt,
			eventType.NoAutoOpen,
			eventType.AutoClaim,
			eventType.Disabled,
			eventType.ActionURL,
			eventType.Logo,
			eventType.QRCodeValue,
		)
	}

	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("insert event types [%+v]: %w", eventTypes, err)
	}

	return nil
}

func (q *eventTypes) Update(fields map[string]any) error {
	stmt := q.updater.SetMap(fields)

	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("update event type with map %+v: %w", fields, err)
	}

	return nil
}

func (q *eventTypes) Transaction(f func() error) error {
	return q.db.Transaction(f)
}

func (q *eventTypes) Select() ([]models.EventType, error) {
	var res []models.EventType

	if err := q.db.Select(&res, q.selector); err != nil {
		return nil, fmt.Errorf("select event types: %w", err)
	}

	return res, nil
}

func (q *eventTypes) Get(name string) (*models.EventType, error) {
	stmt := q.selector.Where(squirrel.Eq{"name": name})

	var res models.EventType
	if err := q.db.Get(&res, stmt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get event type: %w", err)
	}

	return &res, nil
}

func (q *eventTypes) FilterByNames(names ...string) data.EventTypesQ {
	q.selector = q.selector.Where(squirrel.Eq{"name": names})
	q.updater = q.updater.Where(squirrel.Eq{"name": names})
	return q
}
