package ipmitool

import log "git.jd.com/cps-golang/log"

func SetBootDevice(logger *log.Logger, ip, username, password, boot_device string) (res string, ret_code int, err error) {
	res, ret_code, err = execCommand(logger, ipmiToolPath, "-I", "lanplus", "-H", ip, "-U", username, "-P", password, "chassis", "bootdev", boot_device)
	return
}
