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

// ModifyProjectRequest modify project request
//
// swagger:model ModifyProjectRequest
type ModifyProjectRequest struct {

	// 项目名称 1~64字符，只支持数字、大小写字母、英文下划线“_”及中划线“-”
	// Required: true
	ProjectName *string `json:"projectName"`
}

// Validate validates this modify project request
func (m *ModifyProjectRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModifyProjectRequest) validateProjectName(formats strfmt.Registry) error {

	if err := validate.Required("projectName", "body", m.ProjectName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this modify project request based on context it is used
func (m *ModifyProjectRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModifyProjectRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModifyProjectRequest) UnmarshalBinary(b []byte) error {
	var res ModifyProjectRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}