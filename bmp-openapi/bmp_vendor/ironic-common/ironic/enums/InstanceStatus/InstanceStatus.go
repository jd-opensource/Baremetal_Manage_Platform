package InstanceStatus

const (
	/**
	 * 创建中
	 */
	CREATING = "creating"

	/**
	 * 开机中
	 */
	STARTING = "starting"

	/**
	 * 运行
	 */
	RUNNING = "running"

	/**
	 * 关机中
	 */
	STOPPING = "stopping"

	/**
	 * 关机
	 */
	STOPPED = "stopped"

	/**
	 * 重启中
	 */
	RESTARTING = "restarting"

	/**
	 * 重置密码中
	 */
	RESETTING_PASSWORD = "resetting_password"

	/**
	 * 销毁中
	 */
	DESTROYING = "destroying"

	/**
	 * 已销毁
	 */
	DESTROYED = "destroyed"

	/**
	 * 错误
	 */
	ERROR = "error"

	/**
	 * 调整配置中
	 */
	UPGRADING = "upgrading"

	/**
	 * 重装系统中
	 */
	REINSTALLING = "reinstalling"

	// 创建失败
	CREATE_ERROR = "Creation failed"
	// 开机失败
	START_ERROR = "Startup failed"
	// 关机失败
	STOP_ERROR = "Shutdown failed"
	// 重启失败
	RESTART_ERROR = "Reboot failed"
	// 销毁失败
	DESTROY_ERROR = "Delete failed"
	// 重装失败
	REINSTALL_ERROR = "Reinstall failed"
	//重置密码失败
	RESETPASSWD_ERROR = "Resetpasswd failed"
)

var Instance_To_Error_Status map[string]string = map[string]string{
	// "CollectHardwareInfo" : NewCollectHardwareInfoActor()
	"CreateInstances": CREATE_ERROR,
	// "CutDeviceStock" : NewCutDeviceStockActor()
	"DeleteInstance": DESTROY_ERROR,
	// "InstanceResetPassword" : NewInstanceResetPasswordActor()
	// "ModifyBandwidth" : NewModifyBandwidthActor()
	// "ReinstallInstance" : NewReinstallInstanceActor()
	// "ResetSwitchConfig" : NewResetSwitchConfigActor()
	// "RestartDhcp" : NewRestartDhcpActor()
	"RestartInstances": RESTART_ERROR,
	"StartInstances":   START_ERROR,
	"StopInstances":    STOP_ERROR,
	"ReinstallInstance":     REINSTALL_ERROR,
	"InstanceResetPassword": RESETPASSWD_ERROR,
}

var Error_Instance_To_Reset_Status map[string]string = map[string]string{
	START_ERROR:   STOPPED,
	STOP_ERROR:    RUNNING,
	RESTART_ERROR: RUNNING,
	DESTROY_ERROR: STOPPED,
}
