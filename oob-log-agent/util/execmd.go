package util

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"
	"syscall"

	log "coding.jd.com/aidc-bmp/bmp_log"
)

// ExecCommand function to execute command
func ExecCommand(logger *log.Logger, cmdStr string, params ...string) (retStr string, ret int, err error) {

	defer func() {
		r := recover()
		if r != nil {
			logger.Warnf("ExecCommand catch exception, error:%v, stack:%s", r, GetCurrentGoroutineStack())
			ret = -1
			err = errors.New("ExecCommand panic when exec")
		}
	}()

	logger.Infof("ExecCommand cmdStr:%s, param:+%v", cmdStr, params)
	// check command injection
	r := regexp.MustCompile(`[&\|;]+`)
	if matched := r.MatchString(cmdStr); matched {
		return "", -1, errors.New("command can not contain [&|;], ignore execute")
	}
	for _, param := range params {
		if matched := r.MatchString(param); matched {
			return "", -1, errors.New("params can not contain [&|;], ignore execute")
		}
	}

	cmd := exec.Command(cmdStr, params...)

	// execute
	out, err := cmd.CombinedOutput()
	// fmt.Printf("stdout>%v\n", string(out))

	// get exit code
	logger.Infof("%+v ExecCommand cmdStr res debug1 cmd.out:%s", params, string(out))
	logger.Infof("%+v ExecCommand cmdStr res debug1 cmd.ProcessState:%+v", params, cmd.ProcessState)
	logger.Infof("%+v ExecCommand cmdStr res debug2 cmd.ProcessState.Sys:%+v", params, cmd.ProcessState.Sys())
	ret = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	if ret != 0 {
		err = errors.New(strings.TrimSpace(string(out[:])))
	}
	// fmt.Printf("cmd>%+v \n", cmd)

	return string(out[:]), ret, err
}
