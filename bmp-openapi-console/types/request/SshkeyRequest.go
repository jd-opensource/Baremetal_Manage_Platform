package request

import (
	"regexp"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

const REGEX_RSA_PUBLIC_KEY = "^ssh-rsa AAAAB3NzaC1yc2.*"

type CreateSshkeyRequest struct {
	// 秘钥名称
	// required: true
	Name string `json:"name" validate:"required"`
	// 公钥内容，格式：^ssh-rsa AAAAB3NzaC1yc2.*
	// required: true
	Key string `json:"key" validate:"required"`
}
type ModifySshkeyRequest struct {
	// 秘钥名称
	// Extensions:
	// x-nullable: true
	Name *string `json:"name" validate:"omitempty,max=64"`
}

func (req *CreateSshkeyRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	//name条件：1~64 字符，只支持数字、大小写字母、英文下划线“_”及中划线
	specialMatch := regexp.MustCompile("^[a-zA-Z0-9_-]{1,64}$").MatchString
	if !specialMatch(req.Name) {
		panic(constant.BuildInvalidArgumentWithArgs("name格式不正确", "Name invalidate"))
	}
	if match, _ := regexp.MatchString(REGEX_RSA_PUBLIC_KEY, req.Key); !match {
		logger.Warn("CreateKeypairRequest.PublicKey invalid", req.Key)
		panic(constant.BuildInvalidArgumentWithArgs("公钥不合法", "publickey invalid"))
	}

}

func (req *ModifySshkeyRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	//name条件：1~64 字符，只支持数字、大小写字母、中英文下划线“_”及中划线
	specialMatch := regexp.MustCompile("^[a-zA-Z0-9_-]{1,64}$").MatchString
	if req.Name != nil && !specialMatch(*req.Name) {
		panic(constant.BuildInvalidArgumentWithArgs("name格式不正确", "Name invalidate"))
	}

}

type QuerySshkeysRequest struct {
	// 秘钥名称
	Name string `json:"name"` // 秘钥名称
	Pageable
	// 是否显示全部
	IsAll string `json:"isAll"`
}

func (req *QuerySshkeysRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
