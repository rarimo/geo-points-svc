/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "github.com/iden3/go-rapidsnark/types"

type WithdrawAttributes struct {
	// Amount of points to withdraw
	Amount int64 `json:"amount"`
	// Query ZK passport verification proof.
	Proof types.ZKProof `json:"proof"`
}
