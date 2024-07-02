package data

type Referral struct {
	ID        string `db:"id"`
	Nullifier string `db:"nullifier"`
	UsageLeft int32  `db:"usage_left"`
}

type ReferralsQ interface {
	New() ReferralsQ
	Insert(...Referral) error

	Select() ([]Referral, error)
	Get(id string) (*Referral, error)
	Count() (uint64, error)

	Update(usageLeft int) (*Referral, error)

	FilterByNullifier(string) ReferralsQ
}