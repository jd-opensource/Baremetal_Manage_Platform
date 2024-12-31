package controllers

import (
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

// swagger:route GET /deviceTypes/volumesRaids deviceType describeVolumesRaids
//
// DescribeVolumesRaids 根据机型获取卷和raid信息 (deviceTypeRaid->volumesRaids)
//
//     Responses:
//       200: describeVolumesRaids
//       default: ErrorResponse

func (c *DeviceTypeController) DescribeVolumesRaids() {
	req := &requestTypes.QueryVolumesRaidsRequest{
		DeviceTypeID: c.GetString("deviceTypeId"),
		VolumeType:   c.GetString("volumeType"),
	}
	res, err := deviceTypeLogic.QueryVolumesRaids(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
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
