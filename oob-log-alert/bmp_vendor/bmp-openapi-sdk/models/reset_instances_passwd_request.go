// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResetInstancesPasswdRequest 批量重置实例密码
//
// swagger:model ResetInstancesPasswdRequest
type ResetInstancesPasswdRequest struct {

	// 实例id列表
	InstanceIDs []string `json:"instance_ids"`

	// 新密码
	Password string `json:"password,omitempty"`
}

// Validate validates this reset instances passwd request
func (m *ResetInstancesPasswdRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this reset instances passwd request based on context it is used
func (m *ResetInstancesPasswdRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResetInstancesPasswdRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResetInstancesPasswdRequest) UnmarshalBinary(b []byte) error {
	var res ResetInstancesPasswdRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
