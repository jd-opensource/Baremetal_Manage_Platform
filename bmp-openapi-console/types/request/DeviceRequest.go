package request

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type CreateDevicesRequest struct {
	// 机房uuid
	// required: true
	IDcID string `json:"idcId" validate:"required"` // 机房uuid
	// 机型uuid
	// required: true
	DeviceTypeID string `json:"deviceTypeId"  validate:"required"` // 设备类型uuid
	// 设备信息
	// required: true
	Devices []*CreateDeviceSpec `json:"devices" validate:"required"`
}

func (req *CreateDevicesRequest) Validate(logger *log.Logger) {
	//validate(req, logger)
	//if err := validator.New().Struct(c); err != nil {
	//	logger.Warn("CreateDevicesRequest.Validate error:", err.Error())
	//	panic(constant.INVALID_ARGUMENT)
	//}
	for _, v := range req.Devices { //符合属性需要单独调用
		v.Validate(logger)
	}
}

type CreateDeviceSpec struct {
	InstanceID string `json:"instanceId"` // 实例id
	// 设备SN
	// required: true
	Sn string `json:"sn" validate:"required,min=1,max=128"` // 设备SN
	// 设备状态，默认已入库
	ManageStatus string `json:"manageStatus"` // 设备状态: 未装机，已装机
	// 品牌
	// required: true
	Brand string `json:"brand" validate:"required,min=1,max=128"` //品牌
	// 型号
	// required: true
	Model string `json:"model" validate:"required,min=1,max=128"` //型号
	// 机柜编码
	// required: true
	Cabinet string `json:"cabinet" validate:"required,min=1,max=128"` // 机柜编码
	// 所在U位
	// required: true
	UPosition string `json:"uPosition" validate:"required,min=1,max=128"` // U位
	// 带外IP
	// required: true
	IloIP string `json:"iloIp" validate:"required,min=1,max=128"` // 带外管理IP
	// 带外账号
	IloUser string `json:"iloUser"` // 带外账号
	// 带外密码
	IloPassword string `json:"iloPassword"` // 带外账号密码
	// MAC1（eth0）
	// required: true
	Mac1 string `json:"mac1" validate:"required,min=1,max=128"` // MAC1（eth0）
	// MAC2（eth1）
	// required: true
	Mac2 string `json:"mac2" validate:"required,min=1,max=128"` // MAC2（eth1）
	// 交换机1ip
	// required: true
	SwitchIP1 string `json:"switchIp1" validate:"required,min=1,max=128"` // 交换机1ip
	// 交换机1port
	// required: true
	SwitchPort1 string `json:"switchPort1" validate:"required,min=1,max=128"` // 交换机1port
	// 交换机2ip
	// required: true
	SwitchIP2 string `json:"switchIp2" validate:"required,min=1,max=128"` // 交换机2ip
	// 交换机2port
	// required: true
	SwitchPort2 string `json:"switchPort2" validate:"required,min=1,max=128"` // 交换机2port
	// 交换机1登录账号
	SwitchUser1 string `json:"switchUser1"` // 交换机1登录账号，如果为空，取所在机房的值
	// 交换机1登录密码
	SwitchPassword1 string `json:"switchPassword1"` // 交换机1登录密码
	// 交换机2登录账号
	SwitchUser2 string `son:"switchUser2"` // 交换机2登录账号，如果为空，取所在机房的值
	// 交换机2登录密码
	SwitchPassword2 string `json:"switchPassword2"` // 交换机2登录密码
	// 描述
	Description string `json:"description"` // 描述
	// 网口交换机IP
	SwitchIP string `json:"switchIp"` // 网口交换机IP
	// 子网掩码
	// required: true
	Mask string `json:"mask" validate:"required,min=1,max=128"` // 子网掩码
	// 网关地址
	// required: true
	Gateway string `json:"gateway" validate:"required,min=1,max=128"` // 网关地址
	// 内网IPV4
	// required: true
	PrivateIPv4 string `json:"privateIpv4" validate:"required,min=1,max=128"` // 内网IPV4
	// 内网IPV6
	PrivateIPv6 string `json:"privateIpv6"` // 内网IPV6
	// adapter_id
	// required: true
	AdapterID *int `json:"adapterId" validate:"required,min=0,max=10000"` // adapter_id
	// raid工具：（megacli128等）
	RaidDriver string `json:"raidDriver" validate:"max=64"` // raid工具：（megacli128等）
	// 系统盘1背板号
	// required: true
	Enclosure1 string `json:"enclosure1" validate:"required,min=0,max=10000"` //系统盘1背板号
	// 系统盘1槽位
	// required: true
	Slot1 *int `json:"slot1" validate:"required,min=0,max=10000"` //系统盘1槽位
	// 系统盘2背板号
	Enclosure2 string `json:"Enclosure2" validate:"max=64"` //系统盘2背板号
	// 系统盘2槽位
	Slot2 int `json:"slot2" validate:"max=10000"` //系统盘2槽位
	//以后还能还有数据盘，todo
}

func (req *CreateDeviceSpec) Validate(logger *log.Logger) {
	validate(req, logger)
}

type ModifyAllDevicesRequest struct {
	InstanceID   string `json:"instanceId"` // 实例id
	ManageStatus string `json:"manageStatus"`
	UserId       string `json:"userId"`
	UserName     string `json:"userName"`
}
type ModifyDevicesRequest struct {
	// 修改设备描述
	// Extensions:
	// x-nullable: true
	Description *string `json:"description" validate:"omitempty,max=256"`
}

func (req *ModifyDevicesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type MountDevicesRequest struct {
	// 设备uuid
	// required: true
	DeviceID string `json:"deviceId" validate:"required"`
}

func (req *MountDevicesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type UnMountDevicesRequest struct {
	// 设备uuid
	// required: true
	DeviceID string `json:"deviceId" validate:"required"`
}

func (req *UnMountDevicesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type QueryDevicesRequest struct {
	Pageable
	// 机房id
	IDcID string `json:"idcId"`
	// 设备sn
	Sn string `json:"sn"`
	// 实例id
	InstanceID string `json:"instanceId"`
	// 机型id
	DeviceTypeID string `json:"deviceTypeId"`
	// 设备管理状态
	ManageStatus string `json:"manageStatus"`
	// 带外ip
	IloIP string `json:"iloIp"`
	// ipv4
	IPV4 string `json:"ipv4"`
	// ipv6
	IPV6 string `json:"ipv6"`
	// 机型类型
	DeviceSeries string `json:"deviceSeries"`
	// 所属用户id
	UserID string `json:"userId"`
	// 所属用户
	UserName string `json:"userName"`
	// 是否显示全部，1不分页
	IsAll string `json:"isAll"`
}

func (req *QueryDevicesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
