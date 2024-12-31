package instanceLogic

import (
	"fmt"
	"time"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/imageLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	"git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/ImageType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/InstanceStatus"
	log "git.jd.com/cps-golang/log"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func CallBack(logger *log.Logger, instance_id string) {
	if instance_id == "" {
		logger.Warn("instanceLogic CallBack empty instance_id!!!")
		return
	}

	//var callbackMsg interface{}
	logid := logger.GetPoint("logid").(string)
	region := "" //logger.GetPoint("region").(string)
	az := ""     //logger.GetPoint("az").(string)
	fmt.Println(logid, region, az)
	instance, err := GetByInstanceId(logger, instance_id)
	if instance == nil {
		logger.Warn("instanceLogic CallBack empty instanceId "+instance_id, err.Error())
		return
	}
	fmt.Println("实例信息为", instance.InstanceID, instance.InstanceName, instance.Status, "准备写入redis")
	if err != nil {
		logger.Warn("instanceLogic CallBack GetByInstanceId sql error:", instance_id, err.Error())
	}
	image, err := imageLogic.GetByImageId(logger, instance.ImageID)
	if err != nil {
		logger.Warnf("instanceLogic CallBack GetByImageId sql error, instance_id:%s, image_id:%s, error:%s:", instance_id, instance.ImageID, err.Error())
	}
	// device, err := deviceLogic.GetById(logger, instance.DeviceID)
	// if err != nil {
	// 	logger.Warn("instanceLogic CallBack GetByDeviceId sql error:", err.Error())
	// }
	//下面的代码是根据将该实例状态当前的状态，得出指令执行结束以后的状态，并且将这个状态写入redis，等待crontab在接下来的时间内执行完毕（不要实时更新实例状态，有可能因为重启等操作，会有一定时间的延迟），并将实例状态更新
	switch strings.ToUpper(instance.Status) {
	case "CREATING":
		logger.Info("instancelogic callback CREATING......", instance_id)
		//callbackMsg = commonCpsEvent.BuildCreateInstanceCallbackMessage(logid, region, az, instance_id, "instance.PublicIP")
		if ImageType.IsWindows(image.ImageName) {
			emulatorInstanceStatus(logger, instance_id, InstanceStatus.RUNNING, Constants.TIMEOUT_INSTANCE_RESTART_WINDOWS)
		} else {
			emulatorInstanceStatus(logger, instance_id, InstanceStatus.RUNNING, Constants.TIMEOUT_INSTANCE_RESTART)
		}
		logger.Info("instance create finished, emulator status to running, instance_id: ", instance_id)
		break
	case "STARTING":
		logger.Info("instancelogic callback STARTING......", instance_id)
		//callbackMsg = commonCpsEvent.BuildStartInstanceCallbackMessage(logid, region, az, instance_id)
		emulatorInstanceStatus(logger, instance_id, InstanceStatus.RUNNING, Constants.TIMEOUT_INSTANCE_START)
		logger.Info("instance start finished, emulator status to running, instance_id: ", instance_id)
	case "STOPPING":
		logger.Info("instancelogic callback STOPPING......", instance_id)
		//callbackMsg = commonCpsEvent.BuildStopInstanceCallbackMessage(logid, region, az, instance_id)
		emulatorInstanceStatus(logger, instance_id, InstanceStatus.STOPPED, Constants.TIMEOUT_INSTANCE_STOP)
		logger.Info("instance stop finished, emulator status to stopped, instance_id: ", instance_id)
		break
	case "REINSTALLING":
		logger.Info("instancelogic callback REINSTALLING......", instance_id)
		//callbackMsg = commonCpsEvent.BuildReinstallInstanceCallbackMessage(logid, region, az, instance_id)
		if ImageType.IsWindows(image.ImageName) {
			emulatorInstanceStatus(logger, instance_id, InstanceStatus.RUNNING, Constants.TIMEOUT_INSTANCE_RESTART_WINDOWS)
		} else {
			emulatorInstanceStatus(logger, instance_id, InstanceStatus.RUNNING, Constants.TIMEOUT_INSTANCE_RESTART)
		}
		logger.Info("instance reinstall finished, emulator status to running, instance_id: ", instance_id)
		break
	case "RESTARTING":
		logger.Info("instancelogic callback RESTARTING......", instance_id)
		//callbackMsg = commonCpsEvent.BuildRestartInstanceCallbackMessage(logid, region, az, instance_id)
		emulatorInstanceStatus(logger, instance_id, InstanceStatus.RUNNING, Constants.TIMEOUT_INSTANCE_RESTART)
		logger.Info("instance restart finished, emulator status to running, instance_id: ", instance_id)
		break
	case "UPGRADING":
		logger.Info("instancelogic callback UPGRADING......", instance_id)
		//callbackMsg = commonCpsEvent.BuildUpgradeInstanceCallbackMessage(logid, region, az, instance_id)
		instance_key := fmt.Sprintf("%s:%s", Constants.REDIS_KEY_STATUS, instance_id)
		status, err := redis.GetObjectFromRedis(instance_key)
		if err != nil {
			logger.Warn("UPGRADING callback get from redis error, key:", instance_key, err.Error())
			return
		}
		instance.Status = status
		instance.UpdatedTime = int(time.Now().Unix())
		if err := instanceDao.UpdateInstanceById(logger, instance); err != nil {
			logger.Warn("UPGRADING callback save instance error:", err.Error())
			return
		}
		if err := redis.DelObjectFromRedis([]string{instance_key}); err != nil {
			logger.Warn("UPGRADING callback del redis error:", instance_key, err.Error())
			return
		}
		logger.Info("instance upgrade finished, emulator status to:", status, "instance_id:", instance_id)
		break
	case "DESTROYING":
		logger.Info("instancelogic callback DESTROYING......", instance_id)
		//callbackMsg = commonCpsEvent.BuildDestroyInstanceCallbackMessage(logid, region, az, instance_id, "instance.PublicIP")
		instance.Status = InstanceStatus.DESTROYED
		instance.IsDel = 1
		instance.DeletedTime = int(time.Now().Unix())
		instance.UpdatedTime = int(time.Now().Unix())
		if err := instanceDao.UpdateInstanceById(logger, instance); err != nil {
			logger.Warn("DESTROYING callback save instance error:", err.Error())
			return
		}
		device := &deviceDao.Device{
			InstanceID:   instance_id,
			ManageStatus: "putaway", //实例销毁以后，设备状态变为已上架（可售卖），同时将设备对应的用户信息置空
		}
		if err := deviceDao.UpdateDeviceByInstanceId(logger, device); err != nil {
			logger.Warnf("instance destroy finished, change status to destroyed err. instance_id:%s, err:%s", instance_id, err.Error())
		}
		logger.Info("instance destroy finished, change status to destroyed. instance_id:", instance_id)
		break
	case "RESETTING_PASSWORD":
		logger.Info("instancelogic callback RESETTING_PASSWORD......", instance_id)
		//callbackMsg = commonCpsEvent.BuildInstanceResetPasswordCallbackMessage(logid, region, az, instance_id)
		emulatorInstanceStatus(logger, instance_id, InstanceStatus.STOPPED,
			Constants.TIMEOUT_INSTANCE_STOP)
		logger.Info("instance reset password finished, emulator status to stopped. instance_id:", instance_id)
		break
	default:
		break
	}
	//v, _ := json.Marshal(callbackMsg)
	//if err := rabbitMq.SendToOpenAPI(callbackMsg); err != nil {
	//	logger.Warnf("instanceLogic callback send msg err, callbackMsg: %s, error: %s", string(v), err.Error())
	//}
	//
	//logger.Infof("rabbitMq.SendToOpenAPI success, callbackMsg: %s", string(v))
	logger.Info("command exec all finished,crontab redis will update the instance status in the several minutes")

}

func GetByInstanceId(logger *log.Logger, instance_id string) (*instanceDao.Instance, error) {
	return instanceDao.GetInstanceByUuid(logger, instance_id)
}

func emulatorInstanceStatus(logger *log.Logger, instance_id string, status string, timeoutTime int) {
	check_key := fmt.Sprintf("%s:%s", Constants.REDIS_KEY_CHECK, instance_id)
	instance_key := fmt.Sprintf("%s:%s", Constants.REDIS_KEY_INSTANCE, instance_id)
	if err := redis.SetObjectToRedis(check_key, status); err != nil {
		logger.Warn("emulatorInstanceStatus set error:", check_key, status)
	}
	if _, err := redis.SetNxObjectToRedis(instance_key, status, timeoutTime); err != nil {
		logger.Warn("emulatorInstanceStatus set error:", instance_key, status)
	}

}

func resourceRecovery(logger *log.Logger, instance_entity *instanceDao.Instance) error {
	// 设备可售卖
	device_entity, err := deviceDao.GetDeviceById(logger, instance_entity.DeviceID)
	if err != nil {
		logger.Warn("resourceRecovery GetDeviceById sql error:", err.Error())
		return err
	}
	device_entity.ManageStatus = "putaway"
	device_entity.UpdatedTime = int(time.Now().Unix())
	if err := deviceDao.UpdateDeviceById(logger, device_entity); err != nil {
		logger.Warn("resourceRecovery UpdateDeviceById sql error:", err.Error())
		return err
	}
	//if instance_entity.NetworkType == NetworkType.BASIC {
	//	// 回收公网IP
	//	if instance_entity.EnableInternet == 0 {
	//		ip_entity, err := wanIpDao.GetByAzAndIP(logger, device_entity.Az, instance_entity.PublicIP)
	//		if err != nil {
	//			logger.Warn("resourceRecovery GetByAzAndIP sql error:", err.Error())
	//		}
	//		if ip_entity != nil {
	//			ip_entity.Status = 0
	//			ip_entity.UpdateTime = time.Now()
	//			if err := wanIpDao.UpdateWanIpById(logger, ip_entity); err != nil {
	//				logger.Warn("resourceRecovery UpdateWanIpById sql error:", err.Error())
	//			}
	//		}
	//	}
	//	// 删除子网
	//	count, err := instanceDao.CountByTenantAndSubnetId(logger, instance_entity.Tenant, instance_entity.SubnetID)
	//	if err != nil {
	//		logger.Warn("resourceRecovery CountByTenantAndSubnetId sql error:", err.Error())
	//	}
	//	if count == 1 {
	//		subnet_entity, err := subnetDao.GetBySubnetIdAndTenant(logger, instance_entity.SubnetID, instance_entity.Tenant)
	//		if err != nil {
	//			logger.Warn("resourceRecovery GetBySubnetIdAndTenant sql error:", err.Error())
	//		}
	//		if subnet_entity != nil {
	//			subnet_entity.IsDel = 1
	//			subnet_entity.UpdateTime = time.Now()
	//			if err := subnetDao.UpdateSubnetById(logger, subnet_entity); err != nil {
	//				logger.Warn("resourceRecovery UpdateSubnetById sql error:", err.Error())
	//			}
	//		}
	//	}
	//}
	return nil

}
