package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/CustomInfoLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"git.jd.com/cps-golang/ironic-common/exception"
)

// CustomController operations for custom info
type CustomController struct {
	BaseController
}

// SetCustomInfo ...
// swagger:route POST /custom/setCustomInfo custom setCustomInfo
// SetCustomInfo 设置自定义表头
//     Responses:
//       200: setCustomInfo
//       default: ErrorResponse

func (c *CustomController) SetCustomInfo() {
	defer c.CatchException()
	req := &requestTypes.SetCustomInfoRequest{}
	// req.Validate()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		panic(exception.Exception{Msg: err.Error()})
	}
	res := CustomInfoLogic.SetCustomInfo(c.logPoints, *req)
	c.Res.Result = res
}

// GetCustomInfo ...
// swagger:route GET /custom/getCustomInfo custom getCustomInfo
// GetCustomInfo 获取自定义表头
//     Responses:
//       200: getCustomInfo
//       default: ErrorResponse

func (c *CustomController) GetCustomInfo() {
	defer c.CatchException()
	req := &requestTypes.QueryCustomInfoRequest{
		PageKey: c.GetString("pageKey"),
		Reload:  c.GetString("reload"),
	}
	// req.Validate()
	res := CustomInfoLogic.GetCustomInfo(c.logPoints, *req)
	c.Res.Result = res
}
