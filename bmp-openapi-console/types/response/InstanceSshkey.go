package response

type InstancesSshkeyResponse struct {
	ID          uint64 `json:"-"`           // ID
	InstanceID  string `json:"instanceId"`  // 实例ID
	SSHkeyID    string `json:"sshkeyId"`    // sshkeyid
	CreatedBy   string `json:"createdBy"`   // 创建者
	UpdatedBy   string `json:"updatedBy"`   // 更新者
	CreatedTime string `json:"createdTime"` // 创建时间
	UpdatedTime string `json:"updatedTime"` // 更新时间
}

type InstancesSshkeyInfoResponse struct {
	// 实例Id列表
	InstanceIds []string `json:"instanceIds"`
	//SshkeyIds   []string `json:"sshkeyIds"`   // 秘钥Id
}
