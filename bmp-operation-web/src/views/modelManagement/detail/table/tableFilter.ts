import store from 'store/index.ts';
import {CurrencyType} from '@utils/publicType';
import {paginationOperate} from 'utils/publicClass.ts';

class FilterOperate {
    // store库存储的oss数据类
    ossStore: any = store.ossDataInfo();
    // modelDetailStore();
    // tableRef: any;
    tableRef: {
        [x: string]: unknown;
        value?: {
            clearFilter(): Function;
        };
    } = {};
    filterEmptyInfo: { // store
        filterParams(
            arg0: {[x: string]: number | string},
            arg1: any,
            arg2: any,
            arg3: {source: string; osType: string;}
    ): Promise<void>;
        deleteEmtpyData(arg0: {source: string; osType: string;}): unknown;
    } = store.filterEmpty;
    // 复杂数据类型
    reactiveArr: {
        filterParams: {source: string; osType: string;};
    } = reactive<{
        filterParams: {source: string; osType: string;};
    }>({
        filterParams: {
            source: '',
            osType: ''
        }
    });
    modelDetail: any;

    // 构造器
    constructor (modelDetail: any) {
        this.modelDetail = modelDetail;
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
     * 操作系统平台筛选
    */
    operateSysFilter = (): boolean => {
        // return row.osType = (reactiveArr as any).osType[value - 1].text;
        // return row.osType === this.ossStore?.osType[value - 1]?.text;
        // value: number, row: {osType: string}
        return true;
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
     * @param {Object} filter 需要过滤的数据
    */
    filterChange = (filter: CurrencyType) => {
        const filterParams: any = {
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
            this.modelDetail.setFilterData(res, 'filter')
        });
    };
};

export default FilterOperate;
