package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	desc "gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product"
)

type (
	// Meta ...
	Meta map[string]interface{}

	// Seller ...
	Seller struct {
		OGRN  string `json:"ogrn,omitempty"`
		Title string `json:"title,omitempty"`
	}

	// Brand ...
	Brand struct {
		Title string `json:"title,omitempty"`
	}
)

// Value ...
func (m Meta) Value() (driver.Value, error) {
	if len(m) == 0 {
		return nil, nil
	}

	v, err := json.Marshal(m)

	return string(v), err
}

// Scan ...
func (m *Meta) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	type tmp struct {
		Brand   *Brand    `json:"brand"`
		Sellers []*Seller `json:"sellers"`
	}
	var base tmp

	if err := json.Unmarshal(b, &base); err != nil {
		return err
	}

	(*m)[BrandKey] = base.Brand
	(*m)[SellersKey] = base.Sellers

	return nil
}

// Fill ...
func (m Meta) Fill(meta map[string]*anypb.Any) error {
	for key, any := range meta {
		switch key {
		case BrandKey:
			var b desc.Brand
			if err := anypb.UnmarshalTo(any, &b, proto.UnmarshalOptions{}); err != nil {
				return err
			}

			m[BrandKey] = &Brand{Title: b.Title}
		case SellersKey:
			var s desc.Sellers
			if err := anypb.UnmarshalTo(any, &s, proto.UnmarshalOptions{}); err != nil {
				return err
			}

			sellers := make([]*Seller, 0, len(s.Sellers))
			for _, ls := range s.Sellers {
				sellers = append(sellers, &Seller{Title: ls.Title, OGRN: ls.Ogrn})
			}

			m[SellersKey] = sellers
		}
	}

	return nil
}

// Brand ...
func (m Meta) Brand() (*Brand, error) {
	data, ok := m[BrandKey]
	if !ok {
		return nil, errors.New("meta brand key is not found")
	}

	brand, ok := data.(*Brand)
	if !ok {
		return nil, errors.New("meta brand assertion failed")
	}

	return brand, nil
}

// Sellers ...
func (m Meta) Sellers() ([]*Seller, error) {
	data, ok := m[SellersKey]
	if !ok {
		return nil, errors.New("meta sellers key is not found")
	}

	sellers, ok := data.([]*Seller)
	if !ok {
		return nil, errors.New("meta sellers assertion failed")
	}

	return sellers, nil
}

// ToProto ...
func (m Meta) ToProto() (map[string]*anypb.Any, error) {
	res := make(map[string]*anypb.Any)

	for key := range m {
		switch key {
		case BrandKey:
			brand, err := m.Brand()
			if err != nil {
				return res, err
			}

			any, err := anypb.New(&desc.Brand{Title: brand.Title})
			if err != nil {
				return res, err
			}

			res[BrandKey] = any
		case SellersKey:
			data, err := m.Sellers()
			if err != nil {
				return res, err
			}

			sellers := make([]*desc.Seller, 0, len(data))
			for _, s := range data {
				sellers = append(sellers, &desc.Seller{Title: s.Title, Ogrn: s.OGRN})
			}

			any, err := anypb.New(&desc.Sellers{
				Sellers: sellers,
			})
			if err != nil {
				return res, err
			}

			res[SellersKey] = any
		}
	}

	return res, nil
}
