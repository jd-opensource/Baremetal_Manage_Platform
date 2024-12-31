package request

import (
	"fmt"
	"regexp"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"

	log "coding.jd.com/aidc-bmp/bmp_log"
)

type CreateUserRequest struct {

	// 角色uuid
	// required: true
	RoleID string `json:"roleId"`
	// 用户名，唯一, 1~64字符，只支持数字、大小写字母、英文下划线“_”及中划线“-”
	// required: true
	UserName string `json:"userName" validate:"required,min=1,max=64"`
	// 邮箱
	Email string `json:"email" validate:"required,email"`
	// 国家地区码，如86
	PhonePrefix string `json:"phonePrefix" validate:"required,min=1,max=64"`
	// 手机号
	PhoneNumber string `json:"phoneNumber" validate:"required,min=1,max=64"`
	// 用户默认项目uuid
	DefaultProjectID string `json:"defaultProjectId"`
	// 语言[en_US, zh_CN]
	Language string `json:"language" validate:"omitempty,oneof=en_US zh_CN"`
	// 时区，Asia/Shanghai
	Timezone string `json:"timezone"`
	// 密码 明文 密码复杂度校验，大写，小写，数字，特殊字符至少出现3种，且密码长度{8,30}
	// required: true
	Password string `json:"password"`
	// 描述
	Description string `json:"description" validate:"max=256"`
}

func (req *CreateUserRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	//matched, err := regexp.MatchString(`^[a-zA-Z-_0-9]{1,32}$`, req.UserName)
	//if err != nil {
	//	logger.Warn("CreateUserRequest.Validate error:", err.Error())
	//	panic(constant.BuildInvalidArgumentWithArgs(err.Error(), err.Error()))
	//}
	//if !matched {
	//	logger.Warn("CreateUserRequest.Validate userName error")
	//	panic(constant.BuildInvalidArgumentWithArgs("用户名格式不正确", "userName invalidate"))
	//}

	matchedPhone, err := regexp.MatchString(`^[0-9]+$`, req.PhoneNumber)
	if err != nil {
		logger.Warn("CreateUserRequest.Validate error:", err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(err.Error(), err.Error()))
	}
	if !matchedPhone {
		logger.Warn("CreateUserRequest.Validate phoneNumber error")
		panic(constant.BuildInvalidArgumentWithArgs("手机格式不正确", "phoneNumber invalidate"))
	}

	if req.Timezone == "" {
		req.Timezone = "Asia/Shanghai"
	}
	passwd := req.Password
	//密码复杂度校验，大写，小写，数字，特殊字符至少出现3种，且密码长度{8,30}
	lowercaseMatch := regexp.MustCompile(`[a-z]`).MatchString
	uppercaseMatch := regexp.MustCompile(`[A-Z]`).MatchString
	digitMatch := regexp.MustCompile(`\d`).MatchString
	//()`~!@#$%&*_-+=|{}[]:";'<>,.?/
	specialMatch := regexp.MustCompile("[()`~!@#$%&*\\_\\-+=|{}\\[\\]:\";'<>,.?/]").MatchString
	total := 0
	if lowercaseMatch(passwd) {
		total += 1
	}
	if uppercaseMatch(passwd) {
		total += 1
	}
	if digitMatch(passwd) {
		total += 1
	}
	if specialMatch(passwd) {
		total += 1
	}

	if total >= 3 && len(passwd) >= 8 && len(passwd) <= 30 {
		return
	}
	panic(constant.BuildInvalidArgumentWithArgs("密码格式不正确", "password invalidate"))
}

type ModifyUserRequest struct {
	// 邮箱
	// Extensions:
	// x-nullable: true
	Email *string `json:"email" validate:"omitempty,email"`
	// 国家地区码，如86
	// Extensions:
	// x-nullable: true
	PhonePrefix *string `json:"phonePrefix" validate:"omitempty,min=1,max=64"`
	// 手机号
	// Extensions:
	// x-nullable: true
	PhoneNumber *string `json:"phoneNumber" validate:"omitempty,min=1,max=64"`
	// 描述
	// Extensions:
	// x-nullable: true
	Description *string `json:"description" validate:"omitempty,min=0,max=256"`
}

func (req *ModifyUserRequest) Validate(logger *log.Logger) {

	validate(req, logger)
	//matched, err := regexp.MatchString(`^[a-zA-Z-_0-9]{1,32}$`, req.UserName)
	//if err != nil {
	//	logger.Warn("ModifyUserRequest.Validate error:", err.Error())
	//	panic(constant.BuildInvalidArgumentWithArgs(err.Error(), err.Error()))
	//}
	//if !matched {
	//	logger.Warn("ModifyUserRequest.Validate userName error")
	//	panic(constant.BuildInvalidArgumentWithArgs("用户名格式不正确", "userName invalidate"))
	//}

	if req.PhoneNumber != nil {
		matchedPhone, err := regexp.MatchString(`^[0-9]+$`, *req.PhoneNumber)
		if err != nil {
			logger.Warn("ModifyUserRequest.Validate error:", err.Error())
			panic(constant.BuildInvalidArgumentWithArgs(err.Error(), err.Error()))
		}
		if !matchedPhone {
			logger.Warn("ModifyUserRequest.Validate phoneNumber error")
			panic(constant.BuildInvalidArgumentWithArgs("手机格式不正确", "phoneNumber invalidate"))
		}
	}

}

type ModifyUserPasswordRequest struct {
	// 旧密码，明文
	// required: true
	OldPassword string `json:"oldPassword"`
	// 新密码 明文，密码复杂度校验，大写，小写，数字，特殊字符至少出现3种，且密码长度{8,30}
	// required: true
	Password string `json:"password"`
}

func (req *ModifyUserPasswordRequest) Validate(logger *log.Logger) {

	passwd := req.Password
	//密码复杂度校验，大写，小写，数字，特殊字符至少出现3种，且密码长度{8,30}
	lowercaseMatch := regexp.MustCompile(`[a-z]`).MatchString
	uppercaseMatch := regexp.MustCompile(`[A-Z]`).MatchString
	digitMatch := regexp.MustCompile(`\d`).MatchString
	//()`~!@#$%&*_-+=|{}[]:";'<>,.?/
	specialMatch := regexp.MustCompile("[()`~!@#$%&*\\_\\-+=|{}\\[\\]:\";'<>,.?/]").MatchString
	total := 0
	if lowercaseMatch(passwd) {
		total += 1
	}
	if uppercaseMatch(passwd) {
		total += 1
	}
	if digitMatch(passwd) {
		total += 1
	}
	if specialMatch(passwd) {
		total += 1
	}

	if total >= 3 && len(passwd) >= 8 && len(passwd) <= 30 {
		return
	}
	panic(constant.BuildInvalidArgumentWithArgs("密码格式不正确", "password invalidate"))

}

type ModifyLocalUserRequest struct {
	// 邮箱
	// Extensions:
	// x-nullable: true
	Email *string `json:"email" validate:"omitempty,email"`
	// 国家地区码，如86
	// Extensions:
	// x-nullable: true
	PhonePrefix *string `json:"phonePrefix" validate:"omitempty,min=0,max=64"`
	// 手机号
	// Extensions:
	// x-nullable: true
	PhoneNumber *string `json:"phoneNumber" validate:"omitempty,min=0,max=64"`
	// 用户默认项目uuid
	// Extensions:
	// x-nullable: true
	DefaultProjectID *string `json:"defaultProjectId" validate:"omitempty,min=0,max=64"`
	// 语言[en_US, zh_CN]
	// Extensions:
	// x-nullable: true
	Language *string `json:"language" validate:"omitempty,oneof=en_US zh_CN"` // 语言（中文/English）
	// 时区 Asia/Shanghai
	// Extensions:
	// x-nullable: true
	Timezone *string `json:"timezone" validate:"omitempty,min=0,max=64"`
}

func (req *ModifyLocalUserRequest) Validate(logger *log.Logger) {

	validate(req, logger)
	if req.PhonePrefix != nil {
		var phoneReg string
		if *req.PhonePrefix == "086" { //中国大陆手机号
			phoneReg = `^(?:(?:\+|00)86)?1(?:(?:3[\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[189]))\d{8}$`
		} else if *req.PhonePrefix == "852" { //中国香港
			phoneReg = `^([5|6|9])\d{7}$`
		} else if *req.PhonePrefix == "853" { //中国澳门
			phoneReg = `^6\d{7}$`
		} else if *req.PhonePrefix == "886" { //中国台湾
			phoneReg = `^[0][9]\d{8}$`
		} else {
			phoneReg = `^[0-9]{1,}$`
		}

		fmt.Println("dddd", *req.PhonePrefix)
		if req.PhoneNumber != nil {
			matchedPhone, err := regexp.MatchString(phoneReg, *req.PhoneNumber)
			if err != nil {
				logger.Warn("ModifyUserRequest.Validate error:", err.Error())
				panic(constant.BuildInvalidArgumentWithArgs(err.Error(), err.Error()))
			}
			if !matchedPhone {
				logger.Warn("ModifyUserRequest.Validate phoneNumber error")
				panic(constant.BuildInvalidArgumentWithArgs("手机格式不正确", "phoneNumber invalidate"))
			}
		}
	}
}

type GetUserByNameRequest struct {

	// 用户名，唯一
	// required: true
	UserName string `json:"userName"`
}

func (req *GetUserByNameRequest) Validate(logger *log.Logger) {
	//matched, err := regexp.MatchString(`^[a-zA-Z-_0-9]{1,32}$`, req.UserName)
	//if err != nil {
	//	logger.Warn("ModifyUserRequest.Validate error:", err.Error())
	//	panic(constant.INVALID_ARGUMENT)
	//}
	//if !matched {
	//	logger.Warn("ModifyUserRequest.Validate projectName error")
	//	panic(constant.INVALID_ARGUMENT)
	//}
}

type QueryUsersRequest struct {
	// 角色uuid
	RoleID string `json:"roleId"`
	// 项目uuid
	DefaultProjectID string `json:"defaultProjectId"`
	// 用户名
	UserName string `json:"userName"`
	Pageable
	// 是否显示全部, isAll=1表示全部
	IsAll string `json:"isAll"`
}

func (req *QueryUsersRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type VerifyUserRequest struct {
	// 用户名，唯一
	// required: true
	UserName string `json:"userName" validate:"required"` //
	// 密码
	// required: true
	Password string `json:"password" validate:"required"` //
	// 角色uuid
	// required: true
	RoleID string `json:"roleId" validate:"required"`
}

func (req *VerifyUserRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
