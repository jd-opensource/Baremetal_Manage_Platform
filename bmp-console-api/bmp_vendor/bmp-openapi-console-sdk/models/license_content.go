// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LicenseContent license content
//
// swagger:model LicenseContent
type LicenseContent struct {

	// end time
	EndTime int64 `json:"end_time,omitempty"`

	// hard ware
	HardWare string `json:"hardwareinfo,omitempty"`

	// license name
	LicenseName string `json:"license_name,omitempty"`

	// license type
	LicenseType string `json:"license_type,omitempty"`

	// models
	Models []string `json:"models"`

	// nodes num
	NodesNum int64 `json:"nodes_num,omitempty"`

	// start time
	StartTime int64 `json:"start_time,omitempty"`

	// used nodes num
	UsedNodesNum int64 `json:"used_nodes_num,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this license content
func (m *LicenseContent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this license content based on context it is used
func (m *LicenseContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LicenseContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseContent) UnmarshalBinary(b []byte) error {
	var res LicenseContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
