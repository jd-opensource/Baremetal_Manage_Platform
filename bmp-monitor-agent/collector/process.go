package collector

import (
	"coding.jd.com/bmp/cps-agent/api/models"

	sdkmodels "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"
	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/host"
)

func init() {
	Register("process", &ProcessCollector{})
}

type ProcessCollector struct {
}

func (c *ProcessCollector) Collect() []sdkmodels.DataPointX {

	var dps []sdkmodels.DataPointX
	var processTags interface{} = map[string]interface{}{}

	if v, err := host.Info(); err != nil {
		logs.Warn("Collect HOST Error: %s", err.Error())
		return dps
	} else {
		/*
			METRIC_PROCESS_TOTAL           string = "cps.process.total"
		*/

		dps = append(dps, sdkmodels.DataPointX{
			Metric: models.METRIC_PROCESS_TOTAL,
			Value:  v.Procs,
			Tags:   &processTags,
		})

	}

	return dps

}
