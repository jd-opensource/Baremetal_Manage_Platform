package ipmitool

import log "git.jd.com/cps-golang/log"

func UserList(logger *log.Logger, ip, username, password string) (res string, ret_code int, err error) {
	res, ret_code, err = execCommand(logger, ipmiToolPath, "-c", "-I", "lanplus", "-H", ip, "-U", username, "-P", password, "user", "list")
	return
}
