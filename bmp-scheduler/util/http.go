package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

func DoHttpPost(url string, header map[string]string, data interface{}) ([]byte, error) {
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
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()
	//务必带 X-Jdcloud-Requestid
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
	return respBytes, nil
}
