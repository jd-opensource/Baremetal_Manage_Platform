import store from 'store/index.ts'
import {OsStoreType} from '@utils/publicType';
// import {OSSType, ReactiveArrType, FilterParamsType} from '../typeManagement';
import {paginationOperate} from 'utils/publicClass.ts';

/**
 * 筛选操作
*/
class FilterOperate {
    // store库存储的oss数据类
    ossStore: OsStoreType = store.ossDataInfo();
    tableRef: {
        [x: string]: unknown;
        value: {
            clearFilter(): unknown;
        }
    } = {value: {clearFilter(): void{}}};
    filterEmptyInfo: { // store
        filterParams(
            arg0: {[x: string]: number | string},
            arg1: OsStoreType,
            arg2: any,
            arg3: any
    ): Promise<void>;
        deleteEmtpyData(arg0: any): unknown;
    } = store.filterEmpty;
    // 复杂数据类型
    reactiveArr = reactive({
        filterParams: {
            status: ''
        }
    });
    fn: any;

    // 构造器
    constructor (fn: any) {
        this.fn = fn;
        // this.ossStore?.getOssData();
    };

    statusFilter = (): boolean => {
        return true;
    };


    /**
     * @param {Object} filter 需要过滤的数据
    */
    filterChange = (filter: {[x: string]: number | string}) => {
        const filterParams = {
            status: 'status'
        };
        this.filterEmptyInfo.filterParams(filter, this.ossStore, this.reactiveArr, filterParams)
        .then((res: any) => {
            this.reactiveArr.filterParams = res.filterParams;
            if (paginationOperate.pageNumber.value > 1) {
                paginationOperate.pageNumber.value = 1;
                return;
            }
            this.fn('search');
        });
    };

    /**
     * 获取表格ref
     * @param {Object} tableEl 表格信息
    */
    getTableRef = (tableEl: {
        [x: string]: unknown;
        value: {
            clearFilter(): unknown;
        }
    }): void => {
        this.tableRef = tableEl;
    };
};

export default FilterOperate;
