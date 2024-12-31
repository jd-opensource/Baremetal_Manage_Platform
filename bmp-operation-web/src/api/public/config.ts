/**
 * @file
 * @author
*/

import {methodsTotal} from 'utils/index.ts';
import {TotalUrlPathType} from './typeManagement'; // url地址类

// url地址集合
const urlPath: TotalUrlPathType =  {
    resourcesUrl: '/resources', // 精确查询
    securityVerificationUrl: '/users/login/validate', // 安全验证
    customListUrl: '/custom/getCustomInfo', // 自定义列表
    setCustomListUrl: '/custom/setCustomInfo', // 设置自定义列表
    surveillanceCustomListUrl: methodsTotal.humpSplit('/v1/oobAlert/custom/getCustomInfo'), // 监控自定义列表
    surveillanceSetCustomListUrl: methodsTotal.humpSplit('/v1/oobAlert/custom/setCustomInfo'), // 监控设置自定义列表
    getOssUrl: '/oss', // oss数据
    queryImagesUrl: '/images/queryImagesByDeviceType', // images类型
    systemRaidsUrl: '/raids/queryRaidsByDeviceTypeIDAndVolumeType', // raids
    querySystemPartitionUrl: '/partition/queryDefaultSystemPartitions', // 系统盘分区
    keyInfoUrl: '/keypair', // 密钥
    localUrl: '/user/local ', // 权限列表
};

export default urlPath;
