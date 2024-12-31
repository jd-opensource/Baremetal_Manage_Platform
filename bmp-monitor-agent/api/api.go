package api

import (
	"crypto/sha512"
	"fmt"
	"time"

	"coding.jd.com/bmp/cps-agent/api/models"

	"github.com/astaxie/beego/logs"
)

// PutToMonitor post MonitorPut to monitor
func PutToMonitor(url, access, token string, o models.MonitorPut) error {
	if url == "" || access == "" || token == "" {
		return fmt.Errorf("parameter nil")
	}

	req := Post(url)
	if req == nil {
		return fmt.Errorf("Init HttpClient Error")
	}

	if body, err := o.JSON(); err != nil {
		logs.Warn("PutToMonitor JSON Error: %+v", err)
		return err
	} else {
		req.Body(body).Header("Accept", "application/json").Header("Content-Type", "application/json; charset=utf-8")
	}

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	str := access + token + timestamp
	tmp := fmt.Sprintf("Sign %s %x %s", access, sha512.Sum512([]byte(str)), timestamp)
	req.Header("Authorization", tmp)

	req.Debug(true)

	resp := models.MonitorPutResponse{}
	if err := req.ToJSON(&resp); err != nil {
		logs.Warn("PutToMonitor Response Error: %+v", err)
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("PutToMonitor Response Error %+v", resp)
	}

	return nil
}

// Put post ProxyPut object to proxy
func Put(url string, o models.ProxyPut) error {
	if url == "" {
		return fmt.Errorf("url nil")
	}

	req := ProxyPost(url)
	if req == nil {
		return fmt.Errorf("Init HttpClient Error")
	}

	if body, err := o.JSON(); err != nil {
		return err
	} else {
		req.Body(body).Header("Content-Type", "application/json")
	}

	resp := &ProxyResponse{}
	if err := req.ToJSON(resp); err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf("Error: Code=%d, Message=%s", resp.Error.Code, resp.Error.Message)
	}

	return nil
}

// Heartbeat post AgentStatus object to proxy
func Heartbeat(url string, o models.AgentStatus) error {
	if url == "" {
		return fmt.Errorf("url nil")
	}

	req := ProxyPost(url)
	if req == nil {
		return fmt.Errorf("Init HttpClient Error")
	}

	if body, err := o.JSON(); err != nil {
		return err
	} else {
		req.Body(body).Header("Content-Type", "application/json")
	}

	resp := &ProxyResponse{}
	if err := req.ToJSON(resp); err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf("Error: Code=%d, Message=%s", resp.Error.Code, resp.Error.Message)
	}

	return nil
}

// PutToProxy post MonitorPut object to proxy, used to send to monitor through proxy
func PutToProxy(url string, o models.MonitorPut, ca, cert, key []byte) error {
	if url == "" {
		return fmt.Errorf("url nil")
	}

	req := ProxyPost(url)
	if req == nil {
		return fmt.Errorf("Init HttpClient Error")
	}

	if body, err := o.JSON(); err != nil {
		return err
	} else {
		req.Body(body).Header("Content-Type", "application/json")
	}

	resp := &ProxyResponse{}
	if err := req.ToJSON(resp); err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf("Error: Code=%d, Message=%s", resp.Error.Code, resp.Error.Message)
	}

	return nil
}

// GetInstanceID get InstanceID from cps-api
func GetInstanceID(url, sn string) (*models.Instance, error) {
	if sn == "" || url == "" {
		return nil, fmt.Errorf("url nil")
	}

	req := Get(url)
	if req == nil {
		return nil, fmt.Errorf("Init HttpClient Error")
	}

	instance := &models.Instance{}
	if err := req.Param("sn", sn).ToJSON(instance); err != nil {
		return nil, err
	}

	return instance, nil
}
