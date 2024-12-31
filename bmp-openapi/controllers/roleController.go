package controllers

import (
	"encoding/json"

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

// CreateRole ...
// swagger:route POST /roles role createRole
//
// CreateRole 创建角色(暂不启用)
//     Responses:
//       200: createRole
//       default: ErrorResponse

func (c *RoleController) CreateRole() {
	defer c.CatchException()
	req := &requestTypes.CreateRoleRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateRole parse pRolet body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	uuid, err := roleLogic.CreateRole(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.RoleId{
		RoleId: uuid,
	}
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

// ModifyRole ...
// swagger:route PUT /roles/{role_id} role modifyRole
//
// ModifyRole 修改角色信息(暂不启用)
//     Responses:
//       200: modifyRole
//       default: ErrorResponse

func (c *RoleController) ModifyRole() {
	defer c.CatchException()
	roleId := c.Ctx.Input.Param(":role_id")
	req := &requestTypes.ModifyRoleRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyRole parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	//req.Validate(c.logPoints)
	if err := roleLogic.ModifyRole(c.logPoints, req, roleId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// DeleteRole ...
// swagger:route DELETE /roles/{role_id} role deleteRole
//
// DeleteRole 删除角色(暂不启用)
//     Responses:
//       200: deleteRole
//       default: ErrorResponse

func (c *RoleController) DeleteRole() {
	defer c.CatchException()
	roleId := c.Ctx.Input.Param(":role_id")
	if err := roleLogic.DeleteRole(c.logPoints, roleId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
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
