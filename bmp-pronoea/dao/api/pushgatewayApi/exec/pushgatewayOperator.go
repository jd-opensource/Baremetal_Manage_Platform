package exec

import (
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/dao/api"
	"git.jd.com/bmp-pronoea/dao/api/pushgatewayApi"
	"github.com/astaxie/beego/logs"
)

func WipePushgatewayDataApi(requestId string) error {
	api.Locker.Lock()
	defer api.Locker.Unlock()

	a, err := api.GetApi(api.PUSHGATEWAY_API)
	if err != nil {
		return err
	}
	base := a.GetBaseApi()

	currentTime := util.GetCurrentTimestamp()
	headers := a.GetRequestHeader(currentTime)
	base.Options = &api.Options{
		Headers: headers,
	}

	resp, err := api.Send(requestId, a, pushgatewayApi.WIPE_ALL_DATA)
	if err != nil {
		return err
	}

	logs.Info(requestId, fmt.Sprintf("delete pushgateway data res %v", resp))

	return nil
}
