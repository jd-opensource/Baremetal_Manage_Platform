package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createUserProject
type CreateUserProjectRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateProjectRequest
}
