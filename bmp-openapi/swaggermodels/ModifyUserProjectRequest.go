package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyUserProject
type ModifyUserProjectRequest struct {
	ReadRequestHeader

	// in: path
	ProjectID string `json:"project_id"`

	// in: body
	Body requestTypes.ModifyProjectRequest
}
