package deviceTypeLogic

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/diskDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rVolumeRaidDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/volumeDao"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceTypeDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/idcDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/imageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/osDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rDeviceTypeImageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rDeviceTypeImagePartitionDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/raidDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/imageLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

var K int = 1

func Create(logger *log.Logger, param *requestTypes.CreateDeviceTypeRequest) (string, error) {
	_, err := idcDao.GetIdcById(logger, param.IDcID)
	if err != nil {
		logger.Warn("GetIdcByUuid sql error:", param.IDcID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs("机房id不存在", "idcId not exist"))
	}
	//入库机型
	deviceTypeList, _ := deviceTypeDao.QueryAllDeviceTypes(logger, map[string]interface{}{
		"name":   param.Name,
		"is_del": baseLogic.IS_NOT_DEL,
	})
	if len(deviceTypeList) != 0 {
		panic(constant.BuildInvalidArgumentWithArgs("机型名称已存在", "deviceType name exist"))
	}
	//入库机型规格
	deviceTypeList, _ = deviceTypeDao.QueryAllDeviceTypes(logger, map[string]interface{}{
		"device_type": param.DeviceType,
		"is_del":      baseLogic.IS_NOT_DEL,
	})
	if len(deviceTypeList) != 0 {
		panic(constant.BuildInvalidArgumentWithArgs("机型规格已存在", "deviceType exist"))
	}
	entity := BuildDeviceTypeEntity(logger, param)
	if err := addVolume(logger, entity.DeviceTypeID, param); err != nil {
		logger.Warn("AddVolume error:", param.Volumes, err.Error())
		return "", err
	}
	//入库关联raid信息，表 r_device_type_raid，ToDo   1.2标准版以后不再这么关联机型和raid  而是改为了卷和raid的关联
	//if err := addDeviceTypeRaid(logger, entity.DeviceTypeID, param.RaidID, param.SystemVolumeUnit, param.SystemVolumeType, param.SystemVolumeInterfaceType, param.SystemVolumeSize, param.SystemVolumeAmount); err != nil {
	//	logger.Warn("AddDeviceTypeRaid sql error:", param.DeviceType, err.Error())
	//	return "", err
	//}
	if _, err := deviceTypeDao.AddDeviceType(logger, entity); err != nil {
		logger.Warn("AddDeviceType sql error:", param.DeviceType, err.Error())
		return "", err
	}
	return entity.DeviceTypeID, nil
}
func addVolume(logger *log.Logger, deviceTypeId string, deviceType *requestTypes.CreateDeviceTypeRequest) error {
	for _, v := range deviceType.Volumes {
		volume := &volumeDao.Volume{
			VolumeID:      util.GetUuid("vl-"),
			DeviceTypeID:  deviceTypeId,
			VolumeName:    v.VolumeName,
			VolumeType:    v.VolumeType,
			DiskType:      v.DiskType,
			InterfaceType: v.InterfaceType,
			VolumeSize:    v.VolumeSize,
			VolumeUnit:    v.VolumeUnit,
			VolumeAmount:  v.VolumeAmount,
			//RaidCan:       v.RaidCan,
			//Raid:          v.Raid,

			CreatedBy:   logger.GetPoint("username").(string),
			CreatedTime: int(time.Now().Unix()),
			UpdatedBy:   logger.GetPoint("username").(string),
			UpdatedTime: int(time.Now().Unix()),
		}
		if _, err := volumeDao.AddVolume(logger, volume); err != nil {
			logger.Warn("AddVolume sql error:", err.Error())
			return err
		}
		if v.Raid != "" {
			raidIdList := strings.Split(v.Raid, ",")
			for _, raidId := range raidIdList {
				raidInfo, err := raidDao.GetRaidById(logger, raidId)
				if err != nil {
					//fmt.Println(raidInfo == nil, err.Error())
					logger.Warn("raid sql error", err.Error())
					continue
				}
				rVolumeRaid := &rVolumeRaidDao.RVolumeRaid{
					VolumeID:     volume.VolumeID,
					VolumeType:   v.VolumeType,
					DeviceTypeID: deviceTypeId,
					RaidCan:      v.RaidCan,
					RaidID:       raidId,
					RaidName:     raidInfo.Name,
					CreatedBy:    logger.GetPoint("username").(string),
					CreatedTime:  int(time.Now().Unix()),
					UpdatedBy:    logger.GetPoint("username").(string),
					UpdatedTime:  int(time.Now().Unix()),
				}
				if _, err := rVolumeRaidDao.AddRVolumeRaid(logger, rVolumeRaid); err != nil {
					logger.Warn("Add rVolumeRaidDao sql error:", err.Error())
					return err
				}
			}
		} else {
			rVolumeRaid := &rVolumeRaidDao.RVolumeRaid{
				VolumeID:     volume.VolumeID,
				VolumeType:   v.VolumeType,
				DeviceTypeID: deviceTypeId,
				RaidCan:      v.RaidCan,
				RaidID:       "",
				RaidName:     "",
				CreatedBy:    logger.GetPoint("username").(string),
				CreatedTime:  int(time.Now().Unix()),
				UpdatedBy:    logger.GetPoint("username").(string),
				UpdatedTime:  int(time.Now().Unix()),
			}
			if _, err := rVolumeRaidDao.AddRVolumeRaid(logger, rVolumeRaid); err != nil {
				logger.Warn("Add rVolumeRaidDao sql error:", err.Error())
				return err
			}
		}
	}
	return nil
}
func modifyVolume(logger *log.Logger, deviceTypeId string, deviceType *requestTypes.ModifyDeviceTypeRequest) error {
	for _, v := range deviceType.Volumes {
		volume := &volumeDao.Volume{
			VolumeID:      util.GetUuid("vl-"),
			DeviceTypeID:  deviceTypeId,
			VolumeName:    v.VolumeName,
			VolumeType:    v.VolumeType,
			DiskType:      v.DiskType,
			InterfaceType: v.InterfaceType,
			VolumeSize:    v.VolumeSize,
			VolumeUnit:    v.VolumeUnit,
			VolumeAmount:  v.VolumeAmount,
			//RaidCan:       v.RaidCan,
			//Raid:          v.Raid,

			CreatedBy:   logger.GetPoint("username").(string),
			CreatedTime: int(time.Now().Unix()),
			UpdatedBy:   logger.GetPoint("username").(string),
			UpdatedTime: int(time.Now().Unix()),
		}
		if _, err := volumeDao.AddVolume(logger, volume); err != nil {
			logger.Warn("AddVolume sql error:", err.Error())
			return err
		}
		if v.Raid != "" {
			raidIdList := strings.Split(v.Raid, ",")
			for _, raidId := range raidIdList {
				raidInfo, err := raidDao.GetRaidById(logger, raidId)
				if err != nil {
					//fmt.Println(raidInfo == nil, err.Error())
					logger.Warn("raid sql error", err.Error())
					continue
				}
				rVolumeRaid := &rVolumeRaidDao.RVolumeRaid{
					VolumeID:     volume.VolumeID,
					VolumeType:   v.VolumeType,
					DeviceTypeID: deviceTypeId,
					RaidCan:      v.RaidCan,
					RaidID:       raidId,
					RaidName:     raidInfo.Name,
					CreatedBy:    logger.GetPoint("username").(string),
					CreatedTime:  int(time.Now().Unix()),
					UpdatedBy:    logger.GetPoint("username").(string),
					UpdatedTime:  int(time.Now().Unix()),
				}
				if _, err := rVolumeRaidDao.AddRVolumeRaid(logger, rVolumeRaid); err != nil {
					logger.Warn("Add rVolumeRaidDao sql error:", err.Error())
					return err
				}
			}
		} else {
			rVolumeRaid := &rVolumeRaidDao.RVolumeRaid{
				VolumeID:     volume.VolumeID,
				VolumeType:   v.VolumeType,
				DeviceTypeID: deviceTypeId,
				RaidCan:      v.RaidCan,
				RaidID:       "",
				RaidName:     "",
				CreatedBy:    logger.GetPoint("username").(string),
				CreatedTime:  int(time.Now().Unix()),
				UpdatedBy:    logger.GetPoint("username").(string),
				UpdatedTime:  int(time.Now().Unix()),
			}
			if _, err := rVolumeRaidDao.AddRVolumeRaid(logger, rVolumeRaid); err != nil {
				logger.Warn("Add rVolumeRaidDao sql error:", err.Error())
				return err
			}
		}
	}
	return nil
}

func BuildDeviceTypeEntity(logger *log.Logger, param *requestTypes.CreateDeviceTypeRequest) *deviceTypeDao.DeviceType {
	//var systemVolumeSize, dataVolumeSize int
	//if param.SystemVolumeUnit == "TB" {
	//	systemVolumeSize = param.SystemVolumeSize * K
	//} else {
	//	systemVolumeSize = param.SystemVolumeSize
	//}
	//if param.DataVolumeUnit == "TB" {
	//	dataVolumeSize = param.DataVolumeSize * K
	//} else {
	//	dataVolumeSize = param.DataVolumeSize
	//}
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
	//if param.DataVolumeType == "" {
	//	if param.DataVolumeInterfaceType != "" || param.DataVolumeAmount != 0 || param.DataVolumeSize != 0 {
	//		panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
	//	}
	//}
	//if param.DataVolumeInterfaceType == "" {
	//	if param.DataVolumeType != "" || param.DataVolumeAmount != 0 || param.DataVolumeSize != 0 {
	//		panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
	//	}
	//}
	//if param.DataVolumeAmount == 0 {
	//	if param.DataVolumeInterfaceType != "" || param.DataVolumeType != "" || param.DataVolumeSize != 0 {
	//		panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
	//	}
	//}
	//if param.DataVolumeSize == 0 {
	//	if param.DataVolumeInterfaceType != "" || param.DataVolumeAmount != 0 || param.DataVolumeType != "" {
	//		panic(constant.BuildInvalidArgumentWithArgs("数据盘不合法", "data volume invalidate"))
	//	}
	//}

	var bootmode string
	if param.BootMode == "" {
		//用户缺失此参数时，如果x86的默认为bios
		if param.Architecture == "x86_64" {
			bootmode = "Legacy/BIOS"
		} else {
			bootmode = "UEFI"
		}
	} else {
		bootmode = param.BootMode
	}

	return &deviceTypeDao.DeviceType{
		IDcID:        param.IDcID,
		DeviceTypeID: util.GetUuid("dt-"),
		DeviceType:   param.DeviceType,
		Name:         param.Name,
		DeviceSeries: param.DeviceSeries,
		Height:       param.Height,
		Architecture: param.Architecture,
		Description:  param.Description,

		CPUAmount:       param.CPUAmount,
		CPUCores:        param.CPUCores,
		CPUFrequency:    param.CPUFrequency,
		CPUManufacturer: param.CPUManufacturer,
		CPUModel:        param.CPUModel,

		MemAmount:       param.MemAmount,
		MemSize:         param.MemSize,
		MemFrequency:    param.MemFrequency,
		MemType:         param.MemType,
		NicAmount:       param.NicAmount,
		NicRate:         param.NicRate,
		InterfaceMode:   param.InterfaceMode,
		GpuAmount:       param.GpuAmount,
		GpuManufacturer: param.GpuManufacturer,
		GpuModel:        param.GpuModel,

		//SystemVolumeType:          param.SystemVolumeType,
		//SystemVolumeInterfaceType: param.SystemVolumeInterfaceType,
		//SystemVolumeAmount:        param.SystemVolumeAmount,
		//RaidCan:                   param.RaidCan,
		//SystemVolumeSize:          systemVolumeSize,
		//SystemVolumeUnit:          param.SystemVolumeUnit,
		//
		//DataVolumeType:          param.DataVolumeType,
		//DataVolumeInterfaceType: param.DataVolumeInterfaceType,
		//DataVolumeAmount:        param.DataVolumeAmount,
		//DataVolumeSize:          dataVolumeSize,
		//DataVolumeUnit:          param.DataVolumeUnit,
		CreatedBy:   logger.GetPoint("username").(string),
		UpdatedBy:   logger.GetPoint("username").(string),
		CreatedTime: int(time.Now().Unix()),
		UpdatedTime: int(time.Now().Unix()),
		CpuSpec:     param.CpuSpec,
		MemSpec:     param.MemSpec,
		BootMode:    bootmode,
		IsNeedRaid:  param.IsNeedRaid,
	}
}
func QueryDeviceTypes(logger *log.Logger, param *requestTypes.QueryDeviceTypesRequest, p util.Pageable) ([]*responseTypes.DeviceType, int64, error) {
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"idc_id": param.IdcID,
		"is_del": 0,
	}
	if param.Name != "" {
		query["name.like"] = "%" + param.Name + "%"
	}
	if param.DeviceType != "" {
		query["device_type.in"] = strings.Split(param.DeviceType, ",")
	}
	if param.DeviceTypeID != "" {
		query["device_type_id.in"] = strings.Split(param.DeviceTypeID, ",")
	}
	if param.DeviceSeries != "" {
		query["device_series.in"] = strings.Split(param.DeviceSeries, ",")
	}

	count, err := deviceTypeDao.QueryDeviceTypesCount(logger, query)
	if err != nil {
		logger.Warnf("DescribeDeviceTypes_GetDeviceTypeCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	dts := []*deviceTypeDao.DeviceType{}
	if param.IsAll == baseLogic.IS_ALL {
		dts, err = deviceTypeDao.QueryAllDeviceTypes(logger, query)
	} else {
		dts, err = deviceTypeDao.QueryDeviceTypes(logger, query, offset, limit)
	}
	if err != nil {
		logger.Warnf("QueryDeviceTypes sql err, param:%s, error:%s", util.ObjToJson(param), err.Error())
		return nil, 0, err
	}
	res := []*responseTypes.DeviceType{}
	deviceTypeIds := []string{}
	for _, value := range dts {
		deviceTypeIds = append(deviceTypeIds, value.DeviceTypeID)
	}

	deviceTypeId2Stock := map[string]int{}
	if len(deviceTypeIds) != 0 {
		deviceTypeId2Stock, err = deviceLogic.GetStockByDeviceTypeIds(logger, deviceTypeIds)
		if err != nil {
			logger.Warnf("GetStockByDeviceTypeIds error:%s", err.Error())
		}
	}

	for _, value := range dts {
		v := DeviceTypeEntity2DeviceType(logger, value)
		stock, ok := deviceTypeId2Stock[v.DeviceTypeID]
		if ok {
			v.StockAvailable = stock
		} else {
			v.StockAvailable = 0
		}
		res = append(res, v)
	}
	return res, count, err
}
func QueryDeviceType(logger *log.Logger, param *requestTypes.QueryDeviceTypesRequest) (*responseTypes.DeviceType, error) {
	query := map[string]interface{}{
		"is_del": 0,
	}
	//if param.Id != 0 {
	//	query["id"] = param.Id
	//}
	if param.DeviceType != "" {
		query["device_type"] = param.DeviceType
	}
	if param.DeviceTypeID != "" {
		query["device_type_id"] = param.DeviceTypeID
	}
	dts, err := deviceTypeDao.GetOneDeviceType(logger, query)
	if err != nil {
		logger.Warnf("QueryDeviceType sql err, param:%s, error:%s", util.ObjToJson(param), err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	res := DeviceTypeEntity2DeviceType(logger, dts)
	return res, err
}
func QueryDeviceTypeById(logger *log.Logger, deviceTypeId string) (*responseTypes.DeviceType, error) {

	deviceType, err := deviceTypeDao.GetDeviceTypeById(logger, deviceTypeId)
	if err != nil {
		logger.Warnf("QueryDeviceType sql err, deviceTypeId:%s, error:%s", deviceTypeId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	res := DeviceTypeEntity2DeviceType(logger, deviceType)
	return res, err
}
func ModifyDeviceType(logger *log.Logger, param *requestTypes.ModifyDeviceTypeRequest, deviceTypeID string) error {

	entity, err := deviceTypeDao.GetOneDeviceType(logger, map[string]interface{}{
		"device_type_id": deviceTypeID,
		"is_del":         baseLogic.IS_NOT_DEL,
	})
	if err != nil {
		logger.Warn("GetDeviceTypeByUuid sql error:", deviceTypeID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if param.IDcID != nil {
		_, err = idcDao.GetIdcById(logger, *param.IDcID)
		if err != nil {
			logger.Warn("GetIdcByUuid sql error:", *param.IDcID, err.Error())
			panic(constant.BuildInvalidArgumentWithArgs("机房id不存在", "idcId not exist"))
		}
	}
	//入库机型名称
	if param.Name != nil {
		deviceTypeName, _ := deviceTypeDao.GetOneDeviceType(logger, map[string]interface{}{
			"name":   *param.Name,
			"is_del": baseLogic.IS_NOT_DEL,
		})
		if deviceTypeName != nil && deviceTypeID != deviceTypeName.DeviceTypeID {
			panic(constant.BuildInvalidArgumentWithArgs("机型名称已存在", "deviceType name exist"))
		}
	}
	//入库机型规格
	if param.DeviceType != nil {
		deviceType, _ := deviceTypeDao.GetOneDeviceType(logger, map[string]interface{}{
			"device_type": *param.DeviceType,
			"is_del":      baseLogic.IS_NOT_DEL,
		})
		if deviceType != nil && deviceTypeID != deviceType.DeviceTypeID {
			panic(constant.BuildInvalidArgumentWithArgs("机型规格已存在", "deviceType exist"))
		}
	}

	//
	//if param.SystemVolumeType == nil || !util.InArray(*param.SystemVolumeType, []string{"SSD", "HDD"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("SystemVolumeType 不合法", "SystemVolumeType invalidate"))
	//}
	//if param.SystemVolumeInterfaceType == nil || !util.InArray(*param.SystemVolumeInterfaceType, []string{"SATA", "SAS", "NVME"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("SystemVolumeInterfaceType 不合法", "SystemVolumeInterfaceType invalidate"))
	//}
	//if param.DataVolumeType != nil && !util.InArray(*param.DataVolumeType, []string{"SSD", "HDD"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("DataVolumeType 不合法", "DataVolumeType invalidate"))
	//}
	//if param.DataVolumeInterfaceType != nil && !util.InArray(*param.DataVolumeInterfaceType, []string{"SATA", "SAS", "NVME"}) {
	//	panic(constant.BuildInvalidArgumentWithArgs("DataVolumeInterfaceType 不合法", "DataVolumeInterfaceType invalidate"))
	//}
	//var systemVolumeSize, dataVolumeSize int
	//if param.SystemVolumeUnit != nil && *param.SystemVolumeUnit == "TB" {
	//	systemVolumeSize = *param.SystemVolumeSize * K
	//} else {
	//	if param.SystemVolumeSize != nil {
	//		systemVolumeSize = *param.SystemVolumeSize
	//	}
	//}
	//if param.DataVolumeUnit != nil && *param.DataVolumeUnit == "TB" {
	//	dataVolumeSize = *param.DataVolumeSize * K
	//} else {
	//	if param.DataVolumeSize != nil {
	//		dataVolumeSize = *param.DataVolumeSize
	//	}
	//}
	if param.IDcID != nil {
		entity.IDcID = *param.IDcID
	}
	if param.DeviceType != nil {
		entity.DeviceType = *param.DeviceType
	}
	if param.Name != nil {
		entity.Name = *param.Name
	}
	if param.DeviceSeries != nil {
		entity.DeviceSeries = *param.DeviceSeries
	}
	if param.Height != nil {
		entity.Height = *param.Height
	}
	if param.Architecture != nil {
		entity.Architecture = *param.Architecture
	}
	if param.Description != nil {
		entity.Description = *param.Description
	}
	if param.CPUAmount != nil {
		entity.CPUAmount = *param.CPUAmount
	}
	if param.CPUCores != nil {
		entity.CPUCores = *param.CPUCores
	}
	if param.CPUFrequency != nil {
		entity.CPUFrequency = *param.CPUFrequency
	}
	if param.CPUManufacturer != nil {
		entity.CPUManufacturer = *param.CPUManufacturer
	}
	if param.CPUModel != nil {
		entity.CPUModel = *param.CPUModel
	}
	if param.MemAmount != nil {
		entity.MemAmount = *param.MemAmount
	}
	if param.MemSize != nil {
		entity.MemSize = *param.MemSize
	}
	if param.MemFrequency != nil {
		entity.MemFrequency = *param.MemFrequency
	}
	if param.MemType != nil {
		entity.MemType = *param.MemType
	}
	if param.NicAmount != nil {
		entity.NicAmount = *param.NicAmount
	}
	if param.NicRate != nil {
		entity.NicRate = *param.NicRate
	}
	if param.InterfaceMode != nil {
		entity.InterfaceMode = *param.InterfaceMode
	}
	if param.GpuAmount != nil {
		entity.GpuAmount = *param.GpuAmount
	}
	if param.GpuManufacturer != nil {
		entity.GpuManufacturer = *param.GpuManufacturer
	}
	if param.GpuModel != nil {
		entity.GpuModel = *param.GpuModel
	}
	//if param.SystemVolumeType != nil {
	//	entity.SystemVolumeType = *param.SystemVolumeType
	//}
	//if param.SystemVolumeInterfaceType != nil {
	//	entity.SystemVolumeInterfaceType = *param.SystemVolumeInterfaceType
	//}
	//if param.SystemVolumeAmount != nil {
	//	entity.SystemVolumeAmount = *param.SystemVolumeAmount
	//}
	//if param.SystemVolumeSize != nil {
	//	entity.SystemVolumeSize = systemVolumeSize
	//}
	//if param.SystemVolumeUnit != nil {
	//	entity.SystemVolumeUnit = *param.SystemVolumeUnit
	//}

	//if param.RaidCan != nil {
	//	entity.RaidCan = *param.RaidCan
	//}

	//if param.DataVolumeType != nil {
	//	entity.DataVolumeType = *param.DataVolumeType
	//}
	//if param.DataVolumeInterfaceType != nil {
	//	entity.DataVolumeInterfaceType = *param.DataVolumeInterfaceType
	//}
	//if param.DataVolumeAmount != nil {
	//	entity.DataVolumeAmount = *param.DataVolumeAmount
	//}
	//if param.DataVolumeSize != nil {
	//	entity.DataVolumeSize = dataVolumeSize
	//}
	//if param.DataVolumeUnit != nil {
	//	entity.DataVolumeUnit = *param.DataVolumeUnit
	//}
	if param.CpuSpec != nil {
		entity.CpuSpec = *param.CpuSpec
	}
	if param.MemSpec != nil {
		entity.MemSpec = *param.MemSpec
	}
	if param.BootMode != "" {
		entity.BootMode = param.BootMode
	}
	if param.IsNeedRaid != nil {
		entity.IsNeedRaid = *param.IsNeedRaid
	}
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())

	if err := deviceTypeDao.UpdateDeviceTypeByDeviceTypeID(logger, entity, deviceTypeID); err != nil {
		logger.Warnf("ModifyDeviceType UpdateDeviceTypeById sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return err
	}
	//接下来修改卷信息
	//先删掉之前的卷
	if err := volumeDao.UpdateVolumeByWhere(logger, map[string]interface{}{
		"device_type_id": deviceTypeID,
		"is_del":         baseLogic.IS_NOT_DEL,
	}, &volumeDao.Volume{
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
		DeletedTime: int(time.Now().Unix()),
		IsDel:       baseLogic.IS_DEL,
	}); err != nil {
		logger.Warnf("del Volume error deviceTypeId:%s", entity.DeviceTypeID, err.Error())
		return err
	}
	//再删掉卷和raid的关联关系
	if err := rVolumeRaidDao.UpdateRVolumeRaidByWhere(logger, map[string]interface{}{
		"device_type_id": deviceTypeID,
		"is_del":         baseLogic.IS_NOT_DEL,
	}, &rVolumeRaidDao.RVolumeRaid{
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
		DeletedTime: int(time.Now().Unix()),
		IsDel:       baseLogic.IS_DEL,
	}); err != nil {
		logger.Warnf("del Volume error deviceTypeId:%s", entity.DeviceTypeID, err.Error())
		return err
	}
	if err := modifyVolume(logger, entity.DeviceTypeID, param); err != nil {
		logger.Warn("AddVolume error:", param.Volumes, err.Error())
		return err
	}
	//如果传了raidID
	//if param.RaidID != nil {
	//	var systemVolumeType, systemVolumeInterfaceType, systemVolumeUnit string
	//	var systemVolumeAmount, systemVolumeSize int
	//	if param.SystemVolumeType != nil {
	//		systemVolumeType = *param.SystemVolumeType
	//	} else {
	//		systemVolumeType = entity.SystemVolumeType
	//	}
	//	if param.SystemVolumeInterfaceType != nil {
	//		systemVolumeInterfaceType = *param.SystemVolumeInterfaceType
	//	} else {
	//		systemVolumeInterfaceType = entity.SystemVolumeInterfaceType
	//	}
	//	if param.SystemVolumeUnit != nil {
	//		systemVolumeUnit = *param.SystemVolumeUnit
	//	} else {
	//		systemVolumeUnit = entity.SystemVolumeUnit
	//	}
	//	if param.SystemVolumeAmount != nil {
	//		systemVolumeAmount = *param.SystemVolumeAmount
	//	} else {
	//		systemVolumeAmount = entity.SystemVolumeAmount
	//	}
	//	if param.SystemVolumeSize != nil {
	//		systemVolumeSize = *param.SystemVolumeSize
	//	} else {
	//		systemVolumeSize = entity.SystemVolumeSize
	//	}
	//	if err := addDeviceTypeRaid(logger, deviceTypeID, *param.RaidID, systemVolumeUnit, systemVolumeType, systemVolumeInterfaceType, systemVolumeSize, systemVolumeAmount); err != nil {
	//		logger.Warn("AddDeviceTypeRaid sql error:", deviceTypeID, err.Error())
	//		return err
	//	}
	//}
	//如果不传raidId，就不做修改
	//修改device表的device_series
	if param.DeviceSeries != nil {
		if err := deviceLogic.ModifyDeviceByWhere(logger, "device_series", *param.DeviceSeries, deviceTypeID); err != nil {
			return err
		}
	}
	return nil
}
func DeleteDeviceType(logger *log.Logger, deviceTypeID string) error {
	entity, err := deviceTypeDao.GetOneDeviceType(logger, map[string]interface{}{
		"device_type_id": deviceTypeID,
		"is_del":         baseLogic.IS_NOT_DEL,
	})
	if err != nil {
		logger.Warn("GetDeviceTypeByUuid sql error:", deviceTypeID, err.Error())
		return err
	}
	device, err := deviceDao.GetAllDevice(logger, map[string]interface{}{
		"device_type_id": deviceTypeID,
		"is_del":         baseLogic.IS_NOT_DEL,
	})
	if len(device) != 0 {
		panic(constant.BuildInvalidArgumentWithArgs("机型下面还有设备，不能删除", "It has some devices in the deviceType,please delete device first"))
	}
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	entity.DeletedTime = int(time.Now().Unix())
	entity.IsDel = baseLogic.IS_DEL

	if err := deviceTypeDao.UpdateDeviceTypeByDeviceTypeID(logger, entity, deviceTypeID); err != nil {
		logger.Warnf("DeleteDeviceType sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return err
	}
	//rDeviceTypeRaid, err := rDeviceTypeRaidDao.GetOneRDeviceTypeRaid(logger, map[string]interface{}{
	//	"device_type_id": deviceTypeID,
	//	"is_del":         baseLogic.IS_NOT_DEL,
	//})
	//if rDeviceTypeRaid != nil { //如果机型和raid关联关系存在，那么删除机型的同时，把和raid的关系也一起删掉
	//	if err := rDeviceTypeRaidDao.UpdateRDeviceTypeRaidByWhere(logger, map[string]interface{}{
	//		"device_type_id": deviceTypeID,
	//		"is_del":         baseLogic.IS_NOT_DEL,
	//	}, &rDeviceTypeRaidDao.RDeviceTypeRaid{
	//		UpdatedBy:   logger.GetPoint("username").(string),
	//		UpdatedTime: int(time.Now().Unix()),
	//		DeletedTime: int(time.Now().Unix()),
	//		IsDel:       baseLogic.IS_DEL,
	//	}); err != nil {
	//		logger.Warnf("del rdeviceTypeRaid error deviceTypeId:%s", entity.DeviceTypeID, err.Error())
	//		return err
	//	}
	//}
	//删除机型镜像的绑定关系
	rDeviceTypeImage, err := rDeviceTypeImageDao.GetOneRDeviceTypeImage(logger, map[string]interface{}{
		"device_type_id": deviceTypeID,
		"is_del":         baseLogic.IS_NOT_DEL,
	})
	if rDeviceTypeImage != nil {
		if err := rDeviceTypeImageDao.UpdateMultiRDeviceTypeImage(logger, map[string]interface{}{
			"device_type_id": deviceTypeID,
			"is_del":         baseLogic.IS_NOT_DEL,
		}, map[string]interface{}{
			"is_del":       baseLogic.IS_DEL,
			"deleted_time": int(time.Now().Unix()),
		}); err != nil {
			return err
		}
	}

	return nil
}
func DeviceTypeEntity2DeviceType(logger *log.Logger, entity *deviceTypeDao.DeviceType) *responseTypes.DeviceType {
	language := logger.GetPoint("language").(string)
	tz := logger.GetPoint("timezone").(string)
	//deviceTypeRaid, err := rDeviceTypeRaidDao.GetAllRDeviceTypeRaid(logger, map[string]interface{}{
	//	"device_type_id": entity.DeviceTypeID,
	//	"is_del":         0,
	//})
	//if err != nil {
	//	panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	//}
	//raid := ""
	//if len(deviceTypeRaid) != 0 {
	//	for _, v := range deviceTypeRaid {
	//		raid = raid + v.RaidID + ","
	//	}
	//	raid = strings.TrimRight(raid, ",")
	//}
	////计算raid对应map中的Name
	//array := []string{} // 拆分查询
	//deviceTypeRaidName := ""
	////获取raid的信息
	//raidMap := make(map[string]responseTypes.Raid) //	请求一次，避免循环请求
	//raidList, raidsErr := raidLogic.QueryRaidsAll(logger)
	//if raidsErr == nil && len(raidList) != 0 {
	//	for _, vv := range raidList {
	//		raidMap[vv.RaidID] = *vv
	//	}
	//}
	//if len(raidList) != 0 {
	//	array = strings.Split(raid, ",")
	//}
	//if len(array) != 0 {
	//	raid_ := ""
	//	for _, str := range array {
	//		raid_ = raid_ + raidMap[str].Name + "," // RAID0,RAID1
	//	}
	//	deviceTypeRaidName = raid_[0 : len(raid_)-1]
	//}
	info := deviceLogic.GetDeviceTypeInfo(logger, entity)

	deviceSeries := baseLogic.DeviceTypeSeries[entity.DeviceSeries]
	if language == baseLogic.EN_US {
		deviceSeries = baseLogic.DeviceTypeSeriesEn[entity.DeviceSeries]
	}
	idc, err := idcDao.GetIdcById(logger, entity.IDcID) //获取机房信息
	if err != nil {
		idc = &idcDao.Idc{}
	}
	idcName := idc.Name
	if language == baseLogic.EN_US {
		idcName = idc.NameEn
	}
	//var systemVolumeSize, dataVolumeSize int
	//if entity.SystemVolumeUnit == "TB" {
	//	systemVolumeSize = entity.SystemVolumeSize / K
	//} else {
	//	systemVolumeSize = entity.SystemVolumeSize
	//}
	//if entity.DataVolumeUnit == "TB" {
	//	dataVolumeSize = entity.DataVolumeSize / K
	//} else {
	//	dataVolumeSize = entity.DataVolumeSize
	//}
	instance, _ := instanceDao.GetAllInstance(logger, map[string]interface{}{
		"device_type_id": entity.DeviceTypeID,
	})
	device, _ := deviceDao.GetAllDevice(logger, map[string]interface{}{
		"device_type_id": entity.DeviceTypeID,
	})
	volumeList, _ := volumeDao.GetAllVolume(logger, map[string]interface{}{
		"device_type_id": entity.DeviceTypeID,
	})
	volumeArr := []responseTypes.Volume{}
	for _, v := range volumeList {
		volume := responseTypes.Volume{}
		volume.VolumeID = v.VolumeID
		volume.VolumeName = v.VolumeName
		volume.VolumeType = v.VolumeType
		volume.DiskType = v.DiskType
		volume.VolumeSize = v.VolumeSize
		volume.VolumeUnit = v.VolumeUnit
		volume.VolumeAmount = v.VolumeAmount
		volume.InterfaceType = v.InterfaceType

		rVolumeRaid, _ := rVolumeRaidDao.GetAllRVolumeRaid(logger, map[string]interface{}{
			"volume_id": v.VolumeID,
			"is_del":    0,
			//"device_type_id": entity.DeviceTypeID,
		})
		raidStr := ""
		raidIdStr := ""
		raidCan := ""
		if len(rVolumeRaid) > 0 {
			for _, value := range rVolumeRaid {
				raidStr = raidStr + value.RaidName + ","
				raidIdStr = raidIdStr + value.RaidID + ","
			}
			raidCan = rVolumeRaid[0].RaidCan
		}
		volume.RaidCan = raidCan
		volume.Raid = strings.Trim(raidStr, ",")
		volume.RaidId = strings.Trim(raidIdStr, ",")
		volumeArr = append(volumeArr, volume)

	}
	return &responseTypes.DeviceType{
		ID:               entity.ID,
		IDcID:            entity.IDcID,
		IDcName:          idcName,
		IDcNameEn:        idc.NameEn,
		DeviceType:       entity.DeviceType,
		DeviceTypeID:     entity.DeviceTypeID,
		Name:             entity.Name,
		Description:      entity.Description,
		DeviceSeries:     entity.DeviceSeries,
		DeviceSeriesName: deviceSeries,
		Architecture:     entity.Architecture,
		Height:           entity.Height,
		CPUAmount:        entity.CPUAmount,
		CPUCores:         entity.CPUCores,
		CPUFrequency:     entity.CPUFrequency,
		CPUManufacturer:  entity.CPUManufacturer,
		CPUModel:         entity.CPUModel,

		MemType:       entity.MemType,
		MemAmount:     entity.MemAmount,
		MemSize:       entity.MemSize,
		MemFrequency:  entity.MemFrequency,
		NicAmount:     entity.NicAmount,
		NicRate:       entity.NicRate,
		InterfaceMode: entity.InterfaceMode,
		//SystemVolumeType:          entity.SystemVolumeType,
		//SystemVolumeInterfaceType: entity.SystemVolumeInterfaceType,
		//SystemVolumeAmount:        entity.SystemVolumeAmount,
		//RaidCan:                   entity.RaidCan,
		//SystemVolumeSize:          systemVolumeSize,
		//SystemVolumeUnit:          entity.SystemVolumeUnit,
		//DataVolumeType:            entity.DataVolumeType,
		//DataVolumeInterfaceType:   entity.DataVolumeInterfaceType,
		//DataVolumeAmount:          entity.DataVolumeAmount,
		//DataVolumeSize:            dataVolumeSize,
		//DataVolumeUnit:            entity.DataVolumeUnit,
		GpuAmount:       entity.GpuAmount,
		GpuManufacturer: entity.GpuManufacturer,
		GpuModel:        entity.GpuModel,
		CreatedTime:     util.TimestampToString(int64(entity.CreatedTime), tz),
		UpdatedTime:     util.TimestampToString(int64(entity.UpdatedTime), tz),
		CreatedBy:       entity.CreatedBy,
		UpdatedBy:       entity.UpdatedBy,
		//Raid:            strings.Trim(raid, ","),
		CpuSpec:  entity.CpuSpec,
		MemSpec:  entity.MemSpec,
		BootMode: entity.BootMode,

		//拼装信息
		CpuInfo: info["cpu_"],
		MemInfo: info["mem_"],
		//SvInfo:        info["sv_"],
		//DvInfo:        info["dv_"],
		GpuInfo:       info["gpu_"],
		NicInfo:       info["nic_"],
		InstanceCount: len(instance),
		DeviceCount:   len(device),
		IsNeedRaid:    entity.IsNeedRaid,
		Volumes:       volumeArr,
	}
}

func AssociatedImage(logger *log.Logger, param *requestTypes.AssociateImageRequest) error {
	toInsert := []*rDeviceTypeImageDao.RDeviceTypeImage{}
	toUpdate := []*rDeviceTypeImageDao.RDeviceTypeImage{}
	toDelete := []*rDeviceTypeImageDao.RDeviceTypeImage{}
	deviceTypeID := param.DeviceTypeID
	deviceType, err := deviceTypeDao.GetDeviceTypeById(logger, deviceTypeID)
	if err != nil {
		logger.Warnf("QueryDeviceType sql err, param:%s, error:%s", util.ObjToJson(param), err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	all, _ := rDeviceTypeImageDao.GetByDeviceTypeID(logger, deviceTypeID)
	for _, v := range all {
		if !util.InArray(v.ImageID, param.ImageIDs) {
			//如果数据库中的记录不在参数中，说明是要被删除的
			entity := &rDeviceTypeImageDao.RDeviceTypeImage{
				DeviceTypeID: deviceTypeID,
				ImageID:      v.ImageID,
				UpdatedBy:    logger.GetPoint("username").(string),
				UpdatedTime:  int(time.Now().Unix()),
				DeletedTime:  int(time.Now().Unix()),
				IsDel:        baseLogic.IS_DEL,
			}
			toDelete = append(toDelete, entity)
		}
		//否则就是在数组中，走下面逻辑
	}
	for _, imageID := range param.ImageIDs {
		//fmt.Println(deviceTypeID, imageID)
		image, err := imageDao.GetImageByUuid(logger, imageID)
		if err != nil {
			logger.Warnf("镜像不存在，机型id:%s,镜像id：%s,error:%s", deviceTypeID, imageID, err.Error())
			return err
		}
		os, err := osDao.GetOsByUuid(logger, image.OsID)
		if err != nil {
			logger.Warnf("OS不存在，机型id:%s,OS id：%s,error:%s", deviceTypeID, image.OsID, err.Error())
			return err
		}
		if deviceType.Architecture != os.Architecture {
			panic(constant.BuildInvalidArgumentWithArgs("机型和镜像体系架构不匹配", "architecture not match"))
		}
		_, err = rDeviceTypeImageDao.GetByDeviceTypeIDAndImageID(logger, deviceTypeID, imageID)
		if err != nil {
			logger.Warn("AssociatedImage GetByDeviceTypeIDAndImageID sql err:", deviceTypeID, imageID, err.Error())
			entity := &rDeviceTypeImageDao.RDeviceTypeImage{
				DeviceTypeID: deviceTypeID,
				ImageID:      imageID,
				CreatedBy:    logger.GetPoint("username").(string),
				CreatedTime:  int(time.Now().Unix()),
				UpdatedTime:  int(time.Now().Unix()),
			}
			toInsert = append(toInsert, entity)
		} else {
			entity := &rDeviceTypeImageDao.RDeviceTypeImage{
				DeviceTypeID: deviceTypeID,
				ImageID:      imageID,
				UpdatedBy:    logger.GetPoint("username").(string),
				UpdatedTime:  int(time.Now().Unix()),
			}
			toUpdate = append(toUpdate, entity)
		}

		//先将已有的机型和镜像分区关系删除，然后再新增
		err = rDeviceTypeImagePartitionDao.UpdateMultiRDeviceTypeImagePartition(logger, map[string]interface{}{
			"image_id":       imageID,
			"device_type_id": deviceTypeID,
			"is_del":         0,
		}, map[string]interface{}{
			"updated_by":   logger.GetPoint("username").(string),
			"updated_time": int(time.Now().Unix()),
			"deleted_time": int(time.Now().Unix()),
			"is_del":       baseLogic.IS_DEL,
		})
		if err != nil {
			logger.Warnf("删除机型镜像分区表关系错误，机型id:%s,镜像id：%s,error:%s", deviceTypeID, imageID, err.Error())
			return err
		}

		//“/ ” :50GiB，xfs;swap:10GiB，swap
		//partition := image.SystemPartition //标准版只支持系统盘默认分区，数据盘在专业版才有，先预留
		//if json.Unmarshal([]byte(image.SystemPartition),&Partition)
		var partitionsObj = []requestTypes.Partition{} //注意，这里前端给的参数默认的是json字符串，创建实例那边默认是数组
		if image.SystemPartition != "" {
			if err := json.Unmarshal([]byte(image.SystemPartition), &partitionsObj); err != nil {
				return err
			}
		}
		for i := 0; i < len(partitionsObj); i++ { //数组的结构体partitionObj，不能用range遍历？？
			_, err = rDeviceTypeImagePartitionDao.AddRDeviceTypeImagePartition(logger, &rDeviceTypeImagePartitionDao.RDeviceTypeImagePartition{
				ImageID:             imageID,
				DeviceTypeID:        deviceTypeID,
				BootMode:            deviceType.BootMode,
				PartitionType:       baseLogic.PartitionType[partitionsObj[i].MountPoint],
				PartitionSize:       partitionsObj[i].Size,
				PartitionFsType:     partitionsObj[i].FsType,
				PartitionLabel:      baseLogic.PartitionLabel[partitionsObj[i].MountPoint],
				PartitionMountPoint: partitionsObj[i].MountPoint,
				SystemDiskLabel:     baseLogic.SystemDiskLabel[partitionsObj[i].MountPoint],
				DataDiskLabel:       baseLogic.DataDiskLabel[partitionsObj[i].MountPoint],

				CreatedBy:   logger.GetPoint("username").(string),
				UpdatedBy:   logger.GetPoint("username").(string),
				CreatedTime: int(time.Now().Unix()),
				UpdatedTime: int(time.Now().Unix()),
			})
			if err != nil {
				logger.Warnf("添加机型镜像分区表关系错误，机型id:%s,镜像id：%s,error:%s", deviceTypeID, imageID, err.Error())
				return err
			}
		}
	}
	//fmt.Println(util.ObjToJson(toInsert), util.ObjToJson(toUpdate), util.ObjToJson(toDelete))
	if len(all)+len(toInsert) > 20 {
		logger.Warnf("AssociatedImage all amount %d,inset amount:%d", len(all), len(toInsert))
		panic(constant.BuildInvalidArgumentWithArgs("关联镜像不能超过20个", "AssociatedImage amount can't exceed 20"))
	}
	if len(toInsert) > 0 {
		if _, err := rDeviceTypeImageDao.AddMultiRDeviceTypeImage(logger, toInsert); err != nil {
			logger.Warn("AssociatedImage AddMultiRDeviceTypeImage sql err:", err.Error())
			return err
		}
	}
	if len(toUpdate) > 0 {
		if err := rDeviceTypeImageDao.AssociatedImageMultiRDeviceTypeImage(logger, toUpdate); err != nil {
			logger.Warn("AssociatedImage AssociatedImageMultiRDeviceTypeImage sql err:", err.Error())
			return err
		}
	}
	// 如果需要将原来的删掉，执行以下代码，现在和前端约定，前端给什么，之前已经选中了的，不做处理！
	//if len(toDelete) > 0 {
	//	if err := rDeviceTypeImageDao.AssociatedImageMultiRDeviceTypeImage(logger, toDelete); err != nil {
	//		logger.Warn("AssociatedImage AssociatedImageMultiRDeviceTypeImage sql err:", err.Error())
	//		return err
	//	}
	//}
	//绑定镜像分区，入库关联表
	return nil
}

func DissociatedImage(logger *log.Logger, param *requestTypes.DissociatedImageRequest, isDelDeviceType bool) error {
	if isDelDeviceType { //如果是镜像删除机型，需要增加一个额外的判断，判断机型下面所有的实例，没有创建中的，如果有，不允许删除
		instanceList, _ := instanceDao.GetAllInstance(logger, map[string]interface{}{
			"device_type_id": param.DeviceTypeID,
			"image_id":       param.ImageID,
		})
		for _, v := range instanceList {
			if v.Status == baseLogic.CREATING {
				panic(constant.BuildInvalidArgumentWithArgs("机型下面存在创建中的实例，不允许移除", "instance status exist creating"))
			}
		}
	}
	_, err := rDeviceTypeImageDao.GetByDeviceTypeIDAndImageID(logger, param.DeviceTypeID, param.ImageID)
	if err != nil {
		logger.Warn("DissociatedImage GetByDeviceTypeAndImageId sql err:", param.DeviceTypeID, param.ImageID, err.Error())
		return err
	}
	entity := &rDeviceTypeImageDao.RDeviceTypeImage{
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
		DeletedTime: int(time.Now().Unix()),
		IsDel:       baseLogic.IS_DEL,
	}
	if err := rDeviceTypeImageDao.UpdateRDeviceTypeImageById(logger, entity, param.DeviceTypeID, param.ImageID); err != nil {
		logger.Warn("DissociatedImage UpdateRDeviceTypeImageById sql error:", param.DeviceTypeID, param.ImageID, err.Error())
		return err
	}
	err = rDeviceTypeImagePartitionDao.UpdateMultiRDeviceTypeImagePartition(logger, map[string]interface{}{
		"image_id":       param.ImageID,
		"device_type_id": param.DeviceTypeID,
		"is_del":         0,
	}, map[string]interface{}{
		"updated_by":   logger.GetPoint("username").(string),
		"updated_time": int(time.Now().Unix()),
		"deleted_time": int(time.Now().Unix()),
		"is_del":       baseLogic.IS_DEL,
	})
	if err != nil {
		logger.Warnf("删除机型镜像分区表关系错误，机型id:%s,镜像id：%s,error:%s", param.DeviceTypeID, param.ImageID, err.Error())
		return err
	}
	return nil
}
func QueryDeviceTypeImage(logger *log.Logger, param *requestTypes.QueryDeviceTypeImageRequest, p util.Pageable) ([]*responseTypes.Image, int64, error) {
	tz := logger.GetPoint("timezone").(string)
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"image_id":       param.ImageID,
		"device_type_id": param.DeviceTypeID,
		"image_name":     param.ImageName,
		"version":        param.Version,
		"os_id":          param.OsID,
		//"image_ids":   strings.Join(param.ImageIds, ","),
		"source":       param.Source,
		"architecture": param.Architecture,
		"os_type":      param.OsType,
		"is_del":       baseLogic.IS_NOT_DEL,
	}
	//queryList := query
	//query["count"] = true
	_, count, err := rDeviceTypeImageDao.GetDeviceTypeImageAndOsList(logger, query, true, 0, 0)
	//os.Exit(1)
	if err != nil {
		logger.Warnf("DescribeDeviceTypes_GetDeviceTypeCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	entityList := []*imageDao.DImage{}
	if param.IsAll == baseLogic.IS_ALL {
		//entityList, err = imageDao.GetMultiImage(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
		fmt.Println("获取全部", query)
		entityList, _, err = rDeviceTypeImageDao.GetDeviceTypeImageAndOsList(logger, query, false, 0, 0)
	} else {
		//entityList, err = imageDao.GetMultiImage(logger, query, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
		entityList, _, err = rDeviceTypeImageDao.GetDeviceTypeImageAndOsList(logger, query, false, offset, limit)
	}
	if err != nil {
		logger.Warn("imageDao.QueryByDeviceOsType sql error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	images := []*responseTypes.Image{}
	for _, image := range entityList {
		isBind := false
		if param.DeviceTypeID != "" {
			deviceTypeImageNum, _ := rDeviceTypeImageDao.QueryDeviceTypeImagesCount(logger, map[string]interface{}{
				"image_id":       image.ImageID,
				"device_type_id": param.DeviceTypeID,
				"is_del":         baseLogic.IS_NOT_DEL,
			})
			if deviceTypeImageNum != 0 {
				isBind = true
			}
		}
		//fmt.Println(util.InterfaceToJson(image), image.ImageID, image.CreatedTime)
		//os, err := osLogic.GetByOsId(logger, image.OsID)
		//if err != nil && os == nil {
		//	os = &responseTypes.Os{} //如果没有找到os，默认为空
		//}
		deviceTypeNum, err := rDeviceTypeImageDao.QueryDeviceTypeImagesCount(logger, map[string]interface{}{
			"image_id": image.ImageID,
			"is_del":   baseLogic.IS_NOT_DEL,
		})
		if err != nil {
			logger.Warn("获取镜像对应的机型个数错误", image.OsID, err.Error())
			return nil, 0, err
		}
		images = append(images, imageLogic.ImageEntityWithOs2Image(logger, image, int(deviceTypeNum), isBind, tz))
	}
	return images, count, nil
}

//func QueryDeviceTypeImage(logger *log.Logger, param *requestTypes.QueryDeviceTypeImageRequest) ([]*responseTypes.Image, error) {
//	//offset, limit := p.Pageable2offset()
//	query := map[string]interface{}{
//		"device_type_id": param.DeviceTypeID,
//		"image_id":       param.ImageID,
//		"is_del":         0,
//	}
//	rDeviceTypeImage, err := rDeviceTypeImageDao.GetMultiRDeviceTypeImage(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 1000)
//	if err != nil {
//		logger.Warnf("DescribeDeviceTypes_GetDeviceTypeCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
//		return nil, err
//	}
//	images := []*responseTypes.Image{}
//	for _, value := range rDeviceTypeImage {
//		image, err := imageLogic.GetByImageId(logger, value.ImageID)
//		if err != nil {
//			logger.Warnf("GetImageById  error, imageId:%s, error:%s", value.ImageID, err.Error())
//			return nil, err
//		}
//		if param.Architecture != "" && param.OsType != "" {
//			arr := strings.Split(param.Architecture, ",")
//			arr1 := strings.Split(param.OsType, ",")
//			if util.InArray(image.Architecture, arr) && util.InArray(image.OsType, arr1) {
//				images = append(images, image)
//			} else {
//				continue
//			}
//		}
//		if param.Architecture != "" && param.OsType == "" {
//			arr := strings.Split(param.Architecture, ",")
//			if !util.InArray(image.Architecture, arr) {
//				continue
//			} else {
//				images = append(images, image)
//			}
//		}
//		if param.Architecture == "" && param.OsType != "" {
//			arr := strings.Split(param.OsType, ",")
//			if !util.InArray(image.OsType, arr) {
//				continue
//			} else {
//				images = append(images, image)
//			}
//		}
//		if param.Architecture == "" && param.OsType == "" {
//			images = append(images, image)
//		}
//
//	}
//	return images, err
//}
//func QueryDeviceTypeRaid(logger *log.Logger, param *requestTypes.QueryDeviceTypeRaidRequest) ([]*responseTypes.RDeviceTypeRaid, error) {
//	//offset, limit := p.Pageable2offset()
//	tz := logger.GetPoint("timezone").(string)
//	query := map[string]interface{}{
//		"device_type_id": param.DeviceTypeID,
//		"volume_type":    param.VolumeType,
//		"raid_id":        param.RaidID,
//		"is_del":         0,
//	}
//	rDeviceTypeRaid, err := rDeviceTypeRaidDao.GetAllRDeviceTypeRaid(logger, query)
//	if err != nil {
//		logger.Warnf("DescribeDeviceTypes_GetDeviceTypeCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
//		return nil, err
//	}
//	raid := &responseTypes.RDeviceTypeRaid{}
//	raids := []*responseTypes.RDeviceTypeRaid{}
//	for _, value := range rDeviceTypeRaid {
//		raid = RDeviceTypeRaidEntityToRDeviceTypeRaid(value, tz)
//		raids = append(raids, raid)
//	}
//	return raids, err
//}

//表实例转为返回
//func RDeviceTypeRaidEntityToRDeviceTypeRaid(entity *rDeviceTypeRaidDao.RDeviceTypeRaid, tz string) *responseTypes.RDeviceTypeRaid {
//	return &responseTypes.RDeviceTypeRaid{
//		ID:                   entity.ID,
//		RaidID:               entity.RaidID,
//		DeviceTypeID:         entity.DeviceTypeID,
//		VolumeType:           entity.VolumeType,
//		VolumeDetail:         entity.VolumeDetail,
//		AvailableValue:       entity.AvailableValue,
//		SystemPartitionCount: entity.SystemPartitionCount,
//		DiskType:             entity.DiskType,
//		DiskInterfaceType:    entity.DiskInterfaceType,
//		CreatedBy:            entity.CreatedBy,
//		UpdatedBy:            entity.UpdatedBy,
//		CreatedTime:          util.TimestampToString(int64(entity.CreatedTime), tz),
//		UpdatedTime:          util.TimestampToString(int64(entity.UpdatedTime), tz),
//	}
//}
func QueryDeviceTypeImagePartition(logger *log.Logger, param *requestTypes.QueryDeviceTypeImagePartitionRequest) (*responseTypes.QueryDeviceTypeImagePartitionResponse, error) {
	entities, err := rDeviceTypeImagePartitionDao.GetByDeviceTypeAndImageId(logger, param.DeviceTypeID, param.ImageID)
	if err != nil {
		logger.Warnf("GetByDeviceTypeAndImageId error:", param.DeviceTypeID, param.ImageID)
		return nil, err
	}
	res := &responseTypes.QueryDeviceTypeImagePartitionResponse{
		SystemPartition: []responseTypes.Partition{},
		DataPartition:   []responseTypes.Partition{},
	}
	for _, entity := range entities {
		if strings.EqualFold("data", entity.PartitionType) {
			res.DataPartition = append(res.DataPartition, responseTypes.Partition{
				Size:       entity.PartitionSize,
				FsType:     entity.PartitionFsType,
				MountPoint: entity.PartitionMountPoint,
				//Label:      entity.PartitionLabel,
			})
		} else {
			res.SystemPartition = append(res.SystemPartition, responseTypes.Partition{
				Size:       entity.PartitionSize,
				FsType:     entity.PartitionFsType,
				MountPoint: entity.PartitionMountPoint,
				//Label:      entity.PartitionLabel,
			})
		}
	}

	return res, nil

}
func Trim(str string) string {
	s := strings.TrimRight(str, "0")
	s = strings.TrimRight(s, ".")
	return s
}

func DescribeVolumesByDeviceType(logger *log.Logger, param *requestTypes.DescribeVolumesByDeviceTypeRequest) ([]*responseTypes.VolumeIt, error) {

	volumes := []*responseTypes.VolumeIt{}
	q1 := map[string]interface{}{
		"is_del":         0,
		"device_type_id": param.DeviceTypeID,
	}
	volumeEntities, err := volumeDao.GetAllVolume(logger, q1)
	if err != nil {
		logger.Warnf("DescribeVolumesByDeviceType.GetAllVolume error, device_type_id:%s, error:%s", param.DeviceTypeID, err.Error())
	}
	volume := responseTypes.Volume{}
	for _, v := range volumeEntities {
		volume.VolumeID = v.VolumeID
		volume.VolumeName = v.VolumeName
		volume.VolumeType = v.VolumeType
		volume.VolumeSize = v.VolumeSize
		volume.VolumeUnit = v.VolumeUnit

		volume.VolumeAmount = v.VolumeAmount
		volume.DiskType = v.DiskType
		volume.InterfaceType = v.InterfaceType

		volumeId := v.VolumeID
		q2 := map[string]interface{}{
			"is_del":    0,
			"volume_id": v.VolumeID,
		}
		volumnRaidEntities, err := rVolumeRaidDao.GetAllRVolumeRaid(logger, q2)
		if err != nil {
			logger.Warnf("DescribeVolumesByDeviceType.GetVolumeById error, volume_id:%s, error:%s", volumeId, err.Error())
		}
		volume.RaidCan = volumnRaidEntities[0].RaidCan
		raids := []*raidDao.Raid{}
		for _, v2 := range volumnRaidEntities {
			raids = append(raids, &raidDao.Raid{
				RaidID: v2.RaidID,
				Name:   v2.RaidName,
			})
		}

		// diskEntities, err := diskDao.GetDiskWithVolumeId(logger, volumeId)
		// if err != nil {
		// 	logger.Warnf("DescribeVolumesByDeviceType.GetVolumeById error, volume_id:%s, error:%s", volumeId, err.Error())
		// }

		vt := &responseTypes.VolumeIt{
			Volume: volume,
			Raids:  raids,
			Disks:  []*diskDao.Disk{}, //给前端一个空[]
		}
		volumes = append(volumes, vt)
	}
	return volumes, nil

}
