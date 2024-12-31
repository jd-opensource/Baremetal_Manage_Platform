package controllers

import (
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

		IloIP:      c.GetString("iloIp"),
		InstanceID: c.GetString("instanceId"),
		UserID:     c.GetString("user_id"),
		UserName:   c.GetString("userName"),
		IsAll:      c.GetString("isAll"),
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
