package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/roleLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// RoleController operations for role
type RoleController struct {
	BaseController
}

// GetRoleList ...
// swagger:route GET /roles role describeRoles
// DescribeRoles 获取角色列表
//
//     Responses:
//       200: describeRoles
//       default: ErrorResponse

func (c *RoleController) DescribeRoles() {
	defer c.CatchException()
	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	req := &requestTypes.QueryRolesRequest{
		//RoleName: c.GetString("roleName"),
		IsAll: c.GetString("isAll"),
	}
	req.Validate(c.logPoints)
	res, count, err := roleLogic.GetRoleList(c.logPoints, *req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.RoleList{
		Roles:      res,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: count,
	}

}

// DescribeRole ...
// swagger:route GET /roles/{role_id} role describeRole
//
// DescribeRole 获取角色详情
//     Responses:
//       200: describeRole
//       default: ErrorResponse

func (c *RoleController) DescribeRole() {
	defer c.CatchException()
	roleId := c.Ctx.Input.Param(":role_id")
	res, err := roleLogic.GetRoleById(c.logPoints, roleId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.RoleInfo{
		Role: res,
	}
}

// swagger:route GET /roles/roleInfo/current role currentRole
//
// CurrentRole 获取当前登录用户的角色
//     Responses:
//       200: currentRole
//       default: ErrorResponse

func (c *RoleController) CurrentRole() {
	defer c.CatchException()

	res, err := roleLogic.CurrentRole(c.logPoints)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.RoleInfo{
		Role: res,
	}
}
