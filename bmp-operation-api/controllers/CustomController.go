package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/CustomInfoLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	"git.jd.com/cps-golang/ironic-common/exception"
)

// CustomController operations for custom info
type CustomController struct {
	BaseController
}

// FilterList ...
// swagger:route GET /filterList custom filterList
// FilterList 设置自定义信息
//     Responses:
//       200: filterList
//       default: ErrorResponse
func (c *CustomController) FilterList() {
	defer c.CatchException()
	language := c.logPoints.GetPoint("language").(string)
	core := "物理核,"
	if language == BaseLogic.EN_US {
		core = "cores,"
	}
	architecture := response.ArchitectureList{
		Arm:       "ARM64(aarch64)",
		I386:      "i386",
		X86_64:    "x86_64",
		LoongArch: "LoongArch™",
	}

	bootMode := []string{
		"UEFI",
		"Legacy/BIOS",
	}

	res := response.FilterList{
		Architecture: architecture,
		BootMode:     bootMode,
	}

	cpu := response.CPU{}
	for k, v := range BaseLogic.Cpulist {
		cpu.Value = k + 1
		arr := strings.Split(v, ",")
		cpu.Label = arr[0] + " " + arr[1] + "(" + arr[2] + "*" + arr[3] + core + arr[4] + "GHz)"
		cpu.Info.CPUManufacturer = arr[0]
		cpu.Info.CPUModel = arr[1]
		cpu.Info.CPUAmount = s2i(arr[2])
		cpu.Info.CPUCores = s2i(arr[3])
		cpu.Info.CPUFrequency = arr[4]
		res.CPU = append(res.CPU, cpu)
	}

	mem := response.Mem{}
	for k, v := range BaseLogic.Memlist {
		mem.Value = k + 1
		arr := strings.Split(v, " ")
		mem.Label = arr[0] + "GB(" + arr[1] + "GB*" + arr[2] + ")" + arr[3] + " " + arr[4] + "MHz"
		mem.Info.MemType = arr[3]
		mem.Info.MemFrequency = s2i(arr[4])
		mem.Info.MemSize = s2i(arr[1])
		mem.Info.MemAmount = s2i(arr[2])
		res.Mem = append(res.Mem, mem)
	}

	res.RaidRules = CustomInfoLogic.GetRaidconf(c.logPoints)

	c.Res.Result = res
}
func s2i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// SetCustomInfo ...
// swagger:route POST /custom/setCustomInfo custom setCustomInfo
// SetCustomInfo 设置自定义信息
//     Responses:
//       200: setCustomInfo
//       default: ErrorResponse

func (c *CustomController) SetCustomInfo() {
	defer c.CatchException()
	req := &requestTypes.SetCustomInfoRequest{}
	// req.Validate()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		panic(exception.Exception{Msg: err.Error()})
	}
	res := CustomInfoLogic.SetCustomInfo(c.logPoints, *req)
	c.Res.Result = res
}

// GetCustomInfo ...
// swagger:route GET /custom/getCustomInfo custom getCustomInfo
// GetCustomInfo 获取自定义信息
//     Responses:
//       200: getCustomInfo
//       default: ErrorResponse

func (c *CustomController) GetCustomInfo() {
	defer c.CatchException()
	req := &requestTypes.QueryCustomInfoRequest{
		PageKey: c.GetString("pageKey"),
		Reload:  c.GetString("reload"),
	}
	// req.Validate()
	res := CustomInfoLogic.GetCustomInfo(c.logPoints, *req)
	c.Res.Result = res
}
