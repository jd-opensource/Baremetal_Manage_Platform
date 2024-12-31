package idcLogic

import (
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/idcDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

func GetOneIdc(logger *log.Logger) (*responseTypes.Idc, error) {

	var tz string
	if logger.GetPoint("timezone") == nil {
		tz = "Asia/Shanghai"
	} else {
		tz = logger.GetPoint("timezone").(string)
	}
	query := map[string]interface{}{
		"is_del": 0,
	}
	idc, err := idcDao.GetOneIdc(logger, query)
	if err != nil {
		return nil, err
	}
	return IdcEntity2Idc(logger, idc, tz), nil
}

func QueryIdcs(logger *log.Logger, param *requestTypes.QueryIdcsRequest, p util.Pageable) ([]*responseTypes.Idc, int64, error) {
	tz := logger.GetPoint("timezone").(string)
	offset, limit := p.Pageable2offset()
	//tenant := logger.GetPoint("tenant").(string)
	//region := logger.GetPoint("region").(string)
	//az := logger.GetPoint("az").(string)
	//name := param.IdcNo
	query := map[string]interface{}{
		"is_del": 0,
	}
	//if tenant != "" {
	//	query["tenant"] = tenant
	//}
	count, err := idcDao.GetIdcCount(logger, query)
	if err != nil {
		logger.Warnf("DescribeIdcs_GetIdcCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	entityList := []*idcDao.Idc{}
	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = idcDao.GetMultiIdc(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
	} else {
		entityList, err = idcDao.GetMultiIdc(logger, query, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
	}
	if err != nil {
		logger.Warnf("DescribeIdcs_QueryByParam error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	Idcs := []*responseTypes.Idc{}
	for _, entity := range entityList {
		Idcs = append(Idcs, IdcEntity2Idc(logger, entity, tz))
	}
	return Idcs, count, nil
}

func IdcEntity2Idc(logger *log.Logger, k *idcDao.Idc, tz string) *responseTypes.Idc {
	v := &responseTypes.Idc{
		ID:              k.ID,
		IDcID:           k.IDcID,
		Name:            k.Name,
		NameEn:          k.NameEn,
		Shortname:       k.Shortname,
		Level:           k.Level,
		Address:         k.Address,
		IloUser:         k.IloUser,
		IloPassword:     k.IloPassword,
		SwitchUser1:     k.SwitchUser1,
		SwitchPassword1: k.SwitchPassword1,
		SwitchUser2:     k.SwitchUser2,
		SwitchPassword2: k.SwitchPassword2,
		CreateTime:      util.TimestampToString(int64(k.CreatedTime), tz),
		UpdateTime:      util.TimestampToString(int64(k.UpdatedTime), tz),
		CreatedBy:       k.CreatedBy,
		UpdatedBy:       k.UpdatedBy,
	}
	//if logger.GetPoint("language").(string) == baseLogic.EN_US {
	//	v.Name = k.NameEn
	//}
	return v
}
func CreateIdc(logger *log.Logger, param *requestTypes.CreateIdcRequest) (string, error) {

	//a, b := util.AesCBCEncrypt([]byte("hello"), []byte("mykeydsssdsdssds"))
	//fmt.Println(a, string(a), "err:", b)
	//
	//c, d := util.AesCBCDecrypt(a, []byte("mykeydsssdsdssds"))
	//fmt.Println(c, string(c), "err:", d)

	if idc, _ := idcDao.GetOneIdc(logger, map[string]interface{}{
		"name":   param.Name,
		"is_del": baseLogic.IS_NOT_DEL,
	}); idc != nil {
		panic(constant.BuildInvalidArgumentWithArgs("机房名称已存在", "idcName exist"))
	}
	if idc, _ := idcDao.GetOneIdc(logger, map[string]interface{}{
		"name_en": param.NameEn,
		"is_del":  baseLogic.IS_NOT_DEL,
	}); idc != nil {
		panic(constant.BuildInvalidArgumentWithArgs("机房英文名称已存在", "idcNameEn exist"))
	}

	entity := &idcDao.Idc{
		IDcID:           util.GetUuid("idc-"),
		Name:            param.Name,
		NameEn:          param.NameEn,
		Shortname:       param.Shortname,
		Level:           param.Level,
		Address:         param.Address,
		IloUser:         param.IloUser,
		IloPassword:     param.IloPassword,
		SwitchUser1:     param.SwitchUser1,
		SwitchPassword1: param.SwitchPassword1,
		SwitchUser2:     param.SwitchUser2,
		SwitchPassword2: param.SwitchPassword2,
		CreatedBy:       logger.GetPoint("username").(string),
		CreatedTime:     int(time.Now().Unix()),
		UpdatedBy:       logger.GetPoint("username").(string),
		UpdatedTime:     int(time.Now().Unix()),
	}
	_, err := idcDao.CreateIdc(logger, entity)
	if err != nil {
		logger.Warnf("CreateIdc AddIdc sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return "", err
	}
	return entity.IDcID, nil

}
func ModifyIdc(logger *log.Logger, param *requestTypes.ModifyIdcRequest, idcId string) error {

	if idc, _ := idcDao.GetOneIdc(logger, map[string]interface{}{
		"name":   param.Name,
		"is_del": baseLogic.IS_NOT_DEL,
	}); idc != nil && idc.IDcID != idcId && *param.Name != "" {
		panic(constant.BuildInvalidArgumentWithArgs("机房名称已存在", "idcName exist"))
	}
	if idc, _ := idcDao.GetOneIdc(logger, map[string]interface{}{
		"name_en": param.NameEn,
		"is_del":  baseLogic.IS_NOT_DEL,
	}); idc != nil && idc.IDcID != idcId && *param.NameEn != "" {
		panic(constant.BuildInvalidArgumentWithArgs("机房英文名称已存在", "idcNameEn exist"))
	}

	entity, err := idcDao.GetIdcById(logger, idcId)
	if err != nil {
		logger.Warn("GetIdcByUuid sql error:", idcId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if param.Name != nil {
		entity.Name = *param.Name
	}
	if param.NameEn != nil {
		entity.NameEn = *param.NameEn
	}
	if param.Shortname != nil {
		entity.Shortname = *param.Shortname
	}
	if param.Level != nil {
		entity.Level = *param.Level
	}
	if param.Address != nil {
		entity.Address = *param.Address
	}
	if param.IloUser != nil {
		entity.IloUser = *param.IloUser
	}
	if param.IloPassword != nil {
		entity.IloPassword = *param.IloPassword
	}
	if param.SwitchUser1 != nil {
		entity.SwitchUser1 = *param.SwitchUser1
	}
	if param.SwitchPassword1 != nil {
		entity.SwitchPassword1 = *param.SwitchPassword1
	}
	if param.SwitchUser2 != nil {
		entity.SwitchUser2 = *param.SwitchUser2
	}
	if param.SwitchPassword2 != nil {
		entity.SwitchPassword2 = *param.SwitchPassword2
	}

	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())

	if err := idcDao.UpdateIdcById(logger, entity, idcId); err != nil {
		logger.Warnf("ModifyIdc UpdateIdcById sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return err
	}
	return nil
}
func QueryIdcById(logger *log.Logger, idcId string) (*responseTypes.IdcInfo, error) {
	tz := logger.GetPoint("timezone").(string)
	idc, err := idcDao.GetIdcById(logger, idcId)
	if err != nil {
		logger.Warnf("DescribeIdc dao error, error:%s", err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	Idc := IdcEntity2Idc(logger, idc, tz)
	return &responseTypes.IdcInfo{
		Idc: Idc,
	}, nil
}
func DeleteIdc(logger *log.Logger, idcId string) error {
	entity, err := idcDao.GetIdcById(logger, idcId)
	if err != nil {
		logger.Warn("GetIdcByUuid sql error:", idcId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	entity.DeletedTime = int(time.Now().Unix())
	entity.IsDel = baseLogic.IS_DEL

	if err := idcDao.UpdateIdcById(logger, entity, idcId); err != nil {
		logger.Warnf("DeleteIdc UpdateIdcById sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return err
	}
	return nil
}
