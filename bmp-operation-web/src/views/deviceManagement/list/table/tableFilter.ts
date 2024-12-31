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
        filterParams: {},
        // selectArr: []
    });
    fn: any;

    // 构造器
    constructor (fn: any) {
        that = this;
        that.fn = fn;
        this.ossStore?.getModelList();
        this.ossStore?.getOssData();
    };

    /**
     * 获取表格ref
     * @param {Object} tableEl 表格信息
    */
    getTableRef = (tableEl: {[x: string]: unknown}): void => {
        this.tableRef = tableEl;
    };

    /**
     * 机型名称filter
     * @param {number} value 当前选中的value值
     * @param {Object} row 当前选中的这一项
     * @return {boolean} xxx 返回对应信息
    */
    nameFilter = (): boolean => {
        // value: number, row: {deviceTypeName?: string}
        return true;
        // return row.deviceTypeName === this.ossStore?.deviceTypeId[value - 1]?.text;
    };

    /**
     * 机型filter
     * @param {number} value 当前选中的value值
     * @param {Object} row 当前选中的这一项
     * @return {boolean} xxx 返回对应信息
    */
    deviceSeriesFilter = (): boolean => {
        // value: number, row: {deviceSeriesName: string}
        return true;
        // return row.deviceSeriesName === this.ossStore?.deviceSeries[value - 1]?.text;
    };

    statusFilter = () => {
        return true;
    }

    /**
     * 管理状态filter
     * @param {number} value 当前选中的value值
     * @param {Object} row 当前选中的这一项
     * @return {boolean} xxx 返回对应信息
    */
    manageStatusFilter = (): boolean => {
        // value: number, row: {manageStatusName: string}
        return true;
        // return row.manageStatusName === this.textLocale(this.ossStore?.manageStatus[value - 1]?.text);
    };

    textLocale (str: string) {  
        return str.toLowerCase().split(/\s+/).map((item) => {  
            return item.slice(0, 1).toLocaleLowerCase() + item.slice(1);  
        }).join(' ');  
    };

    /**
     * @param {Object} filter 需要过滤的数据
    */
    filterChange = (filter: {[x: string]: number | string}) => {
        this.tableRef.value?.clearSelection()
        const filterParams: FilterParamsType = {
            manageStatus: 'manageStatus',
            deviceSeries: 'deviceSeries',
            deviceTypeId: 'deviceTypeId',
            collectStatus: 'collectStatus'
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
