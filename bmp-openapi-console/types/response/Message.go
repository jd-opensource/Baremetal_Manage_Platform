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
