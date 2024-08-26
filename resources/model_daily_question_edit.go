/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestionEdit struct {
	Key
	Attributes DailyQuestionEditAttributes `json:"attributes"`
}
type DailyQuestionEditResponse struct {
	Data     DailyQuestionEdit `json:"data"`
	Included Included          `json:"included"`
}

type DailyQuestionEditListResponse struct {
	Data     []DailyQuestionEdit `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *DailyQuestionEditListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionEditListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestionEdit - returns DailyQuestionEdit from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestionEdit(key Key) *DailyQuestionEdit {
	var dailyQuestionEdit DailyQuestionEdit
	if c.tryFindEntry(key, &dailyQuestionEdit) {
		return &dailyQuestionEdit
	}
	return nil
}
