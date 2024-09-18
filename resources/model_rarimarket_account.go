/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type RarimarketAccount struct {
	Key
	Attributes RarimarketAccountAttributes `json:"attributes"`
}
type RarimarketAccountResponse struct {
	Data     RarimarketAccount `json:"data"`
	Included Included          `json:"included"`
}

type RarimarketAccountListResponse struct {
	Data     []RarimarketAccount `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *RarimarketAccountListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *RarimarketAccountListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustRarimarketAccount - returns RarimarketAccount from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRarimarketAccount(key Key) *RarimarketAccount {
	var rarimarketAccount RarimarketAccount
	if c.tryFindEntry(key, &rarimarketAccount) {
		return &rarimarketAccount
	}
	return nil
}
