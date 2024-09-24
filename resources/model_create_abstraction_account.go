/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type CreateAbstractionAccount struct {
	Key
	Attributes CreateAbstractionAccountAttributes `json:"attributes"`
}
type CreateAbstractionAccountRequest struct {
	Data     CreateAbstractionAccount `json:"data"`
	Included Included                 `json:"included"`
}

type CreateAbstractionAccountListRequest struct {
	Data     []CreateAbstractionAccount `json:"data"`
	Included Included                   `json:"included"`
	Links    *Links                     `json:"links"`
	Meta     json.RawMessage            `json:"meta,omitempty"`
}

func (r *CreateAbstractionAccountListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateAbstractionAccountListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateAbstractionAccount - returns CreateAbstractionAccount from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateAbstractionAccount(key Key) *CreateAbstractionAccount {
	var createAbstractionAccount CreateAbstractionAccount
	if c.tryFindEntry(key, &createAbstractionAccount) {
		return &createAbstractionAccount
	}
	return nil
}
