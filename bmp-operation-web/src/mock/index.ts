/**
 * @file
 * @author
*/

import axios, {AxiosError, AxiosResponse, AxiosInstance, AxiosRequestConfig} from 'axios';
import {ElMessage} from 'element-plus';

const config = {
    baseURL: '',
    timeout: 10 * 1000, // 超时时间
    headers: { // 请求头
        'Content-Type': 'application/json; charset=utf-8'
    }
};

class MockHttpRequest {
    service: AxiosInstance;

    constructor (config: AxiosRequestConfig<never>) {
        // 创建axios实例 
		this.service = axios.create(config);
        this.requestIntercept();
        this.responseIntercept();
    };

    requestIntercept = () => {
        this.service?.interceptors?.request?.use(
            (config: AxiosRequestConfig): AxiosRequestConfig => {
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
                return response.data;
            },
            (err): Promise<unknown> | void => {
                ElMessage.error(err);
                return Promise.reject(err);
            }
        );
    };
};

export default new MockHttpRequest(config);