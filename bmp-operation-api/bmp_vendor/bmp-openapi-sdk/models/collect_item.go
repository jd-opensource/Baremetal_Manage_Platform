// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CollectItem collect item
//
// swagger:model CollectItem
type CollectItem struct {

	// 可选参数。 传参:True, False 。True表示传入的raid_driver值将覆盖已适配机器的raid_driver
	AllowOverride *bool `json:"allowOverride,omitempty"`

	// sn
	// Required: true
	Sn *string `json:"sn"`
}

// Validate validates this collect item
func (m *CollectItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectItem) validateSn(formats strfmt.Registry) error {

	if err := validate.Required("sn", "body", m.Sn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this collect item based on context it is used
func (m *CollectItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CollectItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectItem) UnmarshalBinary(b []byte) error {
	var res CollectItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
