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

	// beego.InsertFilter("/v1/pro/*", beego.BeforeRouter, license.CheckLicense)

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
		beego.NSRouter("/", &controllers.DeviceController{}, "post:CreateDevice"),
		beego.NSRouter("/", &controllers.DeviceController{}, "post:DescribeDevices"),
		//设备关联机型
		beego.NSRouter("/deviceType/associate", &controllers.DeviceController{}, "put:DeviceAssociateDeviceType"),
		beego.NSRouter("/?:device_id", &controllers.DeviceController{}, "put:ModifyDevice"),
		beego.NSRouter("/mount", &controllers.DeviceController{}, "put:MountDevice"),
		beego.NSRouter("/unmount", &controllers.DeviceController{}, "put:UnMountDevice"),
		beego.NSRouter("/?:device_id", &controllers.DeviceController{}, "delete:DeleteDevice"),
		beego.NSRouter("/stock", &controllers.DeviceController{}, "get:DescribeDeviceStock"),
		beego.NSRouter("/?:device_id", &controllers.DeviceController{}, "get:DescribeDevice"),
		beego.NSRouter("/?:device_id/remove", &controllers.DeviceController{}, "put:RemoveDevice"),
		// 设备详情-磁盘
		beego.NSRouter("/?:device_id/disksDetail", &controllers.DeviceController{}, "get:DescribeDeviceDisks"),
		// 关联设备磁盘
		beego.NSRouter("/disks/associateDisks", &controllers.DeviceController{}, "put:AssociateDeviceDisks"),
		// 获取设备关联的磁盘信息
		beego.NSRouter("/disks/describeAssociateDisks", &controllers.DeviceController{}, "get:GetAssociatedDisks"),
	)
	beego.AddNamespace(deviceNS)

	CommandNS := beego.NewNamespace("/commands",
		beego.NSRouter("/", &controllers.CommandController{}, "get:Query"),
		beego.NSRouter("/cancel", &controllers.CommandController{}, "put:CancelCommand"),
		beego.NSRouter("/retryCommand", &controllers.CommandController{}, "put:RetryCommand"),
	)
	beego.AddNamespace(CommandNS)

	DeviceTypeNS := beego.NewNamespace("/v1/deviceTypes",
		beego.NSRouter("/", &controllers.DeviceTypeController{}, "post:CreateDeviceType"),
		beego.NSRouter("/", &controllers.DeviceTypeController{}, "get:DescribeDeviceTypes"),
		beego.NSRouter("/?:device_type_id", &controllers.DeviceTypeController{}, "get:DescribeDeviceType"),
		beego.NSRouter("/?:device_type_id/describeVolumes", &controllers.DeviceTypeController{}, "get:DescribeVolumesByDeviceType"),
		beego.NSRouter("/?:device_type_id", &controllers.DeviceTypeController{}, "put:ModifyDeviceType"),
		beego.NSRouter("/?:device_type_id", &controllers.DeviceTypeController{}, "delete:DeleteDeviceType"),
		beego.NSRouter("/associatedImage", &controllers.DeviceTypeController{}, "post:AssociatedImage"),
		beego.NSRouter("/dissociatedImage", &controllers.DeviceTypeController{}, "delete:DissociatedImage"),
		beego.NSRouter("/deviceTypeImage", &controllers.DeviceTypeController{}, "get:DescribeDeviceTypeImages"),
		beego.NSRouter("/deviceTypeRaid", &controllers.DeviceTypeController{}, "get:DescribeDeviceTypeRaids"),
		beego.NSRouter("/deviceTypeImagePartition", &controllers.DeviceTypeController{}, "get:DescribeDeviceTypeImagePartitions"),
	)
	beego.AddNamespace(DeviceTypeNS)

	IdcNS := beego.NewNamespace("v1/idcs",
		beego.NSRouter("/devices", &controllers.IdcController{}, "post:CreateDevicesFromIdc"),
		//beego.NSRouter("/bm/deviceBatchImport", &controllers.IdcController{}, "post:DeviceBatchImport"), //ImportBms

		beego.NSRouter("/", &controllers.IdcController{}, "post:CreateIdc"),
		beego.NSRouter("/", &controllers.IdcController{}, "get:DescribeIdcs"),
		beego.NSRouter("/?:idc_id", &controllers.IdcController{}, "get:DescribeIdc"),
		beego.NSRouter("/?:idc_id", &controllers.IdcController{}, "put:ModifyIdc"),
		beego.NSRouter("/?:idc_id", &controllers.IdcController{}, "delete:DeleteIdc"),
	)

	beego.AddNamespace(IdcNS)

	OsNS := beego.NewNamespace("/v1/oss",
		beego.NSRouter("/", &controllers.OsController{}, "post:CreateOS"),
		beego.NSRouter("/", &controllers.OsController{}, "get:DescribeOSs"),
		beego.NSRouter("/?:os_id", &controllers.OsController{}, "get:DescribeOS"),
		beego.NSRouter("/?:os_id", &controllers.OsController{}, "put:ModifyOS"),
		beego.NSRouter("/?:os_id", &controllers.OsController{}, "delete:DeleteOS"),
	)
	beego.AddNamespace(OsNS)

	ImageNS := beego.NewNamespace("/v1/images",
		beego.NSRouter("/", &controllers.ImageController{}, "post:CreateImage"),
		beego.NSRouter("/", &controllers.ImageController{}, "get:DescribeImages"),
		beego.NSRouter("/?:image_id", &controllers.ImageController{}, "get:DescribeImage"),
		beego.NSRouter("/?:image_id", &controllers.ImageController{}, "put:ModifyImage"),
		beego.NSRouter("/?:image_id", &controllers.ImageController{}, "delete:DeleteImage"),
		beego.NSRouter("/imageDeviceTypes", &controllers.ImageController{}, "get:DescribeImageDeviceTypes"),
		beego.NSRouter("/associatedDeviceType", &controllers.ImageController{}, "post:AssociatedDeviceType"),
		beego.NSRouter("/dissociatedDeviceType", &controllers.ImageController{}, "delete:DissociatedDeviceType"),
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

	RaidNS := beego.NewNamespace("/v1/raids",
		beego.NSRouter("/", &controllers.RaidController{}, "post:CreateRaid"),
		beego.NSRouter("/deviceTypeRaid", &controllers.RaidController{}, "post:CreateDeviceTypeRaid"),
		beego.NSRouter("/deviceTypeRaid", &controllers.RaidController{}, "delete:DeleteDeviceTypeRaid"),
		beego.NSRouter("/", &controllers.RaidController{}, "get:DescribeRaids"), //QueryRaids
		//beego.NSRouter("/queryRaidsAll", &controllers.RaidController{}, "get:QueryRaidsAll"),
		beego.NSRouter("/?:raid_id", &controllers.RaidController{}, "get:DescribeRaid"),
		beego.NSRouter("/?:raid_id", &controllers.RaidController{}, "put:ModifyRaid"),
		beego.NSRouter("/?:raid_id", &controllers.RaidController{}, "delete:DeleteRaid"),
	)
	beego.AddNamespace(RaidNS)

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
		// 重置密码
		beego.NSRouter("/?:instance_id/resetpasswd", &controllers.InstanceController{}, "put:ResetPasswd"), //重置密码
		// //批量重置密码
		beego.NSRouter("/batch/:param(instances[:]{1}resetPasswd)", &controllers.InstanceController{}, "put:ResetInstancesPasswd"),
		// 实例重装
		beego.NSRouter("/?:instance_id/reinstallInstance", &controllers.InstanceController{}, "put:ReinstallInstance"),
	)
	beego.AddNamespace(InstanceNS)

	//offline接口，给带内监控用的
	monitorNs := beego.NewNamespace("/v1/offline",
		beego.NSRouter("/instance/queryBySn", &controllers.InstanceController{}, "get:QueryInstanceBySn"),
		// beego.NSRouter("/license/timeout", &controllers.OfflineController{}, "get:LicenseTimeout"),
	)

	beego.AddNamespace(monitorNs)

	UserNS := beego.NewNamespace("/v1/users",
		beego.NSRouter("/", &controllers.UserController{}, "post:CreateUser"),
		beego.NSRouter("/verify", &controllers.UserController{}, "post:VerifyUser"),
		beego.NSRouter("/", &controllers.UserController{}, "get:DescribeUsers"),
		beego.NSRouter("/getUserByName", &controllers.UserController{}, "get:DescribeUserByName"),
		beego.NSRouter("/?:user_id", &controllers.UserController{}, "get:DescribeUser"),
		beego.NSRouter("/?:user_id", &controllers.UserController{}, "put:ModifyUser"),
		beego.NSRouter("/?:user_id", &controllers.UserController{}, "delete:DeleteUser"),
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
		beego.NSRouter("/", &controllers.RoleController{}, "post:CreateRole"),
		beego.NSRouter("/", &controllers.RoleController{}, "get:DescribeRoles"),
		beego.NSRouter("/?:role_id", &controllers.RoleController{}, "get:DescribeRole"),
		beego.NSRouter("/roleInfo/current", &controllers.RoleController{}, "get:CurrentRole"),
		beego.NSRouter("/?:role_id", &controllers.RoleController{}, "put:ModifyRole"),
		beego.NSRouter("/?:role_id", &controllers.RoleController{}, "delete:DeleteRole"),
	)
	beego.AddNamespace(RoleSN)

	ProjectSN := beego.NewNamespace("/v1/user/projects",
		beego.NSRouter("/", &controllers.ProjectController{}, "post:CreateUserProject"),
		beego.NSRouter("/", &controllers.ProjectController{}, "get:DescribeUserProjects"),
		beego.NSRouter("/?:project_id/share", &controllers.ProjectController{}, "put:ShareUserProject"),
		beego.NSRouter("/?:project_id/cancelShare", &controllers.ProjectController{}, "put:CancelShareUserProject"),
		beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "get:DescribeUserProject"),
		beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "put:ModifyUserProject"),
		beego.NSRouter("/?:project_id", &controllers.ProjectController{}, "delete:DeleteUserProject"),
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

	CollectNS := beego.NewNamespace("/v1/collect",
		beego.NSRouter("/collectDeviceInfo", &controllers.DeviceController{}, "post:CollectDeviceInfo"),
	)
	beego.AddNamespace(CollectNS)

	// license := beego.NewNamespace("/v1/license",
	// 	beego.NSRouter("/requestToken", &controllers.LicenseController{}, "get:DescribeRequestToken"),
	// 	beego.NSRouter("/content", &controllers.LicenseController{}, "get:DescribeLicenseContent"),
	// 	beego.NSRouter("/upload", &controllers.LicenseController{}, "post:UploadLicense"),
	// 	beego.NSRouter("/currentVersion", &controllers.LicenseController{}, "get:DescribeCurrentVersion"),
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
		beego.NSRouter("/dialMail", &controllers.MessageController{}, "post:DialMail"),
		beego.NSRouter("/saveIsPushMail", &controllers.MessageController{}, "post:SaveIsPushMail"),
		beego.NSRouter("/mailDetail", &controllers.MessageController{}, "get:MailDetail"),
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
			beego.NSRouter("/deleteRule", &controllers.MonitorRuleController{}, "delete:EditRule"),
		)

	beego.AddNamespace(monitorRuleNs)

	//监控报警相关
	monitorAlertNs :=
		beego.NewNamespace("/v1/monitorAlert",
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

	//内部测试用
	innerNs :=
		beego.NewNamespace("/v1/innerTest",
			beego.NSRouter("/mockMultiAuditLogs", &controllers.AuditLogsController{}, "post:MockMultiAuditLogs"),
		)

	beego.AddNamespace(innerNs)

}
