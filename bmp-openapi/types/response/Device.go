package response

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/diskDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/raidDao"
)

type Device struct {
	// 机房uuid
	// required: true
	IdcID string `json:"idcId"`
	//机型uuid
	DeviceTypeID string `json:"deviceTypeId"`
	// 设备ID编号
	ID int `json:"id"`
	// 设备SN
	Sn string `json:"sn"`
	// 设备uuid
	DeviceID string `json:"deviceId"`
	// 实例id
	InstanceID string `json:"instanceId"`
	// 设备状态: 未装机，已装机
	ManageStatus string `json:"manageStatus"`
	// 设备状态: 未装机，已装机
	ManageStatusName string `json:"manageStatusName"`
	// 实例是否锁定 锁定locked 未锁定unlocked
	Locked string `json:"locked"`
	//品牌
	Brand string `json:"brand"`
	//型号
	Model string `json:"model"`
	// 设备状态变更失败原因
	Reason string `json:"reason"`
	// 机柜编码
	Cabinet string `json:"cabinet"`
	// U位
	UPosition string `json:"uPosition"`
	// 带外管理IP
	IloIP string `json:"iloIp"`
	// 带外账号
	IloUser string `json:"iloUser"`
	// 带外账号密码
	IloPassword string `json:"iloPassword"`
	// MAC1（eth0）
	Mac1 string `gorm:"column:mac1" json:"mac1"`
	// MAC2（eth2）
	Mac2 string `gorm:"column:mac2" json:"mac2"`
	// 交换机1ip
	SwitchIP1 string `gorm:"column:switch_ip1" json:"switchIp1"`
	// 交换机1port
	SwitchPort1 string `gorm:"column:switch_port1" json:"switchPort1"`
	// 交换机2ip
	SwitchIP2 string `gorm:"column:switch_ip2" json:"switchIp2"`
	// 交换机2port
	SwitchPort2 string `gorm:"column:switch_port2" json:"switchPort2"`
	// 交换机1登录账号，如果为空，取所在机房的值
	SwitchUser1 string `json:"switchUser1"`
	// 交换机1登录密码
	SwitchPassword1 string `json:"switchPassword1"`
	// 交换机2登录账号，如果为空，取所在机房的值
	SwitchUser2 string `json:"switchUser2"`
	// 交换机2登录密码
	SwitchPassword2 string `json:"switchPassword2"`
	// 描述
	Description string `json:"description"`
	// 网口交换机IP
	SwitchIP string `json:"switchIp"`
	// 子网掩码
	Mask string `json:"mask"`
	// ETH1子网掩码
	Eth1Mask string `json:"eth1Mask"`
	// 网关地址
	Gateway string `json:"gateway"`
	// 内网IPV4
	PrivateIPv4 string `json:"privateIpv4"`
	// 内网IPV6
	PrivateIPv6 string `json:"privateIpv6"`
	// eth1内网IPV4
	PrivateEth1IPv4 string `json:"privateEth1Ipv4"`
	// eth1内网IPV6
	PrivateEth1IPv6 string `json:"privateEth1Ipv6"`
	// adapter_id
	AdapterID int `json:"adapterId"`
	// raid工具：（megacli64等）
	RaidDriver string `json:"raidDriver"`
	//添加到disk表, 系统盘1背板号
	Enclosure1 string `json:"enclosure1"`
	//系统盘1槽位
	Slot1 int `json:"slot1"`
	//系统盘2背板号
	Enclosure2 string `json:"Enclosure2"`
	//系统盘2槽位
	Slot2 int `json:"slot2"`
	//以后还能还有数据盘，todo

	// 创建者
	CreatedBy string `json:"createdBy"`
	// 更新者
	UpdatedBy string `json:"updatedBy"`
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// 更新时间
	UpdatedTime string `json:"updatedTime"`
	//DeletedTime     int    `json:"deletedTime"`         // 删除时间戳
	//IsDel           int8   `json:"isDel"`                     // 是否删除0未删除 1已删除

	//以下是通过别的表查询出来的数据，用于返回给上层

	//机房名称
	IdcName string `json:"idcName"`
	//idcname
	IDcNameEn string `json:"idcNameEn"`
	//机型名称
	DeviceTypeName string `json:"deviceTypeName"`

	// 实例名称
	InstanceName string `json:"instanceName"`
	// 实例镜像名称
	ImageName string `json:"imageName"`
	// 实例状态
	InstanceStatus string `json:"instanceStatus"`
	// 实例状态
	InstanceStatusName string `json:"instanceStatusName"`
	// 实例失败原因
	InstanceReason string `json:"instanceReason"`
	// 实例归属用户uuid
	UserId string `json:"userId"`
	// 实例归属用户名称
	UserName string `json:"userName"`
	// 实例创建时间
	InstanceCreatedTime string `json:"instanceCreatedTime"`
	// 实例描述
	InstanceDescription string `json:"instanceDescription"`
	//机型配置

	//型号，机型规格cps.normal
	DeviceType string `json:"deviceType"`
	//计算型，存储型
	DeviceSeries string `json:"deviceSeries"`
	//计算型，存储型
	DeviceSeriesName string `json:"deviceSeriesName"`
	// 体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64
	Architecture string `json:"architecture"`
	// cpu数量
	CPUAmount int `json:"cpuAmount"`
	// 单个cpu内核数
	CPUCores int `json:"cpuCores"`
	// 单个cpu路数
	CPURoads int8 `json:"cpuRoads"`
	// cpu厂商
	CPUManufacturer string `json:"cpuManufacturer"`
	// cpu处理器型号
	CPUModel string `json:"cpuModel"`
	// cpu频率(G)
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
	//SystemVolumeUnit          string `json:"systemVolumeUnit"`          // 系统盘单位（GB，TB）

	// 系统盘数量
	SystemVolumeAmount int `json:"systemVolumeAmount"`
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
	//DataVolumeUnit            string `json:"dataVolumeUnit"`            // 数据盘单位（GB，TB）

	// 数据盘数量
	DataVolumeAmount int `json:"dataVolumeAmount"`

	//拼装信息
	CpuInfo string `json:"cpuInfo"`
	MemInfo string `json:"memInfo"`
	SvInfo  string `json:"svInfo"`
	DvInfo  string `json:"dvInfo"`
	GpuInfo string `json:"gpuInfo"`
	NicInfo string `json:"nicInfo"`
	// "1"表示采集成功，"2"表示未采集,3表示采集中，4表示采集失败
	CollectStatus     string `json:"collectStatus"`
	CollectFailReason string `json:"collectFailReason"`
	IsNeedRaid        string `json:"isNeedRaid"`
}

type DeviceList struct {
	// 设备列表
	Devices    []*Device `json:"devices"`
	PageNumber int64     `json:"pageNumber"`
	PageSize   int64     `json:"pageSize"`
	TotalCount int64     `json:"totalCount"`
}
type DeviceInfo struct {
	// 设备
	Device *Device `json:"device"`
}
type DeviceIds struct {
	// 设备id列表
	DeviceIds []string `json:"deviceIds"`
}
type DeviceStock struct {
	// 库存量
	Stocks int `json:"stocks"`
}

func (d Device) MarshalJSON() ([]byte, error) {
	type Alias Device
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

type DeviceDisks struct {
	Disks   []*diskDao.Disk `json:"disks"`
	Panfu   []*diskDao.Disk `json:"panfu"`
	Volumes []*VolumeIt     `json:"volumes"`
}

type VolumeIt struct {
	Volume
	Disks []*diskDao.Disk `json:"disks"` // disk uuid
	Raids []*raidDao.Raid `json:"raids"`
}
