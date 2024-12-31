import store from 'store/index.ts';

class FilterOperate {
    filterEmptyInfo = store.filterEmpty;
    reactiveArr: any = reactive<any>({
        filterParams: {
            source: '',
            osType: ''
        }
    });
    ossStore: {
        getOssData(): void;
        osType: any;
        source: any;
    } = store.ossDataInfo();
    pagination: any;
    imageList: any;

    constructor (pagination: any, imageList: any) {
        this.pagination = pagination;
        this.imageList = imageList;
        this.ossStore?.getOssData();
    };

    /**
     * 筛选事件
    */
    filterOsType = (): boolean => {
        // value: number, row: {osType: {[x: string]: string}}
        // return row.osType === this.ossStore?.osType[value - 1]?.text;
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
     * 筛选搜索
    */
    filterChange = (filter: {[x: string]: string;}): void => {
        const filterParams: {source: string; osType: string;} = {
            source: 'source',
            osType: 'osType'
        };
        this.filterEmptyInfo.filterParams(filter, this.ossStore, this.reactiveArr, filterParams)
        .then((res: any) => {
            this.reactiveArr.filterParams = res.filterParams;
            if (this.pagination.pageNumber.value > 1) {
                this.pagination.pageNumber.value = 1;
                return;
            }
            this.imageList.setFilterData(res);
        });
    };
};

export default FilterOperate;
