/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type AbstractionAccount struct {
	Key
	Attributes AbstractionAccountAttributes `json:"attributes"`
}
type AbstractionAccountResponse struct {
	Data     AbstractionAccount `json:"data"`
	Included Included           `json:"included"`
}

type AbstractionAccountListResponse struct {
	Data     []AbstractionAccount `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
	Meta     json.RawMessage      `json:"meta,omitempty"`
}

func (r *AbstractionAccountListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *AbstractionAccountListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustAbstractionAccount - returns AbstractionAccount from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAbstractionAccount(key Key) *AbstractionAccount {
	var abstractionAccount AbstractionAccount
	if c.tryFindEntry(key, &abstractionAccount) {
		return &abstractionAccount
	}
	return nil
}
