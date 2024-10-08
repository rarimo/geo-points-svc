/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BonusCodeAttributes struct {
	// For creating personal bonus codes
	Nullifier *string `json:"nullifier,omitempty"`
	// Reward for this bonus code
	Reward *int `json:"reward,omitempty"`
	// Specify how many times bonus code can be scaned. Omit if bonus code must have infinity usage count
	UsageCount *int `json:"usage_count,omitempty"`
}
