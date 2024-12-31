package swaggermodels

// swagger:parameters describeRole
type DescribeRoleRequest struct {
	ReadRequestHeader

	// in: path
	RoleID string `json:"role_id"`
}
