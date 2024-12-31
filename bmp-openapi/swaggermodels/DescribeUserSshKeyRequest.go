package swaggermodels

// swagger:parameters describeUserSshKey
type DescribeUserSshKeyRequest struct {
	ReadRequestHeader

	// in: path
	SshkeyID string `json:"sshkey_id"`
}
