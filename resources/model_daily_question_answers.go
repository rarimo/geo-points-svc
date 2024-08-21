/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestionAnswers struct {
	Key
	Attributes DailyQuestionAnswersAttributes `json:"attributes"`
}
type DailyQuestionAnswersResponse struct {
	Data     DailyQuestionAnswers `json:"data"`
	Included Included             `json:"included"`
}

type DailyQuestionAnswersListResponse struct {
	Data     []DailyQuestionAnswers `json:"data"`
	Included Included               `json:"included"`
	Links    *Links                 `json:"links"`
	Meta     json.RawMessage        `json:"meta,omitempty"`
}

func (r *DailyQuestionAnswersListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionAnswersListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestionAnswers - returns DailyQuestionAnswers from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestionAnswers(key Key) *DailyQuestionAnswers {
	var dailyQuestionAnswers DailyQuestionAnswers
	if c.tryFindEntry(key, &dailyQuestionAnswers) {
		return &dailyQuestionAnswers
	}
	return nil
}
