/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestionStatus struct {
	Key
	Attributes DailyQuestionStatusAttributes `json:"attributes"`
}
type DailyQuestionStatusResponse struct {
	Data     DailyQuestionStatus `json:"data"`
	Included Included            `json:"included"`
}

type DailyQuestionStatusListResponse struct {
	Data     []DailyQuestionStatus `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
	Meta     json.RawMessage       `json:"meta,omitempty"`
}

func (r *DailyQuestionStatusListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionStatusListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestionStatus - returns DailyQuestionStatus from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestionStatus(key Key) *DailyQuestionStatus {
	var dailyQuestionStatus DailyQuestionStatus
	if c.tryFindEntry(key, &dailyQuestionStatus) {
		return &dailyQuestionStatus
	}
	return nil
}
