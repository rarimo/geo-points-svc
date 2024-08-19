/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DailyQuestionUserAccess struct {
	Key
	Attributes DailyQuestionUserAccessAttributes `json:"attributes"`
}
type DailyQuestionUserAccessResponse struct {
	Data     DailyQuestionUserAccess `json:"data"`
	Included Included                `json:"included"`
}

type DailyQuestionUserAccessListResponse struct {
	Data     []DailyQuestionUserAccess `json:"data"`
	Included Included                  `json:"included"`
	Links    *Links                    `json:"links"`
	Meta     json.RawMessage           `json:"meta,omitempty"`
}

func (r *DailyQuestionUserAccessListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DailyQuestionUserAccessListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDailyQuestionUserAccess - returns DailyQuestionUserAccess from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDailyQuestionUserAccess(key Key) *DailyQuestionUserAccess {
	var dailyQuestionUserAccess DailyQuestionUserAccess
	if c.tryFindEntry(key, &dailyQuestionUserAccess) {
		return &dailyQuestionUserAccess
	}
	return nil
}
