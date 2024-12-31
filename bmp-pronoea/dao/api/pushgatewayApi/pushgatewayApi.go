package pushgatewayApi

import (
	"fmt"
	"git.jd.com/bmp-pronoea/dao/api"
	"os"
)

type PushgatewayApi struct {
	api.BaseApi
}

// 方法名
const (
	WIPE_ALL_DATA = "wipeAllData"     // 删除全部数据
)

var apis = map[string]*api.UrlPath{
	WIPE_ALL_DATA: {
		Method: "PUT",
		Path:   "/api/v1/admin/wipe",
	},
}

func (a *PushgatewayApi) ServiceConfig(configHeader string) error {
	networkConfig, err := a.ParseConfigByHeader(configHeader)
	if err != nil {
		return err
	}

	a.Host = networkConfig.Host
	a.Domain = networkConfig.Domain
	a.TimeOut = networkConfig.Timeout
	a.ConnectTimeOut = networkConfig.ConnectTimeout
	a.MaxIdleConnsPerHost = networkConfig.MaxIdleConnsPerHost

	return nil
}

func (a *PushgatewayApi) GetBaseApi() *api.BaseApi {
	return &a.BaseApi
}

func (a *PushgatewayApi) GetRequestHeader(timeStamp int64) map[string]string {
	headers := make(map[string]string, 0)
	headers["Timestamp"] = fmt.Sprintf("%v", timeStamp)
	/*
	headers["Authentication"] = a.Config["auth"]
	md5Base := fmt.Sprintf("%v%v%v", a.Config["ak"], a.Config["sk"], timeStamp)
	sign := util.GetMd5(md5Base)
	headers["Signature"] = sign
	 */

	return headers
}

func NewPushgatewayApi() (*PushgatewayApi, error) {
	a := &PushgatewayApi{
		api.BaseApi{
			Apis:   apis,
			Config: make(map[string]string),
		},
	}
	if err := a.ServiceConfig("pushgateway"); err != nil {
		return nil, err
	}
	return a, nil
}

func init() {
	a, err := NewPushgatewayApi()
	if err != nil {
		fmt.Printf("init pushgateway api error:%s", err.Error())
		os.Exit(1)
	}
	api.ApiMap[api.PUSHGATEWAY_API] = a
}
