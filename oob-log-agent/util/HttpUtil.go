package util

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func doPost(urlstr, param string, header map[string]string) (resp string, err error) {
	u, err := url.Parse(urlstr)
	if err != nil {
		return "", err
	}
	u.RawQuery = u.Query().Encode()

	//跳过证书验证,禁止连接保持
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
		},
	}

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(param))
	if err != nil {
		return "", err
	}

	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//增加header选项
	for k, v := range header {
		req.Header.Add(k, v)
	}

	respBody, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer respBody.Body.Close()

	if respBody.StatusCode != 200 {
		return "", fmt.Errorf("Http Post %s return %d", urlstr, respBody.StatusCode)
	}

	body, err := ioutil.ReadAll(respBody.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// HTTPPost http post
func HTTPPost(urlstr, param string) (resp string, err error) {
	return doPost(urlstr, param, nil)
}

func HTTPPostWithHeader(urlstr, param string, header map[string]string) (resp string, err error) {
	return doPost(urlstr, param, header)
}

func doGet(urlstr string, user string, pass string, header map[string]string) (resp string, err error) {
	u, err := url.Parse(urlstr)
	if err != nil {
		return "", err
	}
	if user != "" {
		u.User = url.UserPassword(user, pass)
	}
	u.RawQuery = u.Query().Encode()

	//跳过证书验证,禁止连接保持
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
		},
	}

	//提交请求
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}

	//增加header选项
	for k, v := range header {
		req.Header.Add(k, v)
	}

	respBody, errGet := c.Do(req)
	if errGet != nil {
		return "", errGet
	}
	defer respBody.Body.Close()

	if respBody.StatusCode != 200 {
		return "", fmt.Errorf("Http Get %s return %d", urlstr, respBody.StatusCode)
	}
	body, err := ioutil.ReadAll(respBody.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
func HTTPGetWithHeader(urlstr string, header map[string]string) (resp string, err error) {
	return doGet(urlstr, "", "", header)
}

func HTTPGetWithBasicAuth(urlstr string, user string, pass string) (resp string, err error) {
	return doGet(urlstr, user, pass, nil)
}

// HTTPGet http Get
func HTTPGet(urlstr string) (resp string, err error) {
	return doGet(urlstr, "", "", nil)
}
