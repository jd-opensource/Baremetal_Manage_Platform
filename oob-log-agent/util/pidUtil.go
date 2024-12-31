package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func pid() {
	iPid := fmt.Sprint(os.Getpid())
	tmpDir := os.TempDir()
	if err := ProcExist(tmpDir); err == nil {
		pidFile, _ := os.Create(tmpDir + "\\process.pid")
		defer pidFile.Close()
		pidFile.WriteString(iPid)
	} else {
		os.Exit(1)
	}
}

// ProcExist 判断进程是否启动
func ProcExist(tmpDir string) (err error) {
	iPidFile, err := os.Open(tmpDir + "\\process.pid")
	defer iPidFile.Close()
	if err == nil {
		filePid, err := ioutil.ReadAll(iPidFile)
		if err == nil {
			pidStr := fmt.Sprintf("%s", filePid)
			pid, _ := strconv.Atoi(pidStr)
			_, err := os.FindProcess(pid)
			if err == nil {
				return errors.New("[WARN] Process has been started")
			}
		}
	}
	return nil
}
