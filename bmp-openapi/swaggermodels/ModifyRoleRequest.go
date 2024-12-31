package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyRole
type ModifyRoleRequest struct {
	ReadRequestHeader

	// in: path
	RoleID string `json:"role_id"`

	// in: body
	Body requestTypes.ModifyRoleRequest
}
