package exec

import (
	"encoding/json"
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/dao/api"
	"git.jd.com/bmp-pronoea/dao/api/bmpApi"
)

func ReportAlertInfoToBmp(requestId string, req *bmpApi.BmpAlertRequest) error {
	api.Locker.Lock()
	defer api.Locker.Unlock()

	a, err := api.GetApi(api.BMP_API)
	if err != nil {
		return err
	}
	base := a.GetBaseApi()

	currentTime := util.GetCurrentTimestamp()
	headers := a.GetRequestHeader(currentTime)
	headers["Traceid"] = requestId

	base.Options = &api.Options{
		Headers: headers,
		Data: map[string]api.IBaseApiRequest{
			api.PARAMS_TYPE_JSON:req,
		},
	}

	resp, err := api.Send(requestId, a, bmpApi.REPORT_ALERT)
	if err != nil {
		return err
	}
	getMetricRangeDataResp := new(bmpApi.BmpAlertResponse)
	if err = json.Unmarshal(resp, getMetricRangeDataResp); err != nil {
		return err
	}

	if getMetricRangeDataResp.Error != nil {
		return fmt.Errorf("report alert info error %v", getMetricRangeDataResp)
	}

	return nil
}
