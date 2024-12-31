/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
import {ParamDataType} from './typeManagement';
const {get, post} = request;

/**
 * 登录接口
 * @param {Object} data 请求需要的参数
*/
const loginAPI = (data: ParamDataType['LoginType']): ReturnType<typeof post> => {
    return post(urlPath.loginUrl, data, false);
};

/**
 * 用户权限接口
 * @param {Object} params 请求需要的参数
*/
const userPurviewAPI = (params: ParamDataType['UserPurviewType']): ReturnType<typeof get> => {
    return get(urlPath.userPurviewUrl, params, {isBusinessProcessing: false});
};

/**
 * 退出登录接口
 * @param {Object} data 请求需要的参数
*/
const loginOutAPI = (data: ParamDataType['LoginOutType']): ReturnType<typeof post> => {
    return post(urlPath.loginOutUrl, data, false);
};

const loginUserApiCount = {
    loginAPI,
    loginOutAPI,
    userPurviewAPI
};

export default loginUserApiCount;
