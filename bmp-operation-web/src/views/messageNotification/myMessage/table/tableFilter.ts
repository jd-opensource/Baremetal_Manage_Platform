import store from 'store/index.ts';
import {paginationOperate} from 'utils/publicClass.ts';
import {OSSType, ReactiveArrType, FilterParamsType} from '../typeManagement';

let that: any = null;

/**
 * 筛选操作
*/
class FilterOperate {
    // store库存储的oss数据类
    ossStore: OSSType = store.ossDataInfo();
    tableRef: {
        [x: string]: unknown;
        value?: {
            toggleRowSelection(arg1: any, arg2: any): any;
            clearFilter(): unknown;
            clearSelection(): any;
        }
    } = {value: {clearFilter(): void{}, toggleRowSelection(): any{}, clearSelection(): any{}}};
    filterEmptyInfo: { // store
        filterParams(
            arg0: {[x: string]: number | string},
            arg1: OSSType,
            arg2: ReactiveArrType,
            arg3: FilterParamsType
    ): Promise<void>;
        deleteEmtpyData(arg0: FilterParamsType): unknown;
    } = store.filterEmpty;
    // 复杂数据类型
    reactiveArr: {
        filterParams: FilterParamsType;
        // selectArr: any;

    } = reactive<{
        filterParams: FilterParamsType;
        // selectArr: any;
    }>({
        filterParams: {
            messageType: '',
            messageSubType: ''
        },
        // selectArr: []
    });
    fn: any;

    // 构造器
    constructor (fn: any) {
        that = this;
        that.fn = fn;
        that.ossStore.getMessagesType()
    };

    /**
     * 获取表格ref
     * @param {Object} tableEl 表格信息
    */
    getTableRef = (tableEl: {[x: string]: unknown}): void => {
        this.tableRef = tableEl;
    };

    messageTypeFilter = (): boolean => {
        return true;
    };

    /**
     * @param {Object} filter 需要过滤的数据
    */
    filterChange = (filter: {[x: string]: number | string}) => {
        this.tableRef.value?.clearSelection()
        const filterParams: FilterParamsType = {
            messageType: 'messageType',
            // messageSubType: 'messageSubType'
            // name: 'name'
        };
        this.filterEmptyInfo.filterParams(filter, this.ossStore, this.reactiveArr, filterParams)
        .then((_: any) => {
            if (paginationOperate.pageNumber.value > 1) {
                paginationOperate.pageNumber.value = 1;
                return;
            }
            this.fn();
        })
    };
};

export default FilterOperate;
