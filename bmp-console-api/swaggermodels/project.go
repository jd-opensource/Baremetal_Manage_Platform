package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

// swagger:parameters createProject
type CreateProjectRequest struct {
	WriteHeader

	// in:body
	Body sdkModels.CreateProjectRequest
}

// CreateProjectResponse is an response struct
// swagger:response createProject
type CreateProjectResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.ProjectId `json:"result"`
		RequestId string                  `json:"requestId"`
	}
}

// swagger:parameters getProjectList
type GetProjectListRequest struct {
	ReadHeader

	// in: query
	requestTypes.QueryProjectsRequest
}

// GetProjectListResponse is an response struct
// swagger:response getProjectList
type GetProjectListResponse struct {
	// in: body
	Body struct {
		Result    sdkModels.ProjectList `json:"result"`
		RequestId string                `json:"requestId"`
	}
}

// swagger:parameters getProject
type GetProjectRequest struct {
	ReadHeader

	// in: path
	ProjectID string `json:"project_id"`
}

// GetProjectResponse is an response struct
// swagger:response getProject
type GetProjectResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.Project `json:"result"`
		RequestId string                `json:"requestId"`
	}
}

// swagger:parameters modifyProject
type ModifyProjectRequest struct {
	WriteHeader

	// in: path
	ProjectID string `json:"project_id"`

	// in:body
	Body sdkModels.ModifyProjectRequest
}

// ModifyProjectResponse is an response struct
// swagger:response modifyProject
type ModifyProjectResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters modifyProjectDescription
type ModifyProjectDescriptionRequest struct {
	WriteHeader

	// in: path
	ProjectID string `json:"project_id"`

	// in:body
	Body sdkModels.ModifyProjectDescriptionRequest
}

// ModifyProjectDescriptionResponse is an response struct
// swagger:response modifyProjectDescription
type ModifyProjectDescriptionResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters deleteProject
type DeleteProjectRequest struct {
	WriteHeader

	// in: path
	ProjectID string `json:"project_id"`
}

// DeleteProjectResponse is an response struct
// swagger:response deleteProject
type DeleteProjectResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters shareProject
type ShareProjectRequest struct {
	WriteHeader

	// in: path
	ProjectID string `json:"project_id"`

	// in:body
	Body requestTypes.ShareProjectRequest
}

// ShareProjectResponse is an response struct
// swagger:response shareProject
type ShareProjectResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters cancelshareProject
type CancelShareProjectRequest struct {
	WriteHeader

	// in: path
	ProjectID string `json:"project_id"`

	// in:body
	Body requestTypes.ShareProjectRequest
}

// CancelShareProjectResponse is an response struct
// swagger:response cancelshareProject
type CancelShareProjectResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters moveProject
type MoveProjectRequest struct {
	WriteHeader

	// in: path
	ProjectID string `json:"project_id"`

	// in:body
	Body requestTypes.MoveProjectRequest
}

// MoveProjectResponse is an response struct
// swagger:response moveProject
type MoveProjectResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters moveInstances
type MoveInstancesRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.MoveInstancesRequest
}

// MoveInstancesResponse is an response struct
// swagger:response moveInstances
type MoveInstancesResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}
