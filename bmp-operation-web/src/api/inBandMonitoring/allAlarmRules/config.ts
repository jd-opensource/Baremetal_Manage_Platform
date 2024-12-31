/**
 * @file
 * @author
*/

import {TotalUrlPathType} from './typeManagement'; // url地址类

const dafaultUrl: string = '/monitorRule/describeRules';

// url地址集合
const urlPath: TotalUrlPathType =  {
    describeRulesUrl: dafaultUrl, // 全部报价规则列表
    describeRulesExportUrl: dafaultUrl, // 全部报价规则列表导出
    describeRulesDetailUrl: '/monitorRule/describeRule' // 全部报警规则详情
};

export default urlPath;
