package swaggermodels

// swagger:parameters getInstancesBySshkey
type GetInstancesBySshkey struct {
	ReadRequestHeader

	// in: path
	SshkeyID string `json:"sshkey_id"`
}
