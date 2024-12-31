package swaggermodels

// swagger:parameters deleteUserSshkey
type DeleteUserSshkeyRequest struct {
	WriteRequestHeader

	// in: path
	SshkeyID string `json:"sshkey_id"`
}
