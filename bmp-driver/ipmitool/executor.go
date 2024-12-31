package ipmitool

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"syscall"

	log "git.jd.com/cps-golang/log"
)

var ipmiToolPath = "/usr/sbin/ipmitool"

// ExecCommand function to execute command
func execCommand(logger *log.Logger, cmdStr string, params ...string) (retStr string, ret int, err error) {
	logger.Info("execCommand:", cmdStr+" "+strings.Join(params, " "))
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
	fmt.Println("指令：", cmdStr+" "+strings.Join(params, " "))
	cmd := exec.Command(cmdStr, params...)

	// execute
	//out, err := cmd.StdoutPipe()
	//if err != nil {
	//	fmt.Println("StdoutPipe error: ", err.Error())
	//	logger.Warn("StdoutPipe error: ", err.Error())
	//	return "", 0, err
	//}
	//// fmt.Printf("stdout>%v\n", string(out))
	//defer out.Close() // 保证关闭输出流
	//for i := 1; i <= 4; i++ {
	//	if err := cmd.Start(); err != nil { // 运行命令
	//		//log.Fatal(err)
	//		if i <= 3 {
	//			fmt.Println("exec Start error: ", err.Error(), i)
	//			logger.Warn("exec Start error: ", err.Error(), i)
	//			time.Sleep(time.Minute)
	//			continue
	//		}
	//		return "", 0, err
	//	}
	//}
	//
	//bytes, _ := ioutil.ReadAll(out)
	//for i := 1; i <= 4; i++ {
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("CombinedOutput error, command:", cmdStr, params, "error:", err.Error(), "out:", string(out))
		logger.Warn("CombinedOutput error, command:", cmdStr, params, "error:", err.Error(), "out:", string(out))
		return string(out[:]), 0, err
	}
	//}
	fmt.Println("exec output: ", string(out[:]))
	logger.Info("exec output: ", string(out[:]))
	//return string(bytes)
	// get exit code
	ret = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	if ret != 0 {
		err = errors.New(strings.TrimSpace(string(out[:])))
		logger.Warn("exec WaitStatus error: ", err.Error())
		return "", 0, err
	}
	// fmt.Printf("cmd>%+v \n", cmd)

	return string(out[:]), ret, err
}
