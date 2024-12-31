package subnetDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// Subnet 子网
type Subnet struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"id"`           // ID
	UUID       string    `gorm:"unique;column:uuid;type:varchar(36);not null" json:"uuid"`                    // 子网ID
	Az         string    `gorm:"column:az;type:varchar(32);not null" json:"az"`                               // az
	Tenant     string    `gorm:"index:i_subnet_tenant;column:tenant;type:varchar(64);not null" json:"tenant"` // 租户
	Name       string    `gorm:"column:name;type:varchar(32)" json:"name"`
	Cidr       string    `gorm:"column:cidr;type:varchar(32);not null" json:"cidr"`                 // 子网网段
	VxlanID    int64     `gorm:"column:vxlan_id;type:bigint(20) unsigned;not null" json:"vxlan_id"` // vxlan_id
	Gateway    string    `gorm:"column:gateway;type:varchar(32);not null" json:"gateway"`           // 子网网关
	Mask       string    `gorm:"column:mask;type:varchar(32);not null" json:"mask"`                 // 子网掩码
	Mac        string    `gorm:"column:mac;type:varchar(32);not null" json:"mac"`                   // MAC
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`      // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`               // 更新时间
	IsDel      int8      `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`              // 是否删除(0-未删, 1-已删)
}

func (t *Subnet) TableName() string {
	return "subnet"
}

// AddSubnet insert a new Subnet into database and returns
// last inserted Id on success.
func AddSubnet(logger *log.Logger, m *Subnet) (id int64, err error) {

	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicRdb, m)
}

// GetSubnetById retrieves Subnet by Id. Returns error if
// Id doesn't exist
func GetSubnetById(logger *log.Logger, id int64) (v *Subnet, err error) {
	v = &Subnet{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetOneSubnet(logger *log.Logger, query map[string]interface{}) (v *Subnet, err error) {
	v = &Subnet{}
	var db = dao.Model(logger, dao.IronicRdb, Subnet{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetSubnetCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Subnet{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return

}

// GetAllRaid retrieves all raid matches certain condition. Returns empty list if
// no records exist
func GetAllSubnet(logger *log.Logger, query map[string]interface{}) (ml []*Subnet, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Subnet{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiSubnet retrieves all Subnet matches certain condition. Returns empty list if
// no records exist
func GetMultiSubnet(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Subnet, err error) {
	var db = dao.Model(logger, dao.IronicRdb, Subnet{})
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

// UpdateSubnet updates Subnet by Id and returns error if
// the record to be updated doesn't exist
func UpdateSubnetById(logger *log.Logger, m *Subnet) (err error) {

	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Subnet{}).Where("id = ?", m.ID).Updates(m).Error
}

func GetBySubnetIdAndTenant(logger *log.Logger, subnet_id, tenant string) (v *Subnet, err error) {
	v = &Subnet{}
	// err = dao.Where(logger, dao.IronicRdb, "uuid = ? and tenant = ? and is_del = 0", subnet_id, tenant).Take(&v).Error
	err = dao.Where(logger, dao.IronicRdb, "uuid = ? and is_del = 0", subnet_id).Take(&v).Error
	if err != nil {
		return nil, err
	}
	return v, nil

}

func QueryByTenant(logger *log.Logger, tenant, az string) (ml []*Subnet, err error) {
	// err = dao.Where(logger, dao.IronicRdb, "az = ? and tenant = ? and is_del = 0", az, tenant).Find(&ml).Error
	err = dao.Where(logger, dao.IronicRdb, "is_del = 0").Find(&ml).Error
	return
}

func CountByParam(logger *log.Logger, az, tenant, name string, subnet_ids []string) (n int64, err error) {
	sql := `select count(1) from subnet t where 1=1`
	// if az != "" {
	// 	sql = sql + fmt.Sprintf(" and t.az = '%s'", az)
	// }
	// if tenant != "" {
	// 	sql = sql + fmt.Sprintf(" and t.tenant = '%s'", tenant)
	// }
	if name != "" {
		sql = sql + fmt.Sprintf(" and t.name LIKE CONCAT('%%', '%s', '%%')", name)
	}
	if len(subnet_ids) > 0 {
		for k, v := range subnet_ids {
			subnet_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t.uuid in (%s)", strings.Join(subnet_ids, ","))
	}
	sql = sql + " and t.is_del = 0"
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&n).Error
	return
}

func QueryByParam(logger *log.Logger, az, tenant, name string, subnet_ids []string, offset, limit int64) (ml []*Subnet, err error) {
	sql := `select * from subnet t where 1=1`
	// if az != "" {
	// 	sql = sql + fmt.Sprintf(" and t.az = '%s'", az)
	// }
	// if tenant != "" {
	// 	sql = sql + fmt.Sprintf(" and t.tenant = '%s'", tenant)
	// }
	if name != "" {
		sql = sql + fmt.Sprintf(" and t.name LIKE CONCAT('%%', '%s', '%%')", name)
	}
	if len(subnet_ids) > 0 {
		for k, v := range subnet_ids {
			subnet_ids[k] = fmt.Sprintf("'%s'", v)
		}
		sql = sql + fmt.Sprintf(" and t.uuid in (%s)", strings.Join(subnet_ids, ","))
	}
	sql = sql + fmt.Sprintf(" and t.is_del = 0 limit %d ,%d", offset, limit)
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}
