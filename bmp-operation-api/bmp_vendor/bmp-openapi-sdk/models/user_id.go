// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserID user Id
//
// swagger:model UserId
type UserID struct {

	// 用户uuid
	UserID string `json:"userId,omitempty"`
}

// Validate validates this user Id
func (m *UserID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user Id based on context it is used
func (m *UserID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserID) UnmarshalBinary(b []byte) error {
	var res UserID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}