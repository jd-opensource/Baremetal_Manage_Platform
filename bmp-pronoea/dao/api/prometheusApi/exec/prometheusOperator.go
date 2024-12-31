package exec

import (
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/dao/api"
	"git.jd.com/bmp-pronoea/dao/api/prometheusApi"
	"github.com/astaxie/beego/logs"
)

func PrometheusReload(requestId string) error {
	api.Locker.Lock()
	defer api.Locker.Unlock()

	a, err := api.GetApi(api.PROMETHEUS_API)
	if err != nil {
		return err
	}
	base := a.GetBaseApi()

	currentTime := util.GetCurrentTimestamp()
	headers := a.GetRequestHeader(currentTime)
	base.Options = &api.Options{
		Headers: headers,
	}

	resp, err := api.Send(requestId, a, prometheusApi.PROMETHEUS_RELOAD)
	logs.Info(requestId, fmt.Sprintf("reload prometheus res %v, err is %v", resp, err))

	if err != nil {
		return err
	}
	return nil
}

func PrometheusReloadByShell(requestId string) error {
	cmdRes, err := util.CmdRunWithOutput("curl -X POST http://127.0.0.1:9090/-/reload")
	if err != nil || util.StringsContains(cmdRes, "fail") {
		logs.Info(requestId,fmt.Sprintf("reload promethus by shell error %v, res:%v  ", err, cmdRes))
		return fmt.Errorf("reload prometus by shell error %v，res:%v ", err, cmdRes)
	}
	logs.Info(requestId, fmt.Sprintf("reload promethus by shell success, res:%v ", cmdRes))
	return nil
}
