package DeviceLogic

import (
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"net"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-operation-api/constant"

	sdkDeviceParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/device"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
	beego "github.com/beego/beego/v2/server/web"

	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-operation-api/services/excel"
	sdk "coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
)

// 【设备管理】【设备列表】
func GetDeviceList(logger *log.Logger, param *requestTypes.QueryDeviceListRequest) (*response.DevicePage, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewDescribeDevicesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithDeviceTypeID(&param.DeviceTypeID)
	req.WithDeviceSeries(&param.DeviceSeries)
	req.WithIloIP(&param.IloIP)
	req.WithInstanceID(&param.InstanceID)
	req.WithInstanceName(&param.InstanceName)
	req.WithManageStatus(&param.ManageStatus)
	req.WithCollectStatus(&param.CollectStatus)
	req.WithSn(&param.Sn)
	req.WithUserID(&param.UserID)
	req.WithUserName(&param.UserName)
	req.WithPageNumber(&param.PageNumber)
	req.WithPageSize(&param.PageSize)
	req.WithIsAll(&param.IsAll)
	logger.Info("DescribeDevices openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.DescribeDevices(req, nil)
	logger.Info("DescribeDevices openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("DescribeDevices openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	resp := response.DevicePage{}
	obj := response.Device{}
	if len(client.Payload.Result.Devices) == 0 {
		resp.Devices = make([]response.Device, 0)
		return &resp, nil
	}

	for _, v := range client.Payload.Result.Devices {
		obj = buildEntity(logger, v, param.Show)
		resp.Devices = append(resp.Devices, obj)
	}
	resp.PageNumber = client.Payload.Result.PageNumber
	resp.PageSize = client.Payload.Result.PageSize
	resp.TotalCount = client.Payload.Result.TotalCount
	return &resp, nil
}

// 导出Excel文件
func ExportDeviceExcel(logger *log.Logger, result []response.Device) (string, string, error) {
	language := logger.GetPoint("language").(string)
	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)
	for _, data := range result {

		//dt, err := DeviceTypeLogic.GetDeviceTypeInfo(logger, &requestTypes.QueryDeviceTypeRequest{
		//	DeviceTypeID: data.DeviceTypeID,
		//})
		//if err != nil {
		//	logger.Warn("deviceType err:", err.Error())
		//	return "", "", err
		//}

		nic_ := BaseLogic.InterfaceMode[data.InterfaceMode] //"单网口"
		if language == BaseLogic.EN_US {
			nic_ = BaseLogic.InterfaceModeEn[data.InterfaceMode]
		}
		sheetData = append(sheetData, []string{
			data.Sn,               // SN
			data.DeviceTypeName,   // 机型名称
			data.DeviceSeriesName, //机型类型
			data.ManageStatusName, // 管理状态，包含：已入库 已上架 已创建 上架中 下架中；
			data.IdcName,          // 机房名称
			data.Cabinet,          // 机柜编码
			data.UPosition,        // 所在U位
			data.Brand,            // 品牌
			data.Model,            // 型号
			data.Architecture,     // 体系架构
			data.ImageName,        // 实例镜像

			data.InstanceName,       // 实例名称
			data.InstanceID,         // 实例ID
			data.InstanceStatusName, // 实例状态
			data.Locked,             // 实例锁定状态
			data.UserName,           // 实例属主
			data.CpuInfo,            // CPU

			data.MemInfo,   // 内存：显示总容量
			data.SvInfo,    // 系统盘
			data.DvInfo,    // 数据盘
			data.GpuInfo,   // GPU
			data.SwitchIP1, // 网口1交换机IP

			data.SwitchIP2,   // 网口2交换机IP
			data.SwitchPort1, // 网口1（eth0）上联端口
			data.SwitchPort2, // 网口2（eth1）上联端口
			nic_,             // 网络设置
			data.IloIP,       // 带外IP
			data.PrivateIPv4, // 内网IPv4
			data.PrivateEth1IPV4,
			data.PrivateIPv6, // 内网IPv6
			data.PrivateEth1IPV6,
			data.CreatedTime, // 创建时间
			data.Description, // 设备描述
		})
	}
	exportExcelHeader := BaseLogic.ExportDeviceHeader
	if language == BaseLogic.EN_US {
		exportExcelHeader = BaseLogic.ExportDeviceHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportExcelHeader, sheetData); err != nil {
		logger.Warn("create excel error:", err.Error())
		return "", "", err
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	// 生成100000到100000+900000之间的随机数，左闭右开
	downloadFileName := "device_list_" + time.Now().Format(BaseLogic.FormatDate) + "_" + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx"
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Warn("保存excel到本地失败" + err.Error())
		return "", "", err
	}
	return fileName, downloadFileName, nil
}

// 【设备管理】【设备详情】
func GetDeviceInfo(logger *log.Logger, param *requestTypes.QueryDeviceInfoRequest) (*response.Device, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewDescribeDeviceParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithTraceID(logid)
	req.WithDeviceID(param.DeviceID)
	logger.Info("DescribeDevice openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.DescribeDevice(req, nil)
	logger.Info("DescribeDevice openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("DescribeDevice openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	if client.Payload.Result == nil {
		return nil, errors.New("device not found")
	}

	obj := buildEntity(logger, client.Payload.Result.Device, param.Show)
	return &obj, nil
}

// 【设备管理】【设备磁盘详情】
func GetDeviceDiskDetail(logger *log.Logger, param *requestTypes.QueryDeviceDisksRequest) (*models.DeviceDisks, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewDescribeDeviceDisksParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithTraceID(logid)
	req.WithDeviceID(param.DeviceID)
	logger.Info("DescribeDeviceDisks openapi sdk req:", util.InterfaceToJson(req))
	res, err := sdk.SdkClient.Device.DescribeDeviceDisks(req, nil)
	logger.Info("DescribeDeviceDisks openapi sdk resp:", util.InterfaceToJson(res))
	if err != nil {
		logger.Warn("DescribeDeviceDisks openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	if res.Payload.Result == nil {
		return nil, errors.New("device not found")
	}

	return res.Payload.Result, nil
}

// 列表和详情共用方法
func buildEntity(logger *log.Logger, v *sdkModels.Device, show string) response.Device {
	//dt, err := DeviceTypeLogic.GetDeviceTypeInfo(logger, &requestTypes.QueryDeviceTypeRequest{
	//	DeviceTypeID: v.DeviceTypeID,
	//})
	//if err != nil {
	//	logger.Warn("deviceType err:", err.Error())
	//	return response.Device{}
	//}

	interfaceModeName := BaseLogic.InterfaceMode[v.InterfaceMode] //"单网口"
	if logger.GetPoint("language").(string) == BaseLogic.EN_US {
		interfaceModeName = BaseLogic.InterfaceModeEn[v.InterfaceMode]
	}
	lockedName := BaseLogic.InstanceLockedName[v.Locked] //"单网口"
	if logger.GetPoint("language").(string) == BaseLogic.EN_US {
		lockedName = BaseLogic.InstanceLockedNameEn[v.Locked]
	}

	obj := response.Device{}
	obj.CpuInfo = v.CPUInfo
	obj.MemInfo = v.MemInfo
	obj.SvInfo = v.SvInfo
	obj.DvInfo = v.DvInfo
	obj.GpuInfo = v.GpuInfo
	obj.NicInfo = v.NicInfo
	obj.InterfaceModeName = interfaceModeName
	obj.LockedName = lockedName

	obj.AdapterID = int(v.AdapterID)
	obj.Architecture = v.Architecture
	obj.Brand = v.Brand
	obj.Model = v.Model
	obj.CPUAmount = v.CPUAmount
	obj.CPUCores = v.CPUCores

	obj.CPUFrequency = v.CPUFrequency
	obj.CPUManufacturer = v.CPUManufacturer
	obj.CPUModel = v.CPUModel
	obj.CPURoads = v.CPURoads
	obj.Reason = v.Reason
	obj.Cabinet = v.Cabinet

	obj.DataVolumeAmount = v.DataVolumeAmount
	obj.DataVolumeInterfaceType = v.DataVolumeInterfaceType
	obj.DataVolumeSize = int(v.DataVolumeSize)
	obj.DataVolumeType = v.DataVolumeType
	obj.Description = v.Description

	obj.DeviceID = v.DeviceID
	obj.DeviceSeries = v.DeviceSeries
	obj.DeviceSeriesName = v.DeviceSeriesName
	obj.DeviceType = v.DeviceType
	obj.DeviceTypeID = v.DeviceTypeID
	obj.DeviceTypeName = v.DeviceTypeName

	obj.Enclosure1 = v.Enclosure1
	obj.Enclosure2 = v.Enclosure2
	obj.Gateway = v.Gateway
	obj.GpuAmount = v.GpuAmount
	obj.GpuManufacturer = v.GpuManufacturer

	obj.GpuModel = v.GpuModel
	obj.ID = int(v.ID)
	obj.IdcID = *v.IdcID
	obj.IdcName = v.IdcName
	obj.IdcNameEn = v.IDcNameEn
	obj.IloIP = v.IloIP

	obj.IloUser = v.IloUser
	obj.ImageName = v.ImageName
	obj.InstanceCreatedTime = util.UtcToTime(v.InstanceCreatedTime)
	obj.InstanceDescription = v.InstanceDescription

	obj.InstanceID = v.InstanceID
	obj.InstanceName = v.InstanceName
	obj.UserID = v.UserID
	obj.UserName = v.UserName
	obj.InstanceStatus = v.InstanceStatus
	obj.InstanceStatusName = v.InstanceStatusName
	obj.InstanceReason = v.InstanceReason
	obj.InterfaceMode = v.InterfaceMode
	obj.Locked = v.Locked

	obj.Mac1 = v.Mac1
	obj.Mac2 = v.Mac2
	// in=已入库|putway=已上架|unputway=已下架|created=已创建|putawaying=上架中|unputwaying=下架中
	obj.ManageStatus = v.ManageStatus         // 英文展示
	obj.ManageStatusName = v.ManageStatusName // 中文展示
	obj.Mask = v.Mask
	obj.MemAmount = int(v.MemAmount)

	obj.MemFrequency = int(v.MemFrequency)
	obj.MemSize = int(v.MemSize)
	obj.MemType = v.MemType
	obj.NicAmount = v.NicAmount
	obj.NicRate = int(v.NicRate)

	obj.PrivateIPv4 = v.PrivateIPV4
	obj.PrivateIPv6 = v.PrivateIPV6
	obj.Eth1Mask = v.Eth1Mask
	obj.PrivateEth1IPV4 = v.PrivateEth1IPV4
	obj.PrivateEth1IPV6 = v.PrivateEth1IPV6
	obj.RaidDriver = v.RaidDriver
	obj.Slot1 = int(v.Slot1)
	obj.Slot2 = int(v.Slot2)

	obj.Sn = v.Sn
	obj.SwitchIP = v.SwitchIP
	obj.SwitchIP1 = v.SwitchIP1
	obj.SwitchIP2 = v.SwitchIP2

	obj.SwitchPort1 = v.SwitchPort1
	obj.SwitchPort2 = v.SwitchPort2
	obj.SwitchUser1 = v.SwitchUser1
	obj.SwitchUser2 = v.SwitchUser2

	obj.SystemVolumeAmount = v.SystemVolumeAmount
	obj.SystemVolumeInterfaceType = v.SystemVolumeInterfaceType
	obj.SystemVolumeSize = int(v.SystemVolumeSize)
	obj.SystemVolumeType = v.SystemVolumeType
	obj.UPosition = v.UPosition

	obj.CreatedBy = v.CreatedBy
	obj.CreatedTime = util.UtcToTime(v.CreatedTime)
	obj.UpdatedBy = v.UpdatedBy
	obj.UpdatedBy = util.UtcToTime(v.UpdatedTime)
	if show == "iloPassword" {
		obj.IloPassword = v.IloPassword
		obj.SwitchPassword1 = ""
		obj.SwitchPassword2 = ""
	}
	if show == "switchPassword1" {
		obj.IloPassword = ""
		obj.SwitchPassword1 = v.SwitchPassword1
		obj.SwitchPassword2 = ""
	}
	if show == "switchPassword2" {
		obj.IloPassword = ""
		obj.SwitchPassword1 = ""
		obj.SwitchPassword2 = v.SwitchPassword2
	}
	obj.CollectStatus = v.CollectStatus
	obj.CollectFailReason = v.CollectFailReason
	obj.IsNeedRaid = v.IsNeedRaid
	return obj
}

// 【设备管理】【设备下架】支持多sn，英文逗号分隔
func UnPutawayDevice(logger *log.Logger, param *requestTypes.UnPutawayDeviceRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewUnMountDeviceParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithTraceID(logid)

	//array := []string{}
	//if len(param.Sns) != 0 {
	//	array = strings.Split(param.Sns, ",")
	//}
	body := sdkModels.UnMountDevicesRequest{
		DeviceID: &param.DeviceID,
	}
	req.WithBody(&body)
	logger.Info("UnMountDevice openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.UnMountDevice(req, nil)
	logger.Info("UnMountDevice openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("UnMountDevice openapi error:", err.Error())
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

// 【设备管理】【设备上架】支持多sn，英文逗号分隔
func PutawayDevice(logger *log.Logger, param *requestTypes.PutawayDeviceRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewMountDeviceParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	//array := []string{}
	//if len(param.Sns) != 0 {
	//	array = strings.Split(param.Sns, ",")
	//}
	body := sdkModels.MountDevicesRequest{
		DeviceID: &param.DeviceID,
	}
	req.WithBody(&body)
	logger.Info("MountDevice openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.MountDevice(req, nil)
	logger.Info("MountDevice openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("MountDevice openapi error:", err.Error())
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

// 【设备管理】【设备修改】
func ModifyDevice(logger *log.Logger, param *requestTypes.ModifyDeviceRequest, deviceId string) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewModifyDeviceParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.DeviceID = deviceId
	body := sdkModels.ModifyDevicesRequest{
		Description: param.Description,
	}
	req.WithBody(&body)
	logger.Info("ModifyDevice openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.ModifyDevice(req, nil)
	logger.Info("ModifyDevice openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("MountDevice openapi error:", err.Error())
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

// 【设备管理】【删除设备】
func DeleteDevice(logger *log.Logger, param *requestTypes.DeleteDeviceRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewDeleteDeviceParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithDeviceID(param.DeviceId)
	logger.Info("DeleteDevice openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.DeleteDevice(req, nil)
	logger.Info("DeleteDevice openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("DeleteDevice openapi error:", err.Error())
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

// 【设备管理】【移除设备】
func RemoveDevice(logger *log.Logger, deviceId string) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewRemoveDeviceParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithDeviceID(deviceId)
	logger.Info("RemoveDevice openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.RemoveDevice(req, nil)
	logger.Info("RemoveDevice openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("RemoveDevice openapi error:", err.Error())
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

func UploadDevices(logger *log.Logger, param *requestTypes.UploadDeviceRequest, header *multipart.FileHeader, requestId string) ([]*response.UploadDevice, error) {
	arr := strings.Split(header.Filename, ".")
	suffix := strings.ToLower(arr[len(arr)-1])
	//language := logger.GetPoint("language").(string)
	if suffix != "xlsx" { //util.InArray(suffix, AvailableSuffix) {
		logger.Warn("file format not support:", suffix)
		panic(constant.BuildInvalidArgumentWithArgs("文件类型不符合,必须是xlsx类型", "file must be xlsx"))
	}
	dir, err := util.GetBaseDir()
	if err != nil {
		logger.Warn("获取当前根目录失败" + err.Error())
		return nil, err
	}
	//prefix := util.GenUuid()
	uploadPath, _ := beego.AppConfig.String("upload.path")
	util.CreateDirIfNotExist(uploadPath)                                       //如果目录不存在，创建一个目录
	localFileName := path.Join(dir, uploadPath, requestId+"_"+header.Filename) //prefix+"-"+header.Filename
	fmt.Println(localFileName)
	if err = util.SaveToLocal(localFileName, param.DeviceFile); err != nil {
		logger.Warn("保存excel到本地失败" + err.Error())
		return nil, err
	}
	f, err := excel.LoadExcel(localFileName)
	if err != nil {
		return nil, err
	}
	defer DelUploadExcel(localFileName)
	list, err := f.GetSheetData()
	if err != nil {
		return nil, err
	}
	logger.Infof("UploadExcel.GetSheetData:%s", util.InterfaceToJson(list))
	devices := []*response.UploadDevice{}

	sns := []string{}
	iloIps := []string{}
	ipv6s := []string{}
	if len(list) > 10002 {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("单次上传不能超过10000条"), fmt.Sprintf("exceed 10000")))
	}
	fmt.Println("the excel parse content=======:", list)
	for k, v := range list {
		if k >= 2 {
			// if len(v) == 25 {
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("AdapterId 不能为空"), fmt.Sprintf("AdapterId empty")))
			// }
			// if len(v) == 26 {
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("Enclosure1 不能为空"), fmt.Sprintf("Enclosure1 empty")))
			// }
			// if len(v) == 27 {
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("Slot1 不能为空"), fmt.Sprintf("Slot1 empty")))
			// }
			// if len(v) < 28 { //slot1v[24]
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("有必填项没有写，请检查"), fmt.Sprintf("empty invalidate,please check")))
			// }
			specialMatch := regexp.MustCompile("^.{1,128}$").MatchString //[一-龥_a-zA-Z0-9_-]
			if !specialMatch(v[0]) {                                     //sn
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("Sn %s 不合法", v[0]), fmt.Sprintf("Sn %s invalidate", v[0])))
			}
			if !specialMatch(v[1]) { //Cabinet机柜
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("Cabinet %s 不合法", v[1]), fmt.Sprintf("Cabinet %s invalidate", v[1])))
			}
			if !specialMatch(v[2]) { //U位
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("UPosition %s 不合法", v[2]), fmt.Sprintf("UPosition %s invalidate", v[2])))
			}
			// if !specialMatch(v[3]) {
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("Brand %s 不合法", v[3]), fmt.Sprintf("Brand %s invalidate", v[3])))
			// }
			// if !specialMatch(v[4]) {
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("Model %s 不合法", v[4]), fmt.Sprintf("Model %s invalidate", v[4])))
			// }
			if _, err := net.ParseMAC(v[3]); err != nil { //mac1(eth0)
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("mac1(eth0) %s 不合法", v[3]), fmt.Sprintf("mac1(eth0) %s invalidate", v[3])))
			}
			// if _, err := net.ParseMAC(v[6]); err != nil {
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("Mac2 %s 不合法", v[6]), fmt.Sprintf("Mac2 %s invalidate", v[6])))
			// }
			// specialMatch = regexp.MustCompile("^[a-zA-Z/0-9_-]{1,128}$").MatchString
			// if !specialMatch(v[9]) {
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("switchport1 %s 不合法", v[9]), fmt.Sprintf("switchport1 %s invalidate", v[9])))
			// }
			// specialMatch = regexp.MustCompile("^[a-zA-Z/0-9_-]{1,128}$").MatchString
			// if !specialMatch(v[10]) {
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("switchport2 %s 不合法", v[10]), fmt.Sprintf("switchport2 %s invalidate", v[10])))
			// }
			if net.ParseIP(v[4]) == nil { //mask
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("mask %s 不合法", v[4]), fmt.Sprintf("mask %s invalidate", v[4])))
			}

			if net.ParseIP(v[5]) == nil { //eth0 private ipv4
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("ipv4(eth0) %s 不合法", v[5]), fmt.Sprintf("eth0 PrivateIPv4 %s invalidate", v[5])))
			}
			if v[6] != "" { //eth0 ipv6
				if strings.Contains(v[6], "/") { //2403:1EC0:8549:60C0::1/64
					arr := strings.Split(v[6], "/")
					ipv6 := arr[0]
					mask := arr[1]
					maskCount, err := strconv.Atoi(mask)
					if err != nil {
						panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("eth0 ipv6 %s 不合法", v[6]), fmt.Sprintf("eth0 ipv6 %s invalidate", v[6])))
					}
					if !strings.Contains(ipv6, ":") {
						panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("eth0 ipv6 %s 不合法", v[6]), fmt.Sprintf("eth0 ipv6 %s invalidate", v[6])))
					}
					if net.ParseIP(ipv6) == nil || (maskCount < 0 || maskCount > 128) {
						panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("eth0 ipv6 %s 不合法", v[6]), fmt.Sprintf("eth0 ipv6 %s invalidate", v[6])))
					}
				} else {
					ipv6 := v[6]
					if !strings.Contains(ipv6, ":") {
						panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("eth0 ipv6 %s 不合法", v[6]), fmt.Sprintf("eth0 ipv6 %s invalidate", v[6])))
					}
					if net.ParseIP(ipv6) == nil {
						panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("eth0 ipv6 %s 不合法", v[6]), fmt.Sprintf("eth0 ipv6 %s invalidate", v[6])))
					}
				}
			}
			if net.ParseIP(v[7]) == nil { //gateway
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("Gateway %s 不合法", v[7]), fmt.Sprintf("Gateway %s invalidate", v[7])))
			}

			if net.ParseIP(v[8]) == nil {
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("IloIP %s 不合法", v[8]), fmt.Sprintf("IloIP %s invalidate", v[8])))
			}

			specialMatch = regexp.MustCompile("^.{1,128}$").MatchString
			if !specialMatch(v[9]) && v[9] != "" { //IloUser
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("IloUser %s 不合法", v[9]), fmt.Sprintf("IloUser %s invalidate", v[9])))
			}
			if !specialMatch(v[10]) && v[10] != "" { //IloPassword
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("IloPassword %s 不合法", v[10]), fmt.Sprintf("IloPassword %s invalidate", v[10])))
			}

			if util.InArrayString(v[0], sns) {
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("文件中sn %s 重复", v[0]), fmt.Sprintf("file sn %s repeat", v[0])))
			}
			sns = append(sns, v[0])
			if util.InArrayString(v[8], iloIps) {
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("文件中带外IP %s 重复", v[8]), fmt.Sprintf("file iloIp %s repeat", v[8])))
			}
			iloIps = append(iloIps, v[8])
			//res, err := rdsUtil.RedisCli.SAdd("sns", v[0]).Result()
			//if res == 0 { //没有设置成功，说明有元素存在了
			//	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("文件中sn %s 重复", v[0]), fmt.Sprintf("file sn %s repeat", v[0])))
			//}
			//fmt.Println("err", res, err, rdsUtil.RedisCli.SMembers("b"))

			if v[6] != "" { //eth0 ipv6
				if util.InArrayString(v[6], ipv6s) {
					panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("文件中eth0 IPV6 %s 重复", v[6]), fmt.Sprintf("file eth0 IPV6 %s repeat", v[6])))
				}
				ipv6s = append(ipv6s, v[6])
			}

			// if v[25] == "" { //adapterid
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("AdapterId %s 不能为空", v[25]), fmt.Sprintf("AdapterId %s empty", v[25])))
			// }
			// if v[26] == "" { //enslouser1
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("Enclosure1 %s 不能为空", v[26]), fmt.Sprintf("Enclosure1 %s empty", v[26])))
			// }
			// if v[27] == "" { //slot1
			// 	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("Slot1 %s 不能为空", v[27]), fmt.Sprintf("Slot1 %s empty", v[27])))
			// }
			// adapterId, _ := strconv.Atoi(v[25])
			// slot1, _ := strconv.Atoi(v[27])
			device := &response.UploadDevice{
				Sn:        v[0],
				Cabinet:   v[1],
				UPosition: v[2],
				// Brand:     v[3],
				// Model:     v[4],

				Mac1: strings.ToLower(v[3]), //mac地址转小写
				// Mac2: strings.ToLower(v[6]),
				// SwitchIP1:   v[7],
				// SwitchIP2:   v[8],
				// SwitchPort1: v[9],
				// SwitchPort2: v[10],

				Mask: v[4],
				// MaskEth1:    v[12],
				Gateway:     v[7],
				IloIP:       v[8],
				IloUser:     v[9],
				IloPassword: v[10],

				// SwitchUser1:     v[17],
				// SwitchPassword1: v[18],
				// SwitchUser2:     v[19],
				// SwitchPassword2: v[20],
				PrivateIPv4: v[5],
				// PrivateEth1IPv4: v[22],
				PrivateIPv6: v[6],
				// PrivateEth1IPv6: v[24],

				// AdapterID: adapterId,
				// Enclosure1: v[26],
				// Slot1:      slot1,
			}
			// if len(v) == 30 { //如果30，说明槽位2都有数据（不能只出现slot2或者enclouser2一个，要么都出现，要么都不出现）
			// 	device.Enclosure2 = v[28]
			// 	slot2, _ := strconv.Atoi(v[29])
			// 	device.Slot2 = slot2
			// }

			devices = append(devices, device)
		}
	}
	return devices, nil
}
func DelUploadExcel(localFileName string) error {
	if err := os.Remove(localFileName); err != nil {
		return err
	}
	return nil
}
func CreateDevices(logger *log.Logger, param *requestTypes.CreateDeviceRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewCreateDeviceParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithTimeout(time.Minute * 10) //超时时间设置
	devices := []*sdkModels.CreateDeviceSpec{}
	for _, v := range param.Devices {
		// adapterId := int64(v.AdapterID)
		// slot1 := int64(v.Slot1)
		deviceSpec := &sdkModels.CreateDeviceSpec{
			Sn:        &v.Sn,
			Cabinet:   &v.Cabinet,
			UPosition: &v.UPosition,
			// Brand:     &v.Brand,
			// Model:     &v.Model,

			Mac1: &v.Mac1,
			// Mac2: &v.Mac2,
			// SwitchIP1:   &v.SwitchIP1,
			// SwitchIP2:   &v.SwitchIP2,
			// SwitchPort1: &v.SwitchPort1,
			// SwitchPort2: &v.SwitchPort2,

			Mask: &v.Mask,
			// MaskEth1:    v.MaskEth1,
			Gateway:     &v.Gateway,
			IloIP:       &v.IloIP,
			IloUser:     v.IloUser,
			IloPassword: v.IloPassword,

			// SwitchUser1:     v.SwitchUser1,
			// SwitchPassword1: v.SwitchPassword1,
			// SwitchUser2:     v.SwitchUser2,
			// SwitchPassword2: v.SwitchPassword2,
			PrivateIPV4: &v.PrivateIPv4,
			// PrivateEth1IPV4: v.PrivateEth1IPv4,
			PrivateIPV6: v.PrivateIPv6,
			// PrivateEth1IPV6: v.PrivateEth1IPv6,

			// RaidDriver: "megacli64", //写死
			// AdapterID: &adapterId,
			// Enclosure1: &v.Enclosure1,
			// Slot1:      &slot1,
			// Enclosure2: v.Enclosure2,
			// Slot2:      int64(v.Slot2),
		}
		devices = append(devices, deviceSpec)
	}
	//return response.DeviceIds{}, err
	body := sdkModels.CreateDevicesRequest{
		IDcID: &param.IdcId,
		// DeviceTypeID: &param.DeviceTypeId,
		Devices: devices,
	}
	req.WithBody(&body)
	logger.Info("openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.CreateDevice(req, nil)
	logger.Info(":", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("CreateDevices openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	if client == nil || client.Payload == nil || client.Payload.Result == nil {
		return nil, errors.New("create device failed")
	}
	return &response.CommonResponse{
		Success: client.Payload.Result.Success,
	}, nil
}

func CollectDevice(logger *log.Logger, param *requestTypes.CollectDevice) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewCollectDeviceInfoParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithTimeout(time.Minute * 10) //超时时间设置
	c := []*sdkModels.CollectItem{
		&sdkModels.CollectItem{
			Sn:            &param.Sn,
			AllowOverride: &param.AllowOverride,
		},
	}
	body := sdkModels.CollectDeviceInfoRequest{
		Collects: c,
	}
	req.WithBody(&body)
	logger.Info("openapi CollectDeviceInfo req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.CollectDeviceInfo(req, nil)
	logger.Info("openapi CollectDeviceInfo resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("openapi CollectDeviceInfo openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	if client == nil || client.Payload == nil || client.Payload.Result == nil {
		return nil, errors.New("openapi CollectDeviceInfo failed")
	}
	return &response.CommonResponse{
		Success: client.Payload.Result.Success,
	}, nil

}

// 设备关联磁盘
func AssociateDeviceDisks(logger *log.Logger, param *requestTypes.AssociateDeviceDisksRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewAssociateDeviceDisksParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithTraceID(logid)

	//array := []string{}
	//if len(param.Sns) != 0 {
	//	array = strings.Split(param.Sns, ",")
	//}
	volumes := []*sdkModels.AssociateDeviceDiskSpec{}
	for _, v := range param.Volumes {
		item := &sdkModels.AssociateDeviceDiskSpec{
			DiskIDs:  v.DiskIDs,
			VolumeID: v.VolumeID,
		}
		volumes = append(volumes, item)
	}
	body := sdkModels.AssociateDeviceDisksRequest{
		DeviceID:     &param.DeviceID,
		DeviceTypeID: &param.DeviceTypeID,
		Volumes:      volumes,
	}
	req.WithBody(&body)
	logger.Info("AssociateDeviceDisks openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.AssociateDeviceDisks(req, nil)
	logger.Info("AssociateDeviceDisks openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("AssociateDeviceDisks openapi error:", err.Error())
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

func GetAssociatedDisks(logger *log.Logger, param *requestTypes.GetAssociatedDisksRequest) ([]*models.Disk, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewGetAssociatedDisksParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithTraceID(logid)
	req.WithDeviceID(param.DeviceID)
	req.WithVolumeID(param.VolumeID)
	req.WithDeviceTypeID(param.DeviceTypeID)
	logger.Info("GetAssociatedDisks openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.GetAssociatedDisks(req, nil)
	logger.Info("GetAssociatedDisks openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("GetAssociatedDisks openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	if client.Payload.Result == nil {
		return nil, errors.New("disks not found")
	}

	return client.Payload.Result, nil
}

// 设备关联机型
func DeviceAssociateDeviceType(logger *log.Logger, param *requestTypes.DeviceAssociateDeviceTypeRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewDeviceAssociateDeviceTypeParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithTraceID(logid)

	body := sdkModels.DeviceAssociateDeviceTypeRequest{
		DeviceID:     &param.DeviceID,
		DeviceTypeID: &param.DeviceTypeID,
	}
	req.WithBody(&body)
	logger.Info("DeviceAssociateDeviceType openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Device.DeviceAssociateDeviceType(req, nil)
	logger.Info("DeviceAssociateDeviceType openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		logger.Info("DeviceAssociateDeviceType openapi error:", err.Error())
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
