package response

type Os struct {
	ID     uint64 `json:"id"`     // ID
	OsID   string `json:"osId"`   // 操作系统uuid
	OsName string `json:"osName"` // 操作系统名称
	OsType string `json:"osType"` // 操作系统分类:linux/windows

	Architecture string `json:"architecture"` // 架构:x86/x64/i386/
	Bits         int    `json:"bits"`         // 指令宽度:64/32位
	OsVersion    string `json:"osVersion"`    // 操作系统版本
	SysUser      string `json:"sysUser"`      // 管理员账户
	CreatedBy    string `json:"createdBy"`    // 创建者
	UpdatedBy    string `json:"updatedBy"`    // 更新者
	CreatedTime  string `json:"createdTime"`  // 创建时间戳
	UpdatedTime  string `json:"updatedTime"`  // 更新时间戳
	DeletedTime  string `json:"deletedTime"`  // 删除时间戳
	IsDel        int8   `json:"isDel"`        // 是否删除0未删除 1已删除
}

type OsList struct {
	Oss []*Os `json:"oss"`
	//PageNumber int64 `json:"pageNumber"`
	//PageSize   int64 `json:"pageSize"`
	//TotalCount int64 `json:"totalCount"`
}
type OsInfo struct {
	Os *Os `json:"os"`
}
type OsId struct {
	OsId string `json:"osId"`
}
