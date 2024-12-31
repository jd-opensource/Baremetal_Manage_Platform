package command

import (
	"fmt"
	"strings"
	"time"

	constant "coding.jd.com/aidc-bmp/bmp-driver/constant"
	ipmitool "coding.jd.com/aidc-bmp/bmp-driver/ipmitool"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
)

type PowerCycle struct {
	*BaseCommand
	driver.PowerCycle
}

func init() {

	commands = append(commands, &PowerCycle{})
}

func (d *PowerCycle) Accept(action string) (Commandor, bool) {
	if action == "PowerCycle" {
		return &PowerCycle{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *PowerCycle) ExecuteBefore() {
	//TODO

}

func (d *PowerCycle) Execute() error {
	res, ret_code, err := ipmitool.SetBootDevice(d.Logger, d.IloIp, d.Username, d.Password, constant.DISK)
	if err != nil {
		d.Logger.Warn("PowerCycle SetBootDevice "+d.IloIp+" execute error:", err.Error())
		return err
	}
	d.Logger.Info("PowerCycle_SetBootDevice_res"+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
	time.Sleep(4 * time.Second)

	res, _, err = ipmitool.PowerStatus(d.Logger, d.IloIp, d.Username, d.Password)
	if err != nil {
		d.Logger.Warn("PowerCycle_PowerStatus "+d.IloIp+" execute error:", err.Error())
		return err
	}
	res = strings.TrimSpace(res)

	// fmt.Println("getIloipPowerStatus debug...", iloip, retStr)
	items := strings.Split(res, " ")
	status := strings.ToLower(items[len(items)-1])
	if status != "on" {
		res, ret_code, err = ipmitool.PowerOn(d.Logger, d.IloIp, d.Username, d.Password)
		if err != nil {
			d.Logger.Warn("PowerCycle_PowerOn "+d.IloIp+" execute error:", err.Error())
			return err
		}
		time.Sleep(3 * time.Second)
	}

	res, ret_code, err = ipmitool.PowerCycle(d.Logger, d.IloIp, d.Username, d.Password)
	if err != nil {
		d.Logger.Warn("PowerCycle_PowerCycle "+d.IloIp+" execute error:", err.Error())
		return err
	}
	time.Sleep(3 * time.Second)
	d.Logger.Info("PowerCycle_PowerCycle_res "+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
	return nil

}

func (d *PowerCycle) ExecuteAfter() {
	//TODO

}
