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

// CreateImageRequest create image request
//
// swagger:model CreateImageRequest
type CreateImageRequest struct {

	// 体系架构
	// Required: true
	Architecture *string `json:"architecture"`

	// 引导模式【UEFI Legacy/BIOS】
	// Required: true
	BootMode *string `json:"boot_mode"`

	// 数据分区信息
	DataPartition string `json:"dataPartition,omitempty"`

	// 描述
	Description string `json:"description,omitempty"`

	// 镜像文件名称
	// Required: true
	Filename *string `json:"filename"`

	// 镜像文件类型
	// Required: true
	Format *string `json:"format"`

	// 文件hash值
	// Required: true
	Hash *string `json:"hash"`

	// 镜像名称
	// Required: true
	ImageName *string `json:"imageName"`

	// 操作系统平台
	// Required: true
	OsType *string `json:"osType"`

	// 镜像类型，预置，自定义
	Source string `json:"source,omitempty"`

	// 系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：50GiB，swap）
	SystemPartition string `json:"systemPartition,omitempty"`

	// 镜像上传地址
	// Required: true
	URL *string `json:"url"`

	// 版本号
	// Required: true
	Version *string `json:"Version"`
}

// Validate validates this create image request
func (m *CreateImageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchitecture(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateImageRequest) validateArchitecture(formats strfmt.Registry) error {

	if err := validate.Required("architecture", "body", m.Architecture); err != nil {
		return err
	}

	return nil
}

func (m *CreateImageRequest) validateBootMode(formats strfmt.Registry) error {

	if err := validate.Required("boot_mode", "body", m.BootMode); err != nil {
		return err
	}

	return nil
}

func (m *CreateImageRequest) validateFilename(formats strfmt.Registry) error {

	if err := validate.Required("filename", "body", m.Filename); err != nil {
		return err
	}

	return nil
}

func (m *CreateImageRequest) validateFormat(formats strfmt.Registry) error {

	if err := validate.Required("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

func (m *CreateImageRequest) validateHash(formats strfmt.Registry) error {

	if err := validate.Required("hash", "body", m.Hash); err != nil {
		return err
	}

	return nil
}

func (m *CreateImageRequest) validateImageName(formats strfmt.Registry) error {

	if err := validate.Required("imageName", "body", m.ImageName); err != nil {
		return err
	}

	return nil
}

func (m *CreateImageRequest) validateOsType(formats strfmt.Registry) error {

	if err := validate.Required("osType", "body", m.OsType); err != nil {
		return err
	}

	return nil
}

func (m *CreateImageRequest) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *CreateImageRequest) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("Version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create image request based on context it is used
func (m *CreateImageRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateImageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateImageRequest) UnmarshalBinary(b []byte) error {
	var res CreateImageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}