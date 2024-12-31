package deviceTypeDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type DeviceType struct {
	// 主键
	ID int `gorm:"primaryKey;column:id" json:"-"`
	// 机房id
	IDcID string `gorm:"column:idc_id" json:"idcId"`
	// 设备类型uuid
	DeviceTypeID string `gorm:"column:device_type_id" json:"deviceTypeId"`
	// 机型名称，如计算效能型,标准计算型
	Name string `gorm:"column:name" json:"name"`
	// 机型规格, cps.c.normal
	DeviceType string `gorm:"column:device_type" json:"deviceType"`
	// 机型类型，如计算型，存储型
	DeviceSeries string `gorm:"column:device_series" json:"deviceSeries"`
	// 体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64
	Architecture string `gorm:"column:architecture" json:"architecture"`
	// 【高度（U）】：显示机型高度
	Height int `gorm:"column:height" json:"height"`
	// 描述
	Description string `gorm:"column:description" json:"description"`
	// cpu预置类型
	CpuSpec string `gorm:"column:cpu_spec" json:"cpuSpec"`
	// cpu数量
	CPUAmount int `gorm:"column:cpu_amount" json:"cpuAmount"`
	// 单个cpu内核数
	CPUCores int `gorm:"column:cpu_cores" json:"cpuCores"`
	// cpu厂商
	CPUManufacturer string `gorm:"column:cpu_manufacturer" json:"cpuManufacturer"`
	// cpu处理器型号
	CPUModel string `gorm:"column:cpu_model" json:"cpuModel"`
	// cpu频率(G)
	CPUFrequency string `gorm:"column:cpu_frequency" json:"cpuFrequency"`
	// cpu预置类型
	MemSpec string `gorm:"column:mem_spec" json:"memSpec"`
	// 内存接口（如DDR3，DDR4）
	MemType string `gorm:"column:mem_type" json:"memType"`
	// 单个内存大小(GB)
	MemSize int `gorm:"column:mem_size" json:"memSize"`
	// 内存数量
	MemAmount int `gorm:"column:mem_amount" json:"memAmount"`
	// 内存主频（MHz)
	MemFrequency int `gorm:"column:mem_frequency" json:"memFrequency"`
	// 网卡数量
	NicAmount int `gorm:"column:nic_amount" json:"nicAmount"`
	// 网卡传输速率(GE)
	NicRate int `gorm:"column:nic_rate" json:"nicRate"`
	// 【网口模式】【网络设置】: bond单网口,dual双网口
	InterfaceMode string `gorm:"column:interface_mode" json:"interfaceMode"`
	// 系统盘类型（SSD，HDD）
	SystemVolumeType string `gorm:"column:system_volume_type" json:"systemVolumeType"`
	// 系统盘接口类型（SATA,SAS,NVME）
	SystemVolumeInterfaceType string `gorm:"column:system_volume_interface_type" json:"systemVolumeInterfaceType"`
	// 系统盘单盘大小
	SystemVolumeSize int `gorm:"column:system_volume_size" json:"systemVolumeSize"`
	// 系统盘单盘大小单位
	SystemVolumeUnit string `gorm:"column:system_volume_unit" json:"systemVolumeUnit"`
	// 系统盘数量
	SystemVolumeAmount int `gorm:"column:system_volume_amount" json:"systemVolumeAmount"`
	// 系统盘是否做raid【RAID/NO RAID】
	RaidCan string `gorm:"column:raid_can" json:"raidCan"`
	// gpu数量
	GpuAmount int `gorm:"column:gpu_amount" json:"gpuAmount"`
	// gpu厂商
	GpuManufacturer string `gorm:"column:gpu_manufacturer" json:"gpuManufacturer"`
	// gpu处理器型号
	GpuModel string `gorm:"column:gpu_model" json:"gpuModel"`
	// 数据盘类型（SSD，HDD）
	DataVolumeType string `gorm:"column:data_volume_type" json:"dataVolumeType"`
	// 数据盘接口类型（SATA,SAS,NVME）
	DataVolumeInterfaceType string `gorm:"column:data_volume_interface_type" json:"dataVolumeInterfaceType"`
	// 数据盘单盘大小
	DataVolumeSize int `gorm:"column:data_volume_size" json:"dataVolumeSize"`
	// 系统盘单盘大小单位
	DataVolumeUnit string `gorm:"column:data_volume_unit" json:"dataVolumeUnit"`
	// 数据盘数量
	DataVolumeAmount int `gorm:"column:data_volume_amount" json:"dataVolumeAmount"`
	// 是否配置了阵列卡
	IsNeedRaid string `gorm:"column:is_need_raid" json:"isNeedRaid"`
	// 创建者
	CreatedBy string `gorm:"column:created_by" json:"createdBy"`
	// 更新者
	UpdatedBy string `gorm:"column:updated_by" json:"updatedBy"`
	// 创建时间
	CreatedTime int `gorm:"column:created_time" json:"createdTime"`
	// 更新时间
	UpdatedTime int `gorm:"column:updated_time" json:"updatedTime"`
	// 删除时间
	DeletedTime int `gorm:"column:deleted_time" json:"deletedTime"`
	// boot类型：UEFI Legacy/BIOS
	BootMode string `gorm:"column:boot_mode" json:"bootMode"`
	// 是否删除0未删除 1已删除
	IsDel int8 `gorm:"column:is_del" json:"isDel"`
}

func (t *DeviceType) TableName() string {
	return "device_type"
}

// AddDeviceType insert a new DeviceType into database and returns
// last inserted Id on success.
func AddDeviceType(logger *log.Logger, m *DeviceType) (id int64, err error) {
	return dao.CreateAndGetId(logger, dao.IronicRdb, m)
}

// GetDeviceTypeById retrieves DeviceType by Id. Returns error if
// Id doesn't exist
func GetDeviceTypeById(logger *log.Logger, deviceTypeId string) (v *DeviceType, err error) {
	v = &DeviceType{}
	err = dao.Where(logger, dao.IronicRdb, "device_type_id = ? and is_del = 0", deviceTypeId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func QueryDeviceTypesCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, DeviceType{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return n, err
	}
	err = db.Count(&n).Error
	return
}
func QueryAllDeviceTypes(logger *log.Logger, query map[string]interface{}) (ml []*DeviceType, err error) { //如果需要排序，在下面的代码中定义
	return GetMultiDeviceType(logger, query, nil, []string{"created_time"}, []string{"desc"}, 0, 0)
}

// GetAllCommand retrieves all Command matches certain condition. Returns empty list if
// no records exist
func QueryDeviceTypes(logger *log.Logger, query map[string]interface{}, offset, limit int64) (ml []*DeviceType, err error) {
	return GetMultiDeviceType(logger, query, nil, []string{"created_time"}, []string{"desc"}, offset, limit)
	/*var db = dao.Model(logger, dao.IronicRdb, DeviceType{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return*/
}

func GetOneDeviceType(logger *log.Logger, query map[string]interface{}) (l *DeviceType, err error) {
	l = &DeviceType{}
	var db = dao.Model(logger, dao.IronicRdb, DeviceType{})
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

// GetMultiDeviceType retrieves all DeviceType matches certain condition. Returns empty list if
// no records exist
func GetMultiDeviceType(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*DeviceType, err error) {

	var db = dao.Model(logger, dao.IronicRdb, DeviceType{})
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

// UpdateDeviceType updates DeviceType by Id and returns error if
// the record to be updated doesn't exist
func UpdateDeviceTypeByDeviceTypeID(logger *log.Logger, m *DeviceType, deviceTypeID string) (err error) {
	//m.UpdateUser = logger.GetPoint("username").(string)
	err = dao.Model(logger, dao.IronicWdb, DeviceType{}).Where("device_type_id = ?", deviceTypeID).Save(m).Error
	return
}

// GetAllDeviceTyp retrieves all DeviceTyp matches certain condition. Returns empty list if
// no records exist
func GetAllDeviceType(logger *log.Logger, query map[string]interface{}) (ml []*DeviceType, err error) {

	var db = dao.Model(logger, dao.IronicRdb, DeviceType{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}
func GetByAzAndDeviceType(logger *log.Logger, az, deviceType string) (l *DeviceType, err error) {
	l = &DeviceType{}
	query := map[string]interface{}{
		"is_del":      "0",
		"device_type": deviceType,
		// "az":          az,
	}
	var db = dao.Model(logger, dao.IronicRdb, DeviceType{})
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

func QueryByAz(logger *log.Logger, az string) (ml []*DeviceType, err error) {
	query := map[string]interface{}{
		"is_del": "0",
		// "az":     az,
	}
	var db = dao.Model(logger, dao.IronicRdb, DeviceType{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}
