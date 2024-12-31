/**
 * @file
 * @author
*/

// 暂时用不到
// import ErrorCode from './error/errorCode';
import axios, {AxiosError, AxiosResponse, AxiosInstance, AxiosRequestConfig, AxiosResponseHeaders} from 'axios';
import * as types from './type';
import Error402 from './error/402';
import router from '@/router/index.ts'; // 路由
import OtherError from './error/otherError';
import InitConfig from './config/initConfig';
import CancelRuquest from 'request/cancel/cancel.ts';
import {locationItem} from 'utils/publicClass.ts';
import {CurrencyType, CurrencyType3} from '@utils/publicType';

/**
 * 封装axios请求
*/
class HttpRequest {
    service: AxiosInstance;
    axiosCancelReq = new CancelRuquest();

    constructor (config: AxiosRequestConfig<string>) {
        // 创建axios实例 
		this.service = axios.create(config);
        this.requestIntercept();
        this.responseIntercept();
    };

    /**
     * 请求拦截器
    */
    requestIntercept = () => {
        this.service?.interceptors?.request?.use(
            (config: AxiosRequestConfig & {islongReq?: boolean}): AxiosRequestConfig => {
                const userName: string = localStorage.getItem('userName')??'';
                locationItem.cookie.set('X-Jdcloud-Username', window.atob(userName));
                // 删除重复请求
                this.axiosCancelReq.removeRequest(config);
                const {name}: {name: string;} = router.currentRoute.value;
                this.axiosCancelReq.cancelReq(name, config);
                return config;
            },
            (error: AxiosError): Promise<never>  => {
                return Promise.reject(error);
            }
        );
    };

    /**
     * 响应拦截器
    */
    responseIntercept = () => {
        this.service?.interceptors?.response?.use(
            (response: AxiosResponse): Promise<unknown> => {
                // 每成功一次，数组边为空，否则其他接口重复请求会数据一直叠加
                this.axiosCancelReq.repeatRequestArr = [];
                const {data, headers} : {data: CurrencyType3, headers: AxiosResponseHeaders} = response;
                // 登录用户名称
                const loginUserName: string = headers['x-jdcloud-username'];
                return Promise.resolve({data, loginUserName});
            },
            (err: {response: types.ErrResponseType}): Promise<unknown> | void => {
                const {response} : {response: types.ErrResponseType} = err;
                const errInfo: types.ErrInfoType = JSON.parse(JSON.stringify(err));
                const {config, message}: types.ErrInfoType = errInfo;
                const errorMsg: string = response?.data?.error?.message?? '';
                // 402登录过期-未登录
                if (Object.is(response?.status, 402)) {
                    new Error402();
                    return Promise.resolve({data: {}});
                }
                if (axios.isCancel(err)) {
                    return new OtherError().cancelError();
                };
                return new OtherError().setError(errorMsg || message, config, response);
            }
        );
    };

    get = (url: string, params: CurrencyType, otherParams: AxiosRequestConfig<types.OtherParamsType>): Promise<void> => {
		return this.service.get(url, {params, ...otherParams});
	};

    post = <T extends object>(url: string, data: CurrencyType, status: boolean, otherParams: T): Promise<void> => {
		return this.service.post(url, data, Object.assign({isBusinessProcessing: status}, {...otherParams}));
	};

    put = <T extends object>(url: string, data: CurrencyType, status: boolean, otherParams: T): Promise<void> => {
		return this.service.put(url, data, Object.assign({isBusinessProcessing: status}, {...otherParams}));
	};

    deleteReq = <T extends object>(url: string, data: CurrencyType, status: boolean, otherParams: T): Promise<void> => {
        return this.service.delete(url, Object.assign({data, isBusinessProcessing: status}, {...otherParams}));
    };
}

export default new HttpRequest(new InitConfig().config);
