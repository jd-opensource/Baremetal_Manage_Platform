// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CollectDeviceInfoRequest collect device info request
//
// swagger:model CollectDeviceInfoRequest
type CollectDeviceInfoRequest struct {

	// 待采集项
	// Required: true
	Collects []*CollectItem `json:"collects"`
}

// Validate validates this collect device info request
func (m *CollectDeviceInfoRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectDeviceInfoRequest) validateCollects(formats strfmt.Registry) error {

	if err := validate.Required("collects", "body", m.Collects); err != nil {
		return err
	}

	for i := 0; i < len(m.Collects); i++ {
		if swag.IsZero(m.Collects[i]) { // not required
			continue
		}

		if m.Collects[i] != nil {
			if err := m.Collects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this collect device info request based on the context it is used
func (m *CollectDeviceInfoRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectDeviceInfoRequest) contextValidateCollects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Collects); i++ {

		if m.Collects[i] != nil {
			if err := m.Collects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CollectDeviceInfoRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectDeviceInfoRequest) UnmarshalBinary(b []byte) error {
	var res CollectDeviceInfoRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}