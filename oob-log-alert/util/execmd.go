package util

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

// ExecCommand function to execute command
func ExecCommand(cmdStr string, params ...string) (retStr string, ret int, err error) {

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
	ret = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	if ret != 0 {
		err = errors.New(strings.TrimSpace(string(out[:])))
	}
	// fmt.Printf("cmd>%+v \n", cmd)

	return string(out[:]), ret, err
}
