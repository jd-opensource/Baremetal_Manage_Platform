package dao

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type Device struct {
	ID              int    `gorm:"primaryKey;column:id" json:"id"`                  // 设备ID编号
	Sn              string `gorm:"column:sn" json:"sn"`                             // 设备SN
	DeviceID        string `gorm:"column:device_id" json:"deviceId"`                // 设备uuid
	InstanceID      string `gorm:"column:instance_id" json:"instanceId"`            // 实例id
	IDcID           string `gorm:"column:idc_id" json:"idcId"`                      // 机型uuid
	DeviceTypeID    string `gorm:"column:device_type_id" json:"deviceTypeId"`       // 设备类型uuid
	ManageStatus    string `gorm:"column:manage_status" json:"manageStatus"`        // 设备状态: 未装机，已装机
	UserId          string `gorm:"column:user_id" json:"userId"`                    // 设备状态: 未装机，已装机
	UserName        string `gorm:"column:user_name" json:"userName"`                // 设备状态: 未装机，已装机
	Brand           string `gorm:"column:brand" json:"brand"`                       // 品牌
	Model           string `gorm:"column:model" json:"model"`                       // 型号
	DeviceSeries    string `gorm:"column:device_series" json:"deviceSeries"`        // 机型类型，如计算型，存储型
	Reason          string `gorm:"column:reason" json:"reason"`                     // 设备状态变更失败原因
	Cabinet         string `gorm:"column:cabinet" json:"cabinet"`                   // 机柜编码
	UPosition       string `gorm:"column:u_position" json:"uPosition"`              // U位
	IloIP           string `gorm:"column:ilo_ip" json:"iloIp"`                      // 带外管理IP
	IloUser         string `gorm:"column:ilo_user" json:"iloUser"`                  // 带外账号
	IloPassword     string `gorm:"column:ilo_password" json:"iloPassword"`          // 带外账号密码
	Mac1            string `gorm:"column:mac1" json:"mac1"`                         // MAC1（eth0）
	Mac2            string `gorm:"column:mac2" json:"mac2"`                         // MAC2（eth2）
	SwitchIP1       string `gorm:"column:switch_ip1" json:"switchIp1"`              // 交换机1ip
	SwitchPort1     string `gorm:"column:switch_port1" json:"switchPort1"`          // 交换机1port
	SwitchIP2       string `gorm:"column:switch_ip2" json:"switchIp2"`              // 交换机2ip
	SwitchPort2     string `gorm:"column:switch_port2" json:"switchPort2"`          // 交换机2port
	SwitchUser1     string `gorm:"column:switch_user1" json:"switchUser1"`          // 交换机1登录账号，如果为空，取所在机房的值
	SwitchPassword1 string `gorm:"column:switch_password1" json:"switchPassword1"`  // 交换机1登录密码
	SwitchUser2     string `gorm:"column:switch_user2" json:"switchUser2"`          // 交换机2登录账号，如果为空，取所在机房的值
	SwitchPassword2 string `gorm:"column:switch_password2" json:"switchPassword2"`  // 交换机2登录密码
	Description     string `gorm:"column:description" json:"description"`           // 描述
	SwitchIP        string `gorm:"column:switch_ip" json:"switchIp"`                // 网口交换机IP
	Mask            string `gorm:"column:mask" json:"mask"`                         // 子网掩码
	MaskEth1        string `gorm:"column:mask_eth1" json:"mask_eth1"`               // eth1子网掩码
	Gateway         string `gorm:"column:gateway" json:"gateway"`                   // 网关地址
	PrivateIPv4     string `gorm:"column:private_ipv4" json:"privateIpv4"`          // 内网IPV4
	PrivateEth1IPv4 string `gorm:"column:private_eth1_ipv4" json:"privateEth1Ipv4"` // eth1内网IPV4
	PrivateIPv6     string `gorm:"column:private_ipv6" json:"privateIpv6"`          // 内网IPV6
	PrivateEth1IPv6 string `gorm:"column:private_eth1_ipv6" json:"privateEth1Ipv6"` // eth1内网IPV6
	Gateway6        string `gorm:"column:gateway6" json:"gateway6"`                 // IPV6网关地址
	AdapterID       int    `gorm:"column:adapter_id" json:"adapterId"`              // adapter_id
	RaidDriver      string `gorm:"column:raid_driver" json:"raidDriver"`            // raid工具：（megacli64等）
	CreatedBy       string `gorm:"column:created_by" json:"createdBy"`              // 创建者
	UpdatedBy       string `gorm:"column:updated_by" json:"updatedBy"`              // 更新者
	CreatedTime     int    `gorm:"column:created_time" json:"createdTime"`          // 创建时间戳
	UpdatedTime     int    `gorm:"column:updated_time" json:"updatedTime"`          // 更新时间戳
	DeletedTime     int    `gorm:"column:deleted_time" json:"deletedTime"`          // 删除时间戳
	IsDel           int8   `gorm:"column:is_del" json:"isDel"`                      // 是否删除0未删除 1已删除
}

func (t *Device) TableName() string {
	return "device"
}

func GetCpssCount(logger *log.Logger, sn, idc string) (n int64, err error) {

	query := map[string]interface{}{
		"is_del": 0,
	}
	if sn != "" {
		query["sn"] = sn
	}
	if idc != "" {
		query["idc_id"] = idc
	}
	var db = Model(logger, IronicRdb, Device{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return

}

func GetCpssByCond(logger *log.Logger, sn, idc string, limit, offset int) (ml []*Device, err error) {
	var db = Model(logger, IronicRdb, Device{})
	query := map[string]interface{}{
		"is_del": 0,
	}
	if sn != "" {
		query["sn"] = sn
	}
	if idc != "" {
		query["idc_id"] = idc
	}
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}

	if offset == 0 && limit == 0 {
		err = db.Find(&ml).Error
		return
	}
	err = db.Order("id desc").Offset(offset).Limit(limit).Find(&ml).Error
	return

}

func GetCpssByCondNoLimit(logger *log.Logger, sn, idc string) (ml []*Device, err error) {
	var db = Model(logger, IronicRdb, Device{})
	query := map[string]interface{}{
		"is_del": 0,
	}
	if sn != "" {
		query["sn"] = sn
	}
	if idc != "" {
		query["idc_id"] = idc
	}
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}

	err = db.Order("id desc").Find(&ml).Error
	return

}

func GetCpssBySns(logger *log.Logger, sns []string) (ml []*Device, err error) {
	var db = Model(logger, IronicRdb, Device{})
	query := map[string]interface{}{
		"is_del": 0,
		"sn.in":  sns,
	}
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}

	err = db.Order("id desc").Find(&ml).Error
	return

}

func GetCpsBySn(logger *log.Logger, sn string) (v *Device, err error) {
	v = &Device{}
	var db = Model(logger, IronicRdb, Device{})
	query := map[string]interface{}{
		"is_del": 0,
		"sn":     sn,
	}
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}

	err = db.Order("id desc").Take(v).Error
	return

}
