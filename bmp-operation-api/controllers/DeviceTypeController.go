package controllers

import (
	"encoding/json"
	"os"

	DeviceTypeLogic "coding.jd.com/aidc-bmp/bmp-operation-api/logic/DeviceTypeLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

type DeviceTypeController struct {
	BaseController
}

// CreateDeviceType【机型管理】【添加机型】
// swagger:route PUT /deviceTypes deviceType createDeviceType
// CreateDeviceType【机型管理】【添加机型】
//     Responses:
//       200: createDeviceType
//       default: ErrorResponse

func (c *DeviceTypeController) CreateDeviceType() {
	defer c.CatchException()
	req := &requestTypes.CreateDeviceTypeRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	result, err := DeviceTypeLogic.AddDeviceType(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// QueryDeviceTypes【机型管理】【机型列表】
// swagger:route GET /deviceTypes deviceType queryDeviceTypes
// QueryDeviceTypes【机型管理】【机型列表】
//     Responses:
//       200: queryDeviceTypes
//       default: ErrorResponse

func (c *DeviceTypeController) QueryDeviceTypes() {
	defer c.CatchException()
	page := requestTypes.PagingRequest{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	req := &requestTypes.QueryDeviceTypesRequest{
		DeviceTypeID:  c.GetString("deviceTypeId"),
		DeviceType:    c.GetString("deviceType"),
		IdcID:         c.GetString("idcId"),
		Name:          c.GetString("name"),
		DeviceSeries:  c.GetString("deviceSeries"),
		IsAll:         c.GetString("isAll"),
		PagingRequest: page,
	}
	result, err := DeviceTypeLogic.GetDeviceTypeList(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}

	if c.GetString("exportType") != "" {
		fileName, downloadFileName, err := DeviceTypeLogic.ExportDeviceTypeExcel(c.logPoints, result.DeviceTypes)
		if err != nil {
			c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, "导出设备列表excel错误"+err.Error(), errorCode.INTERNAL)
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

// QueryDeviceType【机型管理】【机型详情信息】
// swagger:route GET /deviceTypes/queryDeviceType/{device_type_id} deviceType queryDeviceType
// QueryDeviceType【机型管理】【机型详情信息】
//     Responses:
//       200: queryDeviceType
//       default: ErrorResponse

func (c *DeviceTypeController) QueryDeviceType() {
	defer c.CatchException()
	deviceTypeId := c.Ctx.Input.Param(":device_type_id")
	req := &requestTypes.QueryDeviceTypeRequest{
		DeviceTypeID: deviceTypeId,
	}
	result, err := DeviceTypeLogic.GetDeviceTypeInfo(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// QueryDeviceTypeImages【机型管理】【机型详情信息】【关联镜像列表】此接口无分页参数
// swagger:route GET /deviceTypes/deviceTypeImage deviceType queryDeviceTypeImages
// QueryDeviceTypeImages【机型管理】【机型详情信息】【关联镜像列表】此接口无分页参数
//     Responses:
//       200: queryDeviceTypeImages
//       default: ErrorResponse

func (c *DeviceTypeController) QueryDeviceTypeImages() {
	defer c.CatchException()
	req := &requestTypes.QueryDeviceTypeImagePageRequest{
		DeviceTypeID: c.GetString("deviceTypeId"),
		ImageID:      c.GetString("imageId"),
		Architecture: c.GetString("architecture"),
		OsType:       c.GetString("osType"),

		ImageName: c.GetString("imageName"),
		Version:   c.GetString("version"),
		OsID:      c.GetString("osId"),
		//ImageIDs:  strings.Split(c.GetString("imageIds"), ","),
		Source: c.GetString("source"),
		IsAll:  c.GetString("isAll"),
	}
	result, err := DeviceTypeLogic.GetDeviceTypeImageList(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// AssociateDeviceTypeImage【机型管理】【机型详情信息】【关联镜像列表】【添加镜像】
// swagger:route POST /deviceTypes/associatedImage deviceType associateDeviceTypeImage
// AssociateDeviceTypeImage【机型管理】【机型详情信息】【关联镜像列表】【添加镜像】
//     Responses:
//       200: associateDeviceTypeImage
//       default: ErrorResponse

func (c *DeviceTypeController) AssociateDeviceTypeImage() {
	defer c.CatchException()
	req := &requestTypes.AssociateDeviceTypeImageRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("AssociateDeviceTypeImage parse request body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	result, err := DeviceTypeLogic.AssociateDeviceTypeImage(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// DisassociateDeviceTypeImage【机型管理】【机型详情信息】【关联镜像列表】【删除镜像】
// swagger:route POST /deviceTypes/dissociatedImage deviceType disassociateDeviceTypeImage
// DisassociateDeviceTypeImage【机型管理】【机型详情信息】【关联镜像列表】【删除镜像】
//     Responses:
//       200: disassociateDeviceTypeImage
//       default: ErrorResponse

func (c *DeviceTypeController) DisassociateDeviceTypeImage() {
	defer c.CatchException()
	req := &requestTypes.DisassociateDeviceTypeImageRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("DisassociateDeviceTypeImage parse request body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	result, err := DeviceTypeLogic.DisassociateDeviceTypeImage(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// ModifyDeviceType【机型管理】【编辑机型】
// swagger:route PUT /deviceTypes/{device_type_id} deviceType modifyDeviceType
// ModifyDeviceType【机型管理】【编辑机型】
//     Responses:
//       200: modifyDeviceType
//       default: ErrorResponse

func (c *DeviceTypeController) ModifyDeviceType() {
	defer c.CatchException()
	deviceTypeId := c.Ctx.Input.Param(":device_type_id")
	req := &requestTypes.ModifyDeviceTypeRequest{
		DeviceTypeID: deviceTypeId,
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	result, err := DeviceTypeLogic.ModifyDeviceTypeInfo(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// DeleteDeviceType【机型管理】【删除机型】
// swagger:route DELETE /deviceTypes/{device_type_id} deviceType deleteDeviceType
// DeleteDeviceType【机型管理】【删除机型】
//     Responses:
//       200: deleteDeviceType
//       default: ErrorResponse

func (c *DeviceTypeController) DeleteDeviceType() {
	defer c.CatchException()
	deviceTypeId := c.Ctx.Input.Param(":device_type_id")
	req := &requestTypes.DeleteDeviceTypeRequest{
		DeviceTypeID: deviceTypeId,
	}
	result, err := DeviceTypeLogic.DeleteDeviceTypeInfo(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// swagger:route GET /deviceType/getAvailableDeviceTypes deviceType getAvailableDeviceTypes
// GetAvailableDeviceTypes 获取该机房内可用机型列表
//     Responses:
//       200: getAvailableDeviceTypes
//       default: ErrorResponse

func (c *DeviceTypeController) GetAvailableDeviceTypes() {
	defer c.CatchException()
	req := requestTypes.GetAvailableDeviceTypesRequest{
		IdcID: c.GetString("idc_id"),
	}

	// req.Validate(c.logPoints)
	res, err := DeviceTypeLogic.GetAvailableDeviceTypes(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = res

}

// swagger:route GET /deviceTypes/{device_type_id}/describeVolumes deviceType describeVolumesByDeviceType
//
// DescribeVolumesByDeviceType 根据机型获取volumes
//
//     Responses:
//       200: describeVolumesByDeviceType
//       default: ErrorResponse

func (c *DeviceTypeController) DescribeVolumesByDeviceType() {
	deviceTypeId := c.Ctx.Input.Param(":device_type_id")

	res, err := DeviceTypeLogic.DescribeVolumesByDeviceType(c.logPoints, deviceTypeId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}
