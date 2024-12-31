package collector

import (
	"coding.jd.com/bmp/cps-agent/api/models"

	sdkmodels "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"
	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/net"
)

func init() {
	Register("tcp", &TcpCollector{})
}

type TcpCollector struct {
}

func (c *TcpCollector) Collect() []sdkmodels.DataPointX {

	var dps []sdkmodels.DataPointX
	var tcpTags interface{} = map[string]interface{}{}

	if conns, err := net.Connections("tcp"); err != nil {
		logs.Warn("Collect TCP Error: %s", err.Error())
		return dps
	} else {
		total := len(conns)
		established := 0
		for _, conn := range conns {
			if conn.Status == "ESTABLISHED" {
				established++
			}
		}
		/*
			METRIC_TCP_CONNECT_TOTAL       string = "cps.tcp.connect.total"
			METRIC_TCP_CONNECT_ESTABLISHED string = "cps.tcp.connect.established"
		*/

		dps = append(dps, sdkmodels.DataPointX{
			Metric: models.METRIC_TCP_CONNECT_TOTAL,
			Value:  total,
			Tags:   &tcpTags,
		})

		dps = append(dps, sdkmodels.DataPointX{
			Metric: models.METRIC_TCP_CONNECT_ESTABLISHED,
			Value:  established,
			Tags:   &tcpTags,
		})
	}

	return dps

}
