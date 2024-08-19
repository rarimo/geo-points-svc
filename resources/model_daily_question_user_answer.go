/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestionUserAnswer struct {
	Key
	Attributes DailyQuestionUserAnswerAttributes `json:"attributes"`
}
type DailyQuestionUserAnswerResponse struct {
	Data     DailyQuestionUserAnswer `json:"data"`
	Included Included                `json:"included"`
}

type DailyQuestionUserAnswerListResponse struct {
	Data     []DailyQuestionUserAnswer `json:"data"`
	Included Included                  `json:"included"`
	Links    *Links                    `json:"links"`
	Meta     json.RawMessage           `json:"meta,omitempty"`
}

func (r *DailyQuestionUserAnswerListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionUserAnswerListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestionUserAnswer - returns DailyQuestionUserAnswer from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestionUserAnswer(key Key) *DailyQuestionUserAnswer {
	var dailyQuestionUserAnswer DailyQuestionUserAnswer
	if c.tryFindEntry(key, &dailyQuestionUserAnswer) {
		return &dailyQuestionUserAnswer
	}
	return nil
}
