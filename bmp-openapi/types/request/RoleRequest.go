package request

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type CreateRoleRequest struct {
	// 角色英文名称
	// required: true
	RoleNameEn string `json:"roleNameEn" validate:"required"`
	// 角色中文名称
	// required: true
	RoleNameCn string `json:"roleNameCn" validate:"required"`
	// 角色描述英文名称
	// required: true
	DescriptionEn string `json:"descriptionEn" validate:"required"`
	// 角色描述中文名称
	// required: true
	DescriptionCn string `json:"descriptionCn" validate:"required"`
}
type ModifyRoleRequest struct {
	// 角色英文名称
	RoleNameEn string `json:"roleNameEn"`
	// 角色中文名称
	RoleNameCn string `json:"roleNameCn"`
	// 角色描述英文名称
	DescriptionEn string `json:"descriptionEn"`
	// 角色描述中文名称
	DescriptionCn string `json:"descriptionCn"`
}

func (req *CreateRoleRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type QueryRolesRequest struct {
	Pageable
	// 是否显示所有
	IsAll string `json:"isAll"`
}

func (req *QueryRolesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
