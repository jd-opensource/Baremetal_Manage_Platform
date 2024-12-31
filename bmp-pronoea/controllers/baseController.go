package controllers

import (
	"encoding/json"
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/models/alert"
	"git.jd.com/bmp-pronoea/models/rules"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

const (
	/*
	API_OK                    = 10000
	API_ARGUMENTS_ERROR       = 20040
	API_NO_ARGUMENTS          = 20041
	API_UNKOWN_ERROR          = 20050
	API_INTERNAL_ERROR        = 20051

	 */

	HTTP_STATUS_SUCCESS = 200

	HTTP_STATUS_BAD_REQUEST = 400
	HTTP_STATUS_FORBIDDEN   = 403
	HTTP_STATUS_NOT_FOUND   = 404
	HTTP_STATUS_INTERNAL    = 500

	DOWNLOAD = "download"
)

type BaseController struct {
	beego.Controller
	resp HttpCommonResponse
}

type HttpCommonResponse struct {
	Code      int
	RequestId string
	Message   string
	Result    interface{}
}

func (b *BaseController) GetHeader() map[string]string {
	return map[string]string{
		"requestId": b.Ctx.Input.Header("X-Request-Id"),
	}
}

func (b *BaseController) Prepare() {
	requestId, ok := b.GetHeader()["requestId"]
	if !ok || requestId == "" {
		if b.GetRequestId() == "" {
			requestId = util.GetUuid()
		} else {
			requestId = b.GetRequestId()
		}
	}
	b.Ctx.Input.SetData("requestId", requestId)
}

func (b *BaseController) Finish() {

	res := b.Ctx.Input.GetData(DOWNLOAD)
	if res == nil {
		b.Ctx.Output.SetStatus(b.resp.Code)
		b.Data["json"] = b.resp
		b.ServeJSON()

		b, _ := json.Marshal(b.resp)
		logs.Info(fmt.Sprintf("response : %v", string(b)))
	}
}

func (b *BaseController) GetRequestId() string {
	if b.Ctx.Input.GetData("requestId") == nil {
		return ""
	} else {
		return b.Ctx.Input.GetData("requestId").(string)
	}
}

func (b *BaseController) GetAlertModel(receiverType string) (ruleModel alert.IAlert, err error) {
	err = nil

	switch receiverType {
	case alert.RECEIVER_MODELS_BMP:
		ruleModel = BmpAlertManagerModels
		break
	default:
		err = fmt.Errorf("unknown receiver type:%v ", receiverType)
	}

	return ruleModel, err
}

func (b *BaseController) GetRuleModel(ruleType string) (ruleModel rules.IRule, err error) {
	err = nil

	switch ruleType {
	case rules.RULE_MODELS_BMPINBAND:
		ruleModel = BmpInBandRuleModeles
		break
	default:
		err = fmt.Errorf("unknown rule type:%v ", ruleType)
	}

	return ruleModel, err
}

