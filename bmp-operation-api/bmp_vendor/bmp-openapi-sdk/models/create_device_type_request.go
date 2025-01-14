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
	"github.com/go-openapi/validate"
)

// CreateDeviceTypeRequest create device type request
//
// swagger:model CreateDeviceTypeRequest
type CreateDeviceTypeRequest struct {

	// 体系架构，1-64
	// Required: true
	Architecture *string `json:"architecture"`

	// boot模式【UEFI Legacy/BIOS】支持多选，逗号分隔
	BootMode string `json:"boot_mode,omitempty"`

	// cpu数量1,10000
	// Required: true
	CPUAmount *int64 `json:"cpuAmount"`

	// 单个cpu内核数1,10000
	// Required: true
	CPUCores *int64 `json:"cpuCores"`

	// cpu频率(G)1-64
	// Required: true
	CPUFrequency *string `json:"cpuFrequency"`

	// cpu厂商1-64
	// Required: true
	CPUManufacturer *string `json:"cpuManufacturer"`

	// cpu处理器型号1-64
	// Required: true
	CPUModel *string `json:"cpuModel"`

	// CPU 预置规格
	// Required: true
	CPUSpec *string `json:"cpuSpec"`

	// 描述0-256
	Description string `json:"description,omitempty"`

	// 机型类型，如computer storage gpu other
	// Required: true
	DeviceSeries *string `json:"deviceSeries"`

	// 机型规格1-64
	// Required: true
	DeviceType *string `json:"deviceType"`

	// gpu数量0-1000
	GpuAmount int64 `json:"gpuAmount,omitempty"`

	// gpu品牌0-64
	GpuManufacturer string `json:"gpuManufacturer,omitempty"`

	// gpu型号0-64
	GpuModel string `json:"gpuModel,omitempty"`

	// 【高度（U）】：显示机型高度 2,4
	// Required: true
	Height *int64 `json:"height"`

	// 机房id
	// Required: true
	IDcID *string `json:"idcId"`

	// 【网口模式】【网络设置】: bond,dual,single
	// Required: true
	InterfaceMode *string `json:"interfaceMode"`

	// 阵列卡是否需要配置 1 需要配置 2无需配置
	IsNeedRaid string `json:"isNeedRaid,omitempty"`

	// 内存数量1-10000
	// Required: true
	MemAmount *int64 `json:"memAmount"`

	// 内存主频（MHz)1-10000
	// Required: true
	MemFrequency *int64 `json:"memFrequency"`

	// 单个内存大小1,10000(GB)
	// Required: true
	MemSize *int64 `json:"memSize"`

	// 内存 预置规格
	// Required: true
	MemSpec *string `json:"memSpec"`

	// 内存接口（DDR DDR2 DDR3 DDR4 DDR5）
	// Required: true
	MemType *string `json:"memType"`

	// 机型名称1-64
	// Required: true
	Name *string `json:"name"`

	// 网卡数量1-10000
	// Required: true
	NicAmount *int64 `json:"nicAmount"`

	// 网卡传输速率(GE) 1-10000
	// Required: true
	NicRate *int64 `json:"nicRate"`

	// 卷管理
	// Required: true
	Volumes []*VolumeItem `json:"volumes"`
}

// Validate validates this create device type request
func (m *CreateDeviceTypeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchitecture(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUManufacturer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceSeries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDcID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaceMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNicAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNicRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDeviceTypeRequest) validateArchitecture(formats strfmt.Registry) error {

	if err := validate.Required("architecture", "body", m.Architecture); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateCPUAmount(formats strfmt.Registry) error {

	if err := validate.Required("cpuAmount", "body", m.CPUAmount); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateCPUCores(formats strfmt.Registry) error {

	if err := validate.Required("cpuCores", "body", m.CPUCores); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateCPUFrequency(formats strfmt.Registry) error {

	if err := validate.Required("cpuFrequency", "body", m.CPUFrequency); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateCPUManufacturer(formats strfmt.Registry) error {

	if err := validate.Required("cpuManufacturer", "body", m.CPUManufacturer); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateCPUModel(formats strfmt.Registry) error {

	if err := validate.Required("cpuModel", "body", m.CPUModel); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateCPUSpec(formats strfmt.Registry) error {

	if err := validate.Required("cpuSpec", "body", m.CPUSpec); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateDeviceSeries(formats strfmt.Registry) error {

	if err := validate.Required("deviceSeries", "body", m.DeviceSeries); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("deviceType", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateIDcID(formats strfmt.Registry) error {

	if err := validate.Required("idcId", "body", m.IDcID); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateInterfaceMode(formats strfmt.Registry) error {

	if err := validate.Required("interfaceMode", "body", m.InterfaceMode); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateMemAmount(formats strfmt.Registry) error {

	if err := validate.Required("memAmount", "body", m.MemAmount); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateMemFrequency(formats strfmt.Registry) error {

	if err := validate.Required("memFrequency", "body", m.MemFrequency); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateMemSize(formats strfmt.Registry) error {

	if err := validate.Required("memSize", "body", m.MemSize); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateMemSpec(formats strfmt.Registry) error {

	if err := validate.Required("memSpec", "body", m.MemSpec); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateMemType(formats strfmt.Registry) error {

	if err := validate.Required("memType", "body", m.MemType); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateNicAmount(formats strfmt.Registry) error {

	if err := validate.Required("nicAmount", "body", m.NicAmount); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateNicRate(formats strfmt.Registry) error {

	if err := validate.Required("nicRate", "body", m.NicRate); err != nil {
		return err
	}

	return nil
}

func (m *CreateDeviceTypeRequest) validateVolumes(formats strfmt.Registry) error {

	if err := validate.Required("volumes", "body", m.Volumes); err != nil {
		return err
	}

	for i := 0; i < len(m.Volumes); i++ {
		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create device type request based on the context it is used
func (m *CreateDeviceTypeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDeviceTypeRequest) contextValidateVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Volumes); i++ {

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateDeviceTypeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDeviceTypeRequest) UnmarshalBinary(b []byte) error {
	var res CreateDeviceTypeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
