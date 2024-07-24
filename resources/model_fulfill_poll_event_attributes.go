/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "github.com/iden3/go-rapidsnark/types"

type FulfillPollEventAttributes struct {
	// Proof of voting in some poll. The poll ID must be equal to `meta.static.poll_id` from the event.
	Proof types.ZKProof `json:"proof"`
}
