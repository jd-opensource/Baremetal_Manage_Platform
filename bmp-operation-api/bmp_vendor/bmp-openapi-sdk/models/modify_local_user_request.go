// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModifyLocalUserRequest modify local user request
//
// swagger:model ModifyLocalUserRequest
type ModifyLocalUserRequest struct {

	// 用户默认项目uuid
	DefaultProjectID *string `json:"defaultProjectId,omitempty"`

	// 邮箱
	Email *string `json:"email,omitempty"`

	// 语言[en_US, zh_CN]
	Language *string `json:"language,omitempty"`

	// 手机号
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// 国家地区码，如86
	PhonePrefix *string `json:"phonePrefix,omitempty"`

	// 时区 Asia/Shanghai
	Timezone *string `json:"timezone,omitempty"`
}

// Validate validates this modify local user request
func (m *ModifyLocalUserRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this modify local user request based on context it is used
func (m *ModifyLocalUserRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModifyLocalUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModifyLocalUserRequest) UnmarshalBinary(b []byte) error {
	var res ModifyLocalUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}