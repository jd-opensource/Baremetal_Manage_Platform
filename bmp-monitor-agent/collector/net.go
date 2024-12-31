package collector

import (
	"coding.jd.com/bmp/cps-agent/api/models"

	sdkmodels "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"
	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/net"
)

var (
	NETWORK_FLAGS_IGNORE = map[string]struct{}{
		"loopback": {},
	}
	NETWORK_FLAGS_INCLUDE = map[string]struct{}{
		"up": {},
	}
)

func init() {
	Register("network", &NetCollector{})
}

type NetCollector struct {
}

func (c *NetCollector) Collect() []sdkmodels.DataPointX {
	var dps []sdkmodels.DataPointX

	retNI := map[string]net.InterfaceStat{}
	nv, err := net.Interfaces()
	if err != nil {
		logs.Warn("Collect NET Error: %s", err.Error())
		return dps
	}

	for _, tmp := range nv {
		if tmp.HardwareAddr == "" {
			continue
		}

		if ignoreNetworkInterface(tmp.Flags) {
			continue
		}

		retNI[tmp.Name] = tmp

	}

	nio, errIO := net.IOCounters(true)
	if errIO != nil {
		logs.Warn("Collect NETIO Error: %s", errIO.Error())
		return dps
	}

	var allNio = net.IOCountersStat{
		Name: "all",
	}
	for _, v := range nio {
		if _, exist := retNI[v.Name]; exist {
			allNio.BytesSent += v.BytesSent
			allNio.BytesRecv += v.BytesRecv
			allNio.PacketsSent += v.PacketsSent
			allNio.PacketsRecv += v.PacketsRecv
		}
	}
	/*
		METRIC_NETWORK_BYTES_IN        string = "cps.network.bytes.incoming"
		METRIC_NETWORK_BYTES_OUT       string = "cps.network.bytes.outgoing"
		METRIC_NETWORK_PACKETS_IN      string = "cps.network.packets.incoming"
		METRIC_NETWORK_PACKETS_OUT     string = "cps.network.packets.outgoing"
	*/
	//汇总
	var allnetTags interface{} = map[string]interface{}{}
	dps = append(dps,
		sdkmodels.DataPointX{
			Metric: models.METRIC_NETWORK_BYTES_IN,
			Value:  allNio.BytesRecv,
			Tags:   &allnetTags,
		},
		sdkmodels.DataPointX{
			Metric: models.METRIC_NETWORK_BYTES_OUT,
			Value:  allNio.BytesSent,
			Tags:   &allnetTags,
		},
		sdkmodels.DataPointX{
			Metric: models.METRIC_NETWORK_PACKETS_IN,
			Value:  allNio.PacketsRecv,
			Tags:   &allnetTags,
		},
		sdkmodels.DataPointX{
			Metric: models.METRIC_NETWORK_PACKETS_OUT,
			Value:  allNio.PacketsSent,
			Tags:   &allnetTags,
		},
	)

	//分网卡
	for _, v := range nio {
		var netTags interface{} = map[string]interface{}{"device": v.Name}
		if _, exist := retNI[v.Name]; exist {
			dps = append(dps,
				sdkmodels.DataPointX{
					Metric: models.METRIC_NETWORK_BYTES_IN,
					Value:  v.BytesRecv,
					Tags:   &netTags,
				},
				sdkmodels.DataPointX{
					Metric: models.METRIC_NETWORK_BYTES_OUT,
					Value:  v.BytesSent,
					Tags:   &netTags,
				},
				sdkmodels.DataPointX{
					Metric: models.METRIC_NETWORK_PACKETS_IN,
					Value:  v.PacketsRecv,
					Tags:   &netTags,
				},
				sdkmodels.DataPointX{
					Metric: models.METRIC_NETWORK_PACKETS_OUT,
					Value:  v.PacketsSent,
					Tags:   &netTags,
				},
			)
		}
	}

	return dps
}

func ignoreNetworkInterface(ni []string) bool {
	for _, v := range ni {
		if _, exist := NETWORK_FLAGS_IGNORE[v]; exist {
			return true
		}
	}

	for _, v := range ni {
		if _, exist := NETWORK_FLAGS_INCLUDE[v]; exist {
			return false
		}
	}

	return true
}
