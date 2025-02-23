// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MonitorDataItem monitor data item
//
// swagger:model MonitorDataItem
type MonitorDataItem struct {

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this monitor data item
func (m *MonitorDataItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this monitor data item based on context it is used
func (m *MonitorDataItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MonitorDataItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MonitorDataItem) UnmarshalBinary(b []byte) error {
	var res MonitorDataItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
