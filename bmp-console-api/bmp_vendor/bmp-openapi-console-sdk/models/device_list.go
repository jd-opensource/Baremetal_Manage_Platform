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

// DeviceList device list
//
// swagger:model DeviceList
type DeviceList struct {

	// 设备列表
	Devices []*Device `json:"devices"`

	// page number
	PageNumber int64 `json:"pageNumber,omitempty"`

	// page size
	PageSize int64 `json:"pageSize,omitempty"`

	// total count
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Validate validates this device list
func (m *DeviceList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceList) validateDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.Devices) { // not required
		return nil
	}

	for i := 0; i < len(m.Devices); i++ {
		if swag.IsZero(m.Devices[i]) { // not required
			continue
		}

		if m.Devices[i] != nil {
			if err := m.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this device list based on the context it is used
func (m *DeviceList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceList) contextValidateDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Devices); i++ {

		if m.Devices[i] != nil {
			if err := m.Devices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceList) UnmarshalBinary(b []byte) error {
	var res DeviceList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}