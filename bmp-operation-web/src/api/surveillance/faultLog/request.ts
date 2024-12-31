/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
import {ParamDataType} from './typeManagement';
const {get, post} = request;

/**
 * 规则报警日志列表接口
 * @param {Object} params 请求需要的参数
*/
const faultLogAPI = (params: ParamDataType['faultLogListType']): ReturnType<typeof get> => {
    return get(urlPath.faultLogUrl, params, {baseURL: '/oob-alert', isBusinessProcessing: false, timeout: 1000 * 20, islongReq: true});
};

/**
 * 故障报警日志列表导出接口
 * @param {Object} params 请求需要的参数
*/
const faultLogExportAPI = (params: ParamDataType['faultLogExportType']): ReturnType<typeof get> => {
    return get(urlPath.faultLogExportUrl, params, {baseURL: '/oob-alert', isBusinessProcessing: false, responseType: 'blob', timeout: 1000 * 120, islongReq: true});
};


/**
 * 故障报警日志处理接口
 * @param {Object} data 请求需要的参数
*/
const faultLogDealAPI = (data: ParamDataType['faultLogDealType']): ReturnType<typeof post> => {
    return post(urlPath.faultLogDealUrl, data, false, {baseURL: '/oob-alert'});
};


const faultLogApiCount = {
    faultLogAPI,
    faultLogExportAPI,
    faultLogDealAPI
};

export default faultLogApiCount;
