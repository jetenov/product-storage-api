package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	jsonpatch "github.com/evanphx/json-patch/v5"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/validation-service/api/validation"
	psapi "gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product"
)

// Meta attribute keys
const (
	BrandKey   = "brand"
	SellersKey = "sellers"
)

// Product statuses
const (
	CorrectionStatus = "Correction"
	BlockedStatus    = "Blocked"
	CreatedStatus    = "Created"
	FillingStatus    = "Filling"
)

type (
	// Product is a competitor product from ml matcher
	Product struct {
		Attributes            Attributes     `json:"attributes,omitempty"`
		Meta                  Meta           `json:"-"`
		ID                    string         `json:"-"`
		Name                  sql.NullString `json:"name,omitempty"`
		Barcode               sql.NullString `json:"barcode,omitempty"`
		Vat                   sql.NullInt64  `json:"vat,omitempty"`
		Price                 sql.NullInt64  `json:"price,omitempty"`
		Description           sql.NullString `json:"description,omitempty"`
		State                 sql.NullString `json:"-"`
		RootCategoryID        int64          `json:"-"`
		CommercialCategoryID  int64          `json:"-"`
		DescriptionCategoryID int64          `json:"-"`
		Images                Images         `json:"images,omitempty"`
		Fulfillment           sql.NullInt64  `json:"-"`
		MinFulfillment        sql.NullInt64  `json:"-"`
		BlockReason           sql.NullInt64  `json:"block_reason,omitempty"`
		Depth                 sql.NullInt64  `json:"depth,omitempty"`
		Weight                sql.NullInt64  `json:"weight,omitempty"`
		Height                sql.NullInt64  `json:"height,omitempty"`
		Width                 sql.NullInt64  `json:"width,omitempty"`
		CreatedAt             time.Time      `json:"-"`
		UpdatedAt             time.Time      `json:"-"`
		CanonicalIDs          []string       `json:"-"`
	}

	// ProductImages ...
	ProductImages struct {
		ID     string
		Images Images
	}
)

// MarshalJSON ...
func (p Product) MarshalJSON() ([]byte, error) {
	type Alias Product
	return json.Marshal(&struct {
		Name        string `json:"name,omitempty"`
		Barcode     string `json:"barcode,omitempty"`
		Vat         int64  `json:"vat,omitempty"`
		Price       int64  `json:"price,omitempty"`
		Description string `json:"description,omitempty"`
		BlockReason int64  `json:"block_reason,omitempty"`
		Depth       int64  `json:"depth,omitempty"`
		Weight      int64  `json:"weight,omitempty"`
		Height      int64  `json:"height,omitempty"`
		Width       int64  `json:"width,omitempty"`
		Alias
	}{
		Name:        p.Name.String,
		Barcode:     p.Barcode.String,
		Vat:         p.Vat.Int64,
		Price:       p.Price.Int64,
		Description: p.Description.String,
		BlockReason: p.BlockReason.Int64,
		Depth:       p.Depth.Int64,
		Weight:      p.Weight.Int64,
		Height:      p.Height.Int64,
		Width:       p.Width.Int64,
		Alias:       (Alias)(p),
	})
}

// UnmarshalJSON ...
func (p *Product) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, p)
}

// Value ...
func (p Product) Value() (driver.Value, error) {
	v, err := json.Marshal(p)

	return string(v), err
}

// Scan ...
func (p *Product) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &p)
}

// AttributesToValidation ...
func (p *Product) AttributesToValidation() []*validation.Product_Attribute {
	attributes := make([]*validation.Product_Attribute, 0, len(p.Attributes))
	for _, attribute := range p.Attributes {
		attr := validation.Product_Attribute{Id: attribute.ID, ComplexId: attribute.ComplexID}
		for _, v := range attribute.AttributeValue {
			attr.Values = append(attr.Values, &validation.Product_Attribute_Value{Id: v.ID, Value: v.Value})
		}
		attributes = append(attributes, &attr)
	}

	return attributes
}

// ImagesToValidation ...
func (p *Product) ImagesToValidation() []*validation.Product_Image {
	images := make([]*validation.Product_Image, 0, len(p.Images))
	for _, image := range p.Images {
		images = append(images, &validation.Product_Image{Url: image.URL, IsMain: image.IsMain})
	}

	return images
}

// AttributesToProto ...
func (p *Product) AttributesToProto() []*psapi.Attribute {
	attributes := make([]*psapi.Attribute, 0, len(p.Attributes))
	for _, attribute := range p.Attributes {
		attr := psapi.Attribute{Id: attribute.ID, ComplexId: attribute.ComplexID, Name: attribute.Name}
		for _, v := range attribute.AttributeValue {
			attr.Values = append(attr.Values, &psapi.Attribute_Value{Id: v.ID, Value: v.Value})
		}
		attributes = append(attributes, &attr)
	}

	return attributes
}

// ImagesToProto ...
func (p *Product) ImagesToProto() []*psapi.Image {
	images := make([]*psapi.Image, 0, len(p.Images))
	for _, image := range p.Images {
		images = append(images, &psapi.Image{Url: image.URL, CephUrl: image.CephURL, IsMain: image.IsMain})
	}

	return images
}

// MergePatch ...
func (p *Product) MergePatch(from *Product) ([]byte, error) {
	original, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}

	target, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return jsonpatch.CreateMergePatch(original, target)
}
