import {ElTable} from 'element-plus';
import {languageSwitch} from 'utils/index.ts';
import {imageDeviceTypesAPI} from 'api/image/request.ts'; // 添加机型接口
// import publicIndexStore from 'store/index.ts'; // 公共store
import {CurrencyType} from '@utils/publicType';
import ImageStaticData from 'staticData/image/index.ts';

interface ParamsType {
    pageNumber: number;
    pageSize: number;
    imageId: string;
    architecture: string;
    hasCheckbox: string;
};

class ImageDeviceTypes {
    locale = languageSwitch();
    // 表单ref
    tableRef = ref<InstanceType<typeof ElTable>>()
    // table-loading加载态
    tableLoading: Ref<boolean> = ref<boolean>(true);
    // filter-空值参数
    reactiveArr: any = reactive<any>({
        selection: [],
        tableData: [],
        arr: [],
        idList: []
    });
    btnDisabled: Ref<number> = ref<number>(0);
    customPagination: any;
    props: any;

    constructor (customPagination: any, props: any) {
        this.customPagination = customPagination;
        this.props = props;
        this.processingParameters();
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
        });
    };

    processingParameters = () => {
        let params: ParamsType = {
            pageNumber: this.customPagination.pageNumber.value,
            pageSize: this.customPagination.pageSize.value,
            imageId: this.props.imageId,
            architecture: this.props.architecture,
            hasCheckbox: '0'
        };
        this.imageDeviceTypes(params);
    };

    imageDeviceTypes = (params: any) => {
        this.tableLoading.value = true;
        imageDeviceTypesAPI(
            {
                ...params
            }
        )
        .then(({data}: {data: any}) => {
            if (data?.result?.deviceTypes?.length) {
                const newDeviceTypes = data.result.deviceTypes.map((item: {hasCheckbox: boolean; deviceTypeId: string;}, index: number) => {
                    item.hasCheckbox = this.reactiveArr?.arr?.filter((ite: {deviceTypeId: string;}) => ite.deviceTypeId === item.deviceTypeId)[0]??false;
                    ImageStaticData.imageModelAddTipData.forEach((t: any) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    })
                    return item;
                });
                this.reactiveArr.tableData = newDeviceTypes;
                this.customPagination.total.value = data.result.totalCount;
                nextTick(() => {
                    for (const key of this.reactiveArr.tableData) {
                        this.tableRef?.value?.toggleRowSelection(key, key.hasCheckbox);
                    }
                });
                return;
            }
            return Promise.reject('');
        })
        .catch(() => {
            this.reactiveArr.tableData = [];
            this.customPagination.total.value = 0;
        })
        .finally(() => {
            this.tableLoading.value = false;
        });
    };


    // 多选框回显
    getRowKeys = (row: {deviceTypeId: string;}) => {
        if (row.deviceTypeId) return row.deviceTypeId;
    };

    selectAllChange = (rows: CurrencyType[]) => {
        const newRows = rows.filter((item) => !item.hasCheckbox);
        if (newRows.length > 0) { // 全选
            for(let i = 0; i < newRows.length; i++) {
                // 未添加的才插入，否则数据会冗余
                if (!this.reactiveArr.arr.some((s: {deviceTypeId: string;}) => s.deviceTypeId === newRows[i].deviceTypeId)) {
                    this.reactiveArr.arr.push(newRows[i]);
                }
            }
        }
        else { // 取消全选
            for (let i = 0; i< this.reactiveArr.tableData.length; i++) {
                for (let j = 0; j< this.reactiveArr.idList.length; j++){
                    // 移除当前页idList的数据
                    if (this.reactiveArr.idList[j].deviceTypeId === this.reactiveArr.tableData[i].deviceTypeId) {
                        this.reactiveArr.idList.splice(j, 1);
                    }
                }
                for(let k = 0; k < this.reactiveArr.arr.length; k++){
                    // 移除当前页idList的数据
                    if (this.reactiveArr.arr[k].deviceTypeId === this.reactiveArr.tableData[i].deviceTypeId) {
                        this.reactiveArr.arr.splice(k, 1);
                    }
                }
            }
        }
    };

    rowClick = (row: any, _: any, event: any) => {
        this.tableRef.value!.toggleRowSelection(row, event);
        this.selectChange(this.reactiveArr.selection, row);
    };

    /**
     * 单选
    */
    selectChange = (rows: CurrencyType[], row: CurrencyType) => {
        const newRows: boolean | number = rows.filter((item) => !item.hasCheckbox).length &&  rows.indexOf(row) !== -1;
        if (newRows) {
            this.reactiveArr.arr.push(row);
        }
        else {
            for(let i = 0; i < this.reactiveArr.idList.length; i++) {
                if(this.reactiveArr.idList[i].deviceTypeId === row.deviceTypeId){
                    this.reactiveArr.idList.splice(i,1);
                }
            }
            for(let i = 0; i < this.reactiveArr.arr.length; i++){
                if(this.reactiveArr.arr[i].deviceTypeId === row.deviceTypeId){
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
};

export default ImageDeviceTypes;
