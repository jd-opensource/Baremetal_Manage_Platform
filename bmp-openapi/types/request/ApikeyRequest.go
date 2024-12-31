package request

import (
	"regexp"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type CreateApikeyRequest struct {
	// 秘钥对名称
	// required: true
	Name string `json:"name" validate:"required"`
	// 是否支持只读，[0/1], read_only=1 的时候说明这个key是只读key，不能访问写方法
	// required: true
	ReadOnly int8 `json:"readOnly" validate:"oneof=0 1"`
	// Token类型, [system/user]
	// required: true
	Type string `json:"type" validate:"required,oneof=system user"`
}
type ModifyApikeyRequest struct {
	// 秘钥对名称
	// required: true
	Name string `json:"name" validate:"required"` // 秘钥对名称
	// 是否支持只读，[0/1], read_only=1 的时候说明这个key是只读key，不能访问写方法
	// required: true
	ReadOnly int8 `json:"readOnly" validate:"oneof=0 1"`
}
type QueryApikeysRequest struct {
	Pageable
	// 秘钥对名称
	Name string `json:"name"`
	// Token类型, [system/user]
	Type string `json:"type" validate:"omitempty,oneof=system user"`
	// 是否查询全部/导出
	IsAll string `json:"isAll"`
}

func (req *CreateApikeyRequest) Validate(logger *log.Logger) {

	validate(req, logger)

	if req.ReadOnly != 0 && req.ReadOnly != 1 {
		panic(constant.BuildInvalidArgumentWithArgs("readonly范围不正确", "readonly param error"))
	}

	//name条件：1~64 字符，只支持数字、大小写字母、中英文下划线“_”及中划线
	specialMatch := regexp.MustCompile("^[\u4e00-\u9fa5_a-zA-Z0-9_-]{1,64}$").MatchString
	if !specialMatch(req.Name) {
		panic(constant.BuildInvalidArgumentWithArgs("name格式不正确", "Name invalidate"))
	}

}

func (req *ModifyApikeyRequest) Validate(logger *log.Logger) {

	validate(req, logger)

	if req.ReadOnly != 0 && req.ReadOnly != 1 {
		panic(constant.BuildInvalidArgumentWithArgs("readonly范围不正确", "readonly param error"))
	}

	//name条件：1~64 字符，只支持数字、大小写字母、中英文下划线“_”及中划线
	specialMatch := regexp.MustCompile("^[\u4e00-\u9fa5_a-zA-Z0-9_-]{1,64}$").MatchString
	if !specialMatch(req.Name) {
		panic(constant.BuildInvalidArgumentWithArgs("name格式不正确", "Name invalidate"))
	}

}
func (req *QueryApikeysRequest) Validate(logger *log.Logger) {

	validate(req, logger)
	if req.Name != "" {
		//name条件：1~64 字符，只支持数字、大小写字母、中英文下划线“_”及中划线
		specialMatch := regexp.MustCompile("^[\u4e00-\u9fa5_a-zA-Z0-9_-]{1,64}$").MatchString
		if !specialMatch(req.Name) {
			panic(constant.BuildInvalidArgumentWithArgs("name格式不正确", "Name invalidate"))
		}
	}

}
