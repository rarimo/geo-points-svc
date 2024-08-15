/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources
import "gitlab.com/tokend/go/xdr"

type Key struct {
	ID string `json:"id"`
	Type ResourceType `json:"type"`
}

func NewKeyInt64(id int64, resourceType ResourceType) Key {
    return Key{
        ID:   strconv.FormatInt(id, 10),
        Type: resourceType,
    }
}

func (r* Key) GetKey() Key {
    return *r
}

func (r Key) GetKeyP() *Key {
    return &r
}


func (r Key) AsRelation() *Relation {
    return &Relation {
        Data: r.GetKeyP(),
    }
}

