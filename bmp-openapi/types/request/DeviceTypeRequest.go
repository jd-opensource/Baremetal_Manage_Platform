package request

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/raidDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

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
	Architecture string `json:"architecture" validate:"required,oneof=x86_64 ARM64(aarch64) i386 LoongArch"`
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
	NicAmount int `json:"nicAmount" validate:"required,min=1,max=2"`
	// 【网口模式】【网络设置】: bond,dual,single
	// required: true
	InterfaceMode string `json:"interfaceMode" validate:"required,oneof=bond dual single"`
	// 网卡传输速率(GE) 1-10000
	// required: true
	NicRate int `json:"nicRate" validate:"required,min=1,max=10000"`
	// gpu数量0-1000
	GpuAmount int `json:"gpuAmount" validate:"min=0,max=1000"`
	// gpu品牌0-64
	GpuManufacturer string `json:"gpuManufacturer" validate:"min=0,max=64"`
	// gpu型号0-64
	GpuModel string `json:"gpuModel" validate:"min=0,max=64"`
	// CPU 预置规格
	// required: true
	CpuSpec string `json:"cpuSpec" validate:"required,oneof=common user_defined"`
	// 内存 预置规格
	// required: true
	MemSpec string `json:"memSpec" validate:"required,oneof=common user_defined"`
	//boot模式【UEFI Legacy/BIOS】支持多选，逗号分隔
	BootMode string `json:"boot_mode" validate:"omitempty"`
	// 卷管理
	// required: true
	Volumes []VolumeItem `json:"volumes"`
	// 阵列卡是否需要配置 1 需要配置 2无需配置
	IsNeedRaid string `json:"isNeedRaid" validate:"required,oneof=need_raid no_need_raid"`
}
type VolumeItem struct {
	DeviceTypeID  string `json:"deviceTypeId"`
	VolumeName    string `json:"volumeName"`
	VolumeType    string `json:"volumeType"`
	DiskType      string `json:"diskType"`
	InterfaceType string `json:"interfaceType"`
	VolumeSize    string `json:"volumeSize"`
	VolumeUnit    string `json:"volumeUnit"`
	VolumeAmount  int    `json:"volumeAmount"`
	RaidCan       string `json:"raidCan"`
	Raid          string `json:"raid"`
}

func (param *CreateDeviceTypeRequest) Validate(logger *log.Logger) {
	validate(param, logger)
	if param.NicAmount == 1 && param.InterfaceMode == "bond" {
		panic(constant.BuildInvalidArgumentWithArgs("nicAmount和interfaceMode参数错误", "nicAmount not match interfaceMode"))
	}
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
	if param.BootMode != "" {
		items := strings.Split(param.BootMode, ",")
		for _, v := range items {
			b := strings.TrimSpace(v)
			if b != "UEFI" && b != "Legacy/BIOS" {
				panic(constant.BuildInvalidArgumentWithArgs("boot_mode 非法", "boot_mode invalid"))
			}
		}
	}
	ValidVolumes(logger, param.Volumes) //校验卷相关逻辑
}
func ValidVolumes(logger *log.Logger, volumes []VolumeItem) {
	//raidMap := map[string]response.Raid{}
	//raidList,_ := raidLogic.QueryRaidsAll(logger)
	//for _,v := range raidList{
	//	raidMap[v.RaidID] = *v
	//}
	count := 0
	for _, v := range volumes {
		specialMatch := regexp.MustCompile("^[a-zA-Z0-9_-]{1,64}$").MatchString
		if !specialMatch(v.VolumeName) {
			panic(constant.BuildInvalidArgumentWithArgs("卷名称:"+v.VolumeName+",不合法", "volumeName invalid"))
		}
		if v.VolumeType == baseLogic.VOLUME_TYPE_SYSTEM {
			count++
		}
		if v.RaidCan == baseLogic.RAID_CAN_NO_RAID {
			if v.InterfaceType != "NVME" {
				panic(constant.BuildInvalidArgumentWithArgs("RAID配置为NO RAID时，接口类型必须是NVME", "no raid,interfaceType must be NVME"))
			}
			if v.VolumeAmount != 1 {
				logger.Warn("NVME只能有1块盘")
				panic(constant.BuildInvalidArgumentWithArgs("仅支持1块盘", "Only support 1 disk"))
			}
		}
		if v.RaidCan == baseLogic.RAID_CAN_SINGLE_RAID_EN {
			if v.VolumeAmount != 1 {
				logger.Warn("RAID配置为单盘RAID0时，只能有1块盘")
				panic(constant.BuildInvalidArgumentWithArgs("仅支持1块盘", "Only support 1 disk"))
			}
		}
		raidList := strings.Split(v.Raid, ",")
		raidAmountMap := map[string]int{
			"RAID0":  2,
			"RAID1":  2,
			"RAID5":  3,
			"RAID10": 4,
		}
		if len(raidList) == 1 {
			raidInfo, _ := raidDao.GetRaidById(logger, v.Raid)
			if raidInfo != nil {
				if v.RaidCan == baseLogic.RAID_CAN_RAID {
					if v.VolumeAmount < raidAmountMap[raidInfo.Name] {
						logger.Warn(fmt.Sprintf("RAID配置为RAID，RAID模式是%s，最低%d块盘", raidInfo.Name, raidAmountMap[raidInfo.Name]))
						panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("最少%d块盘，数量是所选RAID需要的最大盘数", raidAmountMap[raidInfo.Name]), fmt.Sprintf("At least %d disks", raidAmountMap[raidInfo.Name])))
					}
				}
				//其他模式todo
			}
		} else if len(raidList) > 1 {
			amount := []int{0}
			for _, raidId := range raidList {
				raidInfo, _ := raidDao.GetRaidById(logger, raidId)
				if raidInfo != nil {
					amount = append(amount, raidAmountMap[raidInfo.Name])
					//其他模式todo
				}
			}
			fmt.Println(amount)
			sort.Ints(amount)
			max := amount[len(amount)-1]
			if v.VolumeAmount < max {
				logger.Warn("多种raid模式，最低数量为这几种模式允许的最大值" + strconv.Itoa(max))
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("最少%d块盘，数量是所选RAID需要的最大盘数", max), fmt.Sprintf("At least %d disks", max)))
			}
		}

	}
	if count != 1 {
		panic(constant.BuildInvalidArgumentWithArgs("有且只有一个系统卷", "only one system volume"))
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
	Architecture *string `json:"architecture" validate:"omitempty,oneof=x86_64 'ARM64(aarch64)' i386 LoongArch"`
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
	NicAmount *int `json:"nicAmount" validate:"omitempty,min=0,max=2"`
	// 【网口模式】【网络设置】: 单网口,双网口bond，single
	// Extensions:
	// x-nullable: true
	InterfaceMode *string `json:"interfaceMode" validate:"omitempty,oneof=bond dual single"`
	// 网卡传输速率(GE)
	// Extensions:
	// x-nullable: true
	NicRate *int `json:"nicRate" validate:"omitempty,min=0,max=10000"`
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
	SystemVolumeAmount *int `json:"systemVolumeAmount" validate:"omitempty,min=1,max=2"`
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
	//boot模式【UEFI Legacy/BIOS】支持多选，逗号分隔
	BootMode string `json:"boot_mode" validate:"omitempty"`
	// 卷管理
	// Extensions:
	// x-nullable: true
	Volumes []VolumeItem `json:"volumes"`
	// 阵列卡是否需要配置 1 需要配置 2无需配置
	// Extensions:
	// x-nullable: true
	IsNeedRaid *string `json:"isNeedRaid" validate:"omitempty,oneof=need_raid no_need_raid"`
}

func (param *ModifyDeviceTypeRequest) Validate(logger *log.Logger) {
	validate(param, logger)

	if (param.NicAmount != nil && *param.NicAmount == 1) && (param.InterfaceMode != nil && *param.InterfaceMode == "bond") {
		panic(constant.BuildInvalidArgumentWithArgs("nicAmount和interfaceMode参数错误", "nicAmount not match interfaceMode"))
	}

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

	if param.BootMode != "" {
		items := strings.Split(param.BootMode, ",")
		for _, v := range items {
			b := strings.TrimSpace(v)
			if b != "UEFI" && b != "Legacy/BIOS" {
				panic(constant.BuildInvalidArgumentWithArgs("boot_mode 非法", "boot_mode invalid"))
			}
		}
	}
	ValidVolumes(logger, param.Volumes) //校验卷相关逻辑
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

type QueryDeviceTypeRaidRequest struct {
	// 机型id
	DeviceTypeID string `json:"deviceTypeId"`
	// 系统盘还是数据盘
	VolumeType string `json:"volumeType"`
	// raidID
	RaidID string `json:"raidId"`
	// 是否显示所有
	IsAll string `json:"isAll"`
}

func (req *QueryDeviceTypeRaidRequest) Validate(logger *log.Logger) {
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

type DescribeVolumesByDeviceTypeRequest struct {
	// 机型id
	DeviceTypeID string `json:"deviceTypeId" validate:"required"` //
}

func (req *DescribeVolumesByDeviceTypeRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
