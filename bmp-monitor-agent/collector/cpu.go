package collector

import (
	"coding.jd.com/bmp/cps-agent/api/models"
	"coding.jd.com/bmp/cps-agent/util"

	sdkmodels "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"

	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/cpu"
)

func init() {
	Register("cpu", &CPUCollector{})
}

type CPUCollector struct {
}

func (c *CPUCollector) Collect() []sdkmodels.DataPointX {
	var dps []sdkmodels.DataPointX

	cpuUtil, err := cpu.Percent(0, false)
	if err != nil {
		logs.Warn("Collect CPU Error: %s", err.Error())
		return dps
	}

	var cpuTags interface{} = map[string]interface{}{}
	for _, c := range cpuUtil {
		dps = append(dps, sdkmodels.DataPointX{
			Metric: models.METRIC_CPU_UTIL,
			// ValueType: "float",
			Value: util.FloatRound(c, 2),
			Tags:  &cpuTags,
			// Tags:   map[string]interface{}{"core": fmt.Sprintf("%d", i)},
		})

	}

	return dps
}
