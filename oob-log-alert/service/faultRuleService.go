package service

import (
	"math/rand"
	"strconv"
	"time"

	log "coding.jd.com/aidc-bmp/bmp_log"
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	"coding.jd.com/aidc-bmp/oob-log-alert/util"
)

var owned []string = []string{
	"cpu",
	"disk",
	"nic",
	"mem",
	"power",
	"other",
}

var partsTran = map[string]string{
	"Fan":          "风扇",
	"PCIE Bus":     "PCIE总线",
	"Power Supply": "电源",
	"voltage":      "电压",
	"Temperature":  "温度",
	"CPU":          "CPU",
	"Memory":       "内存",
	"Hard Disk":    "硬盘",
	"NIC":          "网卡",
}

var partsTran2En = map[string]string{
	"风扇":     "Fan",
	"PCIE总线": "PCIE Bus",
	"电源":     "Power Supply",
	"电压":     "voltage",
	"温度":     "Temperature",
	"CPU":    "CPU",
	"内存":     "Memory",
	"硬盘":     "Hard Disk",
	"网卡":     "NIC",
}

func GetAllInUsedRuleIdsGroupByOwn(logger *log.Logger) (map[string][]int, error) {
	faultRules, err := dao.GetFaultPushedRules(logger)
	if err != nil {
		return nil, err
	}
	res := map[string][]int{}
	for _, v := range faultRules {
		_, ok := res[v.Owned]
		if ok {
			res[v.Owned] = append(res[v.Owned], v.ID)
		} else {
			res[v.Owned] = []int{v.ID}
		}
	}
	return res, nil
}

func GetAlertPartList(logger *log.Logger) ([]string, error) {

	/*
		rules, err := models.GetFaultRules()
		if err != nil {
			return nil, err
		}
		parts := map[string]struct{}{}
		for _, v := range rules {
			parts[v.Parts] = struct{}{}
		}
		res := []string{}
		for k := range parts {
			res = append(res, k)
		}
		return res, nil
	*/
	parts := []string{}
	language := logger.GetPoint("language").(string)
	if language != util.EN_US {
		for _, v := range partsTran {
			parts = append(parts, v)
		}
	} else {
		for k, _ := range partsTran {
			parts = append(parts, k)
		}
	}

	return parts, nil
}

func GetAlertLevelList(logger *log.Logger) ([]string, error) {

	rules, err := dao.GetFaultRules(logger)
	if err != nil {
		return nil, err
	}
	levels := map[string]struct{}{}
	for _, v := range rules {
		levels[v.Level] = struct{}{}
	}
	res := []string{}
	for k := range levels {
		res = append(res, k)
	}
	return res, nil
}

func GetRuleList(logger *log.Logger, req dao.RuleListRequest) (*dao.RuleListResponse, error) {
	language := logger.GetPoint("language").(string)
	limit := req.PageSize
	offset := (req.PageNum - 1) * req.PageSize
	if limit == 0 {
		limit = 20
		offset = 0
	}
	if language == util.EN_US {
		req.AlertSpec = partsTran[req.AlertSpec]
	}
	rules, cnt, err := dao.GetRuleList(logger, req.AlertName, req.AlertSpec, req.AlertLevel, limit, offset)
	if err != nil {
		return nil, err
	}
	res := &dao.RuleListResponse{
		Detail: []*dao.RuleListResponseItem{},
		PageResult: dao.PageResult{
			Total:      cnt,
			PageNumber: req.PageNum,
			PageSize:   req.PageSize,
		},
	}
	for _, v := range rules {
		item := &dao.RuleListResponseItem{
			RuleId:     v.ID,
			AlertName:  v.Name,
			AlertSpec:  v.Parts,
			AlertLevel: v.Level,
			AlertDesc:  v.Desc,
			AlertType:  v.FaultType,
			// 判定条件
			AlertCondition: v.Condition,
			// 判定阈值
			AlertThreshold: v.Threshold,
			// 是否推送，0表示关闭，1表示开启
			PushStatus: bool2int(v.IsPush),
			UseStatus:  bool2int(v.IsUse),
		}
		if language == util.EN_US {
			item.AlertSpec = partsTran2En[item.AlertSpec]
		}
		res.Detail = append(res.Detail, item)
	}

	return res, nil
}

// 导出Excel文件
func ExportRulesExcel(logger *log.Logger, result []*dao.RuleListResponseItem) (string, string, error) {
	language := logger.GetPoint("language").(string)
	e := util.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)

	for _, data := range result {
		sheetData = append(sheetData, []string{
			data.AlertName,
			data.AlertSpec,
			data.AlertType,
			data.AlertCondition,
			data.AlertThreshold,
			data.AlertLevel,
			data.AlertDesc,
			util.PushStatusIntDesc(data.PushStatus, language),
			util.PushStatusIntDesc(data.UseStatus, language),
		})
	}
	exportExcelHeader := util.ExportRulesHeader
	if language == util.EN_US {
		exportExcelHeader = util.ExportRulesHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportExcelHeader, sheetData); err != nil {
		logger.Warn("create excel error:", err.Error())
		return "", "", err
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	// 生成100000到100000+900000之间的随机数，左闭右开
	downloadFileName := "alert_rule_list_" + time.Now().Format(util.FormatDate) + "_" + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx"
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Warn("保存excel到本地失败" + err.Error())
		return "", "", err
	}
	return fileName, downloadFileName, nil
}

func bool2int(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func ChangeRulePush(logger *log.Logger, req dao.ChangePushRequest) (bool, error) {
	if err := dao.ChangePushStatusByRuleId(logger, req.RuleID, int8(req.PushStatus)); err != nil {
		return false, err
	}
	return true, nil
}

func ChangeRuleUse(logger *log.Logger, req dao.ChangeUseRequest) (bool, error) {
	if err := dao.ChangeUseStatusByRuleId(logger, req.RuleID, int8(req.UseStatus)); err != nil {
		return false, err
	}
	return true, nil
}

func RuleResetPushAndUse(logger *log.Logger, ids []int) (bool, error) {

	if err := dao.RuleResetPushAndUse(logger, ids); err != nil {
		return false, err
	}
	return true, nil
}
