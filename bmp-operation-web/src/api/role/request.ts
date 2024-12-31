/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
import {ParamDataType} from './typeManagement';
const {get} = request;

/**
 * 角色列表接口
 * @param {Object} params 请求需要的参数
*/
const roleListAPI = (params: ParamDataType['roleListType']): ReturnType<typeof get> => {
    return get(urlPath.roleListUrl, params, {isBusinessProcessing: true});
};

const roleApiCount = {
    roleListAPI
};

export default roleApiCount;
