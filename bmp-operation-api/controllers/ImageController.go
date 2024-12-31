package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"

	imageLogic "coding.jd.com/aidc-bmp/bmp-operation-api/logic/ImageLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	validation "git.jd.com/cps-golang/ironic-common/ironic/common/Validation"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type ImageController struct {
	BaseController
}

func (c *ImageController) UploadImage() {
	defer c.CatchException()
	file, header, err := c.GetFile("imageFile")
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	defer file.Close()
	req := &requestTypes.UploadImageRequest{
		ImageFile: file,
	}
	res, err := imageLogic.UploadImage(c.logPoints, req, header, c.GetRequestID()) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// swagger:route POST /images image createImage
//
//     Responses:
//       200: createImage
//       default: ErrorResponse
func (c *ImageController) CreateImage() {
	defer c.CatchException()
	req := &requestTypes.CreateImageRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("createImage unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	res, err := imageLogic.CreateImage(c.logPoints, req) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// DescribeImages ...
// swagger:route GET /images image describeImages
//
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
		PageNumber:   c.getPageNumber(),
		PageSize:     c.getPageSize(),
		IsAll:        c.GetString("isAll"),
		Architecture: c.GetString("architecture"),
		OsType:       c.GetString("osType"),
	}
	//req.Validate(c.logPoints)

	res, err := imageLogic.DescribeImages(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	fmt.Println(res.Images)
	if c.GetString("exportType") != "" {
		fileName, downloadFileName, err := imageLogic.ExportImage(c.logPoints, res.Images)
		if err != nil {
			c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, "导出excel错误 "+err.Error(), errorCode.INTERNAL)
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
	/*if res.Images == nil {
		res.Images = make([]&response.Image, 0)
	}*/
	c.Res.Result = res
}

// DescribeImage ...
// swagger:route GET /images/{image_id} image describeImage
//
//     Responses:
//       200: describeImage
//       default: ErrorResponse

func (c *ImageController) DescribeImage() {
	defer c.CatchException()
	imageId := c.Ctx.Input.Param(":image_id")
	if match, err := regexp.MatchString(validation.REGEX_ID, imageId); !match {
		if err != nil {
			c.logPoints.Warn("DeleteImage regexp match source error:", err.Error())
		}
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	res, err := imageLogic.DescribeImage(c.logPoints, imageId)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = res
}

// ModifyImage ...
// swagger:route PUT /images/{image_id} image modifyImage
//
//     Responses:
//       200: modifyImage
//       default: ErrorResponse

func (c *ImageController) ModifyImage() {
	defer c.CatchException()
	imageId := c.Ctx.Input.Param(":image_id")
	if match, err := regexp.MatchString(validation.REGEX_ID, imageId); !match {
		if err != nil {
			c.logPoints.Warn("DeleteImage regexp match source error:", err.Error())
		}
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req := &requestTypes.ModifyImageRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warnf("createImage unmarshal error:%s", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	res, err := imageLogic.ModifyImage(c.logPoints, req, imageId)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = res
}

// DeleteImage ...
// swagger:route DELETE /images/{image_id} image deleteImage
//
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
	res, err := imageLogic.DeleteImage(c.logPoints, imageId)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = res
}

// DescribeImageDeviceTypes ...
// swagger:route DELETE /images/imageDeviceTypes image describeImageDeviceTypes
//
//     Responses:
//       200: describeImageDeviceTypes
//       default: ErrorResponse

func (c *ImageController) DescribeImageDeviceTypes() {
	defer c.CatchException()
	req := &requestTypes.QueryImageDeviceTypesRequest{
		ImageID:      c.GetString("imageId"),
		Architecture: c.GetString("architecture"),
		IsAll:        c.GetString("isAll"),
		PageNumber:   c.getPageNumber(),
		PageSize:     c.getPageSize(),
		IsBind:       c.GetString("isBind"),
	}
	result, err := imageLogic.DescribeImageDeviceTypes(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
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
	req := &requestTypes.AssociateImageDeviceTypeRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("AssociateDeviceType parse post body err:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	res, err := imageLogic.AssociatedDeviceType(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route DELETE /deviceTypes/dissociatedDeviceType deviceType dissociatedDeviceType
//
// DissociatedDeviceType 镜像解绑机型
//
//     Responses:
//       200: dissociatedDeviceType
//       default: ErrorResponse

func (c *ImageController) DissociatedDeviceType() {
	defer c.CatchException()
	req := &requestTypes.DisassociateDeviceTypeImageRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("DissociatedImage parse post body err:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	//res, err := deviceTypeLogic.DisassociateDeviceTypeImage(c.logPoints, req) //直接调用机型解绑镜像接口
	res, err := imageLogic.DissociatedDeviceType(c.logPoints, req) //直接调用镜像解绑机型接口，不要调用机型解绑镜像，虽然接口参数一样，但是镜像解绑机型接口还要判断机型的实例是否有创建中的
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route GET /image/queryImagesByDeviceType image queryImagesByDeviceType
// QueryImagesByDeviceType 根据机型获取绑定的镜像列表
//     Responses:
//       200: queryImagesByDeviceType
//       default: ErrorResponse

func (c *ImageController) QueryImagesByDeviceType() {
	defer c.CatchException()

	req := requestTypes.QueryImagesByDeviceTypeRequest{
		IdcID:        c.GetString("idc_id"),
		DeviceTypeID: c.GetString("deviceTypeId"),
	}
	// req.Validate(c.logPoints)
	res, err := imageLogic.QueryImagesByDeviceType(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = res

}
