package deviceLogic

import (
	"strconv"
	"strings"
	"sync"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/imageDao"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceTypeDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/diskDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/idcDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	namespacePrefix "git.jd.com/cps-golang/ironic-common/ironic/common/NamespacePrefixs"

	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"

	deviceDao "coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceDao"
	idcApi "coding.jd.com/aidc-bmp/bmp-openapi/service/idc_api"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type CreateDeviceMq struct {
	Action     string `json:"action"`
	Sn         string `json:"sn"`
	Subnet     string `json:"subnet"`
	SubnetMask string `json:"subnet_mask"`
	Routes     string `json:"routes"`
}

//检测设备状态，给driver发消息 ，验证带外等信息是否正确
type CheckDeviceMq struct {
	Action    string `json:"action"`
	Sn        string `json:"sn"`
	IloIp     string `json:"ilo_ip"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Privilege string `json:"privilege"`
}

var wg sync.WaitGroup

func buildDeviceEntity(logger *log.Logger, c *requestTypes.CreateDeviceSpec, deviceType *deviceTypeDao.DeviceType) *deviceDao.Device {
	ipv6, _ := beego.AppConfig.String("gateway6")
	return &deviceDao.Device{
		DeviceID:     commonUtil.GetRandString("d-", namespacePrefix.IMAGE_ID_LENGTH, false, true, true),
		InstanceID:   "",
		IDcID:        deviceType.IDcID,
		DeviceTypeID: deviceType.DeviceTypeID,
		ManageStatus: baseLogic.IN, //导入设备以后，默认已入库

		Brand:        c.Brand,
		Model:        c.Model,
		DeviceSeries: deviceType.DeviceSeries,
		Cabinet:      c.Cabinet,
		UPosition:    c.UPosition,
		Sn:           c.Sn,
		IloUser:      c.IloUser,
		IloPassword:  c.IloPassword,
		IloIP:        c.IloIP,
		Description:  c.Description,

		Mac1:        c.Mac1,
		Mac2:        c.Mac2,
		SwitchIP1:   c.SwitchIP1,
		SwitchPort1: c.SwitchPort1,
		SwitchIP2:   c.SwitchIP2,
		SwitchPort2: c.SwitchPort2,

		SwitchUser1:     c.SwitchUser1,
		SwitchPassword1: c.SwitchPassword1,
		SwitchUser2:     c.SwitchUser2,
		SwitchPassword2: c.SwitchPassword2,
		SwitchIP:        c.SwitchIP,

		Mask:        c.Mask,
		Gateway:     c.Gateway,
		PrivateIPv4: c.PrivateIPv4,
		PrivateIPv6: c.PrivateIPv6,
		Gateway6:    ipv6,
		AdapterID:   *c.AdapterID,

		RaidDriver:  c.RaidDriver,
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
	//disk, _ := diskDao.GetAllDisk(logger, map[string]interface{}{"device_id": d.DeviceID})
	//if len(disk) == 1 {
	//	encloure1 = disk[0].Enclosure
	//	slot1 = disk[0].Slot
	//} else if len(disk) == 2 {
	//	encloure1 = disk[0].Enclosure
	//	slot1 = disk[0].Slot
	//	encloure2 = disk[1].Enclosure
	//	slot2 = disk[1].Slot
	//}
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

		Mask:        d.Mask,
		Gateway:     d.Gateway,
		PrivateIPv4: d.PrivateIPv4,
		PrivateIPv6: d.PrivateIPv6,
		AdapterID:   d.AdapterID,

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
	sv_ := strconv.Itoa(int(v.SystemVolumeAmount)*int(v.SystemVolumeSize)) + "GB(" + strconv.Itoa(int(v.SystemVolumeSize)) + "GB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount)) + ")"
	if v.SystemVolumeUnit == "TB" {
		sv_ = Trim(strconv.FormatFloat(float64(v.SystemVolumeAmount)*float64(v.SystemVolumeSize)/1000, 'f', 2, 64)) + "TB(" + Trim(strconv.FormatFloat(float64(v.SystemVolumeSize)/1000, 'f', 2, 64)) + "TB " + v.SystemVolumeInterfaceType + " " + v.SystemVolumeType + "*" + strconv.Itoa(int(v.SystemVolumeAmount)) + ")"
	}
	// 数据盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式)
	dv_ := strconv.Itoa(int(v.DataVolumeAmount)*int(v.DataVolumeSize)) + "GB(" + strconv.Itoa(int(v.DataVolumeSize)) + "GB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount)) + ")"
	if v.DataVolumeUnit == "TB" {
		dv_ = Trim(strconv.FormatFloat(float64(v.DataVolumeAmount)*float64(v.DataVolumeSize)/1000, 'f', 2, 64)) + "TB(" + Trim(strconv.FormatFloat(float64(v.DataVolumeSize)/1000, 'f', 2, 64)) + "TB " + v.DataVolumeInterfaceType + " " + v.DataVolumeType + "*" + strconv.Itoa(int(v.DataVolumeAmount)) + ")" //  + "," + deviceTypeRaid + ")"
	}
	if v.DataVolumeAmount == 0 {
		dv_ = ""
	}
	// 显示品牌 显卡型号 个数，例 NVIDIA P40 * 4
	gpu_ := v.GpuManufacturer + " " + v.GpuModel + "*" + strconv.Itoa(int(v.GpuAmount))
	if v.GpuAmount == 0 {
		gpu_ = ""
	}
	nic_ := strconv.Itoa(int(v.NicAmount)) + "*" + strconv.Itoa(int(v.NicRate)) + "GE"
	return map[string]string{
		"cpu_": cpu_,
		"mem_": mem_,
		"sv_":  sv_,
		"dv_":  dv_,
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

	for _, v := range deviceTypeALl {

		info := GetDeviceTypeInfo(logger, v)
		deviceTypeMap[v.DeviceTypeID] = *v
		//fmt.Println(info, util.ObjToJson(info), info["nic_"])
		//拼装数据，给上层api直接使用
		deviceTypeMapInfo[v.DeviceTypeID] = map[string]string{
			"cpu_": info["cpu_"],
			"mem_": info["mem_"],
			"sv_":  info["sv_"],
			"dv_":  info["dv_"],
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

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
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

//设备详情
func GetById(logger *log.Logger, DeviceID string) (*responseTypes.Device, error) {
	entity, err := deviceDao.GetDeviceById(logger, DeviceID)
	if err != nil {
		logger.Warn("GetById GetDeviceById sql error:", DeviceID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	v, err := deviceTypeDao.GetDeviceTypeById(logger, entity.DeviceTypeID)
	if err != nil {
		v = &deviceTypeDao.DeviceType{}
	}
	//目的是获取拼装出来的机型数据
	info := GetDeviceTypeInfo(logger, v)

	//fmt.Println(info, util.ObjToJson(info), info["nic_"])
	//拼装数据，给上层api直接使用
	deviceTypeMapInfo := map[string]map[string]string{}
	deviceTypeMapInfo[v.DeviceTypeID] = map[string]string{
		"cpu_": info["cpu_"],
		"mem_": info["mem_"],
		"sv_":  info["sv_"],
		"dv_":  info["dv_"],
		"gpu_": info["gpu_"],
		"nic_": info["nic_"],
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

//获取所有符合条件的设备，可以用于查询库存，查询所有设备接口使用
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

//查询设备详情单独使用
func DeviceEntity2DeviceOne(logger *log.Logger, d *deviceDao.Device, deviceTypeMapInfo map[string]map[string]string) *responseTypes.Device {

	tz := logger.GetPoint("timezone").(string)
	idc, err := idcDao.GetIdcById(logger, d.IDcID) //获取机房信息
	if err != nil {
		idc = &idcDao.Idc{}
	}
	//idc := idcMap[d.IDcID]
	entity, err := deviceTypeDao.GetOneDeviceType(logger, map[string]interface{}{"device_type_id": d.DeviceTypeID})
	if err != nil {
		entity = &deviceTypeDao.DeviceType{}
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
	disk, _ := diskDao.GetAllDisk(logger, map[string]interface{}{"device_id": d.DeviceID})
	if len(disk) == 1 {
		encloure1 = disk[0].Enclosure
		slot1 = disk[0].Slot
	} else if len(disk) == 2 {
		encloure1 = disk[0].Enclosure
		slot1 = disk[0].Slot
		encloure2 = disk[1].Enclosure
		slot2 = disk[1].Slot
	}
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
	return &responseTypes.Device{
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

		//拼装信息
		CpuInfo: deviceTypeMapInfo[entity.DeviceTypeID]["cpu_"],
		MemInfo: deviceTypeMapInfo[entity.DeviceTypeID]["mem_"],
		SvInfo:  deviceTypeMapInfo[entity.DeviceTypeID]["sv_"],
		DvInfo:  deviceTypeMapInfo[entity.DeviceTypeID]["dv_"],
		GpuInfo: deviceTypeMapInfo[entity.DeviceTypeID]["gpu_"],
		NicInfo: deviceTypeMapInfo[entity.DeviceTypeID]["nic_"],
	}
}

func GetGlobalDeviceCount(logger *log.Logger) (int64, error) {
	q := map[string]interface{}{
		"is_del": 0,
	}
	return deviceDao.GetDeviceCount(logger, q)
}
