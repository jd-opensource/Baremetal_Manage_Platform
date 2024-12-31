package deviceTypeLogic

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/idcDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rVolumeRaidDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/volumeDao"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceTypeDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/imageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rDeviceTypeImageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rDeviceTypeImagePartitionDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/imageLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

var K int = 1

func BuildDeviceTypeEntity(logger *log.Logger, param *requestTypes.CreateDeviceTypeRequest) *deviceTypeDao.DeviceType {
	var systemVolumeSize, dataVolumeSize int
	if param.SystemVolumeUnit == "TB" {
		systemVolumeSize = param.SystemVolumeSize * K
	} else {
		systemVolumeSize = param.SystemVolumeSize
	}
	if param.DataVolumeUnit == "TB" {
		dataVolumeSize = param.DataVolumeSize * K
	} else {
		dataVolumeSize = param.DataVolumeSize
	}
	fmt.Println("开始校验")
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

		SystemVolumeType:          param.SystemVolumeType,
		SystemVolumeInterfaceType: param.SystemVolumeInterfaceType,
		SystemVolumeAmount:        param.SystemVolumeAmount,
		SystemVolumeSize:          systemVolumeSize,
		SystemVolumeUnit:          param.SystemVolumeUnit,

		DataVolumeType:          param.DataVolumeType,
		DataVolumeInterfaceType: param.DataVolumeInterfaceType,
		DataVolumeAmount:        param.DataVolumeAmount,
		DataVolumeSize:          dataVolumeSize,
		DataVolumeUnit:          param.DataVolumeUnit,
		CreatedBy:               logger.GetPoint("username").(string),
		UpdatedBy:               logger.GetPoint("username").(string),
		CreatedTime:             int(time.Now().Unix()),
		UpdatedTime:             int(time.Now().Unix()),
		CpuSpec:                 param.CpuSpec,
		MemSpec:                 param.MemSpec,
		BootMode:                bootmode,
	}
}
func QueryDeviceTypes(logger *log.Logger, param *requestTypes.QueryDeviceTypesRequest, p util.Pageable) ([]*responseTypes.DeviceType, int64, error) {

	language := logger.GetPoint("language").(string)
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
		volumes, err := volumeDao.GetAllVolume(logger, map[string]interface{}{
			"is_del":         0,
			"device_type_id": v.DeviceTypeID,
		})
		if err != nil {
			logger.Warnf("QueryDeviceTypes.GetAllVolumeByDeviceType error, device_type_id:%s, error:%s", v.DeviceTypeID, err.Error())
		}
		for k1, v1 := range volumes {
			if language != baseLogic.EN_US {
				if v1.DiskType == "notLimited" {
					v1.DiskType = "无限制"
				}
				if v1.InterfaceType == "notLimited" {
					v1.InterfaceType = "无限制"
				}

			}
			volumes[k1] = v1
		}
		v.Volumes = volumes
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

func DeviceTypeEntity2DeviceType(logger *log.Logger, entity *deviceTypeDao.DeviceType) *responseTypes.DeviceType {
	language := logger.GetPoint("language").(string)
	tz := logger.GetPoint("timezone").(string)

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

	instance, _ := instanceDao.GetAllInstance(logger, map[string]interface{}{
		"device_type_id": entity.DeviceTypeID,
	})
	device, _ := deviceDao.GetAllDevice(logger, map[string]interface{}{
		"device_type_id": entity.DeviceTypeID,
	})
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

		MemType:         entity.MemType,
		MemAmount:       entity.MemAmount,
		MemSize:         entity.MemSize,
		MemFrequency:    entity.MemFrequency,
		NicAmount:       entity.NicAmount,
		NicRate:         entity.NicRate,
		InterfaceMode:   entity.InterfaceMode,
		GpuAmount:       entity.GpuAmount,
		GpuManufacturer: entity.GpuManufacturer,
		GpuModel:        entity.GpuModel,
		CreatedTime:     util.TimestampToString(int64(entity.CreatedTime), tz),
		UpdatedTime:     util.TimestampToString(int64(entity.UpdatedTime), tz),
		CreatedBy:       entity.CreatedBy,
		UpdatedBy:       entity.UpdatedBy,
		RaidCan:         entity.RaidCan,
		// Raid:                      strings.Trim(raid, ","),
		CpuSpec:  entity.CpuSpec,
		MemSpec:  entity.MemSpec,
		BootMode: entity.BootMode,

		//拼装信息
		CpuInfo:       info["cpu_"],
		MemInfo:       info["mem_"],
		SvInfo:        info["sv_"],
		DvInfo:        info["dv_"],
		GpuInfo:       info["gpu_"],
		NicInfo:       info["nic_"],
		InstanceCount: len(instance),
		DeviceCount:   len(device),
		IsNeedRaid:    entity.IsNeedRaid,
	}
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

func QueryVolumesRaids(logger *log.Logger, param *requestTypes.QueryVolumesRaidsRequest) ([]*responseTypes.VolumeRaids, error) {
	language := logger.GetPoint("language").(string)
	//offset, limit := p.Pageable2offset()
	// tz := logger.GetPoint("timezone").(string)
	query := map[string]interface{}{
		"device_type_id": param.DeviceTypeID,
		"volume_type":    param.VolumeType,
		"is_del":         0,
	}
	volumeEntities, err := volumeDao.GetAllVolume(logger, query)
	if err != nil {
		logger.Warnf("QueryVolumesRaids.GetAllVolume error, device_type_id:%s, error:%s", param.DeviceTypeID, err.Error())
		return nil, err
	}
	res := []*responseTypes.VolumeRaids{}
	for _, v := range volumeEntities {
		if language != baseLogic.EN_US {
			if v.DiskType == "notLimited" {
				v.DiskType = "无限制"
			}
			if v.InterfaceType == "notLimited" {
				v.InterfaceType = "无限制"
			}

		}
		rVolumeRaidEntities, err := rVolumeRaidDao.GetRVolumeRaidsByVId(logger, v.VolumeID)
		if err != nil {
			logger.Warnf("QueryVolumesRaids.GetRVolumeRaidsByVId error, volume_id:%s, error:%s", v.VolumeID, err.Error())
		}
		item := &responseTypes.VolumeRaids{
			Volume: *v,
			Raids:  rVolumeRaidEntities,
		}
		res = append(res, item)

	}
	return res, err
}

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
