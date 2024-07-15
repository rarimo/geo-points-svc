/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type ActivateBalance struct {
	Key
	Attributes ActivateBalanceAttributes `json:"attributes"`
}
type ActivateBalanceRequest struct {
	Data     ActivateBalance `json:"data"`
	Included Included        `json:"included"`
}

type ActivateBalanceListRequest struct {
	Data     []ActivateBalance `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
	Meta     json.RawMessage   `json:"meta,omitempty"`
}

func (r *ActivateBalanceListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ActivateBalanceListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustActivateBalance - returns ActivateBalance from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustActivateBalance(key Key) *ActivateBalance {
	var activateBalance ActivateBalance
	if c.tryFindEntry(key, &activateBalance) {
		return &activateBalance
	}
	return nil
}
