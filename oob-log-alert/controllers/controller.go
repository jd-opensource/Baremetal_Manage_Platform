package controllers

import (
	"strings"

	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	"coding.jd.com/aidc-bmp/oob-log-alert/service"

	"github.com/beego/beego/v2/core/logs"
)

// DescribeInstanceMonitorController controller of api describeInstanceMonitor
type DescribeInstanceMonitorController struct {
	BaseController
}

type InstanceStatusReportController struct {
	BaseController
}

type DailyReportController struct {
	BaseController
}

// Get method
func (ctrl *DescribeInstanceMonitorController) Get() {
	logs.Debug(">>>>>>>>DescribeInstanceMonitorController()")
	defer logs.Debug("<<<<<<<<DescribeInstanceMonitorController()")

	instanceID := ctrl.GetString("instanceID")
	logs.Debug("instanceID = %s", instanceID)

	if strings.ToUpper(strings.TrimSpace(instanceID)) == "" {
		ctrl.SetErrorResponse(ParamInvalid, "param instanceID is blank")
		return
	}

	data, err := service.GetInstanceMonitor(instanceID)
	if err != nil {
		ctrl.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}
	ctrl.Res.Result = data

}

// Get method
func (ctrl *InstanceStatusReportController) Get() {
	logs.Debug(">>>>>>>>InstanceStatusReportController()")
	defer logs.Debug("<<<<<<<<InstanceStatusReportController()")

	if err := service.DoStaticCpsPowerState(); err != nil {
		ctrl.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}
	ctrl.Res.Result = dao.CommonResponse{Success: true}
}

// Get method
func (ctrl *DailyReportController) Get() {
	logs.Debug(">>>>>>>>DailyReportController()")
	defer logs.Debug("<<<<<<<<DailyReportController()")

	if err := service.DailReportEvent(); err != nil {
		ctrl.SetErrorResponse(InternalErrorResp, err.Error())
		return
	}
	ctrl.Res.Result = dao.CommonResponse{Success: true}
}
