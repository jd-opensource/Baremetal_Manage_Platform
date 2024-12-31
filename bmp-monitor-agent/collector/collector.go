package collector

import (
	"time"

	"coding.jd.com/bmp/cps-agent/global"

	sdkmodels "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"

	"github.com/astaxie/beego/logs"
)

var (
	Collectors map[string]Collector
)

func init() {
	Collectors = make(map[string]Collector)
}

type Collector interface {
	Collect() []sdkmodels.DataPointX
}

func Register(name string, c Collector) {
	Collectors[name] = c
}

func Collect(timeout time.Duration) []sdkmodels.DataPointX {
	var dps []sdkmodels.DataPointX

	done := make(chan []sdkmodels.DataPointX)

	for k, v := range Collectors {

		go func(name string, c Collector) {
			logs.Debug("collector %s Start", name)
			done <- c.Collect()
		}(k, v)

		select {
		case <-time.After(timeout):
			logs.Warn("collector %s Timeout", k)
		case dp := <-done:
			if len(dp) > 0 {
				dps = append(dps, dp...)
			}
			logs.Debug("collector %s Finish", k)
		}

	}

	// align timestamp
	curr := time.Now().Unix()
	t := curr - (curr % global.DefaultInt64(global.AgentConfig.Monitor.TimestampAlign, 60))

	for i := range dps {
		dps[i].Timestamp = t
		// if dps[i].Tags == nil {
		// 	dps[i].Tags = &map[string]interface{}{}
		// }
	}

	return dps
}
