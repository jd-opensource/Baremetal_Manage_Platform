import {AxiosError} from 'axios';
import {paginationOperate, RouterOperate, CurrentInstance} from 'utils/publicClass.ts';
import {msgTip, methodsTotal} from 'utils/index.ts';
import {CurrencyType} from '@utils/publicType';
import store from 'store/index.ts';
import HardwareStaticData  from 'staticData/surveillance/hardware/index.ts';
import DeviceStaticData from 'staticData/device/index.ts';

class HardwareStatusListOperate {
    router = new RouterOperate().router;
    searchTip: Ref<boolean> = ref<boolean>(false);
    isLoading: Ref<boolean> = ref<boolean>(true); // loading态
    tableRef: {
        [x: string]: unknown;
        value: any;
    } = {value: {}};
    errorCode: Ref<number> = ref<number>(0);
    // store
    filterEmptyInfo: {
        deleteEmtpyData(arg0?: {}): Function;
    }  = store.filterEmpty;
    reactiveArr: {
        tableData: CurrencyType[];
    } = reactive<{
        tableData: CurrencyType[];
    }>({
        tableData: [] // 表格-机型列表数据
    });
    params: any;
    instanceProxy = new CurrentInstance().proxy;

    constructor (params: any) {
        this.params = params;
        this.init();
    };

      /**
     * 获取表格ref
     * @param {Object} tableEl 表格信息
    */
    getTableRef = (tableEl: {
        [x: string]: unknown
        value: {
            $refs: any;
        }
    }): void => {
        this.tableRef = tableEl;
    };


    refresh = () => {
        if (!this.reactiveArr.tableData.length) {
            this.params.clearClick();
            return;
        }
        this.searchTip.value ? this.search() : this.init();
    };

    init = () => {
        this.getHardwareList();
    };

    search = () => {
        this.getHardwareList(methodsTotal.toLine(this.params.ruleForm))
    };

    getHardwareList = async (param?: any) => {
        this.isLoading.value = true;
        const params = methodsTotal.toLine({
            pageNum: paginationOperate.pageNumber.value,
            pageSize: paginationOperate.pageSize.value
        });
        const newParamse = {...params, ...param}
        this.filterEmptyInfo.deleteEmtpyData(newParamse);
        this.params.searchParams = newParamse;
        try {
            const listApi = await this.instanceProxy.$hardwareStatusApi.hardwareStatusAPI(newParamse);
            if (listApi?.data?.result?.detail?.length) {
                const {detail, totalCount} = methodsTotal.lineConverting(listApi.data.result);
                detail.forEach((item: CurrencyType, index: number) => {
                    HardwareStaticData.hardwareTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    })
                })
                this.reactiveArr.tableData = detail;
                paginationOperate.total.value = totalCount;
                return;
            }
            throw new Error('');
        }
        catch (e) {
            const err = e as AxiosError;
            this.errorCode.value = err?.code as unknown as number;
            this.params.btnStatus.value = this.errorCode.value === 403;
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
            methodsTotal.initScrollLeft(this.tableRef.value);
            if (['module license disabled'].includes(err?.message)) {
                return;
            }
            err?.message && msgTip.error(err.message);
        }
        finally {
            this.isLoading.value = false;
            this.searchTip.value = Object.keys(newParamse).length > 2;
        }
    };

    setStatus = (status: number) => {
        const key: Map<number, string> = new Map([
            [0, this.instanceProxy.$defInfo.imagePath('greenIcon')],
            [1, this.instanceProxy.$defInfo.imagePath('redIcon')]
        ]);
        return key.get(status);
    };

    setManagementStatus = (status: string) => {
        const newStatus = DeviceStaticData.manageMentStatus.find((item: {filterParams: string; text: string}) => item.filterParams === status);
        return newStatus?.text;
    };


    jumpRouter = (sn: string, path: string) => {
        this.router.push(path);
        sessionStorage.setItem('sn', sn);
    }
};

export default HardwareStatusListOperate;
