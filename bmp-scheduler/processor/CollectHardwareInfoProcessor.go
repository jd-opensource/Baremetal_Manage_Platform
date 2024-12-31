package processor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	deviceDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/interfaceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/jdBondInterfaceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/jdSwitchDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	util "coding.jd.com/aidc-bmp/bmp-scheduler/util"
	"git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/InterfaceType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/NetworkType"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
	beego "github.com/beego/beego/v2/server/web"
)

const (
	SWITCH_PORT_REGEX   = `[0-9]/([0-9]+)`
	INTERFACE_ETH0_NAME = "eth0"
	INTERFACE_ETH1_NAME = "eth1"
)

type CollectHardwareInfoProcessor struct {
	BaseProcessor
}

func NewCollectHardwareInfoProcessor() CollectHardwareInfoProcessor {
	b := CollectHardwareInfoProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b CollectHardwareInfoProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

type HardwareInfo struct {
	// megacli64
	RaidDriver string `json:"raid_driver"`
	// 10.208.16.165
	BmcAddress string     `json:"bmc_address"`
	Interfaces Interfaces `json:"interfaces"`
	// Disks        []Disk        `json:"disks"`
	Boot         Boot          `json:"boot"`
	SystemVendor *SystemVendor `json:"system_vendor"`
	Memory       Memory        `json:"memory"`
	CPU          CPU           `json:"cpu"`
	Platform     Platform      `json:"platform"`
}

const MAX_PRIVATE_INTERFACE_INDEX = 20

type Interface struct {
	Product        string `json:"product"`
	Vendor         string `json:"vendor"`
	Name           string `json:"name"`
	HasCarrier     bool   `json:"has_carrier"`
	Ipv4Address    string `json:"ipv4_address"`
	SwitchManageIP string `json:"switch_manage_ip"`
	Biosdevname    string `json:"biosdevname"`
	MacAddress     string `json:"mac_address"`
	SwitchIndex    string `json:"switch_index"`
}

func (i Interface) isPrivateInterface() bool {
	return getSwitchPortIndex(i.SwitchIndex) <= MAX_PRIVATE_INTERFACE_INDEX
}

type Interfaces []Interface

func (in Interfaces) Len() int { return len(in) }

func (in Interfaces) Swap(i, j int) { in[i], in[j] = in[j], in[i] }

func (in Interfaces) Less(i, j int) bool { return in[i].MacAddress < in[j].MacAddress }

type Disk struct {
	Rotational         bool   `json:"rotational"`
	Vendor             string `json:"vendor"`
	Name               string `json:"name"`
	WwnVendorExtension string `json:"wwn_vendor_extension"`
	Hctl               string `json:"hctl"`
	WwnWithExtension   string `json:"wwn_with_extension"`
	ByPath             string `json:"by_path"`
	Model              string `json:"model"`
	Wwn                string `json:"wwn"`
	Serial             string `json:"serial"`
	Size               int64  `json:"size"`
}

type Boot struct {
	CurrentBootMode string `json:"current_boot_mode"`
}

type SystemVendor struct {
	SerialNumber string `json:"serial_number"`
	ProductName  string `json:"product_name"`
	Manufacturer string `json:"manufacturer"`
}

type Memory struct {
	PhysicalMb int64 `json:"physical_mb"`
	Total      int64 `json:"total"`
}

type CPU struct {
	Count        int      `json:"count"`
	Frequency    string   `json:"frequency"`
	Flags        []string `json:"flags"`
	ModelName    string   `json:"model_name"`
	Architecture string   `json:"architecture"`
}

type Platform struct {
	//系统架构 x86_64, aarch64
	Machine string `json:"machine"`
}

func (b CollectHardwareInfoProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {

	logger.Info("CollectHardwareInfoProcessor Callback starting...", msg.Sn)
	defer logger.Info("CollectHardwareInfoProcessor Callback ending...", msg.Sn)
	if strings.ToUpper(msg.Status) == OK {

		//以后再优化成更优雅的方法
		b, _ := json.Marshal(msg.Data)
		saveHardwareInfo(logger, msg.Sn, b)

	}
	b.BaseProcessor.CallBack(logger, command, msg)
}

func saveHardwareInfo(logger *log.Logger, sn string, content []byte) error {

	originContent := string(content)

	logger.Infof("TmpHardwareInfo sn: %s, origin content:%s", sn, string(content))

	var err error
	c := HardwareInfo{}
	if err := json.Unmarshal(content, &c); err != nil {
		logger.Warnf("TmpHardwareInfo parse error, sn:%s, content:%s", sn, originContent)
		return err
	}

	saveHardware2Db(logger, sn, c)

	content, err = json.Marshal(c)
	if err != nil {
		logger.Warnf("TmpHardwareInfo parse2 error, sn:%s, content:%s", sn, originContent)
		return err
	}

	logPath, _ := beego.AppConfig.String("log.path")
	targetDir := filepath.Join(logPath, "hardware_info_dir")
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(targetDir, 0777) //0777也可以os.ModePerm
		os.Chmod(targetDir, 0777)
	}
	targetFile := filepath.Join(targetDir, sn)

	file, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		logger.Warn(targetFile, "文件打开失败：", err)
		return err
	}
	//及时关闭file句柄
	defer file.Close()
	write := bufio.NewWriter(file)
	if _, err := write.WriteString(string(content)); err != nil {
		logger.Warn(targetFile, "CollectHardwareInfoProcessor写入文件失败：", err)
		return err
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	return err
}

func saveHardware2Db(logger *log.Logger, sn string, hardware_info HardwareInfo) error {
	logger.Info("saveHardware2Db starting......", sn)
	defer logger.Info("saveHardware2Db ending......", sn)
	device, err := deviceDao.GetBySn(logger, sn)
	if device == nil {
		logger.Warn(fmt.Sprintf("device not found by sn %s, sql error: %s", sn, err.Error()))
		return err
	}

	//这里对deviceEntity的interface信息进行设置
	if err := parseInterfaces(logger, device, hardware_info.Interfaces); err != nil {
		return err
	}

	system_vendor := hardware_info.SystemVendor
	if system_vendor != nil {

		device.Brand = system_vendor.Manufacturer
		device.Model = system_vendor.ProductName
	}
	device.RaidDriver = hardware_info.RaidDriver
	device.CurrentBootMode = hardware_info.Boot.CurrentBootMode
	device.Architecture = hardware_info.Platform.Machine

	//device.ManufacturerID = manufacturer_entity.ID
	if err := deviceDao.UpdateDeviceAnywhere(logger, device); err != nil {
		logger.Warnf("saveHardware2Db UpdateDeviceAnywhere error, sn:%s, error:%s", device.Sn, err.Error())
		return err
	}

	return nil
}

func isRetail(logger *log.Logger, sn string) bool {
	logger.Info("isRetail ......", sn)

	var info_extra *types.CollectInfoExtra = loadCollectInfoExtra(logger, sn)
	if info_extra == nil {
		return false
	}
	return info_extra.NetworkType == NetworkType.RETAIL
}

func loadCollectInfoExtra(logger *log.Logger, sn string) *types.CollectInfoExtra {
	logger.Info("loadCollectInfoExtra ......", sn)
	extra_json, err := redis.GetObjectFromRedis(fmt.Sprintf(Constants.COLLECT_EXTRA_KEY, sn))
	logger.Info("extra_json is: ", sn, extra_json)
	if err != nil {
		logger.Warn("loadCollectInfoExtra GetObjectFromRedis error:", sn, err.Error())
	}
	if extra_json == "" {
		return nil
	}
	var c = &types.CollectInfoExtra{}
	if err := json.Unmarshal([]byte(extra_json), c); err != nil {
		logger.Warn("loadCollectInfoExtra Unmarshal error:", sn, err.Error())
		return nil
	}
	return c
}

func checkRetailInterfaces(logger *log.Logger, sn string, interfaces Interfaces) {

	logger.Info("checkRetailInterfaces ......", sn)
	//debug......
	// s := `[{"product": "0x10fb", "vendor": "0x8086", "name": "eth1", "has_carrier": true, "ipv4_address": "10.208.14.167", "switch_manage_ip": "10.208.13.11", "biosdevname": "em2", "mac_address": "24:6e:96:b7:b0:da", "switch_index": "Ten-GigabitEthernet2/0/3"}, {"product": "0x10fb", "vendor": "0x8086", "name": "eth0", "has_carrier": true, "ipv4_address": "10.208.14.184", "switch_manage_ip": "10.208.13.11", "biosdevname": "em1", "mac_address": "24:6e:96:b7:b0:d8", "switch_index": "Ten-GigabitEthernet1/0/3"}]`
	// if err := json.Unmarshal([]byte(s), &interfaces); err != nil {
	// 	fmt.Println("mock interface error:", err.Error())
	// }
	//ending debug......

	if len(interfaces) < 2 {
		panic(fmt.Sprintf("CollectHardwareInfo retail switch_index %s not match 2", sn))
	}
	sort.Sort(interfaces)
	for i := 0; i < len(interfaces)-1; i++ {
		first := interfaces[i]
		next := interfaces[i+1]
		if getSwitchPortIndex(first.SwitchIndex) > getSwitchPortIndex(next.SwitchIndex) {
			panic(fmt.Sprintf("CollectHardwareInfo mac&port sort error %s", sn))
		}
	}
}

func getSwitchPortIndex(switch_index string) int {
	fmt.Println(time.Now(), "getSwitchPortIndex ......", switch_index)
	r, _ := regexp.Compile(SWITCH_PORT_REGEX)
	matcher := r.FindAllString(switch_index, -1)
	if len(matcher) == 0 {
		panic(fmt.Sprintf("switch_index %s not match", switch_index))
	}
	n, err := strconv.Atoi(strings.Split(matcher[0], "/")[0])
	if err != nil {
		panic(fmt.Sprintf("switch_index %s not match int", switch_index))
	}
	return n
}

func parseInterfaces(logger *log.Logger, device *deviceDao.Device, interfaces Interfaces) error {
	logger.Info("parseAndSaveInterfaces ......", device.Sn)
	//debug......
	// s := `[{"product": "0x10fb", "vendor": "0x8086", "name": "eth1", "has_carrier": true, "ipv4_address": "10.208.14.167", "switch_manage_ip": "10.208.13.11", "biosdevname": "em2", "mac_address": "24:6e:96:b7:b0:da", "switch_index": "Ten-GigabitEthernet2/0/3"}, {"product": "0x10fb", "vendor": "0x8086", "name": "eth0", "has_carrier": true, "ipv4_address": "10.208.14.184", "switch_manage_ip": "10.208.13.11", "biosdevname": "em1", "mac_address": "24:6e:96:b7:b0:d8", "switch_index": "Ten-GigabitEthernet1/0/3"}]`
	// if err := json.Unmarshal([]byte(s), &interfaces); err != nil {
	// fmt.Println("mock interface error:", err.Error())
	// }
	// if err := deleteInterfaces(logger, sn); err != nil {
	// 	return err
	// }
	var lan_interface_count int = 0
	var wan_interface_count int = 0
	entities := []*interfaceDao.Interface{}
	for _, interface_info := range interfaces {
		if interface_info.SwitchIndex == "" {
			continue
		}
		entity := &interfaceDao.Interface{}
		if interface_info.isPrivateInterface() {
			entity.InterfaceType = InterfaceType.LAN
			if lan_interface_count == 0 {
				entity.InterfaceName = INTERFACE_ETH0_NAME
			} else {
				entity.InterfaceName = INTERFACE_ETH1_NAME
			}
			lan_interface_count += 1
		} else {
			entity.InterfaceType = InterfaceType.WAN
			if wan_interface_count == 0 {
				entity.InterfaceName = INTERFACE_ETH1_NAME
			} else {
				entity.InterfaceName = INTERFACE_ETH0_NAME
			}
			wan_interface_count += 1
		}
		entity.Ipv4Address = interface_info.Ipv4Address
		entity.Sn = device.Sn
		entity.Mac = interface_info.MacAddress
		entity.SwitchIP = interface_info.SwitchManageIP
		entity.SwitchPort = interface_info.SwitchIndex
		entities = append(entities, entity)
	}

	t := (TempDaoInterfaces)(entities)
	sort.Sort(t)
	// if _, err := interfaceDao.AddMultiInterface(logger, ([]*interfaceDao.Interface)(t)); err != nil {
	// 	logger.Warn("parseAndSaveInterfaces AddMultiInterface sql error:", err.Error())
	// 	return err
	// }
	if len(t) == 1 { //只插一根网线
		device.Mac1 = t[0].Mac
		device.SwitchIP = t[0].SwitchIP
		device.SwitchIP1 = t[0].SwitchIP
		device.SwitchPort1 = t[0].SwitchPort
		device.Mac2 = ""
		device.SwitchIP2 = ""
		device.SwitchPort2 = ""
	} else if len(t) == 2 { //插两根网线
		device.Mac1 = t[0].Mac
		device.SwitchIP = t[0].SwitchIP
		device.SwitchIP1 = t[0].SwitchIP
		device.SwitchPort1 = t[0].SwitchPort
		device.Mac2 = t[1].Mac
		device.SwitchIP2 = t[1].SwitchIP
		device.SwitchPort2 = t[1].SwitchPort
		device.PrivateEth1IPv4 = t[1].Ipv4Address
	}
	return nil

}

type TempDaoInterfaces []*interfaceDao.Interface

func (t TempDaoInterfaces) Len() int { return len(t) }

func (t TempDaoInterfaces) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func (t TempDaoInterfaces) Less(i, j int) bool { return t[i].InterfaceName < t[j].InterfaceName }

func deleteInterfaces(logger *log.Logger, sn string) error {
	logger.Info("deleteInterfaces ......", sn)
	query := map[string]interface{}{
		"sn":     sn,
		"is_del": 0,
	}
	updates := map[string]interface{}{
		"is_del":      1,
		"update_time": time.Now(),
	}
	if err := interfaceDao.UpdateMultiInterfaces(logger, query, updates); err != nil {
		logger.Warn("deleteInterfaces UpdateMultiInterfaces sql sn error:", sn, err.Error())
		return err
	}
	logger.Info("deleteInterfaces UpdateMultiInterfaces sql success, sn:", sn)
	return nil

}

func parseAndSaveJDBondInterface(logger *log.Logger, sn string, interfaces Interfaces) error {

	logger.Info("parseAndSaveJDBondInterface ......", sn)
	bond_entity, err := jdBondInterfaceDao.GetJDBondInterfaceBysn(logger, sn)
	if err != nil {
		logger.Warn("parseAndSaveJDBondInterface GetJDBondInterfaceBysn sql error:", err.Error())
	}
	if bond_entity != nil {
		bond_entity.IsDel = 1
		bond_entity.UpdateTime = time.Now()
		if err := jdBondInterfaceDao.UpdateJdBondInterfaceById(logger, bond_entity); err != nil {
			logger.Warn("parseAndSaveJDBondInterface UpdateJdBondInterfaceById sql error:", err.Error())
		}
	}
	switch_ips := map[string]struct{}{}
	for _, interface_info := range interfaces {
		if interface_info.SwitchIndex == "" {
			continue
		}
		switch_ips[interface_info.SwitchManageIP] = struct{}{}
	}
	if len(switch_ips) == 0 {
		return fmt.Errorf("get empty switch_ip from interfaces")
	}
	for switch_ip, _ := range switch_ips {
		//java中对switch_ip加锁，不明白加锁的意义在哪里，先不加
		jd_bond_entity := &jdBondInterfaceDao.JdBondInterface{}
		jd_bond_entity.SwitchIP = switch_ip
		jd_bond_entity.Sn = sn
		entities, err := interfaceDao.GetBySn(logger, sn)
		if err != nil {
			logger.Warn("parseAndSaveJDBondInterface interfaceDao.GetBySn sql error:", err.Error())
			continue
		}
		jd_bond_entity.GroupID = getSwitchPortIndex(entities[0].SwitchPort)
		jd_bond_entity.PrivateIP = getPrivateIp(logger, sn, switch_ip)
		if _, err := jdBondInterfaceDao.AddJdBondInterface(logger, jd_bond_entity); err != nil {
			logger.Warn("jdBondInterfaceDao.AddJdBondInterface sql error:", err.Error())
			return err
		}
	}
	return nil

}

func getPrivateIp(logger *log.Logger, sn, switch_ip string) string {
	query := map[string]interface{}{
		"is_del":    0,
		"switch_ip": switch_ip,
	}
	jd_switch, err := jdSwitchDao.GetOneJdSwitch(logger, query)
	if err != nil {
		logger.Warn("getPrivateIp GetOneJdSwitch sql error:", err.Error())
	}
	if jd_switch == nil || jd_switch.Cidr == "" {
		return ""
	}
	jd_bond_interfaces, err := jdBondInterfaceDao.GetAllJdBondInterface(logger, query)
	if err != nil {
		logger.Warn("getPrivateIp GetAllJdBondInterface sql error:", err.Error())
	}
	used_ips := []string{}
	for _, entity := range jd_bond_interfaces {
		if entity.PrivateIP != "" {
			used_ips = append(used_ips, entity.PrivateIP)
		}
	}
	ip_networks := util.GetAllHostIPs(jd_switch.Cidr)
	var tar_ip string
	for _, ip := range ip_networks {
		if util.InArray(ip, used_ips) {
			continue
		} else {
			tar_ip = ip
			break
		}
	}
	if tar_ip == "" {
		logger.Warn(fmt.Sprintf("switch_ip cidr is use up %s sn %s", switch_ip, sn))
		panic(fmt.Sprintf("CollectHardwareInfo retail %s %s cidr is use up", switch_ip, sn))
	}
	return tar_ip
}

func parseAndSaveJDSwitch(logger *log.Logger, interfaces Interfaces) error {

	logger.Info("parseAndSaveJDSwitch ......")
	switch_ips := map[string]struct{}{}
	for _, interface_info := range interfaces {
		if interface_info.SwitchIndex == "" {
			continue
		}
		switch_ips[interface_info.SwitchManageIP] = struct{}{}
	}
	if len(switch_ips) == 0 {
		return fmt.Errorf("get empty switch_ip from interfaces")
	}

	for switch_ip := range switch_ips {
		query := map[string]interface{}{
			"is_del":    0,
			"switch_ip": switch_ip,
		}
		switch_entity, err := jdSwitchDao.GetOneJdSwitch(logger, query)
		if err != nil {
			logger.Warn("parseAndSaveJDSwitch GetOneJdSwitch sql error:", err.Error())
		}
		if switch_entity != nil {
			continue
		}
		jd_switch_entity := &jdSwitchDao.JdSwitch{}
		jd_switch_entity.SwitchIP = switch_ip
		if _, err := jdSwitchDao.AddJdSwitch(logger, jd_switch_entity); err != nil {
			logger.Warn("parseAndSaveJDSwitch AddJdSwitch sql error:", err.Error())
			return err
		}
	}
	return nil
}

func (b CollectHardwareInfoProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	logger.Info("CollectHardwareInfoProcessor doProcess starting...", command.Sn)
	defer logger.Info("CollectHardwareInfoProcessor doProcess ending...", command.Sn)
	collect_hardware_info := agentTypes.CollectHardwareInfo{
		Sn:     command.Sn,
		Action: command.Action,
	}
	if err := rabbitMq.SendToAgent(command.Sn, collect_hardware_info); err != nil {
		logger.Warnf("CollectHardwareInfoProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(collect_hardware_info), err.Error())
		return
	}
	logger.Infof("CollectHardwareInfoProcessor_SendToAgent_Msg:%s", util.ObjToJson(collect_hardware_info))
}

func (b CollectHardwareInfoProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	logger.Info("CollectHardwareInfoProcessor afterProcess starting...", command.Sn)
	defer logger.Info("CollectHardwareInfoProcessor afterProcess ending...", command.Sn)
	//empty
}
