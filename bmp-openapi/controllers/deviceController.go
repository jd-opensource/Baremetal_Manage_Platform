package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/deviceLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceController operations for Device
type DeviceController struct {
	BaseController
}

// swagger:route GET /devices device describeDevices
//
// DescribeDevices 获取设备列表
//
//     Responses:
//       200: describeDevices
//       default: ErrorResponse

func (c *DeviceController) DescribeDevices() {
	defer c.CatchException()
	param := requestTypes.QueryDevicesRequest{
		IDcID:        c.GetString("idcId"),
		Sn:           c.GetString("sn"),
		DeviceSeries: c.GetString("deviceSeries"),
		DeviceTypeID: c.GetString("deviceTypeId"),
		ManageStatus: c.GetString("manageStatus"),

		IloIP:         c.GetString("iloIp"),
		InstanceID:    c.GetString("instanceId"),
		InstanceName:  c.GetString("instanceName"),
		UserID:        c.GetString("user_id"),
		UserName:      c.GetString("userName"),
		IsAll:         c.GetString("isAll"),
		CollectStatus: c.GetString("collectStatus"),
	}

	res, count, err := deviceLogic.QueryDevices(c.logPoints, param, c.pageable)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.DeviceList{
		Devices:    res,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: count,
	}
}

// swagger:route GET /devices/{device_id} device describeDevice
//
// DescribeDevice 获取设备详情
//
//     Responses:
//       200: describeDevice
//       default: ErrorResponse

func (c *DeviceController) DescribeDevice() {
	defer c.CatchException()
	deviceId := c.Ctx.Input.Param(":device_id")
	//id, err := strconv.ParseInt(idStr, 10, 64)
	//if err != nil {
	//	c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	//	return
	//}
	res, err := deviceLogic.GetById(c.logPoints, deviceId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.DeviceInfo{
		Device: res,
	}
}

// swagger:route POST /devices device createDevice
//
//  CreateDevice 创建设备
//
//     Responses:
//       200: createDevice
//       default: ErrorResponse

func (c *DeviceController) CreateDevice() {
	defer c.CatchException()
	req := &requestTypes.CreateDevicesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	_, err := deviceLogic.Save(c.logPoints, *req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	//c.Res.Result = response.DeviceIds{
	//	DeviceIds: deviceIds,
	//}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /devices/{device_id} device modifyDevice
//
// ModifyDevice 修改设备信息
//
//     Responses:
//       200: modifyDevice
//       default: ErrorResponse

func (c *DeviceController) ModifyDevice() {
	defer c.CatchException()
	req := &requestTypes.ModifyDevicesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := deviceLogic.ModifyDeviceDescription(c.logPoints, req, c.Ctx.Input.Param(":device_id")); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /devices/mount device mountDevice
//
// MountDevice 设备上架
//     Responses:
//       200: mountDevice
//       default: ErrorResponse

func (c *DeviceController) MountDevice() {
	defer c.CatchException()
	req := &requestTypes.MountDevicesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	if err := deviceLogic.ModifyDevice(c.logPoints, req.DeviceID, baseLogic.PUTAWAY); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
	//defer c.CatchException()
	//req := &requestTypes.PutawayDevicesRequest{}
	//if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
	//	c.logPoints.Warn("parse post body error:", err.Error())
	//	c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	//	return
	//}
	//req.Validate(c.logPoints)
	//if err := deviceLogic.Putaway(c.logPoints, req.Sns); err != nil {
	//	c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	//	return
	//}
	//c.Res.Result = response.CommonResponse{
	//	Success: true,
	//}
}

// swagger:route PUT /devices/unmount device unMountDevice
//
// UnMountDevice 设备下架
//
//     Responses:
//       200: unMountDevice
//       default: ErrorResponse

func (c *DeviceController) UnMountDevice() {
	defer c.CatchException()
	req := &requestTypes.UnMountDevicesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	if err := deviceLogic.ModifyDevice(c.logPoints, req.DeviceID, baseLogic.UNPUTAWAY); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route DELETE /devices/{device_id} device deleteDevice
//
// DeleteDevice 删除设备
//     Responses:
//       200: deleteDevice
//       default: ErrorResponse

func (c *DeviceController) DeleteDevice() {
	defer c.CatchException()
	//req := &requestTypes.DeleteDevicesRequest{}
	//if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
	//	c.logPoints.Warn("parse post body error:", err.Error())
	//	c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	//	return
	//}
	//req.Validate(c.logPoints)
	if err := deviceLogic.Delete(c.logPoints, c.Ctx.Input.Param(":device_id")); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /devices/{device_id}/remove device removeDevice
//
// DeleteDevice 移除设备
//     Responses:
//       200: removeDevice
//       default: ErrorResponse

func (c *DeviceController) RemoveDevice() {
	defer c.CatchException()

	device_id := c.Ctx.Input.Param(":device_id")
	if err := deviceLogic.RemoveDevice(c.logPoints, device_id); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route DELETE /devices/stock device describeDeviceStock
// DescribeDeviceStock 获取指定机型的设备库存
//     Responses:
//       200: describeDeviceStock
//       default: ErrorResponse
func (c *DeviceController) DescribeDeviceStock() {
	defer c.CatchException()
	list := deviceLogic.GetAllDevices(c.logPoints, map[string]interface{}{
		"manage_status":  baseLogic.PUTAWAY,
		"device_type_id": c.GetString("deviceTypeId"),
	}, []string{}, []string{}, []string{})
	c.Res.Result = response.DeviceStock{
		Stocks: len(list),
	}
}

// swagger:route POST /collect/collectDeviceInfo device collectDeviceInfo
// CollectDeviceInfo 采集设备信息
//     Responses:
//       200: collectDeviceInfo
//       default: ErrorResponse
func (c *DeviceController) CollectDeviceInfo() {
	defer c.CatchException()
	req := &requestTypes.CollectDeviceInfoRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := deviceLogic.CollectDeviceInfo(c.logPoints, req); err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route GET /devices/{device_id}/disksDetail device describeDeviceDisks
//
// DescribeDeviceDisks 设备详情-磁盘
//
//     Responses:
//       200: describeDeviceDisks
//       default: ErrorResponse

func (c *DeviceController) DescribeDeviceDisks() {
	defer c.CatchException()
	deviceId := c.Ctx.Input.Param(":device_id")
	res, err := deviceLogic.GetDeviceDisksById(c.logPoints, deviceId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
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
	req.Validate(c.logPoints)
	res, err := deviceLogic.AssociateDeviceDisks(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	c.Res.Result = response.CommonResponse{
		Success: res,
	}
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
	req.Validate(c.logPoints)
	res, err := deviceLogic.GetAssociatedDisks(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	c.Res.Result = res
}

// AssociateDeviceType ...
// swagger:route PUT /devices/deviceType/associate device deviceAssociateDeviceType
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
	req.Validate(c.logPoints)
	res, err := deviceLogic.DeviceAssociateDeviceType(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}
