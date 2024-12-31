package response

// import "encoding/json"

type Device struct {

	// 额外拼接的字段，省去前端工作
	CpuInfo string `json:"cpuInfo"`
	MemInfo string `json:"memInfo"` // 内存信息拼接
	SvInfo  string `json:"svInfo"`  // 系统盘信息拼接
	DvInfo  string `json:"dvInfo"`  // 数据盘信息拼接
	GpuInfo string `json:"gpuInfo"`
	NicInfo string `json:"nicInfo"` // 单网卡 or 双网口

	IdcID            string `json:"idcId"`
	DeviceTypeID     string `json:"deviceTypeId"`
	ID               int    `json:"id"`               // 设备ID编号
	Sn               string `json:"sn"`               // 设备SN
	DeviceID         string `json:"deviceId"`         // 设备uuid
	InstanceID       string `json:"instanceId"`       // 实例id
	ManageStatus     string `json:"manageStatus"`     // 设备状态: 未装机，已装机
	ManageStatusName string `json:"manageStatusName"` // 设备状态: 未装机，已装机
	Reason           string `json:"reason"`           // 设备状态变更失败原因
	Cabinet          string `json:"cabinet"`          // 机柜编码
	UPosition        string `json:"uPosition"`        // U位
	IloIP            string `json:"iloIp"`            // 带外管理IP
	IloUser          string `json:"iloUser"`          // 带外账号
	IloPassword      string `json:"iloPassword"`      // 带外账号密码

	Mac1        string `json:"mac1"`        // MAC1（eth0）
	Mac2        string `json:"mac2"`        // MAC2（eth2）
	SwitchIP1   string `json:"switchIp1"`   // 交换机1ip
	SwitchPort1 string `json:"switchPort1"` // 交换机1port
	SwitchIP2   string `json:"switchIp2"`   // 交换机2ip
	SwitchPort2 string `json:"switchPort2"` // 交换机2port

	SwitchUser1     string `json:"switchUser1"`     // 交换机1登录账号，如果为空，取所在机房的值
	SwitchPassword1 string `json:"switchPassword1"` // 交换机1登录密码
	SwitchUser2     string `json:"switchUser2"`     // 交换机2登录账号，如果为空，取所在机房的值
	SwitchPassword2 string `json:"switchPassword2"` // 交换机2登录密码
	Description     string `json:"description"`     // 描述
	SwitchIP        string `json:"switchIp"`        // 网口交换机IP
	Mask            string `json:"mask"`            // 子网掩码
	Eth1Mask        string `json:"eth1Mask"`        // eth1子网掩码
	Gateway         string `json:"gateway"`         // 网关地址
	PrivateIPv4     string `json:"privateIpv4"`     // 内网IPV4
	PrivateEth1IPV4 string `json:"privateEth1Ipv4"` // ETH1内网IPV4
	PrivateIPv6     string `json:"privateIpv6"`     // 内网IPV6
	PrivateEth1IPV6 string `json:"privateEth1Ipv6"` // ETH1内网IPV4
	AdapterID       int    `json:"adapterId"`       // adapter_id
	RaidDriver      string `json:"raidDriver"`      // raid工具：（megacli64等）

	//添加到disk表
	Enclosure1 string `json:"enclosure1"` //系统盘1背板号
	Slot1      int    `json:"slot1"`      //系统盘1槽位
	Enclosure2 string `json:"enclosure2"` //系统盘2背板号
	Slot2      int    `json:"slot2"`      //系统盘2槽位
	//以后还能还有数据盘，todo

	CreatedBy   string `json:"createdBy"`   // 创建者
	UpdatedBy   string `json:"updatedBy"`   // 更新者
	CreatedTime string `json:"createdTime"` // 创建时间戳
	UpdatedTime string `json:"updatedTime"` // 更新时间戳
	//DeletedTime     int    `json:"deletedTime"`         // 删除时间戳
	//IsDel           int8   `json:"isDel"`                     // 是否删除0未删除 1已删除

	//以下是通过别的表查询出来的数据，用于返回给上层
	IdcName        string `json:"idcName"`
	IdcNameEn      string `json:"idcNameEn"`
	DeviceTypeName string `json:"deviceTypeName"`

	//实例信息
	InstanceName       string `json:"instanceName"`
	ImageName          string `json:"imageName"`
	InstanceStatus     string `json:"instanceStatus"`
	InstanceStatusName string `json:"instanceStatusName"`
	// 实例失败原因
	InstanceReason      string `json:"instanceReason"`
	UserID              string `json:"userId"`
	UserName            string `json:"userName"`
	InstanceCreatedTime string `json:"instanceCreatedTime"`
	InstanceDescription string `json:"instanceDescription"`
	// 实例是否锁定 锁定locked 未锁定unlocked
	Locked     string `json:"locked"`
	LockedName string `json:"lockedname"`
	//硬件配置
	Brand                     string `json:"brand"`            //品牌
	Model                     string `json:"model"`            //型号
	DeviceType                string `json:"deviceType"`       //型号，机型规格cps.normal
	DeviceSeries              string `json:"deviceSeries"`     //计算型，存储型
	DeviceSeriesName          string `json:"deviceSeriesName"` //计算型，存储型
	Architecture              string `json:"architecture"`     // 体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64
	CPUAmount                 int64  `json:"cpuAmount"`        // cpu数量
	CPUCores                  int64  `json:"cpuCores"`         // 单个cpu内核数
	CPURoads                  int8   `json:"cpuRoads"`         // 单个cpu路数
	CPUManufacturer           string `json:"cpuManufacturer"`  // cpu厂商
	CPUModel                  string `json:"cpuModel"`         // cpu处理器型号
	CPUFrequency              string `json:"cpuFrequency"`     // cpu频率(G)
	MemType                   string `json:"memType"`          // 内存接口（如DDR3，DDR4）
	MemSize                   int    `json:"memSize"`          // 单个内存大小(GB)
	MemAmount                 int    `json:"memAmount"`        // 内存数量
	MemFrequency              int    `json:"memFrequency"`     // 内存主频（MHz)
	NicAmount                 int64  `json:"nicAmount"`        // 网卡数量
	NicRate                   int    `json:"nicRate"`          // 网卡传输速率(GE)
	InterfaceMode             string `json:"interfaceMode"`    // 【网口模式】【网络设置】: bond单网口,dual双网口
	InterfaceModeName         string `json:"interfaceModeName"`
	SystemVolumeType          string `json:"systemVolumeType"`          // 系统盘类型（SSD，HDD）
	SystemVolumeInterfaceType string `json:"systemVolumeInterfaceType"` // 系统盘接口类型（SATA,SAS,NVME）
	SystemVolumeSize          int    `json:"systemVolumeSize"`          // 系统盘单盘大小
	//SystemVolumeUnit          string `json:"systemVolumeUnit"`          // 系统盘单位（GB，TB）
	SystemVolumeAmount      int64  `json:"systemVolumeAmount"`      // 系统盘数量
	GpuAmount               int64  `json:"gpuAmount"`               // gpu数量
	GpuManufacturer         string `json:"gpuManufacturer"`         // gpu厂商
	GpuModel                string `json:"gpuModel"`                // gpu处理器型号
	DataVolumeType          string `json:"dataVolumeType"`          // 数据盘类型（SSD，HDD）
	DataVolumeInterfaceType string `json:"dataVolumeInterfaceType"` // 数据盘接口类型（SATA,SAS,NVME）
	DataVolumeSize          int    `json:"dataVolumeSize"`          // 数据盘单盘大小
	//DataVolumeUnit            string `json:"dataVolumeUnit"`            // 数据盘单位（GB，TB）
	DataVolumeAmount int64 `json:"dataVolumeAmount"` // 数据盘数量
	// "1"表示已采集，"2"表示未采集,3表示采集中，4表示采集失败
	CollectStatus     string `json:"collectStatus"`
	CollectFailReason string `json:"collectFailReason"`
	IsNeedRaid        string `json:"isNeedRaid"`
}

// 上传设备返回
type UploadDevice struct {

	// 额外拼接的字段，省去前端工作
	Sn        string `json:"sn"`        // 设备SN
	Cabinet   string `json:"cabinet"`   // 机柜编码
	UPosition string `json:"uPosition"` // U位
	// Brand       string `json:"brand"`       //品牌
	// Model       string `json:"model"`       //型号
	IloIP       string `json:"iloIp"`       // 带外管理IP
	IloUser     string `json:"iloUser"`     // 带外账号
	IloPassword string `json:"iloPassword"` // 带外账号密码

	Mac1 string `gorm:"column:mac1" json:"mac1"` // MAC1（eth0）
	// Mac2 string `gorm:"column:mac2" json:"mac2"` // MAC2（eth2）
	/*
		SwitchIP1   string `gorm:"column:switch_ip1" json:"switchIp1"`     // 交换机1ip
		SwitchPort1 string `gorm:"column:switch_port1" json:"switchPort1"` // 交换机1port
		SwitchIP2   string `gorm:"column:switch_ip2" json:"switchIp2"`     // 交换机2ip
		SwitchPort2 string `gorm:"column:switch_port2" json:"switchPort2"` // 交换机2port

		SwitchUser1     string `json:"switchUser1"`     // 交换机1登录账号，如果为空，取所在机房的值
		SwitchPassword1 string `json:"switchPassword1"` // 交换机1登录密码
		SwitchUser2     string `son:"switchUser2"`      // 交换机2登录账号，如果为空，取所在机房的值
		SwitchPassword2 string `json:"switchPassword2"` // 交换机2登录密码
	*/
	//SwitchIP        string `json:"switchIp"`        // 网口交换机IP
	Mask string `json:"mask"` // 子网掩码
	// MaskEth1    string `json:"maskEth1"`    // eth1子网掩码
	Gateway     string `json:"gateway"`     // 网关地址
	PrivateIPv4 string `json:"privateIpv4"` // 内网IPV4
	// PrivateEth1IPv4 string `json:"privateEth1Ipv4"` // Eth1内网IPV4
	PrivateIPv6 string `json:"privateIpv6"` // 内网IPV6
	// PrivateEth1IPv6 string `json:"privateEth1Ipv6"` // 内网eth1IPV6
	// AdapterID       int    `json:"adapterId"`       // adapter_id
	//RaidDriver      string `json:"raidDriver"`      // raid工具：（megacli64等）
	/*
		Enclosure1 string `json:"enclosure1"` //系统盘1背板号
		Slot1      int    `json:"slot1"`      //系统盘1槽位
		Enclosure2 string `json:"enclosure2"` //系统盘2背板号
		Slot2      int    `json:"slot2"`      //系统盘2槽位
	*/

}
type DevicePage struct {
	Devices []Device `json:"devices"`
	PagingResponse
}
type DeviceList struct {
	Devices    []*Device `json:"devices"`
	PageNumber int64     `json:"pageNumber"`
	PageSize   int64     `json:"pageSize"`
	TotalCount int64     `json:"totalCount"`
}
type DeviceInfo struct {
	Device *Device `json:"device"`
}
type DeviceIds struct {
	DeviceIds []string `json:"deviceIds"`
}
type DeviceStock struct {
	Stocks int `json:"stocks"`
}
