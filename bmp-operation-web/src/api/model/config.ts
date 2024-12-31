/**
 * @file
 * @author
*/

import {TotalUrlPathType} from './typeManagement'; // url地址类

const defaultUrl: string = '/deviceTypes';

// url地址集合
const urlPath: TotalUrlPathType =  {
    modelListUrl: defaultUrl, // 机型列表
    modelDetailUrl: `${defaultUrl}/queryDeviceType/`, // 机型详情
    relationImageListUrl: `${defaultUrl}/deviceTypeImage`, // 机型详情-关联镜像列表
    addImageUrl: `${defaultUrl}/associatedImage`, // 机型详情-添加镜像
    addModelUrl: defaultUrl, // 添加机型
    editModelUrl: `${defaultUrl}/`, // 编辑机型
    deleteImageUrl: `${defaultUrl}/dissociatedImage`, // 删除关联镜像
    deleteModelUrl: `${defaultUrl}/`, // 删除机型
    getRaidsUrl: '/raids', // 获取raid
    filterListUrl: '/filterList', // 筛选list
};

export default urlPath;
