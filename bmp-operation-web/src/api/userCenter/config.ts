/**
 * @file
 * @author
*/
import {TotalUrlPathType} from './typeManagement'; // url地址类

const defaultUrl1: string = '/user/';
const defaultUrl2: string = '/apikey';

// url地址集合
const urlPath: TotalUrlPathType =  {
    timeOneUrl: `${defaultUrl1}local/timezones`, // 时区
    userInfoUrl: `${defaultUrl1}local`, // 用户信息
    setUserInfoUrl: `${defaultUrl1}local`, // 修改用户信息
    deleteApiKeyUrl: `${defaultUrl2}/`, // 删除apikey
    createApiKeyUrl: defaultUrl2, // 创建apikey
    getApiKeyUrl: defaultUrl2, // 获取apikey
    revisePasswordUrl: `${defaultUrl1}local/password` // 修改密码
};

export default urlPath;
