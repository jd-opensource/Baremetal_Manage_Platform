/**
 * @file
 * @author
*/

import {TotalUrlPathType} from './typeManagement'; // url地址类

const dafaultUrl: string = '/idc';

// url地址集合
const urlPath: TotalUrlPathType =  {
    idcListUrl: dafaultUrl, // 机房列表
    idcDetailUrl: dafaultUrl + '/', // 机房详情
    editIdcUrl: dafaultUrl + '/', // 编辑机房
};

export default urlPath;
