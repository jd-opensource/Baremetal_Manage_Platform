package apikeyLogic

import (
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/apikeyDao"

	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	namespacePrefix "git.jd.com/cps-golang/ironic-common/ironic/common/NamespacePrefixs"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
)

func ApikeyEntity2Apikey(logger *log.Logger, o *apikeyDao.Apikey) *responseTypes.Apikey {

	tz := logger.GetPoint("timezone").(string)
	return &responseTypes.Apikey{
		ID:          o.ID,
		ApikeyID:    o.ApikeyID,
		Name:        o.Name,
		ReadOnly:    o.ReadOnly,
		Token:       o.Token,
		Type:        o.Type,
		UserID:      o.UserID,
		CreatedBy:   o.CreatedBy,
		CreatedTime: util.TimestampToString(int64(o.CreatedTime), tz),
		UpdatedBy:   o.UpdatedBy,
		UpdatedTime: util.TimestampToString(int64(o.UpdatedTime), tz),
	}
}

func GetApikeyById(logger *log.Logger, apikeyId string) (*responseTypes.Apikey, error) {
	userId := logger.GetPoint("userId").(string)
	entity, err := apikeyDao.GetApikeyById(logger, apikeyId)
	if err != nil {
		logger.Warn("GetapikeyByUuid sql error:", apikeyId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if entity.UserID != userId {
		panic(constant.BuildInvalidArgumentWithArgs("无此apikey的访问权限", "cannot access apiKey"))
	}
	return ApikeyEntity2Apikey(logger, entity), nil
}

func CreateApikey(logger *log.Logger, param *requestTypes.CreateApikeyRequest) (string, error) {

	userId := logger.GetPoint("userId").(string)

	v, _ := apikeyDao.GetAllApikey(logger, map[string]interface{}{
		"user_id": userId,
		"is_del":  baseLogic.IS_NOT_DEL,
		"source":  "console",
	})
	if len(v) >= 2 {
		logger.Warn("two apikey already exists", userId)
		panic(constant.BuildCANCELLEDWithArgs("支持创建2个API密钥", "Only supports creating 2 keypairs"))
	}

	list, _ := apikeyDao.GetAllApikey(logger, map[string]interface{}{
		"name":    param.Name,
		"user_id": userId,
		"is_del":  baseLogic.IS_NOT_DEL,
		"source":  "console",
	})
	if len(list) != 0 {
		logger.Warn("apikey name exist:", param.Name)
		panic(constant.BuildCANCELLEDWithArgs("apikey名称已存在", "apikey name exist"))
	}
	apikeyEntity := &apikeyDao.Apikey{
		Name:        param.Name,
		ReadOnly:    param.ReadOnly,
		Token:       commonUtil.GetRandString("", 32, false, true, true),
		Type:        param.Type,
		Source:      "console",
		UserID:      userId,
		CreatedBy:   logger.GetPoint("username").(string),
		CreatedTime: int(time.Now().Unix()),
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
	}
	apikeyEntity.ApikeyID = commonUtil.GetRandString("apikey-", namespacePrefix.INSTANCE_ID_LENGTH, false, true, true)
	if _, err := apikeyDao.AddApikey(logger, apikeyEntity); err != nil {
		logger.Warnf("Createapikey Addapikey sql error, entity:%s, error:%s", util.ObjToJson(apikeyEntity), err.Error())
		return "", err
	}
	return apikeyEntity.ApikeyID, nil
}

func ModifyApikey(logger *log.Logger, param *requestTypes.ModifyApikeyRequest, apikeyId string) error {
	userId := logger.GetPoint("userId").(string)
	list, _ := apikeyDao.GetAllApikey(logger, map[string]interface{}{
		"name":    param.Name,
		"user_id": userId,
		"is_del":  baseLogic.IS_NOT_DEL,
	})
	if len(list) != 0 {
		logger.Warn("apikey name exist:", param.Name)
		panic(constant.BuildCANCELLEDWithArgs("apikey名称已存在", "apikey name exist"))
	}

	k, err := apikeyDao.GetApikeyById(logger, apikeyId)
	if err != nil {
		logger.Warn("GetapikeyByUuid sql error:", apikeyId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if userId != k.UserID {
		panic(constant.BuildInvalidArgumentWithArgs("无权修改此apikey", "cannot modify apikey"))
	}
	k.Name = param.Name
	k.ReadOnly = param.ReadOnly
	if err := apikeyDao.UpdateApikeyById(logger, k); err != nil {
		logger.Warn("Modifyapikey UpdateapikeyById sql error:", apikeyId, err.Error())
		return err
	}
	return nil
}

func DeleteApikey(logger *log.Logger, apikeyId string) error {
	userId := logger.GetPoint("userId").(string)
	entity, err := apikeyDao.GetApikeyById(logger, apikeyId)
	if err != nil {
		logger.Warn("GetapikeyByUuid sql error:", apikeyId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if entity.UserID != userId {
		panic(constant.BuildInvalidArgumentWithArgs("无权删除此apikey", "cannot delete apikey"))
	}
	entity.IsDel = 1
	entity.UpdatedTime = int(time.Now().Unix())
	entity.DeletedTime = int(time.Now().Unix())
	if err := apikeyDao.UpdateApikeyById(logger, entity); err != nil {
		logger.Warn("Deleteapikey UpdateapikeyById sql error:", apikeyId, err.Error())
		return err
	}
	return nil
}

func GetApikeyList(logger *log.Logger, param requestTypes.QueryApikeysRequest, p util.Pageable) ([]*responseTypes.Apikey, int64, error) {
	userId := logger.GetPoint("userId").(string)
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"name": param.Name,
		//"apikey_id":            param.apikeyID,
		//"default_project_id": param.DefaultProjectID,
		"user_id": userId,
		"is_del":  0,
		"source":  "console",
	}
	count, err := apikeyDao.GetApikeyCount(logger, query)
	if err != nil {
		logger.Warnf("GetapikeyCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	entityList := []*apikeyDao.Apikey{}
	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = apikeyDao.GetMultiApikey(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
	} else {
		entityList, err = apikeyDao.GetMultiApikey(logger, query, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
	}
	if err != nil {
		logger.Warn("QueryByapikeyIds GetAllapikey sql error:", err.Error())
		return nil, 0, err
	}
	apikeys := []*responseTypes.Apikey{}
	for _, entity := range entityList {
		apikeys = append(apikeys, ApikeyEntity2Apikey(logger, entity))
	}
	return apikeys, count, nil
}

func ValidateApiKey(logger *log.Logger, key string) (*apikeyDao.Apikey, error) {
	k, err := apikeyDao.GetApikeyByToken(logger, key)
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, nil
	}
	return k, nil
}
