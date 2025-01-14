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

// MessageList message list
//
// swagger:model MessageList
type MessageList struct {

	// messages
	Messages []*Message `json:"messages"`

	// page number
	PageNumber int64 `json:"pageNumber,omitempty"`

	// page size
	PageSize int64 `json:"pageSize,omitempty"`

	// total count
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Validate validates this message list
func (m *MessageList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageList) validateMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	for i := 0; i < len(m.Messages); i++ {
		if swag.IsZero(m.Messages[i]) { // not required
			continue
		}

		if m.Messages[i] != nil {
			if err := m.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this message list based on the context it is used
func (m *MessageList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageList) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Messages); i++ {

		if m.Messages[i] != nil {
			if err := m.Messages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessageList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageList) UnmarshalBinary(b []byte) error {
	var res MessageList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
