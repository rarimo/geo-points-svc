/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestionsStatus struct {
	Key
	Attributes DailyQuestionsStatusAttributes `json:"attributes"`
}
type DailyQuestionsStatusResponse struct {
	Data     DailyQuestionsStatus `json:"data"`
	Included Included             `json:"included"`
}

type DailyQuestionsStatusListResponse struct {
	Data     []DailyQuestionsStatus `json:"data"`
	Included Included               `json:"included"`
	Links    *Links                 `json:"links"`
	Meta     json.RawMessage        `json:"meta,omitempty"`
}

func (r *DailyQuestionsStatusListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionsStatusListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestionsStatus - returns DailyQuestionsStatus from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestionsStatus(key Key) *DailyQuestionsStatus {
	var dailyQuestionsStatus DailyQuestionsStatus
	if c.tryFindEntry(key, &dailyQuestionsStatus) {
		return &dailyQuestionsStatus
	}
	return nil
}
