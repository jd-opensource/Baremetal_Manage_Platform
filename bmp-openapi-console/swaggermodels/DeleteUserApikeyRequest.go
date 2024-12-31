package swaggermodels

// swagger:parameters deleteUserApikey
type DeleteUserApikeyRequest struct {
	WriteRequestHeader

	// in: path
	ApikeyID string `json:"apikey_id"`
}
