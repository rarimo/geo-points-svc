/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestionAnswer struct {
	Key
	Attributes DailyQuestionAnswerAttributes `json:"attributes"`
}
type DailyQuestionAnswerResponse struct {
	Data     DailyQuestionAnswer `json:"data"`
	Included Included            `json:"included"`
}

type DailyQuestionAnswerListResponse struct {
	Data     []DailyQuestionAnswer `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
	Meta     json.RawMessage       `json:"meta,omitempty"`
}

func (r *DailyQuestionAnswerListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionAnswerListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestionAnswer - returns DailyQuestionAnswer from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestionAnswer(key Key) *DailyQuestionAnswer {
	var dailyQuestionAnswer DailyQuestionAnswer
	if c.tryFindEntry(key, &dailyQuestionAnswer) {
		return &dailyQuestionAnswer
	}
	return nil
}
