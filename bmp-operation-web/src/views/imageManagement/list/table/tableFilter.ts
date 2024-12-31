import store from 'store/index.ts'
import {OSSType, ReactiveArrType, FilterParamsType} from '../typeManagement';
import {paginationOperate} from 'utils/publicClass.ts';

/**
 * 筛选操作
*/
class FilterOperate {
    // store库存储的oss数据类
    ossStore: OSSType = store.ossDataInfo();
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
            arg3: FilterParamsType
    ): Promise<void>;
        deleteEmtpyData(arg0: FilterParamsType): unknown;
    } = store.filterEmpty;
    // 复杂数据类型
    reactiveArr: {
        filterParams: FilterParamsType;
    } = reactive<{
        filterParams: FilterParamsType;
    }>({
        filterParams: {
            source: '',
            architecture: '',
            osType: ''
        }
    });
    fn: any;

    // 构造器
    constructor (fn: any) {
        this.fn = fn;
        this.ossStore?.getOssData();
    };

    /**
     * 镜像类型筛选
     * @param {number} value 当前选中的value值
     * @param {Object} row 当前点击的这一项
     * @param {string} sourceName sourceName
     * @return {string} row.sourceName 赋值
    */
    imageTypeFilter = (): boolean => {
        // value: number, row: {sourceName: string;}
        return true;
        // return row.sourceName === this.ossStore?.source[value - 1]?.text;
    };

    /**
     * 操作系统平台筛选
     * @param {number} value 当前选中的value值
     * @param {Object} row 当前点击的这一项
     * @param {string} osType osType
     * @return {string} row.osType 赋值
    */
    operateSysFilter = (): boolean => {
        // value: number, row: {osType: string;}
        return true;
        // return row.osType === this.ossStore?.osType[value - 1]?.text;
    };

    /**
     * 体系架构筛选
     * @param {number} value 当前选中的value值
     * @param {Object} row 当前点击的这一项
     * @param {string} architecture architecture
     * @return {string} row.architecture 赋值
    */
    filterArchitecture = (): boolean => {
        // value: number, row: {architecture: string;}
        return true;
        // return row.architecture === this.ossStore?.architecture[value - 1]?.text;
    };

    /**
     * @param {Object} filter 需要过滤的数据
    */
    filterChange = (filter: {[x: string]: number | string}) => {
        const filterParams: FilterParamsType = {
            architecture: 'architecture',
            source: 'source',
            osType: 'osType'
        };
        this.filterEmptyInfo.filterParams(filter, this.ossStore, this.reactiveArr, filterParams)
        .then((res: any) => {
            this.reactiveArr.filterParams = res.filterParams;
            if (paginationOperate.pageNumber.value > 1) {
                paginationOperate.pageNumber.value = 1;
                return;
            }
            this.fn();
            // this.imageList.setFilterData(res, 'filter');
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
