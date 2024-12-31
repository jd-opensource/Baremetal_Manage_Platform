package osLogic

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	osDao "coding.jd.com/aidc-bmp/bmp-openapi/dao/osDao"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

func OsEntity2OS(o *osDao.Os, tz string) *responseTypes.Os {
	return &responseTypes.Os{
		ID:        o.ID,
		OsID:      o.OsID,
		OsName:    o.OsName,
		OsType:    o.OsType,
		OsVersion: o.OsVersion,

		Architecture: o.Architecture,
		Bits:         o.Bits,
		CreatedTime:  util.TimestampToString(int64(o.CreatedTime), tz),
	}
}

func GetByOsId(logger *log.Logger, osId string) (*responseTypes.Os, error) {
	tz := logger.GetPoint("timezone").(string)
	entity, err := osDao.GetOsByUuid(logger, osId)
	if err != nil {
		logger.Warn("GetOsByUuid sql error:", osId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	return OsEntity2OS(entity, tz), nil
}
func GetOsByUuid(logger *log.Logger, osid string) (o *osDao.Os, err error) {
	entity, err := osDao.GetOsByUuid(logger, osid)
	return entity, err
}

func GetAndCheckById(logger *log.Logger, os_id string) (*responseTypes.Os, error) {
	os, err := GetByOsId(logger, os_id)
	if err != nil {
		return nil, err
	}
	return os, nil

}

func DescribeOSs(logger *log.Logger, param *requestTypes.QueryOssRequest) ([]*responseTypes.Os, error) {

	tz := logger.GetPoint("timezone").(string)
	query := map[string]interface{}{
		"os_name":    param.OsName,
		"os_type":    param.OsType,
		"os_version": param.OsVersion,
		"is_del":     0,
	}
	os_entity_list, err := osDao.GetAllOs(logger, query) //osDao.GetMultiOs(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 1000)
	if err != nil {
		logger.Warn("QueryByOsIds GetAllOs sql error:", err.Error())
		return nil, err
	}
	oss := []*responseTypes.Os{}
	for _, entity := range os_entity_list {
		oss = append(oss, OsEntity2OS(entity, tz))
	}
	return oss, nil
}
