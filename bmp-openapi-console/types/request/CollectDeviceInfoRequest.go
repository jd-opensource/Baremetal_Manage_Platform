package request

import (
	"net"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"

	log "coding.jd.com/aidc-bmp/bmp_log"
	validator "github.com/go-playground/validator/v10"
)

type CollectDeviceInfoRequest struct {
	// 待采集项
	// required: true
	Collects []CollectItem `json:"collects" validate:"required"`
}

type CollectItem struct {
	// sn
	// required: true
	Sn string `json:"sn"`
	// 带外IP
	// required: true
	IloIP string `json:"iloIp"` // 带外管理IP
	// 带外账号
	// required: true
	IloUser string `json:"iloUser"` // 带外账号
	// 带外密码
	// required: true
	IloPassword string `json:"iloPassword"` // 带外账号密码
	// MAC1（eth0）
	// required: true
	Mac1 string `json:"mac1"` // MAC1（eth0）
	// 内网IPV4
	// required: true
	PrivateIPv4 string `json:"privateIpv4"` // 内网IPV4

	// 掩码
	// required: true
	Mask string `json:"mask"`

	// 网关地址
	// required: true
	Gateway string `json:"gateway"`

	// raid_driver可选参数。传参：sas2ircu，sas3ircu，megacli64，storcli64，perccli64，no_raid。不传参数则只可采集已适配机型，如果未适配则采集失败
	// Extensions:
	// x-nullable: true
	RaidDriver string `json:"raidDriver"`
	// 可选参数。 传参:True, False 。True表示传入的raid_driver值将覆盖已适配机器的raid_driver
	// Extensions:
	// x-nullable: true
	AllowOverride bool `json:"allowOverride"`
}

func (c *CollectDeviceInfoRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(c); err != nil {
		logger.Warn("CollectDeviceInfoRequest Validate error:", err.Error())
		panic(constant.INVALID_ARGUMENT)
	}
	if c.Collects != nil {
		for _, v := range c.Collects {
			if v.RaidDriver != "" {
				arrs := []string{"sas2ircu", "sas3ircu", "megacli64", "storcli64", "perccli64", "no_raid", "arcconf"}
				if !util.InArray(v.RaidDriver, arrs) {
					logger.Warn("CollectDeviceInfoRequest Validate error:raid_driver invalid")
					panic(constant.BuildInvalidArgumentWithArgs("raid_driver invalid", "raid_driver 非法"))
				}
			}
			if net.ParseIP(v.IloIP) == nil {
				logger.Warn("CollectDeviceInfoRequest Validate error:iloip invalid")
				panic(constant.BuildInvalidArgumentWithArgs("iloip invalid", "iloip 非法"))
			}
			if net.ParseIP(v.Mask) == nil {
				logger.Warn("CollectDeviceInfoRequest Validate error:mask invalid")
				panic(constant.BuildInvalidArgumentWithArgs("mask invalid", "mask 非法"))
			}
			if net.ParseIP(v.Gateway) == nil {
				logger.Warn("CollectDeviceInfoRequest Validate error:gateway invalid")
				panic(constant.BuildInvalidArgumentWithArgs("gateway invalid", "gateway 非法"))
			}
			if v.IloUser == "" {
				logger.Warn("CollectDeviceInfoRequest Validate error:ilouser invalid")
				panic(constant.BuildInvalidArgumentWithArgs("ilouser invalid", "ilouser 非法"))
			}
			if v.IloPassword == "" {
				logger.Warn("CollectDeviceInfoRequest Validate error:ilopassword invalid")
				panic(constant.BuildInvalidArgumentWithArgs("ilopassword invalid", "ilopassword 非法"))
			}
			if v.Mac1 == "" {
				logger.Warn("CollectDeviceInfoRequest Validate error:mac1 invalid")
				panic(constant.BuildInvalidArgumentWithArgs("mac1 invalid", "mac1 非法"))
			}
			if net.ParseIP(v.PrivateIPv4) == nil {
				logger.Warn("CollectDeviceInfoRequest Validate error:privateipv4 invalid")
				panic(constant.BuildInvalidArgumentWithArgs("privateipv4 invalid", "privateipv4 非法"))
			}

		}
	}

}
