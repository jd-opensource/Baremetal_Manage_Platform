// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteInstancesRequest 批量删除实例
//
// swagger:model DeleteInstancesRequest
type DeleteInstancesRequest struct {

	// 实例id列表
	InstanceIDs []string `json:"instance_ids"`
}

// Validate validates this delete instances request
func (m *DeleteInstancesRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete instances request based on context it is used
func (m *DeleteInstancesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteInstancesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteInstancesRequest) UnmarshalBinary(b []byte) error {
	var res DeleteInstancesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}