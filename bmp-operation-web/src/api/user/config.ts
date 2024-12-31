/**
 * @file
 * @author
*/

import {TotalUrlPathType} from './typeManagement'; // url地址类

const defaultPath: string = '/users';

// url地址集合
const urlPath: TotalUrlPathType =  {
    userListUrl: defaultPath, // 用户列表
    addUserUrl: defaultPath, // 添加用户
    editUserUrl: `${defaultPath}/`, // 编辑用户
    deleteUserUrl: `${defaultPath}/`, // 删除用户
};

export default urlPath;
