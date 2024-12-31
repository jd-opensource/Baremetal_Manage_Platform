import {AxiosError} from 'axios';
import {msgTip} from 'utils/index.ts';
import {CurrencyType1, PaginationType} from '@utils/publicType';
import {paginationOperate, CurrentInstance} from 'utils/publicClass.ts';
import IdcStaticData from 'staticData/idc/index.ts';

/**
 * 机房列表
*/
class IdcList {
    // init: Ref<boolean> = ref<boolean>(true);
    // loading态
    isLoading: Ref<boolean> = ref<boolean>(true);

    // 复杂响应式数据集合
    reactiveArr: {
        filterParams: {name: string};
        tableData: CurrencyType1[];
        } = reactive<{
            filterParams: {name: string};
            tableData: CurrencyType1[]
        }>({
            filterParams: {
                name: ''
            },
            tableData: [] // 表格数据-机房列表
    });
    proxy = new CurrentInstance().proxy;
    templateData = IdcStaticData.templateData;
    
    // 构造器
    constructor() {
        this.getDataList();
    };

    /**
     * 获取列表数据
    */
    getDataList = async () => {
        this.isLoading.value = true;
        try {
            const params: PaginationType = {
                pageNumber: paginationOperate.pageNumber.value,
                pageSize: paginationOperate.pageSize.value
            };
            const idcInfo = await this.proxy.$idcApi.idcListAPI(params);
            if (idcInfo?.data?.result?.idcs?.length) {
                const {idcs, totalCount} : {idcs: CurrencyType1[], totalCount: number} = idcInfo.data.result;
                idcs.forEach((item: CurrencyType1, index: number) => {
                    IdcStaticData.idcTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    });
                });
                this.reactiveArr.tableData = idcs;
                paginationOperate.total.value = totalCount;
                return;
            }
            throw new Error('');
        }
        catch (e) {
            const err = e as AxiosError;
            paginationOperate.total.value = 0;
            this.reactiveArr.tableData = [];
            err.message && msgTip.error(err.message);
        }
        finally {
            this.isLoading.value = false;
            // this.init.value = false;
        }
    };
};

export default IdcList;
