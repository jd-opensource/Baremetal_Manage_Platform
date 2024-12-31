package request

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"regexp"
)

type CreateIdcRequest struct {
	// 机房名称 1-64
	// required: true
	Name string `json:"name" validate:"required,min=1,max=64"`
	// 机房英文名称 1-64
	// required: true
	NameEn string `json:"nameEn" validate:"required,min=1,max=64"`
	// 机房缩写 1-64
	// required: true
	Shortname string `json:"shortname" validate:"required,min=1,max=64"`
	// 机房等级 1-64
	// required: true
	Level string `json:"level" validate:"required,min=1,max=64"`
	// 机房地址 0-256
	Address string `json:"address" validate:"max=256"`
	// 带外用户名 0-64
	IloUser string `json:"iloUser" validate:"max=64"`
	// 带外密码 0-64
	IloPassword string `json:"iloPassword" validate:"max=64"`
	// 交换机用户1 0-64
	SwitchUser1 string `json:"switchUser1" validate:"max=64"`
	// 交换机密码1 0-64
	SwitchPassword1 string `json:"switchPassword1" validate:"max=64"`
	// 交换机用户2 0-64
	SwitchUser2 string `json:"switchUser2" validate:"max=64"`
	// 交换机密码2 0-64
	SwitchPassword2 string `json:"switchPassword2" validate:"max=64"`
}

func (req *CreateIdcRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	specialMatch := regexp.MustCompile("^[a-zA-Z0-9_`{};:\"!~#$%@^&*()+-]+$").MatchString
	if !specialMatch(req.NameEn) {
		panic(constant.BuildInvalidArgumentWithArgs("机房英文名称不合法", "nameEn invalidate"))
	}
}

type ModifyIdcRequest struct {
	// 机房名称 1-64
	// Extensions:
	// x-nullable: true
	Name *string `json:"name" validate:"omitempty,min=1,max=64"`
	// 机房英文名称 1-64
	// Extensions:
	// x-nullable: true
	NameEn *string `json:"nameEn" validate:"omitempty,min=1,max=64"`
	// 机房缩写 1-64
	// Extensions:
	// x-nullable: true
	Shortname *string `json:"shortname" validate:"omitempty,min=1,max=64"`
	// 机房等级 1-64
	// Extensions:
	// x-nullable: true
	Level *string `json:"level" validate:"omitempty,min=1,max=64"`
	// 机房地址 0-256
	// Extensions:
	// x-nullable: true
	Address *string `json:"address" validate:"omitempty,min=0,max=256"`
	// 带外用户名 0-64
	// Extensions:
	// x-nullable: true
	IloUser *string `json:"iloUser" validate:"omitempty,min=0,max=64"`
	// 带外密码 0-64
	// Extensions:
	// x-nullable: true
	IloPassword *string `json:"iloPassword" validate:"omitempty,min=0,max=64"`
	// 交换机用户1 0-64
	// Extensions:
	// x-nullable: true
	SwitchUser1 *string `json:"switchUser1" validate:"omitempty,min=0,max=64"`
	// 交换机密码1 0-64
	// Extensions:
	// x-nullable: true
	SwitchPassword1 *string `json:"switchPassword1" validate:"omitempty,min=0,max=64"`
	// 交换机用户2 0-64
	// Extensions:
	// x-nullable: true
	SwitchUser2 *string `json:"switchUser2" validate:"omitempty,min=0,max=64"`
	// 交换机2密码 0-64
	// Extensions:
	// x-nullable: true
	SwitchPassword2 *string `json:"switchPassword2" validate:"omitempty,min=0,max=64"`
}

func (req *ModifyIdcRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	specialMatch := regexp.MustCompile("^[a-zA-Z0-9_`{};:\"!~#$%@^&*()+-]+$").MatchString
	if req.NameEn != nil && !specialMatch(*req.NameEn) {
		panic(constant.BuildInvalidArgumentWithArgs("机房英文名称不合法", "nameEn invalidate"))
	}
	//对于必填项校验方法
	//if param.Name != nil && (len(*param.Name) < 1 || len(*param.Name) > 64) {
	//	panic(constant.BuildInvalidArgumentWithArgs("机房名称不合法", "idcName invalidate"))
	//}
	//fmt.Println(len(*param.NameEn))
	//if param.NameEn != nil && (len(*param.NameEn) < 1 || len(*param.NameEn) > 64) {
	//	panic(constant.BuildInvalidArgumentWithArgs("机房英文名称不合法", "idc nameEn invalidate"))
	//}
	//if param.Shortname != nil && (len(*param.Shortname) < 1 || len(*param.Shortname) > 64) {
	//	panic(constant.BuildInvalidArgumentWithArgs("机房名称缩写不合法", "idc shortName invalidate"))
	//}
	//if param.Level != nil && (len(*param.Level) < 1 || len(*param.Level) > 64) {
	//	panic(constant.BuildInvalidArgumentWithArgs("机房等级不合法", "idc level invalidate"))
	//}
	////非必填校验方法
	//if param.Address != nil && len(*param.Address) > 256 {
	//	panic(constant.BuildInvalidArgumentWithArgs("地址长度不合法", "idc address invalidate"))
	//}
	//if param.IloUser != nil && len(*param.IloUser) > 64 {
	//	panic(constant.BuildInvalidArgumentWithArgs("带外用户名长度不合法", "IloUser invalidate"))
	//}
	//if param.IloPassword != nil && len(*param.IloPassword) > 64 {
	//	panic(constant.BuildInvalidArgumentWithArgs("带外密码长度不合法", "IloPassword invalidate"))
	//}
	//if param.SwitchUser1 != nil && len(*param.SwitchUser1) > 64 {
	//	panic(constant.BuildInvalidArgumentWithArgs("SwitchUser1 长度不合法", "SwitchUser1 invalidate"))
	//}
	//if param.SwitchPassword1 != nil && len(*param.SwitchPassword1) > 64 {
	//	panic(constant.BuildInvalidArgumentWithArgs("SwitchPassword1 长度不合法", "SwitchPassword1 invalidate"))
	//}
	//if param.SwitchUser2 != nil && len(*param.SwitchUser2) > 64 {
	//	panic(constant.BuildInvalidArgumentWithArgs("SwitchUser2 长度不合法", "SwitchUser2 invalidate"))
	//}
	//if param.SwitchPassword2 != nil && len(*param.SwitchPassword2) > 64 {
	//	panic(constant.BuildInvalidArgumentWithArgs("SwitchPassword2 长度不合法", "SwitchPassword2 invalidate"))
	//}
}

type QueryIdcsRequest struct {
	// 机房名称
	Name string `json:"name"`
	// 机房英文名称
	NameEn string `json:"nameEn"`
	// 机房等级
	Level string `json:"level"`
	Pageable
	// 是否显示所有
	IsAll string `json:"isAll"`
}

func (req *QueryIdcsRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type QueryIdcRequest struct {
	// 机房ID
	// required: true
	IDcID string `json:"idcId" validate:"required"`
}

func (req *QueryIdcRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
