package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	ColAmount      = "amount"
	ColReferredBy  = "referred_by"
	ColLevel       = "level"
	ColInternalAID = "internal_aid"
	ColExternalAID = "external_aid"
	ColSharedHash  = "shared_hash"

	VerifyInternalType = "internal"
	VerifyExternalType = "external"
)

type Balance struct {
	Nullifier   string  `db:"nullifier"`
	Amount      int64   `db:"amount"`
	CreatedAt   int32   `db:"created_at"`
	UpdatedAt   int32   `db:"updated_at"`
	ReferredBy  *string `db:"referred_by"`
	Level       int     `db:"level"`
	InternalAID *string `db:"internal_aid"`
	ExternalAID *string `db:"external_aid"`
	SharedHash  *string `db:"shared_hash"`
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
	FilterByInternalAID(aid string) BalancesQ
	FilterByExternalAID(aid string) BalancesQ
	FilterBySharedHash(hash string) BalancesQ
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

func (b *Balance) IsVerified() bool {
	return b.InternalAID != nil || b.ExternalAID != nil
}

func (b *Balance) IsDisabled() bool {
	return b.ReferredBy == nil
}
