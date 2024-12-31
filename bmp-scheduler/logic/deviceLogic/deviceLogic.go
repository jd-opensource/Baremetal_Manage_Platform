package deviceLogic

import (
	deviceDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"fmt"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

func GetById(logger *log.Logger, deviceId string) (*deviceDao.Device, error) {
	return deviceDao.GetDeviceById(logger, deviceId)
}

func GetBySn(logger *log.Logger, sn string) (*deviceDao.Device, error) {
	return deviceDao.GetDeviceBySn(logger, sn)
}

func GetDelBySn(logger *log.Logger, sn string) (*deviceDao.Device, error) {
	param := map[string]interface{}{
		"is_del": 1,
		"sn":     sn,
	}
	devices, err := deviceDao.GetMultiDevice(logger, param, nil, []string{"update_time"}, []string{"desc"}, 0, 1)
	if devices != nil {
		return devices[0], nil
	}
	return nil, err
}

func UpdateDeviceBySn(logger *log.Logger, m *deviceDao.Device) (err error) {
	return deviceDao.UpdateDeviceBySn(logger, m)
}

func CheckDevice(logger *log.Logger, event ironicAgentEvent.CallbackCommandMessage) error {
	manageStatus := "putaway" //已上架
	reason := ""
	if event.Status != "OK" { //ERROR
		manageStatus = "putawayfail"
		reason = event.Status
	}
	fmt.Println(manageStatus, reason)
	device := &deviceDao.Device{
		Sn:           event.Sn,
		ManageStatus: manageStatus,
		Reason:       reason,
		IsDel:        0,
	}
	if err := UpdateDeviceBySn(logger, device); err != nil {
		logger.Warn("设备上架中变为已上架 sql error:", event.Sn, event.Action, err.Error())
		return err
	}
	return nil
}
