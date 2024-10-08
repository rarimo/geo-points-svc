/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type QrCodeAttributes struct {
	// For creating personal qr-codes
	Nullifier *string `json:"nullifier,omitempty"`
	// Reward for this qr-code
	Reward *int `json:"reward,omitempty"`
	// Specify how many times qr-code can be scaned. Omit if qr-code must have infinity usage count
	UsageCount *int `json:"usage_count,omitempty"`
}
