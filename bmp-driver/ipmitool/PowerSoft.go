package ipmitool

import log "git.jd.com/cps-golang/log"

func PowerSoft(logger *log.Logger, ip, username, password string) (res string, ret_code int, err error) {
	res, ret_code, err = execCommand(logger, ipmiToolPath, "-I", "lanplus", "-H", ip, "-U", username, "-P", password, "chassis", "power", "soft")
	return
}
