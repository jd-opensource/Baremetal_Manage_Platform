/**
 * @file
 * @author
*/

import {methodsTotal} from 'utils/index.ts';
import {TotalUrlPathType} from './typeManagement'; // url地址类

const defaultUrl: string = methodsTotal.humpSplit('/v1/oobAlert/device/status');

// url地址集合
const urlPath: TotalUrlPathType =  {
    hardwareStatusUrl: defaultUrl, // 硬件报警状态列表
    hardwareStatusExportUrl: defaultUrl // 硬件报警状态列表导出
};

export default urlPath;
