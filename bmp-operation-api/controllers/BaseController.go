package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	baseLogic "coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"

	UserLoigc "coding.jd.com/aidc-bmp/bmp-operation-api/logic/UserLogic"

	util "coding.jd.com/aidc-bmp/bmp-operation-api/util"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
	log "git.jd.com/cps-golang/log"
	"github.com/beego/beego/v2/server/web"

	beego "github.com/beego/beego/v2/server/web"
)

var sessionName string

func Init() {
	sessionName, _ = beego.AppConfig.String("sessionname")
}

var noQueryResponse = map[string]interface{}{
	"success": true,
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Response struct {
	Result    interface{}    `json:"result,omitempty"`
	Error     *ErrorResponse `json:"error,omitempty"`
	RequestId string         `json:"requestId"`
}

// SubnetController operations for Subnet
type BaseController struct {
	web.Controller
	logPoints *log.Logger
	Res       Response
	pageable  util.Pageable
}

func (b *BaseController) SetErrorResponse(httpCode int, errmsg, errno string) {

	b.Ctx.Output.Status = httpCode
	b.Res.Error = &ErrorResponse{
		Code:    httpCode,
		Message: errmsg,
		Status:  errno,
	}
}

func (b *BaseController) Prepare() {
	web.ReadFromRequest(&b.Controller)
	logPath, _ := beego.AppConfig.String("log.path")
	b.logPoints = log.New(logPath + "/bmp-operation.log")
	b.logPoints.SetStableFields([]string{"method", "uri", "header", "request", "response"})
	b.logPoints.Point("uri", b.Ctx.Request.RequestURI)
	b.logPoints.Point("method", b.Ctx.Request.Method)
	b.logPoints.Point("header", b.Ctx.Request.Header)
	b.logPoints.Point("request", string(b.Ctx.Input.RequestBody))
	b.logPoints.TimeStart("all_t")
	b.logPoints.TimeStart("self_t")
	requestID := b.GetRequestID()

	b.logPoints.Point("logid", requestID)
	b.logPoints.Point("language", b.GetLanguage())
	b.Ctx.Output.Header("X-Jdcloud-Request-Id", requestID)
	b.Res = Response{}
	b.pageable = util.Pageable{
		PageNumber: b.getPageNumber(),
		PageSize:   b.getPageSize(),
	}

	if !strings.HasPrefix(b.Ctx.Input.URI(), "/login") && !strings.HasPrefix(b.Ctx.Input.URI(), "/logout") {
		jdc := b.Ctx.GetCookie(sessionName)
		fmt.Println(time.Now().String(), "prepare func get cookie:", jdc, sessionName)
		b.logPoints.Info(time.Now().Local().String(), "prepare func get cookie:", jdc, sessionName)
		if jdc == "" {
			b.logPoints.Warn("no session_id in cookie, the sessionName is: %s, cookies are: %v", sessionName, b.Ctx.Request.Cookies())
			b.Ctx.ResponseWriter.WriteHeader(402)
			fmt.Fprintln(b.Ctx.ResponseWriter, "user not login")
			return
		}
		//err := b.SetSession(jdc, "admin")
		//fmt.Println(err)
		userName := b.GetSession(jdc)
		fmt.Println(time.Now().String(), "prepare func get session.username:", sessionName, jdc, userName)
		b.logPoints.Info(time.Now().Local().String(), "prepare func get session.username:", sessionName, jdc, userName)
		if userName == nil {
			b.logPoints.Warnf("sessionname %s session_id %s not found or timeout, relogin please...", sessionName, jdc)
			b.Ctx.ResponseWriter.WriteHeader(402)
			fmt.Fprintln(b.Ctx.ResponseWriter, "user not login")
			return
		} else {
			b.Ctx.Output.Header("X-Jdcloud-Username", userName.(string))
			b.logPoints.Point("username", userName.(string))
			user, err := UserLoigc.GetUserByName(b.logPoints, userName.(string))
			if err != nil {
				fmt.Println("prepare.UserLoigc.GetUserByName error:", err.Error())
				b.logPoints.Warnf("prepare.UserLoigc.GetUserByName error, username:%s, error:%s", userName.(string), err.Error())
				b.Ctx.ResponseWriter.WriteHeader(402)
				fmt.Fprintln(b.Ctx.ResponseWriter, "user not login")
				return
			}
			fmt.Println("prepare.userid:", user.UserID)
			b.logPoints.Point("userId", user.UserID)
			if user.Timezone == "" {
				user.Timezone = "Asia/Shanghai"
			}
			b.logPoints.Point("timezone", user.Timezone)
			// b.Ctx.Output.Header("X-Jdcloud-UserID", user.UserID)
			b.logPoints.Infof("get session success, session_id: %s, userid: %s", jdc, userName)
		}
	}
	//b.logPoints.Point("userId", "user-tl84cfjq3m3f83fszhf9j3mqst3m")
	//b.logPoints.Point("username", "zcl")
	//fmt.Println("[DEBUG MINPING]req url:", b.Ctx.Request.URL)
	//fmt.Println("[DEBUG MINPING]req method", b.Ctx.Request.Method)
	//fmt.Println("[DEBUG MINPING]req header", b.Ctx.Request.Header)
	//fmt.Println("[DEBUG MINPING]req body:", string(b.Ctx.Input.RequestBody))

}

func (b *BaseController) Finish() {

	//上述情况为正确情况
	b.Res.RequestId = b.GetRequestID()
	b.Data["json"] = b.Res
	b.ServeJSON()

	b.logPoints.TimeEnd("self_t")
	b.logPoints.TimeEnd("all_t")
	b.logPoints.Point("response", b.Data["json"])
	b.logPoints.Flush()

	v, _ := json.Marshal(b.Data["json"])
	fmt.Println("[DEBUG MINPING]resp body:", string(v))

}

func (b *BaseController) GetString(key string) string {
	return b.Controller.GetString(key, "")
}

func (b *BaseController) GetInt(key string) (int, error) {
	return b.Controller.GetInt(key, 0)
}

//业务代码可以直接panic终止运行,后面放开
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

func (b *BaseController) GetRequestID() string {
	if b.Ctx.Request.Header["X-Jdcloud-Request-Id"] != nil {
		return b.Ctx.Request.Header["X-Jdcloud-Request-Id"][0]
	}
	//生成随机字符串logid
	logid := commonUtil.GenerateRandUuid()
	b.Ctx.Request.Header["X-Jdcloud-Request-Id"] = []string{logid}
	return logid
}

func (b *BaseController) GetErp() string {
	if b.Ctx.Request.Header["X-Jdcloud-Erp"] != nil {
		return b.Ctx.Request.Header["X-Jdcloud-Erp"][0]
	}
	return ""
}

func (b *BaseController) GetJDCookie() string {
	if b.Ctx.Request.Header["X-Jdcloud-Cookie"] != nil {
		return b.Ctx.Request.Header["X-Jdcloud-Host"][0]
	}
	return ""
}

func (b *BaseController) GetJDHost() string {
	if b.Ctx.Request.Header["X-Jdcloud-Host"] != nil {
		return b.Ctx.Request.Header["X-Jdcloud-Host"][0]
	}
	return ""
}

// v1/validation success res: {"requestId":"0wzsgfzm-veo8-rjbm-nf53-0gqparr6b9di","error":null,"result":{"loginInfo":{"type":1,"pin":"4cb497f05ae7455c9272","adminPin":null,"credentialInfo":null,"loginUrl":"","loginName":"zyTest"}}}

// failed res:{"requestId":"bygqoepg-k3n1-k8rc-4l52-iho1nr3brmtp","error":null,"result":{"loginInfo":{"type":0,"pin":null,"adminPin":null,"credentialInfo":null,"loginUrl":"http://login-stag.jdcloud.com","loginName":null}}}

type LoginstateValidationRes struct {
	RequestId string `json:"requestId"`
	// Error     string `json:"error"`
	Result struct {
		LoginInfo LoginInfo `json:"loginInfo"`
	} `json:"result"`
}

type LoginInfo struct {
	Type           int             `json:"type"` //0表示未登录，1表示主账号登录，2表示子账号登录，3表示角色登录
	Pin            string          `json:"pin"`
	AdminPin       string          `json:"adminPin"`
	CredentialInfo *CredentialInfo `json:"credentialInfo"`
	LoginUrl       string          `json:"loginUrl"`
	LoginName      string          `json:"loginName"`
}

type CredentialInfo struct {
	AccessKey    string `json:"accessKey"`
	SecretKey    string `json:"secretKey"`
	SessionToken string `json:"sessionToken"`
	Expiration   string `json:"expiration"` //类型是data还是string？返回结果此字段都没有数据，先定义成string
	RolePin      string `json:"rolePin"`
}

func (b *BaseController) GetLanguage() string {
	fmt.Println("语言header", b.Ctx.Request.Header["X-Jdcloud-Language"])
	fmt.Println("语言cookie", b.Ctx.GetCookie("X-Jdcloud-Language"), b.Ctx.GetCookie("path_url"))
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

func (b *BaseController) GetHost() string {
	return b.Ctx.Request.Host
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
		if v, err := strconv.Atoi(l); err == nil {
			return int64(v)
		}
	}
	return 1
}

func (b *BaseController) getPageSize() int64 {
	if l := b.GetString("pageSize"); l != "" {
		if v, err := strconv.Atoi(l); err == nil {
			return int64(v)
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
