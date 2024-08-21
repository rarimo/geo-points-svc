/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestions struct {
	Key
	Attributes DailyQuestionsAttributes `json:"attributes"`
}
type DailyQuestionsResponse struct {
	Data     DailyQuestions `json:"data"`
	Included Included       `json:"included"`
}

type DailyQuestionsListResponse struct {
	Data     []DailyQuestions `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
	Meta     json.RawMessage  `json:"meta,omitempty"`
}

func (r *DailyQuestionsListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionsListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestions - returns DailyQuestions from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestions(key Key) *DailyQuestions {
	var dailyQuestions DailyQuestions
	if c.tryFindEntry(key, &dailyQuestions) {
		return &dailyQuestions
	}
	return nil
}
