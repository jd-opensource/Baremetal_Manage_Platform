package response

type Os struct {
	// ID
	ID uint64 `json:"id"`
	// 操作系统uuid
	OsID string `json:"osId"`
	// 操作系统名称
	OsName string `json:"osName"`
	// 操作系统分类:linux/windows
	OsType string `json:"osType"`
	// 架构:x86/x64/i386/
	Architecture string `json:"architecture"`
	// 指令宽度:64/32位
	Bits int `json:"bits"`
	// 操作系统版本
	OsVersion string `json:"osVersion"`
	// 管理员账户
	SysUser string `json:"sysUser"`
	// 创建者
	CreatedBy string `json:"createdBy"`
	// 更新者
	UpdatedBy string `json:"updatedBy"`
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// 更新时间
	UpdatedTime string `json:"updatedTime"`
	// 删除时间
	DeletedTime string `json:"deletedTime"`
	// 是否删除0未删除 1已删除
	IsDel int8 `json:"isDel"`
}

type OsList struct {
	// 操作系统列表
	Oss []*Os `json:"oss"`
}
type OsInfo struct {
	// 操作系统实体
	Os *Os `json:"os"`
}
type OsId struct {
	// 操作系统uuid
	OsId string `json:"osId"`
}
