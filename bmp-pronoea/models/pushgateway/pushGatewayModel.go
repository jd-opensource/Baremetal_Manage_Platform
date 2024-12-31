package pushgateway

import "git.jd.com/bmp-pronoea/dao/api/pushgatewayApi/exec"

func WipeAllPushgatewayData(requestId string) error {
	return exec.WipePushgatewayDataApi(requestId)
}
