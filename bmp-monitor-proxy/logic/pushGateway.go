package logic

import (
	"fmt"

	"coding.jd.com/bmp/agent-proxy-jdstack/api/models"
	"coding.jd.com/bmp/agent-proxy-jdstack/util"
	log "git.jd.com/cps-golang/log"

	"github.com/beego/beego/v2/server/web"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var gatewayUrl string

func Init() error {
	gateWayHost, err := web.AppConfig.String("prometheus_gateway_host")
	if err != nil {
		return err
	}
	gateWayPort, err := web.AppConfig.String("prometheus_gateway_port")
	if err != nil {
		return err
	}
	gatewayUrl = fmt.Sprintf("http://%s:%s", gateWayHost, gateWayPort)
	return nil
}

func PushGateway(logger *log.Logger, instance *models.Instance, reqObj models.ProxyPut) error {

	datas := reqObj.DataPoints
	counterMetrics := []string{
		"bmp.disk.bytes.read",
		"bmp.disk.bytes.write",
		"bmp.disk.counts.read",
		"bmp.disk.counts.write",
		"bmp.network.bytes.ingress",
		"bmp.network.bytes.egress",
		"bmp.network.packets.ingress",
		"bmp.network.packets.egress",
	}

	// currentCacheData := map[string]interface{}{}

	// prevCache, err := util.HGetAllFromRedis("PushGatewayCache_" + instance.Result.Instance.InstanceID)
	// if err != nil {
	// 	logger.Warnf("PushGateway.HGetAllFromRedis error, instanceId:%s, error:%s", instance.Result.Instance.InstanceID, err.Error())
	// } else {
	// 	tmm, _ := json.Marshal(prevCache)
	// 	logger.Infof("push.getCacheData instanceId:%s, data:%s", instance.Result.Instance.InstanceID, string(tmm))
	// }

	for _, data := range datas {
		isCounter := false
		if util.InArray4String(data.Metric, counterMetrics) {
			isCounter = true
		}

		var vvv float64
		// if isCounter { //counter指标
		// 	var tagJoinKey string = data.Metric
		// 	for k, v := range data.Tags {
		// 		tagJoinKey += fmt.Sprintf("_%s_%s", k, v)
		// 	}
		// 	currentCacheData[tagJoinKey] = data.Value

		// 	s, err := strconv.ParseFloat(prevCache[tagJoinKey], 64)
		// 	logger.Infof("debug here", s)
		// 	if err != nil {
		// 		logger.Warnf("PushGateway.ParseFloat error,key:%s  v:%s, error:%s", tagJoinKey, prevCache[tagJoinKey], err.Error())
		// 		// continue
		// 	} else {
		// 		// vvv = data.Value - s //做差值
		// 		vvv = data.Value //不做差值
		// 	}

		// } else { //gause指标
		// 	vvv = data.Value
		// }
		vvv = data.Value
		var pusher *push.Pusher
		logger.Infof("metric is:%s, tag is:%v, isCounter:%v, value:%f, onput_value:%f", data.Metric, data.Tags, isCounter, vvv, data.Value)
		if isCounter {

			c := prometheus.NewCounter(prometheus.CounterOpts{
				Name: "bmp_monitor_counter",
				Help: "bmp-monitor-counter-data, store it",
			})
			c.Add(vvv)
			pusher = push.New(gatewayUrl, "bmp_monitor_counter"). // push.New("pushgateway地址", "job名称")
										Collector(c)
		} else {
			g := prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "bmp_monitor_gauge",
				Help: "bmp-monitor-gauge-data, store it",
			})
			g.Set(vvv)                                          // set可以设置任意值（float64）
			pusher = push.New(gatewayUrl, "bmp_monitor_gauge"). // push.New("pushgateway地址", "job名称")
										Collector(g)
		}

		pusher.Grouping("instance_id", instance.Result.Instance.InstanceID).Grouping("metric_name", data.Metric)
		for k, v := range data.Tags {
			pusher.Grouping(k, v) // 给指标添加标签，可以添加多个
		}
		// logger.Infof("puuuuuuuu start time:%d", time.Now().Unix())
		err := pusher.Push()
		// logger.Infof("puuuuuuuu end time:%d", time.Now().Unix())
		if err != nil {
			logger.Warnf("Could not push completion time to Pushgateway error:%s", err.Error())
		} else {
			logger.Infof("instance_id:%s metric:%s push to gateway success, url:%s", instance.Result.Instance.InstanceID, data.Metric, gatewayUrl)
		}

	}

	// if len(currentCacheData) > 0 {
	// 	if err := util.HMSetObjectToRedis("PushGatewayCache_"+instance.Result.Instance.InstanceID, currentCacheData); err != nil {
	// 		logger.Warnf("PushGateway.HMSetObjectToRedis error, instanceId:%s, error:%s", instance.Result.Instance.InstanceID, err.Error())
	// 	} else {
	// 		tmm, _ := json.Marshal(currentCacheData)
	// 		logger.Infof("push.setCacheData succ, instanceId:%s, data:%s", instance.Result.Instance.InstanceID, string(tmm))
	// 	}
	// 	if err := util.SetObjectWithExpire("PushGatewayCache_"+instance.Result.Instance.InstanceID, 100); err != nil { //过期时间100s正好卡在1分钟后2分钟前
	// 		logger.Warnf("PushGateway.SetObjectWithExpire error, instanceId:%s, error:%s", instance.Result.Instance.InstanceID, err.Error())
	// 	}
	// }

	return nil

}
