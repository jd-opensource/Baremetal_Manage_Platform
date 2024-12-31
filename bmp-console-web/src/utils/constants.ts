import i18n from 'lib/i18n/index.ts'; // 国际化
/**
 * 国际化
*/
const {global} = i18n;

/**
 * 自定义列表类
*/
type optionArrType = {
    name: string;
    selected: boolean;
    disabled: boolean;
    label?: string;
};

type radioArrType = {
    label: string;
    value: string;
};

/**
 * 通用状态类
*/
type CurrencyStatusType = {
    required: boolean;
    selected: boolean;
};

/**
 * 通用状态
*/
const currencyStatus: CurrencyStatusType = {
    required: false,
    selected: true
};

/**
 * 禁止更改状态
*/
const prohibitChangesStatus: CurrencyStatusType = {
    required: true,
    selected: true
};

/**
 * 自定义数据信息类
*/
type InfoType = Record<string, Required<CurrencyStatusType>>;

/**
 * 通用类-列表
*/
type CurrencyType = {
    text: string;
    label: string | number;
    info: any;
    filterParams: string;
};

/**
 * 类集合
*/
 interface StaticDataType {
    NormalDataType: {
        value: number | string;
        text?: string;
        label?: string | number;
        id?: string | number;
    };
    porjectCustomType: InfoType;
    alarmLogCustomType: InfoType;
    operationCustomType: InfoType;
    CurrencyListType: Required<{value: number | string;}> & Partial<CurrencyType>;
    
};

/**
 * 手机号
*/
export const cellPhoneType: StaticDataType['NormalDataType'][] = [
    {
        value: '086',
        label: global.t('personCentre.phoneData.chineseMainland') // 中国大陆
    },
    {
        value: '852',
        label: global.t('personCentre.phoneData.chineseHongKong') // 中国香港
    },
    {
        value: '853',
        label: global.t('personCentre.phoneData.chineseMacao') // 中国澳门
    },
    {
        value: '886',
        label: global.t('personCentre.phoneData.chineseTaiwan') // 中国台湾
    },
    {
        value: '54',
        label: global.t('personCentre.phoneData.Argentina') // 阿根廷
    },
    {
        value: '61',
        label: global.t('personCentre.phoneData.Australia') // 澳大利亚
    },
    {
        value: '55',
        label: global.t('personCentre.phoneData.Brazil') // 巴西
    },
    {
        value: '49',
        label: global.t('personCentre.phoneData.Germany') // 德国
    },
    {
        value: '33',
        label: global.t('personCentre.phoneData.France') // 法国
    },
    {
        value: '82',
        label: global.t('personCentre.phoneData.SouthKorea') // 韩国
    },
    {
        value: '01',
        label: global.t('personCentre.phoneData.Canada') // 加拿大
    },
    {
        value: '1',
        label: global.t('personCentre.phoneData.USA') // 美国
    },
    {
        value: '52',
        label: global.t('personCentre.phoneData.Mexico') // 墨西哥
    },
    {
        value: '27',
        label: global.t('personCentre.phoneData.SouthAfrica') // 南非
    },
    {
        value: '81',
        label: global.t('personCentre.phoneData.Japan') // 日本
    },
    {
        value: '966',
        label: global.t('personCentre.phoneData.SaudiArabia') // 沙特阿拉伯
    },
    {
        value: '90',
        label: global.t('personCentre.phoneData.Turkey') // 土耳其
    },
    {
        value: '39',
        label: global.t('personCentre.phoneData.Italy') // 意大利
    },
    {
        value: '91',
        label: global.t('personCentre.phoneData.India') // 印度
    },
    {
        value: '62',
        label: global.t('personCentre.phoneData.Indonesia') // 印度尼西亚
    },
    {
        value: '44',
        label: global.t('personCentre.phoneData.UnitedKiongdom') // 英国
    }
];

/**
 * 语言
*/
export const languageType: StaticDataType['NormalDataType'][] = [
    {
        value: 'zh_CN',
        label: global.t('personCentre.language.chinese') // 中文
    },{
        value: 'en_US',
        label: 'English' // 英语
    },

];
/**
 * api权限
*/
export const apiKeyType: StaticDataType['NormalDataType'][] = [
    {
        value: 0,
        label: global.t('personCentre.apiPower.writeAndRead') // 读/写
    },
    {
        value: 1,
        label: global.t('personCentre.apiPower.read') // 只读
    },


];
/**
 * api创建权限
*/
export const apiKeyCreateType: StaticDataType['NormalDataType'][] = [
    {
        value: 0,
        label: global.t('personCentre.apiPower.writeAndRead') // 读/写
    },
    // {
    //     value: 1,
    //     label: global.t('personCentre.apiPower.read') // 只读
    // },


];
/**
 * 添加共享用户权限
*/
export const relateUserPowerType: StaticDataType['NormalDataType'][] = [
    {
        value: 0,
        label: global.t('personCentre.apiPower.writeAndRead') // 读/写
    },
    // {
    //     value: 1,
    //     label: global.t('personCentre.apiPower.read') // 只读
    // },


];
/**
 * 中间态
*/
export const intermediate: string[] = ['creating', 'starting', 'stopping', 'restarting', 'resetting_password', 'destroying', 'upgrading', 'reinstalling','Creating', 'Starting', 'Stopping', 'Restarting', 'Resetting_password', 'Destroying', 'Upgrading', 'Reinstalling'];

/**
 * 镜像类型
*/
export const imageType: string[] = ['CentOS', 'Ubuntu', 'Windows', 'Debian', 'OpenEuler'];

/**
 *  项目列表自定义信息
*/
export const porjectCustomIntro: StaticDataType['porjectCustomType'] = {
    projectName: currencyStatus, // 项目名称
    projectId: prohibitChangesStatus, // id
    owned: currencyStatus, // 项目协作状态
    instanceCount: currencyStatus, // 实例数
    createdTime: currencyStatus, // 创建时间
    updatedTime: currencyStatus, // 更新时间
    operation: prohibitChangesStatus, // 操作

};
/**
 * 项目列表自定义项
 */
export const porjectCustomInfo: optionArrType[]= [
    {
        // 项目名称
        name: global.t('list.project.projectName'),
        selected: true,
        disabled: false,
        label: 'projectName'
    },
    {
        // id
        name: global.t('list.project.projectId'),
        selected: true,
        disabled: true,
        label: 'projectId'
    },
    {
        // 项目协作状态
        name: global.t('list.project.relatedState'),
        selected: true,
        disabled: false,
        label: 'owned'
    },
    {
        // 实例数
        name: global.t(('list.project.instanceNumber')),
        selected: true,
        disabled: false,
        label: 'instanceCount'
    },
    {
        // 创建时间
        name: global.t('personCentre.form.createdTime'),
        selected: true,
        disabled: false,
        label: 'createdTime'

    },
    {
        // 更新时间
        name: global.t('personCentre.form.updatedTime'),
        selected: true,
        disabled: false,
        label: 'updatedTime'

    },
    {
        // 操作
        name: global.t('personCentre.form.operation'),
        selected: true,
        disabled: true,
        label: 'operation'
    }
];

/**
 *  项目列表自定义信息
*/
export const instanceCustomIntro: StaticDataType['porjectCustomType'] = {
    instanceName: prohibitChangesStatus, // 实例名称
    monitor: prohibitChangesStatus, // 监控
    status: currencyStatus, // 状态
    deviceTypeName: currencyStatus, // 机型
    imageName: currencyStatus, // 镜像
    iloIp: currencyStatus, // 带外IP
    privateIpv4: currencyStatus, // 内网IPv4(eth0)
    privateEth1Ipv4: currencyStatus, // 内网IPv4(eth1)
    privateIpv6: currencyStatus, // 内网IPv6(eth0)
    privateEth1Ipv6: currencyStatus, // 内网IPv6(eth1)
    config: currencyStatus, // 配置
    createdTime: currencyStatus, // 创建时间
    operation: prohibitChangesStatus, // 操作

};

/**
 * 实例列表自定义项
 */
 export const instanceCustomInfo: optionArrType[]= [
    {
        // 实例名称/ID
        name: global.t('instance.list.instanceName'),
        selected: true,
        disabled: true,
        label: 'instanceName',
    },
    {
        // 监控
        name: global.t('alarm.monitor.monitor'),
        selected: true,
        disabled: true,
        label: 'monitor',
    },
    {
        // 状态
        name: global.t('instance.list.state'),        
        selected: true,
        disabled: false,
        label: 'status',
    },
    {
        // 机型
        name: global.t(('instance.list.model')),       
        selected: true,
        disabled: false,
        label: 'deviceTypeName',
    },
    {
        // 镜像
        name: global.t(('instance.list.mirror')), 
        selected: true,
        disabled: false,
        label: 'imageName',
    },
    {
        // 带外IP
        name: global.t(('instance.list.outBandIp')),
        selected: true,
        disabled: false,
        label: 'iloIp',
    },
    {
        // 内网IPv4(eth0)
        name: global.t('instance.list.intranetIpv4'),
        selected: true,
        disabled: false,
        label: 'privateIpv4',

    },
    {
        // 内网IPv4(eth1)
        name: global.t('instance.list.intranetIpv4eth1'),
        selected: true,
        disabled: false,
        label: 'privateEth1Ipv4',

    },
    {
        // 内网IPv6(eth0)
        name: global.t('instance.list.intranetIpv6'),
        selected: true,
        disabled: false,
        label: 'privateIpv6',

    },
    {
        // 内网IPv6(eth1)
        name: global.t('instance.list.intranetIpv6eth1'),
        selected: true,
        disabled: false,
        label: 'privateEth1Ipv6',

    },
    {
        // 配置
        name: global.t('instance.list.config'),
        selected: true,
        disabled: false,
        label: 'config',

    },
    {
        // 创建时间
        name: global.t('instance.list.createdTime'),
        selected: true,
        disabled: false,
        label: 'createdTime',

    },
    {
        // 操作
        name: global.t('personCentre.form.operation'),
        selected: true,
        disabled: true,
        label: 'operation',
    }
];

/**
 * 管理状态
*/
export const manageMentStatus: StaticDataType['CurrencyListType'][] = [ // 管理状态filter数据
    {
        filterParams: 'running',
        text: global.t('list.status.running'),
        value: 1
    },
    {
        filterParams: 'stopped',
        text: global.t('list.status.stopped'),
        value: 2
    },
    {
        filterParams: 'Creation failed',
        text: global.t('list.status.creationFailed'),
        value: 3
    },
    
];

/**
 * 锁定状态
 */
export const lockedStatus = {
    'locked': global.t('list.status.locked'),
    'unlocked': global.t('list.status.unlocked')
}

/**
 * 项目协作状态
*/
export const projectRelatedStatus: StaticDataType['CurrencyListType'][] = [ // 管理状态filter数据
    {
        filterParams: '1',
        text: global.t('list.project.owner'),
        value: 1
    },
    {
        filterParams: '2',
        text: global.t('list.project.cooperate'),
        value: 2
    }
    
];

/**
 * 批量操作提示
 */
 export const batchOperationTip = {
    'close': global.t('list.content.bacthCloseContent'),
    'open': global.t('list.content.bacthOpenContent'),
    'restart': global.t('list.content.bacthRestartContent'),
    'resetPassword': global.t('list.content.bacthResetPasswordContent'),
    'rename': global.t('list.content.bacthRenameContent'),
    'delete': global.t('list.content.bacthDeleteContent')
}

/**
 *  报警日志列表自定义信息
*/
export const alarmLogCustomIntro: StaticDataType['alarmLogCustomType'] = {
    logid: prohibitChangesStatus, // 序号
    alert_type: currencyStatus, // 故障类型
    alert_desc: currencyStatus, // 故障描述
    event_time: currencyStatus, // 故障报警时间
    update_time: currencyStatus, // 更新时间
    count: currencyStatus, // 报警次数
    status: currencyStatus, // 状态

};

/**
 * 报警日志列表自定义项
 */
 export const alarmLogCustomInfo: optionArrType[]= [
    {
        // 序号
        name: global.t('instance.detail.number'),
        selected: true,
        disabled: true,
        label: 'logid',
    },
    {
        // 故障类型
        name: global.t('instance.detail.faultType'),        
        selected: true,
        disabled: false,
        label: 'alertType',
    },
    {
        // 故障描述
        name: global.t(('instance.detail.faultDesc')),       
        selected: true,
        disabled: false,
        label: 'alertDesc',
    },
    {
        // 故障报警时间
        name: global.t(('instance.detail.faultAlarmTime')), 
        selected: true,
        disabled: false,
        label: 'eventTime',
    },
    {
        // 更新时间
        name: global.t(('instance.detail.updatedTime')),
        selected: true,
        disabled: false,
        label: 'updateTime',
    },
    {
        // 报警次数
        name: global.t('instance.detail.alarmNumbers'),
        selected: true,
        disabled: false,
        label: 'count',

    },
    {
        // 状态
        name: global.t('instance.detail.status'),
        selected: true,
        disabled: false,
        label: 'status',

    }
];

/**
 * 报警日志状态
 */
 export const alarmStatusConstant = {
    0: global.t('instance.detail.unhandled'), /* 未处理 */
    1: global.t('instance.detail.recovered'), /* 已恢复 */
    2: global.t('instance.detail.ignored'), /* 已忽略 */
}

/**
 * 消息列表分组
 */
 export const groupingMessageListConstant: radioArrType[] = [
    {
        label: global.t('messageCentre.list.all'),
        value: '2'
    },{
        label: global.t('messageCentre.list.unread'),
        value: '0'
    },{
        label: global.t('messageCentre.list.read'),
        value: '1'
    }
 ]

 /**
 *  消息列表自定义信息
*/
export const messageCustomIntro: StaticDataType['porjectCustomType'] = {
    detail: prohibitChangesStatus, // 消息内容    
    finish_time: currencyStatus, // 接收时间
    message_type: currencyStatus, // 消息类型
    message_sub_type: currencyStatus, // 消息子类型
};

/**
 * 消息列表自定义项
 */
 export const messageCustomInfo: optionArrType[]= [
    {
        // 消息内容
        name: global.t('messageCentre.list.messageContent'),
        selected: true,
        disabled: true,
        label: 'detail',
    },
    {
        // 接受时间
        name: global.t('messageCentre.list.receivingTime'),        
        selected: true,
        disabled: false,
        label: 'finish_time',
    },
    {
        // 消息类型
        name: global.t(('messageCentre.list.messageType')),       
        selected: true,
        disabled: false,
        label: 'message_type',
    },
    {
        // 消息子类型
        name: global.t(('messageCentre.list.messageSubtype')), 
        selected: true,
        disabled: false,
        label: 'message_sub_type',
    },
];

/**
 * 通知状态
 */
 export const messageContentConstant = {
    'failed': global.t('messageCentre.detail.faultContent'), /* 失败 */
}

/**
 *  操作日志列表自定义信息
*/
export const operationCustomIntro: StaticDataType['operationCustomType'] = {
    logId: prohibitChangesStatus, // 序号
    operation: currencyStatus, // 操作名称
    username: currencyStatus, // 操作人
    operationTime: currencyStatus, // 操作时间
    operationFeedback: currencyStatus, // 操作反馈
    faultWord: currencyStatus, // 错误码

}

/**
 * 操作日志列表自定义项
 */
 export const operationCustomInfo: optionArrType[]= [
    {
        // LogID
        name: global.t('instance.detail.logId'),
        selected: true,
        disabled: true,
        label: 'logid',
    },
    {
        // 操作名称
        name: global.t('instance.detail.operationName'),        
        selected: true,
        disabled: false,
        label: 'operationName',
    },
    {
        // 操作人
        name: global.t(('instance.detail.operator')),       
        selected: true,
        disabled: false,
        label: 'userName',
    },
    {
        // 操作时间
        name: global.t(('instance.detail.operationTime')), 
        selected: true,
        disabled: false,
        label: 'operateTime',
    },
    {
        // 操作反馈
        name: global.t(('instance.detail.operationFeedback')),
        selected: true,
        disabled: false,
        label: 'result',
    },
    {
        // 错误码
        name: global.t('instance.detail.faultWord'),
        selected: true,
        disabled: false,
        label: 'failReason',

    }
];


/**
 * 操作反馈状态
*/
export const operationLogStatus: StaticDataType['CurrencyListType'][] = [ // 操作反馈filter数据
    {
        filterParams: 'success',
        text: global.t('personCentre.content.operateSuccess'),
        value: 'success'
    },
    {
        filterParams: 'fail',
        text: global.t('personCentre.content.operateFailed'),
        value: 'fail'
    },   
    {
        filterParams: 'doing',
        text: global.t('personCentre.content.operateDoing'),
        value: 'doing'
    }, 
];

/**
 * 操作反馈状态
 */
 export const operationLogStatusConstant = {
    'success': global.t('personCentre.content.operateSuccess'), /* 操作成功 */
    'fail': global.t('personCentre.content.operateFailed'), /* 操作失败 */
    'doing': global.t('personCentre.content.operateDoing'), /* 操作中 */
}

/**
 *  报警规则列表自定义信息
*/
export const alarmRuleCustomIntro: StaticDataType['operationCustomType'] = {
    ruleName: prohibitChangesStatus, // 序号
    ruleId: currencyStatus, // 操作名称
    resourceName: currencyStatus, // 操作人
    instanceCount: currencyStatus, // 操作时间
    triggerDescription: currencyStatus, // 操作反馈
    statusName: currencyStatus, // 错误码
    operation: prohibitChangesStatus, // 操作

}

/**
 * 报警规则列表自定义项
 */
 export const alarmRuleCustomInfo: optionArrType[]= [
    {
        // 规则名称
        name: global.t('alarm.alarmRule.ruleName'),
        selected: true,
        disabled: true,
        label: 'ruleName',
    },
    {
        // 规则ID
        name: global.t('alarm.alarmRule.ruleId'),        
        selected: true,
        disabled: false,
        label: 'ruleId',
    },
    {
        // 资源类型
        name: global.t(('alarm.alarmRule.resourceType')),       
        selected: true,
        disabled: false,
        label: 'resourceName',
    },
    {
        // 关联实例类型
        name: global.t(('alarm.alarmRule.instanceNumber')), 
        selected: true,
        disabled: false,
        label: 'instanceCount',
    },
    {
        // 触发条件
        name: global.t(('alarm.alarmRule.triggeringCondition')),
        selected: true,
        disabled: false,
        label: 'triggerDescription',
    },
    {
        // 状态
        name: global.t('alarm.alarmRule.status'),
        selected: true,
        disabled: false,
        label: 'statusName',

    },
    {
        // 操作
        name: global.t('personCentre.form.operation'),
        selected: true,
        disabled: true,
        label: 'operation'
    }
];

/**
 * 报警资源类型
*/
export const resourceData: StaticDataType['CurrencyListType'][] = [ // 管理状态filter数据
    {
        text: global.t('instance.detail.instance'),
        value: 'instance'
    }
    
];

/**
 * 维度下拉状态
 */
 export const dimentionState = {
    'disk': global.t('alarm.alarmRule.disk'), 
    'mountpoint': global.t('alarm.alarmRule.mountpoint'),
    'nic': global.t('alarm.alarmRule.nic'),

}

/**
 * 实例监控指标类型
*/
export const metricInstanceData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.cpuUtil'),
        value: 'bmp.cpu.util'
    },
    {
        text: global.t('alarm.alarmRule.memoryUsage'),
        value: 'bmp.memory.util'
    },
    {
        text: global.t('alarm.alarmRule.memoryUsed'),
        value: 'bmp.memory.used'
    },
    {
        text: global.t('alarm.alarmRule.diskUsed'),
        value: 'bmp.disk.used'
    },
    {
        text: global.t('alarm.alarmRule.diskUsage'),
        value: 'bmp.disk.util'
    },
    {
        text: global.t('alarm.alarmRule.diskReadTraffic'),
        value: 'bmp.disk.bytes.read'
    },
    {
        text: global.t('alarm.alarmRule.diskWriteTraffic'),
        value: 'bmp.disk.bytes.write'
    },
    {
        text: global.t('alarm.alarmRule.diskReadIOPS'),
        value: 'bmp.disk.counts.read'
    },
    {
        text: global.t('alarm.alarmRule.diskWriteIOPS'),
        value: 'bmp.disk.counts.write'
    },
    {
        text: global.t('alarm.alarmRule.networkIngressTraffic'),
        value: 'bmp.network.bytes.ingress'
    },
    {
        text: global.t('alarm.alarmRule.networkEgressTraffic'),
        value: 'bmp.network.bytes.egress'
    },
    {
        text: global.t('alarm.alarmRule.networkIngressPackets'),
        value: 'bmp.network.packets.ingress'
    },
    {
        text: global.t('alarm.alarmRule.networkEgressPackets'),
        value: 'bmp.network.packets.egress'
    },
    {
        text: global.t('alarm.alarmRule.loadAverage1min'),
        value: 'bmp.avg.load1'
    },
    {
        text: global.t('alarm.alarmRule.loadAverage5min'),
        value: 'bmp.avg.load5'
    },
    {
        text: global.t('alarm.alarmRule.loadAverage15min'),
        value: 'bmp.avg.load15'
    },
    {
        text: global.t('alarm.alarmRule.totalTCPConnections'),
        value: 'bmp.tcp.connect.total'
    },
    {
        text: global.t('alarm.alarmRule.establishedTCPConnections'),
        value: 'bmp.tcp.connect.established'
    },
    {
        text: global.t('alarm.alarmRule.totalProcessCount'),
        value: 'bmp.process.total'
    },
    
    
];

/**
 * 磁盘监控指标类型
*/
export const metricDiskData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.diskUsed'),
        value: 'bmp.disk.used'
    },
    {
        text: global.t('alarm.alarmRule.diskUsage'),
        value: 'bmp.disk.util'
    },
    {
        text: global.t('alarm.alarmRule.diskReadTraffic'),
        value: 'bmp.disk.bytes.read'
    },
    {
        text: global.t('alarm.alarmRule.diskWriteTraffic'),
        value: 'bmp.disk.bytes.write'
    },
    {
        text: global.t('alarm.alarmRule.diskReadIOPS'),
        value: 'bmp.disk.counts.read'
    },
]

/**
 * 磁盘用量监控指标类型
*/
export const metricMountpointData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.diskUsedMountPoint'),
        value: 'bmp.disk.partition.used'
    },
    {
        text: global.t('alarm.alarmRule.diskUsageMountPoint'),
        value: 'bmp.disk.partition.util'
    }
]

/**
 * 网口监控指标类型
*/
export const metricNetworkData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.networkIngressTraffic'),
        value: 'bmp.network.bytes.ingress'
    },
    {
        text: global.t('alarm.alarmRule.networkEgressTraffic'),
        value: 'bmp.network.bytes.egress'
    },
    {
        text: global.t('alarm.alarmRule.networkIngressPackets'),
        value: 'bmp.network.packets.ingress'
    },
    {
        text: global.t('alarm.alarmRule.networkEgressPackets'),
        value: 'bmp.network.packets.egress'
    },
]

/**
 * 周期类型
*/
export const periodData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.oneM'),
        value: 1
    },
    {
        text: global.t('alarm.alarmRule.twoM'),
        value: 2
    },
    {
        text: global.t('alarm.alarmRule.fiveM'),
        value: 5
    },
    {
        text: global.t('alarm.alarmRule.fifteenM'),
        value: 15
    },
    {
        text: global.t('alarm.alarmRule.halfH'),
        value: 30
    },
    {
        text: global.t('alarm.alarmRule.oneH'),
        value: 60
    },
]

/**
 * 计算方式类型
*/
export const calculationData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.max'),
        value: 'max'
    },
    {
        text: global.t('alarm.alarmRule.min'),
        value: 'min'
    },
    {
        text: global.t('alarm.alarmRule.avg'),
        value: 'avg'
    },
    {
        text: global.t('alarm.alarmRule.sum'),
        value: 'sum'
    },
]

/**
 * 比较类型
*/
export const operationData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: '>',
        value: 'gt'
    },
    {
        text: '>=',
        value: 'gte'
    },
    {
        text: '<',
        value: 'lt'
    },
    {
        text: '<=',
        value: 'lte'
    },
    {
        text: '==',
        value: 'eq'
    },
    {
        text: '!=',
        value: 'neq'
    },
]

/**
 * 计算单位类型
*/
export const calculationUnitData: StaticDataType['NormalDataType'][] = [
    { id: 'bmp.cpu.util', value: '%', text: '%' },
    { id: 'bmp.memory.util', value: '%', text: '%' },
    { id: 'bmp.memory.used', value: 'GB', text: 'GB' },
    { id: 'bmp.disk.used', value: 'GB', text: 'GB' },
    { id: 'bmp.disk.util', value: '%', text: '%' },
    { id: 'bmp.disk.bytes.read', value: 'Bps', text: 'Bps' },
    { id: 'bmp.disk.bytes.write', value: 'Bps', text: 'Bps' },
    { id: 'bmp.disk.counts.read', value: '次/秒', text: '次/秒' },
    { id: 'bmp.disk.counts.write', value: '次/秒', text: '次/秒' },
    { id: 'bmp.network.bytes.ingress', value: 'bps', text: 'bps' },
    { id: 'bmp.network.bytes.egress', value: 'bps', text: 'bps' },
    { id: 'bmp.network.packets.ingress', value: '个', text: '个/秒' },
    { id: 'bmp.network.packets.egress', value: '个', text: '个/秒' },
    { id: 'bmp.avg.load1', value: '', text: ' ' },
    { id: 'bmp.avg.load5', value: '', text: ' ' },
    { id: 'bmp.avg.load15', value: '', text: ' ' },
    { id: 'bmp.tcp.connect.total', value: '个', text: '个' },
    { id: 'bmp.tcp.connect.established', value: '个', text: '个' },
    { id: 'bmp.process.total', value: '个', text: '个' },
    { id: 'bmp.disk.partition.used', value: 'GB', text: 'GB' },
    { id: 'bmp.disk.partition.util', value: '%', text: '%' },
];

/**
 * 持续周期类型
*/
export const timesData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.continuePeriod',[1]) ,
        value: 1
    },
    {
        text: global.t('alarm.alarmRule.continuePeriod',[2]),
        value: 2
    },
    {
        text: global.t('alarm.alarmRule.continuePeriod',[3]),
        value: 3
    },
    {
        text: global.t('alarm.alarmRule.continuePeriod',[4]),
        value: 4
    },
    {
        text: global.t('alarm.alarmRule.continuePeriod',[5]),
        value: 5
    },
    {
        text: global.t('alarm.alarmRule.continuePeriod',[10]),
        value: 10
    },
    {
        text: global.t('alarm.alarmRule.continuePeriod',[15]),
        value: 15
    },
    {
        text: global.t('alarm.alarmRule.continuePeriod',[30]),
        value: 30
    },
    {
        text: global.t('alarm.alarmRule.continuePeriod',[60]),
        value: 60
    },
]

/**
 * 告警级别类型
*/
export const noticeLevelData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.normal'),
        value: 1
    },
    {
        text: global.t('alarm.alarmRule.serious'),
        value: 2
    },
    {
        text: global.t('alarm.alarmRule.urgent'),
        value: 3
    },
]

/**
 * 通知周期类型
*/
export const noticePeriodData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.minute',[5]) ,
        value: 5
    },
    {
        text: global.t('alarm.alarmRule.minute',[10]),
        value: 10
    },
    {
        text: global.t('alarm.alarmRule.minute',[15]),
        value: 15
    },
    {
        text: global.t('alarm.alarmRule.minute',[30]),
        value: 30
    },
    {
        text: global.t('alarm.alarmRule.hour',[1]),
        value: 60
    },
    {
        text: global.t('alarm.alarmRule.hour',[3]),
        value: 180
    },
    {
        text: global.t('alarm.alarmRule.hour',[6]),
        value: 360
    },
    {
        text: global.t('alarm.alarmRule.hour',[12]),
        value: 720
    },
    {
        text: global.t('alarm.alarmRule.hour',[24]),
        value: 1440
    },
]

/**
 * 通知条件类型
*/
export const noticeConditionData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.alarm'),
        value: 1
    },
    {
        text: global.t('alarm.alarmRule.beNormal'),
        value: 2
    },
]

/**
 * 接收渠道类型
*/
export const noticeWayData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.inEmail'),
        value: 1
    },
    {
        text: global.t('alarm.alarmRule.email'),
        value: 2
    },
]

/**
 * 报警规则状态
*/
export const alarmRuleStatus: StaticDataType['CurrencyListType'][] = [ // 报警规则状态filter数据
    {
        filterParams: '1',
        text: global.t('alarm.alarmRule.regular'),
        value: 1
    },
    {
        filterParams: '2',
        text: global.t('alarm.alarmRule.forbid'),
        value: 2
    },   
    {
        filterParams: '3',
        text: global.t('alarm.alarmRule.alarm'),
        value: 3
    }, 
];

/**
 * 维度类型
*/
export const dimensionData: StaticDataType['CurrencyListType'][] = [ 
    {
        text: global.t('alarm.alarmRule.instance'),
        value: 'instance'
    },
    {
        text: global.t('alarm.alarmRule.disk'),
        value: 'disk'
    },
    {
        text: global.t('alarm.alarmRule.mountpoint'),
        value: 'mountpoint'
    },
    {
        text: global.t('alarm.alarmRule.nic'),
        value: 'nic'
    },
]



