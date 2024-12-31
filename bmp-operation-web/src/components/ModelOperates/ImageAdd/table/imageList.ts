import {languageSwitch} from 'utils/index.ts';
import {ParamsType} from '../typeManagement'; // ts类
// import publicIndexStore from 'store/index.ts'; // 公共store
import store from 'store/index.ts';
import {CurrencyType} from '@utils/publicType';
import ModelStaticData from 'staticData/model/index.ts';

class ImageList {
    // ui库-语言
    locale = languageSwitch();
    imagesInfo = store.imagesInfo;
    filterEmptyInfo = store.filterEmpty;
    systemPartitionInfo = store.sysPartitionInfo;
    searchTip: Ref<boolean> = ref<boolean>(false);
    // table-loading加载态
    tableLoading: Ref<boolean> = ref<boolean>(true);
    reactiveArr: any = reactive<any>({
        selection: [],
        tableData: [],
        idList: [],
        arr: []
    });
    filter: any = {};
    // 表单ref
    tableRef: Ref<any> = ref<any>();
    btnDisabled: Ref<number> = ref<number>(0);
    customPagination: any;
    props: any;

    constructor (customPagination: any, props: any) {
        this.customPagination = customPagination;
        this.props = props;
        this.init();
        watch(() => this.reactiveArr.arr, (newValue) => {
            this.btnDisabled.value = newValue.filter((item: {hasCheckbox: boolean;}) => !item.hasCheckbox).length
        }, {deep: true})
    };

    /**
     * 获取表格ref
     * @param {Object} tableEl 表格信息
    */
    getTableRef = async (tableEl: any): Promise<void> => {
        await nextTick(() => {
            this.tableRef = tableEl;
        })
    };

    init = (): void => {
        this.tableLoading.value = true;
        this.processingParameters(this.filter);
    };

    setFilterData = (filter: any) => {
        this.filter = filter;
        this.init();
    };

    processingParameters = (filter: any) => {
        const params: ParamsType = {
            pageNumber: this.customPagination.pageNumber.value,
            pageSize: this.customPagination.pageSize.value,
            deviceTypeId: this.props.deviceTypeId,
            architecture: this.props.architecture,
            ...filter.filterParams
        };
        this.filterEmptyInfo.deleteEmtpyData(params);
        this.getImageList(params);
    };

    // checkAll = (event: Event) => {
    //     const filterData = this.reactiveArr.tableData.filter((item: {hasCheckbox: boolean}) => !item.hasCheckbox);
    //     filterData.map(item => {
    //         // toggleRowSelection接收两个参数，第一个是被勾选的数据，第二个是选中状态（注：需要注册 ref 来引用）
    //         if (event) {
    //             for(let i = 0; i < filterData.length; i++) {
    //                 // 未添加的才插入，否则数据会冗余
    //                 if (!this.reactiveArr.arr.some((s: {imageId: string;}) => s.imageId === filterData[i].imageId)) {
    //                     this.reactiveArr.arr.push(filterData[i]);
    //                 }
    //             }
    //         }
    //         else {
    //             for (let i = 0; i< this.reactiveArr.tableData.length; i++) {
    //                 for (let j = 0; j< this.reactiveArr.idList.length; j++){
    //                     // 移除当前页idList的数据
    //                     if (this.reactiveArr.idList[j].imageId === this.reactiveArr.tableData[i].imageId) {
    //                         this.reactiveArr.idList.splice(j, 1);
    //                     }
    //                 }
    //                 for(let k = 0; k < this.reactiveArr.arr.length; k++){
    //                     // 移除当前页idList的数据
    //                     if (this.reactiveArr.arr[k].imageId === this.reactiveArr.tableData[i].imageId) {
    //                         this.reactiveArr.arr.splice(k, 1);
    //                     }
    //                 }
    //             }
    //         }
    //         this.tableRef?.value.toggleRowSelection(item, event);
    //     });
    // };
    rowClick = (row: any, _: any, event: any) => {
        this.tableRef?.value.toggleRowSelection(row, event);
        this.selectChange(this.reactiveArr.selection, row);
    };

    getImageList = (params: ParamsType) => {
        this.imagesInfo.getImage(params, this.reactiveArr).then((result: {images: CurrencyType[], totalCount: number;}) => {
            const {images, totalCount} : {images: CurrencyType[], totalCount: number} = result;
            this.reactiveArr.tableData = images.map((item, index: number) => {
                item.hasCheckbox = this.reactiveArr?.arr?.filter((ite: {imageId: string;}) => ite.imageId === item.imageId)[0]??false;
                ModelStaticData.modelDetailTipData.forEach((t: string) => {
                    Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                })
                const systemPartition = JSON.parse((item?.systemPartition?.length ? item?.systemPartition : '[]'));
                return {
                    ...item,
                    newSystemPartition: systemPartition
                }
            });
            this.customPagination.total.value = totalCount;
            nextTick(() => {
                for (const key of this.reactiveArr.tableData) {
                    this.tableRef?.value?.toggleRowSelection(key, key.hasCheckbox);
                }
            });
        })
        .catch(() => {
            this.reactiveArr.tableData = [];
            this.customPagination.total.value = 0;
        })
        .finally(() => {
            const params = this.filter;
            const val = params?.filterParams && Object.values(params?.filterParams).some((item) => item !== '');
            this.searchTip.value = val;
            this.tableLoading.value = false;
        });
    };

    // 多选框回显
    getRowKeys = (row: {imageId: string;}) => {
        if (row.imageId) return row.imageId;
    };

    /**
     * 单选
    */
    selectChange = (rows: CurrencyType[], row: CurrencyType) => {
        const newRows: boolean | number = rows.filter((item: any) => !item.hasCheckbox).length && rows.indexOf(row) !== -1;
        if (newRows) {
            this.reactiveArr.arr.push(row);
        }
        else {
            for(let i = 0; i < this.reactiveArr.idList.length; i++) {
                if(this.reactiveArr.idList[i].imageId === row.imageId){
                    this.reactiveArr.idList.splice(i,1);
                }
            }
            for(let i = 0; i < this.reactiveArr.arr.length; i++){
                if(this.reactiveArr.arr[i].imageId === row.imageId){
                    this.reactiveArr.arr.splice(i,1);
                }
            }
        }
    };

    /**
     * 当选择项发生变化时会触发该事件
    */
    handleSelectionChange = (val: CurrencyType[]) => {
        this.reactiveArr.selection = val.filter((item) => !item.hasCheckbox);
        // this.selectChange(this.reactiveArr.selection, this.test);
        
    };

    selectAllChange = (rows: CurrencyType[]) => {
        const newRows = rows.filter((item) => !item.hasCheckbox);
        if (newRows.length > 0) { // 全选
        for(let i = 0; i < newRows.length; i++) {
            // 未添加的才插入，否则数据会冗余
            if (!this.reactiveArr.arr.some((s: {imageId: string;}) => s.imageId === newRows[i].imageId)) {
                this.reactiveArr.arr.push(newRows[i]);
            }
        }
        }
        else { // 取消全选
            for (let i = 0; i < this.reactiveArr.tableData.length; i++) {
                for (let j = 0; j< this.reactiveArr.idList.length; j++){
                    // 移除当前页idList的数据
                    if (this.reactiveArr.idList[j].imageId === this.reactiveArr.tableData[i].imageId) {
                        this.reactiveArr.idList.splice(j, 1);
                    }
                }
                for(let k = 0; k < this.reactiveArr.arr.length; k++){
                    // 移除当前页idList的数据
                    if (this.reactiveArr.arr[k].imageId === this.reactiveArr.tableData[i].imageId) {
                        this.reactiveArr.arr.splice(k, 1);
                    }
                }
            }
        }
    };
};

export default ImageList;
