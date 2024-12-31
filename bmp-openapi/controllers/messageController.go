package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/messageLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"

	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type MessageController struct {
	BaseController
}

// HasUnreadMessage ...
// swagger:route GET /messages/hasUnreadMessage message hasUnreadMessage
//
// HasUnreadMessage 获取有没有未读消息
//
//     Responses:
//       200: hasUnreadMessage
//       default: ErrorResponse

func (c *MessageController) HasUnreadMessage() {
	defer c.CatchException()

	res, err := messageLogic.HasUnreadMessage(c.logPoints)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.HasUnreadMessage{
		HasUnread: res,
	}
}

// GetMessageList ...
// swagger:route GET /messages message getMessageList
//
// GetMessageList 获取message列表(分页)
//
//     Responses:
//       200: getMessageList
//       default: ErrorResponse
func (c *MessageController) GetMessageList() {
	defer c.CatchException()

	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}

	req := &requestTypes.QueryMessagesRequest{
		IsAll:          c.GetString("isAll"),
		HasRead:        c.GetString("hasRead"),
		MessageType:    c.GetString("messageType"),
		MessageSubType: c.GetString("messageSubType"),
		Detail:         c.GetString("detail"),
	}
	req.Validate(c.logPoints)

	res, cnt, err := messageLogic.GetPageMessages(c.logPoints, req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.MessageList{
		Messages:   res,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: cnt,
	}
}

// GetMessageStatistic ...
// swagger:route GET /messages/statistic message getMessageStatistic
//
// GetMessageStatistic 获取message总数和未读数
//
//     Responses:
//       200: getMessageStatistic
//       default: ErrorResponse
func (c *MessageController) GetMessageStatistic() {
	defer c.CatchException()

	total, unread, err := messageLogic.GetMessageStatistic(c.logPoints)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = response.MessageStatistic{

		TotalCount:  total,
		UnreadCount: unread,
	}
}

// ReadMessage ...
// swagger:route PUT /messages/doRead message readMessage
//
// ReadMessage 将消息设置为已读(可多条)
//
//     Responses:
//       200: readMessage
//       default: ErrorResponse
func (c *MessageController) ReadMessage() {
	defer c.CatchException()

	req := &requestTypes.ReadMessagesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ReadMessage parse ReadMessagesRequest body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	res, err := messageLogic.ReadMessages(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// GetMessageById ...
// swagger:route GET /messages/getMessageById message getMessageById
//
// GetMessageById 获取消息详情，包括上一条/下一条的messageid
//
//     Responses:
//       200: getMessageById
//       default: ErrorResponse

func (c *MessageController) GetMessageById() {
	defer c.CatchException()

	req := &requestTypes.GetMessageByIdRequest{
		MessageID: c.GetString("messageId"),
	}
	req.Validate(c.logPoints)
	msg, nextMsgId, prevMsgId, err := messageLogic.GetMessageById(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.MessageWithNextPrev{
		Message:       msg,
		NextMessageId: nextMsgId,
		PrevMessageId: prevMsgId,
	}
}

// GetMessageTypes ...
// swagger:route GET /messages/getMessageTypes message getMessageTypes
//
// GetMessageTypes 获取消息类型/子类型
//
//     Responses:
//       200: getMessageTypes
//       default: ErrorResponse

func (c *MessageController) GetMessageTypes() {
	defer c.CatchException()

	res := messageLogic.GetMessageTypes(c.logPoints)

	c.Res.Result = res
}

// DeleteMessage ...
// swagger:route DELETE /messages/delete message deleteMessage
//
// DeleteMessage 删除消息(可多条)
//
//     Responses:
//       200: deleteMessage
//       default: ErrorResponse
func (c *MessageController) DeleteMessage() {
	defer c.CatchException()

	req := &requestTypes.DeleteMessagesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("DeleteMessage parse ReadMessagesRequest body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	res, err := messageLogic.DeleteMessages(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// DialMail ...
// swagger:route POST /messages/dialMail message dialMail
//
// DialMail 测试邮箱设置是否正确
//
//     Responses:
//       200: dialMail
//       default: ErrorResponse
func (c *MessageController) DialMail() {
	defer c.CatchException()

	req := &requestTypes.MailDialRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("MailDial parse MailDialRequest body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	err := messageLogic.DialMail(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// SaveIsPushMail ...
// swagger:route POST /messages/saveIsPushMail message saveIsPushMail
//
// SaveIsPushMail 保存邮箱是否通知的设置
//
//     Responses:
//       200: saveIsPushMail
//       default: ErrorResponse
func (c *MessageController) SaveIsPushMail() {
	defer c.CatchException()

	req := &requestTypes.SaveIsPushMailRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("SaveMail parse SaveIsPushMailRequest body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	res, err := messageLogic.SaveIsPushMail(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// MailDetail ...
// swagger:route GET /messages/mailDetail message mailDetail
//
// MailDetail 获取邮箱设置
//
//     Responses:
//       200: mailDetail
//       default: ErrorResponse
func (c *MessageController) MailDetail() {
	defer c.CatchException()

	res, err := messageLogic.GetMail(c.logPoints)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}
