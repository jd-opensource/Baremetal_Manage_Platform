package deviceDao

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	log "git.jd.com/cps-golang/log"
)

type Device struct {
	ID              int    `gorm:"primaryKey;column:id" json:"-"`                   // 设备ID编号
	Sn              string `gorm:"column:sn" json:"sn"`                             // 设备SN
	DeviceID        string `gorm:"column:device_id" json:"deviceId"`                // 设备uuid
	InstanceID      string `gorm:"column:instance_id" json:"instanceId"`            // 实例id
	IDcID           string `gorm:"column:idc_id" json:"idcId"`                      // 机型uuid
	DeviceTypeID    string `gorm:"column:device_type_id" json:"deviceTypeId"`       // 设备类型uuid
	ManageStatus    string `gorm:"column:manage_status" json:"manageStatus"`        // 设备状态: 未装机，已装机
	Brand           string `gorm:"column:brand" json:"brand"`                       // 品牌
	Model           string `gorm:"column:model" json:"model"`                       // 型号
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
	Gateway6        string `gorm:"column:gateway6" json:"gateway6"`                 // 内网IPV6网关
	AdapterID       int    `gorm:"column:adapter_id" json:"adapterId"`              // adapter_id
	RaidDriver      string `gorm:"column:raid_driver" json:"raidDriver"`            // raid工具：（megacli64等）
	CreatedBy       string `gorm:"column:created_by" json:"createdBy"`              // 创建者
	UpdatedBy       string `gorm:"column:updated_by" json:"updatedBy"`              // 更新者
	CreatedTime     int    `gorm:"column:created_time" json:"createdTime"`          // 创建时间戳
	UpdatedTime     int    `gorm:"column:updated_time" json:"updatedTime"`          // 更新时间戳
	DeletedTime     int    `gorm:"column:deleted_time" json:"deletedTime"`          // 删除时间戳
	IsDel           int8   `gorm:"column:is_del" json:"isDel"`                      // 是否删除0未删除 1已删除
	CurrentBootMode string `gorm:"column:current_boot_mode" json:"currentBootMode"` // 当前启动方式
	// "1"表示已采集，"2"表示未采集. "3"表示采集中,4表示采集失败
	CollectStatus string `gorm:"column:collect_status" json:"collectStatus"`
	//采集失败原因
	CollectFailReason string `gorm:"column:collect_fail_reason" json:"collectFailReason"`
	Architecture      string `gorm:"column:architecture" json:"architecture"` // 体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64

}

func (t *Device) TableName() string {
	return "device"
}

func GetBySn(logger *log.Logger, sn string) (v *Device, err error) {
	v = &Device{}
	err = dao.Where(logger, dao.IronicRdb, "sn = ? and is_del = 0", sn).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetOneDevice(logger *log.Logger, query map[string]interface{}) (l *Device, err error) {
	l = &Device{}
	var db = dao.Model(logger, dao.IronicRdb, Device{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Take(l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}

func GetDeviceCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Device{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return n, err
	}
	err = db.Count(&n).Error
	return
}

// AddDevice insert a new Device into database and returns
// last inserted Id on success.
func AddDevice(logger *log.Logger, m *Device) (id int64, err error) {

	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// AddDevice insert a new Device into database and returns
// last inserted Id on success.
func AddMultiDevice(logger *log.Logger, m []*Device) (id int64, err error) {

	tx := dao.GetGormTx(logger)
	tx.Begin()
	for _, device := range m {

		if err := tx.Create(device).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return int64(len(m)), nil
}

// GetDeviceById retrieves Device by Id. Returns error if
// Id doesn't exist
func GetDeviceById(logger *log.Logger, deviceId string) (v *Device, err error) {
	v = &Device{}
	err = dao.Where(logger, dao.IronicRdb, "device_id = ? and is_del = 0", deviceId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetDeviceBySn retrieves Device by Sn. Returns error if
// Sn doesn't exist
func GetDeviceBySn(logger *log.Logger, sn string) (v *Device, err error) {
	v = &Device{}
	err = dao.Where(logger, dao.IronicRdb, "sn = ? and is_del = 0", sn).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetMultiDevice retrieves all Device matches certain condition. Returns empty list if
// no records exist
func GetMultiDevice(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Device, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Device{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	if len(fields) > 0 {
		db = db.Select(fields)
	}
	orderConditions := []string{}
	for i := 0; i < len(sortby); i++ {
		orderConditions = append(orderConditions, fmt.Sprintf("%s %s", sortby[i], order[i]))
	}
	db = db.Order(strings.Join(orderConditions, ","))
	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

// GetMultiDevice retrieves all Device matches certain condition. Returns empty list if
// no records exist
func GetAllDevice(logger *log.Logger, query map[string]interface{}) (ml []*Device, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Device{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func GetSingleDevice(logger *log.Logger, query map[string]interface{}) (l *Device, err error) {
	l = &Device{}
	var db = dao.Model(logger, dao.IronicRdb, Device{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Take(l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}

func UpdateMultiDevices(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	updates["update_time"] = time.Now()
	var db = dao.Model(logger, dao.IronicWdb, Device{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

// UpdateDevice updates Device by Id and returns error if
// the record to be updated doesn't exist
func UpdateDeviceById(logger *log.Logger, m *Device) (err error) {

	return dao.Model(logger, dao.IronicWdb, Device{}).Where("id = ?", m.ID).Updates(m).Error
}
func UpdateDeviceBySn(logger *log.Logger, m *Device) (err error) {
	sql := `update device set manage_status="` + m.ManageStatus + `",reason="` + m.Reason + `" ,updated_time = ` + strconv.Itoa(int(time.Now().Unix())) + ` where sn="` + m.Sn + `" and is_del = 0`
	err = dao.IronicWdb.Exec(sql).Error
	fmt.Println("UpdateDeviceBySn err", err)
	return
	//return dao.Model(logger, dao.IronicWdb, Device{}).Where("sn=? and is_del=?", m.Sn, "0").Save(m).Error
}
func UpdateDeviceByInstanceId(logger *log.Logger, m *Device) (err error) {
	sql := `update device set instance_id='',user_id = '',user_name = '',manage_status="` + m.ManageStatus + `",reason='' ,updated_time = ` + strconv.Itoa(int(time.Now().Unix())) + ` where instance_id="` + m.InstanceID + `" and is_del = 0`
	err = dao.IronicWdb.Exec(sql).Error
	// err = dao.Raw(logger, dao.IronicRdb, sql).Error
	return
	//return dao.Model(logger, dao.IronicWdb, Device{}).Where("sn=? and is_del=?", m.Sn, "0").Save(m).Error
}
func GetFirstVersionDeviceSns(logger *log.Logger) (sns []string, err error) {
	sql := `select distinct d.sn from device d left join command c on d.sn = c.sn where d.is_del=0 and d.sale_status = 'sold' and d.region = 'cn-north-1' and d.az != 'cn-north-1c'`
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&sns).Error
	return
}

func GetFirstVersionAndReInstallDeviceSns(logger *log.Logger) (sns []string, err error) {

	sql := `select distinct d.sn from device d left join command c on d.sn = c.sn where d.is_del = 0 and d.sale_status = 'sold' and d.region = 'cn-north-1' and d.az != 'cn-north-1c' and c.action = 'WriteImage' and c.status = 'finish' and c.is_del = 0`
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&sns).Error
	return
}

type DeviceStock struct {
	Total      int64  `json:"total"`
	Used       int64  `json:"used"`
	Available  int64  `json:"available"`
	Region     string `json:"region"`
	Az         string `json:"az"`
	DeviceType string `json:"device_type"`
}

func GetDeviceStock(logger *log.Logger, device_type, region, az string) (ml []*DeviceStock, err error) {

	sql := `select region, az, device_type, count(id) as total, sum(if(sale_status = 'sold', 1, 0)) as used, sum(if(sale_status = 'selling', 1, 0)) as available from device where 1=1`
	if region != "" {
		sql = sql + fmt.Sprintf(" and region = '%s'", region)
	}
	if az != "" {
		sql = sql + fmt.Sprintf(" and az = '%s'", az)
	}
	if device_type != "" {
		sql = sql + fmt.Sprintf(" and device_type = '%s'", device_type)
	}
	sql = sql + " and is_del = 0 group by region,az,device_type"
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}

type DDeviceIncludeSwitchIp struct {
	*Device
	SwitchIp string `json:"switch_ip"`
}

func QueryIncludeSwitchIpByParam(logger *log.Logger, az, device_type, sale_status string) (ml []*DDeviceIncludeSwitchIp, err error) {

	sql := `select distinct switch_ip, d.* from device d left join interface i on d.sn = i.sn where 1=1`
	if device_type != "" {
		sql = sql + fmt.Sprintf(" and d.device_type = '%s'", device_type)
	}
	if az != "" {
		sql = sql + fmt.Sprintf(" and d.az = '%s'", az)
	}
	if sale_status != "" {
		sql = sql + fmt.Sprintf(" and d.sale_status = '%s'", sale_status)
	}
	sql = sql + ` and d.is_del = 0and i.is_del = 0`
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return

}

// description为空也有可能
func UpdateDeviceAnywhere(logger *log.Logger, m *Device) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Device{}).Save(m).Error
}
