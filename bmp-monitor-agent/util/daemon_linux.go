// +build linux

package util

import "github.com/VividCortex/godaemon"

// import "fmt"

func Daemon() error {
	_, _, err := godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	// panic("MakeDaemon error:" + err.Error())
	return err
}
