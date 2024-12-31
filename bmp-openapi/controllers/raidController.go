package controllers

import (
	"encoding/json"
	"regexp"

	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/raidLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	validation "git.jd.com/cps-golang/ironic-common/ironic/common/Validation"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// RaidController operations for raid
type RaidController struct {
	BaseController
}

// CreateRaid ...
// swagger:route POST /raids raid createRaid
//
// CreateRaid 创建raid(暂不启用)
//     Responses:
//       200: createRaid
//       default: ErrorResponse

func (c *RaidController) CreateRaid() {
	defer c.CatchException()
	req := &requestTypes.CreateRaidRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateRaid parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	uuid, err := raidLogic.CreateRaid(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.RaidId{
		RaidId: uuid,
	}
}

// swagger:route POST /raids/deviceTypeRaid raid createDeviceTypeRaid
//
// CreateDeviceTypeRaid ... 机型和raid关联(暂不启用)
//     Responses:
//       200: createDeviceTypeRaid
//       default: ErrorResponse

func (c *RaidController) CreateDeviceTypeRaid() {
	//defer c.CatchException()
	//req := &requestTypes.CreateDeviceTypeRaidRequest{}
	//if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
	//	c.logPoints.Warn("CreateDeviceTypeRaid parse post body error:", err.Error())
	//	c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	//	return
	//}
	//req.Validate(c.logPoints)
	//if err := raidLogic.CreateDeviceTypeRaid(c.logPoints, req); err != nil {
	//	c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	//	return
	//}
	c.Res.Result = response.CommonResponse{
		Success: false,
	}

}

// swagger:route DELETE /raids/deviceTypeRaid raid deleteDeviceTypeRaid
//
// DeleteDeviceTypeRaid ... 机型和raid解绑(暂不启用)
//     Responses:
//       200: deleteDeviceTypeRaid
//       default: ErrorResponse

func (c *RaidController) DeleteDeviceTypeRaid() {
	//defer c.CatchException()
	//req := &requestTypes.DeleteDeviceTypeRaidRequest{}
	//if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
	//	c.logPoints.Warn("DeleteDeviceTypeRaidRequest parse post body error:", err.Error())
	//	c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	//	return
	//}
	//req.Validate(c.logPoints)
	//if err := raidLogic.DeleteDeviceTypeRaid(c.logPoints, req); err != nil {
	//	c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	//	return
	//}
	c.Res.Result = response.CommonResponse{
		Success: false,
	}

}

// QueryRaidsAll ...
// swagger:route GET /raids raid describeRaids
//
// DescribeRaids 获取raid列表
//     Responses:
//       200: describeRaids
//       default: ErrorResponse

func (c *RaidController) DescribeRaids() {
	defer c.CatchException()

	data, err := raidLogic.QueryRaidsAll(c.logPoints)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.RaidList{
		Raids: data,
	}
}

// GetRaid ...
// swagger:route GET /raids/{raid_id} raid describeRaid
//
// DescribeRaid 获取raid详情
//     Responses:
//       200: describeRaid
//       default: ErrorResponse

func (c *RaidController) DescribeRaid() {
	defer c.CatchException()
	raidId := c.Ctx.Input.Param(":raid_id")
	if match, err := regexp.MatchString(validation.REGEX_ID, raidId); !match {
		if err != nil {
			c.logPoints.Warn("CreateImage validate volume_type error:", err.Error())
		}
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	raid, err := raidLogic.GetByRaidId(c.logPoints, raidId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = raid

}

// ModifyRaid ...
// swagger:route PUT /raids/{raid_id} raid modifyRaid
// ModifyRaid 修改raid信息(暂不启用)
//
//     Responses:
//       200: modifyRaid
//       default: ErrorResponse

func (c *RaidController) ModifyRaid() {
	defer c.CatchException()

	raidId := c.Ctx.Input.Param(":raid_id")
	if match, err := regexp.MatchString(validation.REGEX_ID, raidId); !match {
		if err != nil {
			c.logPoints.Warn("CreateImage validate volume_type error:", err.Error())
		}
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req := &requestTypes.ModifyRaidRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyRaid parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := raidLogic.ModifyRaid(c.logPoints, req, raidId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// DeleteRaid ...
// swagger:route DELETE /raids/{raid_id} raid deleteRaid
//
// DeleteRaid 删除raid(暂不启用)
//     Responses:
//       200: deleteRaid
//       default: ErrorResponse

func (c *RaidController) DeleteRaid() {
	defer c.CatchException()
	raid_id := c.Ctx.Input.Param(":raid_id")
	if match, err := regexp.MatchString(validation.REGEX_ID, raid_id); !match {
		if err != nil {
			c.logPoints.Warn("CreateImage validate volume_type error:", err.Error())
		}
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	if err := raidLogic.DeleteRaid(c.logPoints, raid_id); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}
