// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MessageWithNextPrev message with next prev
//
// swagger:model MessageWithNextPrev
type MessageWithNextPrev struct {

	// next message Id
	NextMessageID string `json:"nextMessageId,omitempty"`

	// prev message Id
	PrevMessageID string `json:"prevMessageId,omitempty"`

	// message
	Message *Message `json:"message,omitempty"`
}

// Validate validates this message with next prev
func (m *MessageWithNextPrev) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageWithNextPrev) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if m.Message != nil {
		if err := m.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this message with next prev based on the context it is used
func (m *MessageWithNextPrev) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageWithNextPrev) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.Message != nil {
		if err := m.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessageWithNextPrev) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageWithNextPrev) UnmarshalBinary(b []byte) error {
	var res MessageWithNextPrev
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}