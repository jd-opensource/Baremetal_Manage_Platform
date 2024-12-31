package instanceLogic

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/auditLogsDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceTypeDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/idcDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/imageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instancePartitionDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/projectDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rDeviceTypeImageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rDeviceTypeImagePartitionDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rInstanceSshkeyDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rInstanceVolumeRaidDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/raidDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/sharingProjectDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/userDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/volumeDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/auditLogLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/deviceTypeLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/idcLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/imageLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/raidLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/sshkeyLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-openapi/service/rabbit_mq"
	namespacePrefix "git.jd.com/cps-golang/ironic-common/ironic/common/NamespacePrefixs"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/AuditLogsType"
	instanceStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/InstanceStatus"
	rabbitEvent "git.jd.com/cps-golang/ironic-common/ironic/event"
	rabbitIronicMsgApi "git.jd.com/cps-golang/ironic-common/ironic/event/api"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

var NoOpForLock []string = []string{instanceStatus.CREATE_ERROR, instanceStatus.CREATING, instanceStatus.DESTROYING, instanceStatus.DESTROY_ERROR}

func InstanceEntity2Instance(logger *log.Logger, o *instanceDao.Instance, vrs []*response.InstanceVolumeRaid) *responseTypes.Instance {
	//logger := &log.Logger{}

	language := logger.GetPoint("language").(string)

	tz := logger.GetPoint("timezone").(string)
	idc, _ := idcDao.GetIdcById(logger, o.IDcID)
	if idc == nil {
		idc = &idcDao.Idc{}
	}
	device, _ := deviceDao.GetDeviceById(logger, o.DeviceID)
	if device == nil {
		device = &deviceDao.Device{}
	}
	deviceType, _ := deviceTypeDao.GetDeviceTypeById(logger, o.DeviceTypeID)
	if deviceType == nil {
		deviceType = &deviceTypeDao.DeviceType{}
	}
	image, _ := imageDao.GetImageByUuid(logger, o.ImageID)
	if image == nil {
		image = &imageDao.Image{}
	}
	raidSys, _ := raidDao.GetRaidById(logger, o.SystemVolumeRaidID)
	if raidSys == nil {
		raidSys = &raidDao.Raid{}
	}
	raidData, _ := raidDao.GetRaidById(logger, o.DataVolumeRaidID)
	if raidData == nil {
		raidData = &raidDao.Raid{}
	}
	reason := baseLogic.InstanceReason[o.Reason]
	if language == baseLogic.EN_US {
		reason = baseLogic.InstanceReasonEn[o.Reason]
	}
	//字典不全，兜底
	if reason == "" {
		reason = o.Reason
	}

	v := &responseTypes.Instance{
		ID:           o.ID,
		InstanceName: o.InstanceName,
		InstanceID:   o.InstanceID,
		UserID:       o.UserID,
		ProjectID:    o.ProjectID,

		IDcID:    o.IDcID,
		IdcName:  idc.Name,
		DeviceID: o.DeviceID,
		Sn:       device.Sn,

		IloIP:       device.IloIP,
		PrivateIPv4: device.PrivateIPv4,
		PrivateIPv6: device.PrivateIPv6,

		DeviceTypeID:              o.DeviceTypeID,
		DeviceTypeName:            deviceType.Name,
		DeviceType:                deviceType.DeviceType,
		DeviceSeries:              deviceType.DeviceSeries,
		DeviceSeriesName:          baseLogic.DeviceTypeSeries[deviceType.DeviceSeries],
		CPUAmount:                 deviceType.CPUAmount,
		CPUCores:                  deviceType.CPUCores,
		CPUManufacturer:           deviceType.CPUManufacturer,
		CPUModel:                  deviceType.CPUModel,
		CPUFrequency:              deviceType.CPUFrequency,
		MemType:                   deviceType.MemType,
		MemSize:                   deviceType.MemSize,
		MemAmount:                 deviceType.MemAmount,
		MemFrequency:              deviceType.MemFrequency,
		NicAmount:                 deviceType.NicAmount,
		NicRate:                   deviceType.NicRate,
		InterfaceMode:             deviceType.InterfaceMode,
		SystemVolumeType:          deviceType.SystemVolumeType,
		SystemVolumeInterfaceType: deviceType.SystemVolumeInterfaceType,
		SystemVolumeSize:          deviceType.SystemVolumeSize,
		SystemVolumeAmount:        deviceType.SystemVolumeAmount,
		SystemVolumeUnit:          deviceType.SystemVolumeUnit,
		GpuAmount:                 deviceType.GpuAmount,
		GpuManufacturer:           deviceType.GpuManufacturer,
		GpuModel:                  deviceType.GpuModel,
		DataVolumeType:            deviceType.DataVolumeType,
		DataVolumeInterfaceType:   deviceType.DataVolumeInterfaceType,
		DataVolumeSize:            deviceType.DataVolumeSize,
		DataVolumeAmount:          deviceType.DataVolumeAmount,
		DataVolumeUnit:            deviceType.DataVolumeUnit,

		Hostname:   o.Hostname,
		Status:     o.Status,
		Reason:     reason,
		StatusName: baseLogic.InstanceStatus[o.Status],
		ImageID:    o.ImageID,
		ImageName:  image.ImageName,

		SystemVolumeRaidID:   o.SystemVolumeRaidID,
		SystemVolumeRaidName: raidSys.Name,
		Locked:               o.Locked,
		LockedName:           baseLogic.InstanceLock[o.Locked],
		DataVolumeRaidID:     o.DataVolumeRaidID,
		DataVolumeRaidName:   raidData.Name,

		Description: o.Description,

		CreatedBy:           o.CreatedBy,
		CreatedTime:         util.TimestampToString(int64(o.CreatedTime), tz),
		UpdatedBy:           o.UpdatedBy,
		UpdatedTime:         util.TimestampToString(int64(o.UpdatedTime), tz),
		InstanceVolumeRaids: vrs,
	}
	if language == baseLogic.EN_US {
		v.StatusName = baseLogic.InstanceStatusEn[o.Status]
	}

	return v
}

func GetInstanceById(logger *log.Logger, InstanceId string) (*responseTypes.Instance, error) {
	instanceEntity, err := instanceDao.GetInstanceById(logger, InstanceId)
	if err != nil {
		logger.Warnf("GetInstanceByUuid sql error, uuid:%s, error:%s", InstanceId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	vrs := []*response.InstanceVolumeRaid{}
	volumeEntities, err := volumeDao.GetAllVolumeByDeviceTypeId(logger, instanceEntity.DeviceTypeID)
	if err != nil {
		logger.Warnf("GetInstanceById.GetAllVolumeByDeviceTypeId error, InstanceId:%s, error:%s", InstanceId, err.Error())
	}
	//有些是盘符，就没有raid，这种情况行不通
	rInstanceVolumeRaidEntities, err := rInstanceVolumeRaidDao.GetRInstanceVolumeRaidByInstanceId(logger, InstanceId)
	if err != nil {
		logger.Warnf("GetInstanceById.GetRVolumeRaidsByInstanceId error, instance_id:%s, error:%s", InstanceId, err.Error())
	}
	for _, volumeEntity := range volumeEntities {
		var tRaid *raidDao.Raid
		for _, rInstanceVolumeRaidEntity := range rInstanceVolumeRaidEntities {
			if rInstanceVolumeRaidEntity.VolumeID == volumeEntity.VolumeID {
				tRaid = &raidDao.Raid{
					RaidID: rInstanceVolumeRaidEntity.RaidID,
					Name:   rInstanceVolumeRaidEntity.RaidName,
				}
				break
			}
		}
		vrs = append(vrs, &response.InstanceVolumeRaid{
			Volume: volumeEntity,
			Raid:   tRaid,
		})

	}
	return InstanceEntity2Instance(logger, instanceEntity, vrs), nil
}

func CreateInstance(logger *log.Logger, param requestTypes.CreateInstanceRequest) (ids []string, err error) {

	if param.Password == "" && param.SshKeyID == "" {
		panic(constant.BuildInvalidArgumentWithArgs("密码和sshkeyid不能同时为空", "password and sshkeyID both empty"))
	}
	if param.Password != "" && param.SshKeyID != "" {
		panic(constant.BuildInvalidArgumentWithArgs("密码和sshkeyid不能同时存在", "password and sshkeyID invalidate"))
	}
	idc, err := idcLogic.QueryIdcById(logger, param.IdcID)
	if err != nil {
		panic(constant.BuildInvalidArgumentWithArgs("idcId非法", "idc invalidate"))
	}
	project, err := projectDao.GetProjectById(logger, param.ProjectID)
	if err != nil {
		panic(constant.BuildInvalidArgumentWithArgs("projectId非法", "projectId invalidate"))
	}
	if project != nil && project.UserID != logger.GetPoint("userId").(string) {
		panic(constant.BuildInvalidArgumentWithArgs("项目ID和用户ID不匹配", "projectID not match userID"))
	}
	deviceType, err := deviceTypeLogic.QueryDeviceTypeById(logger, param.DeviceTypeID)
	if err != nil {
		panic(constant.BuildInvalidArgumentWithArgs("deviceTypeId非法", "deviceTypeId invalidate"))
	}
	if deviceType.IDcID != param.IdcID {
		panic(constant.BuildInvalidArgumentWithArgs("idcId和deviceTypeId不匹配", "idcId not match deviceTypeId"))
	}
	_, err = imageLogic.GetByImageId(logger, param.ImageID)
	if err != nil {
		panic(constant.BuildInvalidArgumentWithArgs("imageId非法", "imageId invalidate"))
	}
	_, err = raidLogic.GetByRaidId(logger, param.SystemVolumeRaidID)
	if err != nil {
		panic(constant.BuildInvalidArgumentWithArgs("systemVolumeRaidId非法", "systemVolumeRaidId invalidate"))
	}
	//检查devicetype和image是否绑定
	q := map[string]interface{}{
		"image_id":       param.ImageID,
		"device_type_id": param.DeviceTypeID,
		"is_del":         baseLogic.IS_NOT_DEL,
	}
	n, _ := rDeviceTypeImageDao.QueryDeviceTypeImagesCount(logger, q)
	if n == 0 {
		panic(constant.BuildInvalidArgumentWithArgs("此机型和镜像未绑定", "devicetype and image do not associated"))
	}
	//检查devicetype和raid是否绑定

	q = map[string]interface{}{
		"raid_id":        param.SystemVolumeRaidID,
		"device_type_id": param.DeviceTypeID,
		"is_del":         baseLogic.IS_NOT_DEL,
	}
	//if _, err := rDeviceTypeRaidDao.GetOneRDeviceTypeRaid(logger, q); err != nil {
	//	panic(constant.BuildInvalidArgumentWithArgs("此机型系统盘和raid未绑定", "devicetype systemdisk and raid do not associated"))
	//}

	if param.SshKeyID != "" {
		sshkeys := strings.Split(param.SshKeyID, ",")
		for _, sshkey := range sshkeys {
			_, err := sshkeyLogic.GetSshkeyById(logger, strings.TrimSpace(sshkey))
			if err != nil {
				panic(constant.BuildInvalidArgumentWithArgs("sshKeyId非法", "sshKeyId invalidate"))
			}
		}

	}
	deviceConditions := deviceLogic.GetAllDevices(logger, map[string]interface{}{
		"manage_status":  baseLogic.PUTAWAY,
		"is_del":         baseLogic.IS_NOT_DEL,
		"device_type_id": param.DeviceTypeID,
	}, []string{}, []string{}, []string{})
	if err != nil {
		return nil, err
	}
	if len(deviceConditions) < param.Count {
		logger.Warn("设备库存不足")
		panic(constant.BuildInvalidArgumentWithArgs("设备库存不足", "device stock is not enough"))
	}
	userId := logger.GetPoint("userId").(string)

	list, _ := instanceDao.GetAllInstance(logger, map[string]interface{}{
		"instance_name": param.InstanceName,
		"user_id":       userId,
		"is_del":        0,
	})
	if len(list) != 0 {
		logger.Warn("instanceName exist:", param.InstanceName)
		panic(constant.BuildInvalidArgumentWithArgs("实例名称已存在", "projectName exist"))
	}

	if param.HostName == "" {
		list, _ := instanceDao.GetAllInstance(logger, map[string]interface{}{
			"hostname.like": "%" + idc.Idc.Shortname + "-" + deviceType.DeviceType + "%",
			"user_id":       userId,
			"is_del":        baseLogic.IS_NOT_DEL,
		})
		param.HostName = fmt.Sprintf("%s-%s-%d", idc.Idc.Shortname, deviceType.DeviceType, len(list)+1)
		param.HostName = strings.ReplaceAll(param.HostName, ".", "-") //机型规格里面的.替换成-，如c1.normal=> c1-normal
		param.HostName = strings.ReplaceAll(param.HostName, "_", "-")
	}

	//从r_devicetype_image_partition表获取bootmode
	imagePartitions, _ := rDeviceTypeImagePartitionDao.GetByDeviceTypeAndImageId(logger, param.DeviceTypeID, param.ImageID)
	if len(imagePartitions) == 0 {
		panic(constant.BuildInvalidArgumentWithArgs("r_device_type_image_partition配置错误", "r_device_type_image_partition config error"))
	}
	var bootmode string //优先取用户指定的
	if param.BootMode != "" {
		bootmode = param.BootMode
	} else {
		bootmode = imagePartitions[0].BootMode
	}

	//user, err := userLogic.GetUserById(logger, userId)
	//if err != nil {
	//	panic(constant.BuildInvalidArgumentWithArgs("用户id不存在", "userId not exist"))
	//}
	entities := []*instanceDao.Instance{} // instanceCreator.Create(logger, param)
	//ids := []string{}
	for i := 0; i < param.Count; i++ {
		instanceName := fmt.Sprintf("%s-%d", param.InstanceName, i)
		if param.Count == 1 {
			instanceName = param.InstanceName
		}
		//instanceList, err := instanceDao.GetAllInstance(logger, map[string]interface{}{
		//	"instance_name": instanceName,
		//	"user_id":       logger.GetPoint("userId").(string),
		//	"is_del":        baseLogic.IS_NOT_DEL,
		//})
		//if len(instanceList) != 0 {
		//	logger.Warn("instanceName exist:", instanceName)
		//	panic(constant.BuildInvalidArgumentWithArgs("实例名称已存在", "instanceName exist"))
		//}
		entity := &instanceDao.Instance{
			IDcID:              param.IdcID,
			InstanceID:         commonUtil.GetRandString("bmp-", namespacePrefix.IMAGE_ID_LENGTH, false, true, true),
			ProjectID:          param.ProjectID,
			UserID:             userId,
			DeviceTypeID:       param.DeviceTypeID,
			InstanceName:       instanceName,
			Hostname:           param.HostName,
			Description:        param.Description,
			ImageID:            param.ImageID,
			SystemVolumeRaidID: param.SystemVolumeRaidID,
			Status:             instanceStatus.CREATING,
			DeviceID:           deviceConditions[i].DeviceID,
			Locked:             baseLogic.UNLOCKED,
			CreatedBy:          logger.GetPoint("username").(string),
			CreatedTime:        int(time.Now().Unix()),
			UpdatedBy:          logger.GetPoint("username").(string),
			UpdatedTime:        int(time.Now().Unix()),
		}

		//if entity.Hostname == "" {
		//	entity.Hostname = fmt.Sprintf("%s-%s-%d", idc.Idc.Name, deviceType.Name, i)
		//}

		//还要写进实例分区表，实例秘钥表 todo
		partitionsObj := param.SystemPartition
		//if len(partitionsObj) == 0 {
		//	panic(constant.BuildInvalidArgumentWithArgs("系统分区不能为空", "SystemPartition is empty"))
		//}
		for i := 0; i < len(partitionsObj); i++ { //数组的结构体partitionObj，不能用range遍历？？
			_, err = instancePartitionDao.AddInstancePartition(logger, &instancePartitionDao.InstancePartition{
				InstanceID:          entity.InstanceID,
				ImageID:             param.ImageID,
				DeviceTypeID:        param.DeviceTypeID,
				BootMode:            bootmode,
				PartitionType:       baseLogic.PartitionType[partitionsObj[i].MountPoint],
				PartitionSize:       partitionsObj[i].Size,
				PartitionFsType:     partitionsObj[i].FsType,
				PartitionLabel:      baseLogic.PartitionLabel[partitionsObj[i].MountPoint],
				PartitionMountPoint: partitionsObj[i].MountPoint,
				SystemDiskLabel:     baseLogic.SystemDiskLabel[partitionsObj[i].MountPoint],
				DataDiskLabel:       baseLogic.DataDiskLabel[partitionsObj[i].MountPoint],

				CreatedBy:   logger.GetPoint("username").(string),
				CreatedTime: int(time.Now().Unix()),
				UpdatedBy:   logger.GetPoint("username").(string),
				UpdatedTime: int(time.Now().Unix()),
			})
			if err != nil {
				logger.Warnf("add instance partition error,instanceId：%s,partition%s,error:%s", entity.InstanceID, util.InterfaceToJson(param.SystemPartition), err.Error())
				return nil, err
			}
		}
		if param.SshKeyID != "" {
			for _, sshkeyId := range strings.Split(param.SshKeyID, ",") {
				sshkeyId = strings.TrimSpace(sshkeyId)
				if sshkeyId == "" {
					panic(constant.BuildInvalidArgumentWithArgs("秘钥ID不能为空", "sshkeyID is empty"))
				}
				_, err = rInstanceSshkeyDao.AddrInstanceSshkey(logger, &rInstanceSshkeyDao.RInstanceSshkey{
					InstanceID:  entity.InstanceID,
					SSHkeyID:    sshkeyId,
					CreatedBy:   logger.GetPoint("username").(string),
					CreatedTime: int(time.Now().Unix()),
					UpdatedBy:   logger.GetPoint("username").(string),
					UpdatedTime: int(time.Now().Unix()),
				})
				if err != nil {
					logger.Warn("add r_instance_sshkey error：", err.Error())
					return nil, err
				}
			}
		}
		_, err = instanceDao.AddInstance(logger, entity)

		if err != nil {
			logger.Warn("createInstance error：", err.Error())
			return nil, err
		}
		if err := deviceLogic.ModifyDeviceByDeviceId(logger, &requestTypes.ModifyAllDevicesRequest{
			ManageStatus: baseLogic.CREATED,
			InstanceID:   entity.InstanceID,
			UserId:       userId,
			UserName:     logger.GetPoint("username").(string),
		}, deviceConditions[i].DeviceID); err != nil {
			logger.Warn("modifyDevice instanceId error：", err.Error())
			return nil, err
		}
		entities = append(entities, entity)
		//ids = append(ids, entity.InstanceID)
	}
	return afterCreateInstances(logger, entities, param)
}
func afterCreateInstances(logger *log.Logger, entities []*instanceDao.Instance, param requestTypes.CreateInstanceRequest) ([]string, error) {

	instance_ids := []string{}
	for _, v := range entities {
		instance_ids = append(instance_ids, v.InstanceID)
	}
	err := sendCreateInstanceEvent(logger, instance_ids, param.Password) //param.UserData, param.AliasIPs, param.ExtensionAliasIps
	if err != nil {
		return nil, err
	}
	if len(entities) == 1 {
		//instance := entities[0]
		//if param.AliasIPs != nil {
		//	saveAliasIPs(logger, instance.InstanceID, instance.SubnetID, param.AliasIPs)
		//}
		//if param.ExtensionAliasIps != nil {
		//	saveAliasIPs(logger, instance.InstanceID, instance.ExtensionSubnetID, param.ExtensionAliasIps)
		//}
	}
	return instance_ids, nil
}

func sendCreateInstanceEvent(logger *log.Logger, instance_ids []string, password string) error {

	logid := logger.GetPoint("logid").(string)
	msg := rabbitIronicMsgApi.CreateInstancesMessage{
		ApiMessage: rabbitIronicMsgApi.ApiMessage{
			RequestId: logid,
		},
		InstanceIds: instance_ids,
		Password:    password,
		//UserData:          user_data,
		//AliasIps:          alias_ips,
		//ExtensionAliasIps: extension_alias_ips,
	}
	userId := logger.GetPoint("userId").(string)
	event, err := rabbitEvent.NewEvent(msg, logid, userId)
	if err != nil {
		logger.Warn("RetryCommand convert event msg error:", instance_ids, err.Error())
		return err
	}
	if err := rabbitMq.SendToApi2Scheduler(event); err != nil {
		logger.Warnf("RetryCommand UpdateCommandById sql error, msg:%s, error:%s", util.ObjToJson(event), err.Error())
		return err
	} else {
		logger.Infof("sendCreateInstanceEvent SendToScheduler success, msg:%s", util.ObjToJson(event))
	}
	return nil
}
func ModifyInstance(logger *log.Logger, param *requestTypes.ModifyInstanceRequest, instanceId string) error {

	q := map[string]interface{}{
		"is_del": 0,
	}
	its, _ := instanceDao.GetAllInstance(logger, q)
	var existNames []string
	var instance *instanceDao.Instance
	for _, item := range its {
		if item.InstanceID == instanceId {
			instance = item
		}
		existNames = append(existNames, item.InstanceName)
	}

	if instance == nil {
		logger.Warnf("GetInstanceByUuid sql not found, uuid:%s", instanceId)
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	if !InstanceOwnedCheck(logger, instance) {
		panic(constant.PermissionDenyForObject)
	}

	if err := checkInstanceNameUnique(logger, param.InstanceName, existNames); err != nil {
		return err
	}

	//instance.InstanceName = param.InstanceName
	if param.Description != nil {
		instance.Description = *param.Description
	}
	instance.UpdatedBy = logger.GetPoint("username").(string)
	instance.UpdatedTime = int(time.Now().Unix())
	if err := instanceDao.UpdateInstanceAnywhere(logger, instance); err != nil {
		logger.Warn("ModifyInstance UpdateInstanceById sql error:", instanceId, err.Error())
		return err
	}
	return nil
}

// 批量编辑实例名称
func ModifyInstances(logger *log.Logger, param requestTypes.ModifyInstancesRequest) error {

	var instances []*instanceDao.Instance

	q := map[string]interface{}{
		"is_del": 0,
	}
	aits, _ := instanceDao.GetAllInstance(logger, q)
	var existNames []string
	for _, v := range aits {
		if !util.InArrayString(v.InstanceID, param.InstanceIDs) {
			existNames = append(existNames, v.InstanceName)
		} else {
			if !InstanceOwnedCheck(logger, v) {
				panic(constant.PermissionDenyForObject)
			}
			instances = append(instances, v)
		}
	}
	for idx, _ := range instances {
		var instanceName string
		if idx == 0 {
			instanceName = param.InstanceName
		} else {
			instanceName = fmt.Sprintf("%s_%d", param.InstanceName, idx)
		}

		instances[idx].InstanceName = instanceName
		if err := checkInstanceNameUnique(logger, instanceName, existNames); err != nil {
			return err
		}
	}

	for _, v := range instances {
		v.UpdatedBy = logger.GetPoint("username").(string)
		if err := instanceDao.UpdateInstanceById(logger, v); err != nil {
			logger.Warnf("ModifyInstances.UpdateInstanceById error, instance_id:%s, instance_name:%s", v.InstanceID, v.InstanceName)
		}
		auditLogLogic.SaveAuditLogs(logger, v.DeviceID, v.InstanceID, AuditLogsType.AuditLogsEditInstanceName)
	}
	return nil
}

func checkInstanceNameUnique(logger *log.Logger, name string, existsNames []string) error {

	if util.InArrayString(name, existsNames) {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例名称 %s 已经存在", name), fmt.Sprintf("instanceName %s already exist", name)))
	}
	return nil
}

func ModifyProjectInstance(logger *log.Logger, param *requestTypes.ModifyInstanceRequest, instanceId string) error {

	q := map[string]interface{}{
		"is_del": 0,
	}
	its, _ := instanceDao.GetAllInstance(logger, q)
	var existNames []string
	var v *instanceDao.Instance
	for _, item := range its {
		if item.InstanceID == instanceId {
			v = item
		}
		existNames = append(existNames, item.InstanceName)
	}

	if v == nil {
		logger.Warnf("GetInstanceByUuid not found, uuid:%s", instanceId)
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	if !InstanceOwnedCheck(logger, v) {
		panic(constant.PermissionDenyForObject)
	}

	if err := checkInstanceNameUnique(logger, param.InstanceName, existNames); err != nil {
		return err
	}

	if param.Description != nil {
		v.Description = *param.Description
	}
	if param.InstanceName != "" {
		v.InstanceName = param.InstanceName
	}

	if err := instanceDao.UpdateInstanceAnywhere(logger, v); err != nil {
		logger.Warn("ModifyInstance UpdateInstanceAnywhere sql error:", instanceId, err.Error())
		return err
	}

	auditLogLogic.SaveAuditLogs(logger, v.DeviceID, instanceId, AuditLogsType.AuditLogsEditInstanceName)
	return nil
}

func LockInstance(logger *log.Logger, instanceId string) error {
	instance, err := instanceDao.GetInstanceById(logger, instanceId)
	if err != nil {
		logger.Warnf("GetInstanceByUuid sql error, uuid:%s, error:%s", instanceId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if !InstanceOwnedCheck(logger, instance) {
		panic(constant.PermissionDenyForObject)
	}

	if instance.Locked == baseLogic.LOCKED {
		logger.Warnf("Instance %s already locked", instanceId)
		panic(constant.NOT_SUPPORTED)
	}

	if util.InArray(instance.Status, NoOpForLock) {
		logger.Warnf("unsupported lock error status instance %s", instanceId)
		panic(constant.NOT_SUPPORTED)
	}

	entity := &instanceDao.Instance{
		InstanceID:  instanceId,
		Locked:      baseLogic.LOCKED,
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
	}
	if err := instanceDao.UpdateInstanceById(logger, entity); err != nil {
		logger.Warn("ModifyInstance UpdateInstanceById sql error:", instanceId, err.Error())
		return err
	}
	return nil
}

func UnLockInstance(logger *log.Logger, instanceId string) error {
	instance, err := instanceDao.GetInstanceById(logger, instanceId)
	if err != nil {
		logger.Warn("GetInstanceById sql error:", instanceId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	if !InstanceOwnedCheck(logger, instance) {
		panic(constant.PermissionDenyForObject)
	}

	if instance.Locked == baseLogic.UNLOCKED {
		logger.Warnf("instance %s already unlocked.", instanceId)
		panic(constant.NOT_SUPPORTED)
	}

	if util.InArray(instance.Status, NoOpForLock) {
		logger.Warnf("unsupported lock error status instance %s", instanceId)
		panic(constant.NOT_SUPPORTED)
	}

	entity := &instanceDao.Instance{
		InstanceID:  instanceId,
		Locked:      baseLogic.UNLOCKED,
		UpdatedBy:   "admin",
		UpdatedTime: int(time.Now().Unix()),
	}
	if err := instanceDao.UpdateInstanceById(logger, entity); err != nil {
		logger.Warn("ModifyInstance UpdateInstanceById sql error:", instanceId, err.Error())
		return err
	}
	return nil
}

func StartInstance(logger *log.Logger, instanceId string) error {
	instance, err := instanceDao.GetInstanceById(logger, instanceId)
	if err != nil {
		logger.Warnf("GetInstanceByUuid sql error, uuid:%s, error:%s", instanceId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	if !InstanceOwnedCheck(logger, instance) {
		panic(constant.PermissionDenyForObject)
	}

	//可以在此增加一些逻辑判断，比如，只能处在关机模式下才可以开机，其他状态不可以。。。后续再处理这块
	if instance.Status != instanceStatus.STOPPED {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态不是关机", instanceId), fmt.Sprintf("unsupported start instance %s", instanceId)))
	}
	entity := &instanceDao.Instance{
		InstanceID:  instanceId,
		Status:      instanceStatus.STARTING,
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
	}
	if err := instanceDao.UpdateInstanceById(logger, entity); err != nil {
		logger.Warn("StartInstance UpdateInstanceById sql error:", instanceId, err.Error())
		return err
	}
	return sendStartInstanceEvent(logger, []string{instanceId})
}
func sendStartInstanceEvent(logger *log.Logger, instance_ids []string) error {

	logid := logger.GetPoint("logid").(string)
	msg := rabbitIronicMsgApi.StartInstancesMessage{}
	msg.ApiMessage.RequestId = logid
	msg.InstanceIds = instance_ids
	userId := logger.GetPoint("userId").(string)
	event, err := rabbitEvent.NewEvent(msg, logid, userId)
	if err != nil {
		logger.Warn("convert StartInstancesMessage 2 event error:", instance_ids, err.Error())
		return err
	}
	if err := rabbitMq.SendToApi2Scheduler(event); err != nil {
		logger.Warnf("StartInstancesMessage SendToScheduler error, msg:%s, error:%s", util.ObjToJson(msg), err.Error())
		return err
	} else {
		logger.Infof("StartInstancesMessage SendToScheduler success, msg:%s", util.ObjToJson(msg))
	}
	return nil
}
func StopInstance(logger *log.Logger, instanceId string) error {
	instance, err := instanceDao.GetInstanceById(logger, instanceId)
	if err != nil {
		logger.Warnf("GetInstanceByUuid sql error, uuid:%s, error:%s", instanceId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	if !InstanceOwnedCheck(logger, instance) {
		panic(constant.PermissionDenyForObject)
	}

	if instance.Status != instanceStatus.RUNNING {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态不是运行", instanceId), fmt.Sprintf("unsupported stop instance %s", instanceId)))
	}
	entity := &instanceDao.Instance{
		InstanceID:  instanceId,
		Status:      instanceStatus.STOPPING,
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
	}
	if err := instanceDao.UpdateInstanceById(logger, entity); err != nil {
		logger.Warn("StopInstance UpdateInstanceById sql error:", instanceId, err.Error())
		return err
	}
	return sendStopInstanceEvent(logger, []string{instanceId})
}
func sendStopInstanceEvent(logger *log.Logger, instance_ids []string) error {

	logid := logger.GetPoint("logid").(string)
	msg := rabbitIronicMsgApi.StopInstancesMessage{}
	msg.ApiMessage.RequestId = logger.GetPoint("logid").(string)
	msg.InstanceIds = instance_ids
	userId := logger.GetPoint("userId").(string)
	event, err := rabbitEvent.NewEvent(msg, logid, userId)
	if err != nil {
		logger.Warn("convert StopInstancesMessage 2 event error:", instance_ids, err.Error())
		return err
	}
	if err := rabbitMq.SendToApi2Scheduler(event); err != nil {
		logger.Warnf("sendStopInstanceEvent SendToScheduler error, msg:%s, error:%s", util.ObjToJson(event), err.Error())
		return err
	} else {
		logger.Infof("sendStopInstanceEvent SendToScheduler success, msg:%s", util.ObjToJson(event))
	}
	return nil
}
func RestartInstance(logger *log.Logger, instanceId string) error {
	instance, err := instanceDao.GetInstanceById(logger, instanceId)
	if err != nil {
		logger.Warnf("GetInstanceByUuid sql error, uuid:%s, error:%s", instanceId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	if !InstanceOwnedCheck(logger, instance) {
		panic(constant.PermissionDenyForObject)
	}

	if instance.Status != instanceStatus.RUNNING {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态不是运行", instanceId), fmt.Sprintf("unsupported restart instance %s", instanceId)))
	}
	entity := &instanceDao.Instance{
		InstanceID:  instanceId,
		Status:      instanceStatus.RESTARTING,
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
	}
	if err := instanceDao.UpdateInstanceById(logger, entity); err != nil {
		logger.Warn("RestartInstance UpdateInstanceById sql error:", instanceId, err.Error())
		return err
	}
	return sendRestartInstanceEvent(logger, []string{instanceId})
}
func sendRestartInstanceEvent(logger *log.Logger, instance_ids []string) error {

	logid := logger.GetPoint("logid").(string)
	msg := rabbitIronicMsgApi.RestartInstancesMessage{}
	msg.ApiMessage.RequestId = logid
	msg.InstanceIds = instance_ids
	userId := logger.GetPoint("userId").(string)
	event, err := rabbitEvent.NewEvent(msg, logid, userId)
	if err != nil {
		logger.Warn("convert RestartInstancesMessage 2 event error:", instance_ids, err.Error())
		return err
	}
	if err := rabbitMq.SendToApi2Scheduler(event); err != nil {
		logger.Warn("sendRestartInstanceEvent SendToScheduler error, msg:%s, error:%s", util.ObjToJson(event), err.Error())
		return err
	} else {
		logger.Infof("sendRestartInstanceEvent SendToScheduler success, msg:%s", util.ObjToJson(event))
	}
	return nil
}

func ResetInstanceStatus(logger *log.Logger, instanceId string) error {
	instance, err := instanceDao.GetInstanceById(logger, instanceId)
	if err != nil {
		logger.Warnf("GetInstanceByUuid sql error, uuid:%s, error:%s", instanceId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	if !InstanceOwnedCheck(logger, instance) {
		panic(constant.PermissionDenyForObject)
	}

	//可以在此增加一些逻辑判断，比如，只能处在关机模式下才可以开机
	resetArr := []string{instanceStatus.START_ERROR, instanceStatus.STOP_ERROR, instanceStatus.RESTART_ERROR, instanceStatus.DESTROY_ERROR}
	if !util.InArray(instance.Status, resetArr) {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态错误", instanceId), fmt.Sprintf("unsupported start instance %s", instanceId)))
	}
	entity := &instanceDao.Instance{
		InstanceID:  instanceId,
		Status:      instanceStatus.Error_Instance_To_Reset_Status[instance.Status],
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
	}

	if err := commandDao.DeleteCommandByInstanceId(logger, instanceId); err != nil {
		logger.Warn("ResetInstanceStatus DeleteCommandByInstanceId sql error:", instanceId, err.Error())
	}

	if err := instanceDao.UpdateInstanceById(logger, entity); err != nil {
		logger.Warn("ResetInstanceStatus UpdateInstanceById sql error:", instanceId, err.Error())
		return err
	}

	auditLogLogic.SaveAuditLogs(logger, instance.DeviceID, instanceId, AuditLogsType.AuditLogsResetInstanceStatus)

	return nil
}

func DeleteInstance(logger *log.Logger, instanceId string) error {
	entity, err := instanceDao.GetInstanceById(logger, instanceId)
	if err != nil {
		logger.Warnf("GetInstanceByUuid sql error, uuid:%s, error:%s", instanceId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	if !InstanceOwnedCheck(logger, entity) {
		panic(constant.PermissionDenyForObject)
	}

	// 已锁定实例不允许删除
	if entity.Locked == baseLogic.LOCKED {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("已锁定实例 %s 不能被删除", instanceId), fmt.Sprintf("locked instance %s not allowed delete", instanceId)))
	}

	if !util.InArray(entity.Status, []string{instanceStatus.STOPPED, instanceStatus.CREATE_ERROR}) {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态不是关机和创建失败，不允许被删除", instanceId), fmt.Sprintf("unsupported delete instance %s", instanceId)))
	}

	if entity.Status == instanceStatus.CREATE_ERROR { //创建错误的实例，不用走完整的删除逻辑，直接释放各种资源就行
		//释放销毁实例使用的各种资源
		if err := rInstanceSshkeyDao.DeleteInstanceSshkeyByInstanceID(logger, entity.InstanceID); err != nil {
			logger.Warn("DeleteInstance DeleteInstanceSshkeyByInstanceID sql error:", instanceId, err.Error())
		}
		deviceEntiry, err := deviceDao.GetDeviceById(logger, entity.DeviceID)
		if err != nil {
			logger.Warnf("delete error status instance.GetDeviceById error, device_id:%s, error:%s", entity.DeviceID, err.Error())
			return err
		}
		device := &deviceDao.Device{
			Sn:           deviceEntiry.Sn,
			ManageStatus: baseLogic.PUTAWAY, //删除实例以后，设备表的某些状态更改
		}
		if err := deviceDao.UpdateDeviceBySn(logger, device); err != nil {
			logger.Warnf("update device error,deviceSN:%s, error:%s", device.Sn, err.Error())
			return err
		}
		if err := commandDao.DeleteCommandBySn(logger, deviceEntiry.Sn); err != nil {
			logger.Warn("DeleteInstance DeleteCommandBySn sql error:", instanceId, err.Error())
		}

		entity.IsDel = 1
		if err := instanceDao.UpdateInstanceById(logger, entity); err != nil {
			logger.Warn("DeleteInstance UpdateInstanceById sql error:", instanceId, err.Error())
			return err
		}

		return nil
	}

	entity.Status = instanceStatus.DESTROYING
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	entity.DeletedTime = int(time.Now().Unix())
	if err := instanceDao.UpdateInstanceById(logger, entity); err != nil {
		logger.Warn("DeleteInstance UpdateInstanceById sql error:", instanceId, err.Error())
		return err
	}
	//释放销毁实例使用的各种资源
	if err := rInstanceSshkeyDao.DeleteInstanceSshkeyByInstanceID(logger, entity.InstanceID); err != nil {
		logger.Warn("DeleteInstance DeleteInstanceSshkeyByInstanceID sql error:", instanceId, err.Error())
	}
	return sendDeleteInstanceEvent(logger, instanceId)
}

//批量删除实例
func DeleteInstances(logger *log.Logger, instanceIds []string) error {

	instances := []*instanceDao.Instance{}

	for _, instanceId := range instanceIds {
		entity, err := instanceDao.GetInstanceById(logger, instanceId)
		if err != nil {
			logger.Warnf("GetInstanceByUuid sql error, uuid:%s, error:%s", instanceId, err.Error())
			panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
		}

		if !InstanceOwnedCheck(logger, entity) {
			panic(constant.PermissionDenyForObject)
		}

		// 已锁定实例不允许删除
		if entity.Locked == baseLogic.LOCKED {
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("已锁定实例 %s 不能被删除", instanceId), fmt.Sprintf("locked instance %s not allowed delete", instanceId)))
		}

		if !util.InArray(entity.Status, []string{instanceStatus.STOPPED, instanceStatus.CREATE_ERROR}) {
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态不是关机和创建失败，不允许被删除", instanceId), fmt.Sprintf("unsupported delete instance %s", instanceId)))
		}
		instances = append(instances, entity)
	}

	for _, entity := range instances {
		if entity.Status == instanceStatus.CREATE_ERROR { //创建错误的实例，不用走完整的删除逻辑，直接释放各种资源就行
			//释放销毁实例使用的各种资源
			if err := rInstanceSshkeyDao.DeleteInstanceSshkeyByInstanceID(logger, entity.InstanceID); err != nil {
				logger.Warn("DeleteInstance DeleteInstanceSshkeyByInstanceID sql error:", entity.InstanceID, err.Error())
			}
			deviceEntiry, err := deviceDao.GetDeviceById(logger, entity.DeviceID)
			if err != nil {
				logger.Warnf("delete error status instance.GetDeviceById error, device_id:%s, error:%s", entity.DeviceID, err.Error())
				return err
			}
			device := &deviceDao.Device{
				Sn:           deviceEntiry.Sn,
				ManageStatus: baseLogic.PUTAWAY, //删除实例以后，设备表的某些状态更改
			}
			if err := deviceDao.UpdateDeviceBySn(logger, device); err != nil {
				logger.Warnf("update device error,deviceSN:%s, error:%s", device.Sn, err.Error())
				return err
			}
			if err := commandDao.DeleteCommandBySn(logger, deviceEntiry.Sn); err != nil {
				logger.Warn("DeleteInstance DeleteCommandBySn sql error:", entity.InstanceID, err.Error())
			}

			entity.IsDel = 1
			if err := instanceDao.UpdateInstanceById(logger, entity); err != nil {
				logger.Warn("DeleteInstance UpdateInstanceById sql error:", entity.InstanceID, err.Error())
				return err
			}

			return nil
		}

		entity.Status = instanceStatus.DESTROYING
		entity.UpdatedBy = logger.GetPoint("username").(string)
		entity.UpdatedTime = int(time.Now().Unix())
		entity.DeletedTime = int(time.Now().Unix())
		if err := instanceDao.UpdateInstanceById(logger, entity); err != nil {
			logger.Warn("DeleteInstance UpdateInstanceById sql error:", entity.InstanceID, err.Error())
			return err
		}
		//释放销毁实例使用的各种资源
		if err := rInstanceSshkeyDao.DeleteInstanceSshkeyByInstanceID(logger, entity.InstanceID); err != nil {
			logger.Warn("DeleteInstance DeleteInstanceSshkeyByInstanceID sql error:", entity.InstanceID, err.Error())
		}
		if err := sendDeleteInstanceEvent(logger, entity.InstanceID); err != nil {
			logger.Warnf("DeleteInstances.sendDeleteInstanceEvent error, instance_id:%s, error:%s", entity.InstanceID, err.Error())
		}
	}

	return nil

}

func sendDeleteInstanceEvent(logger *log.Logger, instance_id string) error {

	logid := logger.GetPoint("logid").(string)
	msg := rabbitIronicMsgApi.DeleteInstanceMessage{}
	msg.ApiMessage.RequestId = logid
	msg.InstanceId = instance_id
	userId := logger.GetPoint("userId").(string)
	event, err := rabbitEvent.NewEvent(msg, logid, userId)
	if err != nil {
		logger.Warn("convert DeleteInstanceMessage 2 event error:", instance_id, err.Error())
		return err
	}
	if err := rabbitMq.SendToApi2Scheduler(event); err != nil {
		logger.Warn("sendDeleteInstanceEvent SendToScheduler error, msg:%s, error:%s", util.ObjToJson(event), err.Error())
		return err
	} else {
		logger.Infof("sendDeleteInstanceEvent SendToScheduler success, msg:%s", util.ObjToJson(event))
	}
	return nil
}

func GetInstanceCount(logger *log.Logger, query map[string]interface{}) (int64, error) {
	count, err := instanceDao.GetInstanceCount(logger, query)
	if err != nil {
		logger.Warnf("GetInstanceCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return 0, err
	}
	return count, nil
}

func GetInstanceList(logger *log.Logger, param *requestTypes.QueryInstancesRequest, p util.Pageable) ([]*responseTypes.Instance, int64, error) {
	userId := logger.GetPoint("userId").(string)
	offset, limit := p.Pageable2offset()

	var deviceId string
	if param.IloIP != "" || param.IPV4 != "" || param.IPV6 != "" {
		q := requestTypes.QueryDevicesRequest{
			IDcID: param.IdcID,
			IsAll: baseLogic.IS_ALL,
		}
		if param.IloIP != "" {
			q.IloIP = param.IloIP
		}
		if param.IPV4 != "" {
			q.IPV4 = param.IPV4
		}
		if param.IPV6 != "" {
			q.IPV6 = param.IPV6
		}
		devices, _, err := deviceLogic.QueryDevices(logger, q, util.Pageable{})
		if err != nil {
			logger.Warnf("GetInstanceList deviceLogic.QueryDevices error:%s", err.Error())
			return nil, 0, err
		}
		if len(devices) != 1 {
			logger.Warnf("GetInstanceList deviceLogic.QueryDevices not one, data is:%s", util.InterfaceToJson(devices))
			return nil, 0, err
		}
		deviceId = devices[0].DeviceID
	}

	query := map[string]interface{}{
		//"sn":              c.GetString("sn"),
		"project_id": param.ProjectID,
		"user_id":    userId,
		"idc_id":     param.IdcID,
	}
	if param.DeviceID != "" {
		query["device_id"] = param.DeviceID
	} else if deviceId != "" {
		query["device_id"] = deviceId
	}

	if param.InstanceID != "" {
		query["instance_id"] = param.InstanceID
	}
	if param.InstanceName != "" {
		query["instance_name.like"] = "%" + param.InstanceName + "%"
	}

	if param.DeviceTypeID != "" {
		query["device_type_id.in"] = strings.Split(param.DeviceTypeID, ",")
	}
	if param.Status != "" {
		query["status.in"] = strings.Split(param.Status, ",")
	}
	count, err := instanceDao.GetInstanceCount(logger, query)
	if err != nil {
		logger.Warnf("GetInstanceCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	entityList := []*instanceDao.Instance{}
	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = instanceDao.GetMultiInstance(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
	} else {
		entityList, err = instanceDao.GetMultiInstance(logger, query, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
	}
	if err != nil {
		logger.Warn("QueryByInstanceIds GetAllInstance sql error:", err.Error())
		return nil, 0, err
	}

	instances := []*responseTypes.Instance{}
	for _, entity := range entityList {
		instances = append(instances, InstanceEntity2Instance(logger, entity, nil)) //列表不用raidcan信息
	}

	return instances, count, nil
}

func StartInstances(logger *log.Logger, instanceIds []string) error {

	instances, err := instanceDao.GetInstancesByIds(logger, instanceIds)
	if err != nil {
		logger.Warn("GetInstancesByIds sql error:", instanceIds, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	for _, instance := range instances {
		if !InstanceOwnedCheck(logger, instance) {
			panic(constant.PermissionDenyForObject)
		}

		//可以在此增加一些逻辑判断，比如，只能处在关机模式下才可以开机，其他状态不可以。。。后续再处理这块
		if instance.Status != instanceStatus.STOPPED {
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态不是关机", instance.InstanceID), fmt.Sprintf("unsupported start instance %s", instance.InstanceID)))
		}
	}

	query := map[string]interface{}{
		"instance_id.in": instanceIds,
		"is_del":         0,
	}
	updates := map[string]interface{}{
		"status":    instanceStatus.STARTING,
		"updatedBy": logger.GetPoint("username").(string),
	}
	if err := instanceDao.UpdateInstances(logger, query, updates); err != nil {
		logger.Warn("StartInstance UpdateInstanceById sql error:", instanceIds, err.Error())
		return err
	}
	return sendStartInstanceEvent(logger, instanceIds)
}

func StopInstances(logger *log.Logger, instanceIds []string) error {

	instances, err := instanceDao.GetInstancesByIds(logger, instanceIds)
	if err != nil {
		logger.Warn("GetInstancesByIds sql error:", instanceIds, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	for _, instance := range instances {
		if !InstanceOwnedCheck(logger, instance) {
			panic(constant.PermissionDenyForObject)
		}

		//可以在此增加一些逻辑判断，比如，只能处在关机模式下才可以开机，其他状态不可以。。。后续再处理这块
		if instance.Status != instanceStatus.RUNNING {
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态不是运行中", instance.InstanceID), fmt.Sprintf("unsupported stop instance %s", instance.InstanceID)))
		}
	}

	query := map[string]interface{}{
		"instance_id.in": instanceIds,
		"is_del":         0,
	}
	updates := map[string]interface{}{
		"status":    instanceStatus.STOPPING,
		"updatedBy": logger.GetPoint("username").(string),
	}
	if err := instanceDao.UpdateInstances(logger, query, updates); err != nil {
		logger.Warn("StopInstance UpdateInstanceById sql error:", instanceIds, err.Error())
		return err
	}
	return sendStopInstanceEvent(logger, instanceIds)
}

//logger里面的userid为当前操作用户的userId,拥有者和共享者都有权限操作/20231218普通运营平台用户也拥有所有实例的全部操作权限
func InstanceOwnedCheck(logger *log.Logger, instance *instanceDao.Instance) bool {

	userId := logger.GetPoint("userId").(string)

	u, err := userDao.GetUserById(logger, userId)
	if err != nil {
		logger.Warnf("InstanceOwnedCheck.GetUserById error, userid:%s, error:%s", userId, err.Error())
		return false
	}
	if u.RoleID == baseLogic.ROLE_ADMIN_UUID { //超级管理员有权限管理一切
		return true
	}
	if u.RoleID == baseLogic.ROLE_OPERATOR_UUID { //普通运营平台用户也有权限操作所有实例
		return true
	}

	if instance == nil {
		return false
	}
	if instance.UserID == userId {
		return true
	}
	project, err := projectDao.GetProjectById(logger, instance.ProjectID)
	if err == nil {
		return false
	}
	if project.UserID == userId {
		return true
	}

	sha, _ := sharingProjectDao.GetSharingsByProjectId(logger, project.ProjectID)
	for _, s := range sha {
		if s.SharedUserID == userId {
			return true
		}
	}
	return false

}

func RestartInstances(logger *log.Logger, instanceIds []string) error {

	instances, err := instanceDao.GetInstancesByIds(logger, instanceIds)
	if err != nil {
		logger.Warn("GetInstancesByIds sql error:", instanceIds, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	for _, instance := range instances {
		if !InstanceOwnedCheck(logger, instance) {
			panic(constant.PermissionDenyForObject)
		}

		//可以在此增加一些逻辑判断，比如，只能处在关机模式下才可以开机，其他状态不可以。。。后续再处理这块
		if instance.Status != instanceStatus.RUNNING {
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态不是运行中", instance.InstanceID), fmt.Sprintf("unsupported restart instance %s", instance.InstanceID)))
		}
	}

	query := map[string]interface{}{
		"instance_id.in": instanceIds,
		"is_del":         0,
	}
	updates := map[string]interface{}{
		"status":    instanceStatus.RESTARTING,
		"updatedBy": logger.GetPoint("username").(string),
	}
	if err := instanceDao.UpdateInstances(logger, query, updates); err != nil {
		logger.Warn("restartInstances UpdateInstanceById sql error:", instanceIds, err.Error())
		return err
	}
	return sendRestartInstanceEvent(logger, instanceIds)
}

func ResetInstancePasswd(logger *log.Logger, instance_id string, passwd string) error {
	instance, err := instanceDao.GetInstanceById(logger, instance_id)
	if err != nil {
		return err
	}
	if instance.Status != instanceStatus.STOPPED && instance.Status != instanceStatus.RESETPASSWD_ERROR {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态不是已关机或重置失败状态", instance.InstanceID), fmt.Sprintf("unsupported reset password instance %s", instance.InstanceID)))

	}

	if !InstanceOwnedCheck(logger, instance) {
		panic(constant.PermissionDenyForObject)
	}

	if instance.Status == instanceStatus.RESETPASSWD_ERROR { //再次重置密码，需要把上次重置的command都删除掉
		if err := commandDao.DeleteCommandByInstanceIdAndTask(logger, instance_id, "InstanceResetPassword"); err != nil {
			logger.Warn("re_resetpasswd.DeleteCommandByInstanceIdAndTask error:", instance_id, err.Error())
			return err
		}
	}

	instance.Status = instanceStatus.RESETTING_PASSWORD
	instance.UpdatedBy = logger.GetPoint("username").(string)
	instance.UpdatedTime = int(time.Now().Unix())
	if err := instanceDao.UpdateInstanceById(logger, instance); err != nil {
		logger.Warn("UpdateInstanceById error:", instance_id, err.Error())
		return err
	}
	if err := resetPasswordSendMq(logger, instance_id, passwd, "ResetPassword"); err != nil {
		return err
	}

	return nil
}

func ResetInstancesPasswd(logger *log.Logger, instanceIds []string, passwd string) error {
	instanceEntities := []*instanceDao.Instance{}
	for _, v := range instanceIds {
		instance, err := instanceDao.GetInstanceById(logger, v)
		if err != nil {
			return err
		}
		if instance.Status != instanceStatus.STOPPED {
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态不是已关机", instance.InstanceID), fmt.Sprintf("unsupported reset password instance %s", instance.InstanceID)))

		}

		if !InstanceOwnedCheck(logger, instance) {
			panic(constant.PermissionDenyForObject)
		}
		instanceEntities = append(instanceEntities, instance)
	}

	query := map[string]interface{}{
		"instance_id.in": instanceIds,
		"is_del":         0,
	}
	updates := map[string]interface{}{
		"status":    instanceStatus.RESETTING_PASSWORD,
		"updatedBy": logger.GetPoint("username").(string),
	}
	if err := instanceDao.UpdateInstances(logger, query, updates); err != nil {
		logger.Warnf("resetInstancesPasswd UpdateInstances sql error, instanceIds:%s, error:%s", strings.Join(instanceIds, ","), err.Error())
		return err
	}
	for _, instanceId := range instanceIds {
		if err := resetPasswordSendMq(logger, instanceId, passwd, "ResetPassword"); err != nil {
			logger.Warnf("resetInstancesPasswd resetPasswordSendMq error, instanceId:%s, error:%s", instanceId, err.Error())
		}
	}
	return nil
}

func resetPasswordSendMq(logger *log.Logger, instance_id, password, action string) error {

	logid := logger.GetPoint("logid").(string)
	msg := rabbitIronicMsgApi.InstanceResetPasswordMessage{}
	msg.ApiMessage.RequestId = logid
	msg.InstanceId = instance_id
	msg.Password = password
	msg.Action = action
	userId := logger.GetPoint("userId").(string)
	event, err := rabbitEvent.NewEvent(msg, logid, userId)
	if err != nil {
		logger.Warn("convert InstanceResetPasswordMessage 2 event error:", instance_id, err.Error())
		return err
	}
	if err := rabbitMq.SendToApi2Scheduler(event); err != nil {
		logger.Warnf("resetPasswordSendMq SendToScheduler error, msg:%s, error:%s", util.ObjToJson(event), err.Error())
		return err
	} else {
		logger.Infof("resetPasswordSendMq SendToScheduler success, msg:%s", util.ObjToJson(event))
	}
	return nil
}

func ReinstallInstance(logger *log.Logger, instance_id string, param *requestTypes.ReinstallInstanceRequest) error {

	instance, _ := instanceDao.GetInstanceById(logger, instance_id)
	if instance == nil || !InstanceOwnedCheck(logger, instance) {
		panic(constant.PermissionDenyForObject)
	}

	if instance.Status != instanceStatus.STOPPED && instance.Status != instanceStatus.REINSTALL_ERROR {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 状态不是已关机或重装失败状态", instance.InstanceID), fmt.Sprintf("unsupported reset password instance %s", instance.InstanceID)))
	}

	image_instance, _ := imageLogic.GetByImageId(logger, param.ImageID)
	if image_instance == nil {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 所选镜像信息错误", instance.InstanceID), fmt.Sprintf("unsupported image for instance %s", instance.InstanceID)))

	}

	// if image_instance.Format == imageFormat.QCOW2 && image_param.Format != imageFormat.QCOW2 {
	// 	//EMPTY
	// }
	device, err := deviceLogic.GetAndCheckById(logger, instance.DeviceID)
	if err != nil {
		return err
	}

	q := requestTypes.QueryDeviceTypeImageRequest{
		DeviceTypeID: device.DeviceTypeID,
		ImageID:      param.ImageID,
	}
	p := util.Pageable{
		PageSize: 1,
	}
	imageE, _, _ := deviceTypeLogic.QueryDeviceTypeImage(logger, &q, p)
	if imageE == nil {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 所选镜像和机型不匹配", instance.InstanceID), fmt.Sprintf("selected image and device_type not match for instance %s", instance.InstanceID)))

	}
	instance.ImageID = param.ImageID

	var keep_data bool = true //默认保留数据

	if param.SshKeyID != "" {
		ssh, _ := sshkeyLogic.GetSshkeyById(logger, param.SshKeyID)
		if ssh == nil {
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 所选sshkey非法", instance.InstanceID), fmt.Sprintf("validated sshkey for instance %s", instance.InstanceID)))

		}
	}
	if param.HostName != "" {
		instance.Hostname = param.HostName
	}
	if param.Password != "" {
		param.SshKeyID = "" //如果设置了自定义密码，keypairId变为空
	}
	if param.Password == "" && param.SshKeyID == "" {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("实例 %s 所选password和sshkey不能同时为空", instance.InstanceID), fmt.Sprintf("password and sshkey cannot both empty for instance %s", instance.InstanceID)))

	}
	if param.BootMode != "" {
		instance.BootMode = param.BootMode
	}

	if instance.Status == instanceStatus.REINSTALL_ERROR { //再次重装，需要把上次重装的command都删除掉
		if err := commandDao.DeleteCommandByInstanceIdAndTask(logger, instance_id, "ReinstallInstance"); err != nil {
			logger.Warn("re_reinstallinstance.DeleteCommandByInstanceIdAndTask error:", instance_id, err.Error())
			return err
		}

		//中间态的操作日志都删除掉，防止脏数据循环影响后续操作
		q := map[string]interface{}{
			"is_del":      0,
			"instance_id": instance_id,
			"result":      "doing",
		}
		u := map[string]interface{}{
			"is_del": 1,
		}
		if err := auditLogsDao.UpdateMultiAuditLogs(logger, q, u); err != nil {
			logger.Warn("re_reinstallinstance.UpdateMultiAuditLogs error:", instance_id, err.Error())
		}
	}

	instance.Status = instanceStatus.REINSTALLING
	instance.UpdatedTime = int(time.Now().Unix())
	if err := instanceDao.UpdateInstanceById(logger, instance); err != nil {
		logger.Warn("UpdateInstanceById error:", instance_id, err.Error())
		return err
	}
	saveInstancePartitions(logger, param, instance, device.DeviceTypeID)
	userData := getInstanceUserData(logger, requestTypes.CreateInstanceRequest{}, param.InstallBmpAgent)
	return sendReinstallInstanceEvent(logger, instance_id, keep_data, param.Password, userData)
}

func getInstanceUserData(logger *log.Logger, param requestTypes.CreateInstanceRequest, installBmpAgent bool) string {

	if installBmpAgent {

		bmpImageHost, err := beego.AppConfig.String("bmp_image_host")
		if err != nil {
			logger.Warnf("getInstanceUserData.get bmp_image_host error:%s", err.Error())
		}
		bmpImagePort, err := beego.AppConfig.String("bmp_image_port")
		if err != nil {
			logger.Warnf("getInstanceUserData.get bmp_image_port error:%s", err.Error())
		}

		tpl := `#!/bin/bash

		readonly WORKSPACE_DIR=/tmp/
		AGENT_BIN=bmp-agent.bin
		###########################################
		
		cd ${WORKSPACE_DIR}
		echo "10.208.12.72 bmp-agent-proxy" >> /etc/hosts
		AGENT_URL=http://bmp-agent-proxy:10000/${AGENT_BIN}
		
		
		
		curl ${AGENT_URL} -O
		chmod +x ${AGENT_BIN}
		sh +x ${AGENT_BIN}`

		content := strings.Replace(tpl, "10.208.12.72", bmpImageHost, -1)
		content = strings.Replace(content, "10000", bmpImagePort, -1)
		return base64.StdEncoding.EncodeToString([]byte(content))

	}
	return ""

}

func sendReinstallInstanceEvent(logger *log.Logger, instance_id string, keep_data bool, password string, userData string) error {

	logid := logger.GetPoint("logid").(string)
	msg := rabbitIronicMsgApi.ReinstallInstanceMessage{}
	msg.ApiMessage.RequestId = logid
	msg.InstanceId = instance_id
	msg.KeepData = keep_data
	msg.Password = password
	msg.UserData = userData
	userId := logger.GetPoint("userId").(string)
	event, err := rabbitEvent.NewEvent(msg, logid, userId)
	if err != nil {
		logger.Warn("convert ReinstallInstanceMessage 2 event error:", instance_id, err.Error())
		return err
	}
	if err := rabbitMq.SendToApi2Scheduler(event); err != nil {
		logger.Warnf("sendReinstallInstanceEvent SendToScheduler error, msg:%s, error:%s", util.ObjToJson(event), err.Error())
		return err
	} else {
		logger.Infof("sendReinstallInstanceEvent SendToScheduler success, msg:%s", util.ObjToJson(event))
	}
	return nil
}

//针对重装的情况
func saveInstancePartitions(logger *log.Logger, param *requestTypes.ReinstallInstanceRequest, instance *instanceDao.Instance, deviceTypeId string) error {
	partitions, err := buildPartitions(logger, param, instance, deviceTypeId)
	if err != nil {
		logger.Warnf("buildPartitions %s error:%s", instance.InstanceID, err.Error())
		return err
	}
	for idx, v := range partitions {
		v.InstancePartitionID = commonUtil.GetRandString("ipt-", namespacePrefix.IMAGE_ID_LENGTH, false, true, true)
		v.Number = idx + 1
		partitions[idx] = v
	}

	if err := instancePartitionDao.DeletePartitionsByInstanceId(logger, instance.InstanceID); err != nil {
		logger.Warnf("DeletePartitionsByInstanceId %s error:%s", instance.InstanceID, err.Error())
	}
	if len(partitions) > 0 {
		if _, err := instancePartitionDao.AddMultiInstancePartition(logger, partitions); err != nil {
			logger.Warnf("AddMultiInstancePartition error, partitions:%s, error:%s", util.ObjToJson(partitions), err.Error())
			return err
		}
		logger.Infof("AddMultiInstancePartition success, instanceId:%s", instance.InstanceID)
	}
	return nil
}

func buildPartitions(logger *log.Logger, param *requestTypes.ReinstallInstanceRequest, instance *instanceDao.Instance, deviceTypeId string) ([]*instancePartitionDao.InstancePartition, error) {
	if len(param.SystemPartition) == 0 && len(param.DataPartition) == 0 {
		return nil, nil
	}
	image_partitions, err := rDeviceTypeImagePartitionDao.GetByDeviceTypeAndImageId(logger, deviceTypeId, param.ImageID)
	if err != nil {
		logger.Warn("buildPartitions.GetByDeviceTypeAndImageId error:", deviceTypeId, param.ImageID, err.Error())
	}
	if len(image_partitions) == 0 {
		return nil, nil
	}
	imagePartitionEntity := image_partitions[0]
	partitions := []*instancePartitionDao.InstancePartition{}
	if len(param.SystemPartition) > 0 {
		for _, partition := range param.SystemPartition {
			system_partition := PartitionInfo2Entity(partition)
			system_partition.InstanceID = instance.InstanceID
			system_partition.ImageID = instance.ImageID
			system_partition.DeviceTypeID = deviceTypeId
			system_partition.BootMode = imagePartitionEntity.BootMode
			if strings.EqualFold(system_partition.PartitionMountPoint, "/") {
				system_partition.PartitionType = "root"
			} else if strings.EqualFold(system_partition.PartitionMountPoint, "/boot") {
				system_partition.PartitionType = "boot"
			} else {
				system_partition.PartitionType = "system"
			}
			partitions = append(partitions, &system_partition)
		}
	}

	if len(param.DataPartition) > 0 {
		for _, partition := range param.DataPartition {
			data_partition := PartitionInfo2Entity(partition)
			data_partition.InstanceID = instance.InstanceID
			data_partition.ImageID = instance.ImageID
			data_partition.DeviceTypeID = deviceTypeId
			data_partition.BootMode = imagePartitionEntity.BootMode
			data_partition.PartitionType = "data"
			partitions = append(partitions, &data_partition)
		}
	}

	return partitions, nil

}

func PartitionInfo2Entity(param requestTypes.Partition) instancePartitionDao.InstancePartition {
	res := instancePartitionDao.InstancePartition{
		PartitionSize:       param.Size,
		PartitionFsType:     param.FsType,
		PartitionMountPoint: param.MountPoint,
	}
	if strings.EqualFold(param.MountPoint, "/") {
		res.PartitionLabel = "l_root"
	} else if strings.EqualFold(param.MountPoint, "swap") {
		res.PartitionLabel = "l_swap"
	} else {
		label := strings.Replace(param.MountPoint, "/", "l_", -1)
		res.PartitionLabel = label
	}
	res.SystemDiskLabel = "gpt"
	res.DataDiskLabel = "gpt"
	return res
}

func GetInstanceBySn(logger *log.Logger, sn string) (*responseTypes.Instance, error) {
	deviceEntity, err := deviceDao.GetBySn(logger, sn)
	if err != nil {
		logger.Warnf("GetInstanceBySn.GetBySn sql error, sn:%s, error:%s", sn, err.Error())
	}
	deviceId := deviceEntity.DeviceID
	instanceEntity, err := instanceDao.GetInstanceByDeviceId(logger, deviceId)
	if err != nil {
		logger.Warnf("GetInstanceByDeviceId sql error, uuid:%s, error:%s", deviceId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	return &responseTypes.Instance{
		InstanceID:   instanceEntity.InstanceID,
		InstanceName: instanceEntity.InstanceName,
		DeviceID:     instanceEntity.DeviceID,
		Sn:           sn,
		IloIP:        deviceEntity.IloIP,
		DeviceTypeID: deviceEntity.DeviceTypeID,
	}, nil
}
