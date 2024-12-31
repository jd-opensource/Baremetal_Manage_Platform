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

type ModifyProjectDescriptionRequest struct {
	// 项目描述
	// required: true
	Description string `json:"description"`
}

func (req *ModifyProjectDescriptionRequest) Validate(logger *log.Logger) {
	validate(req, logger)

}

type QueryProjectsRequest struct {
	// 项目名称 1~64字符，只支持数字、大小写字母、英文下划线“_”及中划线“-”
	ProjectName string `json:"projectName"`
	//0表示全部，1表示拥有者，2表示被共享者
	Owned int `json:"owned"`
	Pageable
	IsAll string `json:"isAll"`
	//按创建时间asc/desc排列
	OrderByCreatetime string `json:"orderByCreatetime"`
	//新增可选，按用户名查找
	OwnerName string `json:"ownerName"`
	//新增可选，按被转移用户的用户名查找它拥有的项目
	SharerName string `json:"sharerName"`
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
	// 如果部分分享，instance_id逗号分隔； 如果全部分享，传all
	// required: true
	InstanceIDs string `json:"instanceIDs"` //all表示项目下的全部资源,否则instance_id逗号分隔传过来
}

func (req *ShareProjectRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	//name条件：1~64 字符，只支持数字、大小写字母、中英文下划线“_”及中划线
}

type CalcelShareProjectRequest struct {
	// from user_id
	// required: true
	OwnerID string `json:"ownerID" validate:"required"`
	// to user_id
	// required: true
	SharerID string `json:"sharerID" validate:"required"`
}

func (req *CalcelShareProjectRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	//name条件：1~64 字符，只支持数字、大小写字母、中英文下划线“_”及中划线
}

type MoveProjectRequest struct {
	// from user_id
	// required: true
	OwnerID string `json:"ownerID" validate:"required"`
	// to user_id
	// required: true
	MoverID string `json:"moverID" validate:"required"`
	// 【read/write】
	// required: true
	// Premisson string `json:"premission" validate:"required,oneof=read write"`
}

func (req *MoveProjectRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	//name条件：1~64 字符，只支持数字、大小写字母、中英文下划线“_”及中划线
}

type MoveInstancesRequest struct {
	// instanceIDs
	// required: true
	InstanceIDs string `json:"instanceIDs" validate:"required"`
	// from project_id
	// required: true
	OwnerProjectID string `json:"ownerProjectID" validate:"required"`
	// to project_id
	// required: true
	MoverProjectID string `json:"moverProjectID" validate:"required"`

	// from user_id
	// required: true
	OwnerID string `json:"ownerID" validate:"required"`

	// to user_id
	// required: true
	MoverID string `json:"moverID" validate:"required"`

	// 【read/write】
	// required: true
	// Premisson string `json:"premission" validate:"required,oneof=read write"`
}

func (req *MoveInstancesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	//name条件：1~64 字符，只支持数字、大小写字母、中英文下划线“_”及中划线
}
