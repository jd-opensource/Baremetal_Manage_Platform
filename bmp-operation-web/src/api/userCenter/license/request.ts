/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
const {get} = request;

/**
 * 授权证内容接口
 * @param {Object} params 请求需要的参数
*/
const licenseAPI = (params: {}): ReturnType<typeof get> => {
    return get(urlPath.licenseUrl, params, {isBusinessProcessing: false});
};

/**
 * 请求码下载接口
 * @param {Object} params 请求需要的参数
*/
const downloadTokenAPI = (params: {}): ReturnType<typeof get> => {
    return get(urlPath.downloadTokenUrl, params, {timeout: 1000 * 120, isBusinessProcessing: false, responseType: 'blob', islongReq: true});
};

const licenseApiCount = {
    licenseAPI,
    downloadTokenAPI
};

export default licenseApiCount;
