package response

type Sshkey struct {
	// id
	Id int64 `json:"id"`
	// 秘钥uuid
	SshkeyId string `json:"sshkeyId"` // 秘钥uuid
	// 所属用户uuid
	UserId string `json:"userId"`
	// 秘钥名称
	Name string `json:"name"`
	// 公钥，格式：ssh-rsa AAA
	Key string `json:"key"`
	// 公钥指纹
	FingerPrint string `json:"fingerPrint"`
	// 创建者
	CreatedBy string `json:"createdBy"`
	// 更新者
	UpdatedBy string `json:"updatedBy"`
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// 更新时间
	UpdatedTime string `json:"updatedTime"`
}
type SshkeyList struct {
	// sshkey实体列表
	Sshkeys []*Sshkey `json:"sshkeys"`
	// 页数
	PageNumber int64 `json:"pageNumber"`
	// 页大小
	PageSize int64 `json:"pageSize"`
	// 总条数
	TotalCount int64 `json:"totalCount"`
}
type SshkeyInfo struct {
	// sshkey实体
	Sshkey *Sshkey `json:"sshkey"`
}
type SshkeyId struct {
	// sshkey uuid
	SshkeyId string `json:"sshkeyId"`
}
