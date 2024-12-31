package command

import (
	"fmt"
	"time"

	constant "coding.jd.com/aidc-bmp/bmp-driver/constant"
	ipmitool "coding.jd.com/aidc-bmp/bmp-driver/ipmitool"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
)

type PowerReset struct {
	*BaseCommand
	driver.PowerReset
}

func init() {
	commands = append(commands, &PowerReset{})
}

func (d *PowerReset) Accept(action string) (Commandor, bool) {
	if action == "PowerReset" {
		return &PowerReset{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *PowerReset) ExecuteBefore() {
	//TODO

}

func (d *PowerReset) Execute() error {
	res, ret_code, err := ipmitool.SetBootDevice(d.Logger, d.IloIp, d.Username, d.Password, constant.DISK)
	if err != nil {
		d.Logger.Warn("PowerReset SetBootDevice "+d.IloIp+" execute error:", err.Error())
		return err
	}
	d.Logger.Info("PowerReset_SetBootDevice_res "+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
	time.Sleep(4 * time.Second)
	res, ret_code, err = ipmitool.PowerReset(d.Logger, d.IloIp, d.Username, d.Password)
	if err != nil {
		d.Logger.Warn("PowerReset PowerReset "+d.IloIp+" execute error:", err.Error())
		return err
	}
	time.Sleep(3 * time.Second)
	d.Logger.Info("PowerReset_PowerReset_res "+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
	return nil
}

func (d *PowerReset) ExecuteAfter() {
	//TODO

}
