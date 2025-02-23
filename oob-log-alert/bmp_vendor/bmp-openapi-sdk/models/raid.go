// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Raid raid
//
// swagger:model Raid
type Raid struct {

	// 可用分区值，单位GB
	AvailableValue int64 `json:"availableValue,omitempty"`

	// 中文描述
	DescriptionEn string `json:"descriptionEn,omitempty"`

	// description
	DescriptionZh string `json:"descriptionZh,omitempty"`

	// 设备类型uuid
	DeviceTypeID string `json:"deviceTypeId,omitempty"`

	// 磁盘类型：SAS/SATA/SSD/NVME
	DiskType string `json:"diskType,omitempty"`

	// raid 名称
	Name string `json:"name,omitempty"`

	// raid uuid
	RaidID string `json:"raidId,omitempty"`

	// 系统盘分区数量
	SystemPartitionCount int64 `json:"systemPartitionCount,omitempty"`

	// 磁盘详细信息
	VolumeDetail string `json:"volumeDetail,omitempty"`

	// 磁盘类型
	VolumeType string `json:"volumeType,omitempty"`
}

// Validate validates this raid
func (m *Raid) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this raid based on context it is used
func (m *Raid) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Raid) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Raid) UnmarshalBinary(b []byte) error {
	var res Raid
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
