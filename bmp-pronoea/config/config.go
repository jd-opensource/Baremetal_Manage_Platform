package config

import (
	"encoding/json"
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type HttpConfig struct {
	Domain        string  `json:"Domain"`
	Host          string  `json:"Host"`
	Port          string  `json:"Port"`
	ConnectTimeout      int
	Timeout             int
	MaxIdleConnsPerHost int
}

type BmpOpenApiConfig struct {
	BaseConfig *HttpConfig
	Authorization string
}

type PrometheusConfig struct {
	BaseConfig *HttpConfig
	PrometheusRulePath string
}

type PushgatewayConfig struct {
	BaseConfig *HttpConfig
}

type SystemConfig struct {
	Crontabs      map[string]string
	PrometheusRulePath string
}


var SysConfig *SystemConfig

// ParseConfig 加载本地文件的方式
func ParseConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	SysConfig = &SystemConfig{}
	err = json.Unmarshal(data, SysConfig)
	if err != nil {
		return err
	}

	return nil
}


func ParseConfigAppConf() error {
	SysConfig = &SystemConfig{}

	// 初始化 crontab 配置
	SysConfig.Crontabs = initCrontabConfig()

	rulePath := beego.AppConfig.String("prometheus.rulepath")
	if util.IsEmpty(rulePath) {
		return fmt.Errorf("prometheus rule path empty ")
	}
	SysConfig.PrometheusRulePath = rulePath
	return nil
}


