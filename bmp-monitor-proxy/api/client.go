package api

import (
	"crypto/tls"
	"crypto/x509"
	"time"

	"github.com/beego/beego/v2/client/httplib"
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
	// req.DumpBody(true)
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

func NewProxyClient(url, method string, caCrt, cert, key []byte) *ProxyClient {
	pool := x509.NewCertPool()

	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return nil
	}

	tlsConfig := &tls.Config{
		RootCAs:      pool,
		Certificates: []tls.Certificate{cliCrt},
	}

	defaultSetting := httplib.BeegoHTTPSettings{
		UserAgent:        "JDCPS-AGENT",
		ConnectTimeout:   1 * time.Second,
		ReadWriteTimeout: 1 * time.Second,
		Gzip:             true,
		EnableCookie:     false,
		TLSClientConfig:  tlsConfig,
		Retries:          3,
	}

	r := httplib.NewBeegoRequest(url, method)
	if r == nil {
		return nil
	}
	r = r.Setting(defaultSetting)

	return &ProxyClient{BeegoHTTPRequest: r}
}

func ProxyPost(url string, caCrt, cert, key []byte) *ProxyClient {
	return NewProxyClient(url, "POST", caCrt, cert, key)
}

func ProxyGet(url string, caCrt, cert, key []byte) *ProxyClient {
	return NewProxyClient(url, "GET", caCrt, cert, key)
}
