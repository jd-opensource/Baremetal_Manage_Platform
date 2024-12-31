/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
const {get} = request;

/**
 * 报警历史列表接口
 * @param {Object} params 请求需要的参数
*/
const describeAlertsAPI = (params: Parameters<typeof get>): ReturnType<typeof get> => {
    return get(urlPath.describeAlertsUrl, params, {isBusinessProcessing: false});
};

/**
 * 报警历史列表导出接口
 * @param {Object} params 请求需要的参数
*/
const describeAlertsExportAPI = (params: Parameters<typeof get>): ReturnType<typeof get> => {
    const otherParams = {
        isBusinessProcessing: false,
        timeout: 1000 * 120,
        responseType: 'blob',
        islongReq: true
    };
    return get(urlPath.describeAlertsExportUrl, params, {...otherParams});
};

export {
    describeAlertsAPI,
    describeAlertsExportAPI
};
