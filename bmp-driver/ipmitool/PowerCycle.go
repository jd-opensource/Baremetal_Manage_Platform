package ipmitool

import log "git.jd.com/cps-golang/log"

// 注意power cycle 和power reset的区别在于前者从掉电到上电有１秒钟的间隔，而后者是很快上电
func PowerCycle(logger *log.Logger, ip, username, password string) (res string, ret_code int, err error) {
	res, ret_code, err = execCommand(logger, ipmiToolPath, "-I", "lanplus", "-H", ip, "-U", username, "-P", password, "chassis", "power", "cycle")
	return
}
