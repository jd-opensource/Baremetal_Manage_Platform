package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"

	"coding.jd.com/aidc-bmp/oob-log-alert/object"

	"coding.jd.com/aidc-bmp/oob-log-alert/util"

	"coding.jd.com/aidc-bmp/oob-log-alert/dao"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

var statusMap map[string]string = map[string]string{
	"running":             "on",
	"stopped":             "off",
	"creating":            "unknown",
	"starting":            "unknown",
	"stopping":            "unknown",
	"restarting":          "unknown",
	"restarting_password": "unknown",
	"destroying":          "unknown",
	"destroyed":           "unknown",
	"error":               "unknown",
	"upgrading":           "unknown",
	"reinstalling":        "unknown",
}

func GetPowerStatusDiff(idc string, instanceIds []string) ([]*dao.PowerStatus, error) {

	namespace, _ := web.AppConfig.String("redis.namespace")

	cpslist, err := GetCPSFromRedisforAzAndInstanceID(namespace, idc, instanceIds)
	if err != nil {
		logs.Warn("DailReportEvent() GetCPSFromRedis Error:%s", err.Error())
		return nil, err
	}
	updateRes := getInstancePowerStatusDiff(cpslist)

	v, _ := json.Marshal(updateRes)
	logs.Info("getInstancePowerStatusDiff success, idc:%s, instanceIds:%s, res:%s", idc, strings.Join(instanceIds, ","), string(v))

	res := []*dao.PowerStatus{}
	for _, v := range updateRes {
		// mailTpl.Info = append(mailTpl.Info, *v)
		if v.IloStatus == "ipmi unreachable" { //带外不通的忽略
			v.IloStatus = "--"
		}
		res = append(res, v)
	}

	return res, nil
}

func SyncPowerStatusApi(idc string, instanceIds []string) (bool, error) {

	instanceEntries, _ := GetPowerStatusDiff(idc, instanceIds)
	if len(instanceEntries) == 0 {
		logs.Warn("SyncPowerStatusApi GetPowerStatusDiff empty, instanceid:%s", instanceIds)
		return false, errors.New("param.instanceId powerstatus all sync already")
	}
	changes := map[string][]string{}
	for _, v := range instanceEntries {
		var targetStatus string
		if v.IloStatus == "on" {
			targetStatus = "running"
		} else if v.IloStatus == "off" {
			targetStatus = "stopped"
		} else {
			continue
		}
		changes[targetStatus] = append(changes[targetStatus], v.InstanceID)
	}

	for status, instanceIDs := range changes {
		if len(instanceIDs) == 0 {
			continue
		}
		if err := sendChangeStatusRequest(status, instanceIDs); err != nil {
			logs.Warn("sendChangeStatusRequest %s to status %s error:%s", strings.Join(instanceIDs, ","), status, err.Error())
		} else {
			logs.Info("sendChangeStatusRequest %s to status %s succeed.", strings.Join(instanceIDs, ","), status)
		}
	}
	return true, nil
}

func getInstancePowerStatusDiff(cpslist object.CPSObject) []*dao.PowerStatus {
	iloUser, _ := web.AppConfig.String("oob.username")
	iloPass, _ := web.AppConfig.String("oob.password")
	wg := sync.WaitGroup{}
	wg.Add(len(cpslist.CPSs))

	diffResChan := make(chan *dao.PowerStatus)
	updateRes := []*dao.PowerStatus{}
	go func() {
		for {
			i, ok := <-diffResChan
			if !ok { //chan关闭
				break
			}
			updateRes = append(updateRes, i)
		}
	}()

	rateLimitSize := 20 //令牌池大小，用来限制下并发量
	rateLimitchan := make(chan int, rateLimitSize)

	for _, cps := range cpslist.CPSs {
		rateLimitchan <- 1 //取令牌
		go func(cps object.CPSRecord) {
			defer func() {
				wg.Done()
				<-rateLimitchan //任务完成后还令牌
			}()
			if cps.InstanceID == "" { //未装机等边界情况下不检查
				return
			}
			if statusMap[cps.Status] == "unknown" { //实例处于边界状态(装机中，销毁中等)时不做检查
				return
			}
			ilo_status := getIloipPowerStatus(cps.IloIP, iloUser, iloPass)

			fmt.Printf("diff debug, ilo_status:%s, cps_status:%s\n", ilo_status, statusMap[cps.Status])
			if !strings.EqualFold(ilo_status, statusMap[cps.Status]) {
				diffResChan <- &dao.PowerStatus{
					SN:           cps.SN,
					InstanceID:   cps.InstanceID,
					InstanceName: cps.InstanceName,
					IdcId:        cps.IdcId,
					Pin:          cps.Pin,
					IloIp:        cps.IloIP,
					Status:       statusMap[cps.Status],
					IloStatus:    ilo_status,
					Ipv4:         cps.Ipv4,
					DeviceType:   cps.DeviceType,
				}
			}
		}(cps)
	}
	wg.Wait()
	close(diffResChan)

	time.Sleep(5 * time.Second)

	return updateRes
}

func DoStaticCpsPowerState() error {

	defer func() {
		if r := recover(); r != nil {
			t := make([]byte, 1<<16)
			runtime.Stack(t, true)
			// t, _ := json.Marshal(r)
			logs.Warn("DoStaticCpsPowerState.panic msg,", string(t))
		}
	}()

	logs.Trace(">>>>>>>>DoStaticCpsPowerState()")
	defer logs.Trace("<<<<<<<<DoStaticCpsPowerState()")
	cpslist, err := GetCPSFromRedis()
	if err != nil {
		logs.Warn("DailReportEvent() GetCPSFromRedis Error:%s", err.Error())
		return err
	}
	updateRes := getInstancePowerStatusDiff(cpslist)

	type StatusMail struct {
		DiffStatus  []*dao.PowerStatus
		Az          string
		Unreachable []*dao.PowerStatus
	}

	az, _ := web.AppConfig.String("app.az")
	statusMailData := StatusMail{
		DiffStatus:  []*dao.PowerStatus{},
		Unreachable: []*dao.PowerStatus{},
		Az:          az,
	}

	// v, _ := json.Marshal(updateRes)
	// fmt.Println("updateRes debug...", string(v))

	// mailTpl := powerStatusMailTpl{}
	// mailTpl.Info = []models.PowerStatus{}
	for _, v := range updateRes {
		// mailTpl.Info = append(mailTpl.Info, *v)

		if v.IloStatus == "ipmi unreachable" {
			statusMailData.Unreachable = append(statusMailData.Unreachable, v)
		} else {
			statusMailData.DiffStatus = append(statusMailData.DiffStatus, v)
		}
	}

	//对于on/off不一致的调用cps offline接口更改，并在邮件中体现
	//对于ipmi unreachable的只在邮件中体现，不调用接口
	changes := map[string][]string{}
	for _, v := range statusMailData.DiffStatus {
		var targetStatus string
		if v.IloStatus == "on" {
			targetStatus = "running"
		} else if v.IloStatus == "off" {
			targetStatus = "stopped"
		} else {
			continue
		}
		changes[targetStatus] = append(changes[targetStatus], v.InstanceID)
	}

	for status, instanceIDs := range changes {
		if len(instanceIDs) == 0 {
			continue
		}
		if err := sendChangeStatusRequest(status, instanceIDs); err != nil {
			logs.Warn("sendChangeStatusRequest %s to status %s error:%s", strings.Join(instanceIDs, ","), status, err.Error())
		} else {
			logs.Info("sendChangeStatusRequest %s to status %s succeed.", strings.Join(instanceIDs, ","), status)
		}
	}
	// send mail
	// errMail := util.SendMailByTpl(powerStatusTemplate, powerStatusMailSubject, statusMailData)
	// if errMail != nil {
	// 	logs.Warn("Send powerStatus Mail failed: " + errMail.Error())
	// 	return err
	// } else {
	// 	logs.Debug("Send powerStatus Mail succeed")
	// }
	return nil

}

func sendChangeStatusRequest(status string, instanceIDs []string) error {
	url, _ := web.AppConfig.String("app.cps-changestatus-api")
	idc, _ := web.AppConfig.String("app.idc")
	m := map[string]interface{}{
		"idc":          idc,
		"status":       status,
		"instance_ids": strings.Join(instanceIDs, ","),
	}
	param, err := json.Marshal(m)
	if err != nil {
		return err
	}
	token, _ := web.AppConfig.String("app.cps-changestatus-api.token")
	header := map[string]string{
		"X-JDCLOUD-PIN": token,
		"Content-type":  "application/json",
	}

	_, err = util.HTTPPostWithHeader(url, string(param), header)
	return err
}

func getIloipPowerStatus(iloip, username, password string) string {

	retStr, _, err := util.ExecCommand("/bin/bash", "-c", fmt.Sprintf(powerStatusCmd(), iloip, username, password))
	if err != nil {
		logs.Warn("%s getIloipPowerStatus impi Error:%s", iloip, err.Error())
		return "ipmi unreachable"
	}
	retStr = strings.TrimSpace(retStr)

	// fmt.Println("getIloipPowerStatus debug...", iloip, retStr)
	items := strings.Split(retStr, " ")
	return items[len(items)-1]
}

func powerStatusCmd() string {
	ipmiPath, _ := web.AppConfig.String("oob.ipmi-path")
	return ipmiPath + " -I lanplus -H %s -U %s -P %s chassis power status"
}
