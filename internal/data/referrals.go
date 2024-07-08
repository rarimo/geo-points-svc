package data

const (
	StatusInfinity = "infinity"
	StatusActive   = "active"
	StatusAwaiting = "awaiting"
	StatusRewarded = "rewarded"
	StatusConsumed = "consumed"
)

type Referral struct {
	ID        string `db:"id"`
	Nullifier string `db:"nullifier"`
	UsageLeft int32  `db:"usage_left"`
	Infinity  bool   `db:"infinity"`
	Status    string `db:"status"`
}

type ReferralsQ interface {
	New() ReferralsQ
	Insert(...Referral) error

	Select() ([]Referral, error)
	Get(id string) (*Referral, error)
	Count() (uint64, error)
	Consume(id string) error

	WithStatus() ReferralsQ
	Update(usageLeft int, infinity bool) (*Referral, error)

	FilterByNullifier(string) ReferralsQ
	FilterInactive() ReferralsQ
}
