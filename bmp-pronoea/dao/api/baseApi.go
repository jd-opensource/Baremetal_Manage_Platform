package api

import (
	"encoding/json"
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/config"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

const (
	//服务
	PROMETHEUS_API     = "prometheus-api"
	PUSHGATEWAY_API    = "pushgateway-api"
	BMP_API    = "bmp-api"

	//参数类型
	PARAMS_TYPE_QUERY      = "query"
	PARAMS_TYPE_JSON       = "json"
	PARAMS_TYPE_FORM       = "form"
	PARAMS_TYPE_URLENCODED = "urlencoded"
	PARAMS_TYPE_ARRAY_JSON       = "array_json"

)

var (
	ApiMap = make(map[string]IBaseApi)
	Locker sync.Mutex
)

type IBaseApi interface {
	ServiceConfig(string) error
	GetBaseApi() *BaseApi
	GetRequestHeader(int64) map[string]string
}

type BaseApi struct {
	Apis           map[string]*UrlPath
	Host           string
	Domain         string
	Config         map[string]string
	ConnectTimeOut int
	TimeOut        int
	MaxIdleConnsPerHost int
	Options        *Options
}

type UrlPath struct {
	Method string
	Path   string
}

type Options struct {
	Attributes map[string]interface{}
	Data       map[string]IBaseApiRequest
	Headers    map[string]string
}

func (a *BaseApi) getURL(apiName string) (*UrlPath, error) {
	if urlPath, ok := a.Apis[apiName]; ok {
		return urlPath, nil
	}
	return nil, fmt.Errorf("%s apiname not found", apiName)
}

func (a *BaseApi) ParseConfig(configPath string) (interface{}, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		logs.Error("read config failed:" + err.Error())
		return nil, err
	}

	var config interface{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		logs.Error("json parse config failed:" + err.Error())
		return nil, err
	}

	return config, nil
}

func (a *BaseApi) ParseConfigByHeader(header string) (*config.HttpConfig, error) {
	return config.HttpServerConfigByHeader(header)
}


func GetApi(name string) (IBaseApi, error) {
	api, ok := ApiMap[name]
	if !ok {
		return nil, fmt.Errorf("unknown api %s, %v", name, ApiMap)
	}
	return api, nil
}

func Send(requestId string, a IBaseApi, apiName string) ([]byte, error) {
	base := a.GetBaseApi()
	urlPath, err := base.getURL(apiName)
	if err != nil {
		return nil, err
	}
	//url中的替换参数
	u := urlPath.Path
	if len(base.Options.Attributes) > 0 {
		for k, v := range base.Options.Attributes {
			switch v.(type) {
			case string:
				u = strings.Replace(u, "{"+k+"}", v.(string), 1)
			case int:
				u = strings.Replace(u, "{"+k+"}", strconv.Itoa(v.(int)), 1)
			}
		}
	}

	params := make(map[string]string)
	body := ""

	if base.Options.Data != nil {
		//处理query参数
		if data, ok := base.Options.Data[PARAMS_TYPE_QUERY]; ok {
			query, err := data.Decode()
			if err != nil {
				return nil, err
			}
			//利用反射处理不同的结构体变量类型
			typeData := reflect.TypeOf(data).Elem()
			for k, v := range query {
				switch v.(type) {
				case json.Number:
					for i := 0; i < typeData.NumField(); i++ {
						tag := typeData.Field(i).Tag.Get("json")
						ret := strings.Split(tag, util.SEPARATOR_COMMA)
						if ret[0] == k {
							dataType := typeData.Field(i).Type.Name()
							if dataType == "int64" || dataType == "int32" || dataType == "int" {
								res, _ := v.(json.Number).Int64()
								params[k] = strconv.Itoa(int(res))
							} else if dataType == "float64" || dataType == "float32" {
								res, _ := v.(json.Number).Float64()
								params[k] = fmt.Sprintf("%.2f", res)
							} else if dataType == "string" {
								params[k] = v.(json.Number).String()
							}
							break
						}
					}
				case string:
					params[k] = v.(string)
				}
			}
		}
		//处理application/json参数
		if data, ok := base.Options.Data[PARAMS_TYPE_JSON]; ok {
			bodyData, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			body = string(bodyData)
			base.Options.Headers["Content-Type"] = "application/json"
		} else if data, ok = base.Options.Data[PARAMS_TYPE_ARRAY_JSON]; ok {
			bodyData, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			body = fmt.Sprintf("[%s]", string(bodyData))
			base.Options.Headers["Content-Type"] = "application/json"
		} else if data, ok = base.Options.Data[PARAMS_TYPE_URLENCODED]; ok { //处理application/x-www-form-urlencoded参数
			urlencoded, err := data.Decode()
			if err != nil {
				return nil, err
			}
			for k, v := range urlencoded {
				body += k + "=" + fmt.Sprintf("%s", v) + "&"
			}
			body = strings.Trim(body, "&")
			base.Options.Headers["Content-Type"] = "application/x-www-form-urlencoded"
		} else if data, ok = base.Options.Data[PARAMS_TYPE_FORM]; ok { //处理multipart/form-data参数
			form, err := data.Decode()
			if err != nil {
				return nil, err
			}
			for k, v := range form {
				body += k + "=" + fmt.Sprintf("%s", v) + ";"
			}
			body = strings.Trim(body, ";")
			base.Options.Headers["Content-Type"] = "multipart/form-data"
		}
	}

	reqURL := base.Domain + u
	base.Options.Headers["X-Request-Id"] = requestId
	//Locker.Unlock()
	response, err := util.DoHttpRequest(
		requestId,
		reqURL,
		base.Host,
		urlPath.Method,
		body,
		params,
		base.Options.Headers,
		base.ConnectTimeOut,
		base.TimeOut,
		base.MaxIdleConnsPerHost,
	)
	if err != nil {
		logs.Info(requestId, fmt.Sprintf("do http request error:%v, params:%v", err, params))
		return nil, err
	}

	logs.Info(requestId, fmt.Sprintf("do http response:%s", string(response)))
	return response, nil
}
