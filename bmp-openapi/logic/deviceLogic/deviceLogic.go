package deviceLogic

import (
	"fmt"
	"math"
	"mime/multipart"
	"net"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/CpsLogContentCollectionDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/CpsLogItemsDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceHintsDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/diskDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/imageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rDeviceVolumeDisksDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rVolumeRaidDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/raidDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/volumeDao"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceTypeDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/idcDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/auditLogLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

	// "coding.jd.com/aidc-bmp/bmp-openapi/logic/licenseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/raidLogic"
	namespacePrefix "git.jd.com/cps-golang/ironic-common/ironic/common/NamespacePrefixs"

	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/jdBondInterfaceDao"

	deviceDao "coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/interfaceDao"
	idcApi "coding.jd.com/aidc-bmp/bmp-openapi/service/idc_api"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-openapi/service/rabbit_mq"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/AuditLogsType"
	instanceStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/InstanceStatus"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/SaleStatus"
	saleStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/SaleStatus"
	rabbitEvent "git.jd.com/cps-golang/ironic-common/ironic/event"
	rabbitIronicMsgApi "git.jd.com/cps-golang/ironic-common/ironic/event/api"
	"github.com/tealeg/xlsx"
)

type CreateDeviceMq struct {
	Action     string `json:"action"`
	Sn         string `json:"sn"`
	Subnet     string `json:"subnet"`
	SubnetMask string `json:"subnet_mask"`
	Routes     string `json:"routes"`
}

// 检测设备状态，给driver发消息 ，验证带外等信息是否正确
type CheckDeviceMq struct {
	Action    string `json:"action"`
	Sn        string `json:"sn"`
	IloIp     string `json:"ilo_ip"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Privilege string `json:"privilege"`
}

var wg sync.WaitGroup

func buildDeviceEntity(logger *log.Logger, idcId string, c *requestTypes.CreateDeviceSpec) *deviceDao.Device {
	ipv6, _ := beego.AppConfig.String("gateway6")
	return &deviceDao.Device{
		DeviceID:   commonUtil.GetRandString("d-", namespacePrefix.IMAGE_ID_LENGTH, false, true, true),
		InstanceID: "",
		IDcID:      idcId,
		// DeviceTypeID: deviceType.DeviceTypeID,
		ManageStatus: baseLogic.IN, //导入设备以后，默认已入库

		// Brand: c.Brand,
		// Model: c.Model,
		// DeviceSeries: deviceType.DeviceSeries,
		Cabinet:     c.Cabinet,
		UPosition:   c.UPosition,
		Sn:          c.Sn,
		IloUser:     c.IloUser,
		IloPassword: c.IloPassword,
		IloIP:       c.IloIP,
		Description: c.Description,

		Mac1: c.Mac1,
		// Mac2: c.Mac2,
		// SwitchIP1:   c.SwitchIP1,
		// SwitchPort1: c.SwitchPort1,
		// SwitchIP2:   c.SwitchIP2,
		// SwitchPort2: c.SwitchPort2,

		// SwitchUser1:     c.SwitchUser1,
		// SwitchPassword1: c.SwitchPassword1,
		// SwitchUser2:     c.SwitchUser2,
		// SwitchPassword2: c.SwitchPassword2,
		// SwitchIP:        c.SwitchIP,

		Mask: c.Mask,
		// MaskEth1:        c.MaskEth1,
		Gateway:     c.Gateway,
		PrivateIPv4: c.PrivateIPv4,
		// PrivateEth1IPv4: c.PrivateEth1IPv4,
		PrivateIPv6: c.PrivateIPv6,
		// PrivateEth1IPv6: c.PrivateEth1IPv6,
		Gateway6: ipv6,
		// AdapterID:       *c.AdapterID,

		// RaidDriver:  c.RaidDriver,
		CreatedBy:   logger.GetPoint("username").(string),
		UpdatedBy:   logger.GetPoint("username").(string),
		CreatedTime: int(time.Now().Unix()),
		UpdatedTime: int(time.Now().Unix()),
	}
}

func DeviceEntity2Device(logger *log.Logger, d *deviceDao.Device, idcMap map[string]idcDao.Idc, deviceTypeMap map[string]deviceTypeDao.DeviceType, deviceTypeMapInfo map[string]map[string]string, instanceMap map[string]instanceDao.Instance, imageMap map[string]imageDao.Image) *responseTypes.Device {

	tz := logger.GetPoint("timezone").(string)
	//idc, err := idcDao.GetIdcById(logger, d.IDcID) //获取机房信息
	//if err != nil {
	//	idc = &idcDao.Idc{}
	//}
	idc := idcMap[d.IDcID]
	//entity, err := deviceTypeDao.GetOneDeviceType(logger, map[string]interface{}{"device_type_id": d.DeviceTypeID})
	//if err != nil {
	//	entity = &deviceTypeDao.DeviceType{}
	//}
	entity := deviceTypeMap[d.DeviceTypeID]
	//instance, err := instanceDao.GetInstanceById(logger, d.InstanceID)
	//if instance == nil {
	//	instance = &instanceDao.Instance{}
	//}
	instance := instanceMap[d.InstanceID]
	//image, err := imageDao.GetImageByUuid(logger, instance.ImageID)
	//if image == nil {
	//	image = &imageDao.Image{}
	//}
	image := imageMap[instance.ImageID]
	//user, _ := userDao.GetUserById(logger, instance.UserID)
	//if user == nil {
	//	user = &userDao.User{}
	//}
	language := logger.GetPoint("language").(string)
	reason := baseLogic.DeviceReason[d.Reason]
	if language == baseLogic.EN_US {
		reason = baseLogic.DeviceReasonEn[d.Reason]
	}
	instanceStatusName := baseLogic.InstanceStatus[instance.Status]
	if language == baseLogic.EN_US {
		instanceStatusName = baseLogic.InstanceStatusEn[instance.Status]
	}
	instanceReason := baseLogic.InstanceReason[instance.Reason]
	if language == baseLogic.EN_US {
		instanceReason = baseLogic.InstanceReasonEn[instance.Reason]
	}
	deviceManageStatusName := baseLogic.DeviceManageStatus[d.ManageStatus]
	if language == baseLogic.EN_US {
		deviceManageStatusName = baseLogic.DeviceManageStatusEn[d.ManageStatus]
	}
	deviceSeriesName := baseLogic.DeviceTypeSeries[d.DeviceSeries]
	if language == baseLogic.EN_US {
		deviceSeriesName = baseLogic.DeviceTypeSeriesEn[d.DeviceSeries]
	}
	idcName := idc.Name
	if language == baseLogic.EN_US {
		idcName = idc.NameEn
	}
	var (
		iloUser         string
		iloPassword     string
		switchUser1     string
		switchUser2     string
		switchPassword1 string
		switchPassword2 string
		encloure1       string
		slot1           int
		encloure2       string
		slot2           int
	)
	// disk, _ := diskDao.GetAllDisk(logger, map[string]interface{}{"device_id": d.DeviceID})
	// if len(disk) == 1 {
	// 	encloure1 = disk[0].Enclosure
	// 	slot1 = disk[0].Slot
	// } else if len(disk) == 2 {
	// 	encloure1 = disk[0].Enclosure
	// 	slot1 = disk[0].Slot
	// 	encloure2 = disk[1].Enclosure
	// 	slot2 = disk[1].Slot
	// }

	if d.IloUser == "" {
		iloUser = idc.IloUser
	} else {
		iloUser = d.IloUser
	}
	if d.IloPassword == "" {
		iloPassword = idc.IloPassword
	} else {
		iloPassword = d.IloPassword
	}
	if d.SwitchUser1 == "" {
		switchUser1 = idc.SwitchUser1
	} else {
		switchUser1 = d.SwitchUser1
	}
	if d.SwitchPassword1 == "" {
		switchPassword1 = idc.SwitchPassword1
	} else {
		switchPassword1 = d.SwitchPassword1
	}
	if d.SwitchUser2 == "" {
		switchUser2 = idc.SwitchUser2
	} else {
		switchUser2 = d.SwitchUser2
	}
	if d.SwitchPassword2 == "" {
		switchPassword2 = idc.SwitchPassword2
	} else {
		switchPassword2 = d.SwitchPassword2
	}
	collectStatus := "2"
	if d.CollectStatus != "" {
		collectStatus = d.CollectStatus
	}
	return &responseTypes.Device{
		ID:                 d.ID,
		DeviceID:           d.DeviceID,
		IdcID:              d.IDcID,
		InstanceID:         d.InstanceID,
		InstanceName:       instance.InstanceName,
		InstanceStatus:     instance.Status,
		InstanceStatusName: instanceStatusName,
		InstanceReason:     instanceReason,
		Locked:             instance.Locked,
		//InstanceOwer:        user.UserName,
		UserId:              d.UserId,
		UserName:            d.UserName,
		InstanceCreatedTime: util.TimestampToString(int64(instance.CreatedTime), tz),
		InstanceDescription: instance.Description,

		DeviceTypeID:     d.DeviceTypeID,
		ManageStatus:     d.ManageStatus,
		ManageStatusName: deviceManageStatusName,

		ImageName: image.ImageName,

		Brand:       d.Brand,
		Model:       d.Model,
		Reason:      reason,
		Cabinet:     d.Cabinet,
		UPosition:   d.UPosition,
		Sn:          d.Sn,
		Description: d.Description,

		Mac1:        d.Mac1,
		Mac2:        d.Mac2,
		SwitchIP1:   d.SwitchIP1,
		SwitchPort1: d.SwitchPort1,
		SwitchIP2:   d.SwitchIP2,
		SwitchPort2: d.SwitchPort2,
		//如果设备不存在，从idc中获取
		IloUser:         iloUser,
		IloPassword:     iloPassword,
		IloIP:           d.IloIP, //必填
		SwitchUser1:     switchUser1,
		SwitchPassword1: switchPassword1,
		SwitchUser2:     switchUser2,
		SwitchPassword2: switchPassword2,
		SwitchIP:        d.SwitchIP,

		Mask:            d.Mask,
		Gateway:         d.Gateway,
		PrivateIPv4:     d.PrivateIPv4,
		PrivateEth1IPv4: d.PrivateEth1IPv4,
		PrivateIPv6:     d.PrivateIPv6,
		PrivateEth1IPv6: d.PrivateEth1IPv6,
		AdapterID:       d.AdapterID,

		Enclosure1:  encloure1,
		Slot1:       slot1,
		Enclosure2:  encloure2,
		Slot2:       slot2,
		RaidDriver:  d.RaidDriver,
		CreatedBy:   d.CreatedBy,
		UpdatedBy:   d.UpdatedBy,
		CreatedTime: util.TimestampToString(int64(d.CreatedTime), tz),
		UpdatedTime: util.TimestampToString(int64(d.UpdatedTime), tz),

		//以下字段从其他表中获取
		//idc
		IdcName:   idcName,
		IDcNameEn: idc.NameEn,
		//deviceType

		DeviceType:       entity.DeviceType,
		DeviceSeries:     d.DeviceSeries,
		DeviceSeriesName: deviceSeriesName,

		DeviceTypeName:  entity.Name,
		Architecture:    entity.Architecture,
		CPUAmount:       entity.CPUAmount,
		CPUCores:        entity.CPUCores,
		CPUFrequency:    entity.CPUFrequency,
		CPUManufacturer: entity.CPUManufacturer,
		CPUModel:        entity.CPUModel,

		MemType:                   entity.MemType,
		MemAmount:                 entity.MemAmount,
		MemSize:                   entity.MemSize,
		MemFrequency:              entity.MemFrequency,
		NicAmount:                 entity.NicAmount,
		NicRate:                   entity.NicRate,
		InterfaceMode:             entity.InterfaceMode,
		SystemVolumeType:          entity.SystemVolumeType,
		SystemVolumeInterfaceType: entity.DataVolumeInterfaceType,
		SystemVolumeAmount:        entity.SystemVolumeAmount,
		SystemVolumeSize:          entity.SystemVolumeSize,
		DataVolumeType:            entity.DataVolumeType,
		DataVolumeInterfaceType:   entity.SystemVolumeInterfaceType,
		DataVolumeAmount:          entity.DataVolumeAmount,
		DataVolumeSize:            entity.DataVolumeSize,
		GpuAmount:                 entity.GpuAmount,
		GpuManufacturer:           entity.GpuManufacturer,
		GpuModel:                  entity.GpuModel,

		//拼装信息
		CpuInfo: deviceTypeMapInfo[entity.DeviceTypeID]["cpu_"],
		MemInfo: deviceTypeMapInfo[entity.DeviceTypeID]["mem_"],
		SvInfo:  deviceTypeMapInfo[entity.DeviceTypeID]["sv_"],
		DvInfo:  deviceTypeMapInfo[entity.DeviceTypeID]["dv_"],
		GpuInfo: deviceTypeMapInfo[entity.DeviceTypeID]["gpu_"],
		NicInfo: deviceTypeMapInfo[entity.DeviceTypeID]["nic_"],

		CollectStatus:     collectStatus,
		CollectFailReason: d.CollectFailReason,
		IsNeedRaid:        entity.IsNeedRaid,
	}
}
func GetDeviceTypeInfo(logger *log.Logger, v *deviceTypeDao.DeviceType) map[string]string {
	language := logger.GetPoint("language").(string)

	cpu_ := v.CPUManufacturer + " " + v.CPUModel + "(" + strconv.Itoa(int(v.CPUAmount)) + "*" + strconv.Itoa(int(v.CPUCores)) + "物理核," + v.CPUFrequency + "GHz)"
	if language == baseLogic.EN_US {
		cpu_ = v.CPUManufacturer + " " + v.CPUModel + "(" + strconv.Itoa(int(v.CPUAmount)) + "*" + strconv.Itoa(int(v.CPUCores)) + "cores," + v.CPUFrequency + "GHz)"
	}
	mem_ := strconv.Itoa(int(v.MemAmount)*int(v.MemSize)) + "GB(" + strconv.Itoa(int(v.MemSize)) + "GB*" + strconv.Itoa(int(v.MemAmount)) + ")" + v.MemType + " " + strconv.Itoa(int(v.MemFrequency)) + "MHz"
	// 系统盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式)
	// 240GB 240GB SATA SSD*2 RAID1
	//sv_ := strconv.Itoa(int(v.SystemVolumeAmount)*int(v.SystemVolumeSize)) + "GB(" + strconv.Itoa(int(v.SystemVolumeSize)) + "GB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount)) + ")"
	//if v.SystemVolumeUnit == "TB" {
	//	sv_ = Trim(strconv.FormatFloat(float64(v.SystemVolumeAmount)*float64(v.SystemVolumeSize)/1000, 'f', 2, 64)) + "TB(" + Trim(strconv.FormatFloat(float64(v.SystemVolumeSize)/1000, 'f', 2, 64)) + "TB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount)) + ")"
	//}
	//// 数据盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式)
	//dv_ := strconv.Itoa(int(v.DataVolumeAmount)*int(v.DataVolumeSize)) + "GB(" + strconv.Itoa(int(v.DataVolumeSize)) + "GB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount)) + ")"
	//if v.DataVolumeUnit == "TB" {
	//	dv_ = Trim(strconv.FormatFloat(float64(v.DataVolumeAmount)*float64(v.DataVolumeSize)/1000, 'f', 2, 64)) + "TB(" + Trim(strconv.FormatFloat(float64(v.DataVolumeSize)/1000, 'f', 2, 64)) + "TB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount)) + ")" //  + "," + deviceTypeRaid + ")"
	//}
	//if v.DataVolumeAmount == 0 {
	//	dv_ = ""
	//}
	//// 显示品牌 显卡型号 个数，例 NVIDIA P40 * 4
	gpu_ := v.GpuManufacturer + " " + v.GpuModel + "*" + strconv.Itoa(int(v.GpuAmount))
	if v.GpuAmount == 0 {
		gpu_ = ""
	}
	nic_ := strconv.Itoa(int(v.NicAmount)) + "*" + strconv.Itoa(int(v.NicRate)) + "GE"
	return map[string]string{
		"cpu_": cpu_,
		"mem_": mem_,
		//"sv_":  sv_,
		//"dv_":  dv_,
		"gpu_": gpu_,
		"nic_": nic_,
	}
}
func Trim(str string) string {
	s := strings.TrimRight(str, "0")
	s = strings.TrimRight(s, ".")
	return s
}
func QueryDevices(logger *log.Logger, param requestTypes.QueryDevicesRequest, p util.Pageable) ([]*responseTypes.Device, int64, error) {
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"idc_id": param.IDcID,
		"is_del": 0,
		//如果还需要查询条件，继续在下面添加条件
	}
	if param.DeviceSeries != "" {
		query["device_series.in"] = strings.Split(param.DeviceSeries, ",")
	}
	if param.Sn != "" {
		query["sn.like"] = "%" + param.Sn + "%"
	}
	if param.InstanceID != "" {
		query["instance_id"] = param.InstanceID
	}
	var instanceIDs []string
	if param.InstanceName != "" {
		if _, ok := query["instance_name"]; ok {
			q := map[string]interface{}{
				"instance_name": query["instance_name"],
			}
			instances, err := instanceDao.GetAllInstance(logger, q)
			if err != nil {
				logger.Warnf("QueryDevices.GetAllInstance by instance name error, instance_nameL:%s, error:%s", query["instance_name"], err.Error())
			} else {
				for _, v := range instances {
					instanceIDs = append(instanceIDs, v.InstanceID)
				}
			}
		}
	}

	//instance_id和instance_name两个参数都存在时，需要都检测
	if param.InstanceID != "" && param.InstanceName != "" {
		if util.InArray(param.InstanceID, instanceIDs) {
			query["instance_id"] = param.InstanceID
		} else { //两个参数冲突，不给任何结果
			return nil, 0, fmt.Errorf("检测到instance_id 和 instance_id参数不匹配")
		}
	} else if param.InstanceID != "" && param.InstanceName == "" {
		query["instance_id"] = param.InstanceID
	} else if param.InstanceID == "" && param.InstanceName != "" {
		query["instance_id.in"] = instanceIDs
	}

	if param.DeviceTypeID != "" {
		query["device_type_id.in"] = strings.Split(param.DeviceTypeID, ",")
	}
	if param.ManageStatus != "" {
		query["manage_status.in"] = strings.Split(param.ManageStatus, ",")
	}
	if param.UserID != "" {
		query["user_id"] = param.UserID
	}
	if param.UserName != "" {
		query["user_name.like"] = "%" + param.UserName + "%"
	}
	if param.IloIP != "" {
		query["ilo_ip"] = param.IloIP
	}
	if param.IPV4 != "" {
		query["private_ipv4"] = param.IPV4
	}
	if param.IPV6 != "" {
		query["private_ipv6"] = param.IPV6
	}

	if param.CollectStatus != "" {
		query["collect_status"] = param.CollectStatus
	}

	count, err := deviceDao.GetDeviceCount(logger, query)
	if err != nil {
		logger.Warn("GetDeviceCount dao error:", err.Error())
		return nil, 0, err
	}
	entityList := []*deviceDao.Device{}
	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = deviceDao.GetMultiDevice(logger, query, nil, []string{"created_time"}, []string{"desc"}, 0, 0)
	} else {
		entityList, err = deviceDao.GetMultiDevice(logger, query, nil, []string{"created_time"}, []string{"desc"}, offset, limit)
	}
	if err != nil {
		logger.Warn("GetMultiDevice dao error:", err.Error())
		return nil, 0, err
	}
	var deviceList = []*responseTypes.Device{}
	idcALl, _ := idcDao.GetAllIdc(logger, map[string]interface{}{"is_del": baseLogic.IS_NOT_DEL})
	idcMap := map[string]idcDao.Idc{}
	for _, v := range idcALl {
		idcMap[v.IDcID] = *v
	}
	deviceTypeALl, _ := deviceTypeDao.GetAllDeviceType(logger, map[string]interface{}{"is_del": baseLogic.IS_NOT_DEL})
	deviceTypeMap := map[string]deviceTypeDao.DeviceType{}
	deviceTypeMapInfo := map[string]map[string]string{}
	//获取raid的信息
	raidMap := make(map[string]responseTypes.Raid) //	请求一次，避免循环请求
	raidList, raidsErr := raidLogic.QueryRaidsAll(logger)
	if raidsErr == nil && len(raidList) != 0 {
		for _, vv := range raidList {
			raidMap[vv.RaidID] = *vv
		}
	}

	for _, v := range deviceTypeALl {
		//deviceTypeRaid, err := rDeviceTypeRaidDao.GetAllRDeviceTypeRaid(logger, map[string]interface{}{
		//	"device_type_id": v.DeviceTypeID,
		//	"is_del":         0,
		//})
		//if err != nil {
		//	panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
		//}
		//raid := ""
		//if len(deviceTypeRaid) != 0 {
		//	for _, vv := range deviceTypeRaid {
		//		raid = raid + vv.RaidID + ","
		//	}
		//	raid = strings.Trim(raid, ",")
		//}
		////计算raid对应map中的Name
		//array := []string{} // 拆分查询
		//deviceTypeRaidName := ""
		//if len(raidList) != 0 {
		//	array = strings.Split(raid, ",")
		//}
		//if len(array) != 0 {
		//	raid_ := ""
		//	for _, str := range array {
		//		raid_ = raid_ + raidMap[str].Name + "," // RAID0,RAID1
		//	}
		//	deviceTypeRaidName = raid_[0 : len(raid_)-1]
		//}

		info := GetDeviceTypeInfo(logger, v)
		deviceTypeMap[v.DeviceTypeID] = *v
		//fmt.Println(info, util.ObjToJson(info), info["nic_"])
		//拼装数据，给上层api直接使用
		deviceTypeMapInfo[v.DeviceTypeID] = map[string]string{
			"cpu_": info["cpu_"],
			"mem_": info["mem_"],
			//"sv_":  info["sv_"],
			//"dv_":  info["dv_"],
			"gpu_": info["gpu_"],
			"nic_": info["nic_"],
		}

	}
	instanceALl, _ := instanceDao.GetAllInstance(logger, map[string]interface{}{"is_del": baseLogic.IS_NOT_DEL})
	instanceMap := map[string]instanceDao.Instance{}
	for _, v := range instanceALl {
		instanceMap[v.InstanceID] = *v
	}
	imageALl, _ := imageDao.GetAllImage(logger, map[string]interface{}{"is_del": baseLogic.IS_NOT_DEL})
	imageMap := map[string]imageDao.Image{}
	for _, v := range imageALl {
		imageMap[v.ImageID] = *v
	}
	//userALl, _ := userDao.GetAllUser(logger, map[string]interface{}{"is_del": baseLogic.IS_NOT_DEL})
	//userMap := map[string]userDao.User{}
	//for _, v := range userALl {
	//	userMap[v.UserID] = v
	//}

	for _, entity := range entityList {
		device := DeviceEntity2Device(logger, entity, idcMap, deviceTypeMap, deviceTypeMapInfo, instanceMap, imageMap)
		deviceList = append(deviceList, device)
	}
	return deviceList, count, err
}

func Save(logger *log.Logger, params requestTypes.CreateDevicesRequest) ([]string, error) {

	if len(params.Devices) > 10000 {
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("单次上传不能超过10000条"), fmt.Sprintf("exceed 10000")))
	}
	_, err := idcDao.GetIdcById(logger, params.IDcID)
	if err != nil {
		panic(constant.BuildInvalidArgumentWithArgs("机房不存在", "idcId not exist"))
	}

	/*
		deviceType, err := deviceTypeDao.GetOneDeviceType(logger, map[string]interface{}{"device_type_id": params.DeviceTypeID})
		if err != nil {
			panic(constant.BuildInvalidArgumentWithArgs("机型不存在", "deviceTypeId not exist"))
		}

		if deviceType.NicAmount == 2 && deviceType.InterfaceMode != "bond" {
			for _, v := range params.Devices {
				if v.PrivateEth1IPv4 == "" || v.MaskEth1 == "" {
					panic(constant.BuildInvalidArgumentWithArgs("双网口非bond时,privateEth1Ipv4和maskEth1为必填项", "dual model must have privateEth1Ipv4 and maskEth1 params"))
				}
			}
		}
	*/

	//fmt.Println(deviceType)
	deviceIds := []string{}
	list, err := deviceDao.GetAllDevice(logger, map[string]interface{}{
		"is_del": baseLogic.IS_NOT_DEL,
	})
	if err != nil {
		logger.Warnf("createDevice.GetAllDevice error:%s", err.Error())
	}

	// var licenseNum int
	// license, err := licenseLogic.GetLicenseContent(logger)
	// if err != nil {
	// 	logger.Warnf("createDevice.GetLicenseContent error:%s", err.Error())
	// }
	// if license == nil || license.NodesNum == 0 {
	// 	logger.Warn("createDevice license empty")
	// 	panic(constant.BuildInvalidArgumentWithArgs("请联系客服授权", "please buy business license"))
	// } else {
	// 	licenseNum = license.NodesNum
	// }
	// if len(list)+len(params.Devices) > licenseNum {
	// 	logger.Warnf("createDevice nodes exceed license_num, licenseNum:%d, actual num:%d", licenseNum, len(list)+len(params.Devices))
	// 	panic(constant.BuildInvalidArgumentWithArgs("节点数超过授权数,请联系客服", "nodes exceed than license num"))
	// }

	listSn := map[string]string{}
	listIloIp := map[string]string{}
	listIPV6 := map[string]string{}
	for _, v := range list {
		listSn[v.Sn] = v.Sn
		listIloIp[v.IloIP] = v.IloIP
		if v.PrivateIPv6 != "" { //不为空才判断
			listIPV6[v.PrivateIPv6] = v.PrivateIPv6
		}
	}
	sns := []string{}
	iloIps := []string{}
	ipv6s := []string{}
	for _, v := range params.Devices {
		if util.InArrayString(v.Sn, sns) {
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("文件中sn %s 重复", v.Sn), fmt.Sprintf("file sn %s repeat", v.Sn)))
		}
		sns = append(sns, v.Sn)
		if util.InArrayString(v.IloIP, iloIps) {
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("文件中带外IP %s 重复", v.IloIP), fmt.Sprintf("file iloIp %s repeat", v.IloIP)))
		}
		iloIps = append(iloIps, v.IloIP)
		if util.InArrayString(v.PrivateIPv6, ipv6s) {
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("文件中IPV6 %s 重复", v.PrivateIPv6), fmt.Sprintf("file IPV6 %s repeat", v.PrivateIPv6)))
		}
		ipv6s = append(ipv6s, v.PrivateIPv6)
		checkDevice(logger, *v, listSn, listIloIp, listIPV6)
	}
	n := 200 //每个协程处理的数量
	//fmt.Println(57/n, float64(57/n), float64(57)/float64(n), math.Ceil(float64(57)/float64(n)), int(math.Floor(float64(57)/float64(n))))
	count := int(math.Ceil(float64(len(params.Devices)) / float64(n))) //协程数量
	fmt.Println("总数", len(params.Devices), "协程数", count)
	for i := 0; i < count; i++ {
		end := (i + 1) * n
		if end > len(params.Devices) {
			end = len(params.Devices)
		}
		wg.Add(1)
		go saveDevices(logger, params.IDcID, params.Devices[i*n:end])
	}
	wg.Wait()
	return deviceIds, nil
}
func checkDevice(logger *log.Logger, param requestTypes.CreateDeviceSpec, listSn, listIloIp, listIPV6 map[string]string) {
	specialMatch := regexp.MustCompile("^.{1,128}$").MatchString //[一-龥_a-zA-Z0-9_-]
	if !specialMatch(param.Sn) {
		panic(constant.BuildInvalidArgumentWithArgs("Sn 不合法", "Sn invalidate"))
	}
	if !specialMatch(param.Cabinet) {
		panic(constant.BuildInvalidArgumentWithArgs("Cabinet 不合法", "Cabinet invalidate"))
	}
	if !specialMatch(param.UPosition) {
		panic(constant.BuildInvalidArgumentWithArgs("UPosition 不合法", "UPosition invalidate"))
	}
	// if !specialMatch(param.Brand) {
	// 	panic(constant.BuildInvalidArgumentWithArgs("Brand 不合法", "Brand invalidate"))
	// }
	// if !specialMatch(param.Model) {
	// 	panic(constant.BuildInvalidArgumentWithArgs("Model 不合法", "Model invalidate"))
	// }
	if _, err := net.ParseMAC(param.Mac1); err != nil {
		panic(constant.BuildInvalidArgumentWithArgs("Mac1 不合法", "Mac1 invalidate"))
	}
	// if _, err := net.ParseMAC(param.Mac2); err != nil {
	// 	panic(constant.BuildInvalidArgumentWithArgs("Mac2 不合法", "Mac2 invalidate"))
	// }
	// specialMatch = regexp.MustCompile("^[a-zA-Z/0-9_-]{1,128}$").MatchString
	// if !specialMatch(param.SwitchPort1) {
	// 	panic(constant.BuildInvalidArgumentWithArgs("switchport1 不合法", "switchport1 invalidate"))
	// }
	// specialMatch = regexp.MustCompile("^[a-zA-Z/0-9_-]{1,128}$").MatchString
	// if !specialMatch(param.SwitchPort2) {
	// 	panic(constant.BuildInvalidArgumentWithArgs("switchport2 不合法", "switchport2 invalidate"))
	// }
	if net.ParseIP(param.Mask) == nil {
		panic(constant.BuildInvalidArgumentWithArgs("Mask 不合法", "Mask invalidate"))
	}
	if net.ParseIP(param.Gateway) == nil {
		panic(constant.BuildInvalidArgumentWithArgs("Gateway 不合法", "Gateway invalidate"))
	}
	if net.ParseIP(param.IloIP) == nil {
		panic(constant.BuildInvalidArgumentWithArgs("IloIP 不合法", "IloIP invalidate"))
	}
	if strings.Contains(param.PrivateIPv4, ":") {
		panic(constant.BuildInvalidArgumentWithArgs("PrivateIPv4 不合法", "PrivateIPv4 invalidate"))
	}
	if net.ParseIP(param.PrivateIPv4) == nil {
		panic(constant.BuildInvalidArgumentWithArgs("PrivateIPv4 不合法", "PrivateIPv4 invalidate"))
	}
	if param.PrivateIPv6 != "" {
		if strings.Contains(param.PrivateIPv6, "/") { //2403:1EC0:8549:60C0::1/64
			arr := strings.Split(param.PrivateIPv6, "/")
			ipv6 := arr[0]
			mask := arr[1]
			maskCount, err := strconv.Atoi(mask)
			fmt.Println(maskCount, err)
			if err != nil {
				panic(constant.BuildInvalidArgumentWithArgs("PrivateIPv6 不合法", "PrivateIPv6 invalidate"))
			}
			if !strings.Contains(ipv6, ":") {
				panic(constant.BuildInvalidArgumentWithArgs("PrivateIPv6 不合法", "PrivateIPv6 invalidate"))
			}
			if net.ParseIP(ipv6) == nil || (maskCount < 0 || maskCount > 128) {
				panic(constant.BuildInvalidArgumentWithArgs("PrivateIPv6 不合法", "PrivateIPv6 invalidate"))
			}
		} else {
			ipv6 := param.PrivateIPv6
			if !strings.Contains(ipv6, ":") {
				panic(constant.BuildInvalidArgumentWithArgs("PrivateIPv6 不合法", "PrivateIPv6 invalidate"))
			}
			if net.ParseIP(ipv6) == nil {
				panic(constant.BuildInvalidArgumentWithArgs("PrivateIPv6 不合法", "PrivateIPv6 invalidate"))
			}
		}
	}
	// if net.ParseIP(param.SwitchIP1) == nil && param.SwitchIP1 != "" {
	// 	panic(constant.BuildInvalidArgumentWithArgs("SwitchIP1 不合法", "SwitchIP1 invalidate"))
	// }
	// if net.ParseIP(param.SwitchIP2) == nil && param.SwitchIP2 != "" {
	// 	panic(constant.BuildInvalidArgumentWithArgs("SwitchIP1 不合法", "SwitchIP1 invalidate"))
	// }
	specialMatch = regexp.MustCompile("^.{1,128}$").MatchString
	if !specialMatch(param.IloUser) && param.IloUser != "" {
		panic(constant.BuildInvalidArgumentWithArgs("IloUser 不合法", "IloUser invalidate"))
	}
	if !specialMatch(param.IloPassword) && param.IloPassword != "" {
		panic(constant.BuildInvalidArgumentWithArgs("IloPassword 不合法", "IloPassword invalidate"))
	}
	// if !specialMatch(param.SwitchUser1) && param.SwitchUser1 != "" {
	// 	panic(constant.BuildInvalidArgumentWithArgs("SwitchUser1 不合法", "SwitchUser1 invalidate"))
	// }
	// if !specialMatch(param.SwitchPassword1) && param.SwitchPassword1 != "" {
	// 	panic(constant.BuildInvalidArgumentWithArgs("SwitchPassword1 不合法", "SwitchPassword1 invalidate"))
	// }
	// if !specialMatch(param.SwitchUser2) && param.SwitchUser2 != "" {
	// 	panic(constant.BuildInvalidArgumentWithArgs("SwitchUser2 不合法", "SwitchUser2 invalidate"))
	// }
	// if !specialMatch(param.SwitchPassword2) && param.SwitchPassword2 != "" {
	// 	panic(constant.BuildInvalidArgumentWithArgs("SwitchPassword2 不合法", "SwitchPassword2 invalidate"))
	// }
	//list, _ := deviceDao.GetOneDevice(logger, map[string]interface{}{
	//	"sn":     param.Sn,
	//	"is_del": baseLogic.IS_NOT_DEL,
	//})
	//if list != nil {
	//	logger.Warn("sn exist:", param.Sn)
	//	//snRepeat = fmt.Sprintf("sn %s 已存在", param.Sn)
	//	//fmt.Println(snRepeat)
	//	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("sn %s 已存在", param.Sn), fmt.Sprintf("sn %s exist", param.Sn)))
	//}
	//list, _ = deviceDao.GetOneDevice(logger, map[string]interface{}{
	//	"ilo_ip": param.IloIP,
	//	"is_del": baseLogic.IS_NOT_DEL,
	//})
	//if list != nil {
	//	logger.Warn("iloIp exist:", param.IloIP)
	//	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("带外IP %s 已存在", param.IloIP), fmt.Sprintf("iloIp %s exist", param.IloIP)))
	//}
	_, ok := listSn[param.Sn]
	if ok { //如果存在
		logger.Warn("sn exist:", param.Sn)
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("sn %s 已存在", param.Sn), fmt.Sprintf("sn %s exist", param.Sn)))
	}
	_, ok = listIloIp[param.IloIP]
	if ok { //如果存在
		logger.Warn("iloIp exist:", param.IloIP)
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("带外IP %s 已存在", param.IloIP), fmt.Sprintf("iloIp %s exist", param.IloIP)))
	}
	if param.PrivateIPv6 != "" {
		_, ok = listIPV6[param.PrivateIPv6]
		if ok { //如果存在
			logger.Warn("IPV6 exist:", param.PrivateIPv6)
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("IPV6 %s 已存在", param.PrivateIPv6), fmt.Sprintf("IPV6 %s exist", param.PrivateIPv6)))
		}
	}
}
func saveDevices(logger *log.Logger, idcId string, list []*requestTypes.CreateDeviceSpec) ([]string, error) {
	sns := []string{}
	deviceIds := []string{}
	devices := []*deviceDao.Device{}
	// disks1 := map[string]*diskDao.Disk{}
	// disks2 := map[string]*diskDao.Disk{}
	defer wg.Done()
	for _, param := range list {
		if param == nil {
			break
		}
		//checkDevice(logger, *param)
		//fmt.Println("sn", k, param.Sn, param.Sn)

		sns = append(sns, param.Sn)
		entity := buildDeviceEntity(logger, idcId, param)
		entity.CreatedTime = int(time.Now().Unix())
		entity.UpdatedTime = int(time.Now().Unix())
		entity.ManageStatus = baseLogic.IN //默认已入库
		entity.CollectStatus = "2"         //2表示未采集
		devices = append(devices, entity)

		/*
			disks1[entity.DeviceID] = &diskDao.Disk{
				DeviceID:    entity.DeviceID,
				Enclosure:   param.Enclosure1,
				Slot:        *param.Slot1,
				DiskType:    volumeType.SYSTEM, //默认只有系统盘
				Size:        strconv.Itoa(deviceType.SystemVolumeSize),
				SizeUnit:    deviceType.SystemVolumeUnit,
				PdType:      deviceType.SystemVolumeInterfaceType,
				AdapterID:   *param.AdapterID,
				CreatedBy:   logger.GetPoint("username").(string),
				UpdatedBy:   logger.GetPoint("username").(string),
				CreatedTime: int(time.Now().Unix()),
				UpdatedTime: int(time.Now().Unix()),
			}

			disks2[entity.DeviceID] = &diskDao.Disk{
				DeviceID:    entity.DeviceID,
				Enclosure:   param.Enclosure2,
				Slot:        param.Slot2,
				DiskType:    volumeType.SYSTEM, //默认只有系统盘
				Size:        strconv.Itoa(deviceType.SystemVolumeSize),
				SizeUnit:    deviceType.SystemVolumeUnit,
				PdType:      deviceType.SystemVolumeInterfaceType,
				AdapterID:   *param.AdapterID,
				CreatedBy:   logger.GetPoint("username").(string),
				CreatedTime: int(time.Now().Unix()),
				UpdatedBy:   logger.GetPoint("username").(string),
				UpdatedTime: int(time.Now().Unix()),
			}
		*/
		deviceIds = append(deviceIds, entity.DeviceID)
	}
	//countParam := map[string]interface{}{
	//	"is_del": 0,
	//	"sn.in":  sns,
	//}
	//count, err := deviceDao.GetDeviceCount(logger, countParam)
	//if err != nil {
	//	logger.Warn("CreateDevices GetDeviceCount error:", sns, err.Error())
	//	return []string{}, err
	//}
	//if count > 0 {
	//	panic(constant.BuildInvalidArgumentWithArgs("sn 已存在", "device sn may already exists"))
	//}
	//20240516 minping 磁盘信息改成在采集时自动插入
	// _, err := deviceDao.CreateDeviceAndDisk(logger, devices, disks1, disks2)
	_, err := deviceDao.AddMultiDevice(logger, devices)
	if err != nil {
		logger.Warn("AddMultiDevice error:", err.Error())
		return []string{}, err
	}
	//for _, device := range list {
	//	_, _, err = net.ParseCIDR(device.PrivateIPv4)
	//	ipv4 := device.PrivateIPv4
	//	if err != nil { //如果不是cird地址，先根据掩码转成cird,假如给的地址是这种格式，10.208.14.81
	//		mask := net.IPMask(net.ParseIP(device.Mask).To4())
	//		prefixSize, _ := mask.Size()
	//		ipv4 = device.PrivateIPv4 + "/" + strconv.Itoa(prefixSize)
	//	}
	//	subnet := util.NetworkIp(ipv4)
	//	msg := CreateDeviceMq{
	//		Action:     "DHCPConfigAddSubnet",
	//		Subnet:     subnet,
	//		Sn:         device.Sn,
	//		SubnetMask: device.Mask,
	//		Routes:     device.Gateway,
	//	}
	//	res, err := rdsUtil.RedisUtil.SetNX("set_subenet_"+subnet, true, time.Second*3600).Result()
	//	fmt.Println("设置子网", res, err, util.ObjToJson(msg))
	//	//if res {
	//	err = sendCreateDeviceEvent(logger, deviceType.IDcID, &msg) //param.UserData, param.AliasIPs, param.ExtensionAliasIps
	//	if err != nil {
	//		return []string{}, err
	//	}
	//}
	//}

	return deviceIds, err
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func CollectDeviceInfo(logger *log.Logger, collects *requestTypes.CollectDeviceInfoRequest) error {

	logid := logger.GetPoint("logid").(string)
	failsns := []string{}

	userId := logger.GetPoint("userId").(string)

	for _, v := range collects.Collects {

		if err := commandDao.DeleteCommandBySn(logger, v.Sn); err != nil {
			logger.Warnf("CollectDeviceInfo.DeleteCommandBySn error, sn:%s, error:%s", v.Sn, err.Error())
		}

		message := rabbitIronicMsgApi.CollectHardwareInfoMessage{}
		message.ApiMessage.RequestId = logid
		message.Sns = []string{v.Sn}
		message.AllowOverride = v.AllowOverride
		message.NetworkType = "retail"
		collectedStatus := "4" //4表示采集失败
		collectFailReason := "mq msg send error"
		event, err := rabbitEvent.NewEvent(message, logid, userId)
		if err != nil {
			logger.Warnf("RetryCommand convert event msg error, sn:%s, error:%s", v.Sn, err.Error())
			failsns = append(failsns, v.Sn)
		} else {
			if err := rabbitMq.SendToApi2Scheduler(event); err != nil {
				logger.Warnf("CollectDeviceInfo SendToScheduler error, msg:%s, error:%s", util.ObjToJson(event), err.Error())
				failsns = append(failsns, v.Sn)
			} else {
				deviceEntity, err := deviceDao.GetBySn(logger, v.Sn)
				if err != nil {
					logger.Warnf("CollectDeviceInfo.GetBySn error, sn:%s, error:%s", v.Sn, err.Error())
				} else {
					collectedStatus = "3" //3表示采集中
					collectFailReason = ""
					deviceEntity.CollectStatus = collectedStatus
					deviceEntity.CollectFailReason = collectFailReason
					if err := deviceDao.UpdateDeviceAnywhere(logger, deviceEntity); err != nil {
						logger.Warnf("CollectDeviceInfo.UpdateDeviceAnywhere error, sn:%s, error:%s", v.Sn, err.Error())
					}
				}
			}
		}

	}
	if len(failsns) == 0 {
		return nil
	}

	return fmt.Errorf("%s collect error!", strings.Join(failsns, ","))
}

func ModifyDevice(logger *log.Logger, DeviceID, manageStatus string) error {
	if DeviceID == "" {
		panic(constant.BuildInvalidArgumentWithArgs("设备ID不能为空", "deviceId invalidate"))
	}
	device, err := deviceDao.GetOneDevice(logger, map[string]interface{}{
		"device_id": DeviceID,
		"is_del":    baseLogic.IS_NOT_DEL,
	})
	if err != nil {
		logger.Warn("GetDeviceTypeByUuid sql error:", DeviceID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	//manageStatus := param.ManageStatus
	if manageStatus == baseLogic.PUTAWAY { //如果执行上架
		if device.DeviceTypeID == "" {
			panic(constant.BuildInvalidArgumentWithArgs("设备必须绑定机型，才可以执行上架操作", "device must bind device_type before putaway"))
		}
		if device.ManageStatus != baseLogic.IN && device.ManageStatus != baseLogic.PUTAWAYFAIL {
			panic(constant.BuildInvalidArgumentWithArgs("设备必须是已入库或者上架失败，才可以执行上架操作", "device manageStatus not support"))
		}
		manageStatus = baseLogic.PUTAWAYING //上架中
	}
	if manageStatus == baseLogic.UNPUTAWAY { //如果执行下架动作
		if device.ManageStatus != baseLogic.PUTAWAY {
			panic(constant.BuildInvalidArgumentWithArgs("设备必须是已上架，才可以执行下架操作", "device manageStatus error"))
		}
		manageStatus = baseLogic.IN //没有下架中，已下架，直接变成已入库

		auditLogLogic.SaveAuditLogs(logger, DeviceID, "", AuditLogsType.AuditLogsUnPutaway)
	}
	entity := &deviceDao.Device{
		DeviceID:     DeviceID,
		ManageStatus: manageStatus,
		UpdatedBy:    logger.GetPoint("username").(string),
		UpdatedTime:  int(time.Now().Unix()),
	}
	if err := deviceDao.UpdateDeviceById(logger, entity); err != nil {
		logger.Warnf("ModifyDevice sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return err
	}
	//接下来修改其他信息，比如需不需要发mq之类的
	if manageStatus == baseLogic.PUTAWAYING { //如果变成了上架中
		if err := sendPutawayDeviceEvent(logger, device.Sn); err != nil {
			logger.Warnf("sendPutawayDeviceEvent error, sn:%s, err:%s", device.Sn, err.Error())
		} else {
			logger.Infof("sendPutawayDeviceEvent success, sn:%s", device.Sn)
		}
	}
	return nil
}

func sendPutawayDeviceEvent(logger *log.Logger, sn string) error {

	logid := logger.GetPoint("logid").(string)
	userId := logger.GetPoint("userId").(string)
	msg := rabbitIronicMsgApi.PutawayDeviceMessage{
		ApiMessage: rabbitIronicMsgApi.ApiMessage{
			RequestId: logid,
		},
		Sn: sn,
	}
	event, err := rabbitEvent.NewEvent(msg, logid, userId)
	if err != nil {
		logger.Warn("rabbitEvent.NewEvent convert event msg error:", sn, err.Error())
		return err
	}
	if err := rabbitMq.SendToApi2Scheduler(event); err != nil {
		logger.Warnf("sendPutawayDeviceEvent SendToScheduler  error, msg:%s, error:%s", util.ObjToJson(event), err.Error())
		return err
	} else {
		logger.Infof("sendPutawayDeviceEvent SendToScheduler success, msg:%s", util.ObjToJson(event))
	}
	return nil
}

func ModifyDeviceDescription(logger *log.Logger, param *requestTypes.ModifyDevicesRequest, deviceId string) error {
	if deviceId == "" {
		panic(constant.BuildInvalidArgumentWithArgs("设备ID不能为空", "deviceId invalidate"))
	}
	entity, err := deviceDao.GetOneDevice(logger, map[string]interface{}{
		"device_id": deviceId,
		"is_del":    baseLogic.IS_NOT_DEL,
	})
	if err != nil || entity == nil {
		logger.Warn("GetDeviceTypeByUuid sql error:", deviceId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if param.Description != nil {
		entity.Description = *param.Description
	}
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())

	if err := deviceDao.SaveDevice(logger, entity); err != nil {
		logger.Warnf("ModifyDevice sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return err
	}
	return nil
}
func ModifyDeviceByWhere(logger *log.Logger, column string, newData string, deviceTypeId string) error {
	_, err := deviceDao.GetOneDevice(logger, map[string]interface{}{
		"device_type_id": deviceTypeId,
		"is_del":         baseLogic.IS_NOT_DEL,
	})
	if err != nil {
		logger.Info("机型没有设备，不做device_series的修改", deviceTypeId)
		//如果设备表没有对应的机型，直接返回，不做修改
		return nil
	}
	entity := &deviceDao.Device{
		DeviceTypeID: deviceTypeId,
	}
	if column == "device_series" {
		entity.DeviceSeries = newData
	}
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	if err := deviceDao.UpdateDeviceByDeviceTypeId(logger, entity); err != nil {
		logger.Warnf("ModifyDevice sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return err
	}
	return nil
}

func ModifyDeviceByDeviceId(logger *log.Logger, param *requestTypes.ModifyAllDevicesRequest, DeviceID string) error {
	_, err := deviceDao.GetOneDevice(logger, map[string]interface{}{
		"device_id": DeviceID,
		"is_del":    baseLogic.IS_NOT_DEL,
	})
	if err != nil {
		logger.Warn("GetDeviceTypeByUuid sql error:", DeviceID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	entity := &deviceDao.Device{
		DeviceID:     DeviceID,
		ManageStatus: param.ManageStatus,
		InstanceID:   param.InstanceID,
		UserId:       param.UserId,
		UserName:     param.UserName,
		UpdatedBy:    logger.GetPoint("username").(string),
		UpdatedTime:  int(time.Now().Unix()),
	}
	if err := deviceDao.UpdateDeviceById(logger, entity); err != nil {
		logger.Warnf("ModifyDevice sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return err
	}
	//接下来修改其他信息，比如需不需要发mq之类的
	return nil
}

func Delete(logger *log.Logger, DeviceID string) error {
	d, err := deviceDao.GetDeviceById(logger, DeviceID)
	if err != nil {
		logger.Warn("DeleteDevices GetAllDevice sql error:", DeviceID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	s := []string{SaleStatus.PUTAWAYFAIL, SaleStatus.REMOVED, SaleStatus.IN}
	if !util.InArray(d.ManageStatus, s) {
		panic(constant.BuildInvalidArgumentWithArgs("设备sn"+d.Sn+"状态为"+d.ManageStatus+",不允许删除", fmt.Sprintf("device %s is %s,not allow delete", d.Sn, d.ManageStatus)))
	}

	d.UpdatedBy = logger.GetPoint("username").(string)
	d.UpdatedTime = int(time.Now().Unix())
	d.DeletedTime = int(time.Now().Unix())
	d.IsDel = baseLogic.IS_DEL

	err = deviceDao.UpdateDeviceById(logger, d)
	if err != nil {
		logger.Warn("DeleteDevices UpdateMultiDevices sql error:", err.Error())
		return err
	}

	auditLogLogic.DeleteAuditLogsByDeviceId(logger, DeviceID)

	if err := commandDao.DeleteCommandBySn(logger, d.Sn); err != nil {
		logger.Warn("DeleteDevices DeleteCommandBySn sql error:", err.Error())
		return err
	}

	q := map[string]interface{}{
		"is_del": 0,
		"sn":     d.Sn,
	}
	u := map[string]interface{}{
		"is_del": 1,
	}

	if err := CpsLogItemsDao.UpdateCpsLogItemss(logger, q, u); err != nil {
		logger.Warnf("DeleteDevice.UpdateCpsLogItemss error, sn:%s, error:%s", d.Sn, err.Error())
	} else {
		if err := CpsLogContentCollectionDao.UpdateCpsLogContentCollections(logger, q, u); err != nil {
			logger.Warnf("DeleteDevice.UpdateCpsLogContentCollections error, sn:%s, error:%s", d.Sn, err.Error())
		}
	}

	return nil
}

func RemoveDevice(logger *log.Logger, DeviceID string) error {

	d, err := deviceDao.GetDeviceById(logger, DeviceID)
	if err != nil {
		logger.Warn("DeleteDevices GetAllDevice sql error:", DeviceID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if d.ManageStatus != SaleStatus.CREATED {
		panic(constant.BuildInvalidArgumentWithArgs("设备状态不是已创建", "device manageStatus is not created"))
	}
	instance, err := instanceDao.GetInstanceByDeviceId(logger, DeviceID)
	if err != nil {
		logger.Warn("DeleteDevices GetAllDevice sql error:", DeviceID, err.Error())
		panic(constant.NOT_SUPPORTED)
	}

	s := []string{instanceStatus.RUNNING, instanceStatus.STOPPED, instanceStatus.REINSTALL_ERROR, instanceStatus.RESETPASSWD_ERROR}
	if !util.InArray(instance.Status, s) {
		panic(constant.BuildInvalidArgumentWithArgs("实例状态不是运行或者关机或者可移除的错误状态", "instance status is not running  or stopped or supported error status"))
	}

	//}
	//updates := map[string]interface{}{
	//	"is_del":       1,
	//	"updated_by":   logger.GetPoint("username").(string),
	//	"updated_time": int(time.Now().Unix()),
	//}
	//err = deviceDao.UpdateMultiDevices(logger, param, updates)
	d.InstanceID = ""
	d.UserId = ""
	d.UserName = ""
	d.ManageStatus = SaleStatus.REMOVED
	d.UpdatedTime = int(time.Now().Unix())
	d.UpdatedBy = logger.GetPoint("username").(string)
	err = deviceDao.SaveDevice(logger, d)
	if err != nil {
		logger.Warn("removeDevices UpdateMultiDevices sql error:", err.Error())
		return err
	}
	//err = interfaceDao.UpdateMultiInterfaces(logger, param, updates)
	//if err != nil {
	//	logger.Warn("DeleteDevices UpdateMultiInterfaces sql error:", err.Error())
	//	return err
	//}
	instance.IsDel = baseLogic.IS_DEL
	if err := instanceDao.UpdateInstanceById(logger, instance); err != nil {
		logger.Warn("removeDevices deleteInstance sql error:", err.Error())
		return err
	}

	if err := commandDao.DeleteCommandBySn(logger, d.Sn); err != nil {
		logger.Warn("removeDevices DeleteCommandBySn sql error:", err.Error())
		return err
	}

	auditLogLogic.SaveAuditLogs(logger, DeviceID, "", AuditLogsType.AuditLogsRemoveDevice)

	return nil
}

func QueryDeviceStock(logger *log.Logger, region, az, deviceType string) ([]*deviceDao.DeviceStock, error) {
	deviceStocks, err := deviceDao.GetDeviceStock(logger, deviceType, region, az)
	if err != nil {
		logger.Warn("GetDeviceStock GetAllDevice sql error:", deviceType, err.Error())
		return nil, err
	}
	return deviceStocks, nil
}

func GetAndCheckById(logger *log.Logger, DeviceID string) (*responseTypes.Device, error) {
	device, err := GetById(logger, DeviceID)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func GetFirstVersionUnavailableSns(logger *log.Logger) ([]string, error) {
	sns, err := deviceDao.GetFirstVersionDeviceSns(logger)
	if err != nil {
		logger.Warn("GetFirstVersionUnavailableSns GetFirstVersionDeviceSns sql error:", err.Error())
		return nil, err
	}
	reinstall_sns, err := deviceDao.GetFirstVersionAndReInstallDeviceSns(logger)
	if err != nil {
		logger.Warn("GetFirstVersionUnavailableSns GetFirstVersionAndReInstallDeviceSns sql error:", err.Error())
		return nil, err
	}
	res := []string{}
	for _, v := range sns {
		if exist := util.InArray(v, reinstall_sns); !exist {
			res = append(res, v)
		}
	}
	return res, nil
}

func RestartDhcp(logger *log.Logger, az string) error {

	logid := logger.GetPoint("logid").(string)
	userId := logger.GetPoint("userId").(string)
	msg := rabbitIronicMsgApi.RestartDhcpMessage{}
	msg.ApiMessage.RequestId = logger.GetPoint("logid").(string)
	msg.Az = az
	event, err := rabbitEvent.NewEvent(msg, logid, userId)
	if err != nil {
		logger.Warn("RestartDhcp convert RestartDhcpMessage to event error:", err.Error())
		return err
	}
	if err := rabbitMq.SendToApi2Scheduler(event); err != nil {
		logger.Warnf("RestartDhcp SendToScheduler error, msg:%s, error:%s", util.ObjToJson(msg), err.Error())
		return err
	}
	return nil
}

// 设备详情
func GetById(logger *log.Logger, DeviceID string) (*responseTypes.Device, error) {
	entity, err := deviceDao.GetDeviceById(logger, DeviceID)
	if err != nil {
		logger.Warn("GetById GetDeviceById sql error:", DeviceID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	raidMap := make(map[string]responseTypes.Raid) //	请求一次，避免循环请求
	raidList, raidsErr := raidLogic.QueryRaidsAll(logger)
	if raidsErr == nil && len(raidList) != 0 {
		for _, vv := range raidList {
			raidMap[vv.RaidID] = *vv
		}
	}

	//deviceTypeRaid, err := rDeviceTypeRaidDao.GetAllRDeviceTypeRaid(logger, map[string]interface{}{
	//	"device_type_id": entity.DeviceTypeID,
	//	"is_del":         0,
	//})
	//if err != nil {
	//	panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	//}
	//raid := ""
	//if len(deviceTypeRaid) != 0 {
	//	for _, vv := range deviceTypeRaid {
	//		raid = raid + vv.RaidID + ","
	//	}
	//	raid = strings.Trim(raid, ",")
	//}
	////计算raid对应map中的Name
	//array := []string{} // 拆分查询
	//deviceTypeRaidName := ""
	//if len(raidList) != 0 {
	//	array = strings.Split(raid, ",")
	//}
	//if len(array) != 0 {
	//	raid_ := ""
	//	for _, str := range array {
	//		raid_ = raid_ + raidMap[str].Name + "," // RAID0,RAID1
	//	}
	//	deviceTypeRaidName = raid_[0 : len(raid_)-1]
	//}
	deviceTypeMapInfo := map[string]map[string]string{}
	if entity.DeviceTypeID != "" {
		v, err := deviceTypeDao.GetDeviceTypeById(logger, entity.DeviceTypeID)
		if err != nil {
			v = &deviceTypeDao.DeviceType{}
		}
		//目的是获取拼装出来的机型数据
		info := GetDeviceTypeInfo(logger, v)

		//fmt.Println(info, util.ObjToJson(info), info["nic_"])
		//拼装数据，给上层api直接使用

		deviceTypeMapInfo[v.DeviceTypeID] = map[string]string{
			"cpu_": info["cpu_"],
			"mem_": info["mem_"],
			//"sv_":  info["sv_"],
			//"dv_":  info["dv_"],
			"gpu_": info["gpu_"],
			"nic_": info["nic_"],
		}
	}

	device := DeviceEntity2DeviceOne(logger, entity, deviceTypeMapInfo)
	return device, nil
}
func IDevice2DeviceEntity(i idcApi.IDevice) *deviceDao.Device {
	return &deviceDao.Device{
		Sn:        i.Sn,
		IloIP:     i.IloIP,
		Cabinet:   i.Cabinet,
		UPosition: i.UPosition,
	}
}
func SaveDevicesFromIdc(logger *log.Logger, param *requestTypes.CreateIdcDevicesRequest) error {
	device_map := idcApi.QueryDevices(logger, param.Sns)
	toInsert := []*deviceDao.Device{}
	toUpdate := []*deviceDao.Device{}
	for _, i_device := range device_map {
		device_entity, err := deviceDao.GetBySn(logger, i_device.Sn)
		if err != nil {
			logger.Warn("SaveDevicesFromIdc GetBySn sql error:", i_device.Sn, err.Error())
			device_entity = IDevice2DeviceEntity(i_device)
			//device_entity.Region = logger.GetPoint("region").(string)
			//device_entity.Az = logger.GetPoint("az").(string)
			//device_entity.DeviceType = param.DeviceType
			toInsert = append(toInsert, device_entity)
		} else {
			device_entity.UPosition = i_device.UPosition
			device_entity.Cabinet = i_device.Cabinet
			toUpdate = append(toUpdate, device_entity)
		}
	}
	if _, err := deviceDao.AddMultiDevice(logger, toInsert); err != nil {
		logger.Warn("SaveDevicesFromIdc AddMultiDevice sql error:", err.Error())
		return err
	}
	for _, device_entity := range toUpdate {
		if err := deviceDao.UpdateDeviceById(logger, device_entity); err != nil {
			logger.Warn("SaveDevicesFromIdc UpdateDeviceById sql error:", err.Error())
			return err
		}
	}
	return nil
}

func ImportBms(logger *log.Logger, file multipart.File, size int64) error {
	beans, err := parseXlsx(logger, file, size)
	if err != nil {
		return err
	}
	toInsert := []*deviceDao.Device{}
	toUpdate := []*deviceDao.Device{}
	for idx, bean := range beans {
		logger.Point(fmt.Sprintf("ImportBms_line_%d", idx), bean)
		device_entity, err := deviceDao.GetBySn(logger, bean.Sn)
		if err != nil {
			logger.Warn("deviceDao.GetBySn sql error:", bean.Sn, err.Error())
			device_entity = &deviceDao.Device{
				//Region:     bean.Region,
				//Az:         bean.Az,
				//Sn:         bean.Sn,
				//IloIP:      bean.IloIp,
				//DeviceType: bean.DeviceType,
				//Cabinet:    bean.Cabinet,
				//UPosition:  bean.UPosition,
				//SaleStatus: saleStatus.PUTAWAYING,
				//IsDel:      0,
			}
			toInsert = append(toInsert, device_entity)
		} else {
			if device_entity.ManageStatus != saleStatus.PUTAWAYING {
				logger.Warn("ImportBms sn exist:", device_entity.Sn)
				continue
			}
			//device_entity.Region = bean.Region
			//device_entity.Az = bean.Az
			//device_entity.IloIP = bean.IloIp
			//device_entity.DeviceType = bean.DeviceType
			//device_entity.Cabinet = bean.Cabinet
			//device_entity.UPosition = bean.UPosition
			//device_entity.UPosition = bean.UPosition
			//device_entity.SaleStatus = saleStatus.PUTAWAYING
			//device_entity.IsDel = 0
			toUpdate = append(toUpdate, device_entity)
		}
	}
	if _, err := deviceDao.AddMultiDevice(logger, toInsert); err != nil {
		logger.Warn("ImportBms AddMultiDevice sql error:", err.Error())
		return err
	}
	for _, device_entity := range toUpdate {
		if err := deviceDao.UpdateDeviceById(logger, device_entity); err != nil {
			logger.Warn("ImportBms UpdateDeviceById sql error:", err.Error())
			return err
		}
	}

	if err = saveJDBondInterface(logger, beans); err != nil {
		return err
	}
	if err = saveInterfaces(logger, beans); err != nil {
		return err
	}
	return nil
}
func saveJDBondInterface(logger *log.Logger, beans []*requestTypes.DeviceExcelBean) error {
	//data := []*jdBondInterfaceDao.JdBondInterface{}
	for _, bean := range beans {
		entity := &jdBondInterfaceDao.JdBondInterface{
			Sn:    bean.Sn,
			IsDel: 1,
		}
		err := jdBondInterfaceDao.DeleteJdBondInterfaceBySn(logger, entity)
		if err != nil {
			logger.Warn("delete JdBondInterface err:", bean.Sn, err.Error())
			return err
		}
		entity = &jdBondInterfaceDao.JdBondInterface{
			Sn:        bean.Sn,
			PrivateIP: bean.SystemIp,
			SwitchIP:  bean.SwitchIp,
			IsDel:     0,
		}
		_, err = jdBondInterfaceDao.AddJdBondInterface(logger, entity)
		if err != nil {
			logger.Warn("insert JdBondInterface err:", bean.Sn, err.Error())
			return err
		}
	}
	return nil
}
func saveInterfaces(logger *log.Logger, beans []*requestTypes.DeviceExcelBean) error {
	//data := []*jdBondInterfaceDao.JdBondInterface{}
	for _, bean := range beans {
		entity := &interfaceDao.Interface{
			Sn:    bean.Sn,
			IsDel: 1,
		}
		err := interfaceDao.DeleteInterfaceBySn(logger, entity)
		if err != nil {
			logger.Warn("delete interface err:", bean.Sn, err.Error())
			return err
		}
		entity = &interfaceDao.Interface{
			InterfaceName: "eth0", //默认写死，对应的是mac1
			InterfaceType: "lan",  //写死，以后根据需要再改
			Sn:            bean.Sn,
			Mac:           bean.Mac1,
			SwitchIP:      bean.SwitchIp,
			//SwitchPort: nil,
			IsDel: 0,
		}
		_, err = interfaceDao.AddInterface(logger, entity)
		if err != nil {
			logger.Warn("insert interface err:", bean.Sn, err.Error())
			return err
		}
		entity = &interfaceDao.Interface{
			InterfaceName: "eth1", //默认写死，对应的是mac1
			InterfaceType: "lan",  //写死，以后根据需要再改
			Sn:            bean.Sn,
			Mac:           bean.Mac2,
			SwitchIP:      bean.SwitchIp,
			//SwitchPort: "port2",
			IsDel: 0,
		}
		_, err = interfaceDao.AddInterface(logger, entity)
		if err != nil {
			logger.Warn("insert interface err:", bean.Sn, err.Error())
			return err
		}
	}
	return nil
}
func parseXlsx(logger *log.Logger, file multipart.File, size int64) ([]*requestTypes.DeviceExcelBean, error) {
	xlFile, err := xlsx.OpenReaderAt(file, size)
	if err != nil {
		return nil, err
	}
	beans := []*requestTypes.DeviceExcelBean{}
	for _, sheet := range xlFile.Sheets {
		if sheet.Hidden { // 不读隐藏的工作区
			continue
		}
		for i, row := range sheet.Rows {
			if i == 0 { // 头信息
				continue
			}
			bean := &requestTypes.DeviceExcelBean{}
			if err := row.ReadStruct(bean); err != nil {
				logger.Warn("parseXlsx ReadStruct error:", err.Error())
				return nil, err
			}
			beans = append(beans, bean)
		}
		break //sheet(0)
	}
	return beans, nil
}

// 获取所有符合条件的设备，可以用于查询库存，查询所有设备接口使用
func GetAllDevices(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string) []*deviceDao.Device {
	list, err := deviceDao.GetAllDevices(logger, query, fields, sortby, order)
	if err != nil {
		logger.Warn("GetAllDevices error:", err.Error())
		return nil
	}
	return list
}

func GetStockByDeviceTypeIds(logger *log.Logger, deviceTypeIds []string) (map[string]int, error) {
	query := map[string]interface{}{
		"is_del":            0,
		"device_type_id.in": deviceTypeIds,
		"manage_status":     "putaway",
	}
	list, err := deviceDao.GetAllDevices(logger, query, nil, nil, nil)
	if err != nil {
		logger.Warn("GetAllDevices error:", err.Error())
		return nil, err
	}
	res := map[string]int{}
	for _, v := range list {
		_, ok := res[v.DeviceTypeID]
		if ok {
			res[v.DeviceTypeID] += 1
		} else {
			res[v.DeviceTypeID] = 1
		}
	}
	return res, nil
}

// 查询设备详情单独使用
func DeviceEntity2DeviceOne(logger *log.Logger, d *deviceDao.Device, deviceTypeMapInfo map[string]map[string]string) *responseTypes.Device {

	tz := logger.GetPoint("timezone").(string)
	idc, err := idcDao.GetIdcById(logger, d.IDcID) //获取机房信息
	if err != nil {
		idc = &idcDao.Idc{}
	}
	//idc := idcMap[d.IDcID]
	entity := &deviceTypeDao.DeviceType{}
	if d.DeviceTypeID != "" {
		entity, err = deviceTypeDao.GetOneDeviceType(logger, map[string]interface{}{"device_type_id": d.DeviceTypeID})
		if err != nil {
			entity = &deviceTypeDao.DeviceType{}
		}
	}

	//entity := deviceTypMap[d.DeviceTypeID]
	instance, err := instanceDao.GetInstanceById(logger, d.InstanceID)
	if instance == nil {
		instance = &instanceDao.Instance{}
	}
	//instance := instanceMap[d.InstanceID]
	image, err := imageDao.GetImageByUuid(logger, instance.ImageID)
	if image == nil {
		image = &imageDao.Image{}
	}
	//image := imageMap[instance.ImageID]
	//user, _ := userDao.GetUserById(logger, instance.UserID)
	//if user == nil {
	//	user = &userDao.User{}
	//}
	language := logger.GetPoint("language").(string)
	reason := baseLogic.DeviceReason[d.Reason]
	if language == baseLogic.EN_US {
		reason = baseLogic.DeviceReasonEn[d.Reason]
	}
	instanceStatusName := baseLogic.InstanceStatus[instance.Status]
	if language == baseLogic.EN_US {
		instanceStatusName = baseLogic.InstanceStatusEn[instance.Status]
	}
	instanceReason := baseLogic.InstanceReason[instance.Reason]
	if language == baseLogic.EN_US {
		instanceReason = baseLogic.InstanceReasonEn[instance.Reason]
	}
	deviceManageStatusName := baseLogic.DeviceManageStatus[d.ManageStatus]
	if language == baseLogic.EN_US {
		deviceManageStatusName = baseLogic.DeviceManageStatusEn[d.ManageStatus]
	}
	deviceSeriesName := baseLogic.DeviceTypeSeries[d.DeviceSeries]
	if language == baseLogic.EN_US {
		deviceSeriesName = baseLogic.DeviceTypeSeriesEn[d.DeviceSeries]
	}
	idcName := idc.Name
	if language == baseLogic.EN_US {
		idcName = idc.NameEn
	}
	var (
		iloUser         string
		iloPassword     string
		switchUser1     string
		switchUser2     string
		switchPassword1 string
		switchPassword2 string
		encloure1       string
		slot1           int
		encloure2       string
		slot2           int
	)
	// disk, _ := diskDao.GetAllDisk(logger, map[string]interface{}{"device_id": d.DeviceID})
	// if len(disk) == 1 {
	// 	encloure1 = disk[0].Enclosure
	// 	slot1 = disk[0].Slot
	// } else if len(disk) == 2 {
	// 	encloure1 = disk[0].Enclosure
	// 	slot1 = disk[0].Slot
	// 	encloure2 = disk[1].Enclosure
	// 	slot2 = disk[1].Slot
	// }

	if d.IloUser == "" {
		iloUser = idc.IloUser
	} else {
		iloUser = d.IloUser
	}
	if d.IloPassword == "" {
		iloPassword = idc.IloPassword
	} else {
		iloPassword = d.IloPassword
	}
	if d.SwitchUser1 == "" {
		switchUser1 = idc.SwitchUser1
	} else {
		switchUser1 = d.SwitchUser1
	}
	if d.SwitchPassword1 == "" {
		switchPassword1 = idc.SwitchPassword1
	} else {
		switchPassword1 = d.SwitchPassword1
	}
	if d.SwitchUser2 == "" {
		switchUser2 = idc.SwitchUser2
	} else {
		switchUser2 = d.SwitchUser2
	}
	if d.SwitchPassword2 == "" {
		switchPassword2 = idc.SwitchPassword2
	} else {
		switchPassword2 = d.SwitchPassword2
	}
	collectStatus := "2" //2表示未采集
	if d.CollectStatus != "" {
		collectStatus = d.CollectStatus
	}

	r := &responseTypes.Device{
		ID:                  d.ID,
		DeviceID:            d.DeviceID,
		IdcID:               d.IDcID,
		InstanceID:          d.InstanceID,
		InstanceName:        instance.InstanceName,
		InstanceStatus:      instance.Status,
		InstanceStatusName:  instanceStatusName,
		InstanceReason:      instanceReason,
		Locked:              instance.Locked,
		UserId:              d.UserId,
		UserName:            d.UserName,
		InstanceCreatedTime: util.TimestampToString(int64(instance.CreatedTime), tz),
		InstanceDescription: instance.Description,

		DeviceTypeID:     d.DeviceTypeID,
		ManageStatus:     d.ManageStatus,
		ManageStatusName: deviceManageStatusName,

		ImageName: image.ImageName,

		Brand:       d.Brand,
		Model:       d.Model,
		Reason:      reason,
		Cabinet:     d.Cabinet,
		UPosition:   d.UPosition,
		Sn:          d.Sn,
		Description: d.Description,

		Mac1:        d.Mac1,
		Mac2:        d.Mac2,
		SwitchIP1:   d.SwitchIP1,
		SwitchPort1: d.SwitchPort1,
		SwitchIP2:   d.SwitchIP2,
		SwitchPort2: d.SwitchPort2,
		//如果设备不存在，从idc中获取
		IloUser:         iloUser,
		IloPassword:     iloPassword,
		IloIP:           d.IloIP, //必填
		SwitchUser1:     switchUser1,
		SwitchPassword1: switchPassword1,
		SwitchUser2:     switchUser2,
		SwitchPassword2: switchPassword2,
		SwitchIP:        d.SwitchIP,

		Mask:            d.Mask,
		Eth1Mask:        d.MaskEth1,
		Gateway:         d.Gateway,
		PrivateIPv4:     d.PrivateIPv4,
		PrivateEth1IPv4: d.PrivateEth1IPv4,
		PrivateIPv6:     d.PrivateIPv6,
		PrivateEth1IPv6: d.PrivateEth1IPv6,
		AdapterID:       d.AdapterID,

		Enclosure1:  encloure1,
		Slot1:       slot1,
		Enclosure2:  encloure2,
		Slot2:       slot2,
		RaidDriver:  d.RaidDriver,
		CreatedBy:   d.CreatedBy,
		UpdatedBy:   d.UpdatedBy,
		CreatedTime: util.TimestampToString(int64(d.CreatedTime), tz),
		UpdatedTime: util.TimestampToString(int64(d.UpdatedTime), tz),

		//以下字段从其他表中获取
		//idc
		IdcName:   idcName,
		IDcNameEn: idc.NameEn,
		//deviceType

		DeviceType:       entity.DeviceType,
		DeviceSeries:     d.DeviceSeries,
		DeviceSeriesName: deviceSeriesName,

		DeviceTypeName:  entity.Name,
		Architecture:    entity.Architecture,
		CPUAmount:       entity.CPUAmount,
		CPUCores:        entity.CPUCores,
		CPUFrequency:    entity.CPUFrequency,
		CPUManufacturer: entity.CPUManufacturer,
		CPUModel:        entity.CPUModel,

		MemType:                   entity.MemType,
		MemAmount:                 entity.MemAmount,
		MemSize:                   entity.MemSize,
		MemFrequency:              entity.MemFrequency,
		NicAmount:                 entity.NicAmount,
		NicRate:                   entity.NicRate,
		InterfaceMode:             entity.InterfaceMode,
		SystemVolumeType:          entity.SystemVolumeType,
		SystemVolumeInterfaceType: entity.DataVolumeInterfaceType,
		SystemVolumeAmount:        entity.SystemVolumeAmount,
		SystemVolumeSize:          entity.SystemVolumeSize,
		DataVolumeType:            entity.DataVolumeType,
		DataVolumeInterfaceType:   entity.SystemVolumeInterfaceType,
		DataVolumeAmount:          entity.DataVolumeAmount,
		DataVolumeSize:            entity.DataVolumeSize,
		GpuAmount:                 entity.GpuAmount,
		GpuManufacturer:           entity.GpuManufacturer,
		GpuModel:                  entity.GpuModel,

		CollectStatus:     collectStatus,
		CollectFailReason: d.CollectFailReason,
		IsNeedRaid:        entity.IsNeedRaid,
	}
	//拼装信息
	if _, ok := deviceTypeMapInfo[entity.DeviceTypeID]; ok {
		r.CpuInfo = deviceTypeMapInfo[entity.DeviceTypeID]["cpu_"]
		r.MemInfo = deviceTypeMapInfo[entity.DeviceTypeID]["mem_"]
		r.GpuInfo = deviceTypeMapInfo[entity.DeviceTypeID]["gpu_"]
		r.NicInfo = deviceTypeMapInfo[entity.DeviceTypeID]["nic_"]
	}

	//SvInfo:            deviceTypeMapInfo[entity.DeviceTypeID]["sv_"],
	//DvInfo:            deviceTypeMapInfo[entity.DeviceTypeID]["dv_"],

	return r
}

func GetGlobalDeviceCount(logger *log.Logger) (int64, error) {
	q := map[string]interface{}{
		"is_del": 0,
	}
	return deviceDao.GetDeviceCount(logger, q)
}

// 设备详情-磁盘
func GetDeviceDisksById(logger *log.Logger, DeviceID string) (*responseTypes.DeviceDisks, error) {
	deviceEntity, err := deviceDao.GetDeviceById(logger, DeviceID)
	if err != nil {
		logger.Warn("GetById GetDeviceById sql error:", DeviceID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	res := &responseTypes.DeviceDisks{
		Disks:   []*diskDao.Disk{},
		Panfu:   []*diskDao.Disk{},
		Volumes: []*responseTypes.VolumeIt{},
	}

	dq := map[string]interface{}{
		"is_del":    0,
		"device_id": DeviceID,
	}
	disks, err := diskDao.GetAllDisk(logger, dq)
	if err != nil {
		logger.Warnf("GetDeviceDisksById.GetAllDisk error, device_id:%s, error:%s", DeviceID, err.Error())
		return nil, err
	}
	for _, v := range disks {
		if v.Types == "nvme" || v.Types == "controller" {
			res.Disks = append(res.Disks, v)
		} else if v.Types == "panfu" {
			res.Panfu = append(res.Panfu, v)
		}
	}

	if deviceEntity.DeviceTypeID != "" {
		q1 := map[string]interface{}{
			"is_del":         0,
			"device_type_id": deviceEntity.DeviceTypeID,
		}
		volumeEntities, err := volumeDao.GetAllVolume(logger, q1) //device_type_id如果为空，where就忽略了key
		if err != nil {
			logger.Warnf("GetDeviceDisksById.GetAllVolume error, device_type_id:%s, error:%s", deviceEntity.DeviceTypeID, err.Error())
		}
		volume := responseTypes.Volume{}
		for _, v := range volumeEntities {
			volume.VolumeID = v.VolumeID
			volume.VolumeName = v.VolumeName
			volume.VolumeType = v.VolumeType
			volume.VolumeSize = v.VolumeSize
			volume.VolumeUnit = v.VolumeUnit

			volume.VolumeAmount = v.VolumeAmount
			volume.DiskType = v.DiskType
			volume.InterfaceType = v.InterfaceType

			volumeId := v.VolumeID
			q2 := map[string]interface{}{
				"is_del":    0,
				"volume_id": v.VolumeID,
			}
			volumnRaidEntities, err := rVolumeRaidDao.GetAllRVolumeRaid(logger, q2)
			if err != nil {
				logger.Warnf("GetDeviceDisksById.GetVolumeById error, volume_id:%s, error:%s", volumeId, err.Error())
			}
			volume.RaidCan = volumnRaidEntities[0].RaidCan
			raids := []*raidDao.Raid{}
			for _, v2 := range volumnRaidEntities {
				raids = append(raids, &raidDao.Raid{
					RaidID: v2.RaidID,
					Name:   v2.RaidName,
				})
			}

			diskEntities, err := diskDao.GetDiskWithDeviceIdVolumeId(logger, DeviceID, volumeId)
			if err != nil {
				logger.Warnf("GetDeviceDisksById.GetVolumeById error, volume_id:%s, error:%s", volumeId, err.Error())
			}

			vt := &responseTypes.VolumeIt{
				Volume: volume,
				Raids:  raids,
				Disks:  diskEntities,
			}
			res.Volumes = append(res.Volumes, vt)
		}
	}

	return res, nil
}

func AssociateDeviceDisks(logger *log.Logger, param *requestTypes.AssociateDeviceDisksRequest) (bool, error) {

	deviceEntity, err := deviceDao.GetOneDevice(logger, map[string]interface{}{
		"device_id": param.DeviceID,
		"is_del":    baseLogic.IS_NOT_DEL,
	})
	if err != nil {
		logger.Warnf("AssociateDeviceDisks.GetOneDevice sql error, device_id:%s, error:%s", param.DeviceID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("deviceId %s 校验非法", param.DeviceID), fmt.Sprintf("deviceId %s check invalided", param.DeviceID)))
	}

	deviceTypeEntity, err := deviceTypeDao.GetDeviceTypeById(logger, param.DeviceTypeID)
	if err != nil {
		logger.Warnf("GetAssociatedDisks.GetDeviceTypeById sql error, device_type_id:%s, error:%s", param.DeviceTypeID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("device_type_Id %s 校验非法", param.DeviceTypeID), fmt.Sprintf("device_type_Id %s check invalided", param.DeviceTypeID)))
	}

	//manageStatus := param.ManageStatus
	// if deviceEntity.ManageStatus != baseLogic.IN && deviceEntity.ManageStatus != baseLogic.PUTAWAYFAIL {
	// 	panic(constant.BuildInvalidArgumentWithArgs("设备必须是已入库或者上架失败，才可以执行预上架操作", "device manageStatus not support"))
	// }

	//验证设备的多个卷选择的diskid是否有重复
	diskIdCheck := map[string]int{}

	rEntities := []*rDeviceVolumeDisksDao.RDeviceVolumeDisks{}
	fmt.Println(param.Volumes)

	//校验volumeid，diskid，deviceid的归属关系是否正确
	for _, v := range param.Volumes {
		fmt.Println("sssss", v.VolumeID)
		volumeId := v.VolumeID

		volumeEntity, err := volumeDao.GetVolumeById(logger, volumeId)
		if err != nil {
			logger.Warnf("param volumeId %s check invalid, error:%s", volumeId, err.Error())
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("volumeId %s 校验非法", volumeId), fmt.Sprintf("volumeId %s check invalided", volumeId)))
		}
		if volumeEntity.DeviceTypeID != param.DeviceTypeID {
			logger.Warnf("param volumeId %s and devicetype_id %s not match", volumeId, param.DeviceTypeID)
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("volumeId %s 和 devicetype %s 校验非法", volumeId, param.DeviceTypeID), fmt.Sprintf("volumeId %s and devicetype %s check invalid", volumeId, param.DeviceTypeID)))
		}
		rVolumeRaidEntities, err := rVolumeRaidDao.GetRVolumeRaidsByVId(logger, volumeEntity.VolumeID)
		logger.Infof("AssociateDeviceDisks volume_id:%s GetRVolumeRaidsByVId res:%s", volumeEntity.VolumeID, util.InterfaceToJson(rVolumeRaidEntities))
		var raidCan string
		mindisks := 1
		maxdisks := math.MaxInt
		fmt.Println("长度", len(rVolumeRaidEntities))
		if len(rVolumeRaidEntities) == 0 { //没有raid卡，使用盘符的情况
			mindisks = 1
			maxdisks = 1
			fmt.Println("", volumeEntity.VolumeID)
		} else {
			raidCan = rVolumeRaidEntities[0].RaidCan
			fmt.Println(raidCan, 666)
			if util.InArray(raidCan, []string{baseLogic.RAID_CAN_NO_RAID, baseLogic.RAID_CAN_SINGLE_RAID, baseLogic.RAID_CAN_SINGLE_RAID_EN}) { //nvme的情况
				mindisks = 1
				maxdisks = 1
			} else {
				for _, rVolumeRaidEntity := range rVolumeRaidEntities {
					fmt.Println("raidname", rVolumeRaidEntity.RaidName)
					if util.InArray(rVolumeRaidEntity.RaidName, []string{"RAID0", "RAID1"}) {
						mindisks = util.Max(mindisks, 2)
					} else if rVolumeRaidEntity.RaidName == "RAID5" {
						mindisks = util.Max(mindisks, 3)
					} else if rVolumeRaidEntity.RaidName == "RAID10" {
						mindisks = util.Max(mindisks, 4)
					}
				}
			}
		}
		if mindisks < volumeEntity.VolumeAmount {
			mindisks = volumeEntity.VolumeAmount
		}
		logger.Infof("volume:%s require min:%d, max:%d", volumeEntity.VolumeID, mindisks, maxdisks)

		var types string = "controller" //默认controller
		//无需配置raid的机型，只能选盘符。需要配置RAID的，根据以下逻辑判断
		if deviceTypeEntity.IsNeedRaid != "no_need_raid" { // no_need_raid
			//TODO 这个规则需确定
			//raidCan只存在了这个表，每个volumeID下面的raid配置肯定都是一样的，一个voulumeId下面对应多个raid，但是只有一个共同的raid模式
			if raidCan == baseLogic.RAID_CAN_SINGLE_RAID_EN || raidCan == baseLogic.RAID_CAN_RAID { //如果是单盘RAID0或者是RAID
				types = "controller"
			} else if raidCan == baseLogic.RAID_CAN_NO_RAID {
				types = "nvme"
			}
		} else {
			types = "panfu"
		}

		diskIds := v.DiskIDs
		fmt.Println(diskIds, mindisks, maxdisks)
		if len(diskIds) < mindisks { //至少需要磁盘的块数校验
			logger.Warnf("param volumeId %s require at least %d disks", volumeId, mindisks)
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("卷名称：%s,最少%d块盘，数量是所选RAID需要的最大盘数", volumeEntity.VolumeName, mindisks), fmt.Sprintf("volumeName:%s,At least %d disks", volumeEntity.VolumeName, mindisks)))
		}
		if len(diskIds) > maxdisks { //至多需要磁盘的块数校验
			logger.Warnf("param volumeId %s require at most %d disks", volumeId, maxdisks)
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("卷名称：%s,仅支持1块盘", volumeEntity.VolumeName), fmt.Sprintf("volumeName:%s,Only support 1 disk", volumeEntity.VolumeName)))
		}

		for _, diskId := range diskIds {
			diskIdCheck[diskId] += 1
			if diskIdCheck[diskId] > 1 { //设备的多个卷选择的diskid有重复
				logger.Warnf("param diskid  %s repeated, device_id:%s", diskId, param.DeviceID)
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("param diskid  %s repeated", diskId), fmt.Sprintf("param diskid  %s repeated", diskId)))
			}

			diskEntity, err := diskDao.GetDiskByUuid(logger, diskId)
			if err != nil {
				logger.Warnf("param diskId %s check invalid, error:%s", diskId, err.Error())
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("diskId %s 校验非法", diskId), fmt.Sprintf("diskId %s check invalided", diskId)))
			}
			if diskEntity.DeviceID != param.DeviceID {
				logger.Warnf("param diskId and device not match, %s error:%s", diskId, err.Error())
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("diskId %s 和 device %s 校验非法", diskId, param.DeviceID), fmt.Sprintf("diskId %s and device %s check invalid", diskId, param.DeviceID)))
			}
			if diskEntity.Types != types { //磁盘类型(controllers/nvme/panfu)的校验
				logger.Warnf("param disk types invalid, volume_id:%s, disk_id:%s, disk type:%s, require type:%s", volumeId, diskId, diskEntity.Types, types)
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("diskId %s 类型校验非法", diskId), fmt.Sprintf("diskId %s types check invalid", diskId)))
			}

			//无需配置raid的机型，只能选非盘符的。1->need_raid;2->no_need_raid
			if (deviceTypeEntity.IsNeedRaid == "no_need_raid" && diskEntity.Types != "panfu") || (deviceTypeEntity.IsNeedRaid != "no_need_raid" && diskEntity.Types == "panfu") {
				logger.Warnf("param diskId not match devicetype, disk_id:%s, device_type_id:%s", diskId, deviceEntity.DeviceTypeID)
				panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("param disk not match devicetype, disk_id:%s, device_type_id:%s", diskId, deviceEntity.DeviceTypeID), fmt.Sprintf("param diskId not match devicetype, disk_id:%s, device_type_id:%s", diskId, deviceEntity.DeviceTypeID)))
			}

		}
	}

	//删除原来绑定的disk
	q := map[string]interface{}{
		"is_del":    0,
		"device_id": param.DeviceID,
	}
	u := map[string]interface{}{
		"is_del": 1,
	}
	if err := rDeviceVolumeDisksDao.UpdateMultiRDeviceVolumeDisk(logger, q, u); err != nil {
		logger.Warnf("AssociateDeviceDisks.delRDeviceVolumeDisk error, device_id:%s, error:%s", param.DeviceID, err.Error())
	} else {
		logger.Infof("AssociateDeviceDisks.delMultiRDeviceVolumeDisk succ, device_id:%s", param.DeviceID)
	}

	//更改设备的机型
	deviceEntity.DeviceTypeID = param.DeviceTypeID
	if err := deviceDao.UpdateDeviceById(logger, deviceEntity); err != nil {
		logger.Warnf("associateDisk.UpdateDeviceById error, device_id:%s, device_type_id:%s, error:%s", param.DeviceID, param.DeviceTypeID, err.Error())
	} else {
		logger.Infof("AssociateDeviceDisks.UpdateDeviceTypeById succ, device_id:%s, device_type_id:%s", param.DeviceID, param.DeviceTypeID)
	}

	for _, v := range param.Volumes {
		volumeId := v.VolumeID

		volumeEntity, _ := volumeDao.GetVolumeById(logger, volumeId)

		diskIds := v.DiskIDs

		diskEntities := []*diskDao.Disk{}
		for _, diskId := range diskIds {
			diskEntity, _ := diskDao.GetDiskByUuid(logger, diskId)
			diskEntities = append(diskEntities, diskEntity)

		}
		if len(diskEntities) == 0 {
			logger.Warnf("disk not selected, sn:%s, volume_id:%s", deviceEntity.Sn, v.VolumeID)
			continue
		}

		if diskEntities[0].Types == "panfu" || diskEntities[0].Types == "nvme" { //绑定盘符或者nvme盘时，新建或更新device_hints表

			deviceHintsEntity, _ := deviceHintsDao.GetDeviceHintsBySnAndVolumeId(logger, deviceEntity.Sn, v.VolumeID)
			size, _ := strconv.ParseInt(diskEntities[0].Size, 10, 64)
			if deviceHintsEntity == nil { //新建

				entity := &deviceHintsDao.DeviceHints{
					Sn:       deviceEntity.Sn,
					VolumeId: v.VolumeID,
					Size:     size,
					Serial:   diskEntities[0].SerialNumber,
					ByPath:   diskEntities[0].DevicePath,
				}

				entity.VolumeType = volumeEntity.VolumeType

				if _, err := deviceHintsDao.AddDeviceHints(logger, entity); err != nil {
					logger.Warnf("AssociateDeviceDisks.saveDeviceHints error, sn:%s, error:%s", deviceEntity.Sn, err.Error())
				}
			} else { //更新
				deviceHintsEntity.Size = size
				deviceHintsEntity.Serial = diskEntities[0].SerialNumber
				deviceHintsEntity.ByPath = diskEntities[0].DevicePath
				if err := deviceHintsDao.UpdateDeviceHintsById(logger, deviceHintsEntity); err != nil {
					logger.Warnf("AssociateDeviceDisks.UpdateDeviceHintsById error, sn:%s, error:%s", deviceEntity.Sn, err.Error())
				}
			}
		}

		for _, diskId := range v.DiskIDs {
			r := &rDeviceVolumeDisksDao.RDeviceVolumeDisks{
				DeviceID:    param.DeviceID,
				VolumeID:    v.VolumeID,
				VolumeType:  volumeEntity.VolumeType,
				DiskID:      diskId,
				CreatedBy:   logger.GetPoint("username").(string),
				UpdatedBy:   logger.GetPoint("username").(string),
				CreatedTime: int(time.Now().Unix()),
				UpdatedTime: int(time.Now().Unix()),
			}
			rEntities = append(rEntities, r)
		}
	}

	if len(rEntities) > 0 {
		if rDeviceVolumeDisksDao.AddMultiRDeviceVolumeDisk(logger, rEntities); err != nil {
			logger.Warnf("AssociateDeviceDisks.AddMultiRDeviceVolumeDisk error, device_id:%s, error:%s", param.DeviceID, err.Error())
			return false, err
		}
	}

	return true, nil
}

func GetAssociatedDisks(logger *log.Logger, param *requestTypes.GetAssociatedDisksRequest) ([]*response.Disk, error) {
	volumeId := param.VolumeID
	_, err := deviceDao.GetOneDevice(logger, map[string]interface{}{
		"device_id": param.DeviceID,
		"is_del":    baseLogic.IS_NOT_DEL,
	})
	if err != nil {
		logger.Warnf("GetAssociatedDisks.GetOneDevice sql error, device_id:%s, error:%s", param.DeviceID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("deviceId %s 校验非法", param.DeviceID), fmt.Sprintf("deviceId %s check invalided", param.DeviceID)))
	}

	deviceTypeEntity, err := deviceTypeDao.GetDeviceTypeById(logger, param.DeviceTypeID)
	if err != nil {
		logger.Warnf("GetAssociatedDisks.GetDeviceTypeById sql error, device_type_id:%s, error:%s", param.DeviceTypeID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("device_type_Id %s 校验非法", param.DeviceTypeID), fmt.Sprintf("device_type_Id %s check invalided", param.DeviceTypeID)))
	}

	volumeEntity, err := volumeDao.GetVolumeById(logger, volumeId)
	if err != nil {
		logger.Warnf("GetAssociatedDisks.param volumeId %s check invalid, error:%s", volumeId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("volumeId %s 校验非法", volumeId), fmt.Sprintf("volumeId %s check invalided", volumeId)))
	}
	if volumeEntity.DeviceTypeID != param.DeviceTypeID {
		logger.Warnf("GetAssociatedDisks.param volumeId and device not match, %s error:%s", volumeId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("volumeId %s 和 deviceId %s 校验非法", volumeId, param.DeviceID), fmt.Sprintf("volumeId %s and deviceId %s check invalid", volumeId, param.DeviceID)))
	}

	types := []string{}
	rVolumeRaidEntities, err := rVolumeRaidDao.GetRVolumeRaidsByVId(logger, param.VolumeID)
	if err != nil {
		logger.Warnf("GetAssociatedDisks.GetRVolumeRaidById error, volume_id:%s, error:%s", param.VolumeID, err.Error())
	}

	//无需配置raid的机型，只能选盘符。需要配置RAID的，根据以下逻辑判断
	if deviceTypeEntity.IsNeedRaid != "no_need_raid" { // no_need_raid
		//TODO 这个规则需确定
		raidCan := rVolumeRaidEntities[0].RaidCan                                               //raidCan只存在了这个表，每个volumeID下面的raid配置肯定都是一样的，一个voulumeId下面对应多个raid，但是只有一个共同的raid模式
		if raidCan == baseLogic.RAID_CAN_SINGLE_RAID_EN || raidCan == baseLogic.RAID_CAN_RAID { //如果是单盘RAID0或者是RAID
			types = []string{"controller"}
		} else if raidCan == baseLogic.RAID_CAN_NO_RAID {
			types = []string{"nvme"}
		}
	} else {
		types = []string{"panfu"}
	}
	if len(types) == 0 { //默认值
		types = []string{"controller"}
	}
	where := map[string]interface{}{
		"is_del":    0,
		"device_id": param.DeviceID,
		// "volume_size.gte": volumeEntity.VolumeSize,
		"types.in": types,
	}
	if volumeEntity.InterfaceType != baseLogic.NOT_LIMITED {
		where["pd_type"] = strings.ToUpper(volumeEntity.InterfaceType)
	}
	if types[0] == "panfu" { //lsblk命令没有pd_type
		delete(where, "pd_type")
	}
	if volumeEntity.DiskType != baseLogic.NOT_LIMITED {
		where["media_type"] = strings.ToUpper(volumeEntity.DiskType)
	}
	allDiskEntities, err := diskDao.GetAllDisk(logger, where)
	if err != nil {
		logger.Warnf("GetAssociatedDisks.GetAllDisk error, device_id:%s, types:%v, error:%s", param.DeviceID, types, err.Error())
	}

	diskEntities, err := diskDao.GetDiskWithDeviceIdVolumeId(logger, param.DeviceID, volumeId)
	if err != nil {
		logger.Warnf("GetAssociatedDisks.GetDiskWithVolumeId error, device_id:%s, error:%s", param.DeviceID, err.Error())
	}
	diskids := []string{}
	for _, v := range diskEntities {
		diskids = append(diskids, v.DiskId)
	}

	res := []*response.Disk{}
	for _, v := range allDiskEntities {
		//以GB来比较
		diskSize, _ := strconv.ParseFloat(strings.TrimSpace(v.Size), 64)
		volumeSize := volumeEntity.VolumeSize
		volumeSize64 := .0
		if v.SizeUnit == "TB" && volumeEntity.VolumeUnit == "GB" { //TB的换算成GB
			diskSize = 1024 * diskSize
		} else if v.SizeUnit == "GB" && volumeEntity.VolumeUnit == "TB" {
			// ParseFloat 将字符串转换为浮点数
			// str：要转换的字符串
			// bitSize：指定浮点类型（32:float32、64:float64）
			// 如果 str 是合法的格式，而且接近一个浮点值，
			// 则返回浮点数的四舍五入值（依据 IEEE754 的四舍五入标准）
			// 如果 str 不是合法的格式，则返回“语法错误”
			// 如果转换结果超出 bitSize 范围，则返回“超出范围”
			//到float64
			volumeSize64, _ := strconv.ParseFloat(volumeSize, 64)
			volumeSize64 = 1024 * volumeSize64
		}

		if diskSize < volumeSize64 {
			continue
		}

		res = append(res, &response.Disk{
			Disk:     *v,
			Selected: util.InArray(v.DiskId, diskids),
		})
	}

	return res, nil
}

func DeviceAssociateDeviceType(logger *log.Logger, param *requestTypes.DeviceAssociateDeviceTypeRequest) (bool, error) {

	deviceEntity, err := deviceDao.GetOneDevice(logger, map[string]interface{}{
		"device_id": param.DeviceID,
		"is_del":    baseLogic.IS_NOT_DEL,
	})
	if err != nil {
		logger.Warnf("DeviceAssociateDeviceType.GetOneDevice sql error, device_id:%s, error:%s", param.DeviceID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("deviceId %s 校验非法", param.DeviceID), fmt.Sprintf("deviceId %s check invalided", param.DeviceID)))
	}

	//if deviceEntity.InstanceID != "" {
	//	logger.Warnf("DeviceAssociateDeviceType device_id %s with instance_id %s not empty, operation denied ", param.DeviceID, deviceEntity.InstanceID)
	//	panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("AssociateDeviceDisks device_id %s with instance_id %s not empty, operation denied ", param.DeviceID, deviceEntity.InstanceID), fmt.Sprintf("DeviceAssociateDeviceType device_id %s with instance_id %s not empty, operation denied ", param.DeviceID, deviceEntity.InstanceID)))
	//}
	if deviceEntity.ManageStatus == "created" {
		logger.Warnf("DeviceAssociateDeviceType device_id %s has already sold", param.DeviceID)
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("AssociateDeviceDisks device_id %s 已售卖", param.DeviceID), fmt.Sprintf("DeviceAssociateDeviceType device_id %s has already sold", param.DeviceID)))
	}
	if deviceEntity.DeviceTypeID != param.DeviceTypeID { //设备更换了机型，则r_device_volume_disk表相关数据要清空
		deviceTypeEntity, err := deviceTypeDao.GetDeviceTypeById(logger, param.DeviceTypeID)
		if err != nil {
			logger.Warnf("DeviceAssociateDeviceType.GetDeviceTypeById sql error, device_type_id:%s, error:%s", param.DeviceTypeID, err.Error())
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("device_type_Id %s 校验非法", param.DeviceTypeID), fmt.Sprintf("device_type_Id %s check invalided", param.DeviceTypeID)))
		}

		deviceEntity.DeviceTypeID = deviceTypeEntity.DeviceTypeID
		deviceEntity.DeviceSeries = deviceTypeEntity.DeviceSeries //同步一下机型的机型类型字段到device表，目的是冗余这个字段，便于搜索
		if err := deviceDao.UpdateDeviceById(logger, deviceEntity); err != nil {
			logger.Warnf("DeviceAssociateDeviceType.UpdateDeviceById error, device_id:%s, device_type_id:%s, error:%s", param.DeviceID, param.DeviceTypeID, err.Error())
			return false, err
		}

		/*
			前端的绑定磁盘和绑定机型请求过来的顺序不固定，所以这里不能这样删除。。。
			q := map[string]interface{}{
				"is_del":    0,
				"device_id": param.DeviceID,
			}
			u := map[string]interface{}{
				"is_del": 1,
			}
			if err := rDeviceVolumeDisksDao.UpdateMultiRDeviceVolumeDisk(logger, q, u); err != nil {
				logger.Warnf("DeviceAssociateDeviceType.UpdateMultiRDeviceVolumeDisk error, device_id:%s, error:%s", param.DeviceID, err.Error())
			}
		*/
	}

	return true, nil
}
