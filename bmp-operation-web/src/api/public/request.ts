/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
import {ParamDataType} from './typeManagement';
const {get, post} = request;

/**
 * 精确查询
 * @param {Object} params 请求需要的参数
*/
export function resourcesAPI(params: ParamDataType['ResourcesType']): ReturnType<typeof get> {
    return get(urlPath.resourcesUrl, params, {isBusinessProcessing: false});
};

/**
 * 安全验证接口
 * @param {Object} data 请求需要的参数
*/
export function securityVerificationAPI(data: ParamDataType['SecurityVerificationType']): ReturnType<typeof post> {
    return post(urlPath.securityVerificationUrl, data, false);
};

/**
 * 自定义列表接口
 * @param {Object} params 请求需要的参数
*/
export function customListAPI(params: ParamDataType['CustomListType']): ReturnType<typeof get> {
    return get(urlPath.customListUrl, params, {isBusinessProcessing: false});
};

/**
 * 设置自定义列表接口
 * @param {Object} data 请求需要的参数
*/
export function setCustomListAPI(data: ParamDataType['SetCustomListType']): ReturnType<typeof post> {
    return post(urlPath.setCustomListUrl, data, true);
};

/**
 * 监控自定义列表接口
 * @param {Object} params 请求需要的参数
*/
export function surveillanceCustomListAPI(params: ParamDataType['CustomListType']): ReturnType<typeof get> {
    return get(urlPath.surveillanceCustomListUrl, params, {baseURL: '/oob-alert', isBusinessProcessing: false});
};

/**
 * 监控设置自定义列表接口
 * @param {Object} data 请求需要的参数
*/
export function surveillanceSetCustomListAPI(data: ParamDataType['SetCustomListType']): ReturnType<typeof post> {
    return post(urlPath.surveillanceSetCustomListUrl, data, true, {baseURL: '/oob-alert'});
};


/**
 * oss接口
 * @param {Object} params 请求需要的参数
*/
export function getOssAPI(params: ParamDataType['OssTyoe']): ReturnType<typeof get> {
    return get(urlPath.getOssUrl, params, {isBusinessProcessing: false});
};

/**
 * 镜像类型接口
 * @param {Object} params 请求需要的参数
*/
export function queryImagesAPI(params: ParamDataType['QueryImagesType']): ReturnType<typeof get> {
    return get(urlPath.queryImagesUrl, params, {isBusinessProcessing: true});
};

/**
 * raids接口
 * @param {Object} params 请求需要的参数
*/
export function systemRaidsAPI(params: ParamDataType['SystemRaidsType']): ReturnType<typeof get> {
    return get(urlPath.systemRaidsUrl, params, {isBusinessProcessing: true});
};

/**
 * 系统盘分区接口
 * @param {Object} params 请求需要的参数
*/
export function querySystemPartitionAPI(params: ParamDataType['SystemPartitionType']): ReturnType<typeof get> {
    return get(urlPath.querySystemPartitionUrl, params, {isBusinessProcessing: true});
};

/**
 * 密钥信息接口
 * @param {Object} params 请求需要的参数
*/
export function keyInfoAPI(params: ParamDataType['KeyInfoType']): ReturnType<typeof get> {
    return get(urlPath.keyInfoUrl, params, {isBusinessProcessing: true});
};

/**
 * 权限列表
 * @param {Object} params 请求需要的参数
*/
export function localAPI(params: Parameters<typeof get>): ReturnType<typeof get> {
    return get(urlPath.localUrl, params, {isBusinessProcessing: false});
};
