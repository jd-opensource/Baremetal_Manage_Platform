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

// VolumeRaids volume raids
//
// swagger:model VolumeRaids
type VolumeRaids struct {

	// RaidCan       string `gorm:"column:raid_can" json:"raidCan"`             // RAID配置： (RAID,NO RAID)
	// Raid          string `gorm:"column:raid" json:"raid"`                    // RAID模式：RAID1,RIAD10等
	CreatedBy string `json:"createdBy,omitempty"`

	// created time
	CreatedTime int64 `json:"createdTime,omitempty"`

	// deleted time
	DeletedTime int64 `json:"deletedTime,omitempty"`

	// 设备类型uuid
	DeviceTypeID string `json:"deviceTypeId,omitempty"`

	// 硬盘类型（SSD,HDD）
	DiskType string `json:"diskType,omitempty"`

	// 接口类型（SATA,SAS,NVME,不限制）
	InterfaceType string `json:"interfaceType,omitempty"`

	// is del
	IsDel int8 `json:"isDel,omitempty"`

	// raids
	Raids []*RVolumeRaid `json:"raids"`

	// updated by
	UpdatedBy string `json:"updatedBy,omitempty"`

	// updated time
	UpdatedTime int64 `json:"updatedTime,omitempty"`

	// 硬盘数量（最低块数）
	VolumeAmount int64 `json:"volumeAmount,omitempty"`

	// 卷uuid
	VolumeID string `json:"volumeId,omitempty"`

	// 卷名称
	VolumeName string `json:"volumeName,omitempty"`

	// 单盘大小（最小容量）
	VolumeSize string `json:"volumeSize,omitempty"`

	// 卷类型：系统卷，数据卷
	VolumeType string `json:"volumeType,omitempty"`

	// 硬盘单位（GB,TB）
	VolumeUnit string `json:"volumeUnit,omitempty"`
}

// Validate validates this volume raids
func (m *VolumeRaids) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRaids(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeRaids) validateRaids(formats strfmt.Registry) error {
	if swag.IsZero(m.Raids) { // not required
		return nil
	}

	for i := 0; i < len(m.Raids); i++ {
		if swag.IsZero(m.Raids[i]) { // not required
			continue
		}

		if m.Raids[i] != nil {
			if err := m.Raids[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("raids" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("raids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this volume raids based on the context it is used
func (m *VolumeRaids) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRaids(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeRaids) contextValidateRaids(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Raids); i++ {

		if m.Raids[i] != nil {
			if err := m.Raids[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("raids" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("raids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeRaids) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeRaids) UnmarshalBinary(b []byte) error {
	var res VolumeRaids
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}