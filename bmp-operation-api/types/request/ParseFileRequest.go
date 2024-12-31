package request

import "mime/multipart"

type ParseFileRequest struct {
	BaseRequest
	Region       string         `json:"region" valid:"Required"`
	Base64       bool           `json:"base64" valid:"Required"`
	UserDataFile multipart.File `json:"userDataFile" valid:"Required"`
}
