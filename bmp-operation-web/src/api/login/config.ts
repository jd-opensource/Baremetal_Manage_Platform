/**
 * @file
 * @author
*/

import {TotalUrlPathType} from './typeManagement'; // url地址类

// url地址集合
const urlPath: TotalUrlPathType =  {
    loginUrl: '/login', // 登录
    userPurviewUrl: '/roles/roleInfo/current', // 用户权限
    loginOutUrl: '/logout' // 退出登录
};

export default urlPath;
