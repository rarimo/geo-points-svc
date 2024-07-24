package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	ColAmount      = "amount"
	ColReferredBy  = "referred_by"
	ColLevel       = "level"
	ColAnonymousID = "anonymous_id"
	ColSharedHash  = "shared_hash"
	ColIsVerified  = "is_verified"
)

type Balance struct {
	Nullifier   string  `db:"nullifier"`
	Amount      int64   `db:"amount"`
	CreatedAt   int32   `db:"created_at"`
	UpdatedAt   int32   `db:"updated_at"`
	ReferredBy  *string `db:"referred_by"`
	Level       int     `db:"level"`
	AnonymousID *string `db:"anonymous_id"`
	SharedHash  *string `db:"shared_hash"`
	IsVerified  bool    `db:"is_verified"`
	Rank        *int    `db:"rank"`
}

type BalancesQ interface {
	New() BalancesQ
	Insert(Balance) error
	Update(map[string]any) error

	Page(*pgdb.OffsetPageParams) BalancesQ
	Select() ([]Balance, error)
	Get() (*Balance, error)
	// GetWithRank returns balance with rank, filtered by nullifier. No other filters can be applied.
	GetWithRank(nullifier string) (*Balance, error)
	SelectWithRank() ([]Balance, error)

	// WithoutPassportEvent returns balances which already
	// have scanned passport, but there no claimed events
	// for this. Filters are not applied.
	WithoutPassportEvent() ([]WithoutPassportEventBalance, error)
	WithoutReferralEvent() ([]ReferredReferrer, error)

	Count() (int64, error)

	FilterByNullifier(...string) BalancesQ
	FilterDisabled() BalancesQ
	FilterByAnonymousID(id string) BalancesQ
}

type WithoutPassportEventBalance struct {
	Balance
	EventID     string      `db:"event_id"`
	EventStatus EventStatus `db:"event_status"`
}

type ReferredReferrer struct {
	Referred string `db:"referred"`
	Referrer string `db:"referrer"`
}
