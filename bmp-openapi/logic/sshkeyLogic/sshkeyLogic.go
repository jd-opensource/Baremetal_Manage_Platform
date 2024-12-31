package sshkeyLogic

import (
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rInstanceSshkeyDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/sshkeyDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	namespacePrefix "git.jd.com/cps-golang/ironic-common/ironic/common/NamespacePrefixs"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
)

func SshkeyEntity2Sshkey(logger *log.Logger, o *sshkeyDao.Sshkey) *responseTypes.Sshkey {
	tz := logger.GetPoint("timezone").(string)
	return &responseTypes.Sshkey{
		Id:          o.ID,
		SshkeyId:    o.SshkeyId,
		UserId:      o.UserId,
		Name:        o.Name,
		FingerPrint: o.FingerPrint,
		Key:         o.Key,
		CreatedBy:   o.CreatedBy,
		CreatedTime: util.TimestampToString(int64(o.CreatedTime), tz),
		UpdatedBy:   o.UpdatedBy,
		UpdatedTime: util.TimestampToString(int64(o.UpdatedTime), tz),
	}
}

func GetSshkeyById(logger *log.Logger, SshkeyId string) (*responseTypes.Sshkey, error) {
	entity, err := sshkeyDao.GetSshkeyById(logger, SshkeyId)
	if err != nil {
		logger.Warn("GetSshkeyByUuid sql error:", SshkeyId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	return SshkeyEntity2Sshkey(logger, entity), nil
}

func CreateSshkey(logger *log.Logger, param *requestTypes.CreateSshkeyRequest) (string, error) {

	userId := logger.GetPoint("userId").(string)

	v, _ := sshkeyDao.GetAllSshkey(logger, map[string]interface{}{
		"user_id": userId,
		"is_del":  0,
	})
	if len(v) >= 2 {
		logger.Warn("two sshkey already exists", userId)
		panic(constant.BuildCANCELLEDWithArgs("支持创建2个SSH密钥", "Only supports creating 2 keypairs"))

	}

	list, _ := sshkeyDao.GetAllSshkey(logger, map[string]interface{}{
		"name":    param.Name,
		"user_id": userId,
		"is_del":  0,
	})
	if len(list) != 0 {
		logger.Warn("sshkey Name exist:", param.Name)
		panic(constant.BuildCANCELLEDWithArgs("秘钥名称已存在", "sshkey name exist"))

	}

	finger_print, err := util.GetFingerprint(param.Key)
	if err != nil {
		logger.Warn("GetFingerprint error:", param.Key, err.Error())
		return "", err
	}

	sshkeyEntity := &sshkeyDao.Sshkey{
		UserId:      userId,
		Name:        param.Name,
		Key:         param.Key,
		FingerPrint: finger_print,
		CreatedBy:   logger.GetPoint("username").(string),
		CreatedTime: int(time.Now().Unix()),
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
	}
	sshkeyEntity.SshkeyId = commonUtil.GetRandString("sshkey-", namespacePrefix.INSTANCE_ID_LENGTH, false, true, true)
	if _, err := sshkeyDao.AddSshkey(logger, sshkeyEntity); err != nil {
		logger.Warnf("CreateSshkey AddSshkey sql error, entity:%s, error:%s", util.ObjToJson(sshkeyEntity), err.Error())
		return "", err
	}
	return sshkeyEntity.SshkeyId, nil
}

func ModifySshkey(logger *log.Logger, param *requestTypes.ModifySshkeyRequest, SshkeyId string) error {
	userId := logger.GetPoint("userId").(string)
	s, err := sshkeyDao.GetSshkeyById(logger, SshkeyId)
	if err != nil {
		logger.Warn("GetSshkeyByUuid sql error:", SshkeyId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if s.UserId != userId {
		panic(constant.PermissionDenyForObject)
	}

	list, _ := sshkeyDao.GetAllSshkey(logger, map[string]interface{}{
		"name":    param.Name,
		"user_id": userId,
		"is_del":  0,
	})
	if len(list) != 0 && list[0].SshkeyId != SshkeyId {
		logger.Warn("sshkey Name exist:", param.Name)
		panic(constant.BuildCANCELLEDWithArgs("秘钥名称已存在", "sshkey name exist"))
	}
	if param.Name != nil {
		s.Name = *param.Name
	}
	s.UpdatedBy = logger.GetPoint("username").(string)
	s.UpdatedTime = int(time.Now().Unix())
	if err := sshkeyDao.UpdateSshkeyById(logger, s); err != nil {
		logger.Warn("ModifySshkey UpdateSshkeyById sql error:", SshkeyId, err.Error())
		return err
	}
	return nil
}

func DeleteSshkey(logger *log.Logger, sshkeyId string) error {

	userId := logger.GetPoint("userId").(string)
	entity, err := sshkeyDao.GetSshkeyById(logger, sshkeyId)
	if err != nil {
		logger.Warn("GetSshkeyByUuid sql error:", sshkeyId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if entity.UserId != userId {
		panic(constant.PermissionDenyForObject)
	}

	instanctIds := GetInstancesBySshkey(logger, sshkeyId)
	if len(instanctIds) > 0 {
		q := map[string]interface{}{
			"instance_id.in": instanctIds,
			"sshkey_id":      sshkeyId,
			"is_del":         0,
		}
		u := map[string]interface{}{
			"is_del": 1,
		}
		if err := rInstanceSshkeyDao.UpdateMultirInstanceSshkey(logger, q, u); err != nil {
			logger.Warnf("DeleteSshkey.UpdateMultirInstanceSshkey error, instanceIds:%s, sshkeyId:%s, error:%s", strings.Join(instanctIds, ","), sshkeyId, err.Error())
			return err
		}
	}

	entity.IsDel = 1
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	entity.DeletedTime = int(time.Now().Unix())
	if err := sshkeyDao.UpdateSshkeyById(logger, entity); err != nil {
		logger.Warn("DeleteSshkey UpdateSshkeyById sql error:", sshkeyId, err.Error())
		return err
	}
	return nil
}
func GetInstancesBySshkey(logger *log.Logger, sshkeyId string) []string {
	entity, _ := rInstanceSshkeyDao.GetAllrInstanceSshkey(logger, map[string]interface{}{
		"sshkey_id": sshkeyId,
		"is_del":    0,
	})
	//if err != nil {
	//	logger.Warn("GetInstancesBySshkey sql error:", sshkeyId, err.Error())
	//	return err
	//}
	if entity == nil {
		return []string{}
	}
	ids := []string{}
	for _, v := range entity {
		ids = append(ids, v.InstanceID)
	}
	return ids
}
func GetSshkeyList(logger *log.Logger, param requestTypes.QuerySshkeysRequest, p util.Pageable) ([]*responseTypes.Sshkey, int64, error) {
	userId := logger.GetPoint("userId").(string)
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"name":    param.Name,
		"user_id": userId,
		//"Sshkey_id":            param.SshkeyID,
		//"default_project_id": param.DefaultProjectID,
		"is_del": 0,
	}
	count, err := sshkeyDao.GetSshkeyCount(logger, query)
	if err != nil {
		logger.Warnf("GetSshkeyCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	entityList := []*sshkeyDao.Sshkey{}
	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = sshkeyDao.GetMultiSshkey(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
	} else {
		entityList, err = sshkeyDao.GetMultiSshkey(logger, query, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
	}
	if err != nil {
		logger.Warn("QueryBySshkeyIds GetAllSshkey sql error:", err.Error())
		return nil, 0, err
	}
	sshkeys := []*responseTypes.Sshkey{}
	for _, entity := range entityList {
		sshkeys = append(sshkeys, SshkeyEntity2Sshkey(logger, entity))
	}
	return sshkeys, count, nil
}
