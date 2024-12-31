/**
 * @file
 * @author
*/
import urlPath from './config';
import request from 'request/index.ts';
import {ParamDataType} from './typeManagement';
const {get, post, put, deleteReq} = request;

/**
 * 机型列表接口
 * @param {Object} params 请求需要的参数
*/
export function modelListAPI(params: ParamDataType['ModeListType']): ReturnType<typeof get> {
    return get(urlPath.modelListUrl, params, {isBusinessProcessing: false});
};

/**
 * 机型列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function modelListExportAPI(params: ParamDataType['ModelListExportType']): ReturnType<typeof get> {
    return get(urlPath.modelListUrl, params, {isBusinessProcessing: false, timeout: 1000 * 120, responseType: 'blob', islongReq: true});
};

/**
 * 机型详情接口
 * @param {Object} params 请求需要的参数
*/
export function modelDetailAPI(params: ParamDataType['ModelDetailType']): ReturnType<typeof get> {
    return get(urlPath.modelDetailUrl + params.id, params, {isBusinessProcessing: false});
};

/**
 * 关联镜像列表接口
 * @param {Object} params 请求需要的参数
*/
export function relationImageListAPI(params: ParamDataType['RelationImageListType']): ReturnType<typeof get> {
    return get(urlPath.relationImageListUrl, params, {isBusinessProcessing: false});
};

/**
 * 添加镜像接口
 * @param {Object} data 请求需要的参数
*/
export function addImageAPI(data: ParamDataType['AddImageType']): ReturnType<typeof post> {
    return post(urlPath.addImageUrl, data, true);
};

/**
 * 添加机型接口
 * @param {Object} data 请求需要的参数
*/
export function addModelAPI(data: ParamDataType['ModelCurrencyType']): ReturnType<typeof put> {
    return put(urlPath.addModelUrl, data, true);
};

/**
 * 编辑机型接口
 * @param {Object} data 请求需要的参数
*/
export function editModelAPI(data: ParamDataType['ModelCurrencyType']): ReturnType<typeof put> {
    return put(urlPath.editModelUrl + data.deviceTypeId, data, true);
};

/**
 * 删除关联镜像接口
 * @param {Object} data 请求需要的参数
*/
export function deleteImageAPI(data: ParamDataType['DeleteImageType']): ReturnType<typeof put> {
    return post(urlPath.deleteImageUrl, data, true);
};

/**
 * 删除机型接口
 * @param {Object} data 请求需要的参数
*/
export function deleteModelAPI(data: ParamDataType['DeleteModelType']): ReturnType<typeof deleteReq> {
    return deleteReq(urlPath.deleteModelUrl + data.id, data, true);
};

/**
 * raid接口
 * @param {Object} params 请求需要的参数
*/
export function raidsAPI(params: ParamDataType['RaidsType']): ReturnType<typeof get> {
    return get(urlPath.getRaidsUrl, params, {isBusinessProcessing: true});
};

/**
 * filter接口
 * @param {Object} params 请求需要的参数
*/
export function getFilterListAPI(params: ParamDataType['FilterListType']): ReturnType<typeof get> {
    return get(urlPath.filterListUrl, params, {isBusinessProcessing: false});
};

const modelApiCount = {
    modelListAPI,
    modelListExportAPI,
    modelDetailAPI,
    relationImageListAPI,
    addImageAPI,
    addModelAPI,
    editModelAPI,
    deleteImageAPI,
    deleteModelAPI,
    raidsAPI,
    getFilterListAPI
};

export default modelApiCount;
