package rDeviceTypeImageDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceTypeDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/imageDao"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// RDeviceTypeImage device_type/image关系
type RDeviceTypeImage struct {
	ID           uint64 `gorm:"primaryKey;column:id" json:"-"`             // ID
	ImageID      string `gorm:"column:image_id" json:"imageId"`            // 镜像ID
	DeviceTypeID string `gorm:"column:device_type_id" json:"deviceTypeId"` // 设备类型id
	CreatedBy    string `gorm:"column:created_by" json:"createdBy"`        // 创建者
	UpdatedBy    string `gorm:"column:updated_by" json:"updatedBy"`        // 更新者
	CreatedTime  int    `gorm:"column:created_time" json:"createdTime"`    // 创建时间
	UpdatedTime  int    `gorm:"column:updated_time" json:"updatedTime"`    // 更新时间
	DeletedTime  int    `gorm:"column:deleted_time" json:"deletedTime"`    // 删除时间
	IsDel        int8   `gorm:"column:is_del" json:"isDel"`                // 是否删除0未删除 1已删除
}

func (t *RDeviceTypeImage) TableName() string {
	return "r_device_type_image"
}

// AddRDeviceTypeImage insert a new RDeviceTypeImage into database and returns
// last inserted Id on success.
func AddRDeviceTypeImage(logger *log.Logger, m *RDeviceTypeImage) (id int64, err error) {
	//
	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

func AddMultiRDeviceTypeImage(logger *log.Logger, m []*RDeviceTypeImage) (id int64, err error) {

	tx := dao.GetGormTx(logger)
	tx.Begin()
	for _, r := range m {
		//r.CreateTime = time.Now()
		//r.UpdateTime = time.Now()
		if err := tx.Create(r).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return int64(len(m)), nil
}

// GetRDeviceTypeImageById retrieves RDeviceTypeImage by Id. Returns error if
// Id doesn't exist
func GetRDeviceTypeImageById(logger *log.Logger, id int64) (v *RDeviceTypeImage, err error) {
	v = &RDeviceTypeImage{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllCommand retrieves all Command matches certain condition. Returns empty list if
// no records exist
func GetAllRDeviceTypeImage(logger *log.Logger, query map[string]interface{}) (ml []*RDeviceTypeImage, err error) {

	var db = dao.Model(logger, dao.IronicRdb, RDeviceTypeImage{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiRDeviceTypeImage retrieves all RDeviceTypeImage matches certain condition. Returns empty list if
// no records exist
func GetMultiRDeviceTypeImage(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*RDeviceTypeImage, err error) {

	var db = dao.Model(logger, dao.IronicRdb, RDeviceTypeImage{})
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

// UpdateRDeviceTypeImage updates RDeviceTypeImage by Id and returns error if
// the record to be updated doesn't exist
func UpdateRDeviceTypeImageById(logger *log.Logger, m *RDeviceTypeImage, deviceTypeId, imageId string) (err error) {
	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, RDeviceTypeImage{}).Where("device_type_id = ? and image_id = ? and is_del = ?", deviceTypeId, imageId, baseLogic.IS_NOT_DEL).Update(m).Error
}

func UpdateMultiRDeviceTypeImage(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	updates["updated_time"] = int(time.Now().Unix())
	var db = dao.Model(logger, dao.IronicWdb, RDeviceTypeImage{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func AssociatedImageMultiRDeviceTypeImage(logger *log.Logger, items []*RDeviceTypeImage) (err error) {
	//paris := []string{}
	for _, item := range items {
		//paris = append(paris, fmt.Sprintf("('%s', '%s')", item.DeviceTypeID, item.ImageID))
		err = dao.Model(logger, dao.IronicRdb, RDeviceTypeImage{}).Where("device_type_id = ? and image_id = ? and is_del = 0", item.DeviceTypeID, item.ImageID).Update(item).Error
	}
	//fmt.Sprintf(`(device_type_id, image_id) in (%s)`, strings.Join(paris, ","))

	//updates := map[string]interface{}{
	//	"is_del":      0,
	//	"updated_time": time.Now(),
	//}
	return
}

func GetOneRDeviceTypeImage(logger *log.Logger, query map[string]interface{}) (l *RDeviceTypeImage, err error) {
	l = &RDeviceTypeImage{}
	var db = dao.Model(logger, dao.IronicRdb, RDeviceTypeImage{})
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

func GetByDeviceTypeIDAndImageID(logger *log.Logger, deviceTypeID, imageID string) (l *RDeviceTypeImage, err error) {
	l = &RDeviceTypeImage{}
	err = dao.Where(logger, dao.IronicRdb, "device_type_id = ? and image_id = ? and is_del = ?", deviceTypeID, imageID, baseLogic.IS_NOT_DEL).Take(l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}
func GetByDeviceTypeID(logger *log.Logger, deviceTypeID string) (l []*RDeviceTypeImage, err error) {
	err = dao.Where(logger, dao.IronicRdb, "device_type_id = ? and is_del = ?", deviceTypeID, baseLogic.IS_NOT_DEL).Find(&l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}

func GetByDeviceTypeAndImageIdPatition(logger *log.Logger, device_type, image_id string) (ml []*RDeviceTypeImage, err error) {

	sql := fmt.Sprintf("select r from r_device_type_image r where (device_type='%s'  or device_type='common') and image_id='%s' and is_del=0", device_type, image_id)
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}
func QueryDeviceTypeImagesCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, RDeviceTypeImage{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return n, err
	}
	err = db.Count(&n).Error
	return
}

func ConvertIn(str string) string {
	arr := strings.Split(str, ",")
	sqlIn := ""
	for _, v := range arr {
		sqlIn = sqlIn + "'" + v + "'" + ","
	}
	sqlIn = strings.Trim(sqlIn, ",")
	return "(" + sqlIn + ")"
}

// QueryByDeviceOsType ...
func GetDeviceTypeImageAndOsList(logger *log.Logger, param map[string]interface{}, count bool, offset int64, limit int64) (ml []*imageDao.DImage, n int64, err error) {
	raw := "*"
	if count {
		raw = "count(*)"
	}
	sql := `select ` + raw + ` FROM  image i,os o,r_device_type_image r where i.os_id=o.os_id and r.image_id=i.image_id and r.is_del=0 and i.is_del=0 and o.is_del=0`
	sql = sql + ` and r.device_type_id=` + "'" + param["device_type_id"].(string) + "'"
	//image i  放在前面，如果有相同字段，以前面这为主，由于都有添加时间字段，如果要获取镜像的添加时间，所以要以image表的created_time为准
	//fmt.Println(param)
	if param["image_id"] != nil && param["image_id"] != "" {
		sql = sql + " and i.image_id in " + ConvertIn(param["image_id"].(string))
	}
	if param["source"] != nil && param["source"] != "" {
		sql = sql + " and i.source in " + ConvertIn(param["source"].(string))
	}
	if param["image_name"] != nil && param["image_name"] != "" {
		sql = sql + " and i.image_name in " + ConvertIn(param["image_name"].(string))
	}
	if param["version"] != nil && param["version"] != "" {
		sql = sql + " and o.os_version in " + ConvertIn(param["version"].(string))
	}
	if param["os_type"] != nil && param["os_type"] != "" {
		sql = sql + " and o.os_type in " + ConvertIn(param["os_type"].(string))
	}
	if param["architecture"] != nil && param["architecture"] != "" {
		sql = sql + " and o.architecture in " + ConvertIn(param["architecture"].(string))
	}
	sql = sql + " order by i.created_time desc"
	fmt.Println(sql)
	logger.Info("查询机型关联镜像sql:", sql)
	//os.Exit(1)
	if count {
		err = dao.Raw(logger, dao.IronicRdb, sql).Count(&n).Error
		return
	}
	if offset == 0 && limit == 0 {
		err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
		return
	}
	fmt.Println(sql)
	err = dao.Raw(logger, dao.IronicRdb, sql).Offset(offset).Limit(limit).Scan(&ml).Error
	return
}

//查询镜像绑定的机型列表
func GetImageDeviceTypeList(logger *log.Logger, param map[string]interface{}, count bool, offset int64, limit int64) (ml []*deviceTypeDao.DeviceType, n int64, err error) {
	raw := "*"
	if count {
		raw = "count(*)"
	}

	sql := ""
	//查询已绑定的机型
	if param["is_bind"] != nil && param["is_bind"] == "1" {
		sql = `select ` + raw + ` FROM  device_type dt,r_device_type_image r where dt.device_type_id=r.device_type_id  and dt.is_del=0 and r.is_del=0`
		if param["image_id"] != nil && param["image_id"] != "" {
			sql = sql + " and r.image_id in " + ConvertIn(param["image_id"].(string))
		}
	} else { //未绑定
		if param["image_id"] != nil && param["image_id"] != "" {
			sql = `select ` + raw + ` FROM  device_type dt where dt.is_del=0`
			sql = sql + " and dt.device_type_id not in (select device_type_id from r_device_type_image where image_id = '" + param["image_id"].(string) + "' and is_del = 0)"
		}
	}

	//if param["source"] != nil && param["source"] != "" {
	//	sql = sql + " and i.source in " + ConvertIn(param["source"].(string))
	//}
	//if param["image_name"] != nil && param["image_name"] != "" {
	//	sql = sql + " and i.image_name in " + ConvertIn(param["image_name"].(string))
	//}
	//if param["version"] != nil && param["version"] != "" {
	//	sql = sql + " and o.os_version in " + ConvertIn(param["version"].(string))
	//}
	//if param["os_type"] != nil && param["os_type"] != "" {
	//	sql = sql + " and o.os_type in " + ConvertIn(param["os_type"].(string))
	//}
	if param["architecture"] != nil && param["architecture"] != "" {
		sql = sql + " and dt.architecture in " + ConvertIn(param["architecture"].(string))
	}
	sql = sql + " order by dt.created_time desc"
	fmt.Println(sql)
	logger.Info("查询机型关联镜像sql:", sql)
	//os.Exit(1)
	if count {
		err = dao.Raw(logger, dao.IronicRdb, sql).Count(&n).Error
		return
	}
	if offset == 0 && limit == 0 {
		err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
		return
	}
	fmt.Println(sql)
	err = dao.Raw(logger, dao.IronicRdb, sql).Offset(offset).Limit(limit).Scan(&ml).Error
	return
}
