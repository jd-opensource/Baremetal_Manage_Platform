package rules

import (
	"fmt"
	"git.jd.com/bmp-pronoea/common"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/config"
	"git.jd.com/bmp-pronoea/dao/api/prometheusApi/exec"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	RULE_MODELS_BMPINBAND = "bmpInBand"
)

type IRule interface {
	DeleteRules(requestId, ruleIds, source string) error
	WriteRule(requestId, ruleContentJson string) error
	RulesList(requestId, source string) (interface{}, error)
	//WriteAllRule(requestId string) error
	//TransforRuleYaml(requestId, ruleContentJson string) ([]byte, error)
}

type BaseRule struct {
	IRule
	//RulesPath  string     //规则文件目录
}

func (b *BaseRule) checkRulePath() error {
	if util.IsEmpty(config.SysConfig.PrometheusRulePath) {
		return fmt.Errorf("rules path empty ! ")
	}
	return nil
}

// DeleteRule 删除规则
func (b *BaseRule) DeleteRules(requestId, ruleIds, source string) error {
	if err := b.checkRulePath(); err != nil {
		return err
	}

	if util.IsEmpty(ruleIds) || util.IsEmpty(source) {
		return fmt.Errorf("rule ids or source empty! ")
	}

	ruleIdList := strings.Split(ruleIds, common.DELIMITER_COMMA)
	if ruleIdList == nil || len(ruleIdList) <= 0 {
		return fmt.Errorf("rule ids empty! ")
	}

	filePathList := make([]string, 0)
	for _, ruleId := range ruleIdList {
		if util.IsEmpty(ruleId) {
			continue
		}
		fileName := fmt.Sprintf("%v_%v", source, ruleId)
		filePath := fmt.Sprintf("%s/%s", config.SysConfig.PrometheusRulePath, fileName + ".yml")
		filePathList = append(filePathList, filePath)
	}

	if filePathList == nil || len(filePathList) <=0 {
		return fmt.Errorf("rule file path list empty ! ")
	}

	// 循环删除一下
	for _, filePath := range filePathList {
		err := b.DeleteRule(requestId, filePath)
		if err != nil {
			return fmt.Errorf("delete rule file error %v, %v", filePath, err)
		}
	}

	// 重启prometheus
	err := b.ReloadPromethus(requestId)
	if err != nil{
		logs.Info(requestId, fmt.Sprintf("remove rule file success[%v], but reload prometus error %v ", ruleIds, err))
		return fmt.Errorf("remove rule file success[%v], but reload prometus error %v ", ruleIds, err)
	}
	logs.Info(requestId, fmt.Sprintf("remove rule file success[%v], reload prometus success, res:%v  ", requestId, ruleIds))
	return nil
}

func (b *BaseRule) DeleteRule(requestId, filePath string) error {
	if util.IsEmpty(filePath) {
		return fmt.Errorf("rule file path empty ! ")
	}

	if !util.FileIsExist(filePath) {
		logs.Info(requestId, fmt.Sprintf("file not exist[%v], return success  ", filePath))
		return nil
	}

	err := util.RemoveFile(filePath)
	if err != nil {
		logs.Info(requestId,fmt.Sprintf("remove rule file error %v, file:%v", err, filePath))
		return fmt.Errorf("remove rule file error %v, file:%v", err, filePath)
	}
	return nil
}

func (b *BaseRule) WriteRule(requestId, ruleContentJson string) error {
	if err := b.checkRulePath(); err != nil {
		return err
	}
	if util.IsEmpty(ruleContentJson) {
		return fmt.Errorf("rule content empty! ")
	}
	return nil
}

func (b *BaseRule) WriteRuleAndReload(requestId, filePath string, yamlByte []byte, reload bool) error {
	err := util.CheckParent(filePath)
	if err != nil {
		logs.Info(requestId, fmt.Sprintf("write rule file error %v, file:%v", err, filePath))
		return err
	}
	err = ioutil.WriteFile(filePath, yamlByte, 0644)
	if err != nil {
		logs.Info(requestId, fmt.Sprintf("write rule file error %v, file:%v", err, filePath))
		return err
	}

	// 重启
	if reload {
		return b.ReloadPromethus(requestId)
	}

	return nil
}

func (b *BaseRule) ReloadPromethus(requestId string) error {
	err := exec.PrometheusReload(requestId)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseRule) RulesList(requestId, source string) (interface{}, error){
	return nil, nil
}

func (b *BaseRule) GetRulesFileList(requestId, source string) (map[string][]byte, error) {
	if err := b.checkRulePath(); err != nil {
		return nil, err
	}

	fileNameMatch := fmt.Sprintf("/%v_*.yml", source)
	fileLists, err := filepath.Glob(config.SysConfig.PrometheusRulePath + fileNameMatch)
	if err != nil {
		return nil, err
	}

	ruleFileMap := make(map[string][]byte, 0)
	if fileLists == nil || len(fileLists) <= 0 {
		return ruleFileMap, nil
	}

	for _, ruleFilePath := range fileLists {
		bytes, err := ioutil.ReadFile(ruleFilePath)
		if err != nil {
			logs.Info(requestId, fmt.Sprintf("read file:%v error:%v", ruleFilePath, err))
			return nil, err
		}
		ruleFileMap[ruleFilePath] = bytes
	}

	return ruleFileMap, nil
}


