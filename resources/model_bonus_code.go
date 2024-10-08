/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type BonusCode struct {
	Key
	Attributes BonusCodeAttributes `json:"attributes"`
}
type BonusCodeRequest struct {
	Data     BonusCode `json:"data"`
	Included Included  `json:"included"`
}

type BonusCodeListRequest struct {
	Data     []BonusCode     `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *BonusCodeListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *BonusCodeListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustBonusCode - returns BonusCode from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBonusCode(key Key) *BonusCode {
	var bonusCode BonusCode
	if c.tryFindEntry(key, &bonusCode) {
		return &bonusCode
	}
	return nil
}
