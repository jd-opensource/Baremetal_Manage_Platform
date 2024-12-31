package response

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/raidDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/volumeDao"
)

type Instance struct {
	ID int `json:"-"`
	// 实例uuid
	InstanceID string `json:"instanceId"` // 实例ID（uuid）
	// 机房uuid
	IDcID string `json:"idcId"`
	// 机房名称
	IdcName string `json:"idcName"`
	// 实例所属项目UUID
	ProjectID string `json:"projectId"`
	// 实例所属用户UUID
	UserID string `json:"userId"`
	// 实例名称
	InstanceName string `json:"instanceName"`
	// 设备uuid
	DeviceID string `json:"deviceId"`
	// 设备SN
	Sn string `json:"sn"`
	// 带外管理IP
	IloIP string `json:"iloIp"`
	// 内网IPV4
	PrivateIPv4 string `json:"privateIpv4"`
	// 内网IPV6
	PrivateIPv6 string `json:"privateIpv6"`
	// 机型uuid
	DeviceTypeID string `json:"deviceTypeId"`
	// 机型名称，如计算效能型,标准计算型
	DeviceTypeName string `json:"deviceTypeName"`
	// 机型规格, cps.c.normal
	DeviceType string `json:"deviceType"`
	// 机型类型，如computer
	DeviceSeries string `json:"deviceSeries"`
	//机型类型，如计算型，存储型
	DeviceSeriesName string `json:"deviceSeriesName"`
	// cpu数量
	CPUAmount int `json:"cpuAmount"`
	// 单个cpu内核数
	CPUCores int `json:"cpuCores"`
	// cpu厂商
	CPUManufacturer string `json:"cpuManufacturer"`
	// cpu处理器型号
	CPUModel string `json:"cpuModel"`
	// cpu频率(GHz)
	CPUFrequency string `json:"cpuFrequency"`
	// 内存接口（如DDR3，DDR4）
	MemType string `json:"memType"`
	// 单个内存大小(GB)
	MemSize int `json:"memSize"`
	// 内存数量
	MemAmount int `json:"memAmount"`
	// 内存主频（MHz)
	MemFrequency int `json:"memFrequency"`
	// 网卡数量
	NicAmount int `json:"nicAmount"`
	// 网卡传输速率(GE)
	NicRate int `json:"nicRate"`
	// 【网口模式】【网络设置】: bond单网口,dual双网口
	InterfaceMode string `json:"interfaceMode"`
	// 系统盘类型（SSD，HDD）
	SystemVolumeType string `json:"systemVolumeType"`
	// 系统盘接口类型（SATA,SAS,NVME）
	SystemVolumeInterfaceType string `json:"systemVolumeInterfaceType"`
	// 系统盘单盘大小
	SystemVolumeSize int `json:"systemVolumeSize"`
	// 系统盘数量
	SystemVolumeAmount int `json:"systemVolumeAmount"`
	// 系统盘单位
	SystemVolumeUnit string `json:"systemVolumeUnit"`
	// gpu数量
	GpuAmount int `json:"gpuAmount"`
	// gpu厂商
	GpuManufacturer string `json:"gpuManufacturer"`
	// gpu处理器型号
	GpuModel string `json:"gpuModel"`
	// 数据盘类型（SSD，HDD）
	DataVolumeType string `json:"dataVolumeType"`
	// 数据盘接口类型（SATA,SAS,NVME）
	DataVolumeInterfaceType string `json:"dataVolumeInterfaceType"`
	// 数据盘单盘大小
	DataVolumeSize int `json:"dataVolumeSize"`
	// 数据盘数量
	DataVolumeAmount int `json:"dataVolumeAmount"`
	// 系统盘单位
	DataVolumeUnit string `json:"dataVolumeUnit"`
	// 主机名
	Hostname string `json:"hostname"`
	// 运行状态
	Status string `json:"status"`
	// 实例错误状态时的错误原因
	Reason string `json:"reason"`
	// 运行状态中文名字
	StatusName string `json:"statusName"`
	// 镜像uuid
	ImageID string `json:"imageId"`
	// 镜像名称
	ImageName string `json:"imageName"`
	// 系统盘raidId
	SystemVolumeRaidID string `json:"systemVolumeRaidId"`
	// 系统盘raid名称
	SystemVolumeRaidName string `json:"systemVolumeRaidName"`
	//系统盘raid配置
	RaidCan string `json:"raidCan"`
	// 是否锁定解锁锁定:locked,解锁unlocked
	Locked string `json:"locked"`
	// 是否锁定解锁锁定:已解锁,已锁定
	LockedName string `json:"lockedName"`
	// 数据盘raidId
	DataVolumeRaidID string `json:"dataVolumeRaidId"`
	// 数据盘raid名称
	DataVolumeRaidName string `json:"DataVolumeRaidName"`
	// 实例描述
	Description string `json:"description"`
	// 创建者
	CreatedBy string `json:"createdBy"`
	// 更新者
	UpdatedBy string `json:"updatedBy"`
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// 更新时间
	UpdatedTime string `json:"updatedTime"`

	//实例的volume和raid信息
	InstanceVolumeRaids []*InstanceVolumeRaid `json:"volumeRaid"`
}
type InstanceList struct {
	// instance实体数组
	Instances  []*Instance `json:"instances"`
	PageNumber int64       `json:"pageNumber"`
	PageSize   int64       `json:"pageSize"`
	TotalCount int64       `json:"totalCount"`
}
type InstanceInfo struct {
	// instance实体
	Instance *Instance `json:"instance"`
}
type InstanceId struct {
	InstanceId string `json:"instanceId"`
}
type InstanceIds struct {
	// instanceId 列表
	InstanceIds []string `json:"instanceIds"`
}

type InstanceVolumeRaid struct {
	Volume *volumeDao.Volume `json:"volume"`
	Raid   *raidDao.Raid     `json:"raid"`
}
