package util

import (
	"os/exec"
	"path"
)

func CmdRunWithOutput(command string) (string, error) {
	var out []byte
	var err error = nil

	if path.Ext(command) == ".sh" {
		// shell
		out, err = exec.Command("/bin/sh", command).CombinedOutput()
	} else {
		// command
		out, err = exec.Command("/bin/sh", "-c", command).CombinedOutput()
	}

	return string(out), err
}
