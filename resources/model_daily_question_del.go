/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestionDel struct {
	Key
	Attributes DailyQuestionDelAttributes `json:"attributes"`
}
type DailyQuestionDelResponse struct {
	Data     DailyQuestionDel `json:"data"`
	Included Included         `json:"included"`
}

type DailyQuestionDelListResponse struct {
	Data     []DailyQuestionDel `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
	Meta     json.RawMessage    `json:"meta,omitempty"`
}

func (r *DailyQuestionDelListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionDelListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestionDel - returns DailyQuestionDel from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestionDel(key Key) *DailyQuestionDel {
	var dailyQuestionDel DailyQuestionDel
	if c.tryFindEntry(key, &dailyQuestionDel) {
		return &dailyQuestionDel
	}
	return nil
}
