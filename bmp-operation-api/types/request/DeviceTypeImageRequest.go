package request

type QueryDeviceTypeImagePageRequest struct {
	BaseRequest
	PagingRequest
	// 机型id
	DeviceTypeID string `json:"deviceTypeId"`
	// 镜像id
	ImageID string `json:"imageId"`
	// 体系架构
	Architecture string `json:"architecture"`
	// 操作系统平台
	OsType string `json:"osType"`
	// 镜像名称
	ImageName string `json:"imageName"`
	// 版本号
	Version string `json:"version"`
	// 操作系统ID
	OsID string `json:"osId"`
	// 镜像类型，预置，自定义
	Source string `json:"source"`
	// 是否显示全部
	IsAll string `json:"isAll"`
}

type AssociateDeviceTypeImageRequest struct {
	BaseRequest
	DeviceTypeID string `json:"deviceTypeId"`
	ImageIds     string `json:"imageIds"` // 逗号分隔
}

type DisassociateDeviceTypeImageRequest struct {
	BaseRequest
	DeviceTypeID string `json:"deviceTypeId"`
	ImageId      string `json:"imageId"`
}
