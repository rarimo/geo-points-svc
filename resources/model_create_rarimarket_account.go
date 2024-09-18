/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type CreateRarimarketAccount struct {
	Key
	Attributes CreateRarimarketAccountAttributes `json:"attributes"`
}
type CreateRarimarketAccountRequest struct {
	Data     CreateRarimarketAccount `json:"data"`
	Included Included                `json:"included"`
}

type CreateRarimarketAccountListRequest struct {
	Data     []CreateRarimarketAccount `json:"data"`
	Included Included                  `json:"included"`
	Links    *Links                    `json:"links"`
	Meta     json.RawMessage           `json:"meta,omitempty"`
}

func (r *CreateRarimarketAccountListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateRarimarketAccountListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateRarimarketAccount - returns CreateRarimarketAccount from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateRarimarketAccount(key Key) *CreateRarimarketAccount {
	var createRarimarketAccount CreateRarimarketAccount
	if c.tryFindEntry(key, &createRarimarketAccount) {
		return &createRarimarketAccount
	}
	return nil
}
