package collector

import (
	"runtime"

	"coding.jd.com/bmp/cps-agent/api/models"
	"coding.jd.com/bmp/cps-agent/util"

	sdkmodels "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"

	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/load"
)

func init() {
	Register("load", &LoadCollector{})
}

type LoadCollector struct {
}

func (c *LoadCollector) Collect() []sdkmodels.DataPointX {

	var dps []sdkmodels.DataPointX
	var loadTags interface{} = map[string]interface{}{}
	if runtime.GOOS == "windows" {
		return dps
	}

	if v, err := load.Avg(); err != nil {
		logs.Warn("Collect LOAD Error: %s", err.Error())
		return dps
	} else {
		/*
			METRIC_AVG_LOAD1
			METRIC_AVG_LOAD5               string = "cps.avg.load5"
			METRIC_AVG_LOAD15              string = "cps.avg.load15"
		*/

		dps = append(dps,
			sdkmodels.DataPointX{
				Metric: models.METRIC_AVG_LOAD1,
				Value:  util.FloatRound(v.Load1, 2),
				Tags:   &loadTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_AVG_LOAD5,
				Value:  util.FloatRound(v.Load5, 2),
				Tags:   &loadTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_AVG_LOAD15,
				Value:  util.FloatRound(v.Load15, 2),
				Tags:   &loadTags,
			},
		)
	}

	return dps

}
