package command

import (
	"fmt"
	"time"

	ipmitool "coding.jd.com/aidc-bmp/bmp-driver/ipmitool"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
)

type PowerOn struct {
	*BaseCommand
	driver.PowerOn
}

func init() {

	commands = append(commands, &PowerOn{})
}

func (d *PowerOn) Accept(action string) (Commandor, bool) {
	if action == "PowerOn" {
		return &PowerOn{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *PowerOn) ExecuteBefore() {
	//TODO

}

func (d *PowerOn) Execute() error {
	res, ret_code, err := ipmitool.PowerOn(d.Logger, d.IloIp, d.Username, d.Password)
	if err != nil {
		d.Logger.Warn("PowerOn PowerOn "+d.IloIp+" execute error:", err.Error())
		return err
	}
	time.Sleep(3 * time.Second)
	d.Logger.Info("PowerOn_PowerOn_res "+d.IloIp, fmt.Sprintf("%d__%s", ret_code, res))
	return nil
}

func (d *PowerOn) ExecuteAfter() {
	//TODO

}
