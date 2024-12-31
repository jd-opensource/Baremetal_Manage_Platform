//go:build linux
// +build linux

package collector

import (
	"os/exec"
	"path/filepath"
	"strings"

	"coding.jd.com/bmp/cps-agent/api/models"
	"coding.jd.com/bmp/cps-agent/util"

	sdkmodels "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"

	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/disk"
)

var DISK_DEVICE_IGNORE = map[string]struct{}{
	"cgroup":     {},
	"debugfs":    {},
	"devpts":     {},
	"devtmpfs":   {},
	"rpc_pipefs": {},
	"rootfs":     {},
	"overlay":    {},
	"tmpfs":      {},
	"none":       {},
	"nodev":      {},
	"proc":       {},
	"hugetlbfs":  {},
	"mqueue":     {},
}

var FSFILE_PREFIX_IGNORE = []string{
	"/sys",
	"/net",
	"/misc",
	"/proc",
	"/lib",
	"/boot",
}

type DiskStat struct {
	Device      string
	Used        uint64
	Usage       float64
	ReadBytes   uint64
	WriteBytes  uint64
	ReadCounts  uint64
	WriteCounts uint64
}

func (c *DiskCollector) Collect() []sdkmodels.DataPointX {
	var dps []sdkmodels.DataPointX

	// 获取硬盘
	disks, diskErr := getDisks()
	if diskErr != nil {
		logs.Warn("Collect Disk Error: %s", diskErr.Error())
		return dps
	}

	// 获取分区
	var partitions []disk.PartitionStat
	dp, errp := disk.Partitions(false)
	if errp != nil {
		logs.Warn("Collect Partition Error: %s", errp.Error())
		return dps
	}

	for _, p := range dp {
		if _, exist := DISK_DEVICE_IGNORE[p.Device]; exist {
			continue
		}

		if strings.HasPrefix(p.Fstype, "fuse") {
			continue
		}

		if ignoreFsFile(p.Mountpoint) {
			continue
		}

		// Fix Bug JCYFBUG-1467, mountinfo duplicate, recalc disk used bytes
		isDuplicated := false
		for _, p2 := range partitions {
			if strings.EqualFold(p.Device, p2.Device) {
				isDuplicated = true
			}
		}
		if isDuplicated {
			continue
		}
		// END fix Bug JCYFBUG-1467

		partitions = append(partitions, p)
	}

	for _, d := range disks {
		dio, dioErr := disk.IOCounters(d.Device)
		if dioErr != nil {
			logs.Warn("Collect Disk IO Error: %s", dioErr.Error())
			continue
		}

		d.ReadBytes = dio[filepath.Base(d.Device)].ReadBytes
		d.WriteBytes = dio[filepath.Base(d.Device)].WriteBytes
		d.ReadCounts = dio[filepath.Base(d.Device)].ReadCount
		d.WriteCounts = dio[filepath.Base(d.Device)].WriteCount

		total := uint64(0)
		used := uint64(0)
		for _, p := range partitions {

			// /dev/sda1 has /dev/sda prefix
			if strings.HasPrefix(p.Device, d.Device) {

				du, duErr := disk.Usage(p.Mountpoint)
				if duErr != nil {
					logs.Warn("Collect Disk Usage Error: %s", duErr.Error())
					continue
				}

				if du.Total != 0 {
					var diskTags interface{} = map[string]interface{}{"device": p.Mountpoint}
					dps = append(dps,
						sdkmodels.DataPointX{ //新增
							Metric: models.METRIC_DISK_MOUNTPOINT_USED,
							Value:  du.Used,
							Tags:   &diskTags,
						},
						sdkmodels.DataPointX{ //新增
							Metric: models.METRIC_DISK_MOUNTPOINT_UTIL,
							Value:  util.FloatRound(float64(du.Used)/float64(du.Total), 2) * 100,
							Tags:   &diskTags,
						},
					)
				}
				total += du.Total
				used += du.Used
			}
		}
		d.Used = used
		if total != 0 {
			d.Usage = float64(used) / float64(total)
		}

		/*
			METRIC_DISK_USED
			METRIC_DISK_UTIL
			METRIC_DISK_BYTES_READ
			METRIC_DISK_BYTES_WRITE
			METRIC_DISK_IOPS_READ
			METRIC_DISK_IOPS_WRITE
		*/
		dev := filepath.Base(d.Device)
		var diskTags interface{} = map[string]interface{}{"device": dev}
		dps = append(dps,
			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_USED,
				Value:  d.Used,
				Tags:   &diskTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_UTIL,
				Value:  util.FloatRound(d.Usage, 2) * 100,
				Tags:   &diskTags,
			},

			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_BYTES_READ,
				Value:  d.ReadBytes,
				Tags:   &diskTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_BYTES_WRITE,
				Value:  d.WriteBytes,
				Tags:   &diskTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_IOPS_READ,
				Value:  d.ReadCounts,
				Tags:   &diskTags,
			},
			sdkmodels.DataPointX{
				Metric: models.METRIC_DISK_IOPS_WRITE,
				Value:  d.WriteCounts,
				Tags:   &diskTags,
			},
		)
	}

	return dps
}

func ignoreFsFile(fs_file string) bool {
	for _, prefix := range FSFILE_PREFIX_IGNORE {
		if strings.HasPrefix(fs_file, prefix) {
			return true
		}
	}

	return false
}

func getDisks() ([]DiskStat, error) {
	ret := []DiskStat{}

	lsblk, err := exec.LookPath("lsblk")
	if err != nil {
		return ret, err
	}
	cmd := exec.Command(lsblk, "-bn")

	out, err := cmd.Output()
	if err != nil {
		return ret, err
	}

	for _, line := range strings.Split(string(out), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}
		if strings.EqualFold(strings.TrimSpace(fields[5]), "disk") {
			ret = append(ret, DiskStat{Device: "/dev/" + fields[0]})
		}
	}

	return ret, nil
}
