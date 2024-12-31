package request

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type CreateImageRequest struct {
	// 镜像名称
	// required: true
	ImageName string `json:"imageName" validate:"required,min=1,max=200"`
	// 体系架构
	// required: true
	Architecture string `json:"architecture" validate:"required,oneof=x86_64 ARM64(aarch64) i386"`
	// 操作系统平台
	// required: true
	OsType string `json:"osType" validate:"required"`
	// 版本号
	// required: true
	Version string `json:"Version" validate:"required"`
	// 镜像文件类型
	// required: true
	Format string `json:"format" validate:"required"`
	// 镜像文件名称
	// required: true
	Filename string `json:"filename" validate:"required"`
	// 镜像上传地址
	// required: true
	Url string `json:"url" validate:"required"`
	// 文件hash值
	// required: true
	Hash string `json:"hash"`
	// 镜像类型，预置，自定义
	Source string `json:"source" validate:"oneof=common user_defined"`
	// 描述
	Description string `json:"description"`
	// 系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：50GiB，swap）
	SystemPartition string `json:"systemPartition"`
	// 数据分区信息
	DataPartition string `json:"dataPartition"`
}

func (req *CreateImageRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

//镜像默认分区
//[{"format":"swap","point":"swap","size":10240},{"format":"xfs","point":"/","size":51200}]
type Partition struct {
	FsType     string `json:"format"` //format
	MountPoint string `json:"point"`  //point
	Size       int    `json:"size"`
}

type ModifyImageRequest struct {
	// 修改设备描述
	// Extensions:
	// x-nullable: true
	Description *string `json:"description" validate:"omitempty,max=256"`
}

//// 镜像名称
//ImageName string `json:"imageName" validate:"min=1,max=200"`
//// 操作系统id
//OsID string `json:"osId" validate:"omitempty,max=256"`
//// 文件类型
//Format string `json:"format" validate:"omitempty,max=256"`
//// 文件名称
//Filename string `json:"filename" validate:"omitempty,max=256"`
//// 镜像上传地址
//Url string `json:"url" validate:"omitempty,max=500"`
//// 文件hash值
//Hash string `json:"hash" validate:"omitempty,max=256"`
//// 镜像类型，预置，自定义
//Source string `json:"source" validate:"omitempty,oneof=common user_defined"`
//// 修改设备描述
//// Extensions:
//// x-nullable: true
//Description *string `json:"description" validate:"omitempty,max=256"`
//// 系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：50GiB，swap）
//SystemPartition string `json:"systemPartition" validate:"omitempty,max=256"`
//// 数据分区信息
//DataPartition string `json:"dataPartition" validate:"omitempty,max=256"`

func (req *ModifyImageRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type QueryImagesRequest struct {
	// 镜像ID
	ImageID string `json:"imageId"`
	// 镜像名称
	ImageName string `json:"imageName"`
	// 设备类型id
	DeviceTypeID string `json:"deviceTypeId"`
	// 版本号
	Version string `json:"version"`
	// 操作系统ID
	OsID string `json:"osId"`
	// 镜像ID，数组，支持多个
	ImageIDs []string `json:"imageIds"`
	// 镜像类型，预置，自定义
	Source string `json:"source"`
	// 体系架构
	Architecture string `json:"architecture"`
	// 操作系统平台
	OsType string `json:"osType"`
	Pageable
	// 是否显示全部
	IsAll string `json:"isAll"`
}

func (req *QueryImagesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type AssociateImageRequest struct {
	// 设备类型id
	// required: true
	DeviceTypeID string `json:"deviceTypeId" validate:"required"`
	// 镜像ID，数组，支持多个
	// required: true
	ImageIDs []string `json:"imageIds" validate:"required"`
}

func (req *AssociateImageRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type DissociatedImageRequest struct {
	// 设备类型id
	// required: true
	DeviceTypeID string `json:"deviceTypeId" validate:"required"`
	// 镜像ID
	// required: true
	ImageID string `json:"imageId" validate:"required"`
}

func (req *DissociatedImageRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type AssociateDeviceTypeRequest struct {
	// 镜像ID，
	// required: true
	ImageID string `json:"imageId" validate:"required"`
	// 设备类型id，数组，支持多个
	// required: true
	DeviceTypeIDs []string `json:"deviceTypeIds" validate:"required"`
}

func (req *AssociateDeviceTypeRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type QueryImageDeviceTypesRequest struct {
	// 镜像ID
	// required: true
	ImageID string `json:"imageId" validate:"required"`
	// 体系架构
	Architecture string `json:"architecture"`
	Pageable
	//镜像是否绑定了机型，0查询该镜像没有绑定的机型列表 1查询该镜像已经绑定了的机型列表
	IsBind string `json:"isBind"`
	// 是否显示全部
	IsAll string `json:"isAll"`
}

func (req *QueryImageDeviceTypesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
