/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BalanceAttributes struct {
	// Amount of points
	Amount int64 `json:"amount"`
	// Unix timestamp of balance creation
	CreatedAt int32 `json:"created_at"`
	// Whether the user was not referred with some code. If it wasn't - balance is disabled and very limited in functionality.
	IsDisabled *bool `json:"is_disabled,omitempty"`
	// Whether the user has scanned passport. Returned only for the single user.
	IsVerified *bool `json:"is_verified,omitempty"`
	// The level indicates user permissions and features
	Level int `json:"level"`
	// Rank of the user in the full leaderboard. Returned only for the single user.
	Rank *int `json:"rank,omitempty"`
	// Referral codes. Returned only for the single active balance.
	ReferralCodes *[]ReferralCode `json:"referral_codes,omitempty"`
	// Number of invited users. Returned only for the single active balance.
	ReferredUsersCount *int `json:"referred_users_count,omitempty"`
	// Number of users for whom the reward was received. Returned only for the single active balance.
	RewardedReferredUsersCount *int `json:"rewarded_referred_users_count,omitempty"`
	// Unix timestamp of the last points accruing
	UpdatedAt int32 `json:"updated_at"`
}
