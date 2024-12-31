package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters moveUserProject
type MoveProjectRequest struct {
	WriteRequestHeader
	// in: path
	ProjectID string `json:"project_id"`
	// in: body
	Body requestTypes.MoveProjectRequest
}
