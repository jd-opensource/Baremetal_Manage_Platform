package models

const (
	APP_CODE     string = "jcloud"
	SERVICE_CODE string = "cps"

	METRIC_CPU_UTIL                string = "bmp.cpu.util"
	METRIC_MEMORY_USED             string = "bmp.memory.used"
	METRIC_MEMORY_UTIL             string = "bmp.memory.util"
	METRIC_DISK_USED               string = "bmp.disk.used"
	METRIC_DISK_UTIL               string = "bmp.disk.util"
	METRIC_DISK_MOUNTPOINT_USED    string = "bmp.disk.partition.used" //新增
	METRIC_DISK_MOUNTPOINT_UTIL    string = "bmp.disk.partition.util" //新增
	METRIC_DISK_BYTES_READ         string = "bmp.disk.bytes.read"
	METRIC_DISK_BYTES_WRITE        string = "bmp.disk.bytes.write"
	METRIC_DISK_IOPS_READ          string = "bmp.disk.counts.read"
	METRIC_DISK_IOPS_WRITE         string = "bmp.disk.counts.write"
	METRIC_NETWORK_BYTES_IN        string = "bmp.network.bytes.ingress"
	METRIC_NETWORK_BYTES_OUT       string = "bmp.network.bytes.egress"
	METRIC_NETWORK_PACKETS_IN      string = "bmp.network.packets.ingress"
	METRIC_NETWORK_PACKETS_OUT     string = "bmp.network.packets.egress"
	METRIC_AVG_LOAD1               string = "bmp.avg.load1"
	METRIC_AVG_LOAD5               string = "bmp.avg.load5"
	METRIC_AVG_LOAD15              string = "bmp.avg.load15"
	METRIC_TCP_CONNECT_TOTAL       string = "bmp.tcp.connect.total"
	METRIC_TCP_CONNECT_ESTABLISHED string = "bmp.tcp.connect.established"
	METRIC_PROCESS_TOTAL           string = "bmp.process.total"
)
