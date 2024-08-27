/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestionCreate struct {
	Key
	Attributes DailyQuestionCreateAttributes `json:"attributes"`
}
type DailyQuestionCreateResponse struct {
	Data     DailyQuestionCreate `json:"data"`
	Included Included            `json:"included"`
}

type DailyQuestionCreateListResponse struct {
	Data     []DailyQuestionCreate `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
	Meta     json.RawMessage       `json:"meta,omitempty"`
}

func (r *DailyQuestionCreateListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionCreateListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestionCreate - returns DailyQuestionCreate from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestionCreate(key Key) *DailyQuestionCreate {
	var dailyQuestionCreate DailyQuestionCreate
	if c.tryFindEntry(key, &dailyQuestionCreate) {
		return &dailyQuestionCreate
	}
	return nil
}
