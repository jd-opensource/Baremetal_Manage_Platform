/**
 * @file
 * @author
*/

import {TotalUrlPathType} from './typeManagement'; // url地址类

const defaultUrl: string =  '/license/';

// url地址集合
const urlPath: TotalUrlPathType =  {
    licenseUrl: `${defaultUrl}content`, // 授权证内容
    downloadTokenUrl: `${defaultUrl}downloadToken`
};

export default urlPath;
