package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/resourceLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

type ResourceController struct {
	BaseController
}

// swagger:route GET /resources resource describeResources
//
// DescribeResources 根据某个字段值精确查询对应的资源列表（如根据imageName查询镜像列表，根据idcName查询机房列表）
//
//     Responses:
//       200: describeResources
//       default: ErrorResponse
func (c *ResourceController) DescribeResources() {
	defer c.CatchException()
	req := &requestTypes.QueryResourcesRequest{
		Name:       c.GetString("name"),
		DeviceType: c.GetString("deviceType"),
		ImageName:  c.GetString("imageName"),
		UserName:   c.GetString("userName"),
	}
	isExist, err := resourceLogic.DescribeResources(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: isExist,
	}
}
