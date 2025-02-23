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

// MoveInstancesRequest move instances request
//
// swagger:model MoveInstancesRequest
type MoveInstancesRequest struct {

	// instanceIDs
	// Required: true
	InstanceIDs *string `json:"instanceIDs"`

	// to user_id
	// Required: true
	MoverID *string `json:"moverID"`

	// to project_id
	// Required: true
	MoverProjectID *string `json:"moverProjectID"`

	// from user_id
	// Required: true
	OwnerID *string `json:"ownerID"`

	// from project_id
	// Required: true
	OwnerProjectID *string `json:"ownerProjectID"`
}

// Validate validates this move instances request
func (m *MoveInstancesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceIDs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoverID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoverProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveInstancesRequest) validateInstanceIDs(formats strfmt.Registry) error {

	if err := validate.Required("instanceIDs", "body", m.InstanceIDs); err != nil {
		return err
	}

	return nil
}

func (m *MoveInstancesRequest) validateMoverID(formats strfmt.Registry) error {

	if err := validate.Required("moverID", "body", m.MoverID); err != nil {
		return err
	}

	return nil
}

func (m *MoveInstancesRequest) validateMoverProjectID(formats strfmt.Registry) error {

	if err := validate.Required("moverProjectID", "body", m.MoverProjectID); err != nil {
		return err
	}

	return nil
}

func (m *MoveInstancesRequest) validateOwnerID(formats strfmt.Registry) error {

	if err := validate.Required("ownerID", "body", m.OwnerID); err != nil {
		return err
	}

	return nil
}

func (m *MoveInstancesRequest) validateOwnerProjectID(formats strfmt.Registry) error {

	if err := validate.Required("ownerProjectID", "body", m.OwnerProjectID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this move instances request based on context it is used
func (m *MoveInstancesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MoveInstancesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveInstancesRequest) UnmarshalBinary(b []byte) error {
	var res MoveInstancesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
