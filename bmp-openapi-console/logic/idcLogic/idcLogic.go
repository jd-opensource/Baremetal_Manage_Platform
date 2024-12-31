package idcLogic

import (
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
