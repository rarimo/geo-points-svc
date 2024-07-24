/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type FulfillPollEvent struct {
	Key
	Attributes FulfillPollEventAttributes `json:"attributes"`
}
type FulfillPollEventRequest struct {
	Data     FulfillPollEvent `json:"data"`
	Included Included         `json:"included"`
}

type FulfillPollEventListRequest struct {
	Data     []FulfillPollEvent `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
	Meta     json.RawMessage    `json:"meta,omitempty"`
}

func (r *FulfillPollEventListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *FulfillPollEventListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustFulfillPollEvent - returns FulfillPollEvent from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustFulfillPollEvent(key Key) *FulfillPollEvent {
	var fulfillPollEvent FulfillPollEvent
	if c.tryFindEntry(key, &fulfillPollEvent) {
		return &fulfillPollEvent
	}
	return nil
}
