package swaggermodels

// swagger:parameters deleteRole
type DeleteRoleRequest struct {
	WriteRequestHeader

	// in: path
	RoleID string `json:"role_id"`
}
