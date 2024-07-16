/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "github.com/iden3/go-rapidsnark/types"

type VerifyPassportAttributes struct {
	// Unique identifier of the passport.
	AnonymousId string `json:"anonymous_id"`
	// Query ZK passport verification proof. Required for endpoint `/v1/balances/{nullifier}/verifypassport`.
	Proof *types.ZKProof `json:"proof,omitempty"`
}
