import {msgTip} from 'utils/index.ts';
import {CurrencyType} from '@utils/publicType';
import {paginationOperate} from 'utils/publicClass.ts';
import {relationImageListAPI} from 'api/model/request.ts';
import store from 'store/index.ts';
import ModelStaticData from 'staticData/model/index.ts';

class TableDetail {
    searchTip: Ref<boolean> = ref<boolean>(false);
    // ui库内置的国际化
    // locale = languageSwitch();
    errorTip: boolean = false;
    filterEmptyInfo = store.filterEmpty;
    // loading态
    isLoading: Ref<boolean> = ref<boolean>(true);
    // architectureName: LocationQueryValue | LocationQueryValue[] = dataInit.query?.architecture;
    // 是否请求镜像列表-默认true
    isRequestImageList: Ref<boolean> = ref<boolean>(false);
    systemPartitionInfo = store.sysPartitionInfo;
    hasDisabled: Ref<boolean> = ref<boolean>(true);
    reactiveArr: any = reactive<any>({
        tableData: [] // 表格关联镜像数据
    });
    filter: any = {};
    dataInit: any;
    baseInfo: any;
    tableS: any;

    constructor (dataInit: any, baseInfo: any) {
        this.dataInit = dataInit;
        this.baseInfo = baseInfo;
    };

    init = () => {
        this.getImageListData();
    };

    setFilterData = (params: any, type: string) => {
        // this.type = type;
        // if (type === 'search') this.search = params;
        if (type === 'filter') this.filter = params;
        this.getImageListData();
    };

    getImageListData = () => {
        this.isLoading.value = true;
        this.filterEmptyInfo.deleteEmtpyData(this.filter.filterParams);
        // const id = deviceTypeId;
        const params: any = {deviceTypeId: this.dataInit.deviceTypeId, ...this.filter.filterParams};
        relationImageListAPI(
            {
                ...params,
                pageNumber: paginationOperate.pageNumber.value,
                pageSize: paginationOperate.pageSize.value,
                architecture: this.dataInit.query?.architecture
            }
        ).then(({data} : {data: {result: {totalCount: number, deviceTypes: CurrencyType[]}}}) => {
            if (data?.result?.deviceTypes?.length) {
                this.isRequestImageList.value = false;
                this.reactiveArr.tableData = data.result.deviceTypes.map((item: CurrencyType, index: number) => {
                    ModelStaticData.modelDetailTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    })
                    const systemPartition = JSON.parse((item?.systemPartition?.length ? item?.systemPartition : '[]'));
                    return {
                        ...item,
                        newSystemPartition: systemPartition
                    }
                });
                paginationOperate.total.value = data.result.totalCount;
                this.errorTip = false;
                return;
            }
            return Promise.reject('');
        }).catch(({message}: {message: string}) => {
            message && msgTip.error(message);
            paginationOperate.total.value = 0;
            this.reactiveArr.tableData = [];
            this.isRequestImageList.value = true;
            this.errorTip = true;
        }).finally(() => {
            const {source, osType} = this.filter.filterParams || {};
            this.searchTip.value = (source?.length || osType?.length) ? true : false;
            this.isLoading.value = false;
            if (this.dataInit.routerType === 'image' && this.baseInfo.isRequestmodelDetail.value) {
                this.baseInfo.getModelDetail();
            }
            this.getImageAll();
        })
    };

    getImageAll = () => {
        relationImageListAPI(
            {
                deviceTypeId: this.dataInit.deviceTypeId,
                pageNumber: 1,
                pageSize: 20,
                architecture: this.dataInit.query?.architecture
            }
        )
        .then(({data} : {data: {result: {totalCount: number, deviceTypes: CurrencyType[]}}}) => {
            if (data?.result?.deviceTypes?.length) {
                this.hasDisabled.value = data.result.totalCount >= 20;
                return;
            }
            return Promise.reject();
        })
        .catch(() => {
            this.hasDisabled.value = false;
        })
    }
};

export default TableDetail;
