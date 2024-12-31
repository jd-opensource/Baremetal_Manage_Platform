package request

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

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
	// 内容，支持模糊搜索
	Detail string `json:"detail"`
}

func (req *QueryMessagesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type ReadMessagesRequest struct {

	// required: true
	MessageIds []string `json:"messageIds" validate:"required"`
}

func (req *ReadMessagesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type DeleteMessagesRequest struct {

	// required: true
	MessageIds []string `json:"messageIds" validate:"required"`
}

func (req *DeleteMessagesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type GetMessageByIdRequest struct {
	// required: true
	MessageId string `json:"messageId" validate:"required"`
}

func (req *GetMessageByIdRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
