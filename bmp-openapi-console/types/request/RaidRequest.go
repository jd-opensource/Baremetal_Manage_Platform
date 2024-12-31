package request

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type CreateRaidRequest struct {
	// raid名称
	// required: true
	Name string `json:"name" validate:"required,min=1,max=16"`
	// raid描述（中文）
	// required: true
	DescriptionEn string `json:"descriptionEn" validate:"required"`
	// raid描述（英文）
	// required: true
	DescriptionZh string `json:"descriptionZh" validate:"required"`
}

func (req *CreateRaidRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type ModifyRaidRequest struct {
	// raid名称
	// required: true
	Name string `json:"name" validate:"required,min=1,max=16"`
	// raid描述（中文）
	// required: true
	DescriptionEn string `json:"descriptionEn" validate:"required"`
	// raid描述（英文）
	// required: true
	DescriptionZh string `json:"descriptionZh" validate:"required"`
}

func (req *ModifyRaidRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
