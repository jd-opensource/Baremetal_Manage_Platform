/**
 * @file
 * @author
*/

import urlPath from './config';
import store from 'store/index.ts';
import request from 'request/index.ts';
import {ParamDataType} from './typeManagement';
const {get, put} = request;

type ProgressEventType = { lengthComputable: boolean; loaded: number; total: number; }

const reqProgress = (progressEvent: ProgressEventType) => {
    // if (progressEvent.lengthComputable) {
    //     const percentComplete = (progressEvent.loaded / progressEvent.total) * 100;
    //     let val:number = 0;
    //     if (Number(percentComplete.toFixed(2)) >= 100) {
    //         val = 100;
    //     }
    //     else {
    //         val = Number(percentComplete.toFixed(2));
    //     }
    //     store.progressInfo.progressNum = val;
    // }
    if (progressEvent.lengthComputable) {
        const percentComplete = (progressEvent.loaded / progressEvent.total) * 100;
        store.progressInfo.progressNum = Math.min(100, Number(percentComplete.toFixed(2)));
    }
}
/**
 * 机房列表接口
 * @param {Object} params 请求需要的参数
*/
const idcListAPI = (params: ParamDataType['IdcListType']): ReturnType<typeof get> => {
    return get(urlPath.idcListUrl, params, {isBusinessProcessing: false});
};

/**
 * 机房列表导出接口
 * @param {Object} params 请求需要的参数
*/
const idcListExportAPI = (params: ParamDataType['idcListExportType']): ReturnType<typeof get> => {
    const otherParams = {
        isBusinessProcessing: false,
        timeout: 1000 * 120,
        responseType: 'blob',
        islongReq: true,
        onDownloadProgress: (progressEvent: ProgressEventType) => reqProgress(progressEvent)
    };
    return get(urlPath.idcListUrl, params, {...otherParams});
};

/**
 * 机房详情接口
 * @param {Object} params 请求需要的参数
*/
const idcDetailAPI = (params: ParamDataType['IdcDetailType']): ReturnType<typeof get> => {
    return get(urlPath.idcDetailUrl + params.idcid, params, {isBusinessProcessing: false});
};

/**
 * 编辑机房接口
 * @param {Object} data 请求需要的参数
*/
const editIdcAPI = (data: ParamDataType['EditIdcType']): ReturnType<typeof put> => {
    return put(urlPath.editIdcUrl + data.idcId, data, true);
};

const idcApiCount = {
    idcListAPI,
    idcListExportAPI,
    idcDetailAPI,
    editIdcAPI
};

export default idcApiCount;
