package command

import (
	"fmt"
	"time"

	constant "coding.jd.com/aidc-bmp/bmp-driver/constant"
	ipmitool "coding.jd.com/aidc-bmp/bmp-driver/ipmitool"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
)

type SetPXEBoot struct {
	*BaseCommand
	driver.SetPXEBoot
}

func init() {
	commands = append(commands, &SetPXEBoot{})
}

func (d *SetPXEBoot) Accept(action string) (Commandor, bool) {
	if action == "SetPXEBoot" {
		return &SetPXEBoot{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *SetPXEBoot) ExecuteBefore() {

}

func (d *SetPXEBoot) Execute() error {

	res, ret_code, err := ipmitool.SetBootDevice(d.Logger, d.IloIp, d.Username, d.Password, constant.PXE)
	if err != nil {
		d.Logger.Warn("SetPXEBoot SetBootDevice "+d.IloIp+" execute error:", err.Error())
		return err
	}
	d.Logger.Info("SetPXEBoot_SetBootDevice_res"+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
	time.Sleep(3 * time.Second)
	res, ret_code, err = ipmitool.PowerOn(d.Logger, d.IloIp, d.Username, d.Password)
	if err != nil {
		d.Logger.Warn("SetPXEBoot PowerOn "+d.IloIp+" execute error:", err.Error())
		return err
	}
	d.Logger.Info("SetPXEBoot_PowerOn_res"+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
	time.Sleep(3 * time.Second)
	res, ret_code, err = ipmitool.PowerReset(d.Logger, d.IloIp, d.Username, d.Password)
	if err != nil {
		d.Logger.Warn("SetPXEBoot PowerReset "+d.IloIp+" execute error:", err.Error())
		return err
	}
	time.Sleep(10 * time.Second)
	d.Logger.Info("SetPXEBoot_PowerReset_res"+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
	return nil

}

func (d SetPXEBoot) ExecuteAfter() {
	return
}
