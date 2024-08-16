/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "github.com/iden3/go-rapidsnark/types"

type FulfillPollEventAttributes struct {
	// Proof of voting in some poll.
	Proof types.ZKProof `json:"proof"`
	// Vote proposal id
	ProposalId string `json:"proposal_id"`
}
