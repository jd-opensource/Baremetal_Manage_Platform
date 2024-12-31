package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"coding.jd.com/bmp/agent-proxy-jdstack/api/models"
	"github.com/beego/beego/v2/server/web"

	log "git.jd.com/cps-golang/log"
)

// Put post ProxyPut object to proxy
func Put(url string, o models.ProxyPut, ca, cert, key []byte) error {
	if url == "" {
		return fmt.Errorf("url nil")
	}

	req := ProxyPost(url, ca, cert, key)
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

	if resp.Code != 0 {
		return fmt.Errorf("Error: Code=%d, Message=%s", resp.Code, resp.Message)
	}

	return nil
}

// Heartbeat post AgentStatus object to proxy
func Heartbeat(url string, o models.AgentStatus, ca, cert, key []byte) error {
	if url == "" {
		return fmt.Errorf("url nil")
	}

	req := ProxyPost(url, ca, cert, key)
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

	if resp.Code != 0 {
		return fmt.Errorf("Error: Code=%d, Message=%s", resp.Code, resp.Message)
	}

	return nil
}

// PutToProxy post MonitorPut object to proxy, used to send to monitor through proxy
func PutToProxy(url string, o models.MonitorPut, ca, cert, key []byte) error {
	if url == "" {
		return fmt.Errorf("url nil")
	}

	req := ProxyPost(url, ca, cert, key)
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

	if resp.Code != 0 {
		return fmt.Errorf("Error: Code=%d, Message=%s", resp.Code, resp.Message)
	}

	return nil
}

// GetInstanceID get InstanceID from cps-api
func GetInstanceID(logger *log.Logger, sn string) (*models.Instance, error) {

	if sn == "minping-mock-sn" { //10.226.192.114测试专用
		return &models.Instance{
			Result: models.InstanceResult{
				Instance: models.InstanceDetail{
					InstanceID:   "minping-mock-instanceId",
					InstanceName: "minping-mock",
					Sn:           "minping-mock-sn",
				},
			},
		}, nil
	}

	openapi, _ := web.AppConfig.String("openapi_host")
	openapiPort, _ := web.AppConfig.String("openapi_port")
	cpsApi, _ := web.AppConfig.String("api.cps")
	url := fmt.Sprintf("http://%s:%s%s", openapi, openapiPort, cpsApi)

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

	v, _ := json.Marshal(instance)
	logger.Infof("GetInstanceID.instance success, sn:%s, instance:%s", sn, string(v))

	return instance, nil
}

// GetInstanceID get InstanceID from cps-api
func GetPinByTenant(url, tenant string) (*models.PinResult, error) {
	if tenant == "" || url == "" {
		return nil, fmt.Errorf("url nil")
	}

	req := Get(url)
	if req == nil {
		return nil, fmt.Errorf("Init HttpClient Error")
	}

	pinRes := &models.PinResult{}
	if err := req.Param("tenant", tenant).ToJSON(pinRes); err != nil {
		return nil, err
	}
	if pinRes.Result.Pin == "" {
		return nil, errors.New("empty pin")
	}

	return pinRes, nil
}
