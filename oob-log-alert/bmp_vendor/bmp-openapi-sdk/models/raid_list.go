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
)

// RaidList raid list
//
// swagger:model RaidList
type RaidList struct {

	// raid实体列表
	Raids []*Raid `json:"raids"`
}

// Validate validates this raid list
func (m *RaidList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRaids(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RaidList) validateRaids(formats strfmt.Registry) error {
	if swag.IsZero(m.Raids) { // not required
		return nil
	}

	for i := 0; i < len(m.Raids); i++ {
		if swag.IsZero(m.Raids[i]) { // not required
			continue
		}

		if m.Raids[i] != nil {
			if err := m.Raids[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("raids" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("raids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this raid list based on the context it is used
func (m *RaidList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRaids(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RaidList) contextValidateRaids(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Raids); i++ {

		if m.Raids[i] != nil {
			if err := m.Raids[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("raids" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("raids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RaidList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RaidList) UnmarshalBinary(b []byte) error {
	var res RaidList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}