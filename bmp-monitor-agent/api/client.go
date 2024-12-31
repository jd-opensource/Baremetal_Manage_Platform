package api

import (
	"time"

	"github.com/astaxie/beego/httplib"
)

type HttpClient struct {
	*httplib.BeegoHTTPRequest
}

func NewHttpClient(url, method string) *HttpClient {
	req := httplib.NewBeegoRequest(url, method)
	if req == nil {
		return nil
	}

	req.SetTimeout(1*time.Second, 1*time.Second)
	req.SetUserAgent("JDCPS-AGENT")
	req.DumpBody(true)
	req.SetEnableCookie(false)
	req.Retries(3)

	return &HttpClient{BeegoHTTPRequest: req}
}

func Post(url string) *HttpClient {
	return NewHttpClient(url, "POST")
}

func Get(url string) *HttpClient {
	return NewHttpClient(url, "GET")
}

type ProxyClient struct {
	*httplib.BeegoHTTPRequest
}

func NewProxyClient(url, method string) *ProxyClient {

	defaultSetting := httplib.BeegoHTTPSettings{
		UserAgent:        "JDCPS-AGENT",
		ConnectTimeout:   1 * time.Second,
		ReadWriteTimeout: 1 * time.Second,
		Gzip:             true,
		DumpBody:         true,
		EnableCookie:     false,
		Retries:          3,
	}

	r := httplib.NewBeegoRequest(url, method)
	if r == nil {
		return nil
	}
	r = r.Setting(defaultSetting)

	return &ProxyClient{BeegoHTTPRequest: r}
}

func ProxyPost(url string) *ProxyClient {
	return NewProxyClient(url, "POST")
}

func ProxyGet(url string) *ProxyClient {
	return NewProxyClient(url, "GET")
}
