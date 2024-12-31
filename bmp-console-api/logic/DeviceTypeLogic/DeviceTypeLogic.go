package DeviceTypeLogic

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-console-api/constants"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	sdkDeviceTypeParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/device_type"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

func sdkDeviceType2DeviceType(logger *log.Logger, v *sdkModels.DeviceType) responseTypes.DeviceType {
	language := logger.GetPoint("language").(string)

	/*
		CPU：2* 24物理核，2.6GHz

		内存：384GB，2933MHz

		数据盘：2* 480 SSD

		系统盘：2* 2TB NVME

		网卡：2* 25GE

		GPU：4* NVIDIA T4
	*/
	//var cputpl string = "%d*%d 物理核,%s" //
	//var memtpl string = "%d*%d, %d MHZ"
	//var systemtpl string = "%d*%d %s"
	//var datatpl string = "%d*%d %s"
	//var nictpl string = "%d*%d GE"
	//var gputpl string = "%d*%s %s"
	//if language == constants.LANGUAGE_EN {
	//	cputpl = "%d*%d cores,%s"
	//}

	//raidMap := make(map[string]responseTypes.Raid) //	请求一次，避免循环请求

	//raid, _ := raidLogic.QueryRaids(logger, requestTypes.QueryRaidsRequest{
	//	DeviceTypeID: v.DeviceTypeID,
	//})
	//raidStr := ""
	//for _, value := range raid {
	//	raidStr += value.RaidName + ","
	//}
	//raidStr = strings.Trim(raidStr, ",")
	cpu_ := strconv.Itoa(int(v.CPUAmount)) + "*" + strconv.Itoa(int(v.CPUCores)) + "物理核," + v.CPUFrequency + "GHz"
	if language == constants.LANGUAGE_EN {
		cpu_ = strconv.Itoa(int(v.CPUAmount)) + "*" + strconv.Itoa(int(v.CPUCores)) + "cores," + v.CPUFrequency + "GHz"
	}
	mem_ := strconv.Itoa(int(v.MemAmount)*int(v.MemSize)) + "GB " + v.MemType + " " + strconv.Itoa(int(v.MemFrequency)) + "MHz"
	// 系统盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式）
	// 240GB 240GB SATA SSD*2 RAID1
	//sv_ := strconv.Itoa(int(v.SystemVolumeAmount)*int(v.SystemVolumeSize)) + "GB (" + strconv.Itoa(int(v.SystemVolumeSize)) + "GB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount)) + "," + raidStr + ")"
	//if v.SystemVolumeUnit == "TB" {
	//	sv_ = strconv.FormatFloat(float64(v.SystemVolumeAmount)*float64(v.SystemVolumeSize)/1000, 'f', 2, 64) + "TB (" + strconv.FormatFloat(float64(v.SystemVolumeSize)/1000, 'f', 2, 64) + "TB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount)) + "," + raidStr + ")"
	//}
	//// 数据盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式）
	//dv_ := strconv.Itoa(int(v.DataVolumeAmount)*int(v.DataVolumeSize)) + "GB (" + strconv.Itoa(int(v.DataVolumeSize)) + "GB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount)) + "," + raidStr + ")"
	//if v.DataVolumeUnit == "TB" {
	//	dv_ = strconv.FormatFloat(float64(v.DataVolumeAmount)*float64(v.DataVolumeSize)/1000, 'f', 2, 64) + "TB (" + strconv.FormatFloat(float64(v.DataVolumeSize)/1000, 'f', 2, 64) + "TB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount)) + "," + raidStr + ")"
	//}
	//if v.DataVolumeAmount == 0 {
	//	dv_ = ""
	//}

	// 显示品牌 显卡型号 个数，例 NVIDIA P40 * 4
	gpu_ := v.GpuManufacturer + " " + v.GpuModel + "*" + strconv.Itoa(int(v.GpuAmount))
	if v.GpuAmount == 0 {
		gpu_ = ""
	}
	nicInfo := strconv.Itoa(int(v.NicAmount)) + "*" + strconv.Itoa(int(v.NicRate)) + "GE"
	sv_, dv_ := Volumes2SystemDataInfo(logger, v.Volumes)
	return responseTypes.DeviceType{
		Name:         v.Name,
		DeviceTypeId: v.DeviceTypeID,
		//Cpu:              fmt.Sprintf(cputpl, v.CPUAmount, v.CPUCores, v.CPUFrequency),
		//Mem:              fmt.Sprintf(memtpl, v.MemAmount, v.MemSize, v.MemFrequency),
		//System:           fmt.Sprintf(systemtpl, v.SystemVolumeAmount, v.SystemVolumeSize, v.SystemVolumeType),
		//Data:             fmt.Sprintf(datatpl, v.DataVolumeAmount, v.DataVolumeSize, v.DataVolumeType),
		//Nic:              fmt.Sprintf(nictpl, v.NicAmount, v.NicRate),
		//Gpu:              fmt.Sprintf(gputpl, v.GpuAmount, v.GpuManufacturer, v.GpuModel),
		Cpu:              cpu_,
		Mem:              mem_,
		System:           sv_,
		Data:             dv_,
		Nic:              nicInfo,
		Gpu:              gpu_,
		AvailableStock:   v.StockAvailable,
		DeviceSeries:     v.DeviceSeries,
		DeviceSeriesName: v.DeviceSeriesName,
		BootMode:         v.BootMode,
		IsNeedRaid:       v.IsNeedRaid,
	}
}

func Volumes2SystemDataInfo(logger *log.Logger, volumes []*sdkModels.Volume) (string, string) {
	var sv_, dv_ string
	ks := map[string]int{}
	for _, v := range volumes {
		if v.VolumeType == "system" { //系统卷只有一个
			sv_ = v.VolumeSize + v.VolumeUnit + "*" + strconv.Itoa(int(v.VolumeAmount)) + " " + v.InterfaceType + " " + v.DiskType
		} else {
			key := fmt.Sprintf("%s___%s___%s___%s", v.VolumeSize, v.VolumeUnit, v.InterfaceType, v.DiskType)
			ks[key] = 1
		}
	}
	for k, v := range ks {
		items := strings.Split(k, "___")
		dv_ += items[0] + items[1] + "*" + strconv.Itoa(v) + " " + items[2] + " " + items[3] + ";"
	}
	dv_ = strings.TrimRight(dv_, ";")

	return sv_, dv_
}

func GetAvailableDeviceTypes(logger *log.Logger, param requestTypes.GetAvailableDeviceTypesRequest) (*responseTypes.AvailableDeviceType, error) {

	logid := logger.GetPoint("logid").(string)
	userId := logger.GetPoint("userId").(string)
	req := sdkDeviceTypeParameters.NewDescribeDeviceTypesParams()
	req.WithTraceID(logid)
	req.WithIdcID(&param.IdcID)
	req.WithBmpUserID(&userId)
	isAll := "1"
	req.WithIsAll(&isAll)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("QueryDeviceTypes openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.DeviceType.DescribeDeviceTypes(req, nil)
	logger.Info("QueryDeviceTypes openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("QueryDeviceTypes openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	deviceTypes := responseApi.Payload.Result.DeviceTypes
	res := map[string][]responseTypes.DeviceType{}
	for _, v := range deviceTypes {
		if _, ok := res[v.DeviceSeries]; !ok {
			res[v.DeviceSeries] = []responseTypes.DeviceType{sdkDeviceType2DeviceType(logger, v)}
		} else {
			res[v.DeviceSeries] = append(res[v.DeviceSeries], sdkDeviceType2DeviceType(logger, v))
		}
	}
	obj := &responseTypes.AvailableDeviceType{}
	if len(res["computer"]) != 0 {
		obj.Computer = res["computer"]
	}
	if len(res["storage"]) != 0 {
		obj.Storage = res["storage"]
	}
	if len(res["gpu"]) != 0 {
		obj.Gpu = res["gpu"]
	}
	if len(res["other"]) != 0 {
		obj.Other = res["other"]
	}
	fmt.Println(obj)
	return obj, nil

}
func Trim(str string) string {
	s := strings.TrimRight(str, "0")
	s = strings.TrimRight(s, ".")
	return s
}

func GetDeviceTypeByUuid(logger *log.Logger, deviceTypeID string) (*sdkModels.DeviceType, error) {

	logid := logger.GetPoint("logid").(string)
	userId := logger.GetPoint("userId").(string)
	req := sdkDeviceTypeParameters.NewDescribeDeviceTypeParams()
	req.WithTraceID(logid)
	req.WithDeviceTypeID(deviceTypeID)
	req.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("DescribeDeviceType openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.DeviceType.DescribeDeviceType(req, nil)
	logger.Info("DescribeDeviceType openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeDeviceType openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	deviceTypes := responseApi.Payload.Result.DeviceType
	return deviceTypes, nil
}
