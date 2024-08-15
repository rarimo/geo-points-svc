/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestionResult struct {
	Key
	Attributes DailyQuestionResultAttributes `json:"attributes"`
}
type DailyQuestionResultRequest struct {
	Data     DailyQuestionResult `json:"data"`
	Included Included            `json:"included"`
}

type DailyQuestionResultListRequest struct {
	Data     []DailyQuestionResult `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
	Meta     json.RawMessage       `json:"meta,omitempty"`
}

func (r *DailyQuestionResultListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionResultListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestionResult - returns DailyQuestionResult from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestionResult(key Key) *DailyQuestionResult {
	var dailyQuestionResult DailyQuestionResult
	if c.tryFindEntry(key, &dailyQuestionResult) {
		return &dailyQuestionResult
	}
	return nil
}
