/**
 * @file
 * @author
*/
import {methodsTotal} from 'utils/index.ts';
import {TotalUrlPathType} from './typeManagement'; // url地址类

const defaultUrl1: string = '/devices';
const defaultUrl2: string = '/instances';
const defaultUrl3: string = '/v1/oobAlert/';
const defaultUrl4: string = '/auditLogs';

// url地址集合
const urlPath: TotalUrlPathType =  {
    devicesListUrl: `${defaultUrl1}/queryDevices`, // 设备列表
    deviceTemplateUrl: '/data/device_import.xlsx', // 设备模板下载
    devicesUpUrl: `${defaultUrl1}/putaway`, // 设备上架
    devicesDownUrl: `${defaultUrl1}/unPutaway`, // 设备下架
    devicesRemoveUrl: `${defaultUrl1}/`, // 设备移除
    instanceRecoveryUrl: `${defaultUrl2}/`, // 回收实例
    resetInstanceUrl: `${defaultUrl2}/resetStatus/`, // 重置实例
    devicesEditUrl: `${defaultUrl1}/`, // 设备编辑
    devicesDeleteUrl: `${defaultUrl1}/`, // 设备删除
    devicesStartUrl: `${defaultUrl2}/start/`, // 设备开机
    devicesResetPasswordUrl: `${defaultUrl2}/resetPasswd`, // 重置密码
    devicesBatchStartUrl: `${defaultUrl2}/startInstances`, // 设备批量开机
    devicesStopUrl: `${defaultUrl2}/stop/`, // 设备关机
    devicesBatchStopUrl: `${defaultUrl2}/stopInstances`, // 设备批量关机
    devicesRestartUrl: `${defaultUrl2}/restart/`, // 设备重启
    devicesBatchRestartUrl: `${defaultUrl2}/restartInstances`, // 设备批量重启
    deviceBatchResetPasswordUrl: `${defaultUrl2}/batchResetPasswd`, // 批量重置密码
    lockUrl: `${defaultUrl2}/lock/`, // 锁定
    unlockUrl: `${defaultUrl2}/unlock/`, // 解锁
    resetSystemUrl: `${defaultUrl2}/reinstallInstance`, // 重装系统
    devicesDetailUrl: `${defaultUrl1}/queryDevice/`, // 设备详情
    devicesCollectUrl: `${defaultUrl1}/collect`, // 设备采集
    instanceDetailUrl: `${defaultUrl2}/`, // 实例详情
    importDevicesUrl: defaultUrl1, // 导入设备
    batchEditInstanceNameUrl: `${defaultUrl2}/modifyInstances`, // 批量编辑实例名称
    batchRecoveryInstanceUrl: `${defaultUrl2}/deleteInstances`, // 批量回收实例
    deviceStatusUrl: methodsTotal.humpSplit(`${defaultUrl3}device/status`), // 获取设备带外总体状态
    consolelogsUrl: methodsTotal.humpSplit(`${defaultUrl3}consolelogs`), // 设备详情-硬件监控-报警日志列表接口
    disksDetailUrl: `${defaultUrl1}/`, // 磁盘详情
    describeAssociateDisksUrl: `${defaultUrl1}/disks/describeAssociateDisks`, // 获取设备关联磁盘信息
    associateDisksUrl: `${defaultUrl1}/disks/associateDisks`, // 关联设备磁盘
    describeVolumesUrl: `/deviceTypes/`, // 根据机型获取卷列表信息
    associateUrl: `${defaultUrl1}/deviceType/associate`, // 设备关联机型
    auditLogsUrl: defaultUrl4, // 设备详情-操作日志
    auditLogsTypesUrl: `${defaultUrl4}/types`, // 获取操作日志类型的接口
    monitorDataUrl: '/monitorData', // 监控图表
    desrcibeTagsUrl: '/monitorProxy/desrcibeTags', // 监控tag
    desrcibeAgentStatusUrl: '/monitorProxy/desrcibeAgentStatus', // 监控Agent章台
};

export default urlPath;
