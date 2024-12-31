package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/deviceTypeLogic"
	"encoding/json"
	"fmt"
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

// CreateImage ...
// swagger:route POST /images image createImage
// CreateImage 添加镜像
//     Responses:
//       200: createImage
//       default: ErrorResponse

func (c *ImageController) CreateImage() {
	defer c.CatchException()
	req := &requestTypes.CreateImageRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("parse CreateImage body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	uuid, err := imageLogic.CreateImage(c.logPoints, req) //返回id
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.ImageId{
		ImageId: uuid,
	}
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

// ModifyImage ...
// swagger:route PUT /images/{image_id} image modifyImage
// ModifyImage 修改镜像(暂不启用)
//     Responses:
//       200: modifyImage
//       default: ErrorResponse

func (c *ImageController) ModifyImage() {
	defer c.CatchException()
	imageId := c.Ctx.Input.Param(":image_id")
	if match, err := regexp.MatchString(validation.REGEX_ID, imageId); !match {
		c.logPoints.Warn("DescribeImage regexp match image_id false")
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req := &requestTypes.ModifyImageRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateImage parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := imageLogic.ModifyImage(c.logPoints, req, imageId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// DeleteImage ...
// swagger:route DELETE /images/{image_id} image deleteImage
// DeleteImage 删除镜像
//     Responses:
//       200: deleteImage
//       default: ErrorResponse

func (c *ImageController) DeleteImage() {
	defer c.CatchException()
	imageId := c.Ctx.Input.Param(":image_id")
	if match, err := regexp.MatchString(validation.REGEX_ID, imageId); !match {
		if err != nil {
			c.logPoints.Warn("DeleteImage regexp match source error:", err.Error())
		}
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	if err := imageLogic.DeleteImage(c.logPoints, imageId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
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

// swagger:route POST /images/associatedDeviceType image associatedDeviceType
//
// AssociatedDeviceType 镜像绑定机型
//
//     Responses:
//       200: associatedDeviceType
//       default: ErrorResponse

func (c *ImageController) AssociatedDeviceType() {
	defer c.CatchException()
	req := &requestTypes.AssociateDeviceTypeRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("AssociateDeviceType parse post body err:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	for _, deviceTypeId := range req.DeviceTypeIDs {
		reqNew := &requestTypes.AssociateImageRequest{
			DeviceTypeID: deviceTypeId,
			ImageIDs:     []string{req.ImageID},
		}
		fmt.Println(reqNew)
		if err := deviceTypeLogic.AssociatedImage(c.logPoints, reqNew); err != nil {
			c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
			return
		}
	}

	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route DELETE /images/dissociatedDeviceType image dissociatedDeviceType
//
// DissociatedDeviceType 镜像解绑机型
//
//     Responses:
//       200: dissociatedDeviceType
//       default: ErrorResponse

func (c *ImageController) DissociatedDeviceType() {
	defer c.CatchException()
	req := &requestTypes.DissociatedImageRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("DissociatedImage parse post body err:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := deviceTypeLogic.DissociatedImage(c.logPoints, req, true); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}
