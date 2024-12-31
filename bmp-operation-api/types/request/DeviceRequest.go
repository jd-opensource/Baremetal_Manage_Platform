package request

import (
	"mime/multipart"

	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
)

type CreateDeviceRequest struct {
	IdcId string `json:"idcId" validate:"required"`
	// DeviceTypeId string                   `json:"deviceTypeId" validate:"required"`
	Devices []*response.UploadDevice `json:"devices" validate:"required"`
}
type UploadDeviceRequest struct {
	DeviceFile multipart.File `json:"deviceFile"`
}

// 【设备管理】【设备列表】
type QueryDeviceListRequest struct {
	BaseRequest
	PagingRequest

	Sn           string `json:"sn"`
	DeviceTypeID string `json:"deviceTypeId"`
	DeviceSeries string `json:"deviceSeries"`
	ManageStatus string `json:"manageStatus"`
	IloIP        string `json:"iloIp"`
	InstanceID   string `json:"instanceId"`
	InstanceName string `json:"instanceName"`
	UserID       string `json:"userId"`
	UserName     string `json:"userName"`
	Show         string `json:"show"`
	IsAll        string `json:"isAll"`
	// ExportType 非空表示要导出
	ExportType string `json:"exportType"`
	// "1"表示采集成功，"2"表示未采集, "3"表示采集中,4表示采集失败，不传表示全部
	CollectStatus string `json:"collectStatus"`
}

// 【设备管理】【设备详情】
type QueryDeviceInfoRequest struct {
	Show     string `json:"show"`
	DeviceID string `json:"deviceId"`
}

// 【设备管理】【设备磁盘详情】
type QueryDeviceDisksRequest struct {
	DeviceID string `json:"deviceId"`
}

// 【设备管理】【设备下架】支持多sn
type UnPutawayDeviceRequest struct {
	Sns      string `json:"sns"`
	DeviceID string `json:"deviceId"`
}

// 【设备管理】【设备上架】支持多sn
type PutawayDeviceRequest struct {
	Sns      string `json:"sns"`
	DeviceID string `json:"deviceId"`
}

// 【设备管理】【删除设备】
type ModifyDeviceRequest struct {
	Description *string `json:"description"`
}

// 【设备管理】【删除设备】
type DeleteDeviceRequest struct {
	DeviceId string `json:"deviceId"`
}

type CollectDevice struct {
	// sn
	// required: true
	Sn string `json:"sn"`
	// 可选参数。 传参:True, False 。True表示传入的raid_driver值将覆盖已适配机器的raid_driver
	// Extensions:
	// x-nullable: true
	AllowOverride bool `json:"allowOverride"`
}

type AssociateDeviceDisksRequest struct {
	// device uuid
	// required: true
	DeviceID string `json:"deviceId" validate:"required"` // device uuid
	// devicetype uuid
	// required: true
	DeviceTypeID string `json:"deviceTypeId" validate:"required"` // devicetype uuid
	// volumeid和磁盘uuid列表
	// required: true
	Volumes []*AssociateDeviceDiskSpec `json:"volumes" validate:"required"`
}

type AssociateDeviceDiskSpec struct {
	VolumeID string   `json:"volumeId" validate:"required"`
	DiskIDs  []string `json:"diskId" validate:"required"`
}

type GetAssociatedDisksRequest struct {
	// device uuid
	// required: true
	DeviceID string `json:"deviceId" validate:"required"` // device uuid
	// devicetype uuid
	// required: true
	DeviceTypeID string `json:"deviceTypeId" validate:"required"` // devicetype uuid
	// volume uuid
	// required: true
	VolumeID string `json:"volumeId" validate:"required"`
}

type DeviceAssociateDeviceTypeRequest struct {
	// 设备类型id
	// required: true
	DeviceTypeID string `json:"deviceTypeId" validate:"required"`
	// 设备ID
	// required: true
	DeviceID string `json:"deviceId" validate:"required"`
}
