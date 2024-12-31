package models

const (
	//METRIC_NAME_BMP = "bmp_monitor_gauge"
	METRIC_NAME_BMP_GAUGE = "bmp_monitor_gauge"
	METRIC_NAME_BMP_COUNTER = "bmp_monitor_counter"

)

// 告警规则中的单位转换
var MetricNameUnitConversionMap = map[string]float64{
	"bmp.memory.used" : 1000000000,
	"bmp.disk.used" : 1000000000,
}
