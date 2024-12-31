package instanceDao

import (
	"fmt"
	"strings"
	"time"

	log "git.jd.com/cps-golang/log"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
)

type Instance struct {
	ID                 int    `gorm:"primaryKey;column:id" json:"-"`
	InstanceID         string `gorm:"column:instance_id" json:"instanceId"`                   // 实例ID（uuid）
	IDcID              string `gorm:"column:idc_id" json:"idcId"`                             // 机房uuid
	ProjectID          string `gorm:"column:project_id" json:"projectId"`                     // 项目id
	UserID             string `gorm:"column:user_id" json:"userId"`                           // 用户uuid
	InstanceName       string `gorm:"column:instance_name" json:"instanceName"`               // 实例名称
	DeviceID           string `gorm:"column:device_id" json:"deviceId"`                       // 设备uuid
	DeviceTypeID       string `gorm:"column:device_type_id" json:"deviceTypeId"`              // 机型uuid
	Hostname           string `gorm:"column:hostname" json:"hostname"`                        // 主机名
	Status             string `gorm:"column:status" json:"status"`                            // 运行状态
	Reason             string `gorm:"column:reason" json:"reason"`                            // 失败原因
	ImageID            string `gorm:"column:image_id" json:"imageId"`                         // 镜像uuid
	SystemVolumeRaidID string `gorm:"column:system_volume_raid_id" json:"systemVolumeRaidId"` // 系统盘raid
	Locked             string `gorm:"column:locked" json:"locked"`                            // 是否锁定解锁锁定:locked,解锁unlocked
	DataVolumeRaidID   string `gorm:"column:data_volume_raid_id" json:"dataVolumeRaidId"`     // 数据盘raid
	Description        string `gorm:"column:description" json:"description"`                  // 描述
	CreatedBy          string `gorm:"column:created_by" json:"createdBy"`                     // 创建者
	UpdatedBy          string `gorm:"column:updated_by" json:"updatedBy"`                     // 更新者
	CreatedTime        int    `gorm:"column:created_time" json:"createdTime"`                 // 创建时间戳
	UpdatedTime        int    `gorm:"column:updated_time" json:"updatedTime"`                 // 更新时间戳
	DeletedTime        int    `gorm:"column:deleted_time" json:"deletedTime"`                 // 删除时间戳
	IsDel              int8   `gorm:"column:is_del" json:"isDel"`                             // 是否删除(0-未删, 1-已删)
	BootMode           string `gorm:"column:boot_mode" json:"bootMode"`                       // boot类型：UEFI Legacy/BIOS
}

func (t *Instance) TableName() string {
	return "instance"
}

// AddInstance insert a new Instance into database and returns
// last inserted Id on success.
func AddInstance(logger *log.Logger, m *Instance) (id int64, err error) {
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// AddDevice insert a new Device into database and returns
// last inserted Id on success.
func AddMultiInstance(logger *log.Logger, m []*Instance) (id int64, err error) {
	tx := dao.GetGormTx(logger)
	tx.Begin()
	for _, instance := range m {
		if err := tx.Create(instance).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return int64(len(m)), nil
}

// GetInstanceById retrieves Instance by Id. Returns error if
// Id doesn't exist
func GetInstanceByUuid(logger *log.Logger, uuid string) (v *Instance, err error) {
	v = &Instance{}
	err = dao.Where(logger, dao.IronicRdb, "instance_id = ? and is_del = 0", uuid).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetInstanceById retrieves Instance by Id. Returns error if
// Id doesn't exist
func GetInstanceById(logger *log.Logger, id int64) (v *Instance, err error) {
	v = &Instance{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetInstanceCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return n, err
	}
	err = db.Count(&n).Error
	return
}

func GetAllInstance(logger *log.Logger, query map[string]interface{}) (ml []*Instance, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func GetSingleInstance(logger *log.Logger, query map[string]interface{}) (l *Instance, err error) {
	l = &Instance{}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
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

// GetAllInstance retrieves all Instance matches certain condition. Returns empty list if
// no records exist
func GetMultiInstance(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []Instance, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Instance{})
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

// UpdateInstance updates Instance by Id and returns error if
// the record to be updated doesn't exist
func UpdateInstanceById(logger *log.Logger, m *Instance) (err error) {
	return dao.Model(logger, dao.IronicWdb, Instance{}).Where("id = ?", m.ID).Updates(m).Error
}

// description为空也有可能
func UpdateInstanceAnywhere(logger *log.Logger, m *Instance) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Instance{}).Save(m).Error
}

func UpdateSuccessInstanceByInstanceId(logger *log.Logger, m *Instance) (err error) {
	sql := `update instance set reason="` + m.Reason + `" where instance_id="` + m.InstanceID + `" and is_del = 0`
	err = dao.IronicWdb.Exec(sql).Error
	return
}
func UpdateInstanceByInstanceId(logger *log.Logger, m *Instance) (err error) {
	return dao.Model(logger, dao.IronicWdb, Instance{}).Where("instance_id = ? and is_del= ?", m.InstanceID, 0).Updates(m).Error
}
func UpdateMultiInstances(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	updates["update_time"] = time.Now()
	var db = dao.Model(logger, dao.IronicWdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func UpdateKeypairIdAsEmpty(logger *log.Logger, tenant, keypair_id string) (err error) {

	query := map[string]interface{}{
		"tenant":     tenant,
		"keypair_id": keypair_id,
		"is_del":     0,
	}
	var db = dao.Model(logger, dao.IronicWdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}

	updates := map[string]interface{}{
		"keypair_id":  "",
		"update_time": time.Now(),
	}

	err = db.Updates(updates).Error
	return
}

func QueryByTenantAndEipId(logger *log.Logger, tenant, eip_id string) (ml []*Instance, err error) {
	query := map[string]interface{}{
		"tenant": tenant,
		"eip_id": eip_id,
		"is_del": 0,
	}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func QueryByTenantAndExtensionEipId(logger *log.Logger, tenant, eip_id string) (ml []*Instance, err error) {
	query := map[string]interface{}{
		"tenant":           tenant,
		"extension_eip_id": eip_id,
		"is_del":           "0",
	}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func QueryByTenantAndSubnetId(logger *log.Logger, tenant, subnet_id string) (ml []*Instance, err error) {
	query := map[string]interface{}{
		"tenant":    tenant,
		"subnet_id": subnet_id,
		"is_del":    "0",
	}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func QueryByTenantAndExtensionSubnetId(logger *log.Logger, tenant, subnet_id string) (ml []*Instance, err error) {
	query := map[string]interface{}{
		"tenant":              tenant,
		"extension_subnet_id": subnet_id,
		"is_del":              "0",
	}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func QueryByTenantAndIpv6Address(logger *log.Logger, tenant, ipv6_address string) (ml []*Instance, err error) {
	query := map[string]interface{}{
		"tenant":       tenant,
		"ipv6_address": ipv6_address,
		"is_del":       "0",
	}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func QueryByTenantAndExtensionIpv6Address(logger *log.Logger, tenant, ipv6_address string) (ml []*Instance, err error) {
	query := map[string]interface{}{
		"tenant":                 tenant,
		"extension_ipv6_address": ipv6_address,
		"is_del":                 0,
	}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

type DInstanceSaleStatistics struct {
	Tenant     string `json:"tenant"`
	Az         string `json:"az"`
	DeviceType string `json:"device_type"`
	Count      int    `json:"count"`
}

func SaleStatisticsByTenant(logger *log.Logger) (ml []*DInstanceSaleStatistics, err error) {

	sql := `select tenant, az, device_type, count(device_type) as count from instance i left join device d on i.device_id = d.id where i.is_del = 0 group by tenant,az,device_type`
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}

func QueryByInstanceIds(logger *log.Logger, instance_ids []string) (ml []*Instance, err error) {
	query := map[string]interface{}{
		"is_del":  "0",
		"uuid.in": instance_ids,
	}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func QueryByUpdateTimeRange(logger *log.Logger, is_del bool, start_time, end_time time.Time) (ml []*Instance, err error) {
	stime := start_time.Format("2006-01-02 15:04:05")
	etime := end_time.Format("2006-01-02 15:04:05")
	query := map[string]interface{}{
		"update_time.gte": stime,
		"update_time.lt":  etime,
		"is_del":          util.Bool2Int8(is_del),
	}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return

}

func QueryByCreateTimeRange(logger *log.Logger, is_del bool, start_time, end_time time.Time) (ml []*Instance, err error) {

	stime := start_time.Format("2006-01-02 15:04:05")
	etime := end_time.Format("2006-01-02 15:04:05")
	query := map[string]interface{}{
		"create_time.gte": stime,
		"create_time.lt":  etime,
		"is_del":          util.Bool2Int8(is_del),
	}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func CountByTenantAndSubnetId(logger *log.Logger, tenant, subnet_id string) (n int64, err error) {

	query := map[string]interface{}{
		"tenant":    tenant,
		"subnet_id": subnet_id,
		"is_del":    0,
	}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

func CountByParam(logger *log.Logger, query map[string]string) (n int64, err error) {

	sql := `select count(1) from instance t1 left join device t2 on t1.device_id = t2.id where 1=1`
	if query["tenant"] != "" {
		sql = sql + fmt.Sprintf(" and t1.tenant = '%s'", query["tenant"])
	}
	if query["region"] != "" {
		sql = sql + fmt.Sprintf(" and t2.region = '%s'", query["region"])
	}
	if query["az"] != "" {
		sql = sql + fmt.Sprintf(" and t2.az = '%s'", query["az"])
	}

	if query["network_type"] != "" {
		sql = sql + fmt.Sprintf(" and t1.network_type = '%s'", query["network_type"])
	}
	if query["interface_mode"] != "" {
		sql = sql + fmt.Sprintf(" and t1.interface_mode = '%s'", query["interface_mode"])
	}
	if query["subnet_id"] != "" {
		sql = sql + fmt.Sprintf(" and (t1.subnet_id = '%s' or t1.extension_subnet_id = '%s')", query["subnet_id"], query["subnet_id"])
	}
	if query["name"] != "" {
		sql = sql + fmt.Sprintf(" and t1.name LIKE CONCAT('%%', '%s', '%%')", query["name"])
	}
	if query["private_ip"] != "" {
		sql = sql + fmt.Sprintf(" and private_ip LIKE CONCAT('%%', '%s', '%%')", query["private_ip"])
	}
	if query["extension_private_ip"] != "" {
		sql = sql + fmt.Sprintf(" and extension_private_ip LIKE CONCAT('%%', '%s', '%%')", query["extension_private_ip"])
	}
	if query["device_type"] != "" {
		sql = sql + fmt.Sprintf(" and t2.device_type = '%s'", query["device_type"])
	}
	if query["statuses"] != "" {
		statuses := strings.Split(query["statuses"], ",")
		for k, v := range statuses {
			statuses[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.status in (%s)", strings.Join(statuses, ","))
	}
	if query["instance_ids"] != "" {
		instance_ids := strings.Split(query["instance_ids"], ",")
		for k, v := range instance_ids {
			instance_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.uuid in (%s)", strings.Join(instance_ids, ","))
	}
	if query["sns"] != "" {
		sns := strings.Split(query["sns"], ",")
		for k, v := range sns {
			sns[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t2.sn in (%s)", strings.Join(sns, ","))
	}
	if query["eip_ids"] != "" {
		eip_ids := strings.Split(query["eip_ids"], ",")
		for k, v := range eip_ids {
			eip_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and eip_id in (%s)", strings.Join(eip_ids, ","))
	}
	if query["keypair_ids"] != "" {
		keypair_ids := strings.Split(query["keypair_ids"], ",")
		for k, v := range keypair_ids {
			keypair_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and keypair_id in (%s)", strings.Join(keypair_ids, ","))
	}
	if query["enable_internet"] != "" {
		enable_internet := 0
		if query["enable_internet"] == "true" {
			enable_internet = 1
		}
		sql = sql + fmt.Sprintf(" and enable_internet = %d", enable_internet)
	}
	sql = sql + " and t1.is_del = 0"
	err = dao.Raw(logger, dao.IronicRdb, sql).Count(&n).Error
	return
}

func CountInstances(logger *log.Logger, query map[string]string) (n int64, err error) {

	sql := `select count(1) from device t2 left join (select * from instance where is_del=0) t1 on t1.device_id = t2.id where 1=1 `
	if query["tenant"] != "" {
		sql = sql + fmt.Sprintf(" and t1.tenant = '%s'", query["tenant"])
	}
	if query["region"] != "" {
		sql = sql + fmt.Sprintf(" and t2.region = '%s'", query["region"])
	}
	if query["az"] != "" {
		sql = sql + fmt.Sprintf(" and t2.az = '%s'", query["az"])
	}
	if query["device_type"] != "" {
		sql = sql + fmt.Sprintf(" and t2.device_type = '%s'", query["device_type"])
	}
	if query["sale_status"] != "" {
		sql = sql + fmt.Sprintf(" and t2.sale_status = '%s'", query["sale_status"])
	}
	if query["statuses"] != "" {
		statuses := strings.Split(query["statuses"], ",")
		for k, v := range statuses {
			statuses[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.status in (%s)", strings.Join(statuses, ","))
	}
	if query["instance_ids"] != "" {
		instance_ids := strings.Split(query["instance_ids"], ",")
		for k, v := range instance_ids {
			instance_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.uuid in (%s)", strings.Join(instance_ids, ","))
	}
	if query["public_ips"] != "" && query["eip_ids"] != "" {

		public_ips := strings.Split(query["public_ips"], ",")
		for k, v := range public_ips {
			public_ips[k] = fmt.Sprintf("'%s'", v)
		}

		eip_ids := strings.Split(query["eip_ids"], ",")
		for k, v := range eip_ids {
			eip_ids[k] = fmt.Sprintf("'%s'", v)
		}

		sql = sql + fmt.Sprintf(" and (t1.public_ip in (%s) or t1.eip_id in (%s))", strings.Join(public_ips, ","), strings.Join(eip_ids, ","))
	} else if query["public_ips"] != "" {
		public_ips := strings.Split(query["public_ips"], ",")
		for k, v := range public_ips {
			public_ips[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.public_ip in (%s)", strings.Join(public_ips, ","))
	} else if query["eip_ids"] != "" {
		eip_ids := strings.Split(query["eip_ids"], ",")
		for k, v := range eip_ids {
			eip_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.eip_id in (%s)", strings.Join(eip_ids, ","))
	}
	if query["private_ip"] != "" {
		sql = sql + fmt.Sprintf(" and private_ip LIKE CONCAT('%%', '%s', '%%')", query["private_ip"])
	}
	if query["sns"] != "" {
		sns := strings.Split(query["sns"], ",")
		for k, v := range sns {
			sns[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t2.sn in (%s)", strings.Join(sns, ","))
	}
	if query["ilo_ips"] != "" {
		ilo_ips := strings.Split(query["ilo_ips"], ",")
		for k, v := range ilo_ips {
			ilo_ips[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t2.ilo_ip in (%s)", strings.Join(ilo_ips, ","))
	}

	if query["start_time"] != "" && query["end_time"] != "" {
		sql = sql + fmt.Sprintf(" and t1.create_time between '%s' and '%s'", query["start_time"], query["end_time"])
	}
	sql = sql + ` and t2.is_del = 0`
	err = dao.Raw(logger, dao.IronicRdb, sql).Count(&n).Error
	return
}

type DInstance struct {
	Uuid                    string    `json:"uuid"`
	DeviceId                int64     `json:"device_id"`
	Tenant                  string    `json:"tenant"`
	Region                  string    `json:"region"`
	Az                      string    `json:"az"`
	DeviceType              string    `json:"device_type"`
	Name                    string    `json:"name"`
	Hostname                string    `json:"hostname"`
	Description             string    `json:"description"`
	Status                  string    `json:"status"`
	ImageId                 string    `json:"image_id"`
	SystemVolumeRaidId      string    `json:"system_volume_raid_id"`
	DataVolumeRaidId        string    `json:"data_volume_raid_id"`
	NetworkType             string    `json:"network_type"`
	SubnetId                string    `json:"subnet_id"`
	PrivateIp               string    `json:"private_ip"`
	LineType                string    `json:"line_type"`
	PublicIp                string    `json:"public_ip"`
	EnableInternet          bool      `json:"enable_internet"`
	EnableIpv6              bool      `json:"enable_ipv6"`
	Ipv6Address             string    `json:"ipv6_address"`
	EipId                   string    `json:"eip_id"`
	Bandwidth               int       `json:"bandwidth"`
	KeypairId               string    `json:"keypair_id"`
	InterfaceMode           string    `json:"interface_mode"`
	ExtensionSubnetId       string    `json:"extension_subnet_id"`
	ExtensionPrivateIp      string    `json:"extension_private_ip"`
	ExtensionEnableInternet bool      `json:"extension_enable_internet"`
	ExtensionEipId          string    `json:"extension_eip_id"`
	ExtensionEnableIpv6     bool      `json:"extension_enable_ipv6"`
	ExtensionIpv6Address    string    `json:"extension_ipv6_address"`
	Sn                      string    `json:"sn"`
	UPosition               string    `json:"u_position"`
	Cabinet                 string    `json:"cabinet"`
	IloIp                   string    `json:"ilo_ip"`
	SaleStatus              string    `json:"sale_status"`
	SwitchIp                string    `json:"switch_ip"`
	SwitchCidr              string    `json:"switch_cidr"`
	CreateTime              time.Time `json:"create_time"`
	UpdateTime              time.Time `json:"update_time"`
}

func QueryByParam(logger *log.Logger, query map[string]string, offset, limit int64) (ml []*DInstance, err error) {

	sql := `select t1.uuid,t1.device_id,t1.tenant,t1.name,t1.hostname,t1.image_id,t1.system_volume_raid_id,t1.data_volume_raid_id,t1.network_type,t1.subnet_id,t1.private_ip,t1.line_type,t1.public_ip,t1.enable_internet,t1.enable_ipv6,t1.ipv6_address,t1.eip_id,t1.extension_enable_ipv6,t1.extension_ipv6_address,t1.bandwidth,t1.status,t1.description,t1.keypair_id,t1.interface_mode,t1.extension_subnet_id,t1.extension_private_ip,t1.extension_enable_internet,t1.extension_eip_id,t1.create_time,t1.update_time,
	t2.region,t2.az,t2.device_type,t2.sn,t2.u_position,t2.cabinet,t2.ilo_ip,t2.sale_status from instance t1 left join device t2 on t1.device_id = t2.id where 1=1 `
	if query["tenant"] != "" {
		sql = sql + fmt.Sprintf(" and t1.tenant = '%s'", query["tenant"])
	}
	if query["region"] != "" {
		sql = sql + fmt.Sprintf(" and t2.region = '%s'", query["region"])
	}
	if query["az"] != "" {
		sql = sql + fmt.Sprintf(" and t2.az = '%s'", query["az"])
	}

	if query["network_type"] != "" {
		sql = sql + fmt.Sprintf(" and t1.network_type = '%s'", query["network_type"])
	}
	if query["interface_mode"] != "" {
		sql = sql + fmt.Sprintf(" and t1.interface_mode = '%s'", query["interface_mode"])
	}
	if query["subnet_id"] != "" {
		sql = sql + fmt.Sprintf(" and (t1.subnet_id = '%s' or t1.extension_subnet_id = '%s')", query["subnet_id"], query["subnet_id"])
	}
	if query["name"] != "" {
		sql = sql + fmt.Sprintf(" and t1.name LIKE CONCAT('%%', '%s', '%%')", query["name"])
	}
	if query["private_ip"] != "" {
		sql = sql + fmt.Sprintf(" and private_ip LIKE CONCAT('%%', '%s', '%%')", query["private_ip"])
	}
	if query["extension_private_ip"] != "" {
		sql = sql + fmt.Sprintf(" and extension_private_ip LIKE CONCAT('%%', '%s', '%%')", query["extension_private_ip"])
	}
	if query["device_type"] != "" {
		sql = sql + fmt.Sprintf(" and t2.device_type = '%s'", query["device_type"])
	}
	if query["statuses"] != "" {
		statuses := strings.Split(query["statuses"], ",")
		for k, v := range statuses {
			statuses[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.status in (%s)", strings.Join(statuses, ","))
	}
	if query["instance_ids"] != "" {
		instance_ids := strings.Split(query["instance_ids"], ",")
		for k, v := range instance_ids {
			instance_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.uuid in (%s)", strings.Join(instance_ids, ","))
	}
	if query["sns"] != "" {
		sns := strings.Split(query["sns"], ",")
		for k, v := range sns {
			sns[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t2.sn in (%s)", strings.Join(sns, ","))
	}
	if query["eip_ids"] != "" {
		eip_ids := strings.Split(query["eip_ids"], ",")
		for k, v := range eip_ids {
			eip_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and eip_id in (%s)", strings.Join(eip_ids, ","))
	}
	if query["keypair_ids"] != "" {
		keypair_ids := strings.Split(query["keypair_ids"], ",")
		for k, v := range keypair_ids {
			keypair_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and keypair_id in (%s)", strings.Join(keypair_ids, ","))
	}
	if query["enable_internet"] != "" {
		enable_internet := 0
		if query["enable_internet"] == "true" {
			enable_internet = 1
		}
		sql = sql + fmt.Sprintf(" and enable_internet = %d", enable_internet)
	}
	sql = sql + fmt.Sprintf(" and t1.is_del = 0 ORDER BY t1.id DESC LIMIT %d, %d", offset, limit)
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}

func QueryInstances(logger *log.Logger, query map[string]string, offset, limit int64) (ml []*DInstance, err error) {

	sql := `select t1.uuid,t1.device_id,t1.tenant,t1.name,t1.hostname,t1.image_id,t1.system_volume_raid_id,t1.data_volume_raid_id,t1.network_type,t1.subnet_id,t1.private_ip,t1.line_type,t1.public_ip,t1.enable_internet,t1.enable_ipv6,t1.ipv6_address,t1.eip_id,t1.extension_enable_ipv6,t1.extension_ipv6_address,t1.bandwidth,t1.status,t1.description,t1.keypair_id,t1.interface_mode,t1.extension_subnet_id,t1.extension_private_ip,t1.extension_enable_internet,t1.extension_eip_id,t1.create_time,t1.update_time,t2.region,t2.az,t2.device_type,t2.sn,t2.u_position,t2.cabinet,t2.ilo_ip,t2.sale_status from device t2 left join (select * from instance where is_del=0) t1 on t1.device_id = t2.id where 1=1 `
	if query["tenant"] != "" {
		sql = sql + fmt.Sprintf(" and t1.tenant = '%s'", query["tenant"])
	}
	if query["region"] != "" {
		sql = sql + fmt.Sprintf(" and t2.region = '%s'", query["region"])
	}
	if query["az"] != "" {
		sql = sql + fmt.Sprintf(" and t2.az = '%s'", query["az"])
	}
	if query["device_type"] != "" {
		sql = sql + fmt.Sprintf(" and t2.device_type = '%s'", query["device_type"])
	}
	if query["sale_status"] != "" {
		sql = sql + fmt.Sprintf(" and t2.sale_status = '%s'", query["sale_status"])
	}
	if query["statuses"] != "" {
		statuses := strings.Split(query["statuses"], ",")
		for k, v := range statuses {
			statuses[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.status in (%s)", strings.Join(statuses, ","))
	}
	if query["instance_ids"] != "" {
		instance_ids := strings.Split(query["instance_ids"], ",")
		for k, v := range instance_ids {
			instance_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.uuid in (%s)", strings.Join(instance_ids, ","))
	}
	if query["public_ips"] != "" && query["eip_ids"] != "" {

		public_ips := strings.Split(query["public_ips"], ",")
		for k, v := range public_ips {
			public_ips[k] = fmt.Sprintf("'%s'", v)
		}

		eip_ids := strings.Split(query["eip_ids"], ",")
		for k, v := range eip_ids {
			eip_ids[k] = fmt.Sprintf("'%s'", v)
		}

		sql = sql + fmt.Sprintf(" and (t1.public_ip in (%s) or t1.eip_id in (%s))", strings.Join(public_ips, ","), strings.Join(eip_ids, ","))
	} else if query["public_ips"] != "" {
		public_ips := strings.Split(query["public_ips"], ",")
		for k, v := range public_ips {
			public_ips[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.public_ip in (%s)", strings.Join(public_ips, ","))
	} else if query["eip_ids"] != "" {
		eip_ids := strings.Split(query["eip_ids"], ",")
		for k, v := range eip_ids {
			eip_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t1.eip_id in (%s)", strings.Join(eip_ids, ","))
	}
	if query["private_ip"] != "" {
		sql = sql + fmt.Sprintf(" and private_ip LIKE CONCAT('%%', '%s', '%%')", query["private_ip"])
	}
	if query["sns"] != "" {
		sns := strings.Split(query["sns"], ",")
		for k, v := range sns {
			sns[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t2.sn in (%s)", strings.Join(sns, ","))
	}
	if query["ilo_ips"] != "" {
		ilo_ips := strings.Split(query["ilo_ips"], ",")
		for k, v := range ilo_ips {
			ilo_ips[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t2.ilo_ip in (%s)", strings.Join(ilo_ips, ","))
	}

	if query["start_time"] != "" && query["end_time"] != "" {
		sql = sql + fmt.Sprintf(" and t1.create_time between '%s' and '%s'", query["start_time"], query["end_time"])
	}
	sql = sql + fmt.Sprintf(" and t2.is_del = 0 ORDER BY t1.id DESC LIMIT %d, %d", offset, limit)
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}
