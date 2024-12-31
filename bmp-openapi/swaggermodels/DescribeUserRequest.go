package swaggermodels

// swagger:parameters describeUser
type DescribeUserRequest struct {
	ReadRequestHeader

	// in: path
	UserID string `json:"user_id"`
}
