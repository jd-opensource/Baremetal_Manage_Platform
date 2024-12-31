package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/deviceTypeLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type DeviceTypeController struct {
	BaseController
}

// swagger:route POST /deviceTypes deviceType createDeviceType
//
// CreateDeviceType 添加机型
//     Responses:
//       200: createDeviceType
//       default: ErrorResponse

func (c *DeviceTypeController) CreateDeviceType() {
	defer c.CatchException()
	req := &requestTypes.CreateDeviceTypeRequest{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, req)
	if err != nil {
		c.logPoints.Warn("CreateDeviceType parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	uuid, err := deviceTypeLogic.Create(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.DeviceTypeId{
		DeviceTypeId: uuid,
	}
}

// swagger:route GET /deviceTypes deviceType describeDeviceTypes
// DescribeDeviceTypes 获取机型列表
//     Responses:
//       200: describeDeviceTypes
//       default: ErrorResponse

func (c *DeviceTypeController) DescribeDeviceTypes() {
	defer c.CatchException()
	req := &requestTypes.QueryDeviceTypesRequest{
		DeviceTypeID: c.GetString("deviceTypeId"),
		DeviceType:   c.GetString("deviceType"),
		IdcID:        c.GetString("idcId"),
		Name:         c.GetString("name"),
		DeviceSeries: c.GetString("deviceSeries"),
		IsAll:        c.GetString("isAll"),
	}
	req.Validate(c.logPoints)
	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	//device_types, err := deviceTypeLogic.QueryByRequest(c.logPoints, req)
	res, count, err := deviceTypeLogic.QueryDeviceTypes(c.logPoints, req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.DeviceTypeList{
		DeviceTypes: res,
		PageNumber:  c.pageable.PageNumber,
		PageSize:    c.pageable.PageSize,
		TotalCount:  count,
	}
}

// swagger:route GET /deviceTypes/{device_type_id} deviceType describeDeviceType
//
//
// DescribeDeviceType 获取机型详情
//
//     Responses:
//       200: describeDeviceType
//       default: ErrorResponse

func (c *DeviceTypeController) DescribeDeviceType() {
	defer c.CatchException()
	req := &requestTypes.QueryDeviceTypesRequest{
		DeviceTypeID: c.Ctx.Input.Param(":device_type_id"),
	}
	req.Validate(c.logPoints)
	res, err := deviceTypeLogic.QueryDeviceType(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.DeviceTypeInfo{
		DeviceType: res,
	}
}

// ModifyDeviceType ...
// swagger:route PUT /deviceTypes/{device_type_id} deviceType modifyDeviceType
//
// ModifyDeviceType 修改机型信息
//
//     Responses:
//       200: modifyDeviceType
//       default: ErrorResponse

func (c *DeviceTypeController) ModifyDeviceType() {
	defer c.CatchException()
	req := &requestTypes.ModifyDeviceTypeRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyDeviceType parse post body err:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := deviceTypeLogic.ModifyDeviceType(c.logPoints, req, c.Ctx.Input.Param(":device_type_id")); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// DeleteDeviceType ...
// swagger:route DELETE /deviceTypes/{device_type_id} deviceType deleteDeviceType
//
// DeleteDeviceType 删除机型
//
//     Responses:
//       200: deleteDeviceType
//       default: ErrorResponse

func (c *DeviceTypeController) DeleteDeviceType() {
	defer c.CatchException()
	deviceTypeID := c.Ctx.Input.Param(":device_type_id")
	if deviceTypeID == "" {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, "deviceTypeId is not empty", errorCode.INTERNAL)
		return
	}
	if err := deviceTypeLogic.DeleteDeviceType(c.logPoints, deviceTypeID); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// AssociatedImage ...
// swagger:route POST /deviceTypes/associatedImage deviceType associatedImage
//
// AssociatedImage 机型绑定镜像
//
//     Responses:
//       200: associatedImage
//       default: ErrorResponse

func (c *DeviceTypeController) AssociatedImage() {
	defer c.CatchException()
	req := &requestTypes.AssociateImageRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("AssociatedImage parse post body err:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := deviceTypeLogic.AssociatedImage(c.logPoints, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// DissociatedImage ...
// swagger:route DELETE /deviceTypes/dissociatedImage deviceType dissociatedImage
//
// DissociatedImage 机型解绑镜像
//
//     Responses:
//       200: dissociatedImage
//       default: ErrorResponse

func (c *DeviceTypeController) DissociatedImage() {
	defer c.CatchException()
	req := &requestTypes.DissociatedImageRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("DissociatedImage parse post body err:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := deviceTypeLogic.DissociatedImage(c.logPoints, req, false); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route GET /deviceTypes/deviceTypeImage deviceType describeDeviceTypeImages
//
// QueryDeviceTypeImage 根据机型获取镜像
//
//     Responses:
//       200: describeDeviceTypeImages
//       default: ErrorResponse

func (c *DeviceTypeController) DescribeDeviceTypeImages() {
	req := &requestTypes.QueryDeviceTypeImageRequest{
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
	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	res, count, err := deviceTypeLogic.QueryDeviceTypeImage(c.logPoints, req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.ImageList{
		Images:     res,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: count,
	}
}

// swagger:route GET /deviceTypes/deviceTypeRaid deviceType describeDeviceTypeRaids
//
// DescribeDeviceTypeRaids 根据机型获取raid
//
//     Responses:
//       200: describeDeviceTypeRaids
//       default: ErrorResponse

func (c *DeviceTypeController) DescribeDeviceTypeRaids() {
	//req := &requestTypes.QueryDeviceTypeRaidRequest{
	//	DeviceTypeID: c.GetString("deviceTypeId"),
	//	VolumeType:   c.GetString("volumeType"),
	//	RaidID:       c.GetString("raidId"),
	//}
	//res, err := deviceTypeLogic.QueryDeviceTypeRaid(c.logPoints, req)
	//if err != nil {
	//	c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	//	return
	//}
	c.Res.Result = nil //废弃
}

// swagger:route GET /deviceTypes/deviceTypeImagePartition deviceType describeDeviceTypeImagePartitions
//
// QueryDeviceTypeImagePartition 根据机型，镜像，获取分区
//
//     Responses:
//       200: describeDeviceTypeImagePartitions
//       default: ErrorResponse

func (c *DeviceTypeController) DescribeDeviceTypeImagePartitions() {
	req := &requestTypes.QueryDeviceTypeImagePartitionRequest{
		DeviceTypeID: c.GetString("deviceTypeId"),
		ImageID:      c.GetString("imageId"),
	}
	req.Validate(c.logPoints)
	res, err := deviceTypeLogic.QueryDeviceTypeImagePartition(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
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
	req := &requestTypes.DescribeVolumesByDeviceTypeRequest{
		DeviceTypeID: c.Ctx.Input.Param(":device_type_id"),
	}

	req.Validate(c.logPoints)
	res, err := deviceTypeLogic.DescribeVolumesByDeviceType(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}
