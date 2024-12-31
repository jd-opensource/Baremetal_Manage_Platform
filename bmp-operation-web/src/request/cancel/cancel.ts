/**
 * @file
 * @author
*/

import axios, {Canceler, CancelTokenStatic, AxiosRequestConfig} from 'axios';
import {RemoveLongReq} from 'request/RemoveLongReq/remove.ts';
import {CurrencyType, RepeatRequestArrType, CurrencyType3} from '@utils/publicType';

// 接收重复请求的数组
let repeatRequestArr: RepeatRequestArrType[] = [];

//获取取消请求的cancel方法
const Cancel: CancelTokenStatic = axios.CancelToken;

class CancelRuquest {

    /**
     * 删除重复请求
     * @param {Object} config 请求头信息
    */
    removeRequest = (config: AxiosRequestConfig | {islongReq?: boolean; method: string; url: string; params: CurrencyType; data: CurrencyType}) => {
        // 数组遍历 item 当前请求的接口地址+请求类型  index数组的下标 arr要处理的重复请求接口数组
        repeatRequestArr.map((item, index, arr) => {
            // methodUrl 是当前请求的地址+方式，如果参数也相同
            if (item.methodUrl === config.url + '&' + config.method) {
                if (this.judgeParamsEquality(config?.data || config?.params, item.params)) {
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

    cancelReq = (name: string, config: AxiosRequestConfig & {islongReq?: boolean}) => {
        // 这个c就是取消请求的方法
        config.cancelToken = new Cancel((c: Canceler) => {
            const restData: RepeatRequestArrType | never = {methodUrl: config.url + '&' + config.method, params: config?.data || config?.params, f: c};
            // 添加到数组里面，methodUrl-url请求地址+method请求方式
            repeatRequestArr.push(restData);
            new RemoveLongReq(name, config, restData);
        });
    };

    /**
     * 判断请求参数是否相等
     * @param {Object} param1 参数1
     * @param {Object} param2 参数2
     * @return {boolean} true 多次请求 false 未多次请求 
    */
    judgeParamsEquality = (param1: CurrencyType3, param2: CurrencyType3) => {
        for (const key in param1) {
            if (param1[key] !== param2[key]) {
                return false;
            }
        }
        return true;
    };
};

export default CancelRuquest;
