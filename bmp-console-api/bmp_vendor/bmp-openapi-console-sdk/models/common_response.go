// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonResponse common response
//
// swagger:model CommonResponse
type CommonResponse struct {

	// 操作是否成功 [true/false]
	Success bool `json:"success,omitempty"`
}

// Validate validates this common response
func (m *CommonResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common response based on context it is used
func (m *CommonResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonResponse) UnmarshalBinary(b []byte) error {
	var res CommonResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
