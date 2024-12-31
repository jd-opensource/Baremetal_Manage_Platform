package controllers

import (
	"encoding/json"
	"fmt"

	log "coding.jd.com/aidc-bmp/bmp_log"

	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	"coding.jd.com/aidc-bmp/oob-log-alert/service"
	"coding.jd.com/aidc-bmp/oob-log-alert/util"

	"github.com/beego/beego/v2/server/web"
)

const (
	InternalErrorResp int = 500
	ParamInvalid      int = 400
	SuccessResp       int = 200
)

type BaseResponse struct {
	Result    interface{}    `json:"result,omitempty"`
	Error     *ErrorResponse `json:"error,omitempty"`
	RequestId string         `json:"requestId"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseController struct {
	web.Controller
	Res       BaseResponse
	logPoints *log.Logger
}

func (b *BaseController) SetErrorResponse(code int, message string) {
	b.Res.Error = &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

func (ctrl *BaseController) GetPageReqInfo() dao.PageRequest {

	pageNum, _ := ctrl.GetInt64("page_num")
	pageSize, _ := ctrl.GetInt64("page_size")
	if ctrl.GetString("isAll") != "" {
		pageSize = 0 //无限制
	}
	return dao.PageRequest{
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}

func (b *BaseController) GetTraceID() string {
	if b.Ctx.Request.Header["Traceid"] != nil {
		return b.Ctx.Request.Header["Traceid"][0]
	}
	//生成随机字符串logid
	logid := util.GenerateRandUuid()
	b.Ctx.Request.Header["Traceid"] = []string{logid}
	return logid
}

func (b *BaseController) Prepare() {

	web.ReadFromRequest(&b.Controller)
	logPath, _ := web.AppConfig.String("log.path")
	b.logPoints = log.New(logPath + "/oobalert-api.log")
	b.logPoints.SetStableFields([]string{"method", "uri", "header", "request", "response"})
	b.logPoints.Point("uri", b.Ctx.Request.RequestURI)
	b.logPoints.Point("method", b.Ctx.Request.Method)
	b.logPoints.Point("header", b.Ctx.Request.Header)
	b.logPoints.Point("request", string(b.Ctx.Input.RequestBody))
	b.logPoints.TimeStart("all_t")
	b.logPoints.TimeStart("self_t")
	traceId := b.GetTraceID()

	b.logPoints.Point("logid", traceId)
	b.SetUserInfo()
	b.logPoints.Point("language", b.GetLanguage())

	b.Ctx.Output.Header("TraceId", traceId)
	b.Res = BaseResponse{RequestId: traceId}

}
func (b *BaseController) Finish() {

	// b.Data["json"] = b.Res
	// if b.Res.Error != nil {
	// 	b.Ctx.Output.Status = b.Res.Error.Code
	// }

	// v, _ := json.Marshal(b.Data["json"])
	// logs.Debug("response", string(v))

	// b.ServeJSON()
	//

	if b.GetString("exportType") != "" { //test
		return
	}

	b.Res.RequestId = b.GetTraceID()
	b.Data["json"] = b.Res

	if b.Res.Error != nil {
		b.Ctx.Output.Status = b.Res.Error.Code
	}
	b.ServeJSON()

	b.logPoints.TimeEnd("self_t")
	b.logPoints.TimeEnd("all_t")
	b.logPoints.Point("response", b.Data["json"])

	v, _ := json.Marshal(b.Data["json"])
	b.logPoints.Info("response:", string(v))
	fmt.Println("[DEBUG MINPING]response:", string(v))

	b.logPoints.Flush()

}

func (b *BaseController) GetRequestID() string {
	if b.Ctx.Request.Header["X-Jdcloud-Request-Id"] != nil {
		return b.Ctx.Request.Header["X-Jdcloud-Request-Id"][0]
	}
	//生成随机字符串logid
	logid := util.GenerateRandUuid()
	b.Ctx.Request.Header["X-Jdcloud-Request-Id"] = []string{logid}
	return logid
}

func (b *BaseController) GetLanguage() string {
	return b.Ctx.GetCookie("X-Jdcloud-Language")
	//if len(b.Ctx.Request.Header["X-Jdcloud-Language"]) == 0 {
	//	if b.Ctx.GetCookie("X-Jdcloud-Language") == "" {
	//		return ""
	//	} else {
	//		return b.Ctx.GetCookie("X-Jdcloud-Language")
	//	}
	//}
	//return b.Ctx.Request.Header["X-Jdcloud-Language"][0]
}

//设置username,userid等到logger
func (b *BaseController) SetUserInfo() {

	userName := b.Ctx.GetCookie("X-Jdcloud-Username")
	if userName == "" {
		b.logPoints.Warnf("username empty, check it!!!")
		return
	}
	b.logPoints.Point("username", userName)
	userId, _ := service.GetUserIdByName(b.logPoints, userName)
	if userId == "" {
		return
	}
	b.logPoints.Point("userId", userId)
}
