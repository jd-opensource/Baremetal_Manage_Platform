// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeRaID volume ra ID
//
// swagger:model VolumeRaID
type VolumeRaID struct {

	// raid can
	RaidCan string `json:"raidCan,omitempty"`

	// raid uuid
	RaidID string `json:"raidId,omitempty"`

	// 卷 uuid
	VolumeID string `json:"volumeId,omitempty"`

	// system|data
	VolumeType string `json:"volumeType,omitempty"`
}

// Validate validates this volume ra ID
func (m *VolumeRaID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume ra ID based on context it is used
func (m *VolumeRaID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeRaID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeRaID) UnmarshalBinary(b []byte) error {
	var res VolumeRaID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
