// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Sshkey sshkey
//
// swagger:model Sshkey
type Sshkey struct {

	// 创建者
	CreatedBy string `json:"createdBy,omitempty"`

	// 创建时间
	CreatedTime string `json:"createdTime,omitempty"`

	// 公钥指纹
	FingerPrint string `json:"fingerPrint,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// 公钥，格式：ssh-rsa AAA
	Key string `json:"key,omitempty"`

	// 秘钥名称
	Name string `json:"name,omitempty"`

	// 秘钥uuid
	SshkeyID string `json:"sshkeyId,omitempty"`

	// 更新者
	UpdatedBy string `json:"updatedBy,omitempty"`

	// 更新时间
	UpdatedTime string `json:"updatedTime,omitempty"`

	// 所属用户uuid
	UserID string `json:"userId,omitempty"`
}

// Validate validates this sshkey
func (m *Sshkey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sshkey based on context it is used
func (m *Sshkey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Sshkey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Sshkey) UnmarshalBinary(b []byte) error {
	var res Sshkey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}