package crontab

import (
	"runtime"

	log "coding.jd.com/aidc-bmp/bmp_log"

	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	"coding.jd.com/aidc-bmp/oob-log-alert/object"
	"coding.jd.com/aidc-bmp/oob-log-alert/service"
	"coding.jd.com/aidc-bmp/oob-log-alert/util"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func SyncDeviceToRedis() error {

	logPath, _ := web.AppConfig.String("log.path")
	logger := log.New(logPath + "/SyncDeviceToRedisCron.log")
	logger.SetStableFields([]string{"method", "uri", "header", "request", "response"})
	logger.TimeStart("all_t")
	logger.TimeStart("self_t")
	logger.Point("task_type", "SyncDeviceToRedis")
	//没有requestid，临时生成
	logger.Point("logid", util.GenerateRandUuid())
	defer logger.Flush()
	defer logger.TimeEnd("self_t")
	defer logger.TimeEnd("all_t")

	defer func() {
		if r := recover(); r != nil {
			logger.Warnf("SyncDeviceToRedis.Cron catch exception, error:%v, stack:%s", r, getCurrentGoroutineStack())
		}
	}()

	logger.Info("doCheckTimeoutCommandCron starting......")
	rIdc, _ := web.AppConfig.String("app.idc")
	devices, err := dao.GetCpssByCondNoLimit(logger, "", rIdc)
	if err != nil {
		logger.Warnf("SyncDeviceToRedis.GetCpssByCondNoLimit error, idc:%s, err:%s", rIdc, err.Error())
		return err
	}
	cpsList := devices2cpsList(rIdc, devices)
	logger.Infof("devices ready to write to redis:%s", util.ToString(cpsList))
	// remove
	var delSNs []string
	oldCPSList, err := service.GetCPSFromRedis()
	if err != nil {
		logs.Warn("SyncCPSFromAPIByRegionAz() GetCPSFromRedis Error:%s", err.Error())
	} else {
		for _, old := range oldCPSList.CPSs {
			isFound := false
			for _, n := range cpsList.CPSs {
				if old.SN == n.SN {
					isFound = true
					break
				}
			}
			if !isFound {
				delSNs = append(delSNs, old.SN)
			}
		}

		if len(delSNs) > 0 {
			err = service.DelCPSFromRedis(delSNs)
			if err != nil {
				logs.Error("delCPSFromRedis() Error : " + err.Error())
			}
		}
	}

	// logs.Debug(cpsList)
	if len(cpsList.CPSs) == 0 {
		return nil
	}

	idc, _ := dao.GetIdcById(logger, rIdc)
	for k, v := range cpsList.CPSs {
		if v.IloUser == "" || v.IloPass == "" { //设备的带外用户名密码如果为空，用机房的带外用户名密码来代替
			v.IloUser = idc.IloUser
			v.IloPass = idc.IloPassword
		}
		cpsList.CPSs[k] = v
	}

	err = service.WriteCPSListToRedis(cpsList)
	if err != nil {
		logs.Error("WriteCPSListToRedis() Error : " + err.Error())
		return err
	}
	return nil
}

func devices2cpsList(idc string, devices []*dao.Device) object.CPSObject {
	o := object.CPSObject{
		Idc:  idc,
		CPSs: []object.CPSRecord{},
	}
	for _, device := range devices {
		cpsObject := device2cpsObject(device)
		o.CPSs = append(o.CPSs, cpsObject)
	}
	return o
}

func device2cpsObject(device *dao.Device) object.CPSRecord {
	return object.CPSRecord{
		Clazz:      device.DeviceTypeID,
		SN:         device.Sn,
		IdcId:      device.IDcID,
		InstanceID: device.InstanceID,
		// InstanceName: device.in,
		Status:  device.ManageStatus,
		IloIP:   device.IloIP,
		Brand:   device.Brand,
		Model:   device.Model,
		Ipv4:    device.PrivateIPv4,
		IloUser: device.IloUser,
		IloPass: device.IloPassword,
	}
}

// getCurrentGoroutineStack 获取当前Goroutine的调用栈，便于排查panic异常
func getCurrentGoroutineStack() string {

	const defaultStackSize = 4096

	var buf [defaultStackSize]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}
