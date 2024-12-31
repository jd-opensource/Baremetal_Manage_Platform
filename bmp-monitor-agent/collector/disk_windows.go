//go:build windows
// +build windows

package collector

import (
	"coding.jd.com/bmp/cps-agent/api/models"
	"coding.jd.com/bmp/cps-agent/util"

	sdkmodels "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"

	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/disk"
)

// TODO: IOCounters需要修改gopsutil源码
func (c *DiskCollector) Collect() []sdkmodels.DataPointX {
	var dps []sdkmodels.DataPointX

	dp, err := disk.Partitions(true)
	if err != nil {
		logs.Warn("Collect Disk Error: %s", err.Error())
		return dps
	}

	for _, tmp := range dp {
		du, duErr := disk.Usage(tmp.Mountpoint)
		if duErr != nil {
			logs.Warn("Collect Disk Usage Error: %s", duErr.Error())
			continue
		}

		dio, dioErr := disk.IOCounters(tmp.Device)
		if dioErr != nil {
			logs.Warn("Collect Disk IO Error: %s", dioErr.Error())
			continue
		}

		/*
			METRIC_DISK_USED
			METRIC_DISK_UTIL
			METRIC_DISK_BYTES_READ
			METRIC_DISK_BYTES_WRITE
			METRIC_DISK_IOPS_READ
			METRIC_DISK_IOPS_WRITE
		*/
		var diskTags interface{}
		diskTags = map[string]interface{}{"device": tmp.Mountpoint}
		dps = append(dps,
			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_USED,
				Value:  du.Used,
				Tags:   &diskTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_UTIL,
				Value:  util.FloatRound(du.UsedPercent, 2),
				Tags:   &diskTags,
			},
			sdkmodels.DataPointX{ //新增
				Metric: models.METRIC_DISK_MOUNTPOINT_USED,
				Value:  du.Used,
				Tags:   &diskTags,
			},
			sdkmodels.DataPointX{ //新增
				Metric: models.METRIC_DISK_MOUNTPOINT_UTIL,
				Value:  util.FloatRound(du.UsedPercent, 2),
				Tags:   &diskTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_BYTES_READ,
				Value:  dio[tmp.Device].ReadBytes,
				Tags:   &diskTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_BYTES_WRITE,
				Value:  dio[tmp.Device].WriteBytes,
				Tags:   &diskTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_IOPS_READ,
				Value:  dio[tmp.Device].ReadCount,
				Tags:   &diskTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_IOPS_WRITE,
				Value:  dio[tmp.Device].WriteCount,
				Tags:   &diskTags,
			},
		)

	}

	return dps
}
