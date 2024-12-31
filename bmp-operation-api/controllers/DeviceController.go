package controllers

import (
	// "encoding/json"

	"encoding/json"
	"os"

	baseLogic "coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"

	DeviceLogic "coding.jd.com/aidc-bmp/bmp-operation-api/logic/DeviceLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

type DeviceController struct {
	BaseController
}

func (c *DeviceController) UploadDevices() {
	defer c.CatchException()
	file, header, err := c.GetFile("deviceFile")
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	defer file.Close()
	req := &requestTypes.UploadDeviceRequest{
		DeviceFile: file,
	}
	res, err := DeviceLogic.UploadDevices(c.logPoints, req, header, c.GetRequestID()) //返回byte数组
	if err != nil {
		errmsg := "文件不合法"
		if c.logPoints.GetPoint("language") == baseLogic.EN_US {
			errmsg = "file invalidate"
		}
		c.logPoints.Warn("上传文件出错", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, errmsg, errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// swagger:route POST /devices device createDevices
//     Responses:
//       200: createDevices
//       default: ErrorResponse
func (c *DeviceController) CreateDevices() {
	defer c.CatchException()
	req := &requestTypes.CreateDeviceRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("createImage unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	res, err := DeviceLogic.CreateDevices(c.logPoints, req) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// swagger:route POST /devices device createDevices
//     Responses:
//       200: createDevices
//       default: ErrorResponse
func (c *DeviceController) CollectDevice() {
	defer c.CatchException()
	req := &requestTypes.CollectDevice{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("CollectDevice unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	res, err := DeviceLogic.CollectDevice(c.logPoints, req) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// QueryDevices【设备管理】【设备列表】
// swagger:route GET /devices/queryDevices device queryDevices
// QueryDevices【设备管理】【设备列表】
//     Responses:
//       200: queryDevices
//       default: ErrorResponse

func (c *DeviceController) QueryDevices() {
	defer c.CatchException()
	page := requestTypes.PagingRequest{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	req := &requestTypes.QueryDeviceListRequest{
		Sn:            c.GetString("sn"),
		DeviceTypeID:  c.GetString("deviceTypeId"),
		DeviceSeries:  c.GetString("deviceSeries"),
		ManageStatus:  c.GetString("manageStatus"),
		IloIP:         c.GetString("iloIp"),
		InstanceID:    c.GetString("instanceId"),
		InstanceName:  c.GetString("instanceName"),
		UserID:        c.GetString("userId"),
		UserName:      c.GetString("userName"),
		Show:          c.GetString("show"),
		IsAll:         c.GetString("isAll"),
		CollectStatus: c.GetString("collectStatus"),
		PagingRequest: page,
	}
	result, err := DeviceLogic.GetDeviceList(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}

	if c.GetString("exportType") != "" {
		fileName, downloadFileName, err := DeviceLogic.ExportDeviceExcel(c.logPoints, result.Devices)
		if err != nil {
			c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, "导出设备列表excel错误 "+err.Error(), errorCode.INTERNAL)
			return
		}
		//c.Ctx.Input.SetData("download", "yes")
		c.Ctx.Output.Download(fileName, downloadFileName)
		if err := os.Remove(fileName); err != nil {
			c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, "删除上传文件失败"+err.Error(), errorCode.INTERNAL)
			return
		}
		return
	}
	c.Res.Result = result
}

// GetDevice【设备管理】【设备详情】
// swagger:route GET /devices/queryDevice/{device_id} device getDevice
// GetDevice【设备管理】【设备详情】
//     Responses:
//       200: getDevice
//       default: ErrorResponse

func (c *DeviceController) GetDevice() {
	defer c.CatchException()
	deviceId := c.Ctx.Input.Param(":device_id")
	req := &requestTypes.QueryDeviceInfoRequest{
		DeviceID: deviceId,
		Show:     c.GetString("show"),
	}
	result, err := DeviceLogic.GetDeviceInfo(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// UnPutaway【设备管理】【设备下架】支持多sn，英文逗号分隔
// swagger:route PUT /devices/unPutaway device unPutaway
// UnPutaway【设备管理】【设备下架】支持多sn，英文逗号分隔
//     Responses:
//       200: unPutaway
//       default: ErrorResponse

func (c *DeviceController) UnPutaway() {
	defer c.CatchException()

	req := &requestTypes.UnPutawayDeviceRequest{}
	// req.Validate()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("deviceController.UnPutaway.unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}

	result, err := DeviceLogic.UnPutawayDevice(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), err.Error())
		return
	}
	c.Res.Result = result
}

// Putaway【设备管理】【设备上架】支持多sn，英文逗号分隔
// swagger:route PUT /devices/putaway device putaway
// Putaway【设备管理】【设备上架】支持多sn，英文逗号分隔
//     Responses:
//       200: putaway
//       default: ErrorResponse

func (c *DeviceController) Putaway() {
	defer c.CatchException()
	req := &requestTypes.PutawayDeviceRequest{}
	// req.Validate()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("deviceController.Putaway.unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	result, err := DeviceLogic.PutawayDevice(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), err.Error())
		return
	}
	c.Res.Result = result
}

// Modify【设备管理】【设备修改】支持描述信息修改
// swagger:route PUT /devices/{device_id} device modify
// Modify【设备管理】【设备修改】支持描述信息修改
//     Responses:
//       200: modify
//       default: ErrorResponse

func (c *DeviceController) ModifyDevice() {
	defer c.CatchException()
	req := &requestTypes.ModifyDeviceRequest{}
	// req.Validate()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("deviceController.modify.unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	result, err := DeviceLogic.ModifyDevice(c.logPoints, req, c.Ctx.Input.Param(":device_id"))
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), err.Error())
		return
	}
	c.Res.Result = result
}

// DeleteDevice【设备管理】【删除设备】
// swagger:route DELETE /devices/{device_id} device deleteDevice
// DeleteDevice【设备管理】【删除设备】
//     Responses:
//       200: deleteDevice
//       default: ErrorResponse

func (c *DeviceController) DeleteDevice() {
	defer c.CatchException()
	deviceId := c.Ctx.Input.Param(":device_id")
	req := &requestTypes.DeleteDeviceRequest{
		DeviceId: deviceId,
	}
	result, err := DeviceLogic.DeleteDevice(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// swagger:route PUT /devices/{device_id}/remove device removeDevice
// RemoveDevice【设备管理】【移除设备】
//     Responses:
//       200: removeDevice
//       default: ErrorResponse

func (c *DeviceController) RemoveDevice() {
	defer c.CatchException()
	deviceId := c.Ctx.Input.Param(":device_id")
	result, err := DeviceLogic.RemoveDevice(c.logPoints, deviceId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// GetDeviceDiskDetail 设备详情-磁盘
func (c *DeviceController) GetDeviceDiskDetail() {
	defer c.CatchException()
	deviceId := c.Ctx.Input.Param(":device_id")
	req := &requestTypes.QueryDeviceDisksRequest{
		DeviceID: deviceId,
	}
	result, err := DeviceLogic.GetDeviceDiskDetail(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// swagger:route PUT /devices/disks/associateDisks device associateDeviceDisks
//
// AssociateDeviceDisks 设备关联磁盘
//
//     Responses:
//       200: associateDeviceDisks
//       default: ErrorResponse

func (c *DeviceController) AssociateDeviceDisks() {
	defer c.CatchException()
	req := &requestTypes.AssociateDeviceDisksRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("AssociateDeviceDisks parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	res, err := DeviceLogic.AssociateDeviceDisks(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	c.Res.Result = res
}

// swagger:route GET /devices/disks/describeAssociateDisks device getAssociatedDisks
//
// GetAssociatedDisks 获取设备已关联的磁盘
//
//     Responses:
//       200: getAssociatedDisks
//       default: ErrorResponse
func (c *DeviceController) GetAssociatedDisks() {
	defer c.CatchException()
	req := &requestTypes.GetAssociatedDisksRequest{
		DeviceID:     c.GetString("deviceId"),
		VolumeID:     c.GetString("volumeId"),
		DeviceTypeID: c.GetString("deviceTypeId"),
	}
	res, err := DeviceLogic.GetAssociatedDisks(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	c.Res.Result = res
}

// AssociateDeviceType ...
// swagger:route PUT /devices/associateDeviceType device deviceAssociateDeviceType
//
// AssociateDeviceType 设备绑定机型
//
//     Responses:
//       200: deviceAssociateDeviceType
//       default: ErrorResponse

func (c *DeviceController) DeviceAssociateDeviceType() {
	defer c.CatchException()
	req := &requestTypes.DeviceAssociateDeviceTypeRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("AssociateDeviceType parse post body err:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	res, err := DeviceLogic.DeviceAssociateDeviceType(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}
