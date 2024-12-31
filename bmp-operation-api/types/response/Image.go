package response

type Image struct {
	ID        uint64 `json:"-"`         // ID
	ImageID   string `json:"imageId"`   // 镜像uuid
	ImageName string `json:"imageName"` // 镜像名称
	OsID      string `json:"osId"`      // OSID
	Format    string `json:"format"`    // 镜像格式（qcow2、tar）
	//boot模式
	BootMode        string `json:"bootMode"`
	Filename        string `json:"filename"`        // 镜像文件名称
	URL             string `json:"url"`             // 镜像源路径
	Hash            string `json:"hash"`            // 镜像校验码
	Source          string `json:"source"`          // 镜像来源(common通用、customize定制、user_defined自定义)
	SourceName      string `json:"sourceName"`      // 镜像来源(common通用、customize定制、user_defined自定义)
	Description     string `json:"description"`     // 描述
	SystemPartition string `json:"systemPartition"` // 系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：10GiB，swap）
	DataPartition   string `json:"dataPartition"`   // 数据分区信息
	CreatedBy       string `json:"createdBy"`       // 创建者
	UpdatedBy       string `json:"updatedBy"`       // 更新者
	CreatedTime     string `json:"createdTime"`     // 创建时间戳
	UpdatedTime     string `json:"updatedTime"`     // 更新时间戳
	//DeletedTime     int    `json:"deletedTime"`     // 删除时间戳
	//IsDel           int8   `json:"isDel"`           // 是否删除0未删除 1已删除
	//Platform      string `json:"platform"`  // suse/centos/ubuntu
	Architecture  string `json:"architecture"` // 架构:x86/x64/i386/
	OsType        string `json:"osType"`       // 操作系统分类:linux/windows
	OsVersion     string `json:"osVersion"`    // 操作系统版本
	DeviceTypeNum int    `json:"deviceTypeNum"`
	IsBind        bool   `json:"isBind"`
	//Os              OS     `json:"os"`
}
type ImageList struct {
	Images     []Image `json:"images"`
	PageNumber int64   `json:"pageNumber"`
	PageSize   int64   `json:"pageSize"`
	TotalCount int64   `json:"totalCount"`
}
type ImageInfo struct {
	Image *Image `json:"image"`
}
type ImageId struct {
	ImageId string `json:"imageId"`
}
type ImageUpload struct {
	Format   string `json:"format"`
	FileName string `json:"fileName"`
	Url      string `json:"url"`
	Hash     string `json:"hash"`
}
type Partition struct {
	//format,如[swap, xfs]
	FsType string `json:"format"`
	//point,如[swap, /, /var]
	MountPoint string `json:"point"`
	// 分区大小, MB为单位
	Size int `json:"size"`
}
