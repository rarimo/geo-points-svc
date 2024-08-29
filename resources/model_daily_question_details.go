/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestionDetails struct {
	Key
	Attributes DailyQuestionDetailsAttributes `json:"attributes"`
}
type DailyQuestionDetailsResponse struct {
	Data     DailyQuestionDetails `json:"data"`
	Included Included             `json:"included"`
}

type DailyQuestionDetailsListResponse struct {
	Data     []DailyQuestionDetails `json:"data"`
	Included Included               `json:"included"`
	Links    *Links                 `json:"links"`
	Meta     json.RawMessage        `json:"meta,omitempty"`
}

func (r *DailyQuestionDetailsListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionDetailsListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestionDetails - returns DailyQuestionDetails from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestionDetails(key Key) *DailyQuestionDetails {
	var dailyQuestionDetails DailyQuestionDetails
	if c.tryFindEntry(key, &dailyQuestionDetails) {
		return &dailyQuestionDetails
	}
	return nil
}
