/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
const {get} = request;


/**
 * 带内监控-监控图表
 * @param {Object} params 请求需要的参数
*/
export function monitorDataAPI(params: Parameters<typeof get>): ReturnType<typeof get> {
    return get(urlPath.monitorDataUrl, params, {isBusinessProcessing: false});
};
