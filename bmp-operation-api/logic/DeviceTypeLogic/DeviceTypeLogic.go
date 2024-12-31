package DeviceTypeLogic

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	sdkDeviceTypeParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/device_type"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"
	raidLogic "coding.jd.com/aidc-bmp/bmp-operation-api/logic/RaidLogic"
	"coding.jd.com/aidc-bmp/bmp-operation-api/services/excel"
	sdk "coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
)

// 【机型管理】【添加机型】
func AddDeviceType(logger *log.Logger, param *requestTypes.CreateDeviceTypeRequest) (*response.CreateDeviceTypeResult, error) {
	logid := logger.GetPoint("logid").(string)

	req := sdkDeviceTypeParameters.NewCreateDeviceTypeParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	//systemVolumeSize := int64(param.SystemVolumeSize)
	//if param.SystemVolumeUnit == "TB" {
	//	systemVolumeSize = int64(param.SystemVolumeSize * 1000)
	//}
	//dataVolumeSize := int64(param.DataVolumeSize)
	//if param.DataVolumeUnit == "TB" {
	//	dataVolumeSize = int64(param.DataVolumeSize * 1000)
	//}
	//volumeItem := sdkModels.CreateDeviceTypeRequestVolumesItems0{
	//	DeviceTypeID:
	//}
	body := sdkModels.CreateDeviceTypeRequest{
		Architecture:    &param.Architecture,
		CPUAmount:       &param.CPUAmount,
		CPUCores:        &param.CPUCores,
		CPUFrequency:    &param.CPUFrequency,
		CPUManufacturer: &param.CPUManufacturer,
		CPUModel:        &param.CPUModel,

		Description:     param.Description,
		DeviceSeries:    &param.DeviceSeries,
		DeviceType:      &param.DeviceType,
		GpuAmount:       param.GpuAmount,
		GpuManufacturer: param.GpuManufacturer,
		GpuModel:        param.GpuModel,
		Height:          GetInt64(param.Height),
		IDcID:           &param.IDcID,
		InterfaceMode:   &param.InterfaceMode,
		MemAmount:       GetInt64(param.MemAmount),
		MemFrequency:    GetInt64(param.MemFrequency),
		MemSize:         GetInt64(param.MemSize),
		MemType:         &param.MemType,
		Name:            &param.Name,
		NicAmount:       &param.NicAmount,
		NicRate:         GetInt64(param.NicRate),

		CPUSpec:    &param.CpuSpec,
		MemSpec:    &param.MemSpec,
		BootMode:   param.BootMode,
		Volumes:    param.Volumes,
		IsNeedRaid: param.IsNeedRaid,
	}
	req.WithBody(&body)
	logger.Info("CreateDeviceType openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.DeviceType.CreateDeviceType(req, nil)
	logger.Info("CreateDeviceType openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("AddDeviceType openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	if client == nil || client.Payload == nil || client.Payload.Result == nil {
		return nil, errors.New("add device info failed")
	}

	res := &response.CreateDeviceTypeResult{
		DeviceTypeId: *client.Payload.Result.DeviceTypeID,
	}

	return res, nil
}
func GetInt64(number int) *int64 {
	num := int64(number)
	return &num
}
func f2i(f float64) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", f))
	return i
}
func GetInt8(number int8) *int8 {
	return &number
}

// 【机型管理】【机型列表】
func GetDeviceTypeList(logger *log.Logger, param *requestTypes.QueryDeviceTypesRequest) (*response.DeviceTypePage, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceTypeParameters.NewDescribeDeviceTypesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithDeviceSeries(&param.DeviceSeries)
	req.WithDeviceType(&param.DeviceType)
	req.WithDeviceTypeID(&param.DeviceTypeID)
	req.WithIdcID(&param.IdcID)
	req.WithIsAll(&param.IsAll)
	req.WithName(&param.Name)
	req.WithPageNumber(&param.PageNumber)
	req.WithPageSize(&param.PageSize)
	logger.Info("DescribeDeviceTypes openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.DeviceType.DescribeDeviceTypes(req, nil)
	logger.Info("DescribeDeviceTypes openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil || client.Payload == nil {
		logger.Info("DescribeDeviceTypeList openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	resp := response.DeviceTypePage{}
	obj := response.DeviceType{}
	if len(client.Payload.Result.DeviceTypes) == 0 {
		resp.DeviceTypes = make([]response.DeviceType, 0)
		return &resp, nil
	}

	raidMap := make(map[string]response.Raid) //	请求一次，避免循环请求
	raid, raidsErr := raidLogic.QueryRaidsAll(logger)
	if raidsErr == nil && len(raid.Raids) != 0 {
		for _, v := range raid.Raids {
			raidMap[v.RaidID] = *v
		}
	}

	for _, v := range client.Payload.Result.DeviceTypes {
		//idc, err := idcLogic.GetIdcInfo(logger, idcReq)
		//if err == nil {
		//	idcName = idc.Name
		//	//idcRegion = idc.NameEn
		//}
		imageCount := ""
		imagesReq := &requestTypes.QueryDeviceTypeImagePageRequest{
			DeviceTypeID: v.DeviceTypeID,
		}
		images, imagesErr := GetDeviceTypeImageList(logger, imagesReq)
		if imagesErr == nil {
			imageCount = strconv.Itoa(len(images.DeviceTypeImages))
		}

		obj = buildEntity(logger, v)
		obj.ImageCount = imageCount
		//array := []string{} // 拆分查询
		//if len(v.Raid) != 0 {
		//	array = strings.Split(v.Raid, ",")
		//}
		//if len(array) != 0 {
		//	raid := ""
		//	for _, str := range array {
		//		raid = raid + raidMap[str].Name + "," // RAID0,RAID1
		//	}
		//	obj.Raid = raid[0 : len(raid)-1]
		//}

		resp.DeviceTypes = append(resp.DeviceTypes, obj)
	}
	resp.PageNumber = client.Payload.Result.PageNumber
	resp.PageSize = client.Payload.Result.PageSize
	resp.TotalCount = client.Payload.Result.TotalCount
	return &resp, nil
}

// 导出Excel文件
func ExportDeviceTypeExcel(logger *log.Logger, result []response.DeviceType) (string, string, error) {
	language := logger.GetPoint("language").(string)
	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)
	for _, v := range result {

		//obj := response.DeviceType{}
		//raidMap := make(map[string]response.Raid) //	请求一次，避免循环请求
		//raid, raidsErr := raidLogic.QueryRaidsAll(logger)
		//if raidsErr == nil && len(raid.Raids) != 0 {
		//	for _, v := range raid.Raids {
		//		raidMap[v.RaidID] = *v
		//	}
		//}
		//array := []string{} // 拆分查询
		//if len(v.Raid) != 0 {
		//	array = strings.Split(v.Raid, ",")
		//}
		//if len(array) != 0 {
		//	raid_ := ""
		//	for _, str := range array {
		//		raid_ = raid_ + raidMap[str].Name + "," // RAID0,RAID1
		//	}
		//	obj.Raid = raid_[0 : len(raid_)-1]
		//}

		//fmt.Println("1", v.Raid, "2", obj.Raid)
		//os.Exit(1)
		//cpu_ := v.v.CPUManufacturer + " " + v.CPUModel + "(" + strconv.Itoa(int(v.CPUAmount)) + "*" + strconv.Itoa(int(v.CPUCores)) + " 物理核, " + v.CPUFrequency + "GHz)"
		//mem_ := strconv.Itoa(int(v.MemAmount)*int(v.MemSize)) + "GB " + v.MemType + " " + strconv.Itoa(int(v.MemFrequency)) + "MHz"
		//// 系统盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式)
		//// 240GB 240GB SATA SSD*2 RAID1
		//sv_ := strconv.Itoa(int(v.SystemVolumeAmount)*int(v.SystemVolumeSize)) + "GB(" + strconv.Itoa(int(v.SystemVolumeSize)) + "GB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount)) + "," + v.Raid + ")"
		//if v.SystemVolumeUnit == "TB" {
		//	sv_ = strconv.FormatFloat(float64(v.SystemVolumeAmount)*float64(v.SystemVolumeSize), 'f', 2, 64) + "TB(" + strconv.FormatFloat(float64(v.SystemVolumeSize), 'f', 2, 64) + "TB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount)) + "," + v.Raid + ")"
		//}
		//// 数据盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式)
		//dv_ := strconv.Itoa(int(v.DataVolumeAmount)*int(v.DataVolumeSize)) + "GB(" + strconv.Itoa(int(v.DataVolumeSize)) + "GB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount)) + "," + v.Raid + ")"
		//if v.DataVolumeUnit == "TB" {
		//	dv_ = strconv.FormatFloat(float64(v.DataVolumeAmount)*float64(v.DataVolumeSize), 'f', 2, 64) + "TB(" + strconv.FormatFloat(float64(v.DataVolumeSize), 'f', 2, 64) + "TB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount)) + "," + v.Raid + ")"
		//}
		//if v.DataVolumeAmount == 0 {
		//	dv_ = ""
		//}
		//// 显示品牌 显卡型号 个数，例 NVIDIA P40 * 4
		//gpu_ := v.GpuManufacturer + " " + v.GpuModel + "*" + strconv.Itoa(int(v.GpuAmount))
		//if v.GpuAmount == 0 {
		//	gpu_ = ""
		//}
		//nicInfo := strconv.Itoa(int(v.NicAmount)) + "*" + strconv.Itoa(int(v.NicRate)) + "GE"

		nic_ := BaseLogic.InterfaceMode[v.InterfaceMode]
		if language == BaseLogic.EN_US {
			nic_ = BaseLogic.InterfaceModeEn[v.InterfaceMode]
		}

		imageCount := ""
		imagesReq := &requestTypes.QueryDeviceTypeImagePageRequest{
			DeviceTypeID: v.DeviceTypeID,
		}
		images, imagesErr := GetDeviceTypeImageList(logger, imagesReq)
		if imagesErr == nil {
			imageCount = strconv.Itoa(len(images.DeviceTypeImages))
		}

		deviceCount := "--"
		if v.DeviceCount > 0 {
			deviceCount = strconv.Itoa(v.DeviceCount)
		}

		sheetData = append(sheetData, []string{
			v.Name,
			v.DeviceSeriesName, // 机型类型：显示所属的机型类型
			v.IdcName,          // 机房名称：显示所属的机房
			v.IdcNameEn,        // 机房英文名称： cn-north-1
			v.Architecture,     // 体系架构：显示机型的体系架构
			v.BootMode,
			v.DeviceType, // 机型规格

			imageCount, // 暂时默认0 支持镜像数量
			deviceCount,
			v.CpuInfo, // CPU
			v.MemInfo, // 内存
			//v.SvInfo,  // 系统盘
			//v.RaidCan,
			//v.Raid,   //  系统盘RAID
			//v.DvInfo, // 数据盘
			//"RAID1", // TODO 缺少此字段 数据盘 RAID
			v.NicInfo, // 网卡
			nic_,      // 网络设置
			v.GpuInfo, // GPU

			v.Description, // 描述：显示针对机型的描述

		})
	}
	exportExcelHeader := BaseLogic.ExportDeviceTypeHeader
	if language == BaseLogic.EN_US {
		exportExcelHeader = BaseLogic.ExportDeviceTypeHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportExcelHeader, sheetData); err != nil {
		logger.Warn("create excel error:", err.Error())
		return "", "", err
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	// 生成100000到100000+900000之间的随机数，左闭右开
	downloadFileName := "device_type_list_" + time.Now().Format(BaseLogic.FormatDate) + "_" + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx"
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Warn("保存excel到本地失败" + err.Error())
		return "", "", err
	}
	return fileName, downloadFileName, nil
}

// 【机型管理】【机型详情信息】
func GetDeviceTypeInfo(logger *log.Logger, param *requestTypes.QueryDeviceTypeRequest) (*response.DeviceType, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceTypeParameters.NewDescribeDeviceTypeParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithDeviceTypeID(param.DeviceTypeID)
	logger.Info("DescribeDeviceType openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.DeviceType.DescribeDeviceType(req, nil)
	logger.Info("DescribeDeviceType openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil || client.Payload == nil {
		logger.Info("DescribeDeviceTypeInfo openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	if client.Payload.Result == nil {
		return nil, errors.New("device type not found")
	}
	imageCount := ""
	imagesReq := &requestTypes.QueryDeviceTypeImagePageRequest{
		DeviceTypeID: param.DeviceTypeID,
	}
	images, imagesErr := GetDeviceTypeImageList(logger, imagesReq)
	if imagesErr == nil {
		imageCount = strconv.Itoa(len(images.DeviceTypeImages))
	}
	obj := buildEntity(logger, client.Payload.Result.DeviceType)
	obj.ImageCount = imageCount
	return &obj, nil
}
func buildEntity(logger *log.Logger, v *sdkModels.DeviceType) response.DeviceType {
	language := logger.GetPoint("language").(string)
	interfaceModeName := BaseLogic.InterfaceMode[v.InterfaceMode]
	if language == BaseLogic.EN_US {
		interfaceModeName = BaseLogic.InterfaceModeEn[v.InterfaceMode]
	}
	obj := response.DeviceType{}
	/*raidMap := make(map[string]response.Raid) //	请求一次，避免循环请求
	raid, raidsErr := raidLogic.QueryRaidsAll(logger)
	if raidsErr == nil && len(raid.Raids) != 0 {
		for _, v := range raid.Raids {
			raidMap[v.RaidID] = *v
		}
	}
	array := []string{} // 拆分查询
	if len(v.Raid) != 0 {
		array = strings.Split(v.Raid, ",")
	}
	if len(array) != 0 {
		raid_ := ""
		for _, str := range array {
			_name := raidMap[str].Name
			if strings.ToUpper(_name) == "NORAID" && v.SystemVolumeAmount == 2 && v.RaidCan == "RAID" {
				if language == BaseLogic.EN_US {
					_name = "RAID0-stripping"
				} else {
					_name = "单盘RAID0"
				}
			}
			if v.RaidCan == "NORAID" { //配合前端展示
				_name = "--"
			}
			raid_ = raid_ + _name + "," // RAID0,RAID1

		}
		obj.Raid = raid_[0 : len(raid_)-1]
	}
	obj.RaidCan = v.RaidCan*/
	//
	//cpu_ := v.CPUManufacturer + " " + v.CPUModel + "(" + strconv.Itoa(int(v.CPUAmount)) + "*" + strconv.Itoa(int(v.CPUCores)) + "物理核," + v.CPUFrequency + "GHz)"
	//if language == BaseLogic.EN_US {
	//	cpu_ = v.CPUManufacturer + " " + v.CPUModel + "(" + strconv.Itoa(int(v.CPUAmount)) + "*" + strconv.Itoa(int(v.CPUCores)) + "cores," + v.CPUFrequency + "GHz)"
	//}
	//mem_ := strconv.Itoa(int(v.MemAmount)*int(v.MemSize)) + "GB(" + strconv.Itoa(int(v.MemSize)) + "GB*" + strconv.Itoa(int(v.MemAmount)) + ")" + v.MemType + " " + strconv.Itoa(int(v.MemFrequency)) + "MHz"
	//// 系统盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式)
	//// 240GB 240GB SATA SSD*2 RAID1
	//sv_ := strconv.Itoa(int(v.SystemVolumeAmount)*int(v.SystemVolumeSize)) + "GB(" + strconv.Itoa(int(v.SystemVolumeSize)) + "GB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount)) + "," + obj.Raid + ")"
	//if v.SystemVolumeUnit == "TB" {
	//	sv_ = Trim(strconv.FormatFloat(float64(v.SystemVolumeAmount)*float64(v.SystemVolumeSize)/1000, 'f', 2, 64)) + "TB(" + Trim(strconv.FormatFloat(float64(v.SystemVolumeSize)/1000, 'f', 2, 64)) + "TB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount)) + "," + obj.Raid + ")"
	//}
	//// 数据盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式)
	//dv_ := strconv.Itoa(int(v.DataVolumeAmount)*int(v.DataVolumeSize)) + "GB(" + strconv.Itoa(int(v.DataVolumeSize)) + "GB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount)) + "," + obj.Raid + ")"
	//if v.DataVolumeUnit == "TB" {
	//	dv_ = Trim(strconv.FormatFloat(float64(v.DataVolumeAmount)*float64(v.DataVolumeSize)/1000, 'f', 2, 64)) + "TB(" + Trim(strconv.FormatFloat(float64(v.DataVolumeSize)/1000, 'f', 2, 64)) + "TB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount)) + "," + obj.Raid + ")"
	//}
	//if v.DataVolumeAmount == 0 {
	//	dv_ = ""
	//}
	//// 显示品牌 显卡型号 个数，例 NVIDIA P40 * 4
	//gpu_ := v.GpuManufacturer + " " + v.GpuModel + "*" + strconv.Itoa(int(v.GpuAmount))
	//if v.GpuAmount == 0 {
	//	gpu_ = ""
	//}
	//nicInfo := strconv.Itoa(int(v.NicAmount)) + "*" + strconv.Itoa(int(v.NicRate)) + "GE"
	//
	//nic_ := BaseLogic.InterfaceMode[v.InterfaceMode]
	//if language == BaseLogic.EN_US {
	//	nic_ = BaseLogic.InterfaceModeEn[v.InterfaceMode]
	//}

	//obj.CpuInfo = cpu_
	//obj.MemInfo = mem_
	//obj.SvInfo = sv_
	//obj.DvInfo = dv_
	//obj.GpuInfo = gpu_
	//obj.NicInfo = nicInfo
	//obj.Nic = nic_
	obj.CpuInfo = v.CPUInfo
	obj.MemInfo = v.MemInfo
	obj.SvInfo = v.SvInfo
	obj.DvInfo = v.DvInfo
	obj.GpuInfo = v.GpuInfo
	obj.NicInfo = v.NicInfo

	obj.IDcID = v.IDcID
	obj.IdcName = v.IDcName
	obj.IdcNameEn = v.IDcNameEn
	obj.DeviceTypeID = v.DeviceTypeID
	obj.Name = v.Name
	obj.DeviceType = v.DeviceType
	obj.DeviceSeries = v.DeviceSeries

	obj.DeviceSeriesName = v.DeviceSeriesName
	obj.Architecture = v.Architecture
	obj.Height = int(v.Height)
	obj.Description = v.Description
	obj.CPUAmount = v.CPUAmount

	obj.CPUCores = v.CPUCores
	obj.CPUManufacturer = v.CPUManufacturer
	obj.CPUModel = v.CPUModel
	obj.CPUFrequency = v.CPUFrequency
	obj.MemSize = int(v.MemSize)

	obj.MemAmount = int(v.MemAmount)
	obj.MemType = v.MemType
	obj.MemFrequency = int(v.MemFrequency)
	obj.NicAmount = v.NicAmount
	obj.NicRate = int(v.NicRate)
	obj.InterfaceMode = v.InterfaceMode
	obj.InterfaceModeName = interfaceModeName

	obj.GpuAmount = v.GpuAmount

	obj.GpuManufacturer = v.GpuManufacturer
	obj.GpuModel = v.GpuModel

	obj.CreatedBy = v.CreatedBy
	obj.CreatedTime = util.UtcToTime(v.CreatedTime)
	obj.UpdatedBy = v.UpdatedBy
	obj.UpdatedBy = util.UtcToTime(v.UpdatedTime)

	obj.CPUSpec = v.CPUSpec
	obj.MemSpec = v.MemSpec
	obj.InstanceCount = int(v.InstanceCount)
	obj.DeviceCount = int(v.DeviceCount)
	obj.BootMode = v.BootMode
	volumes := []response.Volume{}

	if v.Volumes != nil {
		volumeResp := response.Volume{}
		for _, value := range v.Volumes {
			volumeResp.VolumeID = value.VolumeID
			volumeResp.VolumeAmount = value.VolumeAmount
			volumeResp.VolumeName = value.VolumeName
			volumeResp.VolumeUnit = value.VolumeUnit
			volumeResp.VolumeSize = value.VolumeSize

			volumeResp.VolumeType = value.VolumeType
			volumeResp.RaidCan = value.RaidCan
			volumeResp.Raid = value.Raid
			volumeResp.RaidID = value.RaidID
			volumeResp.InterfaceType = value.InterfaceType
			volumeResp.DiskType = value.DiskType
			volumes = append(volumes, volumeResp)
		}
	}
	obj.Volumes = volumes
	obj.IsNeedRaid = v.IsNeedRaid
	return obj
}

// 【机型管理】【机型详情信息】【关联镜像列表】
func GetDeviceTypeImageList(logger *log.Logger, param *requestTypes.QueryDeviceTypeImagePageRequest) (*response.DeviceTypeImageList, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceTypeParameters.NewDescribeDeviceTypeImagesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithDeviceTypeID(&param.DeviceTypeID)
	req.WithImageID(&param.ImageID)
	req.WithArchitecture(&param.Architecture)
	req.WithOsType(&param.OsType)
	req.WithSource(&param.Source)
	req.WithVersion(&param.Version)
	req.WithImageName(&param.ImageName)

	logger.Info("DescribeDeviceTypeImages openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.DeviceType.DescribeDeviceTypeImages(req, nil)
	fmt.Println("cli", client)
	logger.Info("DescribeDeviceTypeImages openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("DescribeDeviceTypeImages openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	resp := response.DeviceTypeImageList{}
	if len(client.Payload.Result.Images) == 0 {
		resp.DeviceTypeImages = make([]response.DeviceTypeImage, 0)
		return &resp, nil
	}
	obj := response.DeviceTypeImage{}
	for _, v := range client.Payload.Result.Images {
		obj.ID = v.ID
		obj.ImageID = v.ImageID
		obj.ImageName = v.ImageName
		obj.OsID = v.OsID
		obj.Format = v.Format

		obj.Filename = v.Filename
		obj.URL = v.URL
		obj.Hash = v.Hash
		obj.Source = v.Source
		obj.SourceName = v.SourceName
		obj.Description = v.Description

		obj.SystemPartition = v.SystemPartition
		obj.DataPartition = v.DataPartition
		obj.OsType = v.OsType
		obj.OsVersion = v.OsVersion
		obj.DeviceTypeNum = int(v.DeviceTypeNum)

		obj.Architecture = v.Architecture
		//obj.PreImage = "预置镜像"
		obj.IsBind = v.IsBind
		obj.CreatedBy = v.CreatedBy
		obj.CreatedTime = util.UtcToTime(v.CreatedTime)
		obj.UpdatedBy = v.UpdatedBy
		obj.UpdatedBy = util.UtcToTime(v.UpdatedTime)

		resp.DeviceTypeImages = append(resp.DeviceTypeImages, obj)
	}
	resp.PageNumber = client.Payload.Result.PageNumber
	resp.PageSize = client.Payload.Result.PageSize
	resp.TotalCount = client.Payload.Result.TotalCount
	return &resp, nil
}

// 【机型管理】【机型详情信息】【关联镜像列表】【添加镜像】
func AssociateDeviceTypeImage(logger *log.Logger, param *requestTypes.AssociateDeviceTypeImageRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceTypeParameters.NewAssociatedImageParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	array := []string{}
	if len(param.ImageIds) != 0 {
		array = strings.Split(param.ImageIds, ",")
	}
	body := sdkModels.AssociateImageRequest{
		DeviceTypeID: &param.DeviceTypeID,
		ImageIDs:     array,
	}
	req.WithBody(&body)
	logger.Info("AssociatedImage openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.DeviceType.AssociatedImage(req, nil)
	logger.Info("AssociatedImage openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil || client.Payload == nil {
		logger.Info("AssociateDeviceTypeImage openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}

// 【机型管理】【机型详情信息】【关联镜像列表】【删除镜像】
func DisassociateDeviceTypeImage(logger *log.Logger, param *requestTypes.DisassociateDeviceTypeImageRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceTypeParameters.NewDissociatedImageParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	body := sdkModels.DissociatedImageRequest{
		DeviceTypeID: &param.DeviceTypeID,
		ImageID:      &param.ImageId,
	}
	req.WithBody(&body)
	logger.Info("DissociatedImage openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.DeviceType.DissociatedImage(req, nil)
	logger.Info("DissociatedImage openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil || client.Payload == nil {
		logger.Info("DisassociateDeviceTypeImage openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}

// 【机型管理】【编辑机型】
func ModifyDeviceTypeInfo(logger *log.Logger, param *requestTypes.ModifyDeviceTypeRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceTypeParameters.NewModifyDeviceTypeParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithDeviceTypeID(param.DeviceTypeID)

	body := sdkModels.ModifyDeviceTypeRequest{
		Architecture:    param.Architecture,
		CPUAmount:       param.CPUAmount,
		CPUCores:        param.CPUCores,
		CPUFrequency:    param.CPUFrequency,
		CPUManufacturer: param.CPUManufacturer,
		CPUModel:        param.CPUModel,
		//DataVolumeAmount:        param.DataVolumeAmount,
		//DataVolumeInterfaceType: param.DataVolumeInterfaceType,
		//
		//DataVolumeType:            param.DataVolumeType,
		//DataVolumeUnit:            param.DataVolumeUnit,
		Description:     param.Description,
		DeviceSeries:    param.DeviceSeries,
		DeviceType:      param.DeviceType,
		GpuAmount:       param.GpuAmount,
		GpuManufacturer: param.GpuManufacturer,
		GpuModel:        param.GpuModel,
		Height:          param.Height,
		IDcID:           param.IDcID,
		InterfaceMode:   param.InterfaceMode,
		MemAmount:       param.MemAmount,
		MemFrequency:    param.MemFrequency,
		MemSize:         param.MemSize,
		MemType:         param.MemType,
		Name:            param.Name,
		NicAmount:       param.NicAmount,
		NicRate:         param.NicRate,
		RaidID:          param.RaidID,
		CPUSpec:         param.CpuSpec,
		MemSpec:         param.MemSpec,
		//SystemVolumeAmount:        param.SystemVolumeAmount,
		//SystemVolumeInterfaceType: param.SystemVolumeInterfaceType,
		//
		//SystemVolumeType: param.SystemVolumeType,
		//SystemVolumeUnit: param.SystemVolumeUnit,
		BootMode:   param.BootMode,
		Volumes:    param.Volumes,
		IsNeedRaid: param.IsNeedRaid,
	}
	//if param.SystemVolumeSize != nil && param.SystemVolumeUnit != nil {
	//	systemVolumeSize := int64(*param.SystemVolumeSize)
	//	if *param.SystemVolumeUnit == "TB" {
	//		systemVolumeSize = int64(*param.SystemVolumeSize * 1000)
	//	}
	//	body.SystemVolumeSize = &systemVolumeSize
	//
	//}
	//if param.RaidCan != "" {
	//	body.RaidCan = &param.RaidCan
	//}

	//if param.DataVolumeSize != nil && param.DataVolumeUnit != nil {
	//	dataVolumeSize := int64(*param.DataVolumeSize)
	//	if *param.DataVolumeUnit == "TB" {
	//		dataVolumeSize = int64(*param.DataVolumeSize * 1000)
	//	}
	//	body.DataVolumeSize = &dataVolumeSize
	//
	//}

	req.WithBody(&body)
	logger.Info("ModifyDeviceType openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.DeviceType.ModifyDeviceType(req, nil)
	logger.Info("ModifyDeviceType openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("ModifyDeviceTypeInfo openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	if client == nil || client.Payload == nil || client.Payload.Result == nil {
		return nil, errors.New("modify device info failed")
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}

// 【机型管理】【删除】
func DeleteDeviceTypeInfo(logger *log.Logger, param *requestTypes.DeleteDeviceTypeRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceTypeParameters.NewDeleteDeviceTypeParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithDeviceTypeID(param.DeviceTypeID)
	logger.Info("DeleteDeviceType openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.DeviceType.DeleteDeviceType(req, nil)
	logger.Info("DeleteDeviceType openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil || client.Payload == nil {
		logger.Info("DeleteDeviceTypeInfo openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	if client.Payload.Result == nil {
		return nil, errors.New("device type not found")
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}
func Trim(str string) string {
	s := strings.TrimRight(str, "0")
	s = strings.TrimRight(s, ".")
	return s
}

func GetAvailableDeviceTypes(logger *log.Logger, param requestTypes.GetAvailableDeviceTypesRequest) (*response.AvailableDeviceType, error) {

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
	responseApi, err := sdk.SdkClient.DeviceType.DescribeDeviceTypes(req, nil)
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
	res := map[string][]response.DeviceTypeReinstall{}
	for _, v := range deviceTypes {
		if _, ok := res[v.DeviceSeries]; !ok {
			res[v.DeviceSeries] = []response.DeviceTypeReinstall{sdkDeviceType2DeviceType(logger, v)}
		} else {
			res[v.DeviceSeries] = append(res[v.DeviceSeries], sdkDeviceType2DeviceType(logger, v))
		}
	}
	obj := &response.AvailableDeviceType{}
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

func sdkDeviceType2DeviceType(logger *log.Logger, v *sdkModels.DeviceType) response.DeviceTypeReinstall {
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
	if language == BaseLogic.EN_US {
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

	/*sv_ := strconv.Itoa(int(v.SystemVolumeSize)) + "GB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount))
	if v.SystemVolumeUnit == "TB" {
		sv_ = Trim(strconv.FormatFloat(float64(v.SystemVolumeSize)/1000, 'f', 2, 64)) + "TB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount))
	}
	// 数据盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式）
	dv_ := strconv.Itoa(int(v.DataVolumeSize)) + "GB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount))
	if v.DataVolumeUnit == "TB" {
		dv_ = Trim(strconv.FormatFloat(float64(v.DataVolumeSize)/1000, 'f', 2, 64)) + "TB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount))
	}
	if v.DataVolumeAmount == 0 {
		dv_ = ""
	}*/

	// 显示品牌 显卡型号 个数，例 NVIDIA P40 * 4
	gpu_ := v.GpuManufacturer + " " + v.GpuModel + "*" + strconv.Itoa(int(v.GpuAmount))
	if v.GpuAmount == 0 {
		gpu_ = ""
	}
	nicInfo := strconv.Itoa(int(v.NicAmount)) + "*" + strconv.Itoa(int(v.NicRate)) + "GE"

	return response.DeviceTypeReinstall{
		Name:         v.Name,
		DeviceTypeId: v.DeviceTypeID,
		//Cpu:              fmt.Sprintf(cputpl, v.CPUAmount, v.CPUCores, v.CPUFrequency),
		//Mem:              fmt.Sprintf(memtpl, v.MemAmount, v.MemSize, v.MemFrequency),
		//System:           fmt.Sprintf(systemtpl, v.SystemVolumeAmount, v.SystemVolumeSize, v.SystemVolumeType),
		//Data:             fmt.Sprintf(datatpl, v.DataVolumeAmount, v.DataVolumeSize, v.DataVolumeType),
		//Nic:              fmt.Sprintf(nictpl, v.NicAmount, v.NicRate),
		//Gpu:              fmt.Sprintf(gputpl, v.GpuAmount, v.GpuManufacturer, v.GpuModel),
		Cpu: cpu_,
		Mem: mem_,
		/*System:           sv_,
		Data:             dv_,*/
		Nic:              nicInfo,
		Gpu:              gpu_,
		AvailableStock:   v.StockAvailable,
		DeviceSeries:     v.DeviceSeries,
		DeviceSeriesName: v.DeviceSeriesName,
		BootMode:         v.BootMode,
	}
}

func DescribeVolumesByDeviceType(logger *log.Logger, deviceTypeId string) ([]*sdkModels.VolumeIt, error) {

	logid := logger.GetPoint("logid").(string)
	userId := logger.GetPoint("userId").(string)
	req := sdkDeviceTypeParameters.NewDescribeVolumesByDeviceTypeParams()
	req.WithTraceID(logid)
	req.WithBmpUserID(&userId)
	req.WithDeviceTypeID(deviceTypeId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("DescribeVolumesByDeviceType openapi request:", util.InterfaceToJson(req))
	responseApi, err := sdk.SdkClient.DeviceType.DescribeVolumesByDeviceType(req, nil)
	logger.Info("DescribeVolumesByDeviceType openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeVolumesByDeviceType openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result, nil

}
