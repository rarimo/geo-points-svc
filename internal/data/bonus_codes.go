package data

import (
	"database/sql"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	ColNullifier  = "nullifier"
	ColUsageCount = "usage_count"
	ColInfinity   = "infinity"
)

type BonusCode struct {
	ID         string         `db:"id"`
	Nullifier  sql.NullString `db:"nullifier"`
	Reward     int            `db:"reward"`
	UsageCount int            `db:"usage_count"`
	Infinity   bool           `db:"infinity"`

	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

type BonusCodesQ interface {
	New() BonusCodesQ
	Insert(...BonusCode) error
	Update(map[string]any) error

	Page(*pgdb.OffsetPageParams) BonusCodesQ

	Get() (*BonusCode, error)
	Select() ([]BonusCode, error)
	Count() (uint64, error)

	FilterByID(...string) BonusCodesQ
	FilterByNullifier(...string) BonusCodesQ
}
