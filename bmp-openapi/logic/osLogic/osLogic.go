package osLogic

import (
	"fmt"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	osDao "coding.jd.com/aidc-bmp/bmp-openapi/dao/osDao"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	namespacePrefix "git.jd.com/cps-golang/ironic-common/ironic/common/NamespacePrefixs"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
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

func CreateOS(logger *log.Logger, param *requestTypes.CreateOSRequest) (string, error) {
	list, _ := osDao.GetAllOs(logger, map[string]interface{}{
		"os_name": param.OsName,
		"is_del":  baseLogic.IS_NOT_DEL,
	})
	if len(list) != 0 {
		logger.Warn("osName exist:", param.OsName)
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("os名称 %s 已存在", param.OsName), fmt.Sprintf("osName %s exist", param.OsName)))
	}
	os_entity := &osDao.Os{
		OsName:       param.OsName,
		OsType:       param.OsType,
		Architecture: param.Architecture,
		Bits:         param.Bits,
		OsVersion:    param.OsVersion,
		SysUser:      param.SysUser,
		CreatedBy:    logger.GetPoint("username").(string),
		CreatedTime:  int(time.Now().Unix()),
		UpdatedBy:    logger.GetPoint("username").(string),
		UpdatedTime:  int(time.Now().Unix()),
	}
	os_entity.OsID = commonUtil.GetRandString(namespacePrefix.IMAGE_ID_PREFIX, namespacePrefix.INSTANCE_ID_LENGTH, false, true, true)
	if _, err := osDao.AddOs(logger, os_entity); err != nil {
		logger.Warnf("CreateOs AddOs sql error, entity:%s, error:%s", util.ObjToJson(os_entity), err.Error())
		return "", err
	}
	return os_entity.OsID, nil
}

func ModifyOS(logger *log.Logger, param *requestTypes.ModifyOSRequest, osId string) error {

	entity, err := osDao.GetOsByUuid(logger, osId)
	if err != nil {
		logger.Warn("GetOsByUuid sql error:", osId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}

	entity.OsName = param.OsName
	entity.OsType = param.OsType

	entity.Architecture = param.Architecture
	entity.Bits = param.Bits
	entity.OsVersion = param.OsVersion
	entity.SysUser = param.SysUser
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	if err := osDao.UpdateOsById(logger, entity); err != nil {
		logger.Warn("ModifyOS UpdateOsById sql error:", osId, err.Error())
		return err
	}
	return nil
}

func DeleteOS(logger *log.Logger, osId string) error {
	entity, err := osDao.GetOsByUuid(logger, osId)
	if err != nil {
		logger.Warn("GetOsByUuid sql error:", osId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	entity.IsDel = 1
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	entity.DeletedTime = int(time.Now().Unix())
	if err := osDao.UpdateOsById(logger, entity); err != nil {
		logger.Warn("DeleteOS UpdateOsById sql error:", osId, err.Error())
		return err
	}
	return nil
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
