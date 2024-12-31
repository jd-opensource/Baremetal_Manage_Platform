/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
import {ParamDataType} from './typeManagement';
const {get, put, deleteReq} = request;

/**
 * 时区
 * @param {Object} params 请求需要的参数
*/
export function timeOneAPI(params: ParamDataType['TimeOneType']): ReturnType<typeof get> {
    return get(urlPath.timeOneUrl, params, {isBusinessProcessing: true});
};

/**
 * 用户信息
 * @param {Object} params 请求需要的参数
*/
export function userInfoAPI(params: ParamDataType['UserInfoType']): ReturnType<typeof get> {
    return get(urlPath.userInfoUrl, params, {isBusinessProcessing: false});
};

/**
 * 修改用户信息
 * @param {Object} data 请求需要的参数
*/
export function setUserInfoAPI(data: ParamDataType['UserInfoType']): ReturnType<typeof put> {
    return put(urlPath.setUserInfoUrl, data, false);
};

/**
 * 获取apikey
 * @param {Object} params 请求需要的参数
*/
export function getApiKeyAPI(params: ParamDataType['GetApiKeyType']): ReturnType<typeof get> {
    return get(urlPath.getApiKeyUrl, params, {isBusinessProcessing: true});
};

/**
 * 创建apikey
 * @param {Object} data 请求需要的参数
*/
export function createApiKeyAPI(data: ParamDataType['CreateApiKeyType']): ReturnType<typeof put> {
    return put(urlPath.createApiKeyUrl, data, true);
};

/**
 * 删除apikey
 * @param {Object} data 请求需要的参数
*/
export function deleteApiKeyAPI(data: ParamDataType['DeleteApiKeyType']): ReturnType<typeof deleteReq> {
    return deleteReq(urlPath.deleteApiKeyUrl + data.apiKey, data, true);
};

/**
 * 修改密码
 * @param {Object} data 请求需要的参数
*/
export function revisePasswordAPI(data: ParamDataType['RevisePasswordType']): ReturnType<typeof put> {
    return put(urlPath.revisePasswordUrl, data, false);
};
