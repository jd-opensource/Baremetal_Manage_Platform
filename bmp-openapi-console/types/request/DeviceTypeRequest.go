package request

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type CreateDeviceTypeRequest struct {
	// 机房id
	// required: true
	IDcID string `json:"idcId"`
	// 机型名称1-64
	// required: true
	Name string `json:"name" validate:"required,min=1,max=64"`
	// 机型规格1-64
	// required: true
	DeviceType string `json:"deviceType" validate:"required,min=1,max=64"`
	// 机型类型，如computer storage gpu other
	// required: true
	DeviceSeries string `json:"deviceSeries" validate:"required,oneof=computer storage gpu other"`
	// 体系架构，1-64
	// required: true
	Architecture string `json:"architecture" validate:"required,oneof=x86_64 ARM64(aarch64) i386"`
	// 【高度（U）】：显示机型高度 2,4
	// required: true
	Height int `json:"height" validate:"required,oneof=2 4"`
	// 描述0-256
	Description string `json:"description" validate:"max=256"`
	// cpu数量1,10000
	// required: true
	CPUAmount int `json:"cpuAmount" validate:"required,min=1,max=10000"`
	// 单个cpu内核数1,10000
	// required: true
	CPUCores int `json:"cpuCores" validate:"required,min=1,max=10000"`
	// cpu厂商1-64
	// required: true
	CPUManufacturer string `json:"cpuManufacturer" validate:"required,min=1,max=64"`
	// cpu处理器型号1-64
	// required: true
	CPUModel string `json:"cpuModel" validate:"required,min=1,max=64"`
	// cpu频率(G)1-64
	// required: true
	CPUFrequency string `json:"cpuFrequency" validate:"required,min=1,max=64"`
	// 内存接口（DDR DDR2 DDR3 DDR4 DDR5）
	// required: true
	MemType string `json:"memType" validate:"required,oneof=DDR DDR2 DDR3 DDR4 DDR5"`
	// 单个内存大小1,10000(GB)
	// required: true
	MemSize int `json:"memSize" validate:"required,min=1,max=10000"`
	// 内存数量1-10000
	// required: true
	MemAmount int `json:"memAmount" validate:"required,min=1,max=10000"`
	// 内存主频（MHz)1-10000
	// required: true
	MemFrequency int `json:"memFrequency" validate:"required,min=1,max=10000"`
	// 网卡数量1-10000
	// required: true
	NicAmount int `json:"nicAmount" validate:"required,min=1,max=10000"`
	// 网卡传输速率(GE) 1-10000
	// required: true
	NicRate int `json:"nicRate" validate:"required,min=1,max=10000"`
	// 【网口模式】【网络设置】: bond,dual
	// required: true
	InterfaceMode string `json:"interfaceMode" validate:"required,oneof=bond dual"`
	// 系统盘类型（SSD，HDD）
	// required: true
	SystemVolumeType string `json:"systemVolumeType" validate:"required,oneof=SSD HDD"`
	// 系统盘接口类型（SATA,SAS,NVME）
	// required: true
	SystemVolumeInterfaceType string `json:"systemVolumeInterfaceType" validate:"required,oneof=SATA SAS NVME"`
	// 系统盘单盘大小1-100000
	// required: true
	SystemVolumeSize int `json:"systemVolumeSize" validate:"required,min=1,max=100000"`
	// 系统盘单盘大小单位 GB,TB
	// required: true
	SystemVolumeUnit string `json:"systemVolumeUnit" validate:"required,oneof=GB TB"`
	// 系统盘数量1-10000
	// required: true
	SystemVolumeAmount int `json:"systemVolumeAmount" validate:"required,min=1,max=10000"`
	// gpu数量0-1000
	GpuAmount int `json:"gpuAmount" validate:"min=0,max=1000"`
	// gpu品牌0-64
	GpuManufacturer string `json:"gpuManufacturer" validate:"min=0,max=64"`
	// gpu型号0-64
	GpuModel string `json:"gpuModel" validate:"min=0,max=64"`
	// 数据盘类型0-64
	DataVolumeType string `json:"dataVolumeType" validate:"min=0,max=64"`
	// 数据盘接口类型0-64
	DataVolumeInterfaceType string `json:"dataVolumeInterfaceType" validate:"min=0,max=64"`
	// 数据盘单盘大小0-100000
	DataVolumeSize int `json:"dataVolumeSize" validate:"min=0,max=100000"`
	// 数据盘单盘大小单位GB TB
	DataVolumeUnit string `json:"dataVolumeUnit" validate:"min=0,max=2"`
	// 数据盘数量0-10000
	DataVolumeAmount int `json:"dataVolumeAmount" validate:"min=0,max=10000"`
	// 系统盘raid，支持多选,英文逗号分隔
	// required: true
	RaidID string `json:"raidId" validate:"required"`
	// CPU 预置规格
	// required: true
	CpuSpec string `json:"cpuSpec" validate:"required,oneof=common user_defined"`
	// 内存 预置规格
	// required: true
	MemSpec string `json:"memSpec" validate:"required,oneof=common user_defined"`
	//boot模式
	BootMode string `json:"boot_mode" validate:"omitempty"`
}

func (param *CreateDeviceTypeRequest) Validate(logger *log.Logger) {
	validate(param, logger)
	if param.GpuManufacturer == "" {
		if param.GpuModel != "" || param.GpuAmount != 0 {
			panic(constant.BuildInvalidArgumentWithArgs("gpu不合法", "gpu invalidate"))
		}
	}
	if param.GpuModel == "" {
		if param.GpuManufacturer != "" || param.GpuAmount != 0 {
			panic(constant.BuildInvalidArgumentWithArgs("gpu不合法", "gpu invalidate"))
		}
	}
	if param.GpuAmount == 0 {
		if param.GpuManufacturer != "" || param.GpuModel != "" {
			panic(constant.BuildInvalidArgumentWithArgs("gpu不合法", "gpu invalidate"))
		}
	}
	if param.DataVolumeType == "" {
		if param.DataVolumeInterfaceType != "" || param.DataVolumeAmount != 0 || param.DataVolumeSize != 0 {
			panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
		}
	}
	if param.DataVolumeInterfaceType == "" {
		if param.DataVolumeType != "" || param.DataVolumeAmount != 0 || param.DataVolumeSize != 0 {
			panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
		}
	}
	if param.DataVolumeAmount == 0 {
		if param.DataVolumeInterfaceType != "" || param.DataVolumeType != "" || param.DataVolumeSize != 0 {
			panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
		}
	}
	if param.DataVolumeSize == 0 {
		if param.DataVolumeInterfaceType != "" || param.DataVolumeAmount != 0 || param.DataVolumeType != "" {
			panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
		}
	}
}

type ModifyDeviceTypeRequest struct {
	// 机房id
	// Extensions:
	// x-nullable: true
	IDcID *string `json:"idcId" validate:"omitempty,min=1,max=64"`
	// 机型名称
	// Extensions:
	// x-nullable: true
	Name *string `json:"name" validate:"omitempty,min=1,max=64"`
	// 机型规格
	// Extensions:
	// x-nullable: true
	DeviceType *string `json:"deviceType" validate:"omitempty,min=1,max=64"`
	// 机型类型，如计算型，存储型
	// Extensions:
	// x-nullable: true
	DeviceSeries *string `json:"deviceSeries" validate:"omitempty,oneof=computer storage gpu other"`
	// 体系架构，如i386/x86_64/ ARM64(aarch64)，默认 x86_64
	// Extensions:
	// x-nullable: true
	Architecture *string `json:"architecture" validate:"omitempty,oneof=x86_64 'ARM64(aarch64)' i386"`
	// 【高度（U）】：显示机型高度
	// Extensions:
	// x-nullable: true
	Height *int `json:"height" validate:"omitempty,oneof=2 4"`
	// 描述
	// Extensions:
	// x-nullable: true
	Description *string `json:"description" validate:"omitempty,min=0,max=256"`
	// cpu数量
	// Extensions:
	// x-nullable: true
	CPUAmount *int `json:"cpuAmount" validate:"omitempty,min=0,max=10000"`
	// 单个cpu内核数
	// Extensions:
	// x-nullable: true
	CPUCores *int `json:"cpuCores" validate:"omitempty,min=0,max=10000"`
	// cpu厂商
	// Extensions:
	// x-nullable: true
	CPUManufacturer *string `json:"cpuManufacturer" validate:"omitempty,min=1,max=64"`
	// cpu处理器型号
	// Extensions:
	// x-nullable: true
	CPUModel *string `json:"cpuModel" validate:"omitempty,min=1,max=64"`
	// cpu频率(G)
	// Extensions:
	// x-nullable: true
	CPUFrequency *string `json:"cpuFrequency" validate:"omitempty,min=1,max=64"`
	// 内存接口（如DDR3，DDR4）
	// Extensions:
	// x-nullable: true
	MemType *string `json:"memType" validate:"omitempty,oneof=DDR DDR2 DDR3 DDR4 DDR5"`
	// 单个内存大小(GB)
	// Extensions:
	// x-nullable: true
	MemSize *int `json:"memSize" validate:"omitempty,min=0,max=10000"`
	// 内存数量
	// Extensions:
	// x-nullable: true
	MemAmount *int `json:"memAmount" validate:"omitempty,min=0,max=10000"`
	// 内存主频（MHz)
	// Extensions:
	// x-nullable: true
	MemFrequency *int `json:"memFrequency" validate:"omitempty,min=0,max=10000"`
	// 网卡数量
	// Extensions:
	// x-nullable: true
	NicAmount *int `json:"nicAmount" validate:"omitempty,min=0,max=10000"`
	// 网卡传输速率(GE)
	// Extensions:
	// x-nullable: true
	NicRate *int `json:"nicRate" validate:"omitempty,min=0,max=10000"`
	// 【网口模式】【网络设置】: 单网口,双网口bond
	// Extensions:
	// x-nullable: true
	InterfaceMode *string `json:"interfaceMode" validate:"omitempty,oneof=bond dual"`
	// 系统盘类型（SSD，HDD）
	// Extensions:
	// x-nullable: true
	SystemVolumeType *string `json:"systemVolumeType" validate:"omitempty,oneof=SSD HDD"`
	// 系统盘接口类型（SATA,SAS,NVME）
	// Extensions:
	// x-nullable: true
	SystemVolumeInterfaceType *string `json:"systemVolumeInterfaceType" validate:"omitempty,oneof=SATA SAS NVME"`
	// 系统盘单盘大小
	// Extensions:
	// x-nullable: true
	SystemVolumeSize *int `json:"systemVolumeSize" validate:"omitempty,min=0,max=100000"`
	// 系统盘单盘大小单位
	// Extensions:
	// x-nullable: true
	SystemVolumeUnit *string `json:"systemVolumeUnit" validate:"omitempty,oneof=GB TB"`
	// 系统盘数量
	// Extensions:
	// x-nullable: true
	SystemVolumeAmount *int `json:"systemVolumeAmount" validate:"omitempty,min=0,max=10000"`
	// gpu数量
	// Extensions:
	// x-nullable: true
	GpuAmount *int `json:"gpuAmount" validate:"omitempty,min=0,max=10000"`
	// gpu厂商
	// Extensions:
	// x-nullable: true
	GpuManufacturer *string `json:"gpuManufacturer" validate:"omitempty,min=0,max=64"`
	// gpu处理器型号
	// Extensions:
	// x-nullable: true
	GpuModel *string `json:"gpuModel" validate:"omitempty,min=0,max=64"`
	// 数据盘类型
	// Extensions:
	// x-nullable: true
	DataVolumeType *string `json:"dataVolumeType" validate:"omitempty,oneof='' SSD HDD"`
	// 数据盘接口类型
	// Extensions:
	// x-nullable: true
	DataVolumeInterfaceType *string `json:"dataVolumeInterfaceType" validate:"omitempty,oneof='' SATA SAS NVME"`
	// 数据盘单盘大小
	// Extensions:
	// x-nullable: true
	DataVolumeSize *int `json:"dataVolumeSize" validate:"omitempty,min=0,max=100000"`
	// 数据盘单盘大小单位
	// Extensions:
	// x-nullable: true
	DataVolumeUnit *string `json:"dataVolumeUnit" validate:"omitempty,oneof='' GB TB"`
	// 数据盘数量
	// Extensions:
	// x-nullable: true
	DataVolumeAmount *int `json:"dataVolumeAmount" validate:"omitempty,min=0,max=10000"`
	// 系统盘raid，支持多选,英文逗号分隔
	// Extensions:
	// x-nullable: true
	RaidID *string `json:"raidId"`
	// CPU 预置规格
	// Extensions:
	// x-nullable: true
	CpuSpec *string `json:"cpuSpec" validate:"omitempty,oneof=common user_defined"`
	// 内存 预置规格
	// Extensions:
	// x-nullable: true
	MemSpec *string `json:"memSpec" validate:"omitempty,oneof=common user_defined"`
}

func (param *ModifyDeviceTypeRequest) Validate(logger *log.Logger) {
	validate(param, logger)
	//fmt.Println("开始校验")
	////对于必填项校验方法
	//fmt.Println(util.ObjToJson(param))
	//if param.Name != nil && (len(*param.Name) < 1 || len(*param.Name) > 64) {
	//	panic(constant.BuildInvalidArgumentWithArgs("机型名称不合法", "Name invalidate"))
	//}
	//if param.DeviceType != nil && (len(*param.DeviceType) < 1 || len(*param.DeviceType) > 64) {
	//	panic(constant.BuildInvalidArgumentWithArgs("DeviceType 不合法", "DeviceType invalidate"))
	//}
	//if param.DeviceSeries != nil && !util.InArray(*param.DeviceSeries, []string{"computer", "storage", "gpu", "other"}) {
	//	//fmt.Println(!util.InArray(param.DeviceSeries, []string{"computer", "storage", "gpu", "other"}))
	//	panic(constant.BuildInvalidArgumentWithArgs("DeviceSeries 不合法", "DeviceSeries invalidate"))
	//}
	//if param.Architecture != nil && (len(*param.Architecture) < 1 || len(*param.Architecture) > 64) {
	//	panic(constant.BuildInvalidArgumentWithArgs("Architecture 不合法", "Architecture invalidate"))
	//}
	//if param.Height != nil && !util.InArray(*param.Height, []int{2, 4}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("Height 不合法", "Height invalidate"))
	//}
	//
	//if param.CPUAmount != nil && (*param.CPUAmount < 1 || *param.CPUAmount > 10000) {
	//	panic(constant.BuildInvalidArgumentWithArgs("CPUAmount 不合法", "CPUAmount invalidate"))
	//}
	//if param.CPUCores != nil && (*param.CPUCores < 1 || *param.CPUCores > 10000) {
	//	panic(constant.BuildInvalidArgumentWithArgs("CPUCores 不合法", "CPUCores invalidate"))
	//}
	//if param.CPUManufacturer != nil && (len(*param.CPUManufacturer) < 1 || len(*param.CPUManufacturer) > 64) {
	//	panic(constant.BuildInvalidArgumentWithArgs("CPUManufacturer 不合法", "CPUManufacturer invalidate"))
	//}
	//if param.CPUModel != nil && (len(*param.CPUModel) < 1 || len(*param.CPUModel) > 64) {
	//	panic(constant.BuildInvalidArgumentWithArgs("CPUModel 不合法", "CPUModel invalidate"))
	//}
	//if param.CPUFrequency != nil && (len(*param.CPUFrequency) < 1 || len(*param.CPUFrequency) > 64) {
	//	panic(constant.BuildInvalidArgumentWithArgs("CPUFrequency 不合法", "CPUFrequency invalidate"))
	//}
	//
	//if param.MemType != nil && !util.InArray(*param.MemType, []string{"DDR", "DDR2", "DDR3", "DDR4", "DDR5"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("MemType 不合法", "MemType invalidate"))
	//}
	//if param.MemSize != nil && (*param.MemSize < 1 || *param.MemSize > 10000) {
	//	panic(constant.BuildInvalidArgumentWithArgs("MemSize 不合法", "MemSize invalidate"))
	//}
	//if param.MemAmount != nil && (*param.MemAmount < 1 || *param.MemAmount > 1000) {
	//	panic(constant.BuildInvalidArgumentWithArgs("MemAmount 不合法", "MemAmount invalidate"))
	//}
	//if param.MemFrequency != nil && (*param.MemFrequency < 1 || *param.MemFrequency > 10000) {
	//	panic(constant.BuildInvalidArgumentWithArgs("MemFrequency 不合法", "MemFrequency invalidate"))
	//}
	//if param.NicAmount != nil && (*param.NicAmount < 1 || *param.NicAmount > 10000) {
	//	panic(constant.BuildInvalidArgumentWithArgs("NicAmount 不合法", "NicAmount invalidate"))
	//}
	//if param.NicRate != nil && (*param.NicRate < 1 || *param.NicRate > 10000) {
	//	panic(constant.BuildInvalidArgumentWithArgs("NicRate 不合法", "NicRate invalidate"))
	//}
	//
	//if param.InterfaceMode != nil && !util.InArray(*param.InterfaceMode, []string{"bond", "dual"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("InterfaceMode 不合法", "InterfaceMode invalidate"))
	//}
	//
	//if param.SystemVolumeType != nil && !util.InArray(*param.SystemVolumeType, []string{"SSD", "HDD"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("SystemVolumeType 不合法", "SystemVolumeType invalidate"))
	//}
	//if param.SystemVolumeInterfaceType != nil && !util.InArray(*param.SystemVolumeInterfaceType, []string{"SATA", "SAS", "NVME"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("SystemVolumeInterfaceType 不合法", "SystemVolumeInterfaceType invalidate"))
	//}
	//if param.SystemVolumeSize != nil && (*param.SystemVolumeSize < 1 || *param.SystemVolumeSize > 100000) {
	//	panic(constant.BuildInvalidArgumentWithArgs("SystemVolumeSize 不合法", "SystemVolumeSize invalidate"))
	//}
	//if param.SystemVolumeAmount != nil && (*param.SystemVolumeAmount < 1 || *param.SystemVolumeAmount > 10000) {
	//	panic(constant.BuildInvalidArgumentWithArgs("SystemVolumeAmount 不合法", "SystemVolumeAmount invalidate"))
	//}
	//if param.SystemVolumeUnit != nil && !util.InArray(*param.SystemVolumeUnit, []string{"GB", "TB"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("SystemVolumeUnit 不合法", "SystemVolumeUnit invalidate"))
	//}
	//
	////非必填校验方法
	//if param.GpuManufacturer != nil && len(*param.GpuManufacturer) > 64 {
	//	panic(constant.BuildInvalidArgumentWithArgs("GpuManufacturer 不合法", "GpuManufacturer invalidate"))
	//}
	//if param.GpuModel != nil && len(*param.GpuModel) > 64 {
	//	panic(constant.BuildInvalidArgumentWithArgs("GpuModel 不合法", "GpuModel invalidate"))
	//}
	//if param.GpuAmount != nil && *param.GpuAmount > 10000 {
	//	panic(constant.BuildInvalidArgumentWithArgs("GpuAmount 不合法", "GpuAmount invalidate"))
	//}
	//if param.DataVolumeType != nil && !util.InArray(*param.DataVolumeType, []string{"", "SSD", "HDD"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("DataVolumeType 不合法", "DataVolumeType invalidate"))
	//}
	//if param.DataVolumeInterfaceType != nil && !util.InArray(*param.DataVolumeInterfaceType, []string{"", "SATA", "SAS", "NVME"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("DataVolumeInterfaceType 不合法", "DataVolumeInterfaceType invalidate"))
	//}
	//if param.DataVolumeSize != nil && *param.DataVolumeSize > 100000 {
	//	panic(constant.BuildInvalidArgumentWithArgs("DataVolumeSize 不合法", "DataVolumeSize invalidate"))
	//}
	//if param.DataVolumeAmount != nil && *param.DataVolumeAmount > 10000 {
	//	panic(constant.BuildInvalidArgumentWithArgs("DataVolumeAmount 不合法", "DataVolumeAmount invalidate"))
	//}
	//if param.DataVolumeUnit != nil && !util.InArray(*param.DataVolumeUnit, []string{"", "GB", "TB"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("DataVolumeUnit 不合法", "DataVolumeUnit invalidate"))
	//}
	//
	//其他逻辑校验

	if param.GpuManufacturer != nil && *param.GpuManufacturer == "" {
		if param.GpuModel != nil && *param.GpuModel != "" || param.GpuAmount != nil && *param.GpuAmount != 0 {
			panic(constant.BuildInvalidArgumentWithArgs("gpu不合法", "gpu invalidate"))
		}
	}
	if param.GpuModel != nil && *param.GpuModel == "" {
		if param.GpuManufacturer != nil && *param.GpuManufacturer != "" || param.GpuAmount != nil && *param.GpuAmount != 0 {
			panic(constant.BuildInvalidArgumentWithArgs("gpu不合法", "gpu invalidate"))
		}
	}
	if param.GpuAmount != nil && *param.GpuAmount == 0 {
		if param.GpuManufacturer != nil && *param.GpuManufacturer != "" || param.GpuModel != nil && *param.GpuModel != "" {
			panic(constant.BuildInvalidArgumentWithArgs("gpu不合法", "gpu invalidate"))
		}
	}
	if param.DataVolumeType != nil && *param.DataVolumeType == "" {
		if param.DataVolumeInterfaceType != nil && *param.DataVolumeInterfaceType != "" || param.DataVolumeAmount != nil && *param.DataVolumeAmount != 0 || param.DataVolumeSize != nil && *param.DataVolumeSize != 0 {
			panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
		}
	}
	if param.DataVolumeInterfaceType != nil && *param.DataVolumeInterfaceType == "" {
		if param.DataVolumeType != nil && *param.DataVolumeType != "" || param.DataVolumeAmount != nil && *param.DataVolumeAmount != 0 || param.DataVolumeSize != nil && *param.DataVolumeSize != 0 {
			panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
		}
	}
	if param.DataVolumeAmount != nil && *param.DataVolumeAmount == 0 {
		if param.DataVolumeInterfaceType != nil && *param.DataVolumeInterfaceType != "" || param.DataVolumeType != nil && *param.DataVolumeType != "" || param.DataVolumeSize != nil && *param.DataVolumeSize != 0 {
			panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
		}
	}
	if param.DataVolumeSize != nil && *param.DataVolumeSize == 0 {
		if param.DataVolumeInterfaceType != nil && *param.DataVolumeInterfaceType != "" || param.DataVolumeAmount != nil && *param.DataVolumeAmount != 0 || param.DataVolumeType != nil && *param.DataVolumeType != "" {
			panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
		}
	}
}

type QueryDeviceTypesRequest struct {
	// 机型id
	DeviceTypeID string `json:"deviceTypeId"`
	// 机型规格
	DeviceType string `json:"deviceType"`
	// 机房id
	IdcID string `json:"idcId"`
	// 机型名称
	Name string `json:"name"`
	// 机型类型
	DeviceSeries string `json:"deviceSeries"`
	Pageable
	IsAll string `json:"isAll"`
}

func (req *QueryDeviceTypesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type QueryDeviceTypeImageRequest struct {
	// 机型id
	DeviceTypeID string `json:"deviceTypeId"`
	// 镜像id
	ImageID string `json:"imageId"`
	// 体系架构
	Architecture string `json:"architecture"`
	// 操作系统平台
	OsType string `json:"osType"`
	// 镜像名称
	ImageName string `json:"imageName"`
	// 版本号
	Version string `json:"version"`
	// 操作系统ID
	OsID string `json:"osId"`
	// 镜像类型，预置，自定义
	Source string `json:"source"`
	Pageable
	// 是否显示全部
	IsAll string `json:"isAll"`
}

func (req *QueryDeviceTypeImageRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type QueryVolumesRaidsRequest struct {
	// 机型id
	DeviceTypeID string `json:"deviceTypeId"`
	// 系统盘还是数据盘
	VolumeType string `json:"volumeType"`
	// 是否显示所有
	IsAll string `json:"isAll"`
}

func (req *QueryVolumesRaidsRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type QueryDeviceTypeImagePartitionRequest struct {
	// 机型id
	DeviceTypeID string `json:"deviceTypeId"` //validate:"required"
	// 镜像id
	ImageID string `json:"imageId"` //validate:"required"
	// 是否显示所有，isAll=1显示所有
	IsAll string `json:"isAll"`
}

func (req *QueryDeviceTypeImagePartitionRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
