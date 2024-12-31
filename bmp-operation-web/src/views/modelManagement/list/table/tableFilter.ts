import store from 'store/index.ts'
import {OSSType, ReactiveArrType} from '../typeManagement';
import {paginationOperate} from 'utils/publicClass.ts';

/**
 * 筛选操作
*/
class FilterOperate {
    // store库存储的oss数据类
    ossStore: OSSType = store.ossDataInfo();
    // modelDetailStore();
    tableRef: {
        [x: string]: unknown;
        value: {
            clearFilter(): unknown;
        }
    } = {value: {clearFilter(): void{}}};
    filterEmptyInfo: { // store
        filterParams(
            arg0: {[x: string]: number | string},
            arg1: OSSType,
            arg2: ReactiveArrType,
            arg3: {deviceSeries?: string;}
    ): Promise<void>;
        deleteEmtpyData(arg0: {deviceSeries?: string;}): unknown;
    } = store.filterEmpty;
    // 复杂数据类型
    reactiveArr: {
        filterParams: {deviceSeries?: string;};
    } = reactive<{
        filterParams: {deviceSeries?: string;};
    }>({
        filterParams: {}
    });
    fn: any;

    // 构造器
    constructor (fn: any) {
        this.fn = fn;
        this.ossStore?.getOssData();
    };

    /**
     * 获取表格ref
     * @param {Object} tableEl 表格信息
    */
    getTableRef = (tableEl: {
        [x: string]: unknown
        value: {
            clearFilter(): unknown;
        }
    }): void => {
        this.tableRef = tableEl;
    };

    /**
     * 机型filter
     * @return {boolean} xxx 返回对应信息
    */
    deviceSeriesFilter = (): boolean => {
        return true;
    };

    /**
     * @param {Object} filter 需要过滤的数据
    */
    filterChange = (filter: {[x: string]: number | string}) => {
        const filterParams: {deviceSeries: string;} = {
            deviceSeries: 'deviceSeries'
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
        // .catch(() => {
        //     this.reactiveArr.filterParams = {};
        //     this.tableRef?.value?.clearFilter();
        //     this.filterEmptyInfo.deleteEmtpyData(this.reactiveArr.filterParams);
        //     this.modelList.setFilterData(null, null);
        // });
    };
};

export default FilterOperate;
