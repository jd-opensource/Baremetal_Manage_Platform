package controllers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	"coding.jd.com/aidc-bmp/oob-log-alert/service"

	"github.com/beego/beego/v2/core/logs"
)

// LogController controller of api log
type LogController struct {
	BaseController
}

// swagger:route GET /v1/oob-alert/logs ooblog getLogCollections
//
// GetLogCollections 故障报警日志-运营平台
//
//     Responses:
//       200: getLogCollections
//       default: ErrorResponse
func (ctrl *LogController) GetLogCollections() {

	logs.Debug(">>>>>>>>LogController.GetLogs()")
	defer logs.Debug("<<<<<<<<LogController.GetLogs()")

	req := dao.GetLogCollectionsRequest{
		IdcId:       ctrl.GetString("idc_id"),
		Sn:          ctrl.GetString("sn"),
		Level:       ctrl.GetString("level"),
		Spec:        ctrl.GetString("spec"),
		DealStatus:  ctrl.GetString("deal_status"),
		StartTime:   ctrl.GetString("start_time"),
		EndTime:     ctrl.GetString("end_time"),
		PageRequest: ctrl.GetPageReqInfo(),
	}

	if req.IdcId != "" {
		matched, err := regexp.MatchString(`^[a-zA-Z-0-9]{1,}$`, req.IdcId)
		if err != nil {
			ctrl.SetErrorResponse(ParamInvalid, err.Error())
			return
		}
		if !matched {
			ctrl.SetErrorResponse(ParamInvalid, "idc invalid")
			return
		}
	}

	if req.DealStatus != "" && req.DealStatus != "0" && req.DealStatus != "1" {
		ctrl.SetErrorResponse(ParamInvalid, "deal_status invalid")
		return
	}

	if req.Level != "" {
		matched, err := regexp.MatchString(`^[a-zA-Z]{1,}$`, req.Level)
		if err != nil {
			ctrl.SetErrorResponse(ParamInvalid, err.Error())
			return
		}
		if !matched {
			ctrl.SetErrorResponse(ParamInvalid, "level invalid")
			return
		}
	}

	if req.Sn != "" {
		matched, err := regexp.MatchString(`^[0-9a-zA-Z-_]{1,}$`, req.Sn)
		if err != nil {
			ctrl.SetErrorResponse(ParamInvalid, err.Error())
			return
		}
		if !matched {
			ctrl.SetErrorResponse(ParamInvalid, "sn invalid")
			return
		}
	}

	logs.Debug("GetLogs.req = %+v", req)

	data, err := service.GetLogCollections(ctrl.logPoints, req)
	if err != nil {
		ctrl.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}

	if ctrl.GetString("exportType") != "" {
		fileName, downloadFileName, err := service.ExportLogCollectionsExcel(ctrl.logPoints, data.Detail)
		if err != nil {
			ctrl.SetErrorResponse(InternalErrorResp, "导出故障报警日志列表excel错误 "+err.Error())
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

// swagger:route GET /v1/oob-alert/consolelogs ooblog getLogCollectionsBySn
//
// GetLogCollectionsBySn 故障报警日志-控制台(只有符合push状态规则的才在控制台展示)
//
//     Responses:
//       200: getLogCollectionsBySn
//       default: ErrorResponse
func (ctrl *LogController) GetLogCollectionsBySn() {

	logs.Debug(">>>>>>>>LogController.GetLogCollectionsBySn()")
	defer logs.Debug("<<<<<<<<LogController.GetLogCollectionsBySn()")

	req := dao.GetLogCollectionsBySnRequest{
		Sn:          ctrl.GetString("sn"),
		PageRequest: ctrl.GetPageReqInfo(),
	}

	logs.Debug("GetLogsBySn.req = %+v", req)

	data, err := service.GetLogCollectionsBySn(ctrl.logPoints, req)
	if err != nil {
		ctrl.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}

	if ctrl.GetString("exportType") != "" {
		fileName, downloadFileName, err := service.ExportLogCollectionBySnExcel(ctrl.logPoints, data.Detail)
		if err != nil {
			ctrl.SetErrorResponse(InternalErrorResp, "导出故障报警日志列表excel错误 "+err.Error())
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

// swagger:route POST /v1/oob-alert/logs/deal ooblog dealLogCollection
//
// DealLogCollection 故障报警日志-处理
//
//     Responses:
//       200: dealLogCollection
//       default: ErrorResponse
func (ctrl *LogController) DealLogCollection() {
	logs.Debug(">>>>>>>>LogController.DealLogCollection()")
	defer logs.Debug("<<<<<<<<LogController.DealLogCollection()")
	req := dao.DealLogCollectionRequest{}
	fmt.Println("DealLogCollection req.body", string(ctrl.Ctx.Input.RequestBody))

	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req); err != nil {
		ctrl.SetErrorResponse(ParamInvalid, err.Error())
		return
	}

	logs.Debug("GetLogsBySn.req = %+v", req)

	err := service.DealLogCollection(ctrl.logPoints, req)
	if err != nil {
		ctrl.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}
	ctrl.Res.Result = dao.CommonResponse{
		Success: true,
	}

}
