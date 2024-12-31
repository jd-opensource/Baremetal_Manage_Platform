package config

import (
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"github.com/astaxie/beego"
)

// HttpServerConfigByHeader 加载http第三方依赖的配置
func HttpServerConfigByHeader(header string) (*HttpConfig, error) {
	port := beego.AppConfig.String(fmt.Sprintf("%v.port", header))
	if util.IsEmpty(port) {
		return nil, fmt.Errorf("%v.port config empty ", header)
	}

	host := beego.AppConfig.String(fmt.Sprintf("%v.host", header))
	if util.IsEmpty(host) {
		return nil, fmt.Errorf("%v.host config empty ", header)
	}

	connecttimeout, err := beego.AppConfig.Int(fmt.Sprintf("%v.connecttimeout", header))
	if err != nil{
		return nil, fmt.Errorf("%v.connecttimeout config error %v ", header, err)
	}
	if connecttimeout <= 0 {
		connecttimeout = 15
	}

	timeout, err := beego.AppConfig.Int(fmt.Sprintf("%v.timeout", header))
	if err != nil{
		return nil, fmt.Errorf("%v.timeout config error %v ", header, err)
	}
	if timeout <= 0 {
		timeout = 20
	}

	maxidleconnsperhost, err := beego.AppConfig.Int(fmt.Sprintf("%v.maxidleconnsperhost", header))
	if err != nil{
		return nil, fmt.Errorf("%v.maxidleconnsperhost config error %v ", header, err)
	}
	if maxidleconnsperhost <= 0 {
		maxidleconnsperhost = 1
	}

	configRes := &HttpConfig{
		Domain: fmt.Sprintf("http://%v:%v", host, port),
		Host:	host,
		Port:   port,
		ConnectTimeout: connecttimeout,
		Timeout: timeout,
		MaxIdleConnsPerHost: maxidleconnsperhost,
	}

	return configRes, nil
}

