package roleDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// Role Role
type Role struct {
	ID            int    `gorm:"primaryKey;column:id" json:"-"`
	RoleID        string `gorm:"column:role_id" json:"roleId"`               // 角色uuid
	RoleNameEn    string `gorm:"column:role_name_en" json:"roleNameEn"`      // 角色名称，唯一
	RoleNameCn    string `gorm:"column:role_name_cn" json:"roleNameCn"`      // 角色名称，唯一
	DescriptionEn string `gorm:"column:description_en" json:"descriptionEn"` // 权限描述
	DescriptionCn string `gorm:"column:description_cn" json:"descriptionCn"` // 权限描述
	CreatedBy     string `gorm:"column:created_by" json:"createdBy"`         // 创建者
	UpdatedBy     string `gorm:"column:updated_by" json:"updatedBy"`         // 更新者
	CreatedTime   int    `gorm:"column:created_time" json:"createdTime"`     // 创建时间戳
	UpdatedTime   int    `gorm:"column:updated_time" json:"updatedTime"`     // 更新时间戳
	DeletedTime   int    `gorm:"column:deleted_time" json:"deletedTime"`     // 删除时间戳
	IsDel         int8   `gorm:"column:is_del" json:"isDel"`                 // 是否删除0未删除 1已删除
}

func (t *Role) TableName() string {
	return "role"
}

// AddRole insert a new Role into database and returns
// last inserted Id on success.
func AddRole(logger *log.Logger, m *Role) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetRoleById retrieves Role by Id. Returns error if
// Id doesn't exist
func GetRoleById(logger *log.Logger, RoleId string) (v *Role, err error) {
	v = &Role{}
	err = dao.Where(logger, dao.IronicRdb, "role_id = ? and is_del = 0", RoleId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetRoleByUuid retrieves Role by Uuid. Returns error if
// Id doesn't exist
func GetRoleCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, Role{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

// GetAllRole retrieves all Role matches certain condition. Returns empty list if
// no records exist
func GetAllRole(logger *log.Logger, query map[string]interface{}) (ml []*Role, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Role{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiRole retrieves all Role matches certain condition. Returns empty list if
// no records exist
func GetMultiRole(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Role, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Role{})
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
	if offset == 0 && limit == 0 {
		err = db.Find(&ml).Error
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

// UpdateRole updates Role by Id and returns error if
// the record to be updated doesn't exist
func UpdateRoleById(logger *log.Logger, m *Role) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Role{}).Where("role_id = ?", m.RoleID).Updates(m).Error
}
