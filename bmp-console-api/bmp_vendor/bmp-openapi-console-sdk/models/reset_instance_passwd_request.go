// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResetInstancePasswdRequest 重置实例密码
//
// swagger:model ResetInstancePasswdRequest
type ResetInstancePasswdRequest struct {

	// 新密码
	Password string `json:"password,omitempty"`
}

// Validate validates this reset instance passwd request
func (m *ResetInstancePasswdRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this reset instance passwd request based on context it is used
func (m *ResetInstancePasswdRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResetInstancePasswdRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResetInstancePasswdRequest) UnmarshalBinary(b []byte) error {
	var res ResetInstancePasswdRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}