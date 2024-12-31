package crontab

import (
	"context"
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/models/pushgateway"
	"github.com/astaxie/beego/logs"
)

func WipeAllPushgatewayData(ctx context.Context) error {
	requestId := util.GetUuid()
	logs.Info(requestId, fmt.Sprintf("start wipe pushgateway data，%v", util.GetCurrentTimeStr()))

	err := pushgateway.WipeAllPushgatewayData(requestId)
	if err != nil {
		logs.Info(requestId, fmt.Sprintf("wipe pushgateway data fail:%v", err))
		return err
	}
	logs.Info(requestId, fmt.Sprintf("wipe pushgateway data success! "))
	return nil
}