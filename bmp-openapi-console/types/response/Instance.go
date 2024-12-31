package response

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
	IloIP    string `json:"iloIp"`
	Mask     string `json:"mask"`
	Eth1Mask string `json:"eth1Mask"`
	// 内网IPV4
	PrivateIPv4 string `json:"privateIpv4"`
	// 内网IPV6
	PrivateIPv6 string `json:"privateIpv6"`
	// eth1内网IPV4
	PrivateEth1IPv4 string `json:"privateEth1Ipv4"`
	// eth1内网IPV6
	PrivateEth1IPv6 string `json:"privateEth1Ipv6"`
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
	// gpu数量
	GpuAmount int `json:"gpuAmount"`
	// gpu厂商
	GpuManufacturer string `json:"gpuManufacturer"`
	// gpu处理器型号
	GpuModel string `json:"gpuModel"`
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
	// 是否锁定解锁锁定:locked,解锁unlocked
	Locked string `json:"locked"`
	// 是否锁定解锁锁定:已解锁,已锁定
	LockedName string `json:"lockedName"`
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
	// 引导模式 [UEFI Legacy/BIOS]
	BootMode string `json:"bootMode"`
	//实例的volume和raid信息
	InstanceVolumeRaids []*InstanceVolumeRaid `json:"volumeRaid"`
	// 所属的机型是否需要配置raid，页面有用
	IsNeedRaid string `json:"isNeedRaid"`
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

//func (d Instance) MarshalJSON() ([]byte, error) {
//	type Alias Instance
//	return json.Marshal(struct {
//		Alias
//		CreateTime string `json:"create_time"`
//		UpdateTime string `json:"update_time"`
//	}{
//		Alias: Alias(d),
//		//CreateTime: d.CreateTime.Format("2006-01-02 15:04:05"),
//		//UpdateTime: d.UpdateTime.Format("2006-01-02 15:04:05"),
//	})
//}

//func Dinstance2Instance(d_instance *instanceDao.DInstance) *Instance {
//	return &Instance{
//		//InstanceId:              d_instance.Uuid,
//		//DeviceId:                d_instance.DeviceId,
//		//Region:                  d_instance.Region,
//		//Az:                      d_instance.Az,
//		//DeviceType:              d_instance.DeviceType,
//		//Tenant:                  d_instance.Tenant,
//		//Name:                    d_instance.Name,
//		//Hostname:                d_instance.Hostname,
//		//Description:             d_instance.Description,
//		//Status:                  d_instance.Status,
//		//ImageId:                 d_instance.ImageId,
//		//SystemVolumeRaidId:      d_instance.SystemVolumeRaidId,
//		//DataVolumeRaidId:        d_instance.DataVolumeRaidId,
//		//NetworkType:             d_instance.NetworkType,
//		//SubnetId:                d_instance.SubnetId,
//		//PrivateIp:               d_instance.PrivateIp,
//		//LineType:                d_instance.LineType,
//		//PublicIp:                d_instance.PublicIp,
//		//EnableInternet:          d_instance.EnableInternet,
//		//EnableIpv6:              d_instance.EnableIpv6,
//		//Ipv6Address:             d_instance.Ipv6Address,
//		//EipId:                   d_instance.EipId,
//		//Bandwidth:               d_instance.Bandwidth,
//		//KeypairId:               d_instance.KeypairId,
//		//InterfaceMode:           d_instance.InterfaceMode,
//		//ExtensionSubnetId:       d_instance.ExtensionSubnetId,
//		//ExtensionPrivateIp:      d_instance.ExtensionPrivateIp,
//		//ExtensionEnableInternet: d_instance.ExtensionEnableInternet,
//		//ExtensionEipId:          d_instance.ExtensionEipId,
//		//ExtensionEnableIpv6:     d_instance.ExtensionEnableIpv6,
//		//ExtensionIpv6Address:    d_instance.ExtensionIpv6Address,
//		//Sn:                      d_instance.Sn,
//		//UPosition:               d_instance.UPosition,
//		//Cabinet:                 d_instance.Cabinet,
//		//IloIp:                   d_instance.IloIp,
//		//SaleStatus:              d_instance.SaleStatus,
//		//CreateTime:              d_instance.CreateTime,
//		//UpdateTime:              d_instance.UpdateTime,
//	}
//}
//
//func DaoInstanceDevice2Instance(entity *instanceDao.Instance, device *Device) *Instance {
//	return &Instance{
//
//		//InstanceId: entity.UUID,
//		//DeviceId:   entity.DeviceID,
//		////Region:                  device.Region,
//		////Az:                      device.Az,
//		////DeviceType:              device.DeviceType,
//		//Tenant:                  entity.Tenant,
//		//Name:                    entity.Name,
//		//Hostname:                entity.Hostname,
//		//Description:             entity.Description,
//		//Status:                  entity.Status,
//		//ImageId:                 entity.ImageID,
//		//SystemVolumeRaidId:      entity.SystemVolumeRaidID,
//		//DataVolumeRaidId:        entity.DataVolumeRaidID,
//		//NetworkType:             entity.NetworkType,
//		//SubnetId:                entity.SubnetID,
//		//PrivateIp:               entity.PrivateIP,
//		//LineType:                entity.LineType,
//		//PublicIp:                entity.PublicIP,
//		//EnableInternet:          util.int2Bool(entity.EnableInternet),
//		//EnableIpv6:              util.int2Bool(entity.EnableIPv6),
//		//Ipv6Address:             entity.IPv6Address,
//		//EipId:                   entity.EipID,
//		//Bandwidth:               entity.Bandwidth,
//		//KeypairId:               entity.KeypairID,
//		//InterfaceMode:           entity.InterfaceMode,
//		//ExtensionSubnetId:       entity.ExtensionSubnetID,
//		//ExtensionPrivateIp:      entity.ExtensionPrivateIP,
//		//ExtensionEnableInternet: util.int2Bool(entity.ExtensionEnableInternet),
//		//ExtensionEipId:          entity.ExtensionEipID,
//		//ExtensionEnableIpv6:     util.int2Bool(entity.ExtensionEnableIPv6),
//		//ExtensionIpv6Address:    entity.ExtensionIPv6Address,
//		//Sn:                      device.Sn,
//		//UPosition:               device.UPosition,
//		//Cabinet:                 device.Cabinet,
//		////IloIp:                   device.IloIp,
//		////SaleStatus:              device.SaleStatus,
//		//CreateTime: entity.CreateTime,
//		//UpdateTime: entity.UpdateTime,
//	}
//}

type InstanceForShare struct {
	ID int `json:"-"`
	// 实例uuid
	InstanceID string `json:"instanceId"` // 实例ID（uuid）
	// 实例名称
	InstanceName string `json:"instanceName"`
	//新增，是否已经被共享过
	HasShared bool `json:"hasShared"`
}

type InstanceForShareList struct {
	// instance实体数组
	Instances []*InstanceForShare `json:"instances"`
}
