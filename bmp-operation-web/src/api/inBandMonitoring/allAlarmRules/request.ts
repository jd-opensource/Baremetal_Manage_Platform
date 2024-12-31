/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
const {get} = request;

/**
 * 全部报警规则列表接口
 * @param {Object} params 请求需要的参数
*/
const describeRulesAPI = (params: Parameters<typeof get>): ReturnType<typeof get> => {
    return get(urlPath.describeRulesUrl, params, {isBusinessProcessing: false});
};

/**
 * 全部报警规则列表导出接口
 * @param {Object} params 请求需要的参数
*/
const describeRulesExportAPI = (params: Parameters<typeof get>): ReturnType<typeof get> => {
    const otherParams = {
        isBusinessProcessing: false,
        timeout: 1000 * 120,
        responseType: 'blob',
        islongReq: true
    };
    return get(urlPath.describeRulesExportUrl, params, {...otherParams});
};

/**
 * 全部报警规则详情接口
 * @param {Object} params 请求需要的参数
*/
const describeRulesDetailAPI = (params: Parameters<typeof get>): ReturnType<typeof get> => {
    return get(urlPath.describeRulesDetailUrl, params, {isBusinessProcessing: false});
};

export {
    describeRulesAPI,
    describeRulesExportAPI,
    describeRulesDetailAPI
};
