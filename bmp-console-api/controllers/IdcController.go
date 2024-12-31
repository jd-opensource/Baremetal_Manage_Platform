package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/logic/IdcLogic"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// IdcController operations for idc
type IdcController struct {
	BaseController
}

// swagger:route GET /idc idc getIdcList
// GetIdcList 获取idc列表
//     Responses:
//       200: getIdcList
//       default: ErrorResponse

func (c *IdcController) GetIdcList() {
	defer c.CatchException()

	// req.Validate(c.logPoints)
	res, err := IdcLogic.GetIdcList(c.logPoints)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	if c.GetString("exportType") != "" {
		fileName, downloadFileName := IdcLogic.ExportIdcList(c.logPoints, res.Idcs)
		//c.Ctx.Input.SetData("download", "yes")
		c.Ctx.Output.Download(fileName, downloadFileName)
		return
	} else {
		c.Res.Result = res
	}
}
