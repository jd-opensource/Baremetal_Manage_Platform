import {defineStore} from 'pinia'; // 定义容器名
import {OsStoreType} from '../typeManagement';
import {CurrencyType} from '@utils/publicType';
import storeName from 'store/storeName.ts';

// defineStore定义容器
// 参数1：是对仓库的命名，名称必须具备唯一性；
// 参数2：配置的选项对象，即state、getters、actions
// 其中state的写法必须是函数，为了避免在服务端交叉请求导致的状态数据污染，而且必须是箭头函数，为了更好的TS类型推导
const filterEmptyStore = defineStore(storeName.filterEmpty, {
    state: () => {
        return {}
    },
    actions: {
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
            reactiveArr: {filterParams: CurrencyType},
            filterParams: CurrencyType
        ) {
            if (!Object.keys(filter).length) return Promise.reject();
            const tag: string[] = [];
            for (const key in filter) {
                if (key === filterParams[key]) {
                    for (const item of filter[key]) {
                        tag.push(ossStore[filterParams[key]][item - 1].filterParams);
                    }
                    // 找到匹配的对应参数，如果多选，数组转字符串，用逗号分隔
                    reactiveArr.filterParams[filterParams[key]] = tag.join(',');
                }
            }
            this.deleteEmtpyData(reactiveArr.filterParams);
            return Promise.resolve(reactiveArr);
        },

        /**
         * 删除空数据
         * @param {Object} params 当前操作的数据
        */
        deleteEmtpyData(params: CurrencyType = {}): void {
            Object.keys(params).forEach((item: string) => {
                if (!params[item]) {
                    Reflect.deleteProperty(params, item);
                    return;
                }
                return params;
            });
        }
    }
});

export default filterEmptyStore;
