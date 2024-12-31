package request

import (
	"git.jd.com/cps-golang/ironic-common/exception"
	log "git.jd.com/cps-golang/log"
	"github.com/go-playground/validator/v10"
)

type CreateOSRequest struct {
	OsName string `json:"osName" validate:"required,min=1,max=200"`
	OsType string `json:"osType" validate:"required"`

	Architecture string `json:"architecture" validate:"required"`
	Bits         int    `json:"bits" validate:"required"`
	OsVersion    string `json:"osVersion" validate:"required"`
	SysUser      string `json:"sysUser"`
}

func (req *CreateOSRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(req); err != nil {
		logger.Warn("CreateOSRequest.Validate error:", err.Error())
		panic(exception.CommonParamValid)
	}
}

type ModifyOSRequest struct {
	OsName string `json:"osName" `
	OsType string `json:"osType" `

	Architecture string `json:"architecture"`
	Bits         int    `json:"bits" `
	OsVersion    string `json:"osVersion"`
	SysUser      string `json:"sysUser"`
}

func (c *ModifyOSRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(c); err != nil {
		logger.Warn("ModifyOSRequest.Validate error:", err.Error())
		panic(exception.CommonParamValid)
	}
}

type QueryOssRequest struct {
	OsName       string `json:"osName"`       // 操作系统名称
	Architecture string `json:"architecture"` //体系架构
	OsType       string `json:"osType"`       // 操作系统分类:linux/windows
	Platform     string `json:"platform"`     // suse/centos/ubuntu
	OsVersion    string `json:"osVersion"`    // 操作系统版本
	IsAll        string `json:"isAll"`
}

func (req *QueryOssRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(req); err != nil {
		logger.Warn("QueryOssRequest.Validate error:", err.Error())
		panic(exception.CommonParamValid)
	}
	//for _, v := range req.Ids {
	//	if match, _ := regexp.MatchString(validation.REGEX_ID, v); !match {
	//		logger.Warn("QueryOssRequest.Ids invalid:", req.Ids)
	//		panic(exception.CommonParamValid)
	//	}
	//}
}
