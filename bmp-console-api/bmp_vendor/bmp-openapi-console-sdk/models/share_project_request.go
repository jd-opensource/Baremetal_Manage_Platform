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

// ShareProjectRequest share project request
//
// swagger:model ShareProjectRequest
type ShareProjectRequest struct {

	// 如果部分分享，instance_id逗号分隔； 如果全部分享，传all
	// Required: true
	InstanceIDs *string `json:"instanceIDs"`

	// from user_id
	// Required: true
	OwnerID *string `json:"ownerID"`

	// to user_id
	// Required: true
	SharerID *string `json:"sharerID"`
}

// Validate validates this share project request
func (m *ShareProjectRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceIDs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharerID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShareProjectRequest) validateInstanceIDs(formats strfmt.Registry) error {

	if err := validate.Required("instanceIDs", "body", m.InstanceIDs); err != nil {
		return err
	}

	return nil
}

func (m *ShareProjectRequest) validateOwnerID(formats strfmt.Registry) error {

	if err := validate.Required("ownerID", "body", m.OwnerID); err != nil {
		return err
	}

	return nil
}

func (m *ShareProjectRequest) validateSharerID(formats strfmt.Registry) error {

	if err := validate.Required("sharerID", "body", m.SharerID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this share project request based on context it is used
func (m *ShareProjectRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ShareProjectRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShareProjectRequest) UnmarshalBinary(b []byte) error {
	var res ShareProjectRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}