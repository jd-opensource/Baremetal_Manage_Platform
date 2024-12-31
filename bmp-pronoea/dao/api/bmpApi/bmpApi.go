package bmpApi

import (
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/dao/api"
	"github.com/astaxie/beego"
	"os"
)

type BmpApi struct {
	api.BaseApi
}

// 方法名
const (
	REPORT_ALERT       = "reportAlert"     // 上报报警
)


var apis = map[string]*api.UrlPath{
	REPORT_ALERT: {
		Method: "POST",
		Path:   "/v1/monitorAlert/addAlert",
	},
}

func (a *BmpApi) ServiceConfig(configHeader string) error {
	networkConfig, err := a.ParseConfigByHeader(configHeader)
	if err != nil {
		return err
	}

	a.Host = networkConfig.Host
	a.Domain = networkConfig.Domain
	a.TimeOut = networkConfig.Timeout
	a.ConnectTimeOut = networkConfig.ConnectTimeout
	a.MaxIdleConnsPerHost = networkConfig.MaxIdleConnsPerHost

	authorization := beego.AppConfig.String(fmt.Sprintf("%v.authorization", configHeader))
	if util.IsEmpty(authorization) {
		return fmt.Errorf("%v.authorization config empty ", configHeader)
	}
	a.Config["Authorization"] = authorization

	return nil
}

func (a *BmpApi) GetBaseApi() *api.BaseApi {
	return &a.BaseApi
}

func (a *BmpApi) GetRequestHeader(timeStamp int64) map[string]string {
	headers := make(map[string]string, 0)
	headers["Authorization"] = a.Config["Authorization"]
	headers["Timestamp"] = fmt.Sprintf("%v", timeStamp)

	return headers
}

func NewBmpApi() (*BmpApi, error) {
	a := &BmpApi{
		api.BaseApi{
			Apis:   apis,
			Config: make(map[string]string),
		},
	}
	if err := a.ServiceConfig("bmpopenapiconsole"); err != nil {
		return nil, err
	}
	return a, nil
}

func init() {
	a, err := NewBmpApi()
	if err != nil {
		fmt.Printf("init bmp api error:%s", err.Error())
		os.Exit(1)
	}
	api.ApiMap[api.BMP_API] = a
}
