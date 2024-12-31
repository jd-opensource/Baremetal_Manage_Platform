import request from 'request/index.ts';
import {AxiosPromise} from 'axios';
import urlPath from 'api/config.ts';

type SysPartitionType = {
    format: string;
    point: string;
    size: number;
}
type SetType<T> = T extends {} ? any : T;

export interface TotalTypes {
    LoginType: {
        username: string;
        password: string;
    },
    UserType:{},
    timeZoneType:{},
    instanceCloseType:{
        instanceId: string;
        idcId: string
    },
    instanceDeleteType:{
        instanceId: string;
    },
    instanceOpenType:{
        instanceId: string;
        idcId: string
    },
    instanceRestartType:{
        instanceId: string;
        idcId: string
    },
    instanceLockType:{
        instanceId: string;
        idcId: string
    },
    instanceUnlockType:{
        instanceId: string;
        idcId: string
    },
    instanceRepasswordType:{
        instanceId: string;
        idcId: string;
        password: string
    },
    apiKeyCreateType:{
        keyName: string;
        readOnly: number;
        type: string;
    },
    apiKeyDeleteType:{
        apikeyId: string
    },
    sshKeyDeleteType:{
        keypairId: string
    },
    sshKeyCreateType:{
        key: string;
        name: string
    },
    sshKeyEditType:{
        name: string;
        keypairId: string
    },
    projectCreateType:{
        isDefault?: 0;
        isSystem?: 0;
        projectId: string;
        projectName: string;
        userId: string;
    },
    projectEditType: {
        projectId: string;
    },
    projectDeleteType:{
        projectId: string;
    },
    az:{},
    deviceType:{
        idcId: string;
    },
    osType:{
        idcId: string;
        deviceTypeId: string;
    },
    raid:{
        deviceTypeId: string;
        volumeType: string;
    },
    systemDiskPartition:{
        imageId: string;
        deviceTypeId: string;
    },
    createInstance:{
        count: number;
        deviceTypeId: string;
        hostname?: string;
        idcId: string;
        imageId: string
        instanceName: string;
        password?: string;
        sshKeyId?: string;
        systemVolumeRaidId: string;
        projectId: string;
        systemPartition: SysPartitionType[];
        username: string
    },
    isDeviceStockEnough:{
        idcId: string;
        deviceTypeId: string;
        count: string
    },
    instanceList:{
        pageNumber?: number;
        pageSize?: number;
        projectId: string;
        status: string;
        exportType?: string;
        isAll?: string
    },
    instanceDetail:{
        instanceId: string;
    },
    areaCode:{},
    setUser:{
        defaultProjectId: string;
        description?: string;
        email: string;
        language: string;
        phoneNumber: string;
        phonePrefix: string;
        timezone: string;
        userName: string;
    },
    changePassword:{
        oldPassword: string;
        password: string;
    },
    sshKey:{
        pageNumber: number;
        pageSize: number;
    },
    apiKey:{
        pageNumber: number;
        pageSize: number;
    },
    projectList:{
        projectName?: string;
        pageNumber: number;
        pageSize: number;
        isAll?: string;
    },
    logOutType:{},
    setCustomListType: Required<{pageKey: string; pageValue: SetType<{}>;}>,
    setAlarmLogCustomListType: Required<{page_key: string; page_value: SetType<{}>;}>,
    customListType: Omit<TotalTypes['setCustomListType'], 'pageValue'>,
    alarmLogCustomListType: {
        page_key: string;
    },
    projectListExportType: {
        exportType: string;
        isAll: string;
    },
    instanceListExportType: {
        exportType: string;
        projectId: string;
        isAll: string;
    },
    editInstanceType: {
        instanceId: string,
        idcId: string,
        description?: string,
        name?: string
    },
    editProjectType: {
        projectId: string,
        description?: string,
        name?: string
    },
    projectDetail: {
        projectId: string;
    },
    relatedUserAddType: {
        projectId: string;
        ownerName: string;
        shareName: string
    },
    relatedUserDeleteType: {
        projectId: string;
        ownerName: string;
        shareName: string
    },
    transferUserType: {
        projectId: string;
        ownerName: string;
        moverName: string
    },
    transferUserCheckType: {
        projectId: string;
        checkedUserName: string;
        operation: string;
    },
    relatedUserCheckType: {
        projectId: string;
        checkedUserName: string;
        operation: string;
    },
    relatedUserList: {

    },
    batchCloseType: {
        instanceIds: []
    },
    batchOpenType: {
        instanceIds: []
    },
    batchRestartType: {
        instanceIds: []
    },
    batchDeleteType: {
        instanceIds: []
    },
    batchRenameType: {
        instanceIds: [],
        instanceName: string
    },
    batchRepasswordType: {
        instanceIds: [],
        password: string
    },
    hardWareStatusType: {
        sn: string
    },
    alarmLogListType: {
        page_size: number,
        page_num: number,
        sn: string
    },
    alarmLogListExportType: {
        
    },
    resystemInstance: {
        bootMode: string;
        instanceId: string;
        instanceName: string;
        hostname?: string;
        imageId: string
        password?: string;
        sshKeyId?: string;
        systemVolumeRaidId: string;
        projectId: string;
        systemPartition: SysPartitionType[];
    },
    shareInstanceListType: {
        projectId: string;
        ownerName: string;
        sharerName: string;
    },
    transferPartUserType: {
        instanceIDs: string,
        moverName: string,
        moverProjectID: string,
        ownerName: string,
        ownerProjectID: string
    },
    messageType: {
        content: string,
        detail: string,
        has_read: number,
        is_del: number,
        message_id: string,
        message_sub_type: string,
        message_type: string,
    },
    messageDetail: {
        content: string,
        detail: string,
        has_read: number,
        is_del: number,
        messageId: string,
        message_sub_type: string,
        message_type: string,
    },
    messageReadType: {
        messageIds: string
    },
    messageDeleteType: {
        messageIds: string
    },
    messageListExportType: {
        exportType: string;
    },
    operationLogType: {
        instanceId: string;
        pageSize: string;
        pageNumber: string;
        operation?: string;
        username?: string
    },
    operationLogListExportType: {
        all: string;
        exportType: string;
    },
    deviceTagType: {
        instanceId: string;
        tagName: string;
    },
    addRuleType: {
        
    },
    alarmRuleListType: {},
    ruleEnableType: {
        ruleId: string
    },
    monitorDataType: [],

}

/**
 * 登录接口
 * @param {Object} data 请求需要的参数
*/
export function loginAPI(data: TotalTypes['LoginType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.loginUrl,
        method: 'post',
        data
    });
}

/**
 * 用户接口
 * @param {Object} params 请求需要的参数
*/
export function userAPI(params: TotalTypes['UserType']): AxiosPromise<any> {
    return request({
        isrepeat: true,
        isBusinessProcessing: false,
        url: urlPath.userUrl,
        method: 'get',
        params
    });
}
/**
 * 时区获取接口
 * @param {Object} params 请求需要的参数
*/
export function timeZoneAPI(params: TotalTypes['timeZoneType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.timeZoneUrl,
        method: 'get',
        params
    });
}

/**
 * 实例关机接口
 * @param {Object} data 请求需要的参数
*/
export function instanceCloseAPI(data: TotalTypes['instanceCloseType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.instanceCloseUrl,
        method: 'post',
        data
    });
}

/**
 * 实例删除接口
 * @param {Object} params 请求需要的参数
*/
export function instanceDeleteAPI(params: TotalTypes['instanceDeleteType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.instanceDeleteUrl + params.instanceId,
        method: 'delete',
        params
    });
}

/**
 * 实例开机接口
 * @param {Object} data 请求需要的参数
*/
export function instanceOpenAPI(data: TotalTypes['instanceOpenType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.instanceOpenUrl,
        method: 'post',
        data
    });
}

/**
 * 实例重启接口
 * @param {Object} data 请求需要的参数
*/
export function instanceRestartAPI(data: TotalTypes['instanceRestartType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.instanceRestartUrl,
        method: 'post',
        data
    });
}

/**
 * 实例锁定接口
 * @param {Object} data 请求需要的参数
*/
export function instanceLockAPI(data: TotalTypes['instanceLockType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.instanceLockUrl,
        method: 'post',
        data
    });
}

/**
 * 实例解除锁定接口
 * @param {Object} data 请求需要的参数
*/
export function instanceUnlockAPI(data: TotalTypes['instanceUnlockType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.instanceUnlockUrl,
        method: 'post',
        data
    });
}

/**
 * 实例重置密码接口
 * @param {Object} data 请求需要的参数
*/
export function instanceRepasswordAPI(data: TotalTypes['instanceRepasswordType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.instanceRepasswordUrl,
        method: 'post',
        data
    });
}

/**
 * api密钥创建接口
 * @param {Object} data 请求需要的参数
*/
export function apiKeyCreateAPI(data: TotalTypes['apiKeyCreateType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.apiKeyCreateUrl,
        method: 'put',
        data
    });
}

/**
 * api密钥删除接口
 * @param {Object} params 请求需要的参数
*/
export function apiKeyDeleteAPI(params: TotalTypes['apiKeyDeleteType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.apiKeyDeleteUrl + params.apikeyId,
        method: 'delete',
        params
    });
}

/**
 * ssh密钥删除接口
 * @param {Object} params 请求需要的参数
*/
export function sshKeyDeleteAPI(params: TotalTypes['sshKeyDeleteType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.sshKeyDeleteUrl + params.keypairId,
        method: 'delete',
        params
    });
}

/**
 * ssh密钥创建接口
 * @param {Object} data 请求需要的参数
*/
export function sshKeyCreateAPI(data: TotalTypes['sshKeyCreateType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.sshKeyCreateUrl,
        method: 'put',
        data
    });
}

/**
 * ssh密钥编辑接口
 * @param {Object} data 请求需要的参数
*/
export function sshKeyEditAPI(data: TotalTypes['sshKeyEditType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.sshKeyEditUrl + data.keypairId,
        method: 'put',
        data
    });
}

/**
 * 项目创建接口
 * @param {Object} data 请求需要的参数
*/
export function projectCreateAPI(data: TotalTypes['projectCreateType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.projectCreateUrl,
        method: 'put',
        data
    });
}

/**
 * 项目编辑接口
 * @param {Object} data 请求需要的参数
*/
export function projectEditAPI(data: TotalTypes['projectEditType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.projectEditUrl + data.projectId,
        method: 'post',
        data
    });
}

/**
 * 项目删除接口
 * @param {Object} params 请求需要的参数
*/
export function projectDeleteAPI(params: TotalTypes['projectDeleteType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.projectDeleteUrl + params.projectId,
        method: 'delete',
        params
    });
}

/**
 * 机房信息接口
 * @param {Object} params 请求需要的参数
*/
export function azAPI(params: TotalTypes['az']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.azUrl,
        method: 'get',
        params
    });
}

/**
 * 机型信息接口
 * @param {Object} params 请求需要的参数
*/
export function deviceTypeAPI(params: TotalTypes['deviceType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.deviceTypeUrl,
        method: 'get',
        params
    });
}

/**
 * 镜像信息接口
 * @param {Object} params 请求需要的参数
*/
export function osTypeAPI(params: TotalTypes['osType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.osTypeUrl,
        method: 'get',
        params
    });
}

/**
 * 系统盘信息接口
 * @param {Object} params 请求需要的参数
*/
export function raidAPI(params: TotalTypes['raid']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.raidUrl,
        method: 'get',
        params
    });
}

/**
 * 系统盘分区信息接口
 * @param {Object} params 请求需要的参数
*/
export function systemDiskPartitionAPI(params: TotalTypes['systemDiskPartition']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.systemDiskPartitionUrl,
        method: 'get',
        params
    });
}

/**
 * 创建实例接口
 * @param {Object} data 请求需要的参数
*/
export function createInstanceAPI(data: TotalTypes['createInstance']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.createInstanceUrl,
        method: 'put',
        data
    });
}

/**
 * 实例库存接口
 * @param {Object} params 请求需要的参数
*/
export function isDeviceStockEnoughAPI(params: TotalTypes['isDeviceStockEnough']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.isDeviceStockEnoughUrl,
        method: 'get',
        params
    });
}

/**
 * 实例列表接口
 * @param {Object} params 请求需要的参数
*/
export function instanceListAPI(params: TotalTypes['instanceList']): AxiosPromise<any> {
    return request({
        isrepeat: true,
        isBusinessProcessing: false,
        url: urlPath.instanceListUrl,
        method: 'get',
        params
    });
}

/**
 * 实例详情接口
 * @param {Object} params 请求需要的参数
*/
export function instanceDetailAPI(params: TotalTypes['instanceDetail']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.instanceDetailUrl + params.instanceId,
        method: 'get',
        params
    });
}

/**
 * 保存用户接口
 * @param {Object} data 请求需要的参数
*/
export function setUserAPI(data: TotalTypes['setUser']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.setUserUrl,
        method: 'put',
        data
    });
}

/**
 * 更改密码接口
 * @param {Object} data 请求需要的参数
*/
export function changePasswordAPI(data: TotalTypes['changePassword']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.changePasswordUrl,
        method: 'put',
        data
    });
}

/**
 * ssh个人密钥接口
 * @param {Object} params 请求需要的参数
*/
export function sshKeyAPI(params: TotalTypes['sshKey']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.sshKeyUrl,
        method: 'get',
        params
    });
}

/**
 * api个人密钥接口
 * @param {Object} params 请求需要的参数
*/
export function apiKeyAPI(params: TotalTypes['apiKey']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.apiKeyUrl,
        method: 'get',
        params
    });
}

/**
 * 项目列表接口
 * @param {Object} params 请求需要的参数
*/
export function projectListAPI(params: TotalTypes['projectList']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.projectListUrl,
        method: 'get',
        params
    });
}

/**
 * 用户登出接口
 * @param {Object} data 请求需要的参数
*/
export function logOutAPI(data: TotalTypes['logOutType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.logOutUrl,
        method: 'post',
        data
    });
}

/**
 * 自定义列表接口
 * @param {Object} data 请求需要的参数
*/
export function setCustomListAPI(data: TotalTypes['setCustomListType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.setCustomListUrl,
        method: 'post',
        data
    });
}

/**
 * 报警日志自定义列表接口
 * @param {Object} data 请求需要的参数
*/
export function setAlarmCustomListAPI(data: TotalTypes['setAlarmLogCustomListType']): AxiosPromise<any> {
    return request({
        baseURL: '/oob-alert',
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.setAlarmCustomListUrl,
        method: 'post',
        data
    });
}
/**
 * 获取自定义列表接口
 * @param {Object} params 请求需要的参数
*/
export function customListAPI(params: TotalTypes['customListType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.customListUrl,
        method: 'get',
        params
    });
}

/**
 * 获取报警日志自定义列表接口
 * @param {Object} params 请求需要的参数
*/
export function alarmCustomListAPI(params: TotalTypes['alarmLogCustomListType']): AxiosPromise<any> {
    return request({
        baseURL: '/oob-alert',
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.alarmCustomListUrl,
        method: 'get',
        params
    });
}

/**
 * 项目列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function projectListExportAPI(params: TotalTypes['projectListExportType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.projectListUrl,
        method: 'get',
        params,
        // 服务器响应的数据类型，默认json
        responseType: 'blob'
    });
};

/**
 * 项目列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function instanceListExportAPI(params: TotalTypes['instanceListExportType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.instanceListUrl,
        method: 'get',
        params,
        // 服务器响应的数据类型，默认json
        responseType: 'blob'
    });
};


/**
 * 编辑实例接口
 * @param {Object} data 请求需要的参数
*/
export function editInstanceAPI(data: TotalTypes['editInstanceType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.editInstanceUrl,
        method: 'post',
        data
    });
}

/**
 * 编辑项目接口
 * @param {Object} data 请求需要的参数
*/
export function editProjectAPI(data: TotalTypes['editProjectType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.editProjectUrl + data.projectId + '/description',
        method: 'post',
        data
    });
}

/**
 * 项目详情接口
 * @param {Object} params 请求需要的参数
*/
export function projectDetailAPI(params: TotalTypes['projectDetail']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.projectDetailUrl + params.projectId,
        method: 'get',
        params
    });
}

/**
 * 添加共享用户接口
 * @param {Object} data 请求需要的参数
*/
export function relatedUserAddAPI(data: TotalTypes['relatedUserAddType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.relatedUserAddUrl + data.projectId + '/share',
        method: 'put',
        data
    });
}

/**
 * 删除共享用户接口
 * @param {Object} data 请求需要的参数
*/
export function relatedUserDeleteAPI(data: TotalTypes['relatedUserDeleteType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.relatedUserDeleteUrl + data.projectId + '/cancelShare',
        method: 'put',
        data
    });
}

/**
 * 转移共享用户接口
 * @param {Object} data 请求需要的参数
*/
export function transferUserAPI(data: TotalTypes['transferUserType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.transferUserUrl + data.projectId + '/move',
        method: 'put',
        data
    });
}

/**
 * 验证转移用户接口
 * @param {Object} params 请求需要的参数
*/
export function transferUserCheckAPI(params: TotalTypes['transferUserCheckType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.transferUserCheckUrl,
        method: 'get',
        params
    });
}

/**
 * 验证关联用户接口
 * @param {Object} params 请求需要的参数
*/
export function relatedUserCheckAPI(params: TotalTypes['relatedUserCheckType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.relatedUserCheckUrl,
        method: 'get',
        params
    });
}

/**
 * 批量关机接口
 * @param {Object} data 请求需要的参数
*/
export function batchCloseAPI(data: TotalTypes['batchCloseType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.batchCloseUrl,
        method: 'post',
        data
    });
}

/**
 * 批量开机接口
 * @param {Object} data 请求需要的参数
*/
export function batchOpenAPI(data: TotalTypes['batchOpenType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.batchOpenUrl,
        method: 'post',
        data
    });
}

/**
 * 批量重启接口
 * @param {Object} data 请求需要的参数
*/
export function batchRestartAPI(data: TotalTypes['batchRestartType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.batchRestartUrl,
        method: 'post',
        data
    });
}

/**
 * 批量删除接口
 * @param {Object} data 请求需要的参数
*/
export function batchDeleteAPI(data: TotalTypes['batchDeleteType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.batchDeleteUrl,
        method: 'delete',
        data
    });
}

/**
 * 批量重置密码接口
 * @param {Object} data 请求需要的参数
*/
export function batchRepasswordAPI(data: TotalTypes['batchRepasswordType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.batchRepasswordUrl,
        method: 'post',
        data
    });
}

/**
 * 批量编辑实例名称接口
 * @param {Object} data 请求需要的参数
*/
export function batchRenameAPI(data: TotalTypes['batchRenameType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.batchRenameUrl,
        method: 'post',
        data
    });
}

/**
 * 硬件设备状态接口
 * @param {Object} params 请求需要的参数
*/
export function hardWareStatusAPI(params: TotalTypes['hardWareStatusType']): AxiosPromise<any> {
    return request({
        baseURL: '/oob-alert',
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.hardWareStatusUrl,
        method: 'get',
        params
    });
}

/**
 * 报警日志列表接口
 * @param {Object} params 请求需要的参数
*/
export function alarmLogListAPI(params: TotalTypes['alarmLogListType']): AxiosPromise<any> {
    return request({
        baseURL: '/oob-alert',
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.alarmLogListUrl,
        method: 'get',
        params
    });
}

/**
 * 报警日志列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function alarmLogListExportAPI(params: TotalTypes['alarmLogListExportType']): AxiosPromise<any> {
    return request({
        baseURL: '/oob-alert',
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.alarmLogListExportUrl,
        method: 'get',
        params,
        // 服务器响应的数据类型，默认json
        responseType: 'blob'
    });
};

/**
 * 重装系统实例接口
 * @param {Object} data 请求需要的参数
*/
export function resystemInstanceAPI(data: TotalTypes['resystemInstance']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.resystemInstanceUrl,
        method: 'post',
        data
    });
}

/**
 * 共享实例列表接口
 * @param {Object} params 请求需要的参数
*/
export function shareInstanceListAPI(params: TotalTypes['shareInstanceListType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.shareInstanceListUrl,
        method: 'get',
        params
    });
}

/**
 * 转移部分共享用户接口
 * @param {Object} data 请求需要的参数
*/
export function transferPartUserAPI(data: TotalTypes['transferPartUserType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.transferPartUserUrl,
        method: 'put',
        data
    });
}

/**
 * 消息列表接口
 * @param {Object} params 请求需要的参数
*/
export function messageListAPI(params: TotalTypes['messageType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.messageListUrl,
        method: 'get',
        params
    });
}

/**
 * 消息详情接口
 * @param {Object} params 请求需要的参数
*/
export function messageDetailAPI(params: TotalTypes['messageDetail']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.messageDetailUrl,
        method: 'get',
        params
    });
}

/**
 * 消息类型接口
 * @param {Object} params 请求需要的参数
*/
export function messageTypeAPI(params: TotalTypes['messageType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.messageTypeUrl,
        method: 'get',
        params
    });
}

/**
 * 消息已读接口
 * @param {Object} data 请求需要的参数
*/
export function messageReadAPI(data: TotalTypes['messageReadType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.messageReadUrl,
        method: 'put',
        data
    });
}


/**
 * 消息类型接口
 * @param {Object} params 请求需要的参数
*/
export function getMessageStatisticAPI(): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.getMessageStatisticUrl,
        method: 'get',
    });
}

/**
 * 消息未读接口
 * @param {Object} params 请求需要的参数
*/
export function messageUnreadAPI(): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.messageUnreadUrl,
        method: 'get',
    });
}

/**
 * 消息删除接口
 * @param {Object} data 请求需要的参数
*/
export function messageDeleteAPI(data: TotalTypes['messageDeleteType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.messageDeleteUrl,
        method: 'delete',
        data
    });
}

/**
 * 消息列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function messageListExportAPI(params: TotalTypes['messageListExportType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.messageListUrl,
        method: 'get',
        params,
        // 服务器响应的数据类型，默认json
        responseType: 'blob'
    });
};

/**
 * 操作日志类型接口
 * @param {Object} params 请求需要的参数
*/
export function operationLogTypeListAPI(): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.operationLogTypeListUrl,
        method: 'get',
    });
}

/**
 * 操作日志列表接口
 * @param {Object} params 请求需要的参数
*/
export function operationLogListAPI(params: TotalTypes['operationLogType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.operationLogListUrl,
        method: 'get',
        params
    });
}

/**
 * 操作日志列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function operationLogListExportAPI(params: TotalTypes['operationLogListExportType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.operationLogListExportUrl,
        method: 'get',
        params,
        // 服务器响应的数据类型，默认json
        responseType: 'blob'
    });
};

/**
 * 添加报警规则deviceTag接口
 * @param {Object} params 请求需要的参数
*/
export function deviceTagAPI(params: TotalTypes['deviceTagType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.deviceTagUrl,
        method: 'get',
        params
    });
}

/**
 * 添加报警规则列表接口
 * @param {Object} data 请求需要的参数
*/
export function addRuleAPI(data: TotalTypes['addRuleType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.addRuleUrl,
        method: 'post',
        data
    });
}

/**
 * 报警规则列表接口
 * @param {Object} params 请求需要的参数
*/
export function alarmRuleListAPI(params: TotalTypes['alarmRuleListType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.alarmRuleListUrl,
        method: 'get',
        params
    });
}

/**
 * 报警规则列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function alarmRuleListExportAPI(params: TotalTypes['operationLogListExportType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.alarmRuleListUrl,
        method: 'get',
        params,
        // 服务器响应的数据类型，默认json
        responseType: 'blob'
    });
};


/**
 * 报警规则详情接口
 * @param {Object} params 请求需要的参数
*/
export function alarmRuleDetailAPI(params: TotalTypes['alarmRuleListType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.alarmRuleDetailUrl,
        method: 'get',
        params
    });
}

/**
 * 报警规则启用接口
 * @param {Object} data 请求需要的参数
*/
export function ruleEnableAPI(data: TotalTypes['ruleEnableType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.ruleEnableUrl,
        method: 'put',
        data
    });
}

/**
 * 报警规则禁用接口
 * @param {Object} data 请求需要的参数
*/
export function ruleDisableAPI(data: TotalTypes['ruleEnableType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.ruleDisableUrl,
        method: 'put',
        data
    });
}

/**
 * 报警规则删除接口
 * @param {Object} data 请求需要的参数
*/
export function ruleDeleteAPI(data: TotalTypes['ruleEnableType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.ruleDeleteUrl,
        method: 'delete',
        data
    });
}

/**
 * 报警规则编辑接口
 * @param {Object} data 请求需要的参数
*/
export function editRuleAPI(data: TotalTypes['ruleEnableType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: true,
        url: urlPath.editRuleUrl,
        method: 'put',
        data
    });
}


/**
 * 获取agent状态接口
 * @param {Object} params 请求需要的参数
*/
export function getAgentStateAPI(params: TotalTypes['alarmRuleListType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.getAgentStateUrl,
        method: 'get',
        params
    });
}

/**
 * 获取报警历史列表接口
 * @param {Object} params 请求需要的参数
*/
export function alarmHistoryAPI(params: TotalTypes['alarmRuleListType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.alarmHistoryUrl,
        method: 'get',
        params
    });
}

/**
 * 操作日志列表导出接口
 * @param {Object} params 请求需要的参数
*/
export function alarmHistoryListExportAPI(params: TotalTypes['operationLogListExportType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.alarmHistoryUrl,
        method: 'get',
        params,
        // 服务器响应的数据类型，默认json
        responseType: 'blob'
    });
};

/**
 * 获取带内监控列表接口
 * @param {Object} params 请求需要的参数
*/
export function monitorDataAPI(params: TotalTypes['monitorDataType']): AxiosPromise<any> {
    return request({
        isrepeat: false,
        isBusinessProcessing: false,
        url: urlPath.monitorDataUrl,
        method: 'get',
        params
    });
}