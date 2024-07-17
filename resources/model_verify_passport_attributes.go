/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "github.com/iden3/go-rapidsnark/types"

type VerifyPassportAttributes struct {
	// Unique identifier of the passport.
	AnonymousId string `json:"anonymous_id"`
	// Query ZK passport verification proof. Required for passport verification endpoint.
	Proof *types.ZKProof `json:"proof,omitempty"`
}
