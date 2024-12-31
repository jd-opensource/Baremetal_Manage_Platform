package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/projectLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// ProjectController operations for project
type ProjectController struct {
	BaseController
}

// CreateUserProject ...
// swagger:route POST /user/projects project createUserProject
// CreateUserProject 创建项目
//
//     Responses:
//       200: createUserProject
//       default: ErrorResponse

func (c *ProjectController) CreateUserProject() {
	defer c.CatchException()
	req := &requestTypes.CreateProjectRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateProject parse pProjectt body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	uuid, err := projectLogic.CreateProject(c.logPoints, req, "")
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.ProjectId{
		ProjectId: uuid,
	}
}

// swagger:route GET /user/projects project describeUserProjects
//
// DescribeUserProjects 获取项目列表
//     Responses:
//       200: describeUserProjects
//       default: ErrorResponse

func (c *ProjectController) DescribeUserProjects() {
	defer c.CatchException()
	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	owned, _ := c.GetInt("owned")
	req := &requestTypes.QueryProjectsRequest{
		ProjectName: c.GetString("projectName"),
		IsAll:       c.GetString("isAll"),
		Owned:       owned,
		OwnerName:   c.GetString("ownerName"),
		SharerName:  c.GetString("sharerName"),
	}
	if c.GetString("orderByCreatetime") == "desc" {
		req.OrderByCreatetime = "desc"
	} else {
		req.OrderByCreatetime = "asc"
	}
	req.Validate(c.logPoints)
	res, count, err := projectLogic.GetProjectList(c.logPoints, *req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.ProjectList{
		Projects:   res,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: count,
	}

}

// swagger:route GET /user/projects/{project_id} project describeUserProject
//
// DescribeUserProject 获取项目详情
//     Responses:
//       200: describeUserProject
//       default: ErrorResponse

func (c *ProjectController) DescribeUserProject() {
	defer c.CatchException()
	ProjectId := c.Ctx.Input.Param(":project_id")
	res, err := projectLogic.GetProjectById(c.logPoints, ProjectId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.ProjectInfo{
		Project: res,
	}
}

// swagger:route PUT /user/projects/{project_id} project modifyUserProject
//
// ModifyUserProject 修改项目信息
//     Responses:
//       200: modifyUserProject
//       default: ErrorResponse

func (c *ProjectController) ModifyUserProject() {
	defer c.CatchException()
	ProjectId := c.Ctx.Input.Param(":project_id")
	req := &requestTypes.ModifyProjectRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyProject parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := projectLogic.ModifyProject(c.logPoints, req, ProjectId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// swagger:route PUT /user/projects/{project_id}/description project modifyUserProjectDescription
//
// ModifyUserProject 修改项目信息
//     Responses:
//       200: modifyUserProjectDescription
//       default: ErrorResponse

func (c *ProjectController) ModifyUserProjectDescription() {
	defer c.CatchException()
	ProjectId := c.Ctx.Input.Param(":project_id")
	req := &requestTypes.ModifyProjectDescriptionRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyProjectDescription parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := projectLogic.ModifyProjectDescription(c.logPoints, req, ProjectId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// swagger:route DELETE /user/projects/{project_id} project deleteUserProject
//
// DeleteUserProject 删除项目
//     Responses:
//       200: deleteUserProject
//       default: ErrorResponse

func (c *ProjectController) DeleteUserProject() {
	defer c.CatchException()
	ProjectId := c.Ctx.Input.Param(":project_id")
	if err := projectLogic.DeleteProject(c.logPoints, ProjectId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /user/projects/{project_id}/share project shareUserProject
//
// ShareUserProject 共享项目
//     Responses:
//       200: shareUserProject
//       default: ErrorResponse

func (c *ProjectController) ShareUserProject() {
	defer c.CatchException()

	ProjectId := c.Ctx.Input.Param(":project_id")
	req := &requestTypes.ShareProjectRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyProject parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := projectLogic.AddShareProject(c.logPoints, ProjectId, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// swagger:route PUT /user/projects/{project_id}/cancelShare project cancelShareUserProject
//
// CancelShareUserProject 取消共享项目
//     Responses:
//       200: cancelShareUserProject
//       default: ErrorResponse

func (c *ProjectController) CancelShareUserProject() {
	defer c.CatchException()

	ProjectId := c.Ctx.Input.Param(":project_id")
	req := &requestTypes.CalcelShareProjectRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyProject parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := projectLogic.CancelShareProject(c.logPoints, ProjectId, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// swagger:route GET /user/projects/{project_id}/describeSharProject project describeShareProject
//
// DescribeShareUserProject 获取共享项目详情
//     Responses:
//       200: describeShareProject
//       default: ErrorResponse

func (c *ProjectController) DescribeShareUserProject() {
	defer c.CatchException()

	ProjectId := c.Ctx.Input.Param(":project_id")

	res, err := projectLogic.DescribeShareProject(c.logPoints, ProjectId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.ShareProjectInfo{
		ProjectID: ProjectId,
		Shares:    res,
	}
}

// swagger:route PUT /user/projects/{project_id}/move project moveUserProject
//
// MoveUserProject 转移项目
//     Responses:
//       200: moveUserProject
//       default: ErrorResponse

func (c *ProjectController) MoveUserProject() {
	defer c.CatchException()

	ProjectId := c.Ctx.Input.Param(":project_id")
	req := &requestTypes.MoveProjectRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("MoveProject parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := projectLogic.MoveProject(c.logPoints, ProjectId, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// swagger:route PUT /user/projects/move/instances project moveUserInstances
//
// MoveUserInstances 转移项目下的某些实例
//     Responses:
//       200: moveUserInstances
//       default: ErrorResponse

func (c *ProjectController) MoveUserInstances() {
	defer c.CatchException()

	req := &requestTypes.MoveInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("MoveProject parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := projectLogic.MoveInstances(c.logPoints, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}
