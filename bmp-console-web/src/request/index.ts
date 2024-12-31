import axios, {
    AxiosInstance,
    AxiosRequestConfig,
    AxiosResponse,
    AxiosRequestHeaders,
    AxiosError,
    CancelTokenStatic,
    AxiosResponseHeaders // axios相应头信息
} from 'axios';
import {ElMessage} from 'element-plus';
import VueCookies from 'vue-cookies';
import router from '@/router/index.ts'; // 路由
interface RepeatRequestArrType {
    methodUrl: string;
    params: any;
    f: Function;
};

/**
 * cookie ts规范校验
*/
const cookie: {
    [x: string]: unknown;
    get?: Function;
    set?: Function;
} = VueCookies;

/**
 * 国际化
 */
 const vueCookiesLanguage: string = cookie?.get && cookie.get('X-Jdcloud-Language');

const server: AxiosInstance = axios.create({
    baseURL: '/console-web',
    timeout: 10 * 1000,
    headers: {
        'Content-Type': 'application/json; charset=utf-8',
        'X-Jdcloud-Language': vueCookiesLanguage || 'zh_CN'
    }
});

// 获取取消请求的cancel方法
const Cancel: CancelTokenStatic = axios.CancelToken;
// 接收重复请求的数组
let repeatRequestArr: RepeatRequestArrType[]  = [];

/**
 * 删除重复请求
 * @param {Object} config 请求头信息
*/
function removeRequest(config: any): void {
    // 数组遍历 item 当前请求的接口地址+请求类型  index数组的下标 arr要处理的重复请求接口数组
    repeatRequestArr.map((item, index, arr) => {
        // methodUrl 是当前请求的地址+方式，如果参数也相同
        if (item.methodUrl === config.url + '&' + config.method) {
            if (judgeParamsEquality(config?.data || config?.params, item.params)) {
                // 走到这里面说明多次请求了
                // f()是内置的取消请求的方法，此处执行
                item.f();
                // 删除当前请求的这一项
                arr.splice(index, 1);
            }
        }
        return config;
    })
};

/**
 * 判断请求参数是否相等
*/
function judgeParamsEquality(param1: any, param2: any) {
    for (const key in param1) {
        if (param1[key] !== param2[key]) {
            return false;
        }
    }
    return true;
};

/**
 * 请求拦截器
*/
server?.interceptors?.request?.use(
    (config: any): AxiosRequestHeaders | object  => {
        if(!config?.isrepeat) {
            // 删除重复请求(isrepreat:false为不可重复，true为可以重复)
            removeRequest(config)
        }
        // 这个c就是取消请求的方法
        config.cancelToken = new Cancel((c: Function) => {
            // 添加到数组里面，methodUrl-url请求地址+method请求方式
            repeatRequestArr.push({methodUrl: config.url + '&' + config.method, params: config?.data || config?.params, f: c})
        });
        return config;
    },
    (error: AxiosError): Promise<never> | Function  => {
        return Promise.reject(error);
    }
);

/**
 * 响应拦截器
*/
server?.interceptors?.response?.use(
    (response: AxiosResponse): Promise<any> | void => {
        // 每成功一次，数组边为空，否则其他接口重复请求会数据一直叠加
        repeatRequestArr = [];
        const {data, headers} : {data: any, headers: AxiosResponseHeaders} = response;
        // 登录用户名称
        const userId: string = headers['x-jdcloud-userid'];
        // 接口成功并且业务不需要自己处理返回
        return Promise.resolve({data, userId});
        
    },
    (err: AxiosError<any> | any): Promise<any> | void => {
        const {response} = err;
        const errInfo = JSON.parse(JSON.stringify(err));
        const {config, message} : {config: AxiosRequestConfig | any, message: string} = errInfo;
        const errorMsg: string = response?.data?.error?.message;
        // 402登录过期-未登录
        if (response?.status === 402) {
            // 跳转地址-用户最后一次访问的地址
            const pathUrl: string = window.btoa(router.currentRoute.value.fullPath);
            // 设置地址-存储到coolie
            cookie?.set && cookie.set('path_url', pathUrl);
            window.open('/Login/login', '_self');
        }
        // isBusinessProcessing 为 true，业务catch不用再做error提示
        else if (response?.config?.isBusinessProcessing || config?.isBusinessProcessing) {
            ElMessage.error(errorMsg || message || '请求失败');
        }
        else {
            // isBusinessProcessing 为 false，需要业务自己处理catch里内容
            return Promise.reject({message: errorMsg || message, status: response?.status});
        }
    }
);

export default server;
