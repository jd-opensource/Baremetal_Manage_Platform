package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"
)

//安装目录固定为 C:\Program Files (x86)\bmp-agent 不可选
const CPS_AGENT_DIR = `C:\Program Files (x86)\bmp-agent`

const (
	RUNNING       = "RUNNING"
	STOPPING      = "STOPPING"
	STOPPED       = "STOPPED"
	START_PENDING = "START_PENDING"
)

/* sc query 输出：
SERVICE_NAME: bmp-agent
TYPE               : 10  WIN32_OWN_PROCESS
STATE              : 1  STOPPED
WIN32_EXIT_CODE    : 0  (0x0)
SERVICE_EXIT_CODE  : 0  (0x0)
CHECKPOINT         : 0x0
WAIT_HINT          : 0x0
*/

type ServiceScInfo struct {
	ServiceName     string
	Type            string
	State           string
	Win32ExitCode   string
	ServiceExitCode string
}

func main() {
	u, _ := user.Current()
	if !strings.HasSuffix(u.Username, "Administrator") {
		fmt.Fprintf(os.Stdout, "请切换至Administrator用户安装bmp-agent服务")
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "bmp-agent安装中...\n")

	v, err := getCpsAgentInfo()
	if err != nil && strings.Contains(err.Error(), "status 1060") {
		//未安装
		if err := installCpsAgent(); err != nil {
			fmt.Fprintf(os.Stdout, "bmpAgent安装失败:"+err.Error())
			os.Exit(1)
		}
		if err := startCpsAgent(); err != nil {
			fmt.Fprintf(os.Stdout, "bmpAgent启动失败:"+err.Error())
			os.Exit(1)
		}
	} else {
		//已安装
		switch v.State {
		case RUNNING:
			if err := stopCpsAgent(); err != nil {
				fmt.Fprintf(os.Stdout, "stopBmpAgent error:"+err.Error())
				os.Exit(1)
			}
			if err := uninstallCpsAgent(); err != nil {
				fmt.Fprintf(os.Stdout, "uninstallBmpAgent error:"+err.Error())
				os.Exit(1)
			}
			if err := installCpsAgent(); err != nil {
				fmt.Fprintf(os.Stdout, "installBmpAgent error:"+err.Error())
				os.Exit(1)
			}
			if err := startCpsAgent(); err != nil {
				fmt.Fprintf(os.Stdout, "startBmpAgent error:"+err.Error())
				os.Exit(1)
			}
		case STOPPED:
			if err := uninstallCpsAgent(); err != nil {
				fmt.Fprintf(os.Stdout, "uninstallBmpAgent error:"+err.Error())
				os.Exit(1)
			}
			if err := installCpsAgent(); err != nil {
				fmt.Fprintf(os.Stdout, "installBmpAgent error:"+err.Error())
				os.Exit(1)
			}
			if err := startCpsAgent(); err != nil {
				fmt.Fprintf(os.Stdout, "startBmpAgent error:"+err.Error())
				os.Exit(1)
			}
		case START_PENDING:
			fmt.Fprintf(os.Stdout, "bmp-agent server is starting, please wait......")
			time.Sleep(60 * time.Second)
		default:
			stopCpsAgent()
			uninstallCpsAgent()
			installCpsAgent()
			startCpsAgent()
		}

	}

	v, err = getCpsAgentInfo()
	if err != nil {
		fmt.Fprintf(os.Stdout, "bmp-agent server unknown error")
		os.Exit(1)
	}

	var finalState = v.State
	if v.State == START_PENDING { //5s检测一次，检测100s
		for i := 0; i < 20; i++ {
			time.Sleep(5 * time.Second)
			if v, _ = getCpsAgentInfo(); v.State == RUNNING {
				finalState = RUNNING
				break
			}
		}
	}

	if finalState != RUNNING {
		fmt.Fprintf(os.Stdout, "bmp-agent安装失败！")
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "bmp-agent安装并启动成功\n")
	time.Sleep(2 * time.Second)
	os.Exit(0)

}

func getCpsAgentInfo() (*ServiceScInfo, error) {
	output, err := execute("cmd", "/C", "sc", "query", "bmp-agent")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	obj := &ServiceScInfo{}
	for _, line := range lines {
		tmp := strings.Split(line, ":")
		if strings.Contains(tmp[0], "SERVICE_NAME") {
			obj.ServiceName = strings.TrimSpace(tmp[1])
		}
		if strings.Contains(tmp[0], "STATE") {
			statuses := strings.Split(strings.TrimSpace(tmp[1]), " ")
			obj.State = statuses[len(statuses)-1]
		}
		if strings.Contains(tmp[0], "ServiceExitCode") {
			obj.ServiceExitCode = strings.TrimSpace(tmp[1])
		}
	}
	return obj, nil

}

func uninstallCpsAgent() error {
	_, err := execute("cmd", "/C", "sc", "delete", "bmp-agent")
	if err != nil {
		return err
	}
	return nil
}

//
func installCpsAgent() error {
	//sc create bmp-agent binPath= "C:\Program Files (x86)\Windows Resource Kits\Tools\srvany.exe" start= auto
	//reg add HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\bmp-agent\Parameters /v Application /t REG_SZ /d "C:\minping\bmp-agent.exe" /f
	execute("cmd", "/C", "mkdir", CPS_AGENT_DIR)
	srvanyFile := fmt.Sprintf(`%s\srvany.exe`, CPS_AGENT_DIR)
	cpsAgentFile := fmt.Sprintf(`%s\bmp-agent.exe`, CPS_AGENT_DIR)
	execute("cmd", "/C", "del", srvanyFile)
	execute("cmd", "/C", "del", cpsAgentFile)
	if _, err := execute("cmd", "/C", "copy", "internal\\srvany.exe", srvanyFile); err != nil {
		return err
	}
	if _, err := execute("cmd", "/C", "copy", "internal\\bmp-agent.exe", cpsAgentFile); err != nil {
		return err
	}
	if _, err := execute("cmd", "/C", "sc", "create", "bmp-agent", "binPath=", fmt.Sprintf(`"%s"`, srvanyFile), "start=", "auto"); err != nil {
		return err
	}
	if _, err := execute("cmd", "/C", "reg", "add", `HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\bmp-agent\Parameters`, `/v`, `Application`, `/t`, `REG_SZ`, `/d`, cpsAgentFile, `/f`); err != nil {
		return err
	}
	return nil
}

func startCpsAgent() error {
	_, err := execute("cmd", "/C", "sc", "start", "bmp-agent")
	if err != nil {
		return err
	}
	return nil
}

func stopCpsAgent() error {
	_, err := execute("cmd", "/C", "sc", "stop", "bmp-agent")
	if err != nil {
		return err
	}
	return nil
}

func execute(name string, arg ...string) (string, error) {
	command := exec.Command(name, arg...)
	output, err := command.CombinedOutput()
	return string(output), err
}
