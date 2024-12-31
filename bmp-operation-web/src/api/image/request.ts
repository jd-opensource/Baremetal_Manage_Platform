/**
 * @file
 * @author
*/
import urlPath from './config';
import request from 'request/index.ts';
import {ParamDataType} from './typeManagement';
const {get, post, put, deleteReq} = request;

/**
 * 镜像列表接口
 * @param {Object} params 请求需要的参数
*/
export function imagesAPI(params: ParamDataType['ImagesType']) {
    return get(urlPath.imagesUrl, params, {isBusinessProcessing: false});
};

/**
 * 镜像列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function imagesExportAPI(params: ParamDataType['ImagesType']) {
    return get(urlPath.imagesUrl, params, {isBusinessProcessing: false, timeout: 1000 * 120, responseType: 'blob', islongReq: true});
};

/**
 * 镜像详情接口
 * @param {Object} params 请求需要的参数
*/
export function imagesDetailAPI(params: ParamDataType['ImagesType']) {
    return get(urlPath.imagesDetailUrl + params.imageId, params, {isBusinessProcessing: false});
};

/**
 * 镜像详情-机型列表-添加机型列表接口
 * @param {Object} params 请求需要的参数
*/
export function imageDeviceTypesAPI(params: ParamDataType['ImagesModelType']) {
    return get(urlPath.imageDeviceTypesUrl, params, {isBusinessProcessing: false});
};

/**
 * 镜像详情-描述编辑接口
 * @param {Object} data 请求需要的参数
*/
export function imageEditAPI(data: ParamDataType['ImageEditType']) {
    return put(urlPath.imageEditUrl + data.imageId, data, true);
};

/**
 * 镜像详情-添加机型接口
 * @param {Object} data 请求需要的参数
*/
export function imageAddModelAPI(data: ParamDataType['ImageAddModelType']) {
    return post(urlPath.imageAddModelUrl, data, true);
};

/**
 * 镜像详情-移除关联机型接口
 * @param {Object} data 请求需要的参数
*/
export function deleteImageModelAPI(data: ParamDataType['DeleteImageModelType']) {
    return deleteReq(urlPath.deleteImageModelUrl, data, true);
};

/**
 * 镜像删除接口
 * @param {Object} data 请求需要的参数
*/
export function imagesDeleteAPI(data: ParamDataType['ImagesDeleteType']) {
    return deleteReq(urlPath.imagesDeleteUrl + data.imageId, data, true);
};

/**
 * 导入镜像接口
 * @param {Object} data 请求需要的参数
*/
export function importImageAPI(data: ParamDataType['ImportImageType']) {
    return post(urlPath.importImageUrl, data, true, {timeout: 1000 * 120, islongReq: true});
};

/**
 * 镜像-oss接口
 * @param {Object} params 请求需要的参数
*/
export function imageOssAPI(params: ParamDataType['ImageOssType']) {
    return get(urlPath.imageOssUrl, params, {isBusinessProcessing: false});
};
