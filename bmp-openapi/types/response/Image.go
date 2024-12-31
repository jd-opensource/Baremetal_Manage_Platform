package response

type Image struct {
	ID uint64 `json:"id"` // ID
	// 镜像uuid
	ImageID string `json:"imageId"`
	// 镜像名称
	ImageName string `json:"imageName"`
	// 源os uuid
	OsID string `json:"osId"`
	// 镜像格式（qcow2、tar）
	Format string `json:"format"`
	//boot模式
	BootMode string `json:"bootMode"`
	// 镜像文件名称
	Filename string `json:"filename"`
	// 镜像源路径
	URL string `json:"url"`
	// 镜像校验码
	Hash string `json:"hash"`
	// 镜像来源(common通用、customize定制、user_defined自定义)
	Source string `json:"source"`
	// 镜像来源(common通用、customize定制、user_defined自定义)
	SourceName string `json:"sourceName"`
	// 描述
	Description string `json:"description"`
	// 系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：10GiB，swap）
	SystemPartition string `json:"systemPartition"`
	// 数据分区信息
	DataPartition string `json:"dataPartition"`
	// 创建者
	CreatedBy string `json:"createdBy"`
	// 更新者
	UpdatedBy string `json:"updatedBy"`
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// 更新时间
	UpdatedTime string `json:"updatedTime"`
	// 架构:x86/x64/i386/
	Architecture string `json:"architecture"`
	//CentOS 7.2 64-bit
	OsName string `json:"osName"`
	// 操作系统分类:linux/windows
	OsType string `json:"osType"`
	// 操作系统版本
	OsVersion string `json:"osVersion"`
	//绑定了机型数量
	DeviceTypeNum int `json:"deviceTypeNum"`
	//是否绑定了某个机型
	IsBind bool `json:"isBind"`
}
type ImageList struct {
	// 镜像实体列表
	Images []*Image `json:"images"`
	// 页数
	PageNumber int64 `json:"pageNumber"`
	// 页大小
	PageSize int64 `json:"pageSize"`
	// 总条数
	TotalCount int64 `json:"totalCount"`
}
type ImageInfo struct {
	// 镜像实体
	Image *Image `json:"image"`
}
type ImageId struct {
	// 镜像uuid
	ImageId string `json:"imageId"`
}
