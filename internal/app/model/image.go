package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type (
	// Image ...
	Image struct {
		URL     string `json:"url,omitempty"`
		CephURL string `json:"ceph_url,omitempty"`
		IsMain  bool   `json:"is_main,omitempty"`
	}

	// Images ..
	Images []*Image
)

// Value ...
func (i Images) Value() (driver.Value, error) {
	if len(i) == 0 {
		return nil, nil
	}

	v, err := json.Marshal(i)

	return string(v), err
}

// Scan ...
func (i *Images) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &i)
}
