package response

type DeviceType struct {
	Name         string `json:"name"`
	DeviceTypeId string `json:"deviceTypeId"`
	Cpu          string `json:"cpu"`
	Mem          string `json:"mem"`
	System       string `json:"system"`
	Data         string `json:"data"`
	Nic          string `json:"nic"`
	Gpu          string `json:"gpu"`
	// IsStockAvailable 是否有库存
	AvailableStock   int64  `json:"availableStock"`
	DeviceSeries     string `json:"deviceSeries"`
	DeviceSeriesName string `json:"deviceSeriesName"`
	//boot模式
	BootMode string `json:"boot_mode"`
	//是否需要做raid
	IsNeedRaid string `json:"isNeedRaid"`
}

type AvailableDeviceType struct {
	Computer []DeviceType `json:"computer,omitempty"`
	Storage  []DeviceType `json:"storage,omitempty"`
	Gpu      []DeviceType `json:"gpu,omitempty"`
	Other    []DeviceType `json:"other,omitempty"`
}
