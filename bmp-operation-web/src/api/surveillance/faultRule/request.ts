/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
import {ParamDataType} from './typeManagement';
const {get, post} = request;

/**
 * 故障报警规则列表接口
 * @param {Object} params 请求需要的参数
*/
export function faultRulesAPI(params: ParamDataType['faultRulesListType']): ReturnType<typeof get> {
    return get(urlPath.faultRulesUrl, params, {baseURL: '/oob-alert', isBusinessProcessing: false});
};

/**
 * 故障报警规则列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function faultRulesExportAPI(params: ParamDataType['faultRulesLisExportType']): ReturnType<typeof get> {
    return get(urlPath.faultRulesExportUrl, params, {baseURL: '/oob-alert', isBusinessProcessing: false, responseType: 'blob', timeout: 1000 * 120, islongReq: true});
};

/**
 * 故障推送push接口
 * @param {Object} data 请求需要的参数
*/
export function faultPushAPI(data: ParamDataType['faultPushType']): ReturnType<typeof post> {
    return post(urlPath.faultPushUrl, data, true, {baseURL: '/oob-alert'});
};

/**
 * 故障启用、禁用接口
 * @param {Object} data 请求需要的参数
*/
export function faultUseAPI(data: ParamDataType['faultUseType']): ReturnType<typeof post> {
    return post(urlPath.faultUseUrl, data, true, {baseURL: '/oob-alert'});
};

/**
 * 恢复默认设置接口
 * @param {Object} data 请求需要的参数
*/
export function faultResetAPI(data: ParamDataType['faultResetType']): ReturnType<typeof post> {
    return post(urlPath.faultResetUrl, data, true, {baseURL: '/oob-alert'});
};

/**
 * 故障等级列表接口
 * @param {Object} params 请求需要的参数
*/
export function faultLevelAPI(params: ParamDataType['faultLevelType']): ReturnType<typeof get> {
    return get(urlPath.faultLevelUrl, params, {baseURL: '/oob-alert', isBusinessProcessing: false});
};

/**
 * 故障配件列表接口
 * @param {Object} params 请求需要的参数
*/
export function faultSpecAPI(params: ParamDataType['faultSpecType']): ReturnType<typeof get> {
    return get(urlPath.faultSpecUrl, params, {baseURL: '/oob-alert', isBusinessProcessing: false});
};
