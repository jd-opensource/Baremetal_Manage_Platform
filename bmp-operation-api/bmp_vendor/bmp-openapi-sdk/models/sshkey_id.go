// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SshkeyID sshkey Id
//
// swagger:model SshkeyId
type SshkeyID struct {

	// sshkey uuid
	SshkeyID string `json:"sshkeyId,omitempty"`
}

// Validate validates this sshkey Id
func (m *SshkeyID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sshkey Id based on context it is used
func (m *SshkeyID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SshkeyID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SshkeyID) UnmarshalBinary(b []byte) error {
	var res SshkeyID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}