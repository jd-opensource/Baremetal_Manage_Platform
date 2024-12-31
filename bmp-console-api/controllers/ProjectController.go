package controllers

import (
	"encoding/json"

	projectLogic "coding.jd.com/aidc-bmp/bmp-console-api/logic/ProjectLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// ProjectController operations for project
type ProjectController struct {
	BaseController
}

// swagger:route PUT /project project createProject
// CreateProject 创建项目
//     Responses:
//       200: createProject
//       default: ErrorResponse

func (c *ProjectController) CreateProject() {
	defer c.CatchException()
	req := &sdkModels.CreateProjectRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateProject parse pProjectt body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	// req.Validate(c.logPoints)
	uuid, err := projectLogic.CreateProject(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = response.ProjectId{
		ProjectId: uuid,
	}
}

// swagger:route GET /project project getProjectList
// GetProjectList 获取项目列表
//     Responses:
//       200: getProjectList
//       default: ErrorResponse

func (c *ProjectController) GetProjectList() {
	defer c.CatchException()
	req := &requestTypes.QueryProjectsRequest{
		ProjectName:       c.GetString("projectName"),
		ExportType:        c.GetString("exportType"),
		IsAll:             c.GetString("isAll"),
		OrderByCreatetime: c.GetString("orderByCreatetime"),
		Owned:             c.GetString("owned"),
		OwnerName:         c.GetString("ownerName"),
		SharerName:        c.GetString("sharerName"),
		PagingRequest: requestTypes.PagingRequest{
			PageNumber: c.getPageNumber(),
			PageSize:   c.getPageSize(),
		},
	}
	req.Validate(c.logPoints)
	res, err := projectLogic.GetProjectList(c.logPoints, *req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	if c.GetString("exportType") != "" {
		fileName, downloadFileName := projectLogic.ExportProject(c.logPoints, res.Projects)
		//c.Ctx.Input.SetData("download", "yes")
		c.Ctx.Output.Download(fileName, downloadFileName)
		return
	} else {
		c.Res.Result = res
	}

}

// swagger:route GET /project/{project_id} project getProject
// GetProject 获取项目详情
//     Responses:
//       200: getProject
//       default: ErrorResponse

func (c *ProjectController) GetProject() {
	defer c.CatchException()
	ProjectId := c.Ctx.Input.Param(":project_id")
	res, err := projectLogic.GetProjectById(c.logPoints, ProjectId)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = res
}

// ModifyProject ...
// swagger:route POST /project/{project_id} project modifyProject
// ModifyProject 修改项目名称
//     Responses:
//       200: modifyProject
//       default: ErrorResponse

func (c *ProjectController) ModifyProject() {
	defer c.CatchException()
	ProjectId := c.Ctx.Input.Param(":project_id")
	req := &sdkModels.ModifyProjectRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyProject parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	// req.Validate(c.logPoints)
	if err := projectLogic.ModifyProject(c.logPoints, req, ProjectId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// ModifyProject ...
// swagger:route POST /project/{project_id}/description project modifyProjectDescription
// ModifyProject 修改项目描述
//     Responses:
//       200: modifyProjectDescription
//       default: ErrorResponse

func (c *ProjectController) ModifyProjectDescription() {
	defer c.CatchException()
	ProjectId := c.Ctx.Input.Param(":project_id")
	req := &sdkModels.ModifyProjectDescriptionRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyProject parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	// req.Validate(c.logPoints)
	if err := projectLogic.ModifyProjectDescription(c.logPoints, req, ProjectId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// swagger:route DELETE /project/{project_id} project deleteProject
// DeleteProject 删除项目
//     Responses:
//       200: deleteProject
//       default: ErrorResponse

func (c *ProjectController) DeleteProject() {
	defer c.CatchException()
	ProjectId := c.Ctx.Input.Param(":project_id")
	if err := projectLogic.DeleteProject(c.logPoints, ProjectId); err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /project/{project_id}/share project shareProject
// ShareProject 共享项目
//     Responses:
//       200: shareProject
//       default: ErrorResponse
func (c *ProjectController) ShareProject() {
	defer c.CatchException()
	ProjectId := c.Ctx.Input.Param(":project_id")
	req := requestTypes.ShareProjectRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.logPoints.Warn("ShareProjectRequest parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	// req.Validate(c.logPoints)
	if err := projectLogic.ShareProject(c.logPoints, ProjectId, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /project/{project_id}/cancelShare project cancelshareProject
// CancelShareProject 取消共享项目
//     Responses:
//       200: cancelshareProject
//       default: ErrorResponse
func (c *ProjectController) CancelShareProject() {
	defer c.CatchException()
	ProjectId := c.Ctx.Input.Param(":project_id")
	req := requestTypes.ShareProjectRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.logPoints.Warn("CancelShareProject parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	// req.Validate(c.logPoints)
	if err := projectLogic.CancelShareProject(c.logPoints, ProjectId, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /project/{project_id}/move project moveProject
// MoveProject 转移项目
//     Responses:
//       200: moveProject
//       default: ErrorResponse
func (c *ProjectController) MoveProject() {
	defer c.CatchException()
	ProjectId := c.Ctx.Input.Param(":project_id")
	req := requestTypes.MoveProjectRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.logPoints.Warn("MoveProjectRequest parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	// req.Validate(c.logPoints)
	if err := projectLogic.MoveProject(c.logPoints, ProjectId, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// swagger:route PUT /project/move/instances project moveInstances
//
// MoveInstances 转移项目下的某些实例
//     Responses:
//       200: moveInstances
//       default: ErrorResponse

func (c *ProjectController) MoveInstances() {
	defer c.CatchException()

	req := &requestTypes.MoveInstancesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("MoveProject parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	// req.Validate(c.logPoints)
	if err := projectLogic.MoveInstances(c.logPoints, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}
