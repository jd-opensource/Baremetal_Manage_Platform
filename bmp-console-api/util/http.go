package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"strings"
	"time"

	log "coding.jd.com/aidc-bmp/bmp_log"
)

/*
func DoHttpGet(logger *log.Logger, url string, header map[string]string, data interface{}) ([]byte, error) {

	rpclogger := log.New("./logs/ironic-console-rpc.log")
	defer rpclogger.Flush()
	defer rpclogger.TimeEnd("cost")
	rpclogger.SetStableFields([]string{"url", "method", "cost", "header", "request", "status", "statusCode", "response"})

	rpclogger.Point("logid", logger.GetPoint("logid"))
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
	fmt.Println(data, param)
	for key, value := range param {
		q.Add(key, value)
	}
	request.URL.RawQuery = q.Encode()
	fmt.Println("get请求的最终url:", request.URL.String())
	defer request.Body.Close()
	//务必带 X-Jdcloud-Requestid
	for k, v := range header {
		//fmt.Println(k,v)
		request.Header.Set(k, v)
	}
	//if header["host"] != "" {
	request.Host = header["Host"]
	//}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	rpclogger.Point("status", resp.Status)
	rpclogger.Point("statusCode", resp.StatusCode)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("状态码：", resp.Status, resp.StatusCode)
	rpclogger.Point("response", string(respBytes))
	return respBytes, nil
}

func DoHttpPost(logger *log.Logger, url string, header map[string]string, data interface{}) ([]byte, error) {

	rpclogger := log.New("./logs/ironic-console-rpc.log")
	defer rpclogger.Flush()
	defer rpclogger.TimeEnd("cost")
	rpclogger.SetStableFields([]string{"url", "method", "cost", "header", "request", "status", "statusCode", "response"})

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
	fmt.Println("状态码：", resp.Status, resp.StatusCode)
	rpclogger.Point("status", resp.Status)
	rpclogger.Point("statusCode", resp.StatusCode)
	//if resp.StatusCode != 200 {
	//	fmt.Println("ironic-api接口异常")
	//	return nil,err
	//}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rpclogger.Point("response", string(respBytes))
	return respBytes, nil
}
*/
func DoHttp(logger *log.Logger, url string, header map[string]string, data interface{}, method string) ([]byte, error) {

	rpclogger := log.New("./log/ironic-console-rpc.log")
	defer rpclogger.Flush()
	defer rpclogger.TimeEnd("cost")
	rpclogger.SetStableFields([]string{"url", "method", "cost", "header", "request", "status", "statusCode", "response"})

	rpclogger.Point("logid", logger.GetPoint("logid"))
	rpclogger.Point("url", url)
	rpclogger.Point("method", method)
	rpclogger.Point("header", header)
	rpclogger.TimeStart("cost")

	var bytesData []byte
	if data != nil {
		fmt.Println("当前参数类型", reflect.TypeOf(data).Kind())
		if reflect.TypeOf(data).Kind() == reflect.String {
			bytesData = []byte(data.(string))
		} else {
			var err error
			bytesData, err = json.Marshal(data)
			if err != nil {
				return nil, err
			}
		}
	}

	// byteHeader, _ := json.Marshal(header)
	// fmt.Println("请求header", string(byteHeader))
	fmt.Println("ironicapi请求url", url)
	fmt.Println("ironicapi请求method", method)
	fmt.Println("ironicapi请求body", string(bytesData))
	rpclogger.Point("request", string(bytesData))
	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	if method == "GET" && data != nil {
		param := InterfaceToMap(data)
		q := request.URL.Query()
		for key, value := range param {
			q.Add(key, value.(string))
		}
		request.URL.RawQuery = q.Encode()
	}
	fmt.Println("ironicapi请求url:", request.URL.String())
	defer request.Body.Close()
	//务必带 X-Jdcloud-Request-Id
	for k, v := range header {
		//fmt.Println(k,v)
		request.Header.Set(k, v)
	}

	fmt.Println("dohttp request header:", request.Header)

	// request.Header.Set("x-jdcloud-erp", "ZXJw")
	if header["Host"] != "" {
		request.Host = header["Host"]
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	rpclogger.Point("status", resp.Status)
	rpclogger.Point("statusCode", resp.StatusCode)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println("状态码：", resp.Status, resp.StatusCode)
	fmt.Println("ironicapi响应:", string(respBytes))
	rpclogger.Point("response", string(respBytes))
	return respBytes, nil
}

func CheckIpPort(ipPort string) error {
	ipPort = strings.TrimLeft(ipPort, "http://")
	ipPort = strings.TrimLeft(ipPort, "https://")
	conn, err := net.DialTimeout("tcp", ipPort, 3*time.Second)
	if err != nil {
		return err
	} else {
		if conn != nil {
			conn.Close()
			return nil
		} else {
			return fmt.Errorf("%s dial failed", ipPort)
		}
	}
}
