package request

import (
	"regexp"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type QueryMessagesRequest struct {
	Pageable
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
	MessageID string `json:"messageId" validate:"required"`
}

func (req *GetMessageByIdRequest) Validate(logger *log.Logger) {
	validate(req, logger)
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
	// 收件人邮箱地址
	To string `json:"to"`
}

func (req *MailDialRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(req.AdminAddr) {
		panic(constant.BuildInvalidArgumentWithArgs("admin_addr格式不正确", "admin_addr invalidate"))
	}
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
	// 收件人邮箱地址
	To string `json:"to"`
}

func (req *MailSaveRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type SaveIsPushMailRequest struct {
	// 是否推送邮件告警，0不推送，1推送
	IsPush string `json:"is_push"`
}

func (req *SaveIsPushMailRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
