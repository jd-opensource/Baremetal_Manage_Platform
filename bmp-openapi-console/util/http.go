package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	log "coding.jd.com/aidc-bmp/bmp_log"
	"github.com/beego/beego/v2/server/web"
)

func DoHttpPost(logger *log.Logger, url string, header map[string]string, data interface{}) ([]byte, error) {

	logPath, _ := web.AppConfig.String("log.path")
	rpclogger := log.New(logPath + "/openapi-rpc.log")
	defer rpclogger.Flush()
	defer rpclogger.TimeEnd("cost")
	rpclogger.SetStableFields([]string{"url", "method", "cost", "header", "request", "response"})

	rpclogger.Point("logid", logger.GetPoint("logid").(string))
	rpclogger.Point("url", url)
	rpclogger.Point("method", "POST")
	rpclogger.Point("header", header)
	rpclogger.TimeStart("cost")

	var bytesData []byte
	if reflect.TypeOf(data).Kind() == reflect.String {
		bytesData = []byte(data.(string))
	} else {
		var err error
		bytesData, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}
	rpclogger.Point("request", string(bytesData))
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()
	//务必带 X-Jdcloud-Request-Id
	for k, v := range header {
		request.Header.Set(k, v)
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	rpclogger.Point("response", string(respBytes))
	return respBytes, nil
}

func DoHttpGet(logger *log.Logger, url string, header map[string]string, data interface{}) ([]byte, error) {

	logPath, _ := web.AppConfig.String("log.path")
	rpclogger := log.New(logPath + "/openapi-rpc.log")
	defer rpclogger.Flush()
	defer rpclogger.TimeEnd("cost")
	rpclogger.SetStableFields([]string{"url", "method", "cost", "header", "request", "response"})

	rpclogger.Point("logid", logger.GetPoint("logid").(string))
	rpclogger.Point("url", url)
	rpclogger.Point("method", "GET")
	rpclogger.Point("header", header)
	rpclogger.TimeStart("cost")

	var bytesData []byte
	if reflect.TypeOf(data).Kind() == reflect.String {
		bytesData = []byte(data.(string))
	} else {
		var err error
		bytesData, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}
	rpclogger.Point("request", string(bytesData))
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("GET", url, reader)
	if err != nil {
		return nil, err
	}
	q := request.URL.Query()
	param := InterfaceToMap(data)
	for key, value := range param {
		q.Add(key, value.(string))
	}
	request.URL.RawQuery = q.Encode()

	defer request.Body.Close()
	//务必带 X-Jdcloud-Request-Id
	for k, v := range header {
		request.Header.Set(k, v)
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	rpclogger.Point("response", string(respBytes))
	return respBytes, nil
}
