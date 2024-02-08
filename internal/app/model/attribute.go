package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type (
	// Attribute ...
	Attribute struct {
		ID             int64   `json:"id,omitempty"`
		ComplexID      int64   `json:"complex_id,omitempty"`
		Name           string  `json:"name,omitempty"`
		AttributeValue []Value `json:"value,omitempty"`
	}

	// Attributes ...
	Attributes []*Attribute

	// Value ...
	Value struct {
		ID    int64  `json:"id,omitempty"`
		Value string `json:"value,omitempty"`
	}
)

// Value ...
func (a Attributes) Value() (driver.Value, error) {
	if len(a) == 0 {
		return nil, nil
	}

	v, err := json.Marshal(a)

	return string(v), err
}

// Scan ...
func (a *Attributes) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
