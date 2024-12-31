package util

import (
	"os"
	"strconv"

	"github.com/astaxie/beego/logs"
)

type Pidfile struct {
	Pathfile string
}

var _localpid *Pidfile

func NewPID(pathfile string) {
	_localpid = New(pathfile)
}

// New creates a new Pidfile and writes the current PID
func New(pathfile string) *Pidfile {
	file, err := os.OpenFile(pathfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeExclusive|os.ModePerm)
	if err != nil {
		logs.Warn("pidfile: failed to open %s (%s)", pathfile, err)
		return nil
	}
	defer file.Close()
	pid := strconv.Itoa(os.Getpid())
	logs.Info("start process pid:%v --> %s", pid, pathfile)
	file.WriteString(pid)
	return &Pidfile{pathfile}
}

func (pf *Pidfile) Remove() {
	err := os.Remove(pf.Pathfile)
	if err != nil {
		logs.Warn("pidfile: failed to remove %s (%s)", pf.Pathfile, err)
	}
}
func RemovePID() {
	defer func() { recover() }()
	if _localpid != nil {
		_localpid.Remove()
	}
}
