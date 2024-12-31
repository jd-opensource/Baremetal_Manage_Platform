package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/RoleLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type RoleController struct {
	BaseController
}

// DescribeOSs ...
// swagger:route GET /oss os describeOSs
//
//     Responses:
//       200: describeOSs
//       default: ErrorResponse

func (c *RoleController) DescribeRoles() {
	defer c.CatchException()
	req := &requestTypes.DescribeRolesRequest{
		IsAll: c.GetString("isAll"),
	}
	res, err := RoleLogic.DescribeRoles(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	if res.Roles == nil {
		res.Roles = make([]*response.Role, 0)
	}
	c.Res.Result = res
}

// swagger:route GET /roles/roleInfo/current role currentRole
//
// CurrentRole 获取当前登录用户的角色
//     Responses:
//       200: currentRole
//       default: ErrorResponse

func (c *RoleController) CurrentRole() {
	defer c.CatchException()

	res, err := RoleLogic.CurrentRole(c.logPoints)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.RoleInfo{
		Role: res,
	}
}
