/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
import {AxiosPromise} from 'axios';
import {ParamDataType, RequestType} from './typeManagement';
import store from 'store/index.ts';
const {get, put, deleteReq, post} = request;

type ProgressEventType = { lengthComputable: boolean; loaded: number; total: number; }

const reqProgress = (progressEvent: ProgressEventType) => {
    if (progressEvent.lengthComputable) {
        const percentComplete = (progressEvent.loaded / progressEvent.total) * 100;
        store.progressInfo.progressNum = Number(percentComplete.toFixed(2));
    }
}

/**
 * 设备列表接口
 * @param {Object} params 请求需要的参数
*/
export function devicesListAPI(params: ParamDataType['DevicesListType']): ReturnType<typeof get> {
    return get(urlPath.devicesListUrl, params, {isBusinessProcessing: false, islongReq: true});
};

/**
 * 设备模板接口
 * @param {Object} params 请求需要的参数
*/
export function deviceTemplateAPI(params: ParamDataType['DeviceTemplateType']): ReturnType<typeof get> {
    return get(urlPath.deviceTemplateUrl, params, {baseURL: '/template-upload', timeout: 1000 * 120, isBusinessProcessing: false, responseType: 'blob', islongReq: true});
};

/**
 * 设备列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function devicesListExportAPI(params: ParamDataType['DevicesListExportType']): ReturnType<typeof get> {
    return get(urlPath.devicesListUrl, params, {timeout: 1000 * 120, isBusinessProcessing: false, responseType: 'blob', islongReq: true});
};

/**
 * 设备详情-硬件监控-报警日志列表接口
 * @param {Object} params 请求需要的参数
*/
export function consolelogsAPI(params: ParamDataType['ConsolelogsType']): ReturnType<typeof get> {
    return get(urlPath.consolelogsUrl, params, {baseURL: '/oob-alert', isBusinessProcessing: false});
};

/**
 * 设备详情-硬件监控-报警日志列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function consolelogsExportAPI(params: ParamDataType['ConsolelogsType']): ReturnType<typeof get> {
    return get(urlPath.consolelogsUrl, params, {baseURL: '/oob-alert', timeout: 1000 * 120, isBusinessProcessing: false, responseType: 'blob', islongReq: true});
};

/**
 * 设备上架接口
 * @param {Object} data 请求需要的参数
*/
export function devicesUpAPI(data: ParamDataType['DeviceUpType']): ReturnType<typeof put> {
    return put(urlPath.devicesUpUrl, data, true);
};

/**
 * 设备采集接口
 * @param {Object} data 请求需要的参数
*/
export function devicesCollectAPI(data: {sn: string; allowOverride: boolean}): ReturnType<typeof post> {
    return post(urlPath.devicesCollectUrl, data, true);
};

/**
 * 设备下架接口
 * @param {Object} data 请求需要的参数
*/
export function devicesDownAPI(data: ParamDataType['DeviceDownType']): ReturnType<typeof put> {
    return put(urlPath.devicesDownUrl, data, true);
};

/**
 * 设备移除接口
 * @param {Object} data 请求需要的参数
*/
export function devicesRemoveAPI(data: ParamDataType['DevicesRemoveType']): ReturnType<typeof put> {
    return put(urlPath.devicesRemoveUrl + data.deviceId + '/remove', data, true);
};

/**
 * 回收实例接口
 * @param {Object} data 请求需要的参数
*/
export function instanceRecoveryAPI(data: ParamDataType['InstanceRecoveryType']): ReturnType<typeof deleteReq> {
    return deleteReq(urlPath.instanceRecoveryUrl + data.instanceId, data, true);
};

/**
 * 重置实例接口
 * @param {Object} data 请求需要的参数
*/
export function resetInstanceAPI(data: ParamDataType['ResetInstanceType']): ReturnType<typeof post> {
    return post(urlPath.resetInstanceUrl + data.instanceId, data, true);
};

/**
 * 设备描述编辑接口
 * @param {Object} data 请求需要的参数
*/
export function devicesEditAPI(data: ParamDataType['DevicesEditType']): ReturnType<typeof put> {
    return put(urlPath.devicesEditUrl + data.id, data, true);
};

/**
 * 设备删除接口
 * @param {Object} data 请求需要的参数
*/
export function devicesDeleteAPI(data: ParamDataType['DeviceDeleteType']): ReturnType<typeof deleteReq> {
    return deleteReq(urlPath.devicesDeleteUrl + data.deviceId, data, true);
};

/**
 * 设备开机接口
 * @param {Object} data 请求需要的参数
*/
export function devicesStartAPI(data: ParamDataType['DevicesStartType']): ReturnType<typeof post> {
    return post(urlPath.devicesStartUrl + data.instanceId, data, true);
};

/**
 * 设备重置密码接口
 * @param {Object} data 请求需要的参数
*/
export function devicesResetPasswordAPI(data: ParamDataType['DevicesResetPasswordType']): ReturnType<typeof post> {
    return post(urlPath.devicesResetPasswordUrl, data, true);
};


/**
 * 设备批量开机接口
 * @param {Object} data 请求需要的参数
*/
export function devicesBatchStartAPI(data: ParamDataType['DevicesStartType']): ReturnType<typeof post> {
    return post(urlPath.devicesBatchStartUrl, data, true);
};

/**
 * 设备关机接口
 * @param {Object} data 请求需要的参数
*/
export function devicesStopAPI(data: ParamDataType['DevicesStopType']): ReturnType<typeof post> {
    return post(urlPath.devicesStopUrl + data.instanceId, data, true);
};

/**
 * 设备批量关机接口
 * @param {Object} data 请求需要的参数
*/
export function devicesBatchStopAPI(data: ParamDataType['DevicesStopType']): ReturnType<typeof post> {
    return post(urlPath.devicesBatchStopUrl, data, true);
};

/**
 * 设备重启接口
 * @param {Object} data 请求需要的参数
*/
export function devicesRestartAPI(data: ParamDataType['DevicesRestartType']): ReturnType<typeof post> {
    return post(urlPath.devicesRestartUrl + data.instanceId, {}, true);
};

/**
 * 设备批量重启接口
 * @param {Object} data 请求需要的参数
*/
export function devicesBatchRestartAPI(data: ParamDataType['DevicesRestartType']): ReturnType<typeof post> {
    return post(urlPath.devicesBatchRestartUrl, data, true);
};

/**
 * 设备批量重置密码接口
 * @param {Object} data 请求需要的参数
*/
export function deviceBatchResetPasswordAPI(data: ParamDataType['BatchResetPasswordType']): ReturnType<typeof post> {
    return post(urlPath.deviceBatchResetPasswordUrl, data, true);
};

/**
 * 锁定接口
 * @param {Object} data 请求需要的参数
*/
export function lockAPI(data: ParamDataType['LockType']): ReturnType<typeof post> {
    return post(urlPath.lockUrl + data.instanceId, data, true);
};

/**
 * 解锁接口
 * @param {Object} data 请求需要的参数
*/
export function unlockAPI(data: ParamDataType['UnlockType']): ReturnType<typeof post> {
    return post(urlPath.unlockUrl + data.instanceId, data, true);
};

/**
 * 重装系统接口
 * @param {Object} data 请求需要的参数
*/
export function resetSystemAPI(data: ParamDataType['ResetSystemType']): ReturnType<typeof post> {
    return post(urlPath.resetSystemUrl, data, true);
};

/**
 * 设备详情接口
 * @param {Object} params 请求需要的参数
*/
export function devicesDetailAPI(params: ParamDataType['DevicesDetailType']): ReturnType<typeof get> {
    return get(urlPath.devicesDetailUrl + params.deviceId, params, {isBusinessProcessing: false});
};

/**
 * 实例详情接口
 * @param {Object} params 请求需要的参数
*/
export function instanceDetailAPI(params: ParamDataType['InstanceDetailType']): ReturnType<typeof get> {
    return get(urlPath.instanceDetailUrl + params.instanceId, params, {isBusinessProcessing: false});
};

/**
 * 导入设备接口
 * @param {Object} data 请求需要的参数
*/
export function importDevicesAPI(data: ParamDataType['ImportDevicesType']): ReturnType<typeof post> {
    return post(urlPath.importDevicesUrl, data, true, {timeout: 1000 * 180, islongReq: true});
};

/**
 * 批量编辑实例名称接口
 * @param {Object} data 请求需要的参数
*/
export function BatchEditInstanceNameAPI(data: ParamDataType['BatchEditInstanceNameType']): ReturnType<typeof post> {
    return post(urlPath.batchEditInstanceNameUrl, data, false);
};

/**
 * 批量回收实例接口
 * @param {Object} data 请求需要的参数
*/
export function BatchRecoveryInstanceAPI(data: ParamDataType['BatchRecoveryInstanceType']): ReturnType<typeof deleteReq> {
    return deleteReq(urlPath.batchRecoveryInstanceUrl, data, true);
};

/**
 * 获取设备带外总体状态接口
 * @param {Object} params 请求需要的参数
*/
export function deviceStatusAPI(params: ParamDataType['deviceStatusType']): AxiosPromise<RequestType['CurrencyType']> {
    return get(urlPath.deviceStatusUrl, params, {isBusinessProcessing: false, baseURL: '/oob-alert'});
};

/**
 * 设备详情-磁盘信息接口
 * @param {Object} params 请求需要的参数
*/
export function disksDetailAPI(params: {deviceId: string}): ReturnType<typeof get> {
    return get(urlPath.disksDetailUrl + params.deviceId + '/disksDetail', {}, {isBusinessProcessing: false});
};

/**
 * 获取设备关联的磁盘信息
 * @param {Object} params 请求需要的参数
*/
export function describeAssociateDisksAPI(params: Parameters<typeof get>): ReturnType<typeof get> {
    return get(urlPath.describeAssociateDisksUrl, params, {isBusinessProcessing: true});
};

/**
 * 关联设备磁盘
 * @param {Object} data 请求需要的参数
*/
export function associateDisksAPI(data: Parameters<typeof put>): ReturnType<typeof put> {
    return put(urlPath.associateDisksUrl, data, true);
};

/**
 * 根据机型获取卷列表信息
 * @param {Object} params 请求需要的参数
*/
export function describeVolumesAPI(params: {deviceTypeId: string}): ReturnType<typeof get> {
    return get(urlPath.describeVolumesUrl + params.deviceTypeId + '/describeVolumes', {}, {isBusinessProcessing: true});
};

/**
 * 设备关联机型
 * @param {Object} data 请求需要的参数
*/
export function associateAPI(data: {deviceTypeId: string}): ReturnType<typeof put> {
    return put(urlPath.associateUrl, data, true);
};

/**
 * 设备详情-操作日志
 * @param {Object} params 请求需要的参数
*/
export function auditLogsAPI(params: Parameters<typeof get>): ReturnType<typeof get> {
    return get(urlPath.auditLogsUrl, params, {isBusinessProcessing: false});
};

/**
 * 设备详情-获取操作日志类型的接口
 * @param {Object} params 请求需要的参数
*/
export function auditLogsTypesAPI(params: Parameters<typeof get>): ReturnType<typeof get> {
    return get(urlPath.auditLogsTypesUrl, params, {isBusinessProcessing: false});
};

/**
 * 设备详情-操作日志-导出
 * @param {Object} params 请求需要的参数
*/
export function auditLogsExportAPI(params: Parameters<typeof get>): ReturnType<typeof get> {
    const otherParams = {
        isBusinessProcessing: false,
        timeout: 1000 * 120,
        responseType: 'blob',
        islongReq: true,
        onDownloadProgress: (progressEvent: ProgressEventType) => reqProgress(progressEvent)
    };
    return get(urlPath.auditLogsUrl, params, {...otherParams});
};

/**
 * 设备详情-监控图表
 * @param {Object} params 请求需要的参数
*/
export function monitorDataAPI(params: Parameters<typeof get>): ReturnType<typeof get> {
    return get(urlPath.monitorDataUrl, params, {isBusinessProcessing: false});
};

/**
 * 设备详情-监控tags
 * @param {Object} params 请求需要的参数
*/
export function desrcibeTagsAPI(params: Parameters<typeof get>): ReturnType<typeof get> {
    return get(urlPath.desrcibeTagsUrl, params, {isBusinessProcessing: false});
};

/**
 * 设备详情-监控agent状态
 * @param {Object} params 请求需要的参数
*/
export function desrcibeAgentStatusAPI(params: Parameters<typeof get>): ReturnType<typeof get> {
    return get(urlPath.desrcibeAgentStatusUrl, params, {isBusinessProcessing: false});
};
