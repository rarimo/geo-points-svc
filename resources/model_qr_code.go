/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type QrCode struct {
	Key
	Attributes QrCodeAttributes `json:"attributes"`
}
type QrCodeRequest struct {
	Data     QrCode   `json:"data"`
	Included Included `json:"included"`
}

type QrCodeListRequest struct {
	Data     []QrCode        `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *QrCodeListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *QrCodeListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustQrCode - returns QrCode from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustQrCode(key Key) *QrCode {
	var qRCode QrCode
	if c.tryFindEntry(key, &qRCode) {
		return &qRCode
	}
	return nil
}
