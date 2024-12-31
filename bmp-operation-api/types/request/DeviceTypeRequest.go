package request

import sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"

// 机型管理 模块请求字段
type CreateDeviceTypeRequest struct {
	//ID                        int    `gorm:"primaryKey;column:id" json:"-"`                                        // 主键
	IDcID string `json:"idcId"`
	//DeviceTypeID              string `json:"deviceTypeId"`              // 设备类型uuid
	Name                      string  `json:"name"`                      // 机型名称，如计算效能型,标准计算型
	DeviceType                string  `json:"deviceType"`                // 机型规格, cps.c.normal
	DeviceSeries              string  `json:"deviceSeries"`              // 机型类型，如计算型，存储型
	Architecture              string  `json:"architecture"`              // 体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64
	Height                    int     `json:"height"`                    // 【高度（U）】：显示机型高度
	Description               string  `json:"description"`               // 描述
	CPUAmount                 int64   `json:"cpuAmount"`                 // cpu数量
	CPUCores                  int64   `json:"cpuCores"`                  // 单个cpu内核数
	CPUManufacturer           string  `json:"cpuManufacturer"`           // cpu厂商
	CPUModel                  string  `json:"cpuModel"`                  // cpu处理器型号
	CPUFrequency              string  `json:"cpuFrequency"`              // cpu频率(G)
	MemType                   string  `json:"memType"`                   // 内存接口（如DDR3，DDR4）
	MemSize                   int     `json:"memSize"`                   // 单个内存大小(GB)
	MemAmount                 int     `json:"memAmount"`                 // 内存数量
	MemFrequency              int     `json:"memFrequency"`              // 内存主频（MHz)
	NicAmount                 int64   `json:"nicAmount"`                 // 网卡数量
	NicRate                   int     `json:"nicRate"`                   // 网卡传输速率(GE)
	InterfaceMode             string  `json:"interfaceMode"`             // 【网口模式】【网络设置】: bond单网口,dual双网口
	SystemVolumeType          string  `json:"systemVolumeType"`          // 系统盘类型（SSD，HDD）
	SystemVolumeInterfaceType string  `json:"systemVolumeInterfaceType"` // 系统盘接口类型（SATA,SAS,NVME）
	SystemVolumeSize          float64 `json:"systemVolumeSize"`          // 系统盘单盘大小
	SystemVolumeUnit          string  `json:"systemVolumeUnit"`          // 系统盘单盘大小单位
	SystemVolumeAmount        int64   `json:"systemVolumeAmount"`        // 系统盘数量
	// 是否做raid，[RAID/NO RAID]
	RaidCan                 string  `json:"raidCan"`
	GpuAmount               int64   `json:"gpuAmount"`               // gpu数量
	GpuManufacturer         string  `json:"gpuManufacturer"`         // gpu厂商
	GpuModel                string  `json:"gpuModel"`                // gpu处理器型号
	DataVolumeType          string  `json:"dataVolumeType"`          // 数据盘类型（SSD，HDD）
	DataVolumeInterfaceType string  `json:"dataVolumeInterfaceType"` // 数据盘接口类型（SATA,SAS,NVME）
	DataVolumeSize          float64 `json:"dataVolumeSize"`          // 数据盘单盘大小
	DataVolumeUnit          string  `json:"dataVolumeUnit"`
	DataVolumeAmount        int64   `json:"dataVolumeAmount"` // 数据盘数量
	CreatedBy               string  `json:"createdBy"`        // 创建者
	UpdatedBy               string  `json:"updatedBy"`        // 更新者
	CreatedTime             int     `json:"createdTime"`      // 创建时间
	UpdatedTime             int     `json:"updatedTime"`      // 更新时间
	RaidID                  string  `json:"raidId"`           //系统盘raid，支持多选
	// CPU 预置规格
	// required: true
	CpuSpec string `json:"cpuSpec" validate:"required,oneof=common user_defined"`
	// 内存 预置规格
	// required: true
	MemSpec string `json:"memSpec" validate:"required,oneof=common user_defined"`

	//boot模式[UEFI Legacy/BIOS]支持多选，逗号分隔
	BootMode string `json:"bootMode" validate:"omitempty"`

	Volumes    []*sdkModels.VolumeItem `json:"volumes"`
	IsNeedRaid string                  `json:"isNeedRaid"`
}

// 【机型管理】【编辑机型】
type ModifyDeviceTypeRequest struct {
	//ID                        int    `gorm:"primaryKey;column:id" json:"-"`                                        // 主键
	IDcID        *string `json:"idcId"`
	DeviceTypeID string  `json:"deviceTypeId"` // 设备类型uuid
	Name         *string `json:"name"`         // 机型名称，如计算效能型,标准计算型
	DeviceType   *string `json:"deviceType"`   // 机型规格, cps.c.normal
	DeviceSeries *string `json:"deviceSeries"` // 机型类型，如计算型，存储型
	Architecture *string `json:"architecture"` // 体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64
	Height       *int64  `json:"height"`       // 【高度（U）】：显示机型高度
	Description  *string `json:"description"`  // 描述
	CPUAmount    *int64  `json:"cpuAmount"`    // cpu数量
	CPUCores     *int64  `json:"cpuCores"`     // 单个cpu内核数

	CPUManufacturer           *string  `json:"cpuManufacturer"`           // cpu厂商
	CPUModel                  *string  `json:"cpuModel"`                  // cpu处理器型号
	CPUFrequency              *string  `json:"cpuFrequency"`              // cpu频率(G)
	MemType                   *string  `json:"memType"`                   // 内存接口（如DDR3，DDR4）
	MemSize                   *int64   `json:"memSize"`                   // 单个内存大小(GB)
	MemAmount                 *int64   `json:"memAmount"`                 // 内存数量
	MemFrequency              *int64   `json:"memFrequency"`              // 内存主频（MHz)
	NicAmount                 *int64   `json:"nicAmount"`                 // 网卡数量
	NicRate                   *int64   `json:"nicRate"`                   // 网卡传输速率(GE)
	InterfaceMode             *string  `json:"interfaceMode"`             // 【网口模式】【网络设置】: bond单网口,dual双网口
	SystemVolumeType          *string  `json:"systemVolumeType"`          // 系统盘类型（SSD，HDD）
	SystemVolumeInterfaceType *string  `json:"systemVolumeInterfaceType"` // 系统盘接口类型（SATA,SAS,NVME）
	SystemVolumeSize          *float64 `json:"systemVolumeSize"`          // 系统盘单盘大小
	SystemVolumeUnit          *string  `json:"systemVolumeUnit"`          // 系统盘单位（GB，TB）
	SystemVolumeAmount        *int64   `json:"systemVolumeAmount"`        // 系统盘数量
	// 是否做raid，[RAID/NO RAID]
	//RaidCan                 string   `json:"raidCan"`
	GpuAmount               *int64   `json:"gpuAmount"`               // gpu数量
	GpuManufacturer         *string  `json:"gpuManufacturer"`         // gpu厂商
	GpuModel                *string  `json:"gpuModel"`                // gpu处理器型号
	DataVolumeType          *string  `json:"dataVolumeType"`          // 数据盘类型（SSD，HDD）
	DataVolumeInterfaceType *string  `json:"dataVolumeInterfaceType"` // 数据盘接口类型（SATA,SAS,NVME）
	DataVolumeSize          *float64 `json:"dataVolumeSize"`          // 数据盘单盘大小
	DataVolumeUnit          *string  `json:"dataVolumeUnit"`          // 数据盘单位（GB，TB）
	DataVolumeAmount        *int64   `json:"dataVolumeAmount"`        // 数据盘数量
	CreatedBy               *string  `json:"createdBy"`               // 创建者
	UpdatedBy               *string  `json:"updatedBy"`               // 更新者
	CreatedTime             *int     `json:"createdTime"`             // 创建时间
	UpdatedTime             *int     `json:"updatedTime"`             // 更新时间

	RaidID *string `json:"raidId"` //系统盘raid，支持多选
	// CPU 预置规格
	CpuSpec *string `json:"cpuSpec" validate:"omitempty,oneof=common user_defined"`
	// 内存 预置规格
	MemSpec *string `json:"memSpec" validate:"omitempty,oneof=common user_defined"`
	//boot模式[UEFI Legacy/BIOS]支持多选，逗号分隔
	BootMode   string                  `json:"bootMode" validate:"omitempty"`
	Volumes    []*sdkModels.VolumeItem `json:"volumes"`
	IsNeedRaid *string                 `json:"isNeedRaid"`
}

// 【机型管理】【机型列表】
type QueryDeviceTypesRequest struct {
	BaseRequest
	PagingRequest
	DeviceTypeID string `json:"deviceTypeId"`
	DeviceType   string `json:"deviceType"` // 机型规格, cps.c.normal
	IdcID        string `json:"idcId"`
	Name         string `json:"name"`         // 机型名称，如计算效能型,标准计算型
	DeviceSeries string `json:"deviceSeries"` // 机型类型，如计算型，存储型
	IsAll        string `json:"isAll"`
	// 非空表示导出
	ExportType string `json:"exportType"`
}

// 【机型管理】【机型详情信息】
type QueryDeviceTypeRequest struct {
	DeviceTypeID string `json:"deviceTypeID"`
}

// 【机型管理】【删除】
type DeleteDeviceTypeRequest struct {
	DeviceTypeID string `json:"deviceTypeID"`
}

type GetAvailableDeviceTypesRequest struct {
	IdcID string `json:"idc_id"`
}
