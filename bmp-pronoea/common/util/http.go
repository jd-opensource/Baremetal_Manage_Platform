package util

import (
	"compress/gzip"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func DoHttpRequest(requestId, reqUrl, host, method, bodyStr string, params, headers map[string]string, connTimeout, timeout, maxIdleConnsPerHost int) ([]byte, error) {

	if connTimeout <= 0 {
		connTimeout = 50000
	}
	if timeout <= 0 {
		timeout = 50000
	}
	if maxIdleConnsPerHost <= 0 {
		maxIdleConnsPerHost = 1
	}

	tr := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, time.Duration(connTimeout)*time.Second)
		},
		MaxIdleConnsPerHost:   maxIdleConnsPerHost,
		ResponseHeaderTimeout: time.Duration(connTimeout) * time.Second,
		DisableKeepAlives:     true,
	}

	client := &http.Client{Transport: tr, Timeout: time.Duration(timeout) * time.Second}

	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}

	var body io.Reader
	body = nil
	if method == "GET" {
		u, err := url.ParseRequestURI(reqUrl)
		if err != nil {
			logs.Warn("%v:parse request uri for %v error:%v", requestId, reqUrl, err)
			return nil, err
		}
		u.RawQuery = values.Encode()
		reqUrl = fmt.Sprintf("%v", u)
	} else {
		if bodyStr == "" {
			body = strings.NewReader(values.Encode())
		} else {
			body = strings.NewReader(bodyStr)
		}
	}

	req, err := http.NewRequest(method, reqUrl, body)
	if err != nil {
		logs.Warn("%v:new request for url:%v error:%v", requestId, reqUrl, err)
		return nil, err
	}
	logs.Trace("%v:begin send request to url:%v, body:%v, param:%v", requestId,reqUrl, body, params)
	if host != "" {
		req.Host = host
		if strings.Contains(reqUrl, "https://") {
			tls := &tls.Config{
				ServerName: host,
			}
			tr.TLSClientConfig = tls
		}
	}

	//req.Header.Set("Content-Type", contentType) "application/json"
	if headers != nil && len(headers) > 0 {
		for headerKey, headerValue := range headers {
			req.Header.Set(headerKey, headerValue)
		}
	}

	// default, accept gzip
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	if err != nil {
		logs.Warn("%v:do request for url:%v error:%v", requestId, reqUrl, err)
		return nil, err
	}
	defer resp.Body.Close()

	// check content type
	rbody := resp.Body
	resContentType := resp.Header.Get("Content-Encoding")
	if strings.Contains(resContentType, "gzip") {
		// decode gzip
		gzipBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			logs.Warn("%v:gzip reader error:%v", requestId, err)
			return nil, err
		}
		rbody = gzipBody
	}

	result, err := ioutil.ReadAll(rbody)
	if err != nil {
		logs.Warn("%v: read response from url:%v error:%v", requestId, reqUrl, err)
		return nil, err
	}

	if resp.StatusCode != 200 &&  resp.StatusCode != 202 {
		logs.Warn("%v: http status code:%v, url:%v, result:%v", requestId, resp.StatusCode, reqUrl, string(result))
		return nil, errors.New(fmt.Sprintf("http response status code error, code:%v", resp.StatusCode))
	}

	logs.Trace("%v:success get result from url:%v, result:%v", requestId, reqUrl, string(result))

	return result, nil
}