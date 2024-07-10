package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gitlab.com/distributed_lab/kit/pgdb"
)

// LocalizationMap maps 2-letter locales to Localized data and implements
// interfaces to work with DB
type LocalizationMap map[string]Localized

type Localized struct {
	Title            string `fig:"title,required" json:"title"`
	Description      string `fig:"description,required" json:"description"`
	ShortDescription string `fig:"short_description,required" json:"short_description"`
}

func (l *LocalizationMap) Value() (driver.Value, error) {
	if l == nil || len(*l) == 0 {
		return nil, nil
	}
	return pgdb.JSONValue(l)
}

func (l *LocalizationMap) Scan(src interface{}) error {
	var data []byte
	switch rawData := src.(type) {
	case []byte:
		data = rawData
	case string:
		data = []byte(rawData)
	case nil:
		return nil
	default:
		return fmt.Errorf("unexpected type for jsonb: %T", src)
	}

	err := json.Unmarshal(data, l)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}

	return nil
}
