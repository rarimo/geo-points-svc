/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateBalanceAttributes struct {
	// Referral code from the link. Supply it to create the active balance, otherwise disabled balance is created, and it can be activated later.  Disabled balance is only allowed to verify passport and get.
	ReferredBy *string `json:"referred_by,omitempty"`
}
