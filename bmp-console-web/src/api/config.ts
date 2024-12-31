interface TotalUrlPathTypes {
    loginUrl: string;
    userUrl: string;
    timeZoneUrl: string;
    instanceCloseUrl: string;
    instanceDeleteUrl: string;
    instanceOpenUrl: string;
    instanceRestartUrl: string;
    instanceLockUrl: string;
    instanceUnlockUrl: string;
    instanceRepasswordUrl: string;
    apiKeyCreateUrl: string;
    apiKeyDeleteUrl: string;
    sshKeyDeleteUrl: string;
    sshKeyCreateUrl: string;
    sshKeyEditUrl: string;
    projectCreateUrl: string;
    projectEditUrl: string;
    projectDeleteUrl: string;
    azUrl: string;
    deviceTypeUrl: string;
    osTypeUrl: string;
    raidUrl: string;
    systemDiskPartitionUrl: string;
    isDeviceStockEnoughUrl: string;
    createInstanceUrl: string;
    instanceListUrl: string;
    instanceDetailUrl: string;
    setUserUrl: string;
    changePasswordUrl: string;
    sshKeyUrl: string;
    apiKeyUrl: string;
    projectListUrl: string;
    logOutUrl: string;
    setCustomListUrl: string;
    setAlarmCustomListUrl: string;
    customListUrl: string;
    alarmCustomListUrl: string;
    editInstanceUrl: string;
    editProjectUrl: string;
    projectDetailUrl: string;
    relatedUserAddUrl: string;
    relatedUserDeleteUrl: string;
    transferUserUrl: string;
    transferUserCheckUrl: string;
    relatedUserCheckUrl: string;
    batchCloseUrl: string;
    batchOpenUrl: string;
    batchRestartUrl: string;
    batchDeleteUrl: string;
    batchRenameUrl: string;
    batchRepasswordUrl: string;
    hardWareStatusUrl: string;
    alarmLogListUrl: string;
    alarmLogListExportUrl: string;
    resystemInstanceUrl: string;
    shareInstanceListUrl: string;
    transferPartUserUrl: string;
    messageListUrl: string;
    messageDetailUrl: string;
    messageTypeUrl: string;
    messageReadUrl: string;
    getMessageStatisticUrl: string;
    messageUnreadUrl: string;
    messageDeleteUrl: string;
    operationLogTypeListUrl: string;
    operationLogListUrl: string;
    operationLogListExportUrl: string;
    deviceTagUrl: string;
    addRuleUrl: string;
    alarmRuleListUrl: string;
    alarmRuleDetailUrl: string;
    ruleEnableUrl: string;
    ruleDisableUrl: string;
    ruleDeleteUrl: string;
    editRuleUrl: string;
    getAgentStateUrl: string;
    alarmHistoryUrl: string;
    monitorDataUrl: string;
};

const urlPath: TotalUrlPathTypes =  {
    loginUrl: '/login', // 登录
    userUrl: '/user', // 用户详情
    timeZoneUrl: '/user/timezones', // 时区信息
    instanceCloseUrl: '/instance/stopInstance', // 关机
    instanceDeleteUrl: '/instance/', // 删除
    instanceOpenUrl: '/instance/startInstance', // 开机
    instanceRestartUrl: '/instance/restartInstance', // 重启
    instanceLockUrl: '/instance/lockInstance', // 锁定
    instanceUnlockUrl: '/instance/unlockInstance', // 解除锁定
    instanceRepasswordUrl: '/instance/resetPasswd', // 重置密码
    apiKeyCreateUrl: '/apikey', // api密钥创建
    apiKeyDeleteUrl: '/apikey/', // api密钥删除 
    sshKeyDeleteUrl: '/keypair/', // ssh密钥删除
    sshKeyCreateUrl: '/keypair', // ssh密钥创建
    sshKeyEditUrl: '/keypair/', // ssh密钥编辑
    projectCreateUrl: '/project',  // 项目创建
    projectEditUrl: '/project/',  // 项目编辑
    projectDeleteUrl: '/project/',  // 项目删除
    azUrl: '/idc', // 机房信息
    deviceTypeUrl: '/deviceType/getAvailableDeviceTypes', // 机型信息
    osTypeUrl: '/image/queryImagesByDeviceType', // 镜像信息
    raidUrl: '/raid/queryRaidsByDeviceTypeIDAndVolumeType', // 系统盘信息
    systemDiskPartitionUrl: '/partition/queryDefaultSystemPartitions', // 系统盘分区信息
    isDeviceStockEnoughUrl: '/device/isDeviceStockEnough', // 查询库存
    createInstanceUrl: '/instance', // 创建实例
    instanceListUrl: '/instance', // 实例列表
    instanceDetailUrl: '/instance/', // 实例详情
    setUserUrl: '/user', // 用户编辑
    changePasswordUrl: '/user/password', // 用户修改密码
    sshKeyUrl: '/keypair', // ssh密钥列表
    apiKeyUrl: '/apikey', // api密钥列表
    projectListUrl: '/project', // 获取项目列表
    logOutUrl: '/logout',  // 登出
    setCustomListUrl: '/custom/setCustomInfo', // 设置自定义列表
    setAlarmCustomListUrl: '/v1/oob-alert/custom/set-custom-info',
    customListUrl: '/custom/getCustomInfo', // 获取自定义列表
    alarmCustomListUrl: '/v1/oob-alert/custom/get-custom-info', // 获取报警日志自定义列表
    editInstanceUrl: '/instance/modifyInstance', // 实例编辑描述
    editProjectUrl: '/project/', //项目编辑描述
    projectDetailUrl: '/project/', //项目详情
    relatedUserAddUrl: '/project/', //添加共享用户
    relatedUserDeleteUrl: '/project/', //删除共享用户
    transferUserUrl: '/project/', //转移用户
    transferUserCheckUrl: '/user/checkUserConsoleAccessByName', //验证转移用户
    relatedUserCheckUrl: '/user/checkUserConsoleAccessByName', //验证共享用户
    batchCloseUrl: '/instance/stopInstances', //批量关机实例
    batchOpenUrl: '/instance/startInstances', //批量开机实例
    batchRestartUrl: '/instance/restartInstances', //批量重启实例
    batchDeleteUrl: '/instance/deleteInstances', //批量删除实例
    batchRenameUrl: '/instance/modifyInstances', //批量编辑实例名称实例
    batchRepasswordUrl: '/instance/batchResetPasswd', //批量重置密码实例
    hardWareStatusUrl: '/v1/oob-alert/device/status', //硬件设备状态
    alarmLogListUrl: '/v1/oob-alert/consolelogs', //报警日志列表
    alarmLogListExportUrl: '/v1/oob-alert/consolelogs', //报警日志导出列表
    resystemInstanceUrl: '/instance/reinstallInstance', //实例重装系统
    shareInstanceListUrl: '/instance/forshare/list', //共享实例列表
    transferPartUserUrl: '/project/move/instances', //转移部分实例
    messageListUrl: '/messages', // 消息列表
    messageDetailUrl: '/messages/getMessageById', // 消息详情
    messageTypeUrl: '/messages/getMessageTypes', // 消息类型
    messageReadUrl: '/messages/doRead', // 消息已读
    getMessageStatisticUrl: '/messages/statistic', // 消息未读数
    messageUnreadUrl: '/messages/hasUnreadMessage', // 消息是否有未读
    messageDeleteUrl: '/messages/delete', // 消息删除
    operationLogTypeListUrl: '/auditLogs/types', // 操作日志类型
    operationLogListUrl: '/auditLogs', // 操作日志列表
    operationLogListExportUrl: '/auditLogs', // 操作日志下载
    deviceTagUrl: '/monitorProxy/desrcibeTags', // devicetag列表
    addRuleUrl: '/monitorRule/addRule', // 添加报警规则
    alarmRuleListUrl: '/monitorRule/describeRules', // 报警规则列表
    alarmRuleDetailUrl: '/monitorRule/describeRule', // 报警规则详情
    ruleEnableUrl: '/monitorRule/enableRule', // 启用报警规则
    ruleDisableUrl: '/monitorRule/disableRule', // 禁用报警规则
    ruleDeleteUrl: '/monitorRule/deleteRule', // 删除报警规则
    editRuleUrl: '/monitorRule/editRule', // 编辑报警规则
    getAgentStateUrl: '/monitorProxy/desrcibeAgentStatus', // 获取agent状态
    alarmHistoryUrl: '/monitorAlert/describeAlerts', // 报警历史列表
    monitorDataUrl: '/monitorData', // 带内监控
}


export default urlPath;
