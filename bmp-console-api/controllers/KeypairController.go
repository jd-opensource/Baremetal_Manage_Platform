package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/KeypairLoigc"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
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

// swagger:route GET /keypair/{keypair_id} keypair getKeypairInfo
// GetKeypairInfo 获取keypair详情
//     Responses:
//       200: getKeypairInfo
//       default: ErrorResponse

func (c *KeypairController) GetKeypairInfo() {
	defer c.CatchException()
	keypairID := c.Ctx.Input.Param(":keypair_id")

	res, err := KeypairLoigc.DescribeKeypair(c.logPoints, keypairID) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// swagger:route PUT /keypair keypair createKeypair
// CreateKeypair 创建keypair
//     Responses:
//       200: createKeypair
//       default: ErrorResponse

func (c *KeypairController) CreateKeypair() {
	defer c.CatchException()
	req := &requestTypes.CreateKeypairRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	res, err := KeypairLoigc.CreateKeypairs(c.logPoints, *req) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// swagger:route PUT /keypair/{keypair_id} keypair modifyKeypair
// ModifyKeypair 修改keypair信息
//     Responses:
//       200: modifyKeypair
//       default: ErrorResponse

func (c *KeypairController) ModifyKeypair() {
	defer c.CatchException()
	keypairID := c.Ctx.Input.Param(":keypair_id")
	req := &requestTypes.ModifyKeypairRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	res, err := KeypairLoigc.ModifyKeypair(c.logPoints, keypairID, req.Name) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// swagger:route DELETE /keypair/{keypair_id} keypair deleteKeypair
// DeleteKeypair 删除keypair
//     Responses:
//       200: deleteKeypair
//       default: ErrorResponse

func (c *KeypairController) DeleteKeypair() {
	defer c.CatchException()
	keypairID := c.Ctx.Input.Param(":keypair_id")
	res, err := KeypairLoigc.DeleteKeypairs(c.logPoints, keypairID) //返回byte数组
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// CheckKeypairName ...
// swagger:route GET /keypair/checkKeypairName keypair checkKeypairName
// CheckKeypairName 查询keypair名称是否已存在
//     Responses:
//       200: checkKeypairName
//       default: ErrorResponse

func (c *KeypairController) CheckKeypairName() {
	defer c.CatchException()
	req := &requestTypes.QueryKeypairsRequest{
		Name: c.GetString("name"),
	}
	res, err := KeypairLoigc.DescribeKeypairs(c.logPoints, *req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	result := response.CommonResponse{}
	if res.TotalCount == 0 { //如果没有这个秘钥名字，返回false，可以创建
		result.Success = false
	} else {
		result.Success = true
	}
	c.Res.Result = result
}
