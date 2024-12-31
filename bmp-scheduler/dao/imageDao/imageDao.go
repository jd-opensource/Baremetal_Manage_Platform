package imageDao

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"fmt"
	log "git.jd.com/cps-golang/log"
	"strings"
)

// Image 镜像
type Image struct {
	ID        uint64 `gorm:"primaryKey;column:id" json:"-"`      // ID
	ImageID   string `gorm:"column:image_id" json:"imageId"`     // 镜像uuid
	ImageName string `gorm:"column:image_name" json:"imageName"` // 镜像名称
	//ImageVersion    string `gorm:"column:image_version" json:"imageVersion"`       // 镜像版本
	OsID            string `gorm:"column:os_id" json:"osId"`                       // OSID
	Format          string `gorm:"column:format" json:"format"`                    // 镜像格式（qcow2、tar）
	Filename        string `gorm:"column:filename" json:"filename"`                // 镜像文件名称
	URL             string `gorm:"column:url" json:"url"`                          // 镜像源路径
	Hash            string `gorm:"column:hash" json:"hash"`                        // 镜像校验码
	Source          string `gorm:"column:source" json:"source"`                    // 镜像来源(common通用、customize定制、user_defined自定义)
	Description     string `gorm:"column:description" json:"description"`          // 描述
	SystemPartition string `gorm:"column:system_partition" json:"systemPartition"` // 系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：10GiB，swap）
	DataPartition   string `gorm:"column:data_partition" json:"dataPartition"`     // 数据分区信息
	CreatedBy       string `gorm:"column:created_by" json:"createdBy"`             // 创建者
	UpdatedBy       string `gorm:"column:updated_by" json:"updatedBy"`             // 更新者
	CreatedTime     int    `gorm:"column:created_time" json:"createdTime"`         // 创建时间戳
	UpdatedTime     int    `gorm:"column:updated_time" json:"updatedTime"`         // 更新时间戳
	DeletedTime     int    `gorm:"column:deleted_time" json:"deletedTime"`         // 删除时间戳
	IsDel           int8   `gorm:"column:is_del" json:"isDel"`                     // 是否删除0未删除 1已删除
}

func (t *Image) TableName() string {
	return "image"
}

// AddImage insert a new Image into database and returns
// last inserted Id on success.
func AddImage(logger *log.Logger, m *Image) (id int64, err error) {
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetImageById retrieves Image by Id. Returns error if
// Id doesn't exist
func GetImageById(logger *log.Logger, id int64) (v *Image, err error) {
	v = &Image{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetImageByUuid retrieves Image by Uuid. Returns error if
// Id doesn't exist
func GetImageByUuid(logger *log.Logger, imageId string) (v *Image, err error) {
	v = &Image{}
	err = dao.Where(logger, dao.IronicRdb, "image_id = ? and is_del = 0", imageId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}
func QueryImagesCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, Image{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return n, err
	}
	err = db.Count(&n).Error
	return
}
func GetAllImage(logger *log.Logger, query map[string]interface{}) (ml []*Image, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Image{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiImage retrieves all Image matches certain condition. Returns empty list if
// no records exist
func GetMultiImage(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Image, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Image{})
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

// UpdateImage updates Image by Id and returns error if
// the record to be updated doesn't exist
func UpdateImageById(logger *log.Logger, m *Image) (err error) {
	return dao.Model(logger, dao.IronicWdb, Image{}).Where("image_id = ?", m.ImageID).Updates(m).Error
}

type DImage struct {
	ID              uint64 `gorm:"primaryKey;column:id" json:"-"`                  // ID
	ImageID         string `gorm:"column:image_id" json:"imageId"`                 // 镜像uuid
	ImageName       string `gorm:"column:image_name" json:"imageName"`             // 镜像名称
	ImageVersion    string `gorm:"column:image_version" json:"imageVersion"`       // 镜像版本
	Format          string `gorm:"column:format" json:"format"`                    // 镜像格式（qcow2、tar）
	Filename        string `gorm:"column:filename" json:"filename"`                // 镜像文件名称
	URL             string `gorm:"column:url" json:"url"`                          // 镜像源路径
	Hash            string `gorm:"column:hash" json:"hash"`                        // 镜像校验码
	Source          string `gorm:"column:source" json:"source"`                    // 镜像来源(common通用、customize定制、user_defined自定义)
	Description     string `gorm:"column:description" json:"description"`          // 描述
	SystemPartition string `gorm:"column:system_partition" json:"systemPartition"` // 系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：10GiB，swap）
	DataPartition   string `gorm:"column:data_partition" json:"dataPartition"`     // 数据分区信息
	OsID            string `json:"OsId"`
	OsName          string `json:"osName"`
	OsVersion       string `json:"osVersion"`
	OsType          string `json:"osType"`
	Platform        string `json:"platform"`
	OsArchitecture  string `json:"architecture"`
	Bits            int32  `json:"bits"`
	CreatedBy       int    `json:"createdBy"`
	UpdatedBy       int    `json:"updatedBy"`
	CreatedTime     int    `json:"createdTime"`
	UpdatedTime     int    `json:"updatedTime"`
}

// QueryByDeviceOsType ...
func QueryByDeviceOsType(logger *log.Logger, param map[string]string) (ml []*DImage, err error) {

	sql := `select i.id,i.image_id,i.image_name,i.image_version,i.format,i.url,i.filename,i.description,i.os_id,i.hash,i.source,
		i.created_by,i.updated_by,i.created_time,i.updated_time FROM  image i`
	if param["device_type"] != "" {
		sql = sql + " LEFT JOIN r_device_type_image rd on i.uuid=rd.image_id"
	}
	if param["region"] != "" {
		sql = sql + " LEFT JOIN r_region_image re on i.uuid=re.image_id"
	}
	sql = sql + " LEFT JOIN os o on i.os_id=o.os_id"
	where := " where 1 = 1"
	if param["device_type"] != "" {
		where = where + " and rd.device_type=" + fmt.Sprintf("'%s'", param["device_type"])
	}
	if param["region"] != "" {
		where = where + " and re.region=" + fmt.Sprintf("'%s'", param["region"])
	}
	if param["os_type"] != "" {
		where = where + " and o.os_type=" + fmt.Sprintf("'%s'", param["os_type"])
	}
	if param["os_version"] != "" {
		where = where + " and o.os_version=" + fmt.Sprintf("'%s'", param["os_version"])
	}
	if param["os_id"] != "" {
		where = where + " and i.os_id=" + fmt.Sprintf("'%s'", param["os_id"])
	}
	if param["name_en"] != "" {
		where = where + fmt.Sprintf(" and i.name_en LIKE '%%s%'", param["name_en"])
	}
	if param["image_ids"] != "" {
		imageIds := strings.Split(param["image_ids"], ",")
		for idx, v := range imageIds {
			imageIds[idx] = fmt.Sprintf("'%s'", v)
		}
		inSql := fmt.Sprintf("(%s)", strings.Join(imageIds, ","))
		where = where + fmt.Sprintf(" and i.uuid in %s", inSql)
	}
	if param["source"] != "" {
		where = where + fmt.Sprintf(" and i.source='%s'", param["source"])
	}
	if param["device_type"] != "" {
		where = where + " and rd.is_del=0"
	}
	if param["region"] != "" {
		where = where + " and re.is_del=0"
	}
	sql = sql + where + " and i.is_del=0 and o.is_del=0"
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}
