// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Raid Raid raid
//
// swagger:model Raid
type Raid struct {

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// created time
	CreatedTime int64 `json:"createdTime,omitempty"`

	// deleted time
	DeletedTime int64 `json:"deletedTime,omitempty"`

	// description en
	DescriptionEn string `json:"descriptionEn,omitempty"`

	// description zh
	DescriptionZh string `json:"descriptionZh,omitempty"`

	// is del
	IsDel int8 `json:"isDel,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// raid ID
	RaidID string `json:"raidId,omitempty"`

	// updated by
	UpdatedBy string `json:"updatedBy,omitempty"`

	// updated time
	UpdatedTime int64 `json:"updatedTime,omitempty"`
}

// Validate validates this raid
func (m *Raid) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this raid based on context it is used
func (m *Raid) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Raid) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Raid) UnmarshalBinary(b []byte) error {
	var res Raid
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}