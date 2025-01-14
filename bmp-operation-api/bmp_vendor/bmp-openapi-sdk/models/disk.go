// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Disk disk
//
// swagger:model Disk
type Disk struct {

	// adapter ID
	AdapterID int64 `json:"adapterId,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// created time
	CreatedTime int64 `json:"createdTime,omitempty"`

	// deleted time
	DeletedTime int64 `json:"deletedTime,omitempty"`

	// device ID
	DeviceID string `json:"deviceId,omitempty"`

	// device path
	DevicePath string `json:"devicePath,omitempty"`

	// disk Id
	DiskID string `json:"diskId,omitempty"`

	// 磁盘类型：system,data
	DiskType string `json:"diskType,omitempty"`

	// enclosure
	Enclosure string `json:"enclosure,omitempty"`

	// index
	Index int64 `json:"index,omitempty"`

	// is del
	IsDel int8 `json:"isDel,omitempty"`

	// media type
	MediaType string `json:"mediaType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pd type
	PdType string `json:"pdType,omitempty"`

	// 是否选中
	Selected bool `json:"selected,omitempty"`

	// serial number
	SerialNumber string `json:"serialNumber,omitempty"`

	// size
	Size string `json:"size,omitempty"`

	// size unit
	SizeUnit string `json:"sizeUnit,omitempty"`

	// slot
	Slot int64 `json:"slot,omitempty"`

	// control/mvme/panfu
	Types string `json:"types,omitempty"`

	// updated by
	UpdatedBy string `json:"updatedBy,omitempty"`

	// updated time
	UpdatedTime int64 `json:"updatedTime,omitempty"`
}

// Validate validates this disk
func (m *Disk) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this disk based on context it is used
func (m *Disk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Disk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Disk) UnmarshalBinary(b []byte) error {
	var res Disk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
