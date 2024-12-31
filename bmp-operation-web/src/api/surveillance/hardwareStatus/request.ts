/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
import {ParamDataType} from './typeManagement';
const {get} = request;

/**
 * 硬件报警状态列表接口
 * @param {Object} params 请求需要的参数
*/
const hardwareStatusAPI = (params: ParamDataType['hardwareStatusType']): ReturnType<typeof get> => {
    return get(urlPath.hardwareStatusUrl, params, {baseURL: '/oob-alert', isBusinessProcessing: false});
};

/**
 * 硬件报警状态导出列表导出接口
 * @param {Object} params 请求需要的参数
*/
const hardwareStatusExportAPI = (params: ParamDataType['hardwareStatusExportType']) => {
    return get(urlPath.hardwareStatusExportUrl, params, {baseURL: '/oob-alert', timeout: 1000 * 120, isBusinessProcessing: false, responseType: 'blob', islongReq: true});
};


const hardwareStatusApiCount = {
    hardwareStatusAPI,
    hardwareStatusExportAPI
};

export default hardwareStatusApiCount;
