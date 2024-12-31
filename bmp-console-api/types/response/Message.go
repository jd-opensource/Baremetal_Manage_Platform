package response

import sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"

type MessageList struct {
	Messages   []*sdkModels.Message `json:"messages"`
	PageNumber int64                `json:"pageNumber"`
	PageSize   int64                `json:"pageSize"`
	TotalCount int64                `json:"totalCount"`
}

type MessageStatistic struct {
	TotalCount  int64 `json:"totalCount"`
	UnreadCount int64 `json:"unreadCount"`
}

type HasUnreadMessage struct {
	HasUnread bool `json:"hasUnread"`
}

type MessageWithNextPrev struct {
	Message       *sdkModels.Message `json:"message"`
	NextMessageId string             `json:"nextMessageId"`
	PrevMessageId string             `json:"prevMessageId"`
}
