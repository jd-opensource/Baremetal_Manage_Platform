package controllers

import (
	"regexp"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"

	imageLogic "coding.jd.com/aidc-bmp/bmp-openapi/logic/imageLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	validation "git.jd.com/cps-golang/ironic-common/ironic/common/Validation"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type ImageController struct {
	BaseController
}

// DescribeImages ...
// swagger:route GET /images image describeImages
// DescribeImages 获取镜像列表
//     Responses:
//       200: describeImages
//       default: ErrorResponse

func (c *ImageController) DescribeImages() {
	defer c.CatchException()
	req := &requestTypes.QueryImagesRequest{
		ImageID:      c.GetString("imageId"),
		DeviceTypeID: c.GetString("deviceTypeId"),
		ImageName:    c.GetString("imageName"),
		Version:      c.GetString("version"),
		OsID:         c.GetString("osId"),
		ImageIDs:     strings.Split(c.GetString("imageIds"), ","),
		Source:       c.GetString("source"),
		Architecture: c.GetString("architecture"),
		OsType:       c.GetString("osType"),
		IsAll:        c.GetString("isAll"),
	}
	req.Validate(c.logPoints)
	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	res, count, err := imageLogic.DescribeImages(c.logPoints, req, p)
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

// DescribeImage ...
// swagger:route GET /images/{image_id} image describeImage
// DescribeImage 获取镜像详情
//     Responses:
//       200: describeImage
//       default: ErrorResponse

func (c *ImageController) DescribeImage() {
	defer c.CatchException()
	image_id := c.Ctx.Input.Param(":image_id")
	if match, err := regexp.MatchString(validation.REGEX_ID, image_id); !match {
		if err != nil {
			c.logPoints.Warn("DescribeImage regexp match source error:", err.Error())
			c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
			return
		}
	}
	//image, err := imageDao.GetImageByUuid(image_id)
	res, err := imageLogic.GetByImageId(c.logPoints, image_id)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.ImageInfo{
		Image: res,
	}
}

// DescribeImageDeviceTypes ...
// swagger:route GET /images/imageDeviceTypes image describeImageDeviceTypes
// DescribeImageDeviceTypes 查看镜像绑定的机型
//     Responses:
//       200: describeImageDeviceTypes
//       default: ErrorResponse

func (c *ImageController) DescribeImageDeviceTypes() {
	defer c.CatchException()
	req := &requestTypes.QueryImageDeviceTypesRequest{
		ImageID:      c.GetString("imageId"),
		Architecture: c.GetString("architecture"),
		IsBind:       c.GetString("isBind"),
		IsAll:        c.GetString("isAll"),
	}
	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	res, count, err := imageLogic.QueryImageDeviceTypes(c.logPoints, req, p)
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
