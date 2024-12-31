package command

import (
	"fmt"
	"os/exec"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-driver/errors"
	ipmitool "coding.jd.com/aidc-bmp/bmp-driver/ipmitool"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
)

type CheckInitConfig struct {
	*BaseCommand
	driver.CheckInitConfig
}

func init() {

	commands = append(commands, &CheckInitConfig{})
}

func (d *CheckInitConfig) Accept(action string) (Commandor, bool) {
	if action == "CheckInitConfig" {
		return &CheckInitConfig{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *CheckInitConfig) ExecuteBefore() {
	//TODO

}

func (d *CheckInitConfig) Execute() (err error) {
	err = d.checkIloipStatus(d.IloIp)
	if err != nil {
		return
	}
	if d.SubnetGateway != "" {
		err = d.checkSubnetGatewayStatus(d.SubnetGateway)
		if err != nil {
			return
		}
	}
	err = d.checkIPMIPrivilege(d.IloIp, d.Username, d.Password, d.Privilege)
	if err != nil {
		return
	}
	return
}

func (d *CheckInitConfig) ExecuteAfter() {
	//TODO

}

func (d *CheckInitConfig) checkIPMIPrivilege(ip, username, password, privilege string) (err error) {
	curPrivilege, err := d.getIPMIUserPrivilege(ip, username, password)
	if err != nil {
		return
	}
	if strings.ToLower(curPrivilege) == strings.ToLower(privilege) {
		return
	}
	errMsg := fmt.Sprintf(errors.USER_PRIVILEGE_ERROR_MESSAGE, username, curPrivilege, privilege)
	err = errors.NewBmpError(errors.USER_PRIVILEGE_ERROR_CODE, errMsg, nil)
	d.Logger.Warn(err.Error())
	return
}

func (d *CheckInitConfig) checkSubnetGatewayStatus(subnetGateway string) (err error) {
	err = d.checkNetWorkStatus(subnetGateway)
	if err != nil {
		err = errors.NewBmpError(errors.PING_SUBNET_GATEWAY_ERROR_CODE, errors.PING_SUBNET_GATEWAY_ERROR_MESSAGE, err)
		d.Logger.Warn(err.Error())
	}
	return
}

func (d *CheckInitConfig) checkIloipStatus(iloIP string) (err error) {
	err = d.checkNetWorkStatus(iloIP)
	if err != nil {
		err = errors.NewBmpError(errors.PING_ILOIP_ERROR_CODE, errors.PING_ILOIP_ERROR_MESSAGE, err)
		d.Logger.Warn(err.Error())
	}
	return
}

func (d *CheckInitConfig) checkNetWorkStatus(address string) (err error) {
	cmd := exec.Command("ping", address, "-c", "1", "-W", "5")
	res, err := cmd.CombinedOutput()
	if err != nil {
		errMsg := fmt.Sprintf(errors.NETWORK_STATUS_ERROR_MESSAGE, string(res[:]))
		err = errors.NewBmpError(errors.NETWORK_STATUS_ERROR_CODE, errMsg, err)
		d.Logger.Warn(err.Error())
	}
	return
}

func (d *CheckInitConfig) getIPMIUserPrivilege(ip, username, password string) (privilege string, err error) {

	res, retCode, err := ipmitool.UserList(d.Logger, ip, username, password)
	if err != nil {
		errMsg := fmt.Sprintf(errors.USERNAME_OR_PASSWORD_ERROR_MESSAGE, res)
		err = errors.NewBmpError(errors.USERNAME_OR_PASSWORD_ERROR_CODE, errMsg, err)
		d.Logger.Warn(err.Error())
		return
	}
	d.Logger.Info("Get ipmi user list res"+ip, fmt.Sprintf("%d__%s", retCode, res))
	//Return csv data format
	//1,,true,false,false,NO ACCESS
	//2,jdroot,true,true,true,USER
	//3,jdcps,true,true,true,ADMINISTRATOR
	userListSlice := strings.Split(string(res), "\n")

	for _, v := range userListSlice {
		vSlice := strings.Split(v, ",")
		if len(vSlice) < 6 {
			continue
		}
		if vSlice[1] == username {
			privilege = vSlice[5]
			return
		}
	}
	errMsg := fmt.Sprintf(errors.USERNAME_OR_PASSWORD_ERROR_MESSAGE, username)
	err = errors.NewBmpError(errors.USERNAME_OR_PASSWORD_ERROR_CODE, errMsg, nil)
	d.Logger.Warn(err.Error())
	return
}
