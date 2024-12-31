package request

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"github.com/go-playground/validator/v10"
)

type CreateDiskRequest struct {
	DeviceID    string `gorm:"column:device_id" json:"deviceId"`       // 设备uuid
	Enclosure   string `gorm:"column:enclosure" json:"enclosure"`      // enclosure
	Slot        int    `gorm:"column:slot" json:"slot"`                // 卡槽槽位
	DiskType    string `gorm:"column:disk_type" json:"diskType"`       // 磁盘类型：system,data
	Size        string `gorm:"column:size" json:"size"`                // 硬盘大小，不确定精度(非nvme盘)
	SizeUnit    string `gorm:"column:size_unit" json:"sizeUnit"`       // 硬盘大小单位 MB GB TB ,1024进制
	PdType      string `gorm:"column:pd_type" json:"pdType"`           // 硬盘类型
	AdapterID   string `gorm:"column:adapter_id" json:"adapterId"`     // 适配ID
	CreatedBy   string `gorm:"column:created_by" json:"createdBy"`     // 创建者
	UpdatedBy   string `gorm:"column:updated_by" json:"updatedBy"`     // 更新者
	CreatedTime int    `gorm:"column:created_time" json:"createdTime"` // 创建时间戳
	UpdatedTime int    `gorm:"column:updated_time" json:"updatedTime"` // 更新时间戳
	//DeletedTime int    `gorm:"column:deleted_time" json:"deletedTime"` // 删除时间戳
	//IsDel       int8   `gorm:"column:is_del" json:"isDel"`             // 是否删除0未删除 1已删除
}

func (c *CreateDiskRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(c); err != nil {
		logger.Warn("CreateDiskRequest validate error:", err.Error())
		panic(constant.INVALID_ARGUMENT)
	}
}
