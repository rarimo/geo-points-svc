/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type EventClaimingState struct {
	Key
	Attributes EventClaimingStateAttributes `json:"attributes"`
}
type EventClaimingStateResponse struct {
	Data     EventClaimingState `json:"data"`
	Included Included           `json:"included"`
}

type EventClaimingStateListResponse struct {
	Data     []EventClaimingState `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
	Meta     json.RawMessage      `json:"meta,omitempty"`
}

func (r *EventClaimingStateListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *EventClaimingStateListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustEventClaimingState - returns EventClaimingState from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEventClaimingState(key Key) *EventClaimingState {
	var eventClaimingState EventClaimingState
	if c.tryFindEntry(key, &eventClaimingState) {
		return &eventClaimingState
	}
	return nil
}
