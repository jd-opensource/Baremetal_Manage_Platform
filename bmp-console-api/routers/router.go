// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/ws/msg", &controllers.WebSocketController{})
	beego.Router("/ws/innerSendmsg", &controllers.WebSocketController{}, "post:Send")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/logout", &controllers.UserController{}, "post:Logout")

	// 后面用户角色校验时需要放开
	// beego.InsertFilter("/v1/*", beego.BeforeRouter, controllers.Authentication)
	// beego.InsertFilter("/api/*", beego.BeforeRouter, controllers.Authentication)

	// beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterPin)

	//----------------------------- console api ----------------------------------------
	customer := beego.NewNamespace("/custom",
		beego.NSRouter("/setCustomInfo", &controllers.CustomController{}, "post:SetCustomInfo"), //暂时没用到，前端自己存
		beego.NSRouter("/getCustomInfo", &controllers.CustomController{}, "get:GetCustomInfo"),
	)
	beego.AddNamespace(customer)

	apikeyNS := beego.NewNamespace("/apikey",
		beego.NSRouter("/", &controllers.ApikeyController{}, "put:CreateApikey"), //td
		beego.NSRouter("/", &controllers.ApikeyController{}, "get:GetApikeyList"),
		beego.NSRouter("/?:apikey_id", &controllers.ApikeyController{}, "delete:DeleteApikey"),
	)
	beego.AddNamespace(apikeyNS)

	deviceNS := beego.NewNamespace("/device",
		beego.NSRouter("/isDeviceStockEnough", &controllers.DeviceController{}, "get:IsDeviceStockEnough"), //td
	)
	beego.AddNamespace(deviceNS)

	deviceTypeNS := beego.NewNamespace("/deviceType",
		beego.NSRouter("/getAvailableDeviceTypes", &controllers.DeviceTypeController{}, "get:GetAvailableDeviceTypes"), //td
	)
	beego.AddNamespace(deviceTypeNS)

	idcNS := beego.NewNamespace("/idc",
		beego.NSRouter("/", &controllers.IdcController{}, "get:GetIdcList"), //td
	)
	beego.AddNamespace(idcNS)

	imageNS := beego.NewNamespace("/image",
		beego.NSRouter("/queryImagesByDeviceType", &controllers.ImageController{}, "get:QueryImagesByDeviceType"), //td
	)
	beego.AddNamespace(imageNS)

	partitionNS := beego.NewNamespace("/partition",
		beego.NSRouter("/queryDefaultSystemPartitions", &controllers.PartitionController{}, "get:QueryDefaultSystemPartitions"), //td
	)
	beego.AddNamespace(partitionNS)

	projectNS := beego.NewNamespace("/project",
		beego.NSRouter("/", &controllers.ProjectController{}, "put:CreateProject"),                                     //td
		beego.NSRouter("/", &controllers.ProjectController{}, "get:GetProjectList"),                                    //td
		beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "get:GetProject"),                            //td
		beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "post:ModifyProject"),                        //修改项目名称
		beego.NSRouter("/?:project_id/description", &controllers.ProjectController{}, "post:ModifyProjectDescription"), //修改项目描述
		beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "delete:DeleteProject"),                      //td
		beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "delete:DeleteProject"),                      //td
		beego.NSRouter("/?:project_id/share", &controllers.ProjectController{}, "put:ShareProject"),                    //td
		beego.NSRouter("/?:project_id/cancelShare", &controllers.ProjectController{}, "put:CancelShareProject"),        //td
		beego.NSRouter("/?:project_id/move", &controllers.ProjectController{}, "put:MoveProject"),
		beego.NSRouter("/move/instances", &controllers.ProjectController{}, "put:MoveInstances"),
	)
	beego.AddNamespace(projectNS)

	raidNS := beego.NewNamespace("/raid",
		beego.NSRouter("/queryRaidsByDeviceTypeIDAndVolumeType", &controllers.RaidController{}, "get:QueryVolumeRaids"), //td
	)
	beego.AddNamespace(raidNS)

	userNS := beego.NewNamespace("/user",
		beego.NSRouter("/", &controllers.UserController{}, "get:GetUserInfo"),            //td
		beego.NSRouter("/password", &controllers.UserController{}, "put:UpdatePassword"), //td
		beego.NSRouter("/", &controllers.UserController{}, "put:UpdateUserInfo"),         //td
		beego.NSRouter("/timezones", &controllers.TimezoneController{}, "get:GetTimezoneList"),
		//获取有控制台权限的用户列表
		beego.NSRouter("/getConsoleAccessUserList", &controllers.UserController{}, "get:GetConsoleAccessUserList"),
		//校验用户是否有控制台权限
		beego.NSRouter("/checkUserConsoleAccess", &controllers.UserController{}, "get:CheckUserConsoleAccess"),
		//校验用户是否有控制台权限,按用户名
		beego.NSRouter("/checkUserConsoleAccessByName", &controllers.UserController{}, "get:CheckUserConsoleAccessByName"),
	)
	beego.AddNamespace(userNS)

	keypairNS := beego.NewNamespace("/keypair",
		beego.NSRouter("/", &controllers.KeypairController{}, "put:CreateKeypair"), //td
		beego.NSRouter("/", &controllers.KeypairController{}, "get:GetKeypairList"),
		beego.NSRouter("/?:keypair_id", &controllers.KeypairController{}, "get:GetKeypairInfo"),
		beego.NSRouter("/?:keypair_id", &controllers.KeypairController{}, "put:ModifyKeypair"),
		beego.NSRouter("/?:keypair_id", &controllers.KeypairController{}, "delete:DeleteKeypair"),
		beego.NSRouter("/checkKeypairName", &controllers.KeypairController{}, "get:CheckKeypairName"), //td
	)
	beego.AddNamespace(keypairNS)

	instanceNS := beego.NewNamespace("/instance",
		beego.NSRouter("/", &controllers.InstanceController{}, "get:GetInstanceList"),              //td
		beego.NSRouter("/?:instance_id", &controllers.InstanceController{}, "get:GetInstanceInfo"), //td
		beego.NSRouter("/", &controllers.InstanceController{}, "put:CreateInstance"),
		beego.NSRouter("/startInstance", &controllers.InstanceController{}, "post:StartInstance"),
		beego.NSRouter("/stopInstance", &controllers.InstanceController{}, "post:StopInstance"),
		beego.NSRouter("/modifyInstance", &controllers.InstanceController{}, "post:ModifyInstance"),
		beego.NSRouter("/restartInstance", &controllers.InstanceController{}, "post:RestartInstance"),
		beego.NSRouter("/resetPasswd", &controllers.InstanceController{}, "post:ResetPasswd"),
		beego.NSRouter("/batchResetPasswd", &controllers.InstanceController{}, "post:ResetInstancesPasswd"),
		beego.NSRouter("/?:instance_id", &controllers.InstanceController{}, "delete:DeleteInstance"),
		beego.NSRouter("/lockInstance", &controllers.InstanceController{}, "post:LockInstance"),
		beego.NSRouter("/unlockInstance", &controllers.InstanceController{}, "post:UnlockInstance"),
		beego.NSRouter("/startInstances", &controllers.InstanceController{}, "post:StartInstances"),     //批量开机
		beego.NSRouter("/stopInstances", &controllers.InstanceController{}, "post:StopInstances"),       //批量关机
		beego.NSRouter("/restartInstances", &controllers.InstanceController{}, "post:RestartInstances"), //批量重启
		beego.NSRouter("/modifyInstances", &controllers.InstanceController{}, "post:ModifyInstances"),   //批量编辑实例名称
		beego.NSRouter("/deleteInstances", &controllers.InstanceController{}, "delete:DeleteInstances"), //批量删除
		beego.NSRouter("/reinstallInstance", &controllers.InstanceController{}, "post:ReinstallInstance"),
		//部分共享的实例列表
		beego.NSRouter("/forshare/list", &controllers.InstanceController{}, "get:GetInstanceListForShare"),
	)
	beego.AddNamespace(instanceNS)

	// licenseNS := beego.NewNamespace("/license",
	// 	beego.NSRouter("/content", &controllers.LicenseController{}, "get:GetLicenceContent"), //td
	// )
	// beego.AddNamespace(licenseNS)

	messageNS := beego.NewNamespace("/messages",
		beego.NSRouter("/hasUnreadMessage", &controllers.MessageController{}, "get:HasUnreadMessage"),
		beego.NSRouter("/", &controllers.MessageController{}, "get:GetMessageList"),
		beego.NSRouter("/statistic", &controllers.MessageController{}, "get:GetMessageStatistic"),
		beego.NSRouter("/doRead", &controllers.MessageController{}, "put:ReadMessage"),
		beego.NSRouter("/delete", &controllers.MessageController{}, "delete:DeleteMessage"),
		beego.NSRouter("/getMessageById", &controllers.MessageController{}, "get:GetMessageById"),
		beego.NSRouter("/getMessageTypes", &controllers.MessageController{}, "get:GetMessageTypes"),
	)
	beego.AddNamespace(messageNS)

	auditLogNS := beego.NewNamespace("/auditLogs",
		beego.NSRouter("/", &controllers.AuditLogsController{}, "get:DescribeAuditLogs"),
		beego.NSRouter("/types", &controllers.AuditLogsController{}, "get:DescribeAuditLogTypes"),
	)
	beego.AddNamespace(auditLogNS)

	//报警规则相关
	monitorRuleNs :=
		beego.NewNamespace("/monitorRule",
			//添加规则
			beego.NSRouter("/addRule", &controllers.MonitorRuleController{}, "post:AddRule"),
			//规则详情
			beego.NSRouter("/describeRule", &controllers.MonitorRuleController{}, "get:DescribeRule"),
			//规则列表
			beego.NSRouter("/describeRules", &controllers.MonitorRuleController{}, "get:DescribeRules"),
			//编辑规则
			beego.NSRouter("/editRule", &controllers.MonitorRuleController{}, "put:EditRule"),
			//启用规则
			beego.NSRouter("/enableRule", &controllers.MonitorRuleController{}, "put:EnableRule"),
			//禁用规则
			beego.NSRouter("/disableRule", &controllers.MonitorRuleController{}, "put:DisableRule"),
			//删除规则
			beego.NSRouter("/deleteRule", &controllers.MonitorRuleController{}, "delete:DeleteRule"),
		)

	beego.AddNamespace(monitorRuleNs)

	//监控报警相关
	monitorAlertNs :=
		beego.NewNamespace("/monitorAlert",
			//告警详情
			beego.NSRouter("/describeAlert", &controllers.MonitorAlertController{}, "get:DescribeAlert"),
			//告警列表
			beego.NSRouter("/describeAlerts", &controllers.MonitorAlertController{}, "get:DescribeAlerts"),
			//删除告警
			beego.NSRouter("/deleteAlert", &controllers.MonitorAlertController{}, "delete:DeleteAlert"),
		)

	beego.AddNamespace(monitorAlertNs)

	//带内监控-图表数据获取
	monitorDataNs :=
		beego.NewNamespace("/monitorData",
			beego.NSRouter("/", &controllers.MonitorDataController{}, "get:GetMonitorData"),
		)

	beego.AddNamespace(monitorDataNs)

	//带内监控-获取agent状态、获取tag列表
	monitorProxyNs :=
		beego.NewNamespace("/monitorProxy",
			beego.NSRouter("/desrcibeAgentStatus", &controllers.MonitorProxyController{}, "get:DesrcibeAgentStatus"),
			beego.NSRouter("/desrcibeTags", &controllers.MonitorProxyController{}, "get:DesrcibeTags"),
		)

	beego.AddNamespace(monitorProxyNs)

}
