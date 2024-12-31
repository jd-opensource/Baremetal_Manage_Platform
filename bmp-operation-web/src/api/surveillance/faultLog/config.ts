/**
 * @file
 * @author
*/

import {methodsTotal} from 'utils/index.ts';
import {TotalUrlPathType} from './typeManagement'; // url地址类

const defaultUrl: string =  methodsTotal.humpSplit('/v1/oobAlert/logs');

// url地址集合
const urlPath: TotalUrlPathType =  {
    faultLogUrl: defaultUrl, // 故障报警日志
    faultLogExportUrl: defaultUrl, // 故障报警日志导出
    faultLogDealUrl: `${defaultUrl}/deal`, // 故障报警日志处理
};

export default urlPath;
