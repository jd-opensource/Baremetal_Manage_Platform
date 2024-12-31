package response

import "coding.jd.com/aidc-bmp/bmp-openapi/dao/messageDao"

type Message struct {
	messageDao.WebMessage
	//独一处的下划线，跟sdk一致
	IDcID       string `json:"idc_id"`
	IDcName     string `json:"idc_name"`
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	DeviceID    string `json:"device_id"`
}

type MessageList struct {
	Messages   []*Message `json:"messages"`
	PageNumber int64      `json:"pageNumber"`
	PageSize   int64      `json:"pageSize"`
	TotalCount int64      `json:"totalCount"`
}

type MessageStatistic struct {
	TotalCount  int64 `json:"totalCount"`
	UnreadCount int64 `json:"unreadCount"`
}

type HasUnreadMessage struct {
	HasUnread bool `json:"hasUnread"`
}

type MessageWithNextPrev struct {
	Message       *Message `json:"message"`
	NextMessageId string   `json:"nextMessageId"`
	PrevMessageId string   `json:"prevMessageId"`
}

type MessageTypesResp map[string][]string

type Mail struct {
	// 邮件服务器地址
	ServerAddr string `json:"server_addr"`
	// 邮件服务器端口
	ServerPort string `json:"server_port"`
	// 管理员邮箱
	AdminAddr string `json:"admin_addr"`
	// 管理员邮箱密码
	AdminPass string `json:"admin_pass"`
	// 是否推送邮件告警，0不推送，1推送
	IsPush string `json:"is_push"`
	// 收件人邮箱
	To string `json:"to"`
	// 邮件功能是否验证过正常。1表示正常，其他表示异常
	IsPass string `json:"is_pass"`
}
