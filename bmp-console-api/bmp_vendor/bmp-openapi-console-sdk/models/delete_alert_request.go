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

// DeleteAlertRequest delete alert request
//
// swagger:model DeleteAlertRequest
type DeleteAlertRequest struct {

	// 告警id
	// Required: true
	AlertID *string `json:"alertId"`
}

// Validate validates this delete alert request
func (m *DeleteAlertRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteAlertRequest) validateAlertID(formats strfmt.Registry) error {

	if err := validate.Required("alertId", "body", m.AlertID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete alert request based on context it is used
func (m *DeleteAlertRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteAlertRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteAlertRequest) UnmarshalBinary(b []byte) error {
	var res DeleteAlertRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}