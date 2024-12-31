package collector

import (
	"coding.jd.com/bmp/cps-agent/api/models"
	"coding.jd.com/bmp/cps-agent/util"

	sdkmodels "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"

	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/mem"
)

func init() {
	Register("mem", &MemCollector{})
}

type MemCollector struct {
}

func (c *MemCollector) Collect() []sdkmodels.DataPointX {

	var dps []sdkmodels.DataPointX
	var memTags interface{} = map[string]interface{}{}
	if v, err := mem.VirtualMemory(); err != nil {
		logs.Warn("Collect MEM Error: %s", err.Error())
		return dps
	} else {
		/*
			METRIC_MEMORY_USED             string = "cps.memory.used"
			METRIC_MEMORY_UTIL            string = "cps.memory.usage"
		*/

		dps = append(dps,
			sdkmodels.DataPointX{
				Metric: models.METRIC_MEMORY_USED,
				Value:  v.Used,
				Tags:   &memTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_MEMORY_UTIL,
				Value:  util.FloatRound(v.UsedPercent, 2),
				Tags:   &memTags,
			},
		)
	}

	return dps

}
