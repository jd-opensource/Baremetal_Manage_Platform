package response

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/volumeDao"
)

type DeviceType struct {
	ID int `json:"id"`
	// 机房uuid
	IDcID string `json:"idcId"`
	// 机房名称
	IDcName string `json:"idcName"`
	// 机房名称
	IDcNameEn string `json:"idcNameEn"`
	// 机型uuid
	DeviceTypeID string `json:"deviceTypeId"`
	// 机型名称，如计算效能型,标准计算型
	Name string `json:"name"`
	// 机型规格, cps.c.normal
	DeviceType string `json:"deviceType"`
	// 机型类型，如计算型，存储型
	DeviceSeries string `json:"deviceSeries"`
	// 机型类型，如计算型，存储型
	DeviceSeriesName string `json:"deviceSeriesName"`
	// 体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64
	Architecture string `json:"architecture"`
	// 显示机型高度
	Height int `json:"height"`
	// 描述
	Description string `json:"description"`
	// cpu数量
	CPUAmount int `json:"cpuAmount"`
	// 单个cpu内核数
	CPUCores int `json:"cpuCores"`
	// cpu厂商
	CPUManufacturer string `json:"cpuManufacturer"`
	// cpu处理器型号
	CPUModel string `json:"cpuModel"`
	// cpu频率(G)
	CPUFrequency string `json:"cpuFrequency"`
	// 内存接口（如DDR3，DDR4）
	MemType string `json:"memType"`
	// 单个内存大小(GB)
	MemSize int `json:"memSize"`
	// 内存数量
	MemAmount int `json:"memAmount"`
	// 内存主频（MHz)
	MemFrequency int `json:"memFrequency"`
	// 网卡数量
	NicAmount int `json:"nicAmount"`
	// 网卡传输速率(GE)
	NicRate int `json:"nicRate"`
	// bond单网口,dual双网口
	InterfaceMode string `json:"interfaceMode"`
	// gpu数量
	GpuAmount int `json:"gpuAmount"`
	// gpu厂商
	GpuManufacturer string `json:"gpuManufacturer"`
	// gpu处理器型号
	GpuModel string `json:"gpuModel"`
	// 机型的卷信息
	Volumes []*volumeDao.Volume `json:"volumes"`
	// 创建者
	CreatedBy string `json:"createdBy"`
	// 更新者
	UpdatedBy string `json:"updatedBy"`
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// 更新时间
	UpdatedTime string `json:"updatedTime"`
	// StockAvailable 可用库存
	StockAvailable int `json:"stockAvailable"`
	// raid配置
	RaidCan string `json:"raidCan"`
	//系统盘raid
	Raid string `json:"raid"`
	// CPU 规格,预置还是其它
	CpuSpec string `json:"cpuSpec"`
	// 内存 规格,预置还是其它
	MemSpec string `json:"memSpec"`

	//拼装信息
	CpuInfo string `json:"cpuInfo"`
	MemInfo string `json:"memInfo"`
	SvInfo  string `json:"svInfo"`
	DvInfo  string `json:"dvInfo"`
	GpuInfo string `json:"gpuInfo"`
	NicInfo string `json:"nicInfo"`

	InstanceCount  int      `json:"instanceCount"`
	DeviceCount    int      `json:"deviceCount"`
	InstanceStatus []string `json:"instanceStatus"`
	//boot模式
	BootMode string `json:"boot_mode"`
	// 是否配置了阵列卡
	IsNeedRaid string `json:"isNeedRaid"`
}

type DeviceTypeList struct {
	// 机型列表
	DeviceTypes []*DeviceType `json:"deviceTypes"`
	PageNumber  int64         `json:"pageNumber"`
	PageSize    int64         `json:"pageSize"`
	TotalCount  int64         `json:"totalCount"`
}
type DeviceTypeInfo struct {
	// 机型信息
	DeviceType *DeviceType `json:"deviceType"`
}
type DeviceTypeId struct {
	// 机型uuid
	// required: true
	DeviceTypeId string `json:"deviceTypeId"`
}

func (d DeviceType) MarshalJSON() ([]byte, error) {
	type Alias DeviceType
	return json.Marshal(struct {
		Alias
		//CreateTime string `json:"create_time"`
		//UpdateTime string `json:"update_time"`
	}{
		Alias: Alias(d),
		//CreateTime: d.CreateTime.Format("2006-01-02 15:04:05"),
		//UpdateTime: d.UpdateTime.Format("2006-01-02 15:04:05"),
	})
}
