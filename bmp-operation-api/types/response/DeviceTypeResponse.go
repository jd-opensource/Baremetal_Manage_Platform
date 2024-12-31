package response

import (
	"encoding/json"
)

type DeviceType struct {
	// 额外拼接的字段，省去前端工作
	CpuInfo    string `json:"cpuInfo"`
	MemInfo    string `json:"memInfo"` // 内存信息拼接
	SvInfo     string `json:"svInfo"`  // 系统盘信息拼接
	DvInfo     string `json:"dvInfo"`  // 数据盘信息拼接
	GpuInfo    string `json:"gpuInfo"`
	NicInfo    string `json:"nicInfo"` // 2*10GE
	Nic        string `json:"nic"`     // 单网卡 or 双网口
	ImageCount string `json:"imageCount"`

	ID               int    `json:"id"`
	IDcID            string `json:"idcId"`
	IdcName          string `json:"idcName"`
	IdcNameEn        string `json:"idcNameEn"`
	DeviceTypeID     string `json:"deviceTypeId"`     // 设备类型uuid
	Name             string `json:"name"`             // 机型名称，如计算效能型,标准计算型
	DeviceType       string `json:"deviceType"`       // 机型规格, cps.c.normal
	DeviceSeries     string `json:"deviceSeries"`     // 机型类型，如计算型，存储型
	DeviceSeriesName string `json:"deviceSeriesName"` // 机型类型，如计算型，存储型
	Architecture     string `json:"architecture"`     // 体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64
	Height           int    `json:"height"`           // 【高度（U）】：显示机型高度
	Description      string `json:"description"`      // 描述
	CPUAmount        int64  `json:"cpuAmount"`        // cpu数量
	CPUCores         int64  `json:"cpuCores"`         // 单个cpu内核数

	CPUManufacturer   string `json:"cpuManufacturer"`   // cpu厂商
	CPUModel          string `json:"cpuModel"`          // cpu处理器型号
	CPUFrequency      string `json:"cpuFrequency"`      // cpu频率(G)
	MemType           string `json:"memType"`           // 内存接口（如DDR3，DDR4）
	MemSize           int    `json:"memSize"`           // 单个内存大小(GB)
	MemAmount         int    `json:"memAmount"`         // 内存数量
	MemFrequency      int    `json:"memFrequency"`      // 内存主频（MHz)
	NicAmount         int64  `json:"nicAmount"`         // 网卡数量
	NicRate           int    `json:"nicRate"`           // 网卡传输速率(GE)
	InterfaceMode     string `json:"interfaceMode"`     // 【网口模式】【网络设置】: bond单网口,dual双网口
	InterfaceModeName string `json:"interfaceModeName"` // 【网口模式】【网络设置】: bond单网口,dual双网口

	// 是否做raid，[RAID/NO RAID]
	RaidCan         string `json:"raidCan"`
	GpuAmount       int64  `json:"gpuAmount"`       // gpu数量
	GpuManufacturer string `json:"gpuManufacturer"` // gpu厂商
	GpuModel        string `json:"gpuModel"`        // gpu处理器型号

	CreatedBy   string `json:"createdBy"`   // 创建者
	UpdatedBy   string `json:"updatedBy"`   // 更新者
	CreatedTime string `json:"createdTime"` // 创建时间
	UpdatedTime string `json:"updatedTime"` // 更新时间
	//DeletedTime               int    `json:"deletedTime"`               // 删除时间
	//IsDel                     int8   `son:"isDel"`                      // 是否删除0未删除 1已删除
	//Raid string `json:"raid"` //系统盘raid

	CPUSpec string `json:"cpuSpec"` // gpu处理器型号
	MemSpec string `json:"memSpec"` // gpu处理器型号

	InstanceCount  int      `json:"instanceCount"`
	DeviceCount    int      `json:"deviceCount"`
	InstanceStatus []string `json:"instanceStatus"`
	//boot模式
	BootMode string `json:"bootMode"`
	//Volumes    []*sdkModels.Volume `json:"volumes"`
	Volumes    []Volume `json:"volumes"`
	IsNeedRaid string   `json:"isNeedRaid"`
}
type Volume struct {

	// disk type
	DiskType string `json:"diskType"`

	// interface type
	InterfaceType string `json:"interfaceType"`

	// raid
	Raid string `json:"raid"`

	// raid can
	RaidCan string `json:"raidCan"`

	// raid Id
	RaidID string `json:"raidId"`

	// volume amount
	VolumeAmount int64 `json:"volumeAmount"`

	// volume ID
	VolumeID string `json:"volumeId"`

	// volume name
	VolumeName string `json:"volumeName"`

	// volume size
	VolumeSize string `json:"volumeSize"`

	// volume type
	VolumeType string `json:"volumeType"`

	// volume unit
	VolumeUnit string `json:"volumeUnit"`
}
type DeviceTypePage struct {
	DeviceTypes []DeviceType `json:"deviceTypes"`
	PagingResponse
}

type DeviceTypeInfo struct {
	DeviceType *DeviceType `json:"deviceType"`
}

type CreateDeviceTypeResult struct {
	DeviceTypeId string `json:"deviceTypeId"`
}

func (d DeviceType) MarshalJSON() ([]byte, error) {
	type Alias DeviceType
	return json.Marshal(struct {
		Alias
		//CreateTime string `json:"create_time"`
		//UpdateTime string `json:"update_time"`
	}{
		Alias: Alias(d),
		//CreateTime: d.CreateTime.Format("2006-01-02 15:04:05"),
		//UpdateTime: d.UpdateTime.Format("2006-01-02 15:04:05"),
	})
}

type AvailableDeviceType struct {
	Computer []DeviceTypeReinstall `json:"computer,omitempty"`
	Storage  []DeviceTypeReinstall `json:"storage,omitempty"`
	Gpu      []DeviceTypeReinstall `json:"gpu,omitempty"`
	Other    []DeviceTypeReinstall `json:"other,omitempty"`
}

type DeviceTypeReinstall struct {
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
}
