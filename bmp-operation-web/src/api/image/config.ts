/**
 * @file
 * @author
*/

import {TotalUrlPathType} from './typeManagement'; // url地址类

const defaultUrl: string = '/images';

// url地址集合
const urlPath: TotalUrlPathType =  {
    imagesUrl: defaultUrl, // 镜像列表
    imagesDetailUrl: `${defaultUrl}/`, // 镜像详情
    imageDeviceTypesUrl: `${defaultUrl}/imageDeviceTypes`, // 镜像详情-机型列表、添加机型列表
    imageAddModelUrl: `${defaultUrl}/associatedDeviceType`, // 镜像详情-机型列表、添加机型
    deleteImageModelUrl: `${defaultUrl}/dissociatedDeviceType`, // 镜像详情-删除关联机型
    imageEditUrl: `${defaultUrl}/`, // 镜像编辑
    imagesDeleteUrl: `${defaultUrl}/`, // 删除镜像
    importImageUrl: defaultUrl, // 导入镜像
    imageOssUrl: '/oss/filter' // 镜像oss数据
};

export default urlPath;
