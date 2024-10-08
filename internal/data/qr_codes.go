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

type QRCode struct {
	ID         string         `db:"id"`
	Nullifier  sql.NullString `db:"nullifier"`
	Reward     int            `db:"reward"`
	UsageCount int            `db:"usage_count"`
	Infinity   bool           `db:"infinity"`

	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

type QRCodesQ interface {
	New() QRCodesQ
	Insert(...QRCode) error
	Update(map[string]any) error

	Page(*pgdb.OffsetPageParams) QRCodesQ

	Get() (*QRCode, error)
	Select() ([]QRCode, error)
	Count() (uint64, error)

	FilterByID(...string) QRCodesQ
	FilterByNullifier(...string) QRCodesQ
}
