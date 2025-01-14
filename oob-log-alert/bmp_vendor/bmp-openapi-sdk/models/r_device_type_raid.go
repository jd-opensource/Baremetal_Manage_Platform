// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RDeviceTypeRaid r device type raid
//
// swagger:model RDeviceTypeRaid
type RDeviceTypeRaid struct {

	// 可用分区值，单位GB
	AvailableValue int64 `json:"availableValue,omitempty"`

	// 创建者
	CreatedBy string `json:"createdBy,omitempty"`

	// 创建时间
	CreatedTime string `json:"createdTime,omitempty"`

	// 设备类型id
	DeviceTypeID string `json:"deviceTypeId,omitempty"`

	// 磁盘类型：SAS/SATA/SSD/NVME
	DiskInterfaceType string `json:"DiskInterfaceType,omitempty"`

	// 磁盘类型：SAS/SATA/SSD/NVME
	DiskType string `json:"diskType,omitempty"`

	// RAID uuid
	RaidID string `json:"raidId,omitempty"`

	// 系统盘noraid模式真实数量;此模式多块盘只能用一块
	SystemPartitionCount int64 `json:"systemPartitionCount,omitempty"`

	// 更新者
	UpdatedBy string `json:"updatedBy,omitempty"`

	// 更新时间
	UpdatedTime string `json:"updatedTime,omitempty"`

	// 磁盘详细信息
	VolumeDetail string `json:"volumeDetail,omitempty"`

	// 磁盘类型 system/data
	VolumeType string `json:"volumeType,omitempty"`
}

// Validate validates this r device type raid
func (m *RDeviceTypeRaid) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r device type raid based on context it is used
func (m *RDeviceTypeRaid) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RDeviceTypeRaid) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RDeviceTypeRaid) UnmarshalBinary(b []byte) error {
	var res RDeviceTypeRaid
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
