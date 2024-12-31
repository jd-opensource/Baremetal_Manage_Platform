package controllers

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	"coding.jd.com/aidc-bmp/oob-log-alert/service"

	"github.com/beego/beego/v2/core/logs"
)

// DeviceController controller of api device
type DeviceController struct {
	BaseController
}

// swagger:route GET /v1/oob-alert/device/status device getDeviceStatus
//
// GetDeviceStatus 获取设备带外总体状态
//
//     Responses:
//       200: getDeviceStatus
//       default: ErrorResponse
func (ctrl *DeviceController) GetDeviceStatus() {
	logs.Debug(">>>>>>>>DeviceController.GetDeviceStatus()")
	defer logs.Debug("<<<<<<<<DeviceController.GetDeviceStatus()")

	req := dao.GetDeviceStatusRequest{
		IdcId:       ctrl.GetString("idc_id"),
		Sn:          ctrl.GetString("sn"),
		PageRequest: ctrl.GetPageReqInfo(),
	}
	logs.Debug("GetDeviceStatus.req = %+v", req)

	data, err := service.GetDeviceStatus(ctrl.logPoints, req)
	if err != nil {
		ctrl.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}

	if ctrl.GetString("exportType") != "" {
		fileName, downloadFileName, err := service.ExportDeviceStatusExcel(ctrl.logPoints, data.Detail)
		if err != nil {
			ctrl.SetErrorResponse(InternalErrorResp, "导出设备带外列表excel错误 "+err.Error())
			return
		}
		//c.Ctx.Input.SetData("download", "yes")
		ctrl.Ctx.Output.Download(fileName, downloadFileName)
		// if err := os.Remove(fileName); err != nil {
		// 	ctrl.SetErrorResponse(InternalErrorResp, "删除上传文件失败"+err.Error())
		// 	return
		// }
		return
	}

	ctrl.Res.Result = data

}
