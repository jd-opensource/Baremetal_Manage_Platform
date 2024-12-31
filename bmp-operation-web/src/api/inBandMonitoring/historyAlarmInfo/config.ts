/**
 * @file
 * @author
*/

import {TotalUrlPathType} from './typeManagement'; // url地址类

const dafaultUrl: string = '/monitorAlert/describeAlerts';

// url地址集合
const urlPath: TotalUrlPathType =  {
    describeAlertsUrl: dafaultUrl, // 报警历史列表
    describeAlertsExportUrl: dafaultUrl // 报警历史列表导出
};

export default urlPath;
