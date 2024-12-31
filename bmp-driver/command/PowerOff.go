package command

import (
	"fmt"
	"strings"
	"time"

	ipmitool "coding.jd.com/aidc-bmp/bmp-driver/ipmitool"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
)

type PowerOff struct {
	*BaseCommand
	driver.PowerOff
}

func init() {

	commands = append(commands, &PowerOff{})
}

func (d *PowerOff) Accept(action string) (Commandor, bool) {
	if action == "PowerOff" {
		return &PowerOff{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *PowerOff) ExecuteBefore() {
	//TODO

}

func (d *PowerOff) Execute() error {

	res, _, err := ipmitool.PowerStatus(d.Logger, d.IloIp, d.Username, d.Password)
	if err != nil {
		d.Logger.Warn("PowerOff PowerStatus "+d.IloIp+" execute error:", err.Error())
		return err
	}
	res = strings.TrimSpace(res)

	// fmt.Println("getIloipPowerStatus debug...", iloip, retStr)
	items := strings.Split(res, " ")
	status := strings.ToLower(items[len(items)-1])
	if status != "off" {
		res, ret_code, err := ipmitool.PowerSoft(d.Logger, d.IloIp, d.Username, d.Password)
		if err != nil {
			d.Logger.Warn("PowerOff PowerSoft "+d.IloIp+" execute error:", err.Error())
			return err
		}
		time.Sleep(1 * time.Second)
		d.Logger.Info("PowerOff_PowerSoft_res "+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
		res, ret_code, err = ipmitool.PowerOff(d.Logger, d.IloIp, d.Username, d.Password)
		if err != nil {
			d.Logger.Warn("PowerOff PowerOff "+d.IloIp+" execute error:", err.Error())
			return err
		}
		d.Logger.Info("PowerOff_PowerOff_res "+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
		time.Sleep(3 * time.Second)

	}
	d.Logger.Info("PowerOff_execute success " + d.IloIp)
	return nil
}

func (d *PowerOff) ExecuteAfter() {
	//TODO

}
