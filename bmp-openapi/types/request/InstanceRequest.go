package request

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"

	log "coding.jd.com/aidc-bmp/bmp_log"
	validation "git.jd.com/cps-golang/ironic-common/ironic/common/Validation"
	"github.com/go-playground/validator/v10"
)

type CreateInstanceRequest struct {
	// 项目ID
	// required: true
	ProjectID string `json:"projectId" validate:"required"`
	// 机房ID
	// required: true
	IdcID string `json:"idcId" validate:"required"`
	// 设备类型ID
	// required: true
	DeviceTypeID string `json:"deviceTypeId" validate:"required"`
	// 镜像ID
	// required: true
	ImageID string `json:"imageId" validate:"required"`
	// 系统盘RAID ID
	// required: true
	SystemVolumeRaidID string `json:"systemVolumeRaidId" validate:"required"`
	// 系统盘分区
	SystemPartition []Partition `json:"systemPartition"`
	// 实例名称
	// required: true
	InstanceName string `json:"instanceName" validate:"required"`
	// 主机名称
	HostName string `json:"hostname"`
	// 密码
	Password string `json:"password"`
	// 秘钥ID
	SshKeyID string `json:"sshKeyId"`
	// 创建数量
	Count int `json:"count" validate:"gte=1"`
	// 描述
	Description string `json:"description" validate:"omitempty,max=256"`

	BootMode string `json:"boot_mode" validate:"omitempty,oneof=uefi bios"`
}

func (req *CreateInstanceRequest) Validate(logger *log.Logger) {
	validate(req, logger)
	//name条件：2-128 个字符，以大小写字母或中文开头，可包含数字、“.”、“_”、“:”或“-”
	specialMatch := regexp.MustCompile("^[.:\u4e00-\u9fa5_a-zA-Z0-9_-]{2,128}$").MatchString
	if !specialMatch(req.InstanceName) {
		panic(constant.BuildInvalidArgumentWithArgs("instanceName格式不正确", "Name invalidate"))
	}

	if req.HostName != "" {
		//hostname条件：长度为 2-64 个字符，允许使用点号(.)分隔字符成多段，每段允许使 用大小写字母、数字或连字符(-)，但不能连续使用点号(.)或连字符(-)。不能以点号(.)或连字符(-)开头或结尾
		specialMatch := regexp.MustCompile("^[.a-zA-Z0-9-]{2,64}$").MatchString
		if !specialMatch(req.HostName) {
			panic(constant.BuildInvalidArgumentWithArgs("hostName格式不正确", "hostName invalidate"))
		}
		if strings.HasPrefix(req.HostName, ".") || strings.HasPrefix(req.HostName, "-") || strings.HasSuffix(req.HostName, ".") || strings.HasSuffix(req.HostName, "-") || strings.Contains(req.HostName, "..") || strings.Contains(req.HostName, "--") {
			panic(constant.BuildInvalidArgumentWithArgs("hostName格式不正确", "hostName invalidate"))
		}

	}

	//if req.UserName != "" {
	//	specialMatch := regexp.MustCompile("^[a-zA-Z0-9_-]{1,64}$").MatchString
	//	if !specialMatch(req.UserName) {
	//		panic(constant.BuildInvalidArgumentWithArgs("userName格式不正确", "userName invalidate"))
	//	}
	//}

	if req.Password != "" {
		//密码复杂度校验，大写，小写，数字，特殊字符至少出现3种，且密码长度{8,30}
		lowercaseMatch := regexp.MustCompile(`[a-z]`).MatchString
		uppercaseMatch := regexp.MustCompile(`[A-Z]`).MatchString
		digitMatch := regexp.MustCompile(`\d`).MatchString
		//()`~!@#$%&*_-+=|{}[]:";'<>,.?/
		specialMatch := regexp.MustCompile("[()`~!@#$%&*\\_\\-+=|{}\\[\\]:\";'<>,.?/]").MatchString
		total := 0
		if lowercaseMatch(req.Password) {
			total += 1
		}
		if uppercaseMatch(req.Password) {
			total += 1
		}
		if digitMatch(req.Password) {
			total += 1
		}
		if specialMatch(req.Password) {
			total += 1
		}

		if total >= 3 && len(req.Password) >= 8 && len(req.Password) <= 30 {
			//pass
		} else {
			panic(constant.BuildInvalidArgumentWithArgs("password格式不正确", "password invalidate"))
		}
	}

	if len(req.SystemPartition) > 0 {
		fsTypes := []string{"swap", "xfs", "ext2", "ext3", "ext4"}
		for _, v := range req.SystemPartition {
			if !util.InArray(v.FsType, fsTypes) {
				logger.Warnf("CreateInstanceRequest.SystemPartition.format invalid:%s", v.FsType)
				panic(constant.BuildInvalidArgumentWithArgs("format格式不正确", "format invalidate"))
			}
			if !util.InArray(v.MountPoint, []string{"swap", "/boot", "/var", "/", "/export", "/data0", "/home"}) {
				logger.Warnf("CreateInstanceRequest.SystemPartition.MountPoint invalid:%s", v.MountPoint)
				panic(constant.BuildInvalidArgumentWithArgs("point格式不正确", "point invalidate"))
			}
			if v.Size <= 0 {
				logger.Warnf("CreateInstanceRequest.SystemPartition.Size invalid:%s", v.MountPoint)
				panic(constant.BuildInvalidArgumentWithArgs("size格式不正确", "size invalidate"))
			}
		}
	}

}

type ModifyInstanceRequest struct {
	// 实例名称
	InstanceName string `json:"instanceName"`
	// 描述
	// Extensions:
	// x-nullable: true
	Description *string `json:"description" validate:"omitempty,max=256"`
}

func (c *ModifyInstanceRequest) Validate(logger *log.Logger) {
	validate(c, logger)
	if c.InstanceName != "" {
		//name条件：2-128 个字符，以大小写字母或中文开头，可包含数字、“.”、“_”、“:”或“-”
		specialMatch := regexp.MustCompile("^[.:\u4e00-\u9fa5_a-zA-Z0-9_-]{2,128}$").MatchString
		if !specialMatch(c.InstanceName) {
			panic(constant.BuildInvalidArgumentWithArgs("instanceName格式不正确", "Name invalidate"))
		}
	}

}

// 批量修改实例名称
type ModifyInstancesRequest struct {
	// 实例名称
	InstanceName string `json:"instanceName" validate:"required"`
	// 实例id列表
	InstanceIDs []string `json:"instance_ids" validate:"required"`
}

func (c *ModifyInstancesRequest) Validate(logger *log.Logger) {
	validate(c, logger)
	if c.InstanceName != "" {
		//name条件：2-128 个字符，以大小写字母或中文开头，可包含数字、“.”、“_”、“:”或“-”
		specialMatch := regexp.MustCompile("^[.:\u4e00-\u9fa5_a-zA-Z0-9_-]{2,128}$").MatchString
		if !specialMatch(c.InstanceName) {
			panic(constant.BuildInvalidArgumentWithArgs("instanceName格式不正确", "Name invalidate"))
		}
	}

}

// 批量删除实例
type DeleteInstancesRequest struct {
	// 实例id列表
	InstanceIDs []string `json:"instance_ids" validate:"required"`
}

func (c *DeleteInstancesRequest) Validate(logger *log.Logger) {
	validate(c, logger)
}

type QueryInstancesRequest struct {
	Pageable
	// 实例ID
	InstanceID string `json:"instanceId"`
	// 实例名称,模糊搜索
	InstanceName string `json:"instanceName"`
	// 设备类型ID
	DeviceTypeID string `json:"deviceTypeId"`
	// 运行状态
	Status string `json:"status"`
	// 设备ID
	DeviceID string `json:"deviceId"`
	// SN
	Sn string `json:"sn"`
	// 机房ID
	IdcID string `json:"idcId"`
	// 项目ID
	ProjectID string `json:"projectId"`
	// 带外ip，精确搜索
	IloIP string `json:"ilo_ip"`
	// ipv4，精确搜索
	IPV4 string `json:"ipv4"`
	// ipv6，精确搜索
	IPV6 string `json:"ipv6"`
	// 是否显示全部，取值为1时表示全部
	IsAll string `json:"isAll"`
}

func (req *QueryInstancesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type RetryInstallInstanceRequest struct {
	Password string `json:"password"`
	UserData string `json:"user_data"`
	//AliasIps []AliasIP `json:"alias_ips"`
}
type RetryReinstallInstanceRequest struct {
	KeepData *bool  `json:"keep_data" validate:"required"`
	Password string `json:"password"`
	UserData string `json:"user_data"`
}

func (c *RetryReinstallInstanceRequest) Validate(logger *log.Logger) {
	validate(c, logger)
}

type RetryStartInstanceRequest struct {
	InstanceIds []string `json:"instance_ids" validate:"required"`
}

func (c *RetryStartInstanceRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(c); err != nil {
		logger.Warn("RetryStartInstanceRequest.Validate error:", err.Error())
		panic(constant.INVALID_ARGUMENT)
	}
	for _, instance_id := range c.InstanceIds {
		if match, _ := regexp.MatchString(validation.REGEX_ID, instance_id); !match {
			logger.Warn("RetryStartInstanceRequest.InstanceIds inValid:", c.InstanceIds)
			panic(constant.INVALID_ARGUMENT)
		}
	}
}

type StartInstancesRequest struct {
	InstanceIds []string `json:"instanceIds" validate:"required"`
}

func (req *StartInstancesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type StopInstancesRequest struct {
	InstanceIds []string `json:"instanceIds" validate:"required"`
}

func (req *StopInstancesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type RestartInstancesRequest struct {
	InstanceIds []string `json:"instanceIds" validate:"required"`
}

func (req *RestartInstancesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

// 重置实例密码
type ResetInstancePasswdRequest struct {
	// 新密码
	Password string `json:"password" validate:"required"`
}

func (req *ResetInstancePasswdRequest) Validate(logger *log.Logger) {
	validate(req, logger)

	if req.Password != "" {
		//密码复杂度校验，大写，小写，数字，特殊字符至少出现3种，且密码长度{8,30}
		lowercaseMatch := regexp.MustCompile(`[a-z]`).MatchString
		uppercaseMatch := regexp.MustCompile(`[A-Z]`).MatchString
		digitMatch := regexp.MustCompile(`\d`).MatchString
		//()`~!@#$%&*_-+=|{}[]:";'<>,.?/
		specialMatch := regexp.MustCompile("[()`~!@#$%&*\\_\\-+=|{}\\[\\]:\";'<>,.?/]").MatchString
		total := 0
		if lowercaseMatch(req.Password) {
			total += 1
		}
		if uppercaseMatch(req.Password) {
			total += 1
		}
		if digitMatch(req.Password) {
			total += 1
		}
		if specialMatch(req.Password) {
			total += 1
		}

		if total >= 3 && len(req.Password) >= 8 && len(req.Password) <= 30 {
			//pass
		} else {
			panic(constant.BuildInvalidArgumentWithArgs("password格式不正确", "password invalidate"))
		}
	}
}

// 批量重置实例密码
type ResetInstancesPasswdRequest struct {
	// 实例id列表
	InstanceIDs []string `json:"instance_ids" validate:"required"`
	// 新密码
	Password string `json:"password" validate:"required"`
}

func (req *ResetInstancesPasswdRequest) Validate(logger *log.Logger) {
	validate(req, logger)

	if req.Password != "" {
		//密码复杂度校验，大写，小写，数字，特殊字符至少出现3种，且密码长度{8,30}
		lowercaseMatch := regexp.MustCompile(`[a-z]`).MatchString
		uppercaseMatch := regexp.MustCompile(`[A-Z]`).MatchString
		digitMatch := regexp.MustCompile(`\d`).MatchString
		//()`~!@#$%&*_-+=|{}[]:";'<>,.?/
		specialMatch := regexp.MustCompile("[()`~!@#$%&*\\_\\-+=|{}\\[\\]:\";'<>,.?/]").MatchString
		total := 0
		if lowercaseMatch(req.Password) {
			total += 1
		}
		if uppercaseMatch(req.Password) {
			total += 1
		}
		if digitMatch(req.Password) {
			total += 1
		}
		if specialMatch(req.Password) {
			total += 1
		}

		if total >= 3 && len(req.Password) >= 8 && len(req.Password) <= 30 {
			//pass
		} else {
			panic(constant.BuildInvalidArgumentWithArgs("password格式不正确", "password invalidate"))
		}
	}
}

type ReinstallInstanceRequest struct {
	// 镜像ID
	// required: true
	ImageID string `json:"imageId" validate:"required"`
	// 系统盘RAID ID
	// required: true
	SystemVolumeRaidID string `json:"systemVolumeRaidId"` //panfu安装时，不传此字段
	// 系统盘分区
	SystemPartition []Partition `json:"systemPartition"`
	// 数据盘分区
	DataPartition []Partition `json:"dataPartition"`
	// 实例名称
	// required: true
	InstanceName string `json:"instanceName" validate:"required"`
	// 主机名称
	HostName string `json:"hostname"`
	// 密码
	Password string `json:"password" validate:"required"`
	// 秘钥ID，运营平台重装，密钥暂时不可用
	SshKeyID string `json:"sshKeyId"`
	// 描述
	Description string `json:"description" validate:"omitempty,max=256"`
	// 引导模式 [UEFI Legacy/BIOS]
	// required: true
	BootMode string `json:"bootMode"`

	//是否安装bmp-agent
	InstallBmpAgent bool `json:"installBmpAgent"`
}

func (req *ReinstallInstanceRequest) Validate(logger *log.Logger) error {
	if err := validator.New().Struct(req); err != nil {
		logger.Warn("ReinstallInstanceRequest.Validate error:", err.Error())
		return fmt.Errorf("parameter validate error:%s", err.Error())
	}

	if match, _ := regexp.MatchString(validation.REGEX_ID, req.ImageID); !match {
		logger.Warn("ReinstallInstanceRequest.ImageId invalid:", req.ImageID)
		return errors.New("parameter ImageId invalid")
	}

	b := strings.TrimSpace(req.BootMode)
	if b != "UEFI" && b != "Legacy/BIOS" {
		panic(constant.BuildInvalidArgumentWithArgs("boot_mode invalid", "boot_mode 非法"))
	}

	return nil

}
