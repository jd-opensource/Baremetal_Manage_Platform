// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"coding.jd.com/aidc-bmp/bmp-operation-api/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	beego.Router("/ws/msg", &controllers.WebSocketController{})
	beego.Router("/ws/innerSendmsg", &controllers.WebSocketController{}, "post:Send")

	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/logout", &controllers.UserController{}, "post:Logout")

	// 后面用户角色校验时需要放开
	// beego.InsertFilter("/*", beego.BeforeRouter, controllers.Authentication)
	// beego.InsertFilter("/api/*", beego.BeforeRouter, controllers.Authentication)

	// beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterPin)

	//----------------------------- operation api ----------------------------------------
	beego.Router("/filterList", &controllers.CustomController{}, "get:FilterList")
	customer := beego.NewNamespace("/custom",
		beego.NSRouter("/setCustomInfo", &controllers.CustomController{}, "post:SetCustomInfo"), //暂时没用到，前端自己存
		beego.NSRouter("/getCustomInfo", &controllers.CustomController{}, "get:GetCustomInfo"),
	)
	beego.AddNamespace(customer)
	UserNS := beego.NewNamespace("/users",
		beego.NSRouter("/", &controllers.UserController{}, "put:AddUser"),
		beego.NSRouter("/", &controllers.UserController{}, "get:GetUserList"),
		beego.NSRouter("/?:user_id", &controllers.UserController{}, "get:GetUserInfo"),
		beego.NSRouter("/?:user_id", &controllers.UserController{}, "post:ModifyUserInfo"),
		beego.NSRouter("/?:user_id", &controllers.UserController{}, "delete:DeleteUserInfo"),
		beego.NSRouter("/login/validate", &controllers.UserController{}, "post:VerifyLoginUser"),
	)
	beego.AddNamespace(UserNS)

	//【设备管理】
	deviceNS := beego.NewNamespace("/devices",
		beego.NSRouter("/queryDevices", &controllers.DeviceController{}, "get:QueryDevices"),
		beego.NSRouter("/queryDevice/?:device_id", &controllers.DeviceController{}, "get:GetDevice"),

		beego.NSRouter("/unPutaway", &controllers.DeviceController{}, "put:UnPutaway"),
		beego.NSRouter("/putaway", &controllers.DeviceController{}, "put:Putaway"),
		//设备关联机型
		beego.NSRouter("/deviceType/associate", &controllers.DeviceController{}, "put:DeviceAssociateDeviceType"),
		beego.NSRouter("/?:device_id", &controllers.DeviceController{}, "delete:DeleteDevice"),
		beego.NSRouter("/?:device_id", &controllers.DeviceController{}, "put:ModifyDevice"),

		beego.NSRouter("/upload", &controllers.DeviceController{}, "post:UploadDevices"),
		beego.NSRouter("/collect", &controllers.DeviceController{}, "post:CollectDevice"),
		beego.NSRouter("/", &controllers.DeviceController{}, "post:CreateDevices"),
		beego.NSRouter("/?:device_id/remove", &controllers.DeviceController{}, "put:RemoveDevice"),
		//设备详情-磁盘
		beego.NSRouter("/?:device_id/disksDetail", &controllers.DeviceController{}, "get:GetDeviceDiskDetail"),
		//获取设备关联的磁盘信息
		beego.NSRouter("/disks/describeAssociateDisks", &controllers.DeviceController{}, "get:GetAssociatedDisks"),
		//设备关联磁盘
		beego.NSRouter("/disks/associateDisks", &controllers.DeviceController{}, "put:AssociateDeviceDisks"),
	)
	beego.AddNamespace(deviceNS)

	// 【机型管理】
	DeviceTypeNS := beego.NewNamespace("/deviceTypes",
		beego.NSRouter("/", &controllers.DeviceTypeController{}, "put:CreateDeviceType"),
		beego.NSRouter("/", &controllers.DeviceTypeController{}, "get:QueryDeviceTypes"),
		beego.NSRouter("/queryDeviceType/?:device_type_id", &controllers.DeviceTypeController{}, "get:QueryDeviceType"),
		beego.NSRouter("/deviceTypeImage", &controllers.DeviceTypeController{}, "get:QueryDeviceTypeImages"),
		beego.NSRouter("/associatedImage", &controllers.DeviceTypeController{}, "post:AssociateDeviceTypeImage"),
		beego.NSRouter("/dissociatedImage", &controllers.DeviceTypeController{}, "post:DisassociateDeviceTypeImage"),
		beego.NSRouter("/?:device_type_id", &controllers.DeviceTypeController{}, "put:ModifyDeviceType"),
		beego.NSRouter("/?:device_type_id", &controllers.DeviceTypeController{}, "delete:DeleteDeviceType"),
		beego.NSRouter("/?:device_type_id/describeVolumes", &controllers.DeviceTypeController{}, "get:DescribeVolumesByDeviceType"),
		//以下是重装使用的接口
		beego.NSRouter("/getAvailableDeviceTypes", &controllers.DeviceTypeController{}, "get:GetAvailableDeviceTypes"), //td
	)
	beego.AddNamespace(DeviceTypeNS)

	OsNS := beego.NewNamespace("/oss",
		beego.NSRouter("/", &controllers.OsController{}, "get:DescribeOSs"),
		beego.NSRouter("/filter", &controllers.OsController{}, "get:DescribeOSsFilter"),
	)
	beego.AddNamespace(OsNS)
	RoleNS := beego.NewNamespace("/roles",
		beego.NSRouter("/", &controllers.RoleController{}, "get:DescribeRoles"),
		beego.NSRouter("/roleInfo/current", &controllers.RoleController{}, "get:CurrentRole"),
	)
	beego.AddNamespace(RoleNS)

	ImageNS := beego.NewNamespace("/images",
		beego.NSRouter("/", &controllers.ImageController{}, "post:CreateImage"),
		beego.NSRouter("/", &controllers.ImageController{}, "get:DescribeImages"),
		beego.NSRouter("/?:image_id", &controllers.ImageController{}, "get:DescribeImage"),
		beego.NSRouter("/?:image_id", &controllers.ImageController{}, "put:ModifyImage"),
		beego.NSRouter("/?:image_id", &controllers.ImageController{}, "delete:DeleteImage"),
		beego.NSRouter("/upload", &controllers.ImageController{}, "post:UploadImage"),
		beego.NSRouter("/imageDeviceTypes", &controllers.ImageController{}, "get:DescribeImageDeviceTypes"),
		beego.NSRouter("/associatedDeviceType", &controllers.ImageController{}, "post:AssociatedDeviceType"),
		beego.NSRouter("/dissociatedDeviceType", &controllers.ImageController{}, "delete:DissociatedDeviceType"),

		beego.NSRouter("/queryImagesByDeviceType", &controllers.ImageController{}, "get:QueryImagesByDeviceType"),
	)
	beego.AddNamespace(ImageNS)

	RaidNS := beego.NewNamespace("/raids",
		//	beego.NSRouter("/", &controllers.RaidController{}, "post:CreateRaid"),
		//	beego.NSRouter("/deviceTypeRaid", &controllers.RaidController{}, "post:CreateDeviceTypeRaid"),
		//	beego.NSRouter("/deviceTypeRaid", &controllers.RaidController{}, "delete:DeleteDeviceTypeRaid"),
		//	beego.NSRouter("/", &controllers.RaidController{}, "get:QueryRaidsAll"), //QueryRaids
		beego.NSRouter("/", &controllers.RaidController{}, "get:QueryRaidsAll"),
		//	beego.NSRouter("/?:raid_id", &controllers.RaidController{}, "get:GetRaid"),
		//	beego.NSRouter("/?:raid_id", &controllers.RaidController{}, "put:ModifyRaid"),
		//	beego.NSRouter("/?:raid_id", &controllers.RaidController{}, "delete:DeleteRaid"),
		beego.NSRouter("/queryRaidsByDeviceTypeIDAndVolumeType", &controllers.RaidController{}, "get:QueryRaidsByDeviceTypeIDAndVolumeType"),
	)
	beego.AddNamespace(RaidNS)

	InstanceNS := beego.NewNamespace("/instances",
		//	beego.NSRouter("/", &controllers.InstanceController{}, "post:CreateInstance"),
		//	beego.NSRouter("/", &controllers.InstanceController{}, "get:GetInstanceList"),
		beego.NSRouter("/?:instance_id", &controllers.InstanceController{}, "get:GetInstanceInfo"),
		//	beego.NSRouter("/?:instance_id", &controllers.InstanceController{}, "put:ModifyInstance"),
		//	beego.NSRouter("/?:instance_id", &controllers.InstanceController{}, "delete:DeleteInstance"),
		//	beego.NSRouter("/lock/?:instance_id", &controllers.InstanceController{}, "put:LockInstance"),       //锁定
		//	beego.NSRouter("/unlock/?:instance_id", &controllers.InstanceController{}, "put:UnLockInstance"),   //解锁定
		beego.NSRouter("/start/?:instance_id", &controllers.InstanceController{}, "post:StartInstance"),             //开机
		beego.NSRouter("/restart/?:instance_id", &controllers.InstanceController{}, "post:RestartInstance"),         //重启
		beego.NSRouter("/stop/?:instance_id", &controllers.InstanceController{}, "post:StopInstance"),               //关机
		beego.NSRouter("/resetStatus/?:instance_id", &controllers.InstanceController{}, "post:ResetInstanceStatus"), //重置实例状态
		beego.NSRouter("/?:instance_id", &controllers.InstanceController{}, "delete:DeleteInstance"),                //回收实例
		beego.NSRouter("/lock/?:instance_id", &controllers.InstanceController{}, "post:LockInstance"),
		beego.NSRouter("/unlock/?:instance_id", &controllers.InstanceController{}, "post:UnlockInstance"),
		beego.NSRouter("/startInstances", &controllers.InstanceController{}, "post:StartInstances"),     //批量开机
		beego.NSRouter("/stopInstances", &controllers.InstanceController{}, "post:StopInstances"),       //批量关机
		beego.NSRouter("/restartInstances", &controllers.InstanceController{}, "post:RestartInstances"), //批量重启
		beego.NSRouter("/modifyInstances", &controllers.InstanceController{}, "post:ModifyInstances"),   //批量编辑实例名称
		beego.NSRouter("/deleteInstances", &controllers.InstanceController{}, "delete:DeleteInstances"), //批量删除
		beego.NSRouter("/resetPasswd", &controllers.InstanceController{}, "post:ResetPasswd"),
		beego.NSRouter("/batchResetPasswd", &controllers.InstanceController{}, "post:ResetInstancesPasswd"),
		beego.NSRouter("/reinstallInstance", &controllers.InstanceController{}, "post:ReinstallInstance"), //重装实例
	)
	beego.AddNamespace(InstanceNS)

	//
	//monitorNs := beego.NewNamespace("/offline") //beego.NSRouter("/instance/queryBySn", &controllers.InstanceController{}, "get:QueryBySn"),
	//
	//beego.AddNamespace(monitorNs)

	//RoleSN := beego.NewNamespace("/roles",
	//	beego.NSRouter("/", &controllers.RoleController{}, "post:CreateRole"),
	//	beego.NSRouter("/", &controllers.RoleController{}, "get:GetRoleList"),
	//	beego.NSRouter("/?:role_id", &controllers.RoleController{}, "get:GetRole"),
	//	beego.NSRouter("/?:role_id", &controllers.RoleController{}, "put:ModifyRole"),
	//	beego.NSRouter("/?:role_id", &controllers.RoleController{}, "delete:DeleteRole"),
	//)
	//beego.AddNamespace(RoleSN)
	//
	//ProjectSN := beego.NewNamespace("/projects",
	//	beego.NSRouter("/", &controllers.ProjectController{}, "post:CreateProject"),
	//	beego.NSRouter("/", &controllers.ProjectController{}, "get:GetProjectList"),
	//	beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "get:GetProject"),
	//	beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "put:ModifyProject"),
	//	beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "delete:DeleteProject"),
	//)
	//beego.AddNamespace(ProjectSN)
	//SshkeySN := beego.NewNamespace("/sshkeys",
	//	beego.NSRouter("/", &controllers.SshkeyController{}, "post:CreateSshkey"),
	//	beego.NSRouter("/", &controllers.SshkeyController{}, "get:GetSshkeyList"),
	//	beego.NSRouter("/?:sshkey_id", &controllers.SshkeyController{}, "get:GetSshkey"),
	//	beego.NSRouter("/?:sshkey_id", &controllers.SshkeyController{}, "put:ModifySshkey"),
	//	beego.NSRouter("/?:sshkey_id", &controllers.SshkeyController{}, "delete:DeleteSshkey"),
	//	beego.NSRouter("/instances/?:sshkey_id", &controllers.SshkeyController{}, "get:GetInstancesBySshkey"),
	//)
	//beego.AddNamespace(SshkeySN)
	//ApikeySN := beego.NewNamespace("/apikeys",
	//	beego.NSRouter("/", &controllers.ApikeyController{}, "post:CreateApikey"),
	//	beego.NSRouter("/", &controllers.ApikeyController{}, "get:GetApikeyList"),
	//	beego.NSRouter("/?:apikey_id", &controllers.ApikeyController{}, "get:GetApikey"),
	//	beego.NSRouter("/?:apikey_id", &controllers.ApikeyController{}, "put:ModifyApikey"),
	//	beego.NSRouter("/?:apikey_id", &controllers.ApikeyController{}, "delete:DeleteApikey"),
	//)
	//beego.AddNamespace(ApikeySN)

	keypairNS := beego.NewNamespace("/keypair",
		beego.NSRouter("/", &controllers.KeypairController{}, "get:GetKeypairList"),
	)
	beego.AddNamespace(keypairNS)

	IdcNS := beego.NewNamespace("/idc",
		beego.NSRouter("/", &controllers.IdcController{}, "put:CreateIdc"),
		beego.NSRouter("/", &controllers.IdcController{}, "get:GetIdcList"),
		beego.NSRouter("/?:idc_id", &controllers.IdcController{}, "get:GetIdcInfo"),
		beego.NSRouter("/?:idc_id", &controllers.IdcController{}, "delete:DeleteIdc"),
		beego.NSRouter("/?:idc_id", &controllers.IdcController{}, "put:ModifyIdcInfo"),
		beego.NSRouter("/VerifyUser", &controllers.IdcController{}, "post:VerifyUser"),
	)
	beego.AddNamespace(IdcNS)

	userNS := beego.NewNamespace("/user/local",
		beego.NSRouter("/", &controllers.UserController{}, "get:GetLocalUserInfo"),            //td
		beego.NSRouter("/password", &controllers.UserController{}, "put:UpdateLocalPassword"), //td
		beego.NSRouter("/", &controllers.UserController{}, "put:UpdateLocalUserInfo"),         //td
		beego.NSRouter("/timezones", &controllers.TimezoneController{}, "get:GetTimezoneList"),
	)
	beego.AddNamespace(userNS)

	apikeyNS := beego.NewNamespace("/apikey",
		beego.NSRouter("/", &controllers.ApikeyController{}, "put:CreateApikey"), //td
		beego.NSRouter("/", &controllers.ApikeyController{}, "get:GetApikeyList"),
		beego.NSRouter("/?:apikey_id", &controllers.ApikeyController{}, "delete:DeleteApikey"),
	)
	beego.AddNamespace(apikeyNS)

	resource := beego.NewNamespace("/resources",
		beego.NSRouter("/", &controllers.ResourceController{}, "get:DescribeResources"),
	)
	beego.AddNamespace(resource)

	partitionNS := beego.NewNamespace("/partition",
		beego.NSRouter("/queryDefaultSystemPartitions", &controllers.PartitionController{}, "get:QueryDefaultSystemPartitions"), //td
	)
	beego.AddNamespace(partitionNS)

	messageNS := beego.NewNamespace("/messages",
		beego.NSRouter("/hasUnreadMessage", &controllers.MessageController{}, "get:HasUnreadMessage"),
		beego.NSRouter("/", &controllers.MessageController{}, "get:GetMessageList"),
		beego.NSRouter("/statistic", &controllers.MessageController{}, "get:GetMessageStatistic"),
		beego.NSRouter("/doRead", &controllers.MessageController{}, "put:ReadMessage"),
		beego.NSRouter("/delete", &controllers.MessageController{}, "delete:DeleteMessage"),
		beego.NSRouter("/getMessageById", &controllers.MessageController{}, "get:GetMessageById"),
		beego.NSRouter("/getMessageTypes", &controllers.MessageController{}, "get:GetMessageTypes"),
		beego.NSRouter("/dialMail", &controllers.MessageController{}, "post:DialMail"),
		beego.NSRouter("/saveIsPushMail", &controllers.MessageController{}, "post:SaveIsPushMail"),
		beego.NSRouter("/describeMail", &controllers.MessageController{}, "get:DescribeMail"),
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
			beego.NSRouter("/deleteRule", &controllers.MonitorRuleController{}, "delete:EditRule"),
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

//TODO queryRaidsByDeviceTypeIDAndVolumeType queryDefaultSystemPartitions 未完成
