package data

import (
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

type Withdrawal struct {
	TxHash    []byte    `db:"tx_hash"`
	Nullifier string    `db:"nullifier"`
	Amount    int64     `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}

type WithdrawalsQ interface {
	New() WithdrawalsQ
	Insert(Withdrawal) (*Withdrawal, error)

	Page(*pgdb.CursorPageParams) WithdrawalsQ
	Select() ([]Withdrawal, error)

	FilterByNullifier(string) WithdrawalsQ
}
