package idc_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"git.jd.com/cps-golang/ironic-common/exception"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
	"gopkg.in/goyy/goyy.v0/util/crypto/hmac"
)

//返回数据格式 https://cf.jd.com/pages/viewpage.action?pageId=103918804
type IdcDevicesResp struct {
	StatusCode int32              `json:"status_code"`
	Data       map[string]IDevice `json:"data"`
}

type IDevice struct {
	Sn           string `json:"sn"`
	IloIP        string `json:"ilo_ip"`
	Cabinet      string `json:"cabinet"`
	UPosition    string `json:"u_position"`
	Department   string `json:"department"`
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	UpDown       string `json:"up_down"`
	Idc          string `json:"idc"`
}

var host, user, token string

func InitIdcApiConfig(idcHost, idcUser, idcToken string) error {
	if idcHost == "" || idcUser == "" || idcToken == "" {
		return errors.New("host|user|token missing...")
	}
	host = idcHost
	user = idcUser
	token = idcToken
	return nil
}

func QueryDevices(logger *log.Logger, sns []string) map[string]IDevice {

	t := fmt.Sprint(time.Now().Unix())
	auth, _ := hmac.Md5Hex(fmt.Sprintf("%s%s%s", user, token, t), "key")
	CMDB_GET_DEVICE_INFO_BY_SNS := fmt.Sprintf("http://%s/v1.0/search/device_info/sns?user=%s&auth=%s&timestamp=%s", host, user, auth, t)

	header := map[string]string{
		"Host":                 host,
		"Content-Type":         "application/json",
		"X-Jdcloud-Request-Id": logger.GetPoint("logid").(string),
	}
	resp, err := util.DoHttpPost(logger, CMDB_GET_DEVICE_INFO_BY_SNS, header, sns)
	if err != nil {
		logger.Warn("idc_api QueryDevices error:", sns, err.Error())
		panic(exception.Exception{httpStatus.INTERNAL_SERVER_ERROR, errorCode.INTERNAL, "queryDevices from idc fail"})
	}
	logger.Point("idc_api_resp", string(resp))
	var res = IdcDevicesResp{}
	if err = json.Unmarshal(resp, &res); err != nil {
		logger.Warn("idc_api QueryDevices response parse error:", sns, err.Error())
		panic(exception.Exception{httpStatus.INTERNAL_SERVER_ERROR, errorCode.INTERNAL, "queryDevices from idc fail"})
	}
	if res.StatusCode != 200 {
		panic(exception.Exception{httpStatus.INTERNAL_SERVER_ERROR, errorCode.INTERNAL, "queryDevices from idc is empty"})
	}
	return res.Data
}
