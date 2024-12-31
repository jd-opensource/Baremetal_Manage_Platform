package service

import (
	"math/rand"
	"strconv"
	"time"

	"coding.jd.com/aidc-bmp/oob-log-alert/util"

	log "coding.jd.com/aidc-bmp/bmp_log"
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

func GetLogCollections(logger *log.Logger, req dao.GetLogCollectionsRequest) (*dao.GetLogCollectionsResponse, error) {

	language := logger.GetPoint("language").(string)
	limit := req.PageSize
	offset := (req.PageNum - 1) * req.PageSize
	if limit == 0 {
		offset = 0
	}
	if language == util.EN_US {
		req.Spec = partsTran[req.Spec]
	}
	logCollectionItems, cnt, err := dao.GetLogCollectionsforOperation(logger, req.Sn, req.Level, req.Spec, req.IdcId, req.StartTime, req.EndTime, req.DealStatus, limit, offset)
	if err != nil {
		return nil, err
	}
	if cnt == 0 {
		return &dao.GetLogCollectionsResponse{
			Detail: []*dao.LogRespItem{},
			PageResult: dao.PageResult{
				Total:      0,
				PageNumber: req.PageNum,
				PageSize:   req.PageSize,
			},
		}, nil
	}

	logger.Info("logCollectionItems res:", util.ToString(logCollectionItems))

	res := &dao.GetLogCollectionsResponse{
		Detail: []*dao.LogRespItem{},
		PageResult: dao.PageResult{
			Total:      cnt,
			PageNumber: req.PageNum,
			PageSize:   req.PageSize,
		},
	}

	var idcEntity *dao.Idc
	if len(logCollectionItems) > 0 {
		idcEntity, _ = dao.GetIdcById(logger, logCollectionItems[0].IdcId)
	}

	for _, v := range logCollectionItems {
		lri := &dao.LogRespItem{
			LogId:        v.ID,
			Sn:           v.Sn,
			DeviceId:     v.DeviceId,
			IloIp:        v.IloIp,
			IdcId:        v.IdcId,
			IdcName:      idcEntity.Name,
			AlertName:    v.Name,
			AlertDesc:    v.Desc,
			AlertType:    v.FaultType,
			AlertPart:    v.Parts,
			AlertLevel:   v.Level,
			AlertContent: v.Content,
			Count:        int(v.Count),
			CollectTime:  v.CollectTime,
			EventTime:    v.EventTime,
			UpdateTime:   v.UpdateTime,
			ManageStatus: v.ManageStatus,
		}
		if language == util.EN_US {
			lri.IdcName = idcEntity.NameEn
			lri.AlertPart = partsTran2En[lri.AlertPart]
		}
		if lri.Count == 0 {
			lri.Status = 1
		}

		res.Detail = append(res.Detail, lri)
	}
	return res, nil

}

// 导出Excel文件
func ExportLogCollectionsExcel(logger *log.Logger, result []*dao.LogRespItem) (string, string, error) {
	language := logger.GetPoint("language").(string)
	e := util.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)

	for _, data := range result {
		sheetData = append(sheetData, []string{
			util.ToString(data.LogId), // SN
			data.Sn,                   // idc
			data.IdcName,
			data.AlertLevel,
			data.AlertType,
			data.AlertPart,
			data.IloIp,
			data.AlertDesc,
			data.AlertContent,
			util.LogDealStatusIntDesc(int(data.Status), language),
			util.ToString(data.Count),
			util.Time2String(data.CollectTime),
			util.Time2String(data.EventTime),
			util.Time2String(data.UpdateTime),
		})
	}
	exportExcelHeader := util.ExportLogCollectionsHeader
	if language == util.EN_US {
		exportExcelHeader = util.ExportLogCollectionsHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportExcelHeader, sheetData); err != nil {
		logger.Warn("create excel error:", err.Error())
		return "", "", err
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	// 生成100000到100000+900000之间的随机数，左闭右开
	downloadFileName := "oob_alert_log_list_" + time.Now().Format(util.FormatDate) + "_" + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx"
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Warn("保存excel到本地失败" + err.Error())
		return "", "", err
	} else {
		logger.Infof("保存excel到本地成功，filename:%s", fileName)
	}
	return fileName, downloadFileName, nil
}

func GetLogCollectionsBySn(logger *log.Logger, req dao.GetLogCollectionsBySnRequest) (*dao.GetLogCollectionsResponse, error) {

	limit := req.PageSize
	offset := (req.PageNum - 1) * req.PageSize
	if limit == 0 {
		offset = 0
	}
	logCollectionItems, cnt, err := dao.GetLogCollectionsforConsole(logger, req.Sn, limit, offset)
	if err != nil {
		return nil, err
	}

	logger.Infof("GetLogCollectionsforConsole origin content for sn:%s is:%s", req.Sn, util.ToString(logCollectionItems))

	//获取sn和cps的对应关系
	sns := []string{}
	for _, v := range logCollectionItems {
		sns = append(sns, v.Sn)
	}
	if len(sns) == 0 {
		return &dao.GetLogCollectionsResponse{
			Detail: []*dao.LogRespItem{},
			PageResult: dao.PageResult{
				Total:      0,
				PageNumber: req.PageNum,
				PageSize:   req.PageSize,
			},
		}, nil
	}
	snToCps, err := GetSnToCpsMap(logger, sns)
	if err != nil {
		return nil, err
	}
	logger.Infof("GetSnToCpsMap origin content for sns:%s is:%s", util.ToString(sns), util.ToString(snToCps))

	//获取collectionId和faultrule的对应关系
	rids := []int{}
	for _, v := range logCollectionItems {
		rids = append(rids, int(v.FaultConfID))
	}
	ridToFaultRule, err := GetRidToFaultRuleMap(logger, rids)
	if err != nil {
		return nil, err
	}

	logger.Infof("GetRidToFaultRuleMap origin content for rids:%s is:%s", util.ToString(rids), util.ToString(ridToFaultRule))

	res := &dao.GetLogCollectionsResponse{
		Detail: []*dao.LogRespItem{},
		PageResult: dao.PageResult{
			Total:      cnt,
			PageNumber: req.PageNum,
			PageSize:   req.PageSize,
		},
	}

	for _, v := range logCollectionItems {
		lri := &dao.LogRespItem{
			LogId:        v.ID,
			Sn:           v.Sn,
			IloIp:        snToCps[v.Sn].IloIP,
			AlertName:    ridToFaultRule[int(v.FaultConfID)].Name,
			AlertDesc:    ridToFaultRule[int(v.FaultConfID)].Desc,
			AlertType:    ridToFaultRule[int(v.FaultConfID)].FaultType,
			AlertPart:    ridToFaultRule[int(v.FaultConfID)].Parts,
			AlertLevel:   ridToFaultRule[int(v.FaultConfID)].Level,
			AlertContent: v.Content,
			Count:        int(v.Count),
			CollectTime:  v.CollectTime,
			EventTime:    v.EventTime,
			UpdateTime:   v.UpdateTime,
		}
		if lri.Count == 0 {
			lri.Status = 1
		}

		res.Detail = append(res.Detail, lri)
	}
	return res, nil
}

// 导出Excel文件
func ExportLogCollectionBySnExcel(logger *log.Logger, result []*dao.LogRespItem) (string, string, error) {
	language := logger.GetPoint("language").(string)
	e := util.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)

	for _, data := range result {
		sheetData = append(sheetData, []string{
			util.ToString(data.LogId), //
			data.AlertType,
			data.AlertDesc,
			util.Time2String(data.EventTime),
			util.Time2String(data.UpdateTime),
			util.ToString(data.Count),
			util.LogDealStatusIntDesc(int(data.Status), language),
		})
	}
	exportExcelHeader := util.ExportLogCollectionBySnHeader
	if language == util.EN_US {
		exportExcelHeader = util.ExportLogCollectionBySnHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportExcelHeader, sheetData); err != nil {
		logger.Warn("create excel error:", err.Error())
		return "", "", err
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	// 生成100000到100000+900000之间的随机数，左闭右开
	downloadFileName := "oob_console_list_" + time.Now().Format(util.FormatDate) + "_" + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx"
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Warn("保存excel到本地失败" + err.Error())
		return "", "", err
	}
	return fileName, downloadFileName, nil
}

func GetSnToCpsMap(logger *log.Logger, sns []string) (map[string]*dao.Device, error) {
	cpss, err := dao.GetCpssBySns(logger, sns)
	if err != nil {
		return nil, err
	}
	res := map[string]*dao.Device{}
	for _, v := range cpss {
		res[v.Sn] = v
	}
	return res, nil
}

func GetRidToFaultRuleMap(logger *log.Logger, rids []int) (map[int]*dao.CpsFaultRules, error) {
	cfrs, err := dao.GetFaultRulesByRids(logger, rids)
	if err != nil {
		return nil, err
	}
	res := map[int]*dao.CpsFaultRules{}
	for _, v := range cfrs {
		res[v.ID] = v
	}
	return res, nil
}

func DealLogCollection(logger *log.Logger, req dao.DealLogCollectionRequest) error {

	cid := req.Cid
	return dao.DealCollectionByID(logger, cid)
}
