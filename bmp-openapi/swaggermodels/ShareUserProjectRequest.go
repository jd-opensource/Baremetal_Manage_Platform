package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters shareUserProject
type ShareUserProjectRequest struct {
	WriteRequestHeader
	// in: path
	ProjectID string `json:"project_id"`
	// in: body
	Body requestTypes.ShareProjectRequest
}
