package controllers

import (
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"

	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/KeypairLoigc"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
)

// DeviceController operations for Device
type KeypairController struct {
	BaseController
}

// swagger:route GET /keypair keypair getKeypairList
// GetKeypairList 获取keypair列表
//     Responses:
//       200: getKeypairList
//       default: ErrorResponse

func (c *KeypairController) GetKeypairList() {
	defer c.CatchException()
	pg := requestTypes.PagingRequest{
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
	}
	req := &requestTypes.QueryKeypairsRequest{
		PagingRequest: pg,
		Name:          c.GetString("name"),
		KeypairIds:    c.GetStrings("keypairId"),
	}

	res, err := KeypairLoigc.DescribeKeypairs(c.logPoints, *req) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}
