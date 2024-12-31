package sharingProjectDao

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// SharingProject 共享项目
type SharingProject struct {
	ID                int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	ProjectID         string `gorm:"column:project_id;type:varchar(255);not null" json:"project_id"`             // 项目uuid
	ProjectName       string `gorm:"column:project_name;type:varchar(255);not null" json:"project_name"`         // 项目名称
	OwnerUserID       string `gorm:"column:owner_user_id;type:varchar(255);not null" json:"owner_user_id"`       // 项目拥有者用户id
	SharedUserID      string `gorm:"column:shared_user_id;type:varchar(255);not null" json:"shared_user_id"`     // 项目共享者用户id
	OwnerUserName     string `gorm:"column:owner_user_name;type:varchar(255);not null" json:"owner_user_name"`   // 项目拥有者用户名
	SharedUserName    string `gorm:"column:shared_user_name;type:varchar(255);not null" json:"shared_user_name"` // 项目拥有者用户名
	IsDefault         int8   `gorm:"column:is_default;type:tinyint(4);not null" json:"is_default"`               // 是否默认项目0否 1是
	CreatedBy         string `gorm:"column:created_by;type:varchar(255);not null" json:"created_by"`             // 创建者
	UpdatedBy         string `gorm:"column:updated_by;type:varchar(255);not null" json:"updated_by"`             // 更新者
	CreatedTime       int    `gorm:"column:created_time;type:int(11);not null" json:"created_time"`              // 创建时间戳
	UpdatedTime       int    `gorm:"column:updated_time;type:int(11);not null" json:"updated_time"`              // 更新时间戳
	DeletedTime       int    `gorm:"column:deleted_time;type:int(11);not null" json:"deleted_time"`              // 删除时间戳
	IsDel             int8   `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`                       // 是否删除0未删除 1已删除
	Premission        string `gorm:"column:premission;type:varchar(64);not null" json:"premission"`              // 读/写权限
	SharedInstanceIDs string `gorm:"column:shared_instance_ids;type:text;not null" json:"shared_instance_ids"`   // sharding instances
}

func (t *SharingProject) TableName() string {
	return "sharing_project"
}

// AddProject insert a new Project into database and returns
// last inserted Id on success.
func AddSharingProject(logger *log.Logger, m *SharingProject) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

func GetSharingsByProjectId(logger *log.Logger, project_id string) (ml []*SharingProject, err error) {

	err = dao.Where(logger, dao.IronicRdb, "project_id = ? and is_del = 0", project_id).Find(&ml).Error
	return
}

func GetSharingProjectIdsBySharding(logger *log.Logger, user_id string) (ml []*SharingProject, err error) {
	err = dao.Where(logger, dao.IronicRdb, "shared_user_id = ? and is_del = 0", user_id).Find(&ml).Error
	return
}

func GetOneSharingProject(logger *log.Logger, project_id string, owner_id string, shard_id string) (l *SharingProject, err error) {
	l = &SharingProject{}
	err = dao.Where(logger, dao.IronicRdb, "project_id = ? and shared_user_id = ? and is_del = 0", project_id, shard_id).Take(l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}

func DeleteOneSharingProject(logger *log.Logger, p *SharingProject) (err error) {
	p.IsDel = 1
	err = dao.Save(logger, dao.IronicWdb, p).Error
	return
}

func DeleteSharingProjectByProjectId(logger *log.Logger, project_id string) (err error) {
	m := map[string]interface{}{
		"is_del": 1,
	}
	return dao.Model(logger, dao.IronicWdb, SharingProject{}).Where("project_id = ? and is_del = 0", project_id).Updates(m).Error

}

// GetAllShareProjects
func GetAllShareProjects(logger *log.Logger, query map[string]interface{}) (ml []*SharingProject, err error) {

	var db = dao.Model(logger, dao.IronicRdb, SharingProject{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func UpdateMultiShareProjects(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	var db = dao.Model(logger, dao.IronicWdb, SharingProject{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func UpdateShareProjectById(logger *log.Logger, m *SharingProject) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, SharingProject{}).Where("id = ?", m.ID).Updates(m).Error
}

func UpdateShareProjectByMap(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	var db = dao.Model(logger, dao.IronicWdb, SharingProject{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}
