package task

import (
	"time"

	// "encoding/json"
	// "fmt"
	// "os"

	"coding.jd.com/bmp/cps-agent/api"
	"coding.jd.com/bmp/cps-agent/api/models"
	"coding.jd.com/bmp/cps-agent/collector"
	"coding.jd.com/bmp/cps-agent/global"

	"github.com/astaxie/beego/logs"
)

type CollectTask struct {
}

func init() {
	Register(&CollectTask{})
}

func (t *CollectTask) Name() string {
	return "CollectTask"
}

func (t *CollectTask) Spec() string {
	return global.DefaultString(global.AgentConfig.Agent.Collect.Scheduler, "0/60 * * * * *")
}

func (t *CollectTask) TaskFunc() error {
	logs.Debug("Collect")

	timeout := time.Duration(global.DefaultInt(global.AgentConfig.Agent.Collect.Timeout, 1)) * time.Second
	dataPoints := collector.Collect(timeout)

	// v, _ := json.Marshal(dataPoints)
	// fmt.Println("collect datapoints is:", string(v))

	// fmt.Println("put url is:", global.AgentConfig.Proxy.URL)

	url := global.AgentConfig.Proxy.URL + "put"
	// global.SerialNumber = "minping-mock-sn" //TODO 记得注释掉
	o := models.ProxyPut{SN: global.SerialNumber, DataPoints: dataPoints}

	if err := api.Put(url, o); err != nil {
		logs.Warn("Post Put Error:%s", err.Error())
	} else {
		logs.Info("Post Put success, url:%s, sn:%s", global.AgentConfig.Proxy.URL, global.SerialNumber)
	}

	return nil
}

func (t *CollectTask) Shutdown() error {
	return nil
}
