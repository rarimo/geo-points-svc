/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestion struct {
	Key
	Attributes DailyQuestionAttributes `json:"attributes"`
}
type DailyQuestionRequest struct {
	Data     DailyQuestion `json:"data"`
	Included Included      `json:"included"`
}

type DailyQuestionListRequest struct {
	Data     []DailyQuestion `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *DailyQuestionListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestion - returns DailyQuestion from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestion(key Key) *DailyQuestion {
	var dailyQuestion DailyQuestion
	if c.tryFindEntry(key, &dailyQuestion) {
		return &dailyQuestion
	}
	return nil
}
