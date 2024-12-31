import {defineStore} from 'pinia'; // 定义容器名
import STORE_NAMES from 'store/storeName.ts';

import {customListAPI, alarmCustomListAPI} from 'api/request.ts';
import {
    porjectCustomInfo, // 项目列表自定义初始数据
    instanceCustomInfo, // 实例自定义初始信息
    alarmLogCustomInfo, // 报警日志列表初始信息
} from 'utils/constants.ts';



/**
 * 自定义列表-数据信息-类
*/
type CustomDataInfoType = {
    pageKey: string;
    data: Record<number, CurrencyStatusType>
};

interface IndexStateType {
    customDataInfo: CustomDataInfoType[];    
};
/**
 * 通用状态类
*/
type CurrencyStatusType = {
    required: boolean;
    selected: boolean;
};

type OsStoreType = {
    filterParams: string;
    showInfo: string;
    text: string;
    value: number;
};

type NameType = {
    name: string;
    label: string;
};

type CheckListArrType = NameType[] & Omit<CurrencyStatusType, 'required'>[] & {disabled: boolean}[];

// defineStore定义容器
// 参数1：是对仓库的命名，名称必须具备唯一性；
// 参数2：配置的选项对象，即state、getters、actions
// 其中state的写法必须是函数，为了避免在服务端交叉请求导致的状态数据污染，而且必须是箭头函数，为了更好的TS类型推导
const publicIndexStore = defineStore(STORE_NAMES.PUBLICK_INDEX, {
    state: (): IndexStateType => {
        return {
            customDataInfo: [
                {
                    pageKey: 'projectList',
                    data: porjectCustomInfo
                },
                {
                    pageKey: 'instanceList',
                    data: instanceCustomInfo
                },
                {
                    pageKey: 'deviceAlertLogList',
                    data: alarmLogCustomInfo, 

                }
            ]
        }
    },
    actions: {
        /**
         * 自定义列表
         * @param {string} pageKey 不同列表不同key
         * @param {Object} reactiveArr.hasCustomInfo 对应列表信息，用来接收接口返回的状态
         * @param {Array} checkListArr 根据接口返回的信息做数据处理
        */
        customList(
            pageKey: string,
            reactiveArr: {
                hasCustomInfo: Record<number, CurrencyStatusType>,
                checkListArr: CheckListArrType
            }
        ) {
            customListAPI(
                {
                    pageKey
                }
            ).then(({data} : {data: {result: {customInfo: Record<string, CurrencyStatusType>}}}) => {
                if (data?.result && Object.keys(data?.result?.customInfo)?.length) {
                    reactiveArr.hasCustomInfo = data.result.customInfo;
                    return;
                }
                return Promise.reject();
            })
            .catch(() => {
                for (const key of this.customDataInfo) {
                    if (key.pageKey === pageKey) {
                        reactiveArr.hasCustomInfo = key.data;
                    }
                }
            })
            .finally(() => {
                // 列表展示数据
                const checkListArr: CheckListArrType = reactiveArr.checkListArr;
                // 遍历列表展示数据
                for (const key of checkListArr) {
                    // 遍历自定义信息数据
                    for (const index in reactiveArr.hasCustomInfo) {
                        // 点击的当前项等于对应数据的那一项
                        if (index === key.label) {
                            // 状态赋值
                            key.disabled = reactiveArr.hasCustomInfo[index].required;
                            key.selected = reactiveArr.hasCustomInfo[index].selected;
                        }
                    }
                }
            });
        },
        alarmLogCustomList(
            page_key: string,
            reactiveArr: {
                hasCustomInfo: Record<number, CurrencyStatusType>,
                checkListArr: CheckListArrType
            },
          
        ) { 
            
            alarmCustomListAPI(
                {
                    page_key
                }
            ).then(({data} : {data: {result: {custom_info: Record<string, CurrencyStatusType>}}}) => {
                if (data?.result && Object.keys(data?.result?.custom_info)?.length) {
                    reactiveArr.hasCustomInfo = data.result.custom_info;
                    return;
                }
                return Promise.reject();
            })
            .catch(() => {
                for (const key of this.customDataInfo) {
                    if (key.pageKey === page_key) {
                        reactiveArr.hasCustomInfo = key.data;
                    }
                }
            })
            .finally(() => {
                // 列表展示数据
                const checkListArr: CheckListArrType = reactiveArr.checkListArr;
                // 遍历列表展示数据
                for (const key of checkListArr) {
                    // 遍历自定义信息数据
                    for (const index in reactiveArr.hasCustomInfo) {
                        // 点击的当前项等于对应数据的那一项
                        if (index === key.label) {
                            // 状态赋值
                            key.disabled = reactiveArr.hasCustomInfo[index].required;
                            key.selected = reactiveArr.hasCustomInfo[index].selected;
                        }
                    }
                }
            });
        },
        /**
         * 过滤参数-filter-筛选
         * @param {Object} filter 筛选的参数
         * @param {Object} ossStore 筛选的值
         * @param {Object} reactiveArr.filterParams 需要筛选的参数
         * @param {Object} filterParams 当前筛选的参数
         * @return {Function} reoslve() 成功的返回
        */
         filterParams(
            filter: {[x: string]: number[]},
            ossStore: Record<string, {[filterParams: string]: OsStoreType}>,
            reactiveArr: {filterParams: {[x: string]: string}},
            filterParams: {[x: string]: string},
            currentStore: any,
        ) {
            if (!Object.keys(filter).length) {
                return Promise.reject();
            }
            else {
                const tag: string[] = [];
                for (const key in filter) {
                    if (key === filterParams[key]) {    
                        for (const item of filter[key]) {
                            if(ossStore[filterParams[key]]) {
                                tag.push(ossStore[filterParams[key]][item - 1].filterParams);
                            } else {
                                tag.push(currentStore[item - 1].filterParams);
                            }
                            
                        }
                        // filter[key]?.forEach((item: number) => {
                        //     tag.push(ossStore[filterParams[key]][item - 1].filterParams);
                        // });
                        // 找到匹配的对应参数，如果多选，数组转字符串，用逗号分隔
                        reactiveArr.filterParams[filterParams[key]] = tag.join(',');
                    }
                }
                this.deleteEmtpyData(reactiveArr.filterParams);
                return Promise.resolve();
            }
        },

        /**
         * 删除空数据
         * @param {Object} params 当前操作的数据
        */
        deleteEmtpyData(params: {[x: string]: string}): void {
            Object.keys(params).forEach((item: string) => {
                if (!params[item]) {
                    delete params[item];
                }
            });
        }
    }
});

export default publicIndexStore;