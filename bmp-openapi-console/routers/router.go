// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	//subnetNS := beego.NewNamespace("/subnets",
	//	beego.NSRouter("/", &controllers.SubnetController{}, "get:QuerySubnets"),
	//	beego.NSRouter("/?:subnet_id", &controllers.SubnetController{}, "get:GetSubnet"),
	//	beego.NSRouter("/", &controllers.SubnetController{}, "post:CreateSubnet"),
	//	beego.NSRouter("/?:subnet_id", &controllers.SubnetController{}, "put:ModifySubnet"),
	//	beego.NSRouter("/?:subnet_id", &controllers.SubnetController{}, "delete:DeleteSubnet"),
	//)
	//beego.AddNamespace(subnetNS)
	//
	//wanSubnetNS := beego.NewNamespace("/wanSubnet",
	//	beego.NSRouter("/", &controllers.WanSubnetController{}, "post:CreateWanSubnet"),
	//)
	//beego.AddNamespace(wanSubnetNS)

	deviceNS := beego.NewNamespace("/v1/devices",
		beego.NSRouter("/", &controllers.DeviceController{}, "get:DescribeDevices"),
		beego.NSRouter("/", &controllers.DeviceController{}, "post:DescribeDevices"),
		beego.NSRouter("/stock", &controllers.DeviceController{}, "get:DescribeDeviceStock"),
		beego.NSRouter("/?:device_id", &controllers.DeviceController{}, "get:DescribeDevice"),
	)
	beego.AddNamespace(deviceNS)

	DeviceTypeNS := beego.NewNamespace("/v1/deviceTypes",
		beego.NSRouter("/", &controllers.DeviceTypeController{}, "get:DescribeDeviceTypes"),
		beego.NSRouter("/?:device_type_id", &controllers.DeviceTypeController{}, "get:DescribeDeviceType"),
		beego.NSRouter("/deviceTypeImage", &controllers.DeviceTypeController{}, "get:DescribeDeviceTypeImages"),
		beego.NSRouter("/volumesRaids", &controllers.DeviceTypeController{}, "get:DescribeVolumesRaids"),
		beego.NSRouter("/deviceTypeImagePartition", &controllers.DeviceTypeController{}, "get:DescribeDeviceTypeImagePartitions"),
	)
	beego.AddNamespace(DeviceTypeNS)

	IdcNS := beego.NewNamespace("v1/idcs",

		beego.NSRouter("/", &controllers.IdcController{}, "get:DescribeIdcs"),
		beego.NSRouter("/?:idc_id", &controllers.IdcController{}, "get:DescribeIdc"),
	)

	beego.AddNamespace(IdcNS)

	OsNS := beego.NewNamespace("/v1/oss",
		beego.NSRouter("/", &controllers.OsController{}, "get:DescribeOSs"),
		beego.NSRouter("/?:os_id", &controllers.OsController{}, "get:DescribeOS"),
	)
	beego.AddNamespace(OsNS)

	ImageNS := beego.NewNamespace("/v1/images",
		beego.NSRouter("/", &controllers.ImageController{}, "get:DescribeImages"),
		beego.NSRouter("/?:image_id", &controllers.ImageController{}, "get:DescribeImage"),
		beego.NSRouter("/imageDeviceTypes", &controllers.ImageController{}, "get:DescribeImageDeviceTypes"),
	)
	beego.AddNamespace(ImageNS)

	RabbitMq := beego.NewNamespace("/rabbitmq",
		beego.NSRouter("/send", &controllers.RabbitMqController{}, "post:Send"),
		beego.NSRouter("/receive", &controllers.RabbitMqController{}, "get:Receive"),
	)
	beego.AddNamespace(RabbitMq)

	/*InterfaceNs := beego.NewNamespace("/interfaces",
		beego.NSRouter("/", &controllers.InterfaceController{}, "post:CreateInterface"),
		beego.NSRouter("/", &controllers.InterfaceController{}, "get:QueryInterfaces"),
		beego.NSRouter("/", &controllers.InterfaceController{}, "delete:DeleteInterface"),
	)
	beego.AddNamespace(InterfaceNs)

	SwitchNs := beego.NewNamespace("/switches",
		beego.NSRouter("/", &controllers.JdSwitchController{}, "put:CreateSwitch"),
		beego.NSRouter("/", &controllers.JdSwitchController{}, "post:UpdateSwitch"),
		beego.NSRouter("/", &controllers.JdSwitchController{}, "get:QuerySwitches"),
	)
	beego.AddNamespace(SwitchNs)*/

	//KeypairNS := beego.NewNamespace("/keypairs",
	//	beego.NSRouter("/", &controllers.KeypairController{}, "post:CreateKeypair"),
	//	beego.NSRouter("/?:keypair_id", &controllers.KeypairController{}, "delete:DeleteKeypair"),
	//	beego.NSRouter("/", &controllers.KeypairController{}, "get:DescribeKeyPairs"),
	//	beego.NSRouter("/?:keypair_id", &controllers.KeypairController{}, "get:DescribeKeyPair"),
	//	beego.NSRouter("/getKeypairByname/?:keypair_name", &controllers.KeypairController{}, "get:DescribeKeyPairByName"),
	//)
	//beego.AddNamespace(KeypairNS)

	/*MailNS := beego.NewNamespace("/mail",
		beego.NSRouter("/addReceiver", &controllers.MailReceiverController{}, "put:AddReceiver"),
		beego.NSRouter("/removeReceiver", &controllers.MailReceiverController{}, "put:RemoveReceiver"),
		beego.NSRouter("/listReceivers", &controllers.MailReceiverController{}, "get:QueryReceivers"),
	)
	beego.AddNamespace(MailNS)*/

	InstanceNS := beego.NewNamespace("/v1/project/instances",
		beego.NSRouter("/", &controllers.InstanceController{}, "post:CreateProjectInstance"),
		beego.NSRouter("/", &controllers.InstanceController{}, "get:DescribeProjectInstances"),
		beego.NSRouter("/?:instance_id", &controllers.InstanceController{}, "get:DescribeProjectInstance"),
		beego.NSRouter("/?:instance_id", &controllers.InstanceController{}, "put:ModifyProjectInstance"),
		beego.NSRouter("/?:instance_id", &controllers.InstanceController{}, "delete:DeleteProjectInstance"),
		beego.NSRouter("/?:instance_id/lock", &controllers.InstanceController{}, "put:LockProjectInstance"),        //锁定
		beego.NSRouter("/?:instance_id/unlock", &controllers.InstanceController{}, "put:UnLockProjectInstance"),    //解锁定
		beego.NSRouter("/?:instance_id/start", &controllers.InstanceController{}, "put:StartProjectInstance"),      //开机
		beego.NSRouter("/?:instance_id/stop", &controllers.InstanceController{}, "put:StopProjectInstance"),        //关机
		beego.NSRouter("/?:instance_id/restart", &controllers.InstanceController{}, "put:RestartProjectInstance"),  //重启
		beego.NSRouter("/?:instance_id/resetpasswd", &controllers.InstanceController{}, "put:ResetPasswd"),         //重置密码
		beego.NSRouter("/?:instance_id/resetStatus", &controllers.InstanceController{}, "put:ResetInstanceStatus"), //重置状态
		//批量开机
		beego.NSRouter("/batch/:param(instances[:]{1}startInstances)", &controllers.InstanceController{}, "put:StartInstances"),
		// //批量关机
		beego.NSRouter("/batch/:param(instances[:]{1}stopInstances)", &controllers.InstanceController{}, "put:StopInstances"),
		// //批量重启
		beego.NSRouter("/batch/:param(instances[:]{1}restartInstances)", &controllers.InstanceController{}, "put:RestartInstances"),
		// //批量编辑实例名称
		beego.NSRouter("/batch/:param(instances[:]{1}modifyInstances)", &controllers.InstanceController{}, "put:ModifyInstances"),
		// //批量删除实例
		beego.NSRouter("/batch/:param(instances[:]{1}deleteInstances)", &controllers.InstanceController{}, "delete:DeleteInstances"),
		// //批量重置密码
		beego.NSRouter("/batch/:param(instances[:]{1}resetPasswd)", &controllers.InstanceController{}, "put:ResetInstancesPasswd"),
		// 实例重装
		beego.NSRouter("/?:instance_id/reinstallInstance", &controllers.InstanceController{}, "put:ReinstallInstance"),
		//部分共享，获取实例列表
		beego.NSRouter("/share/describeInstances", &controllers.InstanceController{}, "get:DescribeInstancesByProjectIdAndOwnerNameAndSharerName"),
	)
	beego.AddNamespace(InstanceNS)

	offlineNs := beego.NewNamespace("/v1/offline",
		beego.NSRouter("/monitorRules/checkDiffFromPronoea", &controllers.MonitorRuleController{}, "get:CheckDiffFromPronoea"),
	)

	beego.AddNamespace(offlineNs)

	UserNS := beego.NewNamespace("/v1/users",
		beego.NSRouter("/verify", &controllers.UserController{}, "post:VerifyUser"),
		beego.NSRouter("/", &controllers.UserController{}, "get:DescribeUsers"),
		beego.NSRouter("/getUserByName", &controllers.UserController{}, "get:DescribeUserByName"),
		beego.NSRouter("/?:user_id", &controllers.UserController{}, "get:DescribeUser"),
		beego.NSRouter("/?:user_id", &controllers.UserController{}, "put:ModifyUser"),
	)
	beego.AddNamespace(UserNS)

	// 控制台个人中心
	LocalUserNS := beego.NewNamespace("/v1/local/users",
		beego.NSRouter("/", &controllers.UserController{}, "get:DescribeLocalUser"),
		beego.NSRouter("/", &controllers.UserController{}, "put:ModifyLocalUser"),
		beego.NSRouter("/password", &controllers.UserController{}, "put:ModifyLocalUserPassword"),
	)
	beego.AddNamespace(LocalUserNS)

	RoleSN := beego.NewNamespace("/v1/roles",
		beego.NSRouter("/", &controllers.RoleController{}, "get:DescribeRoles"),
		beego.NSRouter("/?:role_id", &controllers.RoleController{}, "get:DescribeRole"),
		beego.NSRouter("/roleInfo/current", &controllers.RoleController{}, "get:CurrentRole"),
	)
	beego.AddNamespace(RoleSN)

	ProjectSN := beego.NewNamespace("/v1/user/projects",
		beego.NSRouter("/", &controllers.ProjectController{}, "post:CreateUserProject"),
		beego.NSRouter("/", &controllers.ProjectController{}, "get:DescribeUserProjects"),
		beego.NSRouter("/?:project_id/share", &controllers.ProjectController{}, "put:ShareUserProject"),
		beego.NSRouter("/?:project_id/cancelShare", &controllers.ProjectController{}, "put:CancelShareUserProject"),
		beego.NSRouter("/?:project_id/move", &controllers.ProjectController{}, "put:MoveUserProject"),
		beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "get:DescribeUserProject"),
		beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "put:ModifyUserProject"),
		beego.NSRouter("/?:project_id/description", &controllers.ProjectController{}, "put:ModifyUserProjectDescription"),
		beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "delete:DeleteUserProject"),
		beego.NSRouter("/?:project_id/describeSharProject", &controllers.ProjectController{}, "get:DescribeShareUserProject"),
		beego.NSRouter("/move/instances", &controllers.ProjectController{}, "put:MoveUserInstances"),
	)
	beego.AddNamespace(ProjectSN)
	SshkeySN := beego.NewNamespace("/v1/user/sshkeys",
		beego.NSRouter("/", &controllers.SshkeyController{}, "post:CreateUserSshkey"),
		beego.NSRouter("/", &controllers.SshkeyController{}, "get:DescribeUserSshKeys"),
		beego.NSRouter("/?:sshkey_id", &controllers.SshkeyController{}, "get:DescribeUserSshKey"),
		beego.NSRouter("/?:sshkey_id", &controllers.SshkeyController{}, "put:ModifyUserSshkey"),
		beego.NSRouter("/?:sshkey_id", &controllers.SshkeyController{}, "delete:DeleteUserSshkey"),
		beego.NSRouter("/instances/?:sshkey_id", &controllers.SshkeyController{}, "get:GetInstancesBySshkey"),
	)
	beego.AddNamespace(SshkeySN)
	ApikeySN := beego.NewNamespace("/v1/user/apikeys",
		beego.NSRouter("/", &controllers.ApikeyController{}, "post:CreateUserApikey"),
		beego.NSRouter("/", &controllers.ApikeyController{}, "get:DescribeUserAPIKeys"),
		beego.NSRouter("/?:apikey_id", &controllers.ApikeyController{}, "get:DescribeUserAPIKey"),
		beego.NSRouter("/?:apikey_id", &controllers.ApikeyController{}, "put:ModifyUserApikey"),
		beego.NSRouter("/?:apikey_id", &controllers.ApikeyController{}, "delete:DeleteUserApikey"),
	)
	beego.AddNamespace(ApikeySN)
	resource := beego.NewNamespace("/v1/resources",
		beego.NSRouter("/", &controllers.ResourceController{}, "get:DescribeResources"),
	)
	beego.AddNamespace(resource)

	// license := beego.NewNamespace("/v1/license",
	// 	beego.NSRouter("/content", &controllers.LicenseController{}, "get:DescribeLicenseContent"),
	// )
	// beego.AddNamespace(license)

	messageNS := beego.NewNamespace("/v1/messages",
		beego.NSRouter("/hasUnreadMessage", &controllers.MessageController{}, "get:HasUnreadMessage"),
		beego.NSRouter("/", &controllers.MessageController{}, "get:GetMessageList"),
		beego.NSRouter("/statistic", &controllers.MessageController{}, "get:GetMessageStatistic"),
		beego.NSRouter("/doRead", &controllers.MessageController{}, "put:ReadMessage"),
		beego.NSRouter("/delete", &controllers.MessageController{}, "delete:DeleteMessage"),
		beego.NSRouter("/getMessageById", &controllers.MessageController{}, "get:GetMessageById"),
		beego.NSRouter("/getMessageTypes", &controllers.MessageController{}, "get:GetMessageTypes"),
	)
	beego.AddNamespace(messageNS)

	auditLogNS := beego.NewNamespace("/v1/auditLogs",
		beego.NSRouter("/", &controllers.AuditLogsController{}, "get:DescribeAuditLogs"),
		beego.NSRouter("/types", &controllers.AuditLogsController{}, "get:DescribeAuditLogTypes"),
	)
	beego.AddNamespace(auditLogNS)

	//报警规则相关
	monitorRuleNs :=
		beego.NewNamespace("/v1/monitorRule",
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
		beego.NewNamespace("/v1/monitorAlert",
			//添加告警
			beego.NSRouter("/addAlert", &controllers.MonitorAlertController{}, "post:AddAlert"),
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
		beego.NewNamespace("/v1/monitorData",
			beego.NSRouter("/", &controllers.MonitorDataController{}, "get:GetMonitorData"),
		)

	beego.AddNamespace(monitorDataNs)

	//带内监控-获取agent状态、获取tag列表
	monitorProxyNs :=
		beego.NewNamespace("/v1/monitorProxy",
			beego.NSRouter("/desrcibeAgentStatus", &controllers.MonitorProxyController{}, "get:DesrcibeAgentStatus"),
			beego.NSRouter("/desrcibeTags", &controllers.MonitorProxyController{}, "get:DesrcibeTags"),
		)

	beego.AddNamespace(monitorProxyNs)

}
