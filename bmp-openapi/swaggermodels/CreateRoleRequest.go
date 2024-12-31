package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createRole
type CreateRoleRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateRoleRequest
}
