package AuditLogsType

const (
	//操作消息类型

	AuditLogsPutaway               = "PutawayDevice"         //上架
	AuditLogsUnPutaway             = "UnPutawayDevice"       //下架
	AuditLogsStartInstances        = "StartInstances"        //开机
	AuditLogsStopInstances         = "StopInstances"         //关机
	AuditLogsRestartInstances      = "RestartInstances"      //重启
	AuditLogsResetInstanceStatus   = "ResetInstances"        //重置实例状态（新增）
	AuditLogsRemoveDevice          = "RemoveDevice"          //设备移除（新增）
	AuditLogsDeleteDevice          = "DeleteDevice"          //设备删除（新增）
	AuditLogsDeleteInstance        = "DeleteInstance"        //回收实例
	AuditLogsInstanceResetPassword = "InstanceResetPassword" //重置密码
	AuditLogsReinstallInstance     = "ReinstallInstance"     //重装系统
	AuditLogsEditInstanceName      = "EditInstanceName"      //编辑实例名称（新增）
	AuditLogsCollectHardwareInfo   = "CollectHardwareInfo"   //设备采集
	AuditLogsTypeCreateInstances   = "CreateInstances"       //创建实例
	AuditLogsTypeDropInstances     = "DropInstances"         //删除实例(不走销毁流程，新增)
	AuditLogsTypeLockInstances     = "LockInstances"         //锁定实例(新增)

)
