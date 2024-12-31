package sharingProjectDao

import (
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/projectDao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// SharingProject 共享项目
type SharingProject struct {
	ID             int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	ProjectID      string `gorm:"column:project_id;type:varchar(255);not null" json:"project_id"`             // 项目uuid
	ProjectName    string `gorm:"column:project_name;type:varchar(255);not null" json:"project_name"`         // 项目名称
	OwnerUserID    string `gorm:"column:owner_user_id;type:varchar(255);not null" json:"owner_user_id"`       // 项目拥有者用户id
	SharedUserID   string `gorm:"column:shared_user_id;type:varchar(255);not null" json:"shared_user_id"`     // 项目共享者用户id
	OwnerUserName  string `gorm:"column:owner_user_name;type:varchar(255);not null" json:"owner_user_name"`   // 项目拥有者用户名
	SharedUserName string `gorm:"column:shared_user_name;type:varchar(255);not null" json:"shared_user_name"` // 项目拥有者用户名
	IsDefault      int8   `gorm:"column:is_default;type:tinyint(4);not null" json:"is_default"`               // 是否默认项目0否 1是
	CreatedBy      string `gorm:"column:created_by;type:varchar(255);not null" json:"created_by"`             // 创建者
	UpdatedBy      string `gorm:"column:updated_by;type:varchar(255);not null" json:"updated_by"`             // 更新者
	CreatedTime    int    `gorm:"column:created_time;type:int(11);not null" json:"created_time"`              // 创建时间戳
	UpdatedTime    int    `gorm:"column:updated_time;type:int(11);not null" json:"updated_time"`              // 更新时间戳
	DeletedTime    int    `gorm:"column:deleted_time;type:int(11);not null" json:"deleted_time"`              // 删除时间戳
	IsDel          int8   `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`                       // 是否删除0未删除 1已删除
	Premission     string `gorm:"column:premission;type:varchar(64);not null" json:"premission"`              // 读/写权限
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
	err = dao.Where(logger, dao.IronicRdb, "project_id = ? and owner_user_id = ? and shared_user_id = ? and is_del = 0", project_id, owner_id, shard_id).Take(l).Error
	return
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

func UpdateShareProjects(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	updates["updated_time"] = time.Now().Unix()
	var db = dao.Model(logger, dao.IronicWdb, SharingProject{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return

}

// GetProjectByUuid retrieves Project by Uuid. Returns error if
// Id doesn't exist
func GetShareProjectCountGroupByUser(logger *log.Logger) ([]*projectDao.UserProjectCount, error) {
	ml := []*projectDao.UserProjectCount{}
	sql := `select shared_user_id as user_id, count(1) as count from sharing_project group by shared_user_id, is_del having is_del = 0`
	err := dao.Raw(logger, dao.IronicRdb, sql).Find(&ml).Error
	return ml, err
}
