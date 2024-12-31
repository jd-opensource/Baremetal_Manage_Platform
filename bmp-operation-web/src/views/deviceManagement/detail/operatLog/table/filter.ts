/**
 * @file
 * @author
*/
import store from 'store/index.ts';
import {paginationOperate} from 'utils/publicClass.ts';

class FilterOperate {
    // store库存储的oss数据类
    ossStore = store.ossDataInfo()
    tableRef: {
        [x: string]: unknown;
        value?: {
            clearFilter(): unknown;
        }
    } = {value: {clearFilter(): void{}}};
    filterEmptyInfo: { // store
        filterParams(
            arg0: {[x: string]: number | string},
            arg1: any,
            arg2: {filterParams: {result: string; operation: string}},
            arg3: {result: string; operation: string}
    ): Promise<void>;
        deleteEmtpyData(arg0: {result: string; operation: string}): unknown;
    } = store.filterEmpty;
    // 复杂数据类型
    reactiveArr: {
        filterParams: {result: string; operation: string};
    } = reactive<{
        filterParams: {result: string; operation: string};
    }>({
        filterParams: {
            operation: '',
            result: ''
        }
    });
    fn;

    constructor(fn: Function) {
        this.fn = fn;
        this.ossStore?.getAuditLogsTypes();
    }

    operationFilter = () => {
        return true;
    }

    resultFilter = () => {
        return true;
    }

    /**
     * 获取表格ref
     * @param {Object} tableEl 表格信息
    */
    getTableRef = (tableEl: {[x: string]: unknown}): void => {
        this.tableRef = tableEl;
    };

    /**
     * @param {Object} filter 需要过滤的数据
    */
    filterChange = (filter: {[x: string]: number | string}) => {
        const filterParams = {
            operation: 'operation',
            result: 'result',
        };
        this.filterEmptyInfo.filterParams(filter, this.ossStore, this.reactiveArr, filterParams)
        .then((res: any) => {
            this.reactiveArr.filterParams = res.filterParams;
            if (paginationOperate.pageNumber.value > 1) {
                paginationOperate.pageNumber.value = 1;
                return;
            }
            this.fn();
        });
    };
}

export default FilterOperate;
