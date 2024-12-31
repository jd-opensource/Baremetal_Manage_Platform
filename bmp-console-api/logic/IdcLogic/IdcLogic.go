package IdcLogic

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/excel"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	sdkIdcParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/idc"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"git.jd.com/cps-golang/ironic-common/exception"
)

func GetIdcList(logger *log.Logger) (*sdkModels.IdcList, error) {

	logid := logger.GetPoint("logid").(string)
	userId := logger.GetPoint("userId").(string)
	req := sdkIdcParameters.NewDescribeIdcsParams()
	req.WithTraceID(logid)
	// req.WithIdcID(&param.IdcID)
	req.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("DescribeIdcs openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Idc.DescribeIdcs(req, nil)
	logger.Info("DescribeIdcs openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeIdcs openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result, nil
}

func ExportIdcList(logger *log.Logger, result []*sdkModels.Idc) (fileName, downLoadFileName string) {

	language := logger.GetPoint("language").(string)

	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)
	for _, data := range result {

		sheetData = append(sheetData, []string{
			data.Name,
			data.NameEn,
			data.Level,
			data.Address,
			data.CreateTime,
			data.CreatedBy,
			data.UpdateTime,
			data.UpdatedBy,
		})
	}
	exportIdcHeader := BaseLogic.ExportIdcHeader
	if language == "en_US" {
		exportIdcHeader = BaseLogic.ExportIdcHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportIdcHeader, sheetData); err != nil {
		panic(exception.Exception{Msg: "新建excel失败", Status: BaseLogic.ERROR_CODE})
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	downloadFileName := "idc_list_" + time.Now().Format(BaseLogic.FormatDate) + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx" //生成100000到100000+900000之间的随机数，左闭右开
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Info("保存excel到本地失败" + err.Error())
		panic(exception.Exception{Msg: "保存excel到本地失败", Status: BaseLogic.ERROR_CODE})
	}
	return fileName, downloadFileName
}
