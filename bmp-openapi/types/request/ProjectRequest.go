package request

import (
	"regexp"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"

	log "coding.jd.com/aidc-bmp/bmp_log"
)

type CreateProjectRequest struct {
	// ProjectName 1~64字符，只支持数字、大小写字母、中英文下划线“_”及中划线“-”
	// required: true
	ProjectName string `json:"projectName" validate:"required"` //项目名称
	// 是否作为该用户的默认项目, [0/1], 默认为0
	IsDefault int8 `json:"isDefault" validate:"omitempty,oneof=0 1"`
	// 是否作为系统项目, [0/1], 默认为0
	IsSystem int8 `json:"isSystem" validate:"omitempty,oneof=0 1"`
}

func (req *CreateProjectRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	//name条件：1~64 字符，只支持数字、大小写字母、中英文下划线“_”及中划线
	specialMatch := regexp.MustCompile("^[\u4e00-\u9fa5_a-zA-Z0-9_-]{1,64}$").MatchString
	if !specialMatch(req.ProjectName) {
		panic(constant.BuildInvalidArgumentWithArgs("projectName格式不正确", "Name invalidate"))
	}

}

type ModifyProjectRequest struct {
	// 项目名称 1~64字符，只支持数字、大小写字母、英文下划线“_”及中划线“-”
	// required: true
	ProjectName string `json:"projectName" validate:"required"`
}

func (req *ModifyProjectRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	//name条件：1~64 字符，只支持数字、大小写字母、中英文下划线“_”及中划线
	specialMatch := regexp.MustCompile("^[\u4e00-\u9fa5_a-zA-Z0-9_-]{1,64}$").MatchString
	if !specialMatch(req.ProjectName) {
		panic(constant.BuildInvalidArgumentWithArgs("projectName格式不正确", "Name invalidate"))
	}

}

type QueryProjectsRequest struct {
	// 项目名称 1~64字符，只支持数字、大小写字母、英文下划线“_”及中划线“-”
	ProjectName string `json:"projectName"`
	Owned       int    `json:"owned"` //0表示全部，1表示拥有者，2表示被共享者
	Pageable
	IsAll string `json:"isAll"`
}

func (req *QueryProjectsRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type ShareProjectRequest struct {
	// from user_id
	// required: true
	OwnerID string `json:"ownerID" validate:"required"`
	// to user_id
	// required: true
	SharerID string `json:"sharerID" validate:"required"`
}

func (req *ShareProjectRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	//name条件：1~64 字符，只支持数字、大小写字母、中英文下划线“_”及中划线
}
