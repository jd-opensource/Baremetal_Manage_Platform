package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	util "coding.jd.com/bmp/agent-proxy-jdstack/util"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
	log "git.jd.com/cps-golang/log"
	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	logPoints *log.Logger
	Res       Response
	pageable  util.Pageable
}

type ErrorResponse struct {
	// 错误码
	Code int `json:"code"`
	// 错误信息
	Message string `json:"message"`
	// 错误状态
	Status string `json:"status"`
}

type Response struct {
	// 操作失败结果。成功时有此结构
	Result interface{} `json:"result,omitempty"`
	// 操作失败结果。失败时有此结构
	Error *ErrorResponse `json:"error,omitempty"`
	// 请求traceId
	// required: true
	RequestId string `json:"requestId"`
}

func (b *BaseController) SetErrorResponse(code int, message, status string) {
	b.Res.Error = &ErrorResponse{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

func (b *BaseController) Prepare() {
	web.ReadFromRequest(&b.Controller)
	logPath, _ := web.AppConfig.String("log.path")
	b.logPoints = log.New(logPath + "/monitor-proxy.log")
	b.logPoints.SetStableFields([]string{"method", "uri", "header", "request", "response"})
	b.logPoints.Point("uri", b.Ctx.Request.RequestURI)
	b.logPoints.Point("method", b.Ctx.Request.Method)
	b.logPoints.Point("header", b.Ctx.Request.Header)
	b.logPoints.Point("language", b.GetLanguage())

	b.logPoints.Point("request", string(b.Ctx.Input.RequestBody))
	b.logPoints.TimeStart("all_t")
	b.logPoints.TimeStart("self_t")
	traceId := b.GetTraceID()

	b.logPoints.Point("logid", traceId)
	b.Ctx.Output.Header("TraceId", traceId)
	b.Res = Response{}

	b.pageable = util.Pageable{
		PageNumber: b.getPageNumber(),
		PageSize:   b.getPageSize(),
	}
	b.logPoints.Info("url:", b.Ctx.Request.URL)
	b.logPoints.Info("body:", string(b.Ctx.Input.RequestBody))
	fmt.Println("[DEBUG MINPING]url:", b.Ctx.Request.URL)
	fmt.Println("[DEBUG MINPING]method", b.Ctx.Request.Method)
	fmt.Println("[DEBUG MINPING]header", b.Ctx.Request.Header)
	fmt.Println("[DEBUG MINPING]body:", string(b.Ctx.Input.RequestBody))
	fmt.Println("[DEBUG MINPING]all header", b.GetAllHeader())
}

func (b *BaseController) Finish() {

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
	logid := commonUtil.GenerateRandUuid()
	b.Ctx.Request.Header["X-Jdcloud-Request-Id"] = []string{logid}
	return logid
}

func (b *BaseController) GetLanguage() string {
	if b.Ctx.Request.Header["Bmplanguage"] != nil {
		language := b.Ctx.Request.Header["Bmplanguage"][0]
		return language
	}
	return ""
}

func (b *BaseController) GetTraceID() string {
	if b.Ctx.Request.Header["Traceid"] != nil {
		return b.Ctx.Request.Header["Traceid"][0]
	}
	//生成随机字符串logid
	logid := commonUtil.GenerateRandUuid()
	b.Ctx.Request.Header["Traceid"] = []string{logid}
	return logid
}

func (b *BaseController) getPageNumber() int64 {
	if l := b.GetString("pageNumber"); l != "" {
		if v, err := strconv.ParseInt(l, 10, 64); err == nil {
			return v
		}
	}
	return 1
}

func (b *BaseController) getPageSize() int64 {
	if l := b.GetString("pageSize"); l != "" {
		if v, err := strconv.ParseInt(l, 10, 64); err == nil {
			return v
		}
	}
	return 20
}

func (b *BaseController) GetAllHeader() string {
	v, _ := json.Marshal(b.Ctx.Request.Header)
	return string(v)
}
