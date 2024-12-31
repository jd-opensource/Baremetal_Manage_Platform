package controllers

import (
	"encoding/json"
	"time"

	"coding.jd.com/bmp/agent-proxy-jdstack/api"
	"coding.jd.com/bmp/agent-proxy-jdstack/api/models"
	"coding.jd.com/bmp/agent-proxy-jdstack/logic"

	"coding.jd.com/bmp/agent-proxy-jdstack/global"
	responseTypes "coding.jd.com/bmp/agent-proxy-jdstack/types/response"

	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

type ProxyController struct {
	BaseController
}

// Put 上传metric，将SN转换为实例ID，调用Monitor的Put接口
func (c *ProxyController) Put() {
	c.logPoints.Infof("put api start, time:%d", time.Now().Unix())
	req := c.Ctx.Input.RequestBody

	reqObj := models.ProxyPut{}
	var instance *models.Instance

	if err := json.Unmarshal(req, &reqObj); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	reqObj = modifyPutDataPoint(reqObj)
	c.logPoints.Infof("put api start11, time:%d", time.Now().Unix())
	if v, found := global.Get(reqObj.SN); !found {
		if ins, err := api.GetInstanceID(c.logPoints, reqObj.SN); err != nil || ins == nil {
			c.logPoints.Warnf("%s GetInstanceID Error: %+v", reqObj.SN, err)
			c.SetErrorResponse(httpStatus.BAD_REQUEST, "sn invalid", errorCode.INVALID_ARGUMENT)
			return
		} else {
			// logs.Debug("%s GetInstanceID Response: %+v", reqObj.SN, ins)
			if ins.Error.Code > 0 {
				c.SetErrorResponse(httpStatus.BAD_REQUEST, ins.Error.Message, errorCode.INVALID_ARGUMENT)
				return
			}
			instance = ins
			global.Set(reqObj.SN, instance)
		}
	} else {
		instance = v.(*models.Instance)
	}
	c.logPoints.Infof("put api start22, time:%d", time.Now().Unix())

	if instance == nil || instance.Result.Instance.InstanceID == "" {
		c.logPoints.Warnf("%s GetInstanceID empty", reqObj.SN)
		c.SetErrorResponse(httpStatus.BAD_REQUEST, "sn invalid", errorCode.INVALID_ARGUMENT)
		return
	}

	models.PutDeviceTags2Redis(c.logPoints, instance.Result.Instance.InstanceID, reqObj)
	c.logPoints.Infof("put api start33, time:%d", time.Now().Unix())

	//上报云监控
	err := logic.PushGateway(c.logPoints, instance, reqObj)
	if err != nil {
		c.logPoints.Warnf("ProxyController.Put.pushCloudMonitor error:%s", err.Error())
	}
	c.logPoints.Infof("put api start44, time:%d", time.Now().Unix())

	c.Res.Result = responseTypes.CommonResponse{
		Success: err == nil,
	}
}

func (c *ProxyController) Heartbeat() {

	req := c.Ctx.Input.RequestBody
	agentStatus := &models.AgentStatus{}
	var instance *models.Instance

	if err := json.Unmarshal(req, &agentStatus); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, "body invalid", errorCode.INVALID_ARGUMENT)

		return
	}

	if v, found := global.Get(agentStatus.SN); !found {
		if ins, err := api.GetInstanceID(c.logPoints, agentStatus.SN); err != nil || ins == nil {
			c.logPoints.Warnf("%s GetInstanceID Error: %+v", agentStatus.SN, err)
			c.SetErrorResponse(httpStatus.BAD_REQUEST, "sn invalid", errorCode.INVALID_ARGUMENT)
			return
		} else {
			c.logPoints.Infof("%s GetInstanceID Response: %+v", agentStatus.SN, ins)
			if ins.Error.Code > 0 {
				c.SetErrorResponse(httpStatus.BAD_REQUEST, "sn invalid", errorCode.INVALID_ARGUMENT)
				return
			}
			instance = ins
			global.Set(agentStatus.SN, instance)
		}
	} else {
		instance = v.(*models.Instance)
	}

	global.Set(instance.Result.Instance.InstanceID, agentStatus)

	c.Res.Result = responseTypes.CommonResponse{
		Success: true,
	}
	return

}

// 客户端升级不便，为了应对业务快速迭代，在procy层对于上报过来的数据点做修正，目前的修正包括如下：
// 1,网卡部分，如果tag中有bond0，则只上报bond0，否则全部上报
func modifyPutDataPoint(reqObj models.ProxyPut) models.ProxyPut {

	var exist func(string, []string) bool = func(item string, items []string) bool {
		for _, i := range items {
			if i == item {
				return true
			}
		}
		return false
	}

	newReqObj := models.ProxyPut{
		SN:         reqObj.SN,
		DataPoints: []models.DataPointA{},
	}

	netMetrics := []string{"bmp.network.bytes.ingress", "bmp.network.bytes.egress", "bmp.network.packets.ingress", "bmp.network.packets.egress"}

	nicMetrics := []models.DataPointA{}
	for _, point := range reqObj.DataPoints {
		if exist(point.Metric, netMetrics) {
			nicMetrics = append(nicMetrics, point)
		} else {
			newReqObj.DataPoints = append(newReqObj.DataPoints, point)
		}
	}

	hasBond0 := false
	bond0Metrics := []models.DataPointA{}

	for _, point := range nicMetrics {
		if point.Tags != nil {
			if point.Tags["device"] == "bond0" {
				hasBond0 = true
				bond0Metrics = append(bond0Metrics, point)
			}

		}
	}

	if hasBond0 {
		newReqObj.DataPoints = append(newReqObj.DataPoints, bond0Metrics...)
	} else {
		newReqObj.DataPoints = append(newReqObj.DataPoints, nicMetrics...)
	}

	return newReqObj
}

// func pushCloudMonitor(logger *log.Logger, instance *models.Instance, reqObj models.ProxyPut) (*monitorSdkApis.PutResponse, error) {
// 	accessKey := beego.AppConfig.DefaultString("openapi.ak", "")
// 	secretKey := beego.AppConfig.DefaultString("openapi.sk", "")
// 	credentials := core.NewCredentials(accessKey, secretKey)

// 	config := core.NewConfig()
// 	config.SetScheme(core.SchemeHttp)

// 	gateWay := beego.AppConfig.DefaultString("openapi.host", "")
// 	config.SetEndpoint(gateWay)
// 	config.SetTimeout(20 * time.Second)

// 	monitorClient := monitorSdkClient.NewMonitorClient(credentials)
// 	monitorClient.SetConfig(config)
// 	appCode := beego.AppConfig.DefaultString("openapi.appCode", "")
// 	serviceCode := beego.AppConfig.DefaultString("openapi.serviceCode", "")
// 	region := instance.Result.Region
// 	resourceId := instance.Result.InstanceID
// 	dataPoints := reqObj.DataPoints

// 	req1 := monitorSdkApis.NewPutRequest(appCode, serviceCode, region, resourceId, dataPoints)

// 	resp1, err := monitorClient.Put(req1)
// 	logger.Infof("%s Put endpoint: %s, cloudmonitor Response:%+v", reqObj.SN, gateWay, resp1)
// 	if err != nil {
// 		logger.Warnf("%s cloudmonitor put Error: %+v", reqObj.SN, resp1.Error)
// 		return nil, err
// 	}
// 	return resp1, nil
// }
