package request

type QueryMessagesRequest struct {
	PagingRequest
	// 0：导出全部实力；1：导出选中实例；2：导出搜索结果
	ExportType string `json:"exportType"`
	// 是否显示全部，取值为1时表示全部
	IsAll string `json:"isAll"`
	//0：未读；1：已读；""不传表示全部
	HasRead string `json:"hasRead"`
	// 消息类型
	MessageType string `json:"messageType"`
	// 消息子类型
	MessageSubType string `json:"messageSubType"`
	// 内容模糊搜索
	Detail string `json:"detail"`
}

type ReadMessagesRequest struct {

	// required: true
	MessageIds []string `json:"messageIds" validate:"required"`
}
type DeleteMessagesRequest struct {

	// required: true
	MessageIds []string `json:"messageIds" validate:"required"`
}

type GetMessageByIdRequest struct {
	// required: true
	MessageId string `json:"messageId" validate:"required"`
}

type MailDialRequest struct {
	// 邮件服务器地址
	// required: true
	ServerAddr string `json:"server_addr" validate:"required"`
	// 邮件服务器端口
	// required: true
	ServerPort string `json:"server_port" validate:"required"`
	// 管理员邮箱
	// required: true
	AdminAddr string `json:"admin_addr" validate:"required"`
	// 管理员邮箱密码
	// required: true
	AdminPass string `json:"admin_pass" validate:"required"`
	//收件人邮箱
	To string `json:"to"`
}

type MailSaveRequest struct {
	// 邮件服务器地址
	// required: true
	ServerAddr string `json:"server_addr" validate:"required"`
	// 邮件服务器端口
	// required: true
	ServerPort string `json:"server_port" validate:"required"`
	// 管理员邮箱
	// required: true
	AdminAddr string `json:"admin_addr" validate:"required"`
	// 管理员邮箱密码
	// required: true
	AdminPass string `json:"admin_pass" validate:"required"`
	//收件人邮箱
	To string `json:"to"`
}

type SaveIsPushMailRequest struct {
	// 是否打开推送，"0","1"
	// required: true
	IsPush string `json:"is_push" validate:"required"`
}
