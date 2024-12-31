package command

import (
	"fmt"
	"time"

	"coding.jd.com/aidc-bmp/bmp-driver/constant"
	ipmitool "coding.jd.com/aidc-bmp/bmp-driver/ipmitool"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
)

type SetDISKBoot struct {
	*BaseCommand
	driver.SetDISKBoot
}

func init() {

	commands = append(commands, &SetDISKBoot{})
}

func (d *SetDISKBoot) Accept(action string) (Commandor, bool) {
	if action == "SetDISKBoot" {
		return &SetDISKBoot{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *SetDISKBoot) ExecuteBefore() {
	//TODO

}

func (d *SetDISKBoot) Execute() error {
	res, ret_code, err := ipmitool.SetBootDevice(d.Logger, d.IloIp, d.Username, d.Password, constant.DISK)
	if err != nil {
		d.Logger.Warn("SetDISKBoot SetBootDevice "+d.IloIp+" execute error:", err.Error())
		return err
	}
	d.Logger.Info("SetDISKBoot_SetBootDevice_res"+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
	time.Sleep(3 * time.Second)
	res, ret_code, err = ipmitool.PowerOn(d.Logger, d.IloIp, d.Username, d.Password)
	if err != nil {
		d.Logger.Warn("SetDISKBoot PowerOn "+d.IloIp+" execute error:", err.Error())
		return err
	}
	d.Logger.Info("SetDISKBoot_PowerOn_res"+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
	time.Sleep(3 * time.Second)
	// res, ret_code, err = ipmitool.PowerReset(d.IloIp, d.Username, d.Password)
	// if err != nil {
	// 	d.Logger.Warn("SetDISKBoot PowerReset execute error:", err.Error())
	// 	panic(err)
	// }
	// time.Sleep(3 * time.Second)
	d.Logger.Info("SetDISKBoot_PowerReset_res"+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
	return nil

}

func (d *SetDISKBoot) ExecuteAfter() {
	//TODO

}
