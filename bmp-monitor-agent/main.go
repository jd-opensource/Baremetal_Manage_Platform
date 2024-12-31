package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"coding.jd.com/bmp/cps-agent/global"
	"coding.jd.com/bmp/cps-agent/task"

	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/process"
)

var (
	helpFlag    bool
	versionFlag bool
	configFile  string
)

func init() {
	flag.BoolVar(&helpFlag, "h", false, "show help")
	flag.BoolVar(&versionFlag, "v", false, "show version and build")
	flag.BoolVar(&global.Debug, "d", true, "run in debug mode")
	flag.StringVar(&global.SerialNumber, "s", "", "SerialNumber")
	flag.StringVar(&configFile, "c", "global/config.go", "specify config file")

	flag.Usage = usage
}

func writeFile(line string) {

	logf, err := os.OpenFile("testddd"+"."+time.Now().Format("2006010215"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open filename error:", err.Error())
	}
	defer logf.Close()
	if _, err := io.WriteString(logf, line); err != nil {
		fmt.Println("write line  error:", err.Error())
	}
}

func main() {

	if runtime.GOOS == "windows" {
		pn, _ := process.NewProcess(int32(os.Getppid()))
		pName, _ := pn.Name()
		if pName != "srvany.exe" {
			fmt.Fprint(os.Stdout, "cps-agent.exe不支持独立启动!")
			time.Sleep(2 * time.Second)
			os.Exit(1)
		}
	}

	flag.Parse()

	if helpFlag {
		flag.Usage()
	}

	global.Init()
	if versionFlag {
		version()
	}

	global.InitApp(configFile)

	logs.Notice("========Agent Start========")

	task.Start()
	<-global.SignalChan
	logs.Notice("========Agent Exit========")

	task.Shutdown()
	if runtime.GOOS == "linux" {
		global.UninitApp()
	}

	time.Sleep(500 * time.Millisecond)
	os.Exit(0)
}

func usage() {
	fmt.Fprintf(os.Stdout, `jdcps-agent version: %s
Usage: jdcps-agent [-chvds]

Options:
`, global.VERSION)
	flag.PrintDefaults()
	os.Exit(0)
}

func version() {
	fmt.Fprintf(os.Stdout, "Version:%s.%s\nSerialNumber:%s\n", global.VERSION, global.BUILD_TIME, global.SerialNumber)
	os.Exit(0)
}
