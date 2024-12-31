package config

import (
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"github.com/astaxie/beego"
	"strings"
)

const (
	CRONTAB_WIPE_PUSHGATEWAY = "WipePushgateway"
)

// CrontabList 这里添加其他crontab
var CrontabList = []string{
	CRONTAB_WIPE_PUSHGATEWAY,
}

func initCrontabConfig() map[string]string {
	crontabConfig := make(map[string]string)

	for _, crontabName := range CrontabList {
		crontabConfigDetail := beego.AppConfig.String(fmt.Sprintf("crontab.%v", strings.ToLower(crontabName)))
		if util.IsEmpty(crontabConfigDetail) {
			continue
		}
		crontabConfig[crontabName] = crontabConfigDetail
	}

	return crontabConfig
}
