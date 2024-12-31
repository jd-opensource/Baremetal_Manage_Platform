package request

import (
	"mime/multipart"
)

type CreateImageRequest struct {
	ImageName    string `json:"imageName" validate:"required,min=1,max=200"`
	Architecture string `json:"architecture" validate:"required"`
	OsType       string `json:"osType" validate:"required"`
	Version      string `json:"version" validate:"required"`
	//OsID            string         `json:"osId" validate:"required"`
	//ImageFile       multipart.File `json:"imageFile"`
	Format string `json:"format"`
	//boot模式[UEFI Legacy/BIOS]支持多选，逗号分隔
	BootMode        string `json:"bootMode" validate:"omitempty"`
	FileName        string `json:"fileName"`
	Url             string `json:"url"`
	Hash            string `json:"hash"`
	Description     string `json:"description"`     // 描述
	SystemPartition string `json:"systemPartition"` // 系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：10GiB，swap）
	DataPartition   string `json:"dataPartition"`   // 数据分区信息
}
type UploadImageRequest struct {
	ImageFile multipart.File `json:"imageFile"`
}

//Format          string `json:"format" validate:"required"`
//Filename        string `json:"filename" validate:"required"`
//Url             string `json:"url" validate:"required"`
//Hash            string `json:"hash"`
//Source          string `json:"source"`
//镜像默认分区
//[{"format":"swap","point":"swap","size":10240},{"format":"xfs","point":"/","size":51200}]
type Partition struct {
	FsType     string `json:"format"` //format
	MountPoint string `json:"point"`  //point
	Size       int    `json:"size"`
}

//func (c *CreateImageRequest) Validate(logger *log.Logger) {
//	if err := validator.New().Struct(c); err != nil {
//		logger.Warn("CreateImageRequest validate error:", err.Error())
//		panic(exception.CommonParamValid)
//	}
//}

type ModifyImageRequest struct {
	Description *string `json:"description"` // 描述
}

//func (req *ModifyImageRequest) Validate(logger *log.Logger) {
//	if err := validator.New().Struct(req); err != nil {
//		logger.Warn("ModifyImageRequest.Validate error:", err.Error())
//		panic(exception.CommonParamValid)
//	}
//	if match, _ := regexp.MatchString(validation.REGEX_ID, req.OsID); !match {
//		logger.Warn("ModifyImageRequest.OsId invalid:", req.OsID)
//		panic(exception.CommonParamValid)
//	}
//	if match, _ := regexp.MatchString(validation.REGEX_SOURCE_TYPE, req.Source); !match {
//		logger.Warn("ModifyImageRequest.Source invalid:", req.Source)
//		panic(exception.CommonParamValid)
//	}
//
//}

type QueryImagesRequest struct {
	ImageID      string   `json:"imageId"`
	ImageName    string   `json:"imageName"`
	DeviceTypeID string   `json:"deviceTypeId"`
	Version      string   `json:"version"`
	OsID         string   `json:"osId"`
	ImageIDs     []string `json:"imageIds"`
	Source       string   `json:"source"`
	PageNumber   int64    `json:"pageNumber"`
	PageSize     int64    `json:"pageSize"`
	IsAll        string   `json:"isAll"`
	ExportType   string   `json:"exportType"`
	Architecture string   `json:"architecture"`
	OsType       string   `json:"osType"`
}

//func (req *QueryImagesRequest) Validate(logger *log.Logger) {
//	if err := validator.New().Struct(req); err != nil {
//		logger.Warn("QueryImagesRequest validate error:", err.Error())
//		panic(exception.CommonParamValid)
//	}
//
//	//if req.Source != "" {
//	//	if match, _ := regexp.MatchString(validation.REGEX_SOURCE_TYPE, req.Source); !match {
//	//		logger.Warn("QueryImagesRequest Source invalid")
//	//		panic(exception.CommonParamValid)
//	//	}
//	//}
//
//}

type AssociateImageRequest struct {
	//DeviceTypeImages []DeviceTypeImagePair `json:"device_type_images"`
	DeviceTypeID string   `json:"deviceTypeId"`
	ImageIDs     []string `json:"imageIds"`
}

type DeviceTypeImagePair struct {
	DeviceType string `json:"device_type"`
	ImageId    string `json:"image_id"`
}

//func (c *AssociateImageRequest) Validate(logger *log.Logger) {
//	if err := validator.New().Struct(c); err != nil {
//		logger.Warn("AssociateImageRequest.Validate error:", err.Error())
//		panic(exception.CommonParamValid)
//	}
//}

type DissociatedImageRequest struct {
	DeviceTypeID string `json:"deviceTypeId"`
	ImageID      string `json:"ImageId"`
}

//func (req *DissociatedImageRequest) Validate(logger *log.Logger) {
//	if err := validator.New().Struct(req); err != nil {
//		logger.Warn("DissociatedImageRequest.Validate error:", err.Error())
//		panic(exception.CommonParamValid)
//	}
//}
type QueryImageDeviceTypesRequest struct {
	// 镜像ID
	// required: true
	ImageID string `json:"imageId" validate:"required"`
	// 体系架构
	Architecture string `json:"architecture"`
	PageNumber   int64  `json:"pageNumber"`
	PageSize     int64  `json:"pageSize"`
	//镜像是否绑定了机型，0查询该镜像没有绑定的机型列表 1查询该镜像已经绑定了的机型列表
	IsBind string `json:"isBind"`
	// 是否显示全部
	IsAll string `json:"isAll"`
}

//镜像绑定机型
type AssociateImageDeviceTypeRequest struct {
	BaseRequest
	ImageId       string `json:"ImageId"`
	DeviceTypeIDs string `json:"deviceTypeIds"` // 逗号分隔
}

type QueryImagesByDeviceTypeRequest struct {
	// Pin        string `json:"pin"`
	IdcID        string `json:"idc_id"`
	DeviceTypeID string `json:"deviceTypeId"`
	// OsType     string   `json:"os_type"`
	// OsVersion  string   `json:"os_version"`
	// OsID       string   `json:"os_id"`
	// Source     string   `json:"source"`
	// ImageIds   []string `json:"image_ids"`
}
