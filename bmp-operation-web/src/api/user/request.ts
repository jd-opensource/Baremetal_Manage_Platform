/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
import {ParamDataType} from './typeManagement';
const {get, post, put, deleteReq} = request;

/**
 * 用户列表接口
 * @param {Object} params 请求需要的参数
*/
const userListAPI = (params: ParamDataType['UserListType']): ReturnType<typeof get> => {
    return get(urlPath.userListUrl, params, {isBusinessProcessing: false});
};

/**
 * 用户列表导出接口
 * @param {Object} params 请求需要的参数
*/
const userListExportAPI = (params: ParamDataType['UserListExportType']): ReturnType<typeof get> => {
    return get(urlPath.userListUrl, params, {isBusinessProcessing: false, timeout: 1000 * 120, responseType: 'blob', islongReq: true});
};

/**
 * 添加用户接口
 * @param {Object} data 请求需要的参数
*/
const addUserAPI = (data: ParamDataType['AddUserType']): ReturnType<typeof put> => {
    return put(urlPath.addUserUrl, data, true);
};

/**
 * 编辑用户接口
 * @param {Object} data 请求需要的参数
*/
const editUserAPI = (data: ParamDataType['EditUserType']): ReturnType<typeof post> => {
    return post(urlPath.editUserUrl + data.userId, data, true);
};

/**
 * 删除用户接口
 * @param {Object} data 请求需要的参数
*/
const deleteUserAPI = (data: ParamDataType['DeleteUserType']): ReturnType<typeof deleteReq> => {
    return deleteReq(urlPath.deleteUserUrl + data.userId, data, true);
};

const userApiCount = {
    addUserAPI,
    userListAPI,
    editUserAPI,
    deleteUserAPI,
    userListExportAPI
};

export default userApiCount;
