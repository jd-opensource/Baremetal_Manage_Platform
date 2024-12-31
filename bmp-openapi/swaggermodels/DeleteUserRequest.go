package swaggermodels

// swagger:parameters deleteUser
type DeleteUserRequest struct {
	WriteRequestHeader

	// in: path
	UserID string `json:"user_id"`
}
