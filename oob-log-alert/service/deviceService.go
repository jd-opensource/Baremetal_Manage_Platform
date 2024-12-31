package service

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	log "coding.jd.com/aidc-bmp/bmp_log"

	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	"coding.jd.com/aidc-bmp/oob-log-alert/util"

	"github.com/beego/beego/v2/core/logs"
)

const (
	IN           = "in"           //已入库
	PUTAWAY      = "putaway"      //已上架
	CREATED      = "created"      //已创建
	PUTAWAYING   = "putawaying"   //上架中
	PUTAWAYFAIL  = "putawayfail"  //上架失败
	UNPUTAWAY    = "unputaway"    //下架
	UNPUTAWAYING = "unputawaying" //下架中，将“已上架”的设备执行“下架”操作的中间过程状态
	REMOVED      = "removed"      //已移除

)

var DeviceManageStatus = map[string]string{
	IN:           "已入库",
	PUTAWAY:      "已上架",
	UNPUTAWAY:    "已下架",
	CREATED:      "已创建",
	PUTAWAYING:   "上架中",
	PUTAWAYFAIL:  "上架失败",
	UNPUTAWAYING: "下架中",
	REMOVED:      "已移除",
}

func GetDeviceStatus(logger *log.Logger, req dao.GetDeviceStatusRequest) (*dao.GetDeviceStatusResponse, error) {
	limit := req.PageSize
	offset := (req.PageNum - 1) * req.PageSize
	if limit == 0 {
		offset = 0
	}
	cnt, devices, err := GetDevices(logger, req.IdcId, req.Sn, limit, offset)
	if err != nil {
		return nil, err
	}

	oobstatus, err := GetDeviceOOBstatusBySns(logger, devices)
	if err != nil {
		logs.Debug("GetDeviceOOBstatusBySns error:%s", err.Error())
		return nil, err
	}

	return &dao.GetDeviceStatusResponse{
		Detail: oobstatus,
		PageResult: dao.PageResult{
			Total:      cnt,
			PageNumber: req.PageNum,
			PageSize:   req.PageSize,
		},
	}, nil

}

func GetDevices(logger *log.Logger, idc, sn string, limit, offset int64) (int64, []*dao.DeviceItem, error) {

	language := logger.GetPoint("language").(string)
	if sn != "" {
		matched, err := regexp.MatchString(`^[0-9a-zA-Z]{1,}$`, sn)
		if err != nil {
			return 0, nil, err
		}
		if !matched {
			return 0, nil, errors.New("sn invalid")
		}
	}

	cnt, err := dao.GetCpssCount(logger, sn, idc)
	if err != nil {
		return 0, nil, err
	}
	var ds []*dao.Device
	if limit == 0 {
		ds, err = dao.GetCpssByCondNoLimit(logger, sn, idc)
		if err != nil {
			return 0, nil, err
		}

	} else {
		ds, err = dao.GetCpssByCond(logger, sn, idc, int(limit), int(offset))
		if err != nil {
			return 0, nil, err
		}
	}
	var idcEntity *dao.Idc
	if len(ds) > 0 {
		idcEntity, _ = dao.GetIdcById(logger, ds[0].IDcID)
	}
	devices := []*dao.DeviceItem{}
	for _, d := range ds {
		ds := &dao.DeviceItem{
			IdcId:        d.IDcID,
			IdcName:      idcEntity.Name,
			Sn:           d.Sn,
			DeviceId:     d.DeviceID,
			Brand:        d.Brand,
			Model:        d.Model,
			ManageStatus: d.ManageStatus,
		}
		if language == util.EN_US {
			ds.IdcName = idcEntity.NameEn
		}
		devices = append(devices, ds)

	}
	return cnt, devices, nil
}

func GetDeviceOOBstatusBySns(logger *log.Logger, devices []*dao.DeviceItem) ([]*dao.DeviceStatus, error) {

	sns := []string{}
	for _, d := range devices {
		sns = append(sns, d.Sn)
	}

	ormStatusRes, err := dao.GetOwnedStatusBySns(logger, sns)
	if err != nil {
		return nil, err
	}

	res := map[string]*dao.DeviceStatus{}
	for _, de := range devices {
		res[de.Sn] = &dao.DeviceStatus{
			DeviceItem: *de,
			StatusItem: dao.StatusItem{},
		}
	}
	for _, v := range ormStatusRes {
		switch v.Owned {
		case "cpu":
			if v.Count > 0 {
				res[v.Sn].StatusItem.CpuStatus = 1
			}
		case "mem":
			if v.Count > 0 {
				res[v.Sn].StatusItem.MemStatus = 1
			}
		case "disk":
			if v.Count > 0 {
				res[v.Sn].StatusItem.DiskStatus = 1
			}
		case "power":
			if v.Count > 0 {
				res[v.Sn].StatusItem.PowerStatus = 1
			}
		case "nic":
			if v.Count > 0 {
				res[v.Sn].StatusItem.NicStatus = 1
			}
		case "other":
			if v.Count > 0 {
				res[v.Sn].StatusItem.OtherStatus = 1
			}

		}
	}

	finalRes := []*dao.DeviceStatus{}
	for _, sn := range sns {
		finalRes = append(finalRes, res[sn])
	}
	// for _, de := range res {
	// 	finalRes = append(finalRes, de)

	// }
	return finalRes, nil

}

func inArray(arr []int, v int) bool {
	for _, item := range arr {
		if item == v {
			return true
		}
	}
	return false
}

// 导出Excel文件
func ExportDeviceStatusExcel(logger *log.Logger, result []*dao.DeviceStatus) (string, string, error) {
	language := logger.GetPoint("language").(string)
	e := util.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)

	for _, data := range result {
		item := []string{
			data.Sn,      // SN
			data.IdcName, // idc
			data.Brand,
			data.Model,
			data.ManageStatus,
			util.StatusIntDesc(data.CpuStatus, language),
			util.StatusIntDesc(data.MemStatus, language),
			util.StatusIntDesc(data.DiskStatus, language),
			util.StatusIntDesc(data.NicStatus, language),
			util.StatusIntDesc(data.PowerStatus, language),
			util.StatusIntDesc(data.OtherStatus, language),
		}
		if language == util.CN_ZH {
			item[4] = DeviceManageStatus[data.ManageStatus]
		}
		sheetData = append(sheetData, item)
	}
	exportExcelHeader := util.ExportDeviceStatusHeader
	if language == util.EN_US {
		exportExcelHeader = util.ExportDeviceStatusHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportExcelHeader, sheetData); err != nil {
		logger.Warn("create excel error:", err.Error())
		return "", "", err
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	// 生成100000到100000+900000之间的随机数，左闭右开
	downloadFileName := "device_status_list_" + time.Now().Format(util.FormatDate) + "_" + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx"
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Warn("保存excel到本地失败" + err.Error())
		return "", "", err
	}
	return fileName, downloadFileName, nil
}
