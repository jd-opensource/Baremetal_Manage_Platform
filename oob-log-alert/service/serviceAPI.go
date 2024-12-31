package service

import (
	"encoding/json"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/oob-log-alert/object"

	"coding.jd.com/aidc-bmp/oob-log-alert/util"

	"github.com/beego/beego/v2/core/logs"
)

// Service bussiness service
type Service struct {
}

// GetCPS find CPS
func GetCPS(sn string) (object.CPSRecord, error) {

	var cps object.CPSRecord

	cpsStr, err := util.HGetObjectFromRedis("", sn)
	if err != nil {
		logs.Warn("GetCPSFromRedis Error: %s", err.Error())
		return cps, err
	}

	err = json.Unmarshal([]byte(cpsStr), &cps)
	if err != nil {
		logs.Warn("convert json to CPSRecord error, error:%s", err.Error())
		return cps, err
	}

	return cps, nil

}

// GetCPSFromRedis read CPS list from redis, according IDC ID in config file
func GetCPSFromRedis() (object.CPSObject, error) {

	var cpsObject object.CPSObject

	var cpsMap map[string]string
	var err error
	cpsMap, err = util.HGetAllFromRedis("")
	if len(cpsMap) == 0 {
		time.Sleep(10 * time.Second) //间隔10s重试一次
		cpsMap, err = util.HGetAllFromRedis("")
	}
	if err != nil {
		logs.Warn("get cps list error. %s", err.Error())
	}

	for _, v := range cpsMap {
		var cpsInfo object.CPSRecord
		err := json.Unmarshal([]byte(v), &cpsInfo)
		if err != nil {
			logs.Warn("convert json to CPSRecord error, error:%s", err.Error())
			continue
		}

		cpsObject.CPSs = append(cpsObject.CPSs, cpsInfo)
	}

	//logs.Debug(cpsObject)

	return cpsObject, nil
}

func inArrayStr(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// GetCPSFromRedis read CPS list from redis, according IDC ID in config file
func GetCPSFromRedisforAzAndInstanceID(namespace, idc string, instanceIds []string) (object.CPSObject, error) {

	var cpsObject object.CPSObject

	var cpsMap map[string]string
	var err error
	cpsMap, err = util.HGetAllFromRedisCustom(namespace, idc)
	if len(cpsMap) == 0 {
		time.Sleep(10 * time.Second) //间隔10s重试一次
		cpsMap, err = util.HGetAllFromRedisCustom(namespace, idc)
	}
	if err != nil {
		logs.Warn("get cps list error. %s", err.Error())
	}

	tmp := []object.CPSRecord{}
	for _, v := range cpsMap {
		var cpsInfo object.CPSRecord
		err := json.Unmarshal([]byte(v), &cpsInfo)
		if err != nil {
			logs.Warn("convert json to CPSRecord error, error:%s", err.Error())
			continue
		}

		tmp = append(tmp, cpsInfo)
	}

	//logs.Debug(cpsObject)
	if len(instanceIds) > 0 {
		tmp := cpsObject.CPSs
		for _, v := range tmp {
			if inArrayStr(v.InstanceID, instanceIds) {
				cpsObject.CPSs = append(cpsObject.CPSs, v)
			}
		}
	} else {
		cpsObject.CPSs = tmp
	}

	return cpsObject, nil
}

// writeCPSListToRedis write CPS to Redis
func WriteCPSListToRedis(cpsObject object.CPSObject) error {

	cpsMap := make(map[string]interface{})

	for _, v := range cpsObject.CPSs {
		jsonStr, err := v.JSON()
		if err != nil {
			logs.Warn(err.Error())
			continue
		}
		cpsMap[v.SN] = jsonStr
	}
	// write to Redis
	err := util.HMSetObjectToRedis("", cpsMap)
	if err != nil {
		logs.Warn("Write cps list to Redis. status: %v", err.Error())
		return err
	}

	return nil
}

func DelCPSFromRedis(cpsSNs []string) error {

	// del Hash fields
	err := util.HDelObjectFromRedis("", cpsSNs)
	if err != nil {
		logs.Warn("Del cps list from Redis. ERROR: %v", err.Error())
		return err
	}

	// del monitor keys
	err = util.DelObjectFromRedis(cpsSNs)
	if err != nil {
		logs.Warn("Del cps monitor from Redis. ERROR: %v", err.Error())
		return err
	}

	return nil
}

// GetInstanceMonitor get instance monitor data from redis. construct data
func GetInstanceMonitor(instanceID string) (interface{}, error) {

	all, err := util.HGetAllFromRedis(instanceID)
	if err != nil {
		logs.Warn("Get %s Error:%s", instanceID, err.Error())
		return nil, err
	}

	// lastMonitorTime := all[object.LastMonitorTimeField]
	pDisk, _ := getDisk(all[object.PDisksField])
	cpu, _ := getCPU(all[object.CPUsField])
	mem, _ := getMem(all[object.MemsField])
	nic, _ := getNIC(all[object.NicsField])

	data := map[string]interface{}{
		"instanceID": instanceID,
		// "lastMonitorTime": lastMonitorTime,
		"pdisks": pDisk,
		"cpus":   cpu,
		"mems":   mem,
		"nics":   nic,
	}

	return &data, nil

}

func getDisk(str string) (*object.PhysicalDiskObject, error) {

	if str == "" {
		return nil, nil
	}

	pDisk := object.PhysicalDiskObject{}
	pDisk.PDisks = []object.PhysicalDiskRecord{}

	err := json.Unmarshal([]byte(str), &(pDisk.PDisks))
	if err != nil {
		logs.Warn("Unmarshal Error:%s", err.Error())
	}

	pDisk.Status = true
	for _, v := range pDisk.PDisks {
		if strings.Compare(strings.ToUpper(v.PrimaryStatus), "OK") != 0 {
			pDisk.Status = false
		}
	}
	pDisk.Count = len(pDisk.PDisks)

	return &pDisk, nil
}

func getCPU(str string) (*object.CPUObject, error) {

	if str == "" {
		return nil, nil
	}

	cpu := object.CPUObject{}
	cpu.CPUs = []object.CPURecord{}

	err := json.Unmarshal([]byte(str), &(cpu.CPUs))
	if err != nil {
		logs.Warn("Unmarshal Error:%s", err.Error())
	}

	cpu.Status = true
	for _, v := range cpu.CPUs {
		if strings.Compare(strings.ToUpper(v.PrimaryStatus), "OK") != 0 {
			cpu.Status = false
		}
	}
	cpu.Count = len(cpu.CPUs)

	return &cpu, nil
}

func getMem(str string) (*object.MemObject, error) {

	if str == "" {
		return nil, nil
	}

	mem := object.MemObject{}
	mem.Mems = []object.MemRecord{}

	err := json.Unmarshal([]byte(str), &(mem.Mems))
	if err != nil {
		logs.Warn("Unmarshal Error:%s", err.Error())
	}

	mem.Status = true
	for _, v := range mem.Mems {
		if strings.Compare(strings.ToUpper(v.PrimaryStatus), "OK") != 0 {
			mem.Status = false
		}
	}
	mem.Count = len(mem.Mems)

	return &mem, nil
}

func getNIC(str string) (*object.NicObject, error) {

	if str == "" {
		return nil, nil
	}

	nic := object.NicObject{}
	nic.Nics = []object.NicRecord{}

	err := json.Unmarshal([]byte(str), &(nic.Nics))
	if err != nil {
		logs.Warn("Unmarshal Error:%s", err.Error())
	}

	nic.Status = true
	for _, v := range nic.Nics {
		if strings.Compare(strings.ToUpper(v.LinkStatus), "UP") != 0 {
			nic.Status = false
		}
	}
	nic.Count = len(nic.Nics)

	return &nic, nil
}
