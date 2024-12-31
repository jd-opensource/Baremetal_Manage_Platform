package processor

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/diskDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	util "coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
	beego "github.com/beego/beego/v2/server/web"
)

type D struct {
	Enclosure int     `json:"enclosure"`
	Slot      int     `json:"slot"`
	Size      float64 `json:"size"`
	//GB
	SizeUnit string `json:"size_unit"`
	//接口类型 SAS
	PdType string `json:"pd_type"`
	//磁盘类型 新增， SDD
	MediaType string `json:"media_type"`
}

type Nvme struct {
	//nvme无须做raid enclosure，slot都设置为0即可
	Devices []NvmeDevice `json:"devices"`
}
type NvmeDevice struct {
	// /dev/nvme0n1
	DevicePath string  `json:"device_path"`
	Index      int     `json:"index"`
	Serial     string  `json:"serial"`
	Size       float64 `json:"size"`
	//GB
	SizeUnit string `json:"size_unit"`
	//接口类型 NVME
	PdType string `json:"pd_type"`
	//磁盘类型 新增， SDD
	MediaType string `json:"media_type"`
}

type Cit struct {
	Disks     []D `json:"disks"`
	AdapterID int `json:"adapter_id"` //可能是int，也可能是string
}

type TmpDiskInfo struct {
	// 有的机器可能有多张raid卡的情况，每个cit都代表一个raid卡。本次只支持有一张raid卡
	Controllers []Cit `json:"controllers"`
	Nvme        Nvme  `json:"nvme"`
	//盘符
	Devices []DeviceItem `json:"devices"`
}

type DeviceItem struct {
	// /dev/nvme0n1
	DevicePath string `json:"device_path"`
	//磁盘类型 新增， SDD
	MediaType string  `json:"media_type"`
	Serial    string  `json:"serial"`
	Size      float64 `json:"size"`
	//GB
	SizeUnit string `json:"size_unit"`
}

type CollectDiskLocationsProcessor struct {
	BaseProcessor
}

func NewCollectDiskLocationsProcessor() CollectDiskLocationsProcessor {
	b := CollectDiskLocationsProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b CollectDiskLocationsProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b CollectDiskLocationsProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {

	logger.Info("CollectDiskLocationsProcessor Callback starting...", msg.Sn)
	defer logger.Info("CollectDiskLocationsProcessor Callback ending...", msg.Sn)
	if strings.ToUpper(msg.Status) == OK {
		logger.Infof("CollectDiskLocationsProcessor callback content is:%s", util.ObjToJson(msg.Data))
		//以后再优化成更优雅的方法
		b, err := json.Marshal(msg.Data)
		if err != nil {
			logger.Warn("CollectDiskLocationsProcessor Callback unmarshal error:", msg.Sn, err.Error())
		} else {
			SaveDiskLocationInfo(logger, msg.Sn, b)
		}
	}
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b CollectDiskLocationsProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	logger.Info("CollectDiskLocationsProcessor doProcess starting...", command.Sn)
	defer logger.Info("CollectDiskLocationsProcessor doProcess ending...", command.Sn)

	device, err := deviceDao.GetBySn(logger, command.Sn)
	if err != nil {
		logger.Warn("CollectDiskLocationsProcessor GetById sql error:", command.Sn, err.Error())
	}
	if device == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process CollectDiskLocationsProcessor error,sn %s device not found", command.Sn)})
	}

	collect_disk_location := agentTypes.CollectDiskLocations{
		Sn:      command.Sn,
		Action:  command.Action,
		Version: "2.0",
		// AdapterIds:    []string{strconv.Itoa(device.AdapterID)},
		AllowOverride: false,
		RaidDriver:    device.RaidDriver,
	}
	// if device.AllowOverride != 0 {
	// 	collect_disk_location.AllowOverride = true
	// }
	if err := rabbitMq.SendToAgent(command.Sn, collect_disk_location); err != nil {
		logger.Warnf("CollectDiskLocationsProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(collect_disk_location), err.Error())
		return
	}
	logger.Infof("CollectDiskLocationsProcessor_SendToAgent_Msg:%s", util.ObjToJson(collect_disk_location))
}

func (b CollectDiskLocationsProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	logger.Info("CollectDiskLocationsProcessor afterProcess starting...", command.Sn)
	defer logger.Info("CollectDiskLocationsProcessor afterProcess ending...", command.Sn)
	//empty
}

func saveDiskToDb(logger *log.Logger, sn string, c TmpDiskInfo) error {

	deviceEntity, err := deviceDao.GetBySn(logger, sn)
	if err != nil {
		logger.Warnf("saveDiskToDb.GetBySn error, sn:%s, err:%s", sn, err.Error())
		return err
	}
	if deviceEntity == nil {
		logger.Warnf("saveDiskToDb deviceEntity get nil, sn:%s", sn)
		return errors.New("invalid sn, check it!!!")
	}

	q1 := map[string]interface{}{
		"is_del":    0,
		"device_id": deviceEntity.DeviceID,
	}
	u1 := map[string]interface{}{
		"is_del": 1,
	}
	if err := diskDao.UpdateMultiDisks(logger, q1, u1); err != nil {
		logger.Warnf("saveDiskToDb.delDisks error, device_id:%s, error:%s", deviceEntity.DeviceID, err.Error())
	}

	//普通盘
	var disks []D
	//nvme盘
	var nvmes []NvmeDevice
	//盘符
	var panfus []DeviceItem
	if len(c.Controllers) > 0 {
		deviceEntity.AdapterID = c.Controllers[0].AdapterID
		disks = c.Controllers[0].Disks
	}
	if err := deviceDao.UpdateDeviceAnywhere(logger, deviceEntity); err != nil {
		logger.Warnf("saveDiskToDb.UpdateDevice.HasCollected and AdapterID error, sn:%s, error:%s", deviceEntity.Sn, err.Error())
	}

	nvmes = c.Nvme.Devices
	panfus = c.Devices

	diskEntities := []*diskDao.Disk{}
	for _, disk := range disks {
		d := &diskDao.Disk{
			DiskId:      util.GetUuid("disk-"),
			Name:        fmt.Sprintf("磁盘%d", disk.Slot),
			DeviceID:    deviceEntity.DeviceID,
			Enclosure:   fmt.Sprintf("%d", disk.Enclosure),
			Slot:        disk.Slot,
			Size:        strconv.FormatFloat(disk.Size, 'f', -1, 64),
			SizeUnit:    disk.SizeUnit,
			PdType:      disk.PdType,
			MediaType:   disk.MediaType,
			Types:       "controller",
			CreatedTime: int(time.Now().Unix()),
			UpdatedTime: int(time.Now().Unix()),
		}
		//agent采集过来的是二进制，这里转换成10进制
		if d.SizeUnit == "GB" {
			disk.Size = 1.073741824 * disk.Size
			d.Size = strconv.Itoa(int(disk.Size))
		} else if d.SizeUnit == "TB" {
			disk.Size = 1.1 * disk.Size
			d.Size = strconv.FormatFloat(disk.Size, 'f', 2, 64)
		}

		diskEntities = append(diskEntities, d)
	}
	for _, nvme := range nvmes {
		p := strings.Split(nvme.DevicePath, "/")
		d := &diskDao.Disk{
			DiskId:      util.GetUuid("disk-"),
			Name:        p[len(p)-1],
			DeviceID:    deviceEntity.DeviceID,
			DevicePath:  nvme.DevicePath,
			Serial:      nvme.Serial,
			Size:        strconv.FormatFloat(nvme.Size, 'f', -1, 64),
			SizeUnit:    nvme.SizeUnit,
			PdType:      nvme.PdType,
			MediaType:   nvme.MediaType,
			Types:       "nvme",
			CreatedTime: int(time.Now().Unix()),
			UpdatedTime: int(time.Now().Unix()),
		}
		//agent采集过来的是二进制，这里转换成10进制
		if d.SizeUnit == "GB" {
			nvme.Size = 1.073741824 * nvme.Size
			d.Size = strconv.Itoa(int(nvme.Size))
		} else if d.SizeUnit == "TB" {
			nvme.Size = 1.1 * nvme.Size
			d.Size = strconv.FormatFloat(nvme.Size, 'f', 2, 64)
		}

		diskEntities = append(diskEntities, d)
	}
	for _, panfu := range panfus {
		p := strings.Split(panfu.DevicePath, "/")

		d := &diskDao.Disk{
			DiskId:      util.GetUuid("disk-"),
			Name:        p[len(p)-1],
			DeviceID:    deviceEntity.DeviceID,
			DevicePath:  panfu.DevicePath,
			MediaType:   panfu.MediaType,
			Serial:      panfu.Serial,
			Size:        strconv.FormatFloat(panfu.Size, 'f', -1, 64),
			SizeUnit:    panfu.SizeUnit,
			Types:       "panfu",
			CreatedTime: int(time.Now().Unix()),
			UpdatedTime: int(time.Now().Unix()),
		}
		//agent采集过来的是二进制，这里转换成10进制
		if d.SizeUnit == "GB" {
			panfu.Size = 1.073741824 * panfu.Size
			d.Size = strconv.Itoa(int(panfu.Size))
		} else if d.SizeUnit == "TB" {
			panfu.Size = 1.1 * panfu.Size
			d.Size = strconv.FormatFloat(panfu.Size, 'f', 2, 64)
		}

		diskEntities = append(diskEntities, d)
	}

	if _, err := diskDao.AddMultiDisk(logger, diskEntities); err != nil {
		logger.Warnf("saveDiskToDb.AddMultiDisk error:%s", err.Error())
		return err
	}

	return nil
}

func SaveDiskLocationInfo(logger *log.Logger, sn string, content []byte) error {

	originContent := string(content)
	logger.Infof("SaveDiskLocationInfo sn: %s, origin content:%s", sn, originContent)
	//controllers.disks[0]的struct

	var err error
	c := TmpDiskInfo{}
	if err := json.Unmarshal(content, &c); err != nil {
		logger.Warnf("SaveDiskLocationInfo parse error, sn:%s, content:%s, error:%s", sn, originContent, err.Error())
		return err
	}

	saveDiskToDb(logger, sn, c)

	content, err = json.Marshal(c)
	if err != nil {
		logger.Warnf("SaveDiskLocationInfo parse2 error, sn:%s, content:%s, error:%s", sn, originContent, err.Error())
		return err
	}

	logPath, _ := beego.AppConfig.String("log.path")
	targetDir := filepath.Join(logPath, "disk_location_dir")
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
		logger.Warn(targetFile, "CollectDiskLocationsProcessor写入文件失败：", err)
		return err
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	return err
}
