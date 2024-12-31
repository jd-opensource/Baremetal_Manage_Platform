package RaidLogic

import (
	"errors"
	"strconv"

	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"

	log "coding.jd.com/aidc-bmp/bmp_log"

	sdkDeviceTypeParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/device_type"
	"coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

func QueryVolumesRaids(logger *log.Logger, param requestTypes.QueryVolumeRaidsRequest) ([]*response.VolumeRaid, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkDeviceTypeParameters.NewDescribeVolumesRaidsParams()
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithDeviceTypeID(&param.DeviceTypeID)
	request.WithVolumeType(&param.VolumeType)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	//request.SetOsType(param.ImageType)
	logger.Info("DescribeVolumesRaids openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.DeviceType.DescribeVolumesRaids(request, nil)
	logger.Info("DescribeVolumesRaids openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeVolumesRaids openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	res := []*response.VolumeRaid{}
	for _, v := range responseApi.Payload.Result {
		volumeSize, err := strconv.ParseFloat(v.VolumeSize, 64)
		if err != nil {
			logger.Warnf("QueryVolumesRaids parse VolumeSize error, volume_id:%s, volume_size:%s", v.VolumeID, v.VolumeSize)
		}
		item := &response.VolumeRaid{
			// 卷uuid
			VolumeID: v.VolumeID,
			// 设备类型uuid
			DeviceTypeID: v.DeviceTypeID,
			// 卷名称
			VolumeName: v.VolumeName,
			// 卷类型：系统卷，数据卷
			VolumeType: v.VolumeType,
			// 硬盘类型（SSD,HDD）
			DiskType: v.DiskType,
			// 接口类型（SATA,SAS,NVME,不限制）
			InterfaceType: v.InterfaceType,
			// 单盘大小（最小容量）
			VolumeSize: volumeSize,
			// 硬盘单位（GB,TB）
			VolumeUnit: v.VolumeUnit,
			// 硬盘数量（最低块数）
			VolumeAmount: int(v.VolumeAmount),
			//支持的raid模式
		}
		rs := []*response.Raid{}
		for _, r := range v.Raids {

			rs = append(rs, &response.Raid{
				RaidCan:  r.RaidCan,
				RaidID:   r.RaidID,
				RaidName: r.RaidName,
				//容量大小描述
				Detail: GetDiskDetailByRaid(v, &models.Raid{
					RaidID: r.RaidID,
					Name:   r.RaidName,
				}),
			})
		}
		item.Raids = rs
		res = append(res, item)
	}
	return res, nil

}

func GetDiskDetailByRaid(v *models.VolumeRaids, raid *models.Raid) string {

	volumeDetail := ""
	// avaliableValue := .0
	all := 0.0
	volumeSize, _ := strconv.ParseFloat(v.VolumeSize, 64)
	size := volumeSize
	all = volumeSize * float64(v.VolumeAmount)
	if raid != nil {
		if util.InArray(raid.Name, []string{"RAID1", "RAID10"}) {
			all = all / 2
		} else if util.InArray(raid.Name, []string{"RAID51H"}) {
			all = volumeSize * float64(v.VolumeAmount-2)
		} else if util.InArray(raid.Name, []string{"RAID5"}) {
			all = volumeSize * float64(v.VolumeAmount-1)
		}
	}
	var allStr, sizeStr string
	if v.VolumeUnit == "TB" { //保留两位小数
		allStr = strconv.FormatFloat(all, 'f', 2, 64)
		sizeStr = strconv.FormatFloat(size, 'f', 2, 64)
	} else { //不保留小数
		allStr = strconv.Itoa(int(all))
		sizeStr = strconv.Itoa(int(size))

	}

	volumeDetail = allStr + v.VolumeUnit + "(" +
		(strconv.Itoa(int(v.VolumeAmount)) + "*" + sizeStr + v.VolumeUnit) + " " + v.InterfaceType + ")"

	//NVME另外算法 TODO
	return volumeDetail

}
