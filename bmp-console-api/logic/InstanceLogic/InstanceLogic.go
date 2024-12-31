package InstanceLogic

import (
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"path"
	"strconv"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-console-api/constants"
	"coding.jd.com/aidc-bmp/bmp-console-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-console-api/logic/RaidLogic"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/excel"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"

	log "coding.jd.com/aidc-bmp/bmp_log"
	"git.jd.com/cps-golang/ironic-common/exception"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"

	sdkInstanceParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/instance"
	"coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

type Formula struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type ExtraInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func SdkInstance2RespInstance(logger *log.Logger, v sdkModels.Instance) response.Instance {
	/*
		CPU：Intel E5-2650V4（2*12物理核，2.2GHz）
		内存：256GB DDR4
		系统盘：300GB (2*300GB SAS, RAID1)
		数据盘：72TB (12*6TB SATA, NO RAID)
		网卡：2*25GE
		GPU：P40*4
	*/
	//cputpl := "%s %s (%d*%d物理核, %sGHz)"
	//memtpl := "%dGB %s %dGHZ"
	//systemtpl := "%dGB(%dGB %s %s*%d,%s)"
	//datatpl := "%dGB(%dGB %s %s*%d,%s)"
	//nictpl := "%d*%dGE"
	//gputpl := "%s %s*%d"
	language := logger.GetPoint("language").(string)
	//if language == constants.LANGUAGE_EN {
	//	cputpl = "%s %s (%d*%dcores, %sGHz)"
	//}

	//raidMap := make(map[string]response.DeviceTypeRaid) //	请求一次，避免循环请求

	// raid, _ := raidLogic.QueryRaid(logger, v.SystemVolumeRaidID)
	// raidStr := ""
	// if raid != nil {
	// 	raidStr = raid.Name
	// }

	cpu_ := strconv.Itoa(int(v.CPUAmount)) + "*" + strconv.Itoa(int(v.CPUCores)) + "物理核," + v.CPUFrequency + "GHz"
	if language == constants.LANGUAGE_EN {
		cpu_ = strconv.Itoa(int(v.CPUAmount)) + "*" + strconv.Itoa(int(v.CPUCores)) + "cores," + v.CPUFrequency + "GHz"
	}
	mem_ := strconv.Itoa(int(v.MemAmount)*int(v.MemSize)) + "GB(" + strconv.Itoa(int(v.MemSize)) + "GB*" + strconv.Itoa(int(v.MemAmount)) + ")" + v.MemType + " " + strconv.Itoa(int(v.MemFrequency)) + "MHz"
	//系统盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式）
	//240GB 240GB SATA SSD*2 RAID1
	sv_ := getInstanceSystemDiskInfo(logger, v.InstanceVolumeRaids)

	// 显示品牌 显卡型号 个数，例 NVIDIA P40 * 4
	gpu_ := v.GpuManufacturer + " " + v.GpuModel + "*" + strconv.Itoa(int(v.GpuAmount))
	if v.GpuAmount == 0 {
		gpu_ = ""
	}
	nicInfo := strconv.Itoa(int(v.NicAmount)) + "*" + strconv.Itoa(int(v.NicRate)) + "GE"
	return response.Instance{
		Instance:   v,
		CpuInfo:    cpu_,
		MemInfo:    mem_,
		SystemInfo: sv_,
		NicInfo:    nicInfo,
		GpuInfo:    gpu_,
		Volumes:    v.InstanceVolumeRaids,
	}
}

func getInstanceSystemDiskInfo(logger *log.Logger, ivr []*models.InstanceVolumeRaid) string {
	//TODO,返回的是结构
	for _, v := range ivr {
		if v.Volume == nil {
			logger.Warnf("getInstanceSystemDiskInfo.Volume empty, v:%s", util.InterfaceToJson(v))
			continue
		}
		if v.Raid == nil { //没有raid卡，不做raid
			// logger.Warnf("getInstanceSystemDiskInfo.Raid empty, v:%s", util.InterfaceToJson(v))
			// continue
		}
		if v.Volume.VolumeType == "system" {

			vr := &models.VolumeRaids{
				// 硬盘数量（最低块数）
				VolumeAmount: v.Volume.VolumeAmount,

				// 卷uuid
				VolumeID: v.Volume.VolumeID,

				// 卷名称
				VolumeName: v.Volume.VolumeName,

				// 单盘大小（最小容量）
				VolumeSize: v.Volume.VolumeSize,

				// 卷类型：系统卷，数据卷
				VolumeType: v.Volume.VolumeType,

				// 硬盘单位（GB,TB）
				VolumeUnit: v.Volume.VolumeUnit,

				// 硬盘类型（SSD,HDD）
				DiskType: v.Volume.DiskType,

				// 接口类型（SATA,SAS,NVME,不限制）
				InterfaceType: v.Volume.InterfaceType,
			}
			return RaidLogic.GetDiskDetailByRaid(vr, v.Raid)
		}
	}
	return ""
}

func DescribeInstances(logger *log.Logger, param requestTypes.QueryInstancesRequest) (*response.InstancesResponse, error) {
	// language := logger.GetPoint("language").(string)
	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewDescribeProjectInstancesParams()
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithProjectID(&param.ProjectID)
	request.WithDeviceTypeID(&param.DeviceTypeId)
	request.WithStatus(&param.Status)
	request.WithPageNumber(&param.PageNumber)
	request.WithPageSize(&param.PageSize)
	if param.IloIp != "" {
		request.WithIloIP(&param.IloIp)
	}
	if param.PrivateIp != "" {
		request.WithIPV4(&param.PrivateIp)
	}
	if param.PrivateIpv6 != "" {
		request.WithIPV6(&param.PrivateIpv6)
	}

	if param.InstanceId != "" {
		request.WithInstanceID(&param.InstanceId)
	}

	if param.ExportType != "" || param.IsAll == "1" {
		isAll := "1"
		request.IsAll = &isAll
	}

	if param.Name != "" {
		request.WithInstanceName(&param.Name)
	}

	if param.IsInstallAgent != "" {
		request.WithIsInstallAgent(&param.IsInstallAgent)
	}

	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	logger.Info("DescribeProjectInstances openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.DescribeProjectInstances(request, nil)
	logger.Info("DescribeProjectInstances openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeProjectInstances openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	resp := response.InstancesResponse{}
	resp.PagingResponse.PageNumber = responseApi.Payload.Result.PageNumber
	resp.PagingResponse.PageSize = responseApi.Payload.Result.PageSize
	resp.PagingResponse.TotalCount = responseApi.Payload.Result.TotalCount

	obj := response.Instance{}

	for _, v := range responseApi.Payload.Result.Instances {
		obj = SdkInstance2RespInstance(logger, *v)

		resp.Instances = append(resp.Instances, obj)
	}
	return &resp, nil
}

func DescribeInstance(logger *log.Logger, param requestTypes.QueryInstanceRequest) (*response.InstanceInfo, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewDescribeProjectInstanceParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	request.WithInstanceID(param.InstanceId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("DescribeProjectInstance openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.DescribeProjectInstance(request, nil)
	logger.Info("DescribeProjectInstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeProjectInstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	obj := response.Instance{}
	v := responseApi.Payload.Result.Instance

	obj = SdkInstance2RespInstance(logger, *v)

	return &response.InstanceInfo{
		Instance: obj,
	}, nil
}

// func formatDiskDesc(logger *log.Logger, deviceType string, raidId string, volumnType string) string {

// 	region := logger.GetPoint("region").(string)
// 	pin := logger.GetPoint("pin").(string)
// 	req := requestTypes.QueryRaidsRequest{
// 		Region:     region,
// 		DeviceType: deviceType,
// 	}
// 	raidsResult, err := RaidLogic.DescribeDeviceRaids(logger, req, pin)
// 	if err != nil {
// 		logger.Warnf("RaidLogic.DescribeDeviceRaids error:%s", err.Error())
// 		return ""
// 	}
// 	raids := raidsResult.Raids
// 	for _, raid := range raids {
// 		if raid.RaidID == raidId && strings.EqualFold(volumnType, raid.VolumeType) {
// 			volume_detail := strings.TrimRight(raid.VolumeDetail, ")")
// 			raidType := strings.ReplaceAll(raid.RaidID, "NORAID", "NO RAID")
// 			return fmt.Sprintf("%s, %s)", volume_detail, raidType)
// 		}
// 	}
// 	return ""
// }

func StartInstance(logger *log.Logger, param requestTypes.StartInstanceRequest) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewStartProjectInstanceParams()
	request.WithInstanceID(param.InstanceId)
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("StartProjectInstance openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.StartProjectInstance(request, nil)
	logger.Info("StartProjectInstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("StartProjectInstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

func StopInstance(logger *log.Logger, param requestTypes.StopInstanceRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewStopProjectInstanceParams()
	request.WithInstanceID(param.InstanceId)
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("stopinstance openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.StopProjectInstance(request, nil)
	logger.Info("stopinstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("stopinstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

func RestartInstance(logger *log.Logger, param requestTypes.RestartInstanceRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewRestartProjectInstanceParams()
	request.WithInstanceID(param.InstanceId)
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("restartinstance openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.RestartProjectInstance(request, nil)
	logger.Info("restartinstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("restartinstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

func ModifyInstance(logger *log.Logger, param requestTypes.ModifyInstanceRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewModifyProjectInstanceParams()
	request.WithInstanceID(param.InstanceID)
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	body := sdkModels.ModifyInstanceRequest{
		Description:  param.Description,
		InstanceName: param.Name,
	}
	request.WithBody(&body)
	logger.Info("ModifyProjectInstance openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.ModifyProjectInstance(request, nil)
	logger.Info("ModifyProjectInstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("ModifyProjectInstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

func DeleteInstance(logger *log.Logger, instanceId string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewDeleteProjectInstanceParams()
	request.WithInstanceID(instanceId)
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("deleteinstance openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.DeleteProjectInstance(request, nil)
	logger.Info("deleteinstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("restartinstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

func LockInstance(logger *log.Logger, param requestTypes.LockInstanceRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewLockProjectInstanceParams()
	request.WithInstanceID(param.InstanceId)
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("lockinstance openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.LockProjectInstance(request, nil)
	logger.Info("lockinstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("lockinstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

func UnLockInstance(logger *log.Logger, param requestTypes.UnLockInstanceRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewUnLockProjectInstanceParams()
	request.WithInstanceID(param.InstanceId)
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("unlockinstance openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.UnLockProjectInstance(request, nil)
	logger.Info("unlockinstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("lockinstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

//CreateInstanceApi 测试sdk直接创建使用
func CreateInstanceApi(logger *log.Logger, param *sdkModels.CreateInstanceRequest) ([]string, error) {

	logid := logger.GetPoint("logid").(string)

	req := sdkInstanceParameters.NewCreateProjectInstanceParams()
	req.WithTraceID(logid)
	req.WithBody(param)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("createinstancesApi openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Instance.CreateProjectInstance(req, nil)
	logger.Info("createinstancesApi openapi response:", util.InterfaceToJson(responseApi))
	fmt.Println("createinstancesApi openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("createinstancesApi openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.InstanceIds, nil
}

func getOrCreate(str string) string {
	if str != "" {
		return str
	}
	return commonUtil.GenerateRandUuid()
}

func ExportInstance(logger *log.Logger, result []response.Instance) (fileName, downLoadFileName string) {

	language := logger.GetPoint("language").(string)

	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)

	//cputpl := "%s %s (%d*%d 物理核, %s)"
	//if language == "en_US" {
	//	cputpl = "%s %s (%d*%d cores, %s)"
	//}
	//memtpl := "%d*%d %s"

	for _, data := range result {

		// interfaceMode := BaseLogic.InterfaceMode[data.InterfaceMode]
		//var instanceStatus string
		//if language == "en_US" {
		//	// interfaceMode = BaseLogic.InterfaceModeEn[data.InterfaceMode]
		//	instanceStatus = BaseLogic.InstanceStatusEn[data.Status]
		//} else {
		//	instanceStatus = BaseLogic.InstanceStatus[data.Status]
		//}
		var info string
		//cpu_ := strconv.Itoa(int(data.CPUAmount)) + "*" + strconv.Itoa(int(data.CPUCores)) + " 物理核，" + data.CPUFrequency + "GHz"
		//mem_ := strconv.Itoa(int(data.MemAmount)*int(data.MemSize)) + "GB (" + strconv.Itoa(int(data.MemSize)) + "*" + strconv.Itoa(int(data.MemAmount)) + ") " + data.MemType + "，" + strconv.Itoa(int(data.MemFrequency)) + "MHz"
		info = fmt.Sprintf("CPU:%s, 内存:%s", data.CpuInfo, data.MemInfo)
		if language == constants.LANGUAGE_EN {
			//cpu_ = strconv.Itoa(int(data.CPUAmount)) + "*" + strconv.Itoa(int(data.CPUCores)) + " cores, " + data.CPUFrequency + "GHz"
			info = fmt.Sprintf("CPU:%s, Memory:%s", data.CpuInfo, data.MemInfo)
		}
		//cpuInfo := fmt.Sprintf(cputpl, data.CPUManufacturer, data.CPUModel, data.CPUAmount, data.CPUCores, data.CPUFrequency)
		//memInfo := fmt.Sprintf(memtpl, data.MemAmount, data.MemSize, data.MemType)

		sheetData = append(sheetData, []string{
			fmt.Sprintf("%s/%s", data.InstanceName, data.InstanceID),
			data.StatusName,
			data.DeviceTypeName,
			data.ImageName,
			data.IloIP,
			data.PrivateIPV4,
			data.PrivateEth1IPV4,
			data.PrivateIPV6,
			data.PrivateEth1IPV6,
			info,
			data.CreatedTime,
		})
	}
	exportInstanceHeader := BaseLogic.ExportInstanceHeader
	if language == "en_US" {
		exportInstanceHeader = BaseLogic.ExportInstanceHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportInstanceHeader, sheetData); err != nil {
		panic(exception.Exception{Msg: "新建excel失败", Status: BaseLogic.ERROR_CODE})
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	downloadFileName := "instance_list_" + time.Now().Format(BaseLogic.FormatDate) + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx" //生成100000到100000+900000之间的随机数，左闭右开
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Info("保存excel到本地失败" + err.Error())
		panic(exception.Exception{Msg: "保存excel到本地失败", Status: BaseLogic.ERROR_CODE})
	}
	return fileName, downloadFileName
}
func ParseFile(logger *log.Logger, param requestTypes.ParseFileRequest, header *multipart.FileHeader) string {
	if header.Size > 16*1024 {
		panic(exception.Exception{Msg: BaseLogic.FileError[BaseLogic.UPLOAD_FILE_SIZE_ERROR], ErrorCode: BaseLogic.UPLOAD_FILE_SIZE_ERROR})
	}
	arr := strings.Split(header.Filename, ".")
	suffix := strings.ToLower(arr[len(arr)-1])
	AvailableSuffix := []string{"jpg", "jpeg", "png", "gif"}
	if util.InArray(suffix, AvailableSuffix) {
		panic(exception.Exception{Msg: "file is image", Status: BaseLogic.ERROR_CODE})
	}
	dir, err := util.GetBaseDir()
	if err != nil {
		logger.Info("获取当前根目录失败" + err.Error())
		panic(exception.Exception{Msg: "获取当前根目录失败", Status: BaseLogic.ERROR_CODE})
	}
	prefix := util.GenUuid()
	excel.CreateDirIfNotExist(BaseLogic.LOCAL_FILE_PATH) //如果目录不存在，创建一个目录
	localFileName := path.Join(dir, BaseLogic.LOCAL_FILE_PATH, prefix+"-"+header.Filename)
	if err = util.SaveToLocal(localFileName, param.UserDataFile); err != nil {
		logger.Info("保存excel到本地失败" + err.Error())
		panic(exception.Exception{Msg: "保存文件到本地失败" + err.Error(), Status: BaseLogic.ERROR_CODE})
	}
	contentFile, err := util.ReadFileContent(localFileName)
	content := ""
	if err != nil {
		logger.Info("读取文件失败" + err.Error())
		panic(exception.Exception{Msg: "读取文件失败", Status: BaseLogic.ERROR_CODE})
	}
	contentByte, err := util.Base64Decrypt(string(contentFile))
	if err != nil && param.Base64 { //如果选了64，结果内容不是64
		panic(exception.Exception{Msg: BaseLogic.FileError[BaseLogic.UPLOAD_FILE_NOT_BASE64], ErrorCode: BaseLogic.UPLOAD_FILE_NOT_BASE64})
	}
	if err == nil && !param.Base64 { //如果内容是64，结果还没有勾选64
		panic(exception.Exception{Msg: BaseLogic.FileError[BaseLogic.UPLOAD_FILE_BASE64], ErrorCode: BaseLogic.UPLOAD_FILE_BASE64})
	}
	if err != nil { //说明不是64位编码，原始代码
		content = string(contentFile)
	} else {
		content = string(contentByte)
	}
	//fmt.Println("qq", strings.Index(content, "#!/bin/bash"))
	//必须在文件第一行的开头是如下内容才能通过
	if strings.Index(content, "#!/bin/bash") != 0 && strings.Index(content, "#!/usr/bin/env python") != 0 && strings.Index(content, "(rem cmd|<powershell>)") != 0 {
		panic(exception.Exception{Msg: BaseLogic.FileError[BaseLogic.UPLOAD_FILE_FORMAT_ERROR], ErrorCode: BaseLogic.UPLOAD_FILE_FORMAT_ERROR})
	}
	return string(contentFile)
}
func Trim(str string) string {
	s := strings.TrimRight(str, "0")
	s = strings.TrimRight(s, ".")
	return s
}

// 【设备管理】【批量开机】
func StartInstances(logger *log.Logger, instanceIDs []string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewStartInstancesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.StartInstancesRequest{
		InstanceIds: instanceIDs,
	}
	req.WithBody(&body)

	logger.Info("StartInstances openapi sdk req:", util.InterfaceToJson(req))
	client, err := openApi.SdkClient.Instance.StartInstances(req, nil)
	logger.Info("StartInstances openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return true, nil
}

// 【设备管理】【批量关机】
func StopInstances(logger *log.Logger, instanceIDs []string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewStopInstancesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.StopInstancesRequest{
		InstanceIds: instanceIDs,
	}
	req.WithBody(&body)

	logger.Info("StopInstances openapi sdk req:", util.InterfaceToJson(req))
	client, err := openApi.SdkClient.Instance.StopInstances(req, nil)
	logger.Info("StopInstances openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return true, nil
}

// 【设备管理】【批量重启】
func RestartInstances(logger *log.Logger, instanceIDs []string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewRestartInstancesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.RestartInstancesRequest{
		InstanceIds: instanceIDs,
	}
	req.WithBody(&body)

	logger.Info("RestartInstances openapi sdk req:", util.InterfaceToJson(req))
	client, err := openApi.SdkClient.Instance.RestartInstances(req, nil)
	logger.Info("RestartInstances openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return true, nil
}

// 【设备管理】【批量编辑实例名称】
func ModifyInstances(logger *log.Logger, instanceIDs []string, instanceName string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewModifyInstancesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.ModifyInstancesRequest{
		InstanceIDs:  instanceIDs,
		InstanceName: instanceName,
	}
	req.WithBody(&body)

	logger.Info("ModifyInstances openapi sdk req:", util.InterfaceToJson(req))
	client, err := openApi.SdkClient.Instance.ModifyInstances(req, nil)
	logger.Info("ModifyInstances openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return true, nil
}

// 【设备管理】【批量删除】
func DeleteInstances(logger *log.Logger, instanceIDs []string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewDeleteInstancesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.DeleteInstancesRequest{
		InstanceIDs: instanceIDs,
	}
	req.WithBody(&body)

	logger.Info("DeleteInstances openapi sdk req:", util.InterfaceToJson(req))
	client, err := openApi.SdkClient.Instance.DeleteInstances(req, nil)
	logger.Info("DeleteInstances openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return true, nil
}

func ResetInstancePassword(logger *log.Logger, instanceId string, password string) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewResetProjectInstancePasswdParams()
	request.WithInstanceID(instanceId)
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	body := sdkModels.ResetInstancePasswdRequest{
		Password: password,
	}
	request.WithBody(&body)

	logger.Info("ResetProjectInstancePasswd openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.ResetProjectInstancePasswd(request, nil)
	logger.Info("ResetProjectInstancePasswd openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("ResetProjectInstancePasswd openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

func ResetInstancesPassword(logger *log.Logger, instanceIds []string, password string) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewResetProjectInstancesPasswdParams()
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	body := sdkModels.ResetInstancesPasswdRequest{
		Password:    password,
		InstanceIDs: instanceIds,
	}
	request.WithBody(&body)

	logger.Info("ResetProjectInstancesPasswd openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.ResetProjectInstancesPasswd(request, nil)
	logger.Info("ResetProjectInstancesPasswd openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("ResetProjectInstancesPasswd openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

//ReinstallInstance sdk重装
func ReinstallInstance(logger *log.Logger, param requestTypes.ReinstallInstanceRequest) (bool, error) {

	logid := logger.GetPoint("logid").(string)

	req := sdkInstanceParameters.NewReinstallProjectInstanceParams()
	req.WithTraceID(logid)
	req.WithBody(&param.ReinstallInstanceRequest)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithInstanceID(param.InstanceId)
	logger.Info("ReinstallProjectInstance openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Instance.ReinstallProjectInstance(req, nil)
	logger.Info("ReinstallProjectInstance openapi response:", util.InterfaceToJson(responseApi))
	fmt.Println("ReinstallProjectInstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("ReinstallProjectInstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Success, nil
}

func DescribeInstancesForShare(logger *log.Logger, param requestTypes.QueryInstancesForShareRequest) ([]*sdkModels.InstanceForShare, error) {
	// language := logger.GetPoint("language").(string)
	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewDescribeInstancesByProjectIDAndOwnerNameAndSharerNameParams()
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithProjectID(&param.ProjectID)
	request.WithOwnerName(&param.OwnerName)
	request.WithSharerName(&param.SharerName)

	if param.InstanceId != "" {
		request.WithInstanceID(&param.InstanceId)
	}
	if param.InstanceName != "" {
		request.WithInstanceName(&param.InstanceName)
	}

	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	logger.Info("DescribeInstancesByProjectIDAndOwnerNameAndSharerName openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.DescribeInstancesByProjectIDAndOwnerNameAndSharerName(request, nil)
	logger.Info("DescribeInstancesByProjectIDAndOwnerNameAndSharerName openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeInstancesByProjectIDAndOwnerNameAndSharerName openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Instances, nil
}
