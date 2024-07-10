/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

// Primary event metadata in plain JSON. This is a template to be filled by `dynamic` when it's present.  This structure is also reused as request body to event type creation and update.
type EventStaticMeta struct {
	// Page where you can fulfill the event
	ActionUrl *string `json:"action_url,omitempty"`
	// Whether the event is automatically claimed on fulfillment, or requires manual claim
	AutoClaim   bool   `json:"auto_claim"`
	Description string `json:"description"`
	// Whether the event is disabled in the system. Disabled events can only be retrieved.
	Disabled bool `json:"disabled"`
	// General event expiration date (UTC RFC3339)
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// Event configuration flag:   - active: Events can be opened, fulfilled, claimed   - not_started: Event are not available yet, see `starts_at`   - expired: Event is not available, as it has already expired, see `expires_at`   - disabled: Event is disabled in the system  If event is disabled, it doesn't matter if it's expired or not started: it has `disabled` flag.  Do not specify this field on creation: this structure is reused for request body too.
	Flag string `json:"flag"`
	// Event frequency, which means how often you can fulfill certain task and claim the reward.
	Frequency string `json:"frequency"`
	// Event logo
	Logo *string `json:"logo,omitempty"`
	// Unique event code name
	Name string `json:"name"`
	// Base64-encoded QR code. Must match the code provided in event type.
	QrCodeValue *string `json:"qr_code_value,omitempty"`
	// Reward amount in points
	Reward           int64  `json:"reward"`
	ShortDescription string `json:"short_description"`
	// General event starting date (UTC RFC3339)
	StartsAt *time.Time `json:"starts_at,omitempty"`
	Title    string     `json:"title"`
	// Number of uses. Only available to the administrator.
	UsageCount *int `json:"usage_count,omitempty"`
}
