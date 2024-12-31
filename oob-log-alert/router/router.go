package router

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	INS := beego.NewNamespace("/describeInstanceMonitor",
		// init api controller
		beego.NSRouter("/", &controllers.DescribeInstanceMonitorController{}, "get:Get"),
	)
	beego.AddNamespace(INS)

	InnerNS := beego.NewNamespace("/inner",
		// init api controller
		beego.NSRouter("/describeInstanceMonitor", &controllers.DescribeInstanceMonitorController{}, "get:Get"),
		beego.NSRouter("/trigPowerStatusReport", &controllers.InstanceStatusReportController{}, "get:Get"),
		beego.NSRouter("/trigDailyReport", &controllers.DailyReportController{}, "get:Get"),
	)
	beego.AddNamespace(InnerNS)

	v1NS := beego.NewNamespace("/v1/oob-alert",
		//新增接口
		beego.NSRouter("/device/status", &controllers.DeviceController{}, "get:GetDeviceStatus"),
		// beego.NSRouter("/device/getPowerStatusDiff", &controllers.DeviceController{}, "get:GetPowerStatusDiff"), //获取状态不一致接口
		// beego.NSRouter("/device/syncPowerStatus", &controllers.DeviceController{}, "post:SyncPowerStatus"),      //同步状态不一致接口
		//运营平台
		beego.NSRouter("/logs", &controllers.LogController{}, "get:GetLogCollections"),
		//控制台
		beego.NSRouter("/consolelogs", &controllers.LogController{}, "get:GetLogCollectionsBySn"),
		beego.NSRouter("/logs/deal", &controllers.LogController{}, "post:DealLogCollection"),

		beego.NSRouter("/rules/alert-spec-list", &controllers.RuleController{}, "get:AlertPartList"),
		beego.NSRouter("/rules/alert-level-list", &controllers.RuleController{}, "get:AlertLevelList"),
		beego.NSRouter("/rules", &controllers.RuleController{}, "get:RuleList"),

		beego.NSRouter("/rules/change-push", &controllers.RuleController{}, "post:ChangePush"),
		beego.NSRouter("/rules/change-use", &controllers.RuleController{}, "post:ChangeUse"),
		beego.NSRouter("/rules/reset", &controllers.RuleController{}, "post:RuleResetPushAndUse"),
		beego.NSRouter("/custom/set-custom-info", &controllers.CustomController{}, "post:SetCustomInfo"),
		beego.NSRouter("/custom/get-custom-info", &controllers.CustomController{}, "get:GetCustomInfo"),
	)
	beego.AddNamespace(v1NS)
}
