package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/apikeyLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/userLogic"

	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"

	util "coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
	"github.com/beego/beego/v2/server/web"
)

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

// SubnetController operations for Subnet
type BaseController struct {
	web.Controller
	logPoints *log.Logger
	Res       Response
	pageable  util.Pageable
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
	b.logPoints = log.New(logPath + "/bmp-openapi.log")
	b.logPoints.SetStableFields([]string{"method", "uri", "header", "request", "response"})
	b.logPoints.Point("uri", b.Ctx.Request.RequestURI)
	b.logPoints.Point("method", b.Ctx.Request.Method)
	b.logPoints.Point("header", b.Ctx.Request.Header)
	b.logPoints.Point("language", b.GetLanguage())
	b.logPoints.Point("timezone", b.GetTimezone())

	b.logPoints.Point("request", string(b.Ctx.Input.RequestBody))
	b.logPoints.TimeStart("all_t")
	b.logPoints.TimeStart("self_t")
	traceId := b.GetTraceID()

	b.logPoints.Point("logid", traceId)
	b.Ctx.Output.Header("TraceId", traceId)
	b.Res = Response{}

	//鉴定用户/系统apikey,补全流量的用户信息，操作权限
	if !strings.HasPrefix(b.Ctx.Input.URI(), "/v1/users/verify") && !strings.HasPrefix(b.Ctx.Input.URI(), "/v1/users/getUserByName") && !strings.HasPrefix(b.Ctx.Input.URI(), "/commands/retryCommand") && !strings.HasPrefix(b.Ctx.Input.URI(), "/v1/license/currentVersion") && !strings.HasPrefix(b.Ctx.Input.URI(), "/v1/offline") && !strings.HasPrefix(b.Ctx.Input.URI(), "/v1/innerTest") {

		err := b.ValidateAuth()
		if err != nil {
			b.logPoints.Warn("ValidateLicense error:", err.Messagech)
		} else if strings.HasPrefix(b.Ctx.Input.URI(), "/v1/monitor") {
			// err = licenseLogic.ValidateLicense(b.logPoints, b.Ctx.Input.URI())
			// if err != nil {
			// 	b.logPoints.Warn("ValidateLicense error:", err.Messagech)
			// }

		}

		if err != nil {
			b.Res.Error = &ErrorResponse{
				Code:    err.Code,
				Message: err.Messagech,
				Status:  err.Status,
			}
			if b.logPoints.GetPoint("language").(string) == baseLogic.EN_US {
				b.Res.Error.Message = err.MessageEn
			}

			b.Res.RequestId = traceId
			b.Data["json"] = b.Res
			if b.Res.Error != nil {
				b.Ctx.Output.Status = b.Res.Error.Code
			}
			b.ServeJSON()
			return
		}

	}

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

func (b *BaseController) GetString(key string) string {
	return b.Controller.GetString(key, "")
}

//业务代码可以直接panic终止运行
func (b *BaseController) CatchException() {
	// if r := recover(); r != nil {
	// 	// t := make([]byte, 1<<16)
	// 	// runtime.Stack(t, true)
	// 	t, _ := json.Marshal(r)
	// 	b.logPoints.Warn(string(t))
	// 	if reflect.TypeOf(r).String() == "exception.Exception" {
	// 		b.Res["result"] = r
	// 	} else { //非预期panic输出统一化
	// 		b.Res["result"] = exception.InternalError
	// 	}
	// }
	if r := recover(); r != nil {
		t, _ := json.Marshal(r)
		b.logPoints.Warn(string(t))
		//b.Res["result"] = int(reflect.ValueOf(r).Field(0).Int())
		//if reflect.ValueOf(r).Field(1).String() == "" { //重要返回，如果errno不写，那么默认的errno和msg在这里设置
		//	b.Res["error"] = 40000
		//	b.Res["m"] = reflect.ValueOf(r).Field(2).String()
		//	return
		//}
		//errno, err := strconv.Atoi(reflect.ValueOf(r).Field(1).String()) //如果errno不写
		//if err != nil {                                                  //字符串转整形失败
		//	b.Res["errno"] = 40001
		//	b.Res["errmsg"] = reflect.ValueOf(r).Field(2).String()
		//	return
		//}
		var (
			code      int
			chMessage string
			enMessage string
			status    string
		)
		if reflect.ValueOf(r).Field(0).Int() == 0 {
			code = httpStatus.BAD_REQUEST //默认httpcode 400
		} else {
			code = int(reflect.ValueOf(r).Field(0).Int())
		}
		//en_message
		if reflect.ValueOf(r).Field(1).String() == "" {
			enMessage = errorCode.INVALID_ARGUMENT //参数错误
		} else {
			enMessage = reflect.ValueOf(r).Field(1).String()
		}

		//ch_message
		if reflect.ValueOf(r).Field(2).String() == "" {
			chMessage = errorCode.INVALID_ARGUMENT //参数错误
		} else {
			chMessage = reflect.ValueOf(r).Field(2).String()
		}

		//status
		if reflect.ValueOf(r).Field(3).String() == "" {
			status = errorCode.INVALID_ARGUMENT //参数错误
		} else {
			status = reflect.ValueOf(r).Field(3).String() //具体的报错信息err.error()
		}
		fmt.Println("recover info:", code, chMessage, enMessage, status)
		if b.logPoints.GetPoint("language").(string) == baseLogic.EN_US {
			b.SetErrorResponse(code, enMessage, status)
		} else {
			b.SetErrorResponse(code, chMessage, status)
		}

	}
}

func (b *BaseController) GetAllHeader() string {
	v, _ := json.Marshal(b.Ctx.Request.Header)
	return string(v)
}

func (b *BaseController) ValidateAuth() *constant.RespMsg {

	auth := ""
	if b.Ctx.Request.Header["Authorization"] != nil {
		auth = b.Ctx.Request.Header["Authorization"][0]
	}
	items := strings.Split(auth, "Bearer ")
	if len(items) != 2 {
		b.logPoints.Warnf("ValidateAuth.token type invalided, token:%s", auth)
		err := constant.AUTH_BASIC_AUTH_ERROR
		return &err
	}
	token := items[1]
	apiKeyEntity, err := apikeyLogic.ValidateApiKey(b.logPoints, token)
	if err != nil {
		b.logPoints.Warnf("ValidateAuth.ValidateApiKey error, auth:%s, error:%s", auth, err.Error())
		err := constant.AUTH_BASIC_AUTH_ERROR
		return &err
	}
	if apiKeyEntity == nil {
		b.logPoints.Warnf("ValidateAuth.apiKeyEntity empty,auth:%s", auth)
		err := constant.AUTH_BASIC_AUTH_ERROR
		return &err
	}

	// 只读的apikey不允许做非get操作
	if apiKeyEntity.ReadOnly == 1 && b.logPoints.GetPoint("method").(string) != "GET" {
		b.logPoints.Warn("readonly apikey try do modify operation", token, b.Ctx.Input.Method(), b.Ctx.Input.URI())
		err := constant.PermissionDenyForObject
		return &err
	}

	//如果是个人apikey,可以不传userId
	var userId string
	if apiKeyEntity.Type == "user" {
		b.logPoints.Point("traffic_source", "api")
		if apiKeyEntity.UserID != b.GetUserId() {
			b.logPoints.Warn("userid in auth and header not match")
			// err := constant.AUTH_ERROR_FOR_AUTH_AND_USERID
			// return &err
		}
		userId = apiKeyEntity.UserID
	} else { //系统apikey时,userId必须要经header传过来
		b.logPoints.Point("traffic_source", "web")
		userId = b.GetUserId()
	}
	if userId == "" {
		err := constant.AUTH_NOT_LOGIN
		return &err
	}
	u, err := userLogic.GetUserById(b.logPoints, userId)
	if err != nil {
		b.logPoints.Warn("apikey for user invalid", err.Error())
		err := constant.AUTH_NOT_LOGIN
		return &err
	}

	if apiKeyEntity.Type == "user" && u.RoleID == baseLogic.ROLE_USER_UUID {
		//控制台创建的个人apikey不能访问运营平台的openapi接口
		err := constant.PermissionDenyForObject
		return &err
	}

	if apiKeyEntity.Type == "user" && apiKeyEntity.Source != "operation" {
		//个人apikey且不是在operation上创建的
		err := constant.PermissionDenyForObject
		return &err
	}

	//接口请求的控制台用户，镜像，机型等接口的写权限禁用
	if apiKeyEntity.Type == "user" && u.RoleID == baseLogic.ROLE_USER_UUID && b.logPoints.GetPoint("method").(string) != "GET" {
		forbidUrl := []string{
			"/v1/devices",
			"/v1/deviceTypes",
			"v1/idcs",
			"/v1/oss",
			"/v1/images",
			"/v1/raids",
			"/v1/users",
			"/v1/roles",
		}
		for _, pre := range forbidUrl {
			if strings.HasPrefix(b.Ctx.Input.URI(), pre) {
				b.logPoints.Warnf("api user request permission denied, userId:%s, url:%s", u.UserID, b.Ctx.Input.URI())
				err := constant.PermissionDenyForObject
				return &err
			}
		}
	}

	// 非amdin无运营平台role和users的权限
	if u.RoleID != baseLogic.ROLE_ADMIN_UUID {
		if strings.HasPrefix(b.Ctx.Input.URI(), "/v1/roles") {
			if !strings.HasPrefix(b.Ctx.Input.URI(), "/v1/roles/roleInfo/current") {
				b.logPoints.Warn("role method for non-admin user invalid, userId:", u.UserID)
				err := constant.PermissionDenyForObject
				return &err
			}

		}

		if strings.HasPrefix(b.Ctx.Input.URI(), "/v1/users") {
			b.logPoints.Warn("role method for non-admin user invalid, userId:", u.UserID)
			err := constant.PermissionDenyForObject
			return &err
		}
	}

	fmt.Println("baseauth.username is", u.UserName)
	fmt.Println("baseauth.user_id is", userId)
	fmt.Println("baseauth.timezone is", u.Timezone)
	fmt.Println("baseauth.language is", u.Language)
	fmt.Println("language-header", b.logPoints.GetPoint("language").(string))
	b.logPoints.Point("userId", userId)
	b.logPoints.Point("username", u.UserName)
	b.logPoints.Point("userRole", u.RoleID)
	if b.logPoints.GetPoint("language").(string) == "" { //页面上没有设置语言时，用用户的默认语言
		b.logPoints.Point("language", u.Language)
	}
	fmt.Println("language-final", b.logPoints.GetPoint("language").(string))
	b.logPoints.Point("timezone", u.Timezone)

	return nil
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

func (b *BaseController) GetLanguage() string {
	if b.Ctx.Request.Header["Bmplanguage"] != nil {
		language := b.Ctx.Request.Header["Bmplanguage"][0]
		return language
	}
	return ""
}

func (b *BaseController) GetTimezone() string {
	if b.Ctx.Request.Header["X-Bmp-Timezone"] != nil {
		tz := b.Ctx.Request.Header["X-Bmp-Timezone"][0]
		return tz
	}
	return "Asia/Shanghai"
}

func (b *BaseController) GetUserId() string {
	if b.Ctx.Request.Header["Bmpuserid"] != nil {
		userId := b.Ctx.Request.Header["Bmpuserid"][0]
		return userId
	}
	return ""
}
func (b *BaseController) GetUserName() string {
	user, _ := userLogic.GetUserById(b.logPoints, b.GetUserId()) //报错
	if user != nil {
		return user.UserName
	} else {
		return ""
	}
}
func (b *BaseController) GetFields() (fields []string) {
	// fields: col1,col2,entity.col3
	if v := b.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	return
}

func (b *BaseController) GetSortBy() (sortby []string) {
	// sortby: col1,col2
	if v := b.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	return
}

func (b *BaseController) GetOrder() (order []string) {
	// order: desc,asc
	if v := b.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	return
}

// GetOffset 默认返回0
func (b *BaseController) GetOffset() int64 {
	if o := b.GetString("pageNumber"); o != "" {
		if v, err := strconv.ParseInt(o, 10, 64); err == nil {
			//第2页 每页10条 offset=10
			//第1页， 每页10条，offset=0
			if v > 0 {
				return (v - 1) * b.GetLimit()
			}
		}
	}
	return 0
}

// GetOffset 默认返回10
func (b *BaseController) GetLimit() int64 {
	if l := b.GetString("pageSize"); l != "" {
		if v, err := strconv.ParseInt(l, 10, 64); err == nil {
			return v
		}
	}
	return 20
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
	if strings.HasPrefix(b.Ctx.Request.RequestURI, "/devices") {
		return 10
	}
	if strings.HasPrefix(b.Ctx.Request.RequestURI, "/instances") {
		return 10
	}
	if strings.HasPrefix(b.Ctx.Request.RequestURI, "/subnets") {
		return 10
	}
	return 20
}

func restfulResponseException(code int, status string, msg string) map[string]interface{} {
	return map[string]interface{}{
		"error": map[string]interface{}{
			"code":    code,
			"status":  status,
			"message": msg,
		},
	}
}
