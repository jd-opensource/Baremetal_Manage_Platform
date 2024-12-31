// +build darwin

package util

import "github.com/VividCortex/godaemon"

func Daemon() error {
	_, _, err := godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	return err
}
