package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	"coding.jd.com/aidc-bmp/oob-log-alert/service"
)

// CustomController operations for custom info
type CustomController struct {
	BaseController
}

// SetCustomInfo ...
// swagger:route POST /custom/set-custom-info custom setCustomInfo
// SetCustomInfo 设置自定义信息
//     Responses:
//       200: setCustomInfo
//       default: ErrorResponse

func (c *CustomController) SetCustomInfo() {

	req := dao.SetCustomInfoRequest{}
	// req.Validate()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.logPoints.Warnf("SetCustomInfo param invalid, body:%s, error:%s", string(c.Ctx.Input.RequestBody), err.Error())
		c.SetErrorResponse(ParamInvalid, err.Error())
	}
	res := service.SetCustomInfo(c.logPoints, req)
	c.Res.Result = res
}

// GetCustomInfo ...
// swagger:route GET /custom/get-custom-info custom getCustomInfo
// GetCustomInfo 获取自定义信息
//     Responses:
//       200: getCustomInfo
//       default: ErrorResponse

func (c *CustomController) GetCustomInfo() {

	req := dao.QueryCustomInfoRequest{
		PageKey: c.GetString("page_key"),
		Reload:  c.GetString("reload"),
	}
	// req.Validate()
	res := service.GetCustomInfo(c.logPoints, req)
	c.Res.Result = res
}
