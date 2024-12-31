package request

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type CreateOSRequest struct {
	// 操作系统名称
	// required: true
	OsName string `json:"osName" validate:"required,min=1,max=200"`
	// 操作系统平台
	// required: true
	OsType string `json:"osType" validate:"required"`
	// 体系架构
	// required: true
	Architecture string `json:"architecture" validate:"required"`
	// 位数
	// required: true
	Bits int `json:"bits" validate:"required"`
	// 操作系统版本
	// required: true
	OsVersion string `json:"osVersion" validate:"required"`
	// 系统用户
	SysUser string `json:"sysUser"`
}

func (req *CreateOSRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type ModifyOSRequest struct {
	// 操作系统名称
	OsName string `json:"osName"`
	// 操作系统平台
	OsType string `json:"osType"`
	// 体系架构
	Architecture string `json:"architecture"`
	// 位数
	Bits int `json:"bits" `
	// 操作系统版本
	OsVersion string `json:"osVersion"`
	// 系统用户
	SysUser string `json:"sysUser"`
}

func (req *ModifyOSRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type QueryOssRequest struct {
	// 操作系统名称
	OsName string `json:"osName"`
	// 操作系统平台
	OsType string `json:"osType"`
	// 操作系统版本
	OsVersion string `json:"osVersion"`
	// 是否显示所有
	IsAll string `json:"isAll"`
}

func (req *QueryOssRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
