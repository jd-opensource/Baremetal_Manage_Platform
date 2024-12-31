/**
 * @file
 * @author
*/

import {AxiosError} from 'axios';
import {ParamType, TableType} from '../typeManagement';
import {paginationOperate, CurrentInstance} from 'utils/publicClass.ts';
import {msgTip, methodsTotal} from 'utils/index.ts';
import store from 'store/index.ts';
import ModelStaticData from 'staticData/model/index.ts';

class ModelListOperate {
    searchTip: Ref<boolean> = ref<boolean>(false);
    isLoading: Ref<boolean> = ref<boolean>(true); // loading态
    // store
    filterEmptyInfo: {
        deleteEmtpyData(arg0?: {}): Function;
    }  = store.filterEmpty;
    reactiveArr: {
        tableData: TableType[];
    } = reactive<{
        tableData: TableType[];
    }>({
        tableData: [] // 表格-机型列表数据
    });
    search: any = {};
    filter: any = {};
    proxy = new CurrentInstance().proxy;

    constructor (search: any, filter: any) {
        this.search = search;
        this.filter = filter;
        this.remove();
        this.init();
    };
    
    remove = () => {
        const removeData = ['repeat', 'hasTemplate', 'propsName', 'propsDeviceType'];
        for (const key of removeData) {
            sessionStorage?.removeItem(key)
        }
    };

    init = () => {
        this.getModelList();
    };

    getModelList = async () => {
        this.isLoading.value = true;
        let param: ParamType = {
            pageNumber: paginationOperate.pageNumber.value,
            pageSize: paginationOperate.pageSize.value,
            ...this.filter.reactiveArr.filterParams,
            ...this.search.reactiveArr.searchParams
        };
        this.filterEmptyInfo.deleteEmtpyData(param);
        try {
            const res = await this.proxy.$modelApi.modelListAPI({...param});
            if (res?.data?.result?.deviceTypes?.length) {
                const {deviceTypes, totalCount} : {totalCount: number; deviceTypes: TableType[]} =  res.data.result;
                deviceTypes.forEach((item: TableType, index: number) => {
                    ModelStaticData.modelTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    });
                })
                this.reactiveArr.tableData = deviceTypes;
                paginationOperate.total.value = totalCount;
                return;
            }
            throw new Error('');
        }
        catch (e) {
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
            methodsTotal.initScrollLeft(this.filter.tableRef.value);
            const err = e as AxiosError;
            err?.message && msgTip.error(err.message);
        }
        finally {
            this.isLoading.value = false;
            this.search.hasClear.value = false;
            this.searchTip.value = Object.keys(param).length > 2;
        }
    };
};

export default ModelListOperate;
