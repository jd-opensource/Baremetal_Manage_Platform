// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OsID os Id
//
// swagger:model OsId
type OsID struct {

	// 操作系统uuid
	OsID string `json:"osId,omitempty"`
}

// Validate validates this os Id
func (m *OsID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this os Id based on context it is used
func (m *OsID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OsID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsID) UnmarshalBinary(b []byte) error {
	var res OsID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
