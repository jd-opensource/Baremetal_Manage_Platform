package diskLogic

//func OsEntity2OS(o *osDao.Os) *responseTypes.OS {
//	return &responseTypes.OS{
//		ID:           o.ID,
//		OsID:         o.OsID,
//		OsName:       o.OsName,
//		OsType:       o.OsType,
//		OsVersion:    o.OsVersion,
//		Platform:     o.Platform,
//		Architecture: o.Architecture,
//		Bits:         o.Bits,
//		CreatedTime:  util.TimestampToUtc(int64(o.CreatedTime)),
//	}
//}
//
//func GetByOsId(logger *log.Logger, osId string) (*responseTypes.OS, error) {
//	entity, err := osDao.GetOsByUuid(logger, osId)
//	if err != nil {
//		logger.Warn("GetOsByUuid sql error:", osId, err.Error())
//		return nil, err
//	}
//	return OsEntity2OS(entity), nil
//}
//func GetOsByUuid(logger *log.Logger, osid string) (o *osDao.Os, err error) {
//	entity, err := osDao.GetOsByUuid(logger, osid)
//	return entity, err
//}
//
//func GetAndCheckById(logger *log.Logger, os_id string) (*responseTypes.OS, error) {
//	os, err := GetByOsId(logger, os_id)
//	if err != nil {
//		return nil, err
//	}
//	return os, nil
//
//}
//
//func CreateDisk(logger *log.Logger, param *requestTypes.CreateD) (string, error) {
//	os_entity := &osDao.Os{
//		OsName:       param.OsName,
//		OsType:       param.OsType,
//		Platform:     param.Platform,
//		Architecture: param.Architecture,
//		Bits:         param.Bits,
//		OsVersion:    param.OsVersion,
//		SysUser:      param.SysUser,
//		CreatedBy:    logger.GetPoint("username").(string),
//		CreatedTime:  int(time.Now().Unix()),
//	}
//	os_entity.OsID = commonUtil.GetRandString(namespacePrefix.IMAGE_ID_PREFIX, namespacePrefix.INSTANCE_ID_LENGTH, false, true, true)
//	if _, err := osDao.AddOs(logger, os_entity); err != nil {
//		logger.Warnf("CreateOs AddOs sql error, entity:%s, error:%s", util.ObjToJson(os_entity), err.Error())
//		return "", err
//	}
//	return os_entity.OsID, nil
//}
//
//func ModifyOS(logger *log.Logger, param *requestTypes.ModifyOSRequest, osId string) error {
//
//	entity, err := osDao.GetOsByUuid(logger, osId)
//	if err != nil {
//		logger.Warn("GetOsByUuid sql error:", osId, err.Error())
//		return err
//	}
//
//	entity.OsName = param.OsName
//	entity.OsType = param.OsType
//	entity.Platform = param.Platform
//	entity.Architecture = param.Architecture
//	entity.Bits = param.Bits
//	entity.OsVersion = param.OsVersion
//	entity.SysUser = param.SysUser
//	entity.UpdatedBy = logger.GetPoint("username").(string)
//	entity.UpdatedTime = int(time.Now().Unix())
//	if err := osDao.UpdateOsById(logger, entity); err != nil {
//		logger.Warn("ModifyOS UpdateOsById sql error:", osId, err.Error())
//		return err
//	}
//	return nil
//}
//
//func DeleteOS(logger *log.Logger, osId string) error {
//	entity, err := osDao.GetOsByUuid(logger, osId)
//	if err != nil {
//		logger.Warn("GetOsByUuid sql error:", osId, err.Error())
//		return err
//	}
//	entity.IsDel = 1
//	entity.UpdatedTime = int(time.Now().Unix())
//	entity.DeletedTime = int(time.Now().Unix())
//	if err := osDao.UpdateOsById(logger, entity); err != nil {
//		logger.Warn("DeleteOS UpdateOsById sql error:", osId, err.Error())
//		return err
//	}
//	return nil
//}
//
//func DescribeOSs(logger *log.Logger, param *requestTypes.QueryOssRequest) ([]*responseTypes.OS, error) {
//
//	query := map[string]interface{}{
//		"os_name":    param.OsName,
//		"os_type":    param.OsType,
//		"os_version": param.OsVersion,
//		"platform":   param.Platform,
//		"is_del":     0,
//	}
//	os_entity_list, err := osDao.GetMultiOs(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 1000)
//	if err != nil {
//		logger.Warn("QueryByOsIds GetAllOs sql error:", err.Error())
//		return nil, err
//	}
//	oss := []*responseTypes.OS{}
//	for _, entity := range os_entity_list {
//		oss = append(oss, OsEntity2OS(entity))
//	}
//	return oss, nil
//}
