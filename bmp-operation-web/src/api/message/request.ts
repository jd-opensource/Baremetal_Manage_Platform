/**
 * @file
 * @author
*/

import urlPath from './config';
import request from 'request/index.ts';
import store from 'store/index.ts';

const {get, deleteReq, post, put} = request;
const reqProgress = (progressEvent: any) => {
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
 * 消息列表接口
 * @param {Object} params 请求需要的参数
*/
const messageListAPI = (params: Parameters<typeof get>): ReturnType<typeof get> => {
    return get(urlPath.messageListUrl, params, {isBusinessProcessing: false, islongReq: true});
};

/**
 * 消息列表导出接口
 * @param {Object} params 请求需要的参数
*/
const messageExportAPI = (params: Parameters<typeof get>): ReturnType<typeof get> => {
    return get(urlPath.messageListUrl, params, {isBusinessProcessing: false, timeout: 1000 * 120, responseType: 'blob', islongReq: true,  onDownloadProgress: (progressEvent: any) => reqProgress(progressEvent)});
};

/**
 * 消息列表总数和未读数接口
 * @param {Object} params 请求需要的参数
*/
const messagesStatisticAPI = (params: {}): ReturnType<typeof get> => {
    return get(urlPath.statisticUrl, params, {isBusinessProcessing: false});
};

/**
 * 获取消息类型、子类型接口
 * @param {Object} params 请求需要的参数
*/
const messageTypesAPI = (params: {}): ReturnType<typeof get> => {
    return get(urlPath.messageTypesUrl, params, {isBusinessProcessing: false});
};

/**
 * 获取有没有未读消息
 * @param {Object} params 请求需要的参数
*/
const hasUnreadMessageAPI = (params: {}): ReturnType<typeof get> => {
    return get(urlPath.hasUnreadMessageUrl, params, {isBusinessProcessing: false});
};

/**
 * 获取消息详情id
 * @param {Object} params 请求需要的参数
*/
const messageByIdAPI = (params: {messageId: string}): ReturnType<typeof get> => {
    return get(urlPath.messageByIdUrl, params, {isBusinessProcessing: false});
};

/**
 * 消息标为已读
 * @param {Object} data 请求需要的参数
*/
const doReadAPI = (data: {messageIds: [string]}): ReturnType<typeof put> => {
    return put(urlPath.doReadUrl, data, false);
};

/**
 * 删除消息接口
 * @param {Object} data 请求需要的参数
*/
const deleteAPI = (data: Parameters<typeof deleteReq>): ReturnType<typeof deleteReq> => {
    return deleteReq(urlPath.deleteUrl, data, true);
};

/**
 * 读消息设置
 * @param {Object} params 请求需要的参数
*/
const describeMailAPI = (params: {}): ReturnType<typeof get> => {
    return get(urlPath.describeMailUrl, params, {isBusinessProcessing: false});
};

/**
 * 消息设置-校验邮箱信息
 * @param {Object} data 请求需要的参数
*/
const dialMailAPI = (data: {server_addr: string; server_port: string; admin_addr: string; admin_pass: string; to: string;}) => {
    return post(urlPath.dialMailUrl, data, true, {timeout: 1000 * 30});
};

/**
 * 消息设置-保存消息设置
 * @param {Object} data 请求需要的参数
*/
const saveMailAPI = (data: Parameters<typeof post>): ReturnType<typeof post> => {
    return post(urlPath.saveMailUrl, data, true);
};

/**
 * 消息设置-是否关联邮箱推送
 * @param {Object} data 请求需要的参数
*/
const savelsPushMailAPI = (data: {is_push: string;}) => {
    return post(urlPath.savelsPushMailUrl, data, true);
};

const messageApiCount = {
    messageListAPI,
    messageExportAPI,
    messagesStatisticAPI,
    messageTypesAPI,
    hasUnreadMessageAPI,
    messageByIdAPI,
    doReadAPI,
    deleteAPI,
    describeMailAPI,
    dialMailAPI,
    saveMailAPI,
    savelsPushMailAPI
};

export default messageApiCount;
