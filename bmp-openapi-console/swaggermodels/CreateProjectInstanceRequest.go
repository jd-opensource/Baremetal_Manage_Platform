package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createProjectInstance
type CreateProjectInstanceRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateInstanceRequest
}
