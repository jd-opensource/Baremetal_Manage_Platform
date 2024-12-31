package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/logic/PartitionLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// PartitionController operations for partition
type PartitionController struct {
	BaseController
}

// swagger:route GET /partition/queryDefaultSystemPartitions keypair queryDefaultSystemPartitions
// QueryDefaultSystemPartitions 根据机型和镜像查系统盘默认分区列表
//     Responses:
//       200: queryDefaultSystemPartitions
//       default: ErrorResponse

func (c *PartitionController) QueryDefaultSystemPartitions() {
	defer c.CatchException()

	req := requestTypes.QueryDefaultSystemPartitionRequest{
		ImageID:      c.GetString("imageId"),
		DeviceTypeID: c.GetString("deviceTypeId"),
	}
	if err := req.Validate(c.logPoints); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.PARAMETER_ERROR)
		return
	}
	res, err := PartitionLogic.QueryDefaultSystemPartitions(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.PARAMETER_ERROR)
		return
	}
	c.Res.Result = res

}
