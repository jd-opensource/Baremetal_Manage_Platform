import {AxiosError} from 'axios';
import {CurrencyType} from '@utils/publicType';
import {msgTip, methodsTotal, deepCopy} from 'utils/index.ts';
import {paginationOperate, language, CurrentInstance} from 'utils/publicClass.ts';
import store from 'store/index.ts';
import FalutLogStaticData  from 'staticData/surveillance/faultLog/index.ts';

// 坑，引入mock这个文件，文件流导出数据会有问题
// import 'mock/faultLog/index.ts';
// import {mockFaultLogListAPI} from 'mock/faultLog/request.ts';

class FaultLogListOperate {
    // routerQuery = new RouterOperate().router.currentRoute.value;
    env: string = process.env.VUE_APP_ENV;
    searchTip: Ref<boolean> = ref<boolean>(false);
    isLoading: Ref<boolean> = ref<boolean>(true); // loading态
    // store
    filterEmptyInfo: {
        deleteEmtpyData(arg0?: {}): Function;
    }  = store.filterEmpty;
    tableRef: any;
    errorCode: Ref<number> = ref<number>(0);
    reactiveArr: {
        tableData: any[];
        selectArr: any[];
    } = reactive<{
        tableData: any[];
        selectArr: any[];
    }>({
        tableData: [], // 表格-机型列表数据
        selectArr: []
    });
    params: any;
    proxy = new CurrentInstance().proxy;

    constructor (params: any = {}) {
        this.params = params;
        const sn: string = sessionStorage?.getItem('sn')?? '';
        if (sn?.length) {
            this.params.ruleForm.sn = sn;
            this.search();
            return;
        }
        this.init();
    };

    /**
     * 获取表格ref
     * @param {Object} tableEl 表格信息
    */
    getTableRef = (tableEl: any): void => {
        this.tableRef = tableEl;
    };

    queryParamsClick = (query: any) => {
        const status = Object.values(query).some((item) => item !== '');
        status ? this.search() : this.init();
    };

    refresh = () => {
        if (!this.reactiveArr.tableData.length) {
            this.params.clearClick();
            return;
        }
        this.searchTip.value ? this.search() : this.init();
    };

    init = () => {
        this.isLoading.value = true;
        this.getFaultLogList();
    };

    search = () => {
        this.isLoading.value = true;
        this.getFaultLogList(methodsTotal.toLine(this.params.ruleForm))
    };

    getFaultLogList = async (param?: any) => {
        const params = methodsTotal.toLine({
            pageNum: paginationOperate.pageNumber.value,
            pageSize: paginationOperate.pageSize.value
        });
        const newParams = {...params, ...param}
        this.filterEmptyInfo.deleteEmtpyData(newParams);
        this.params.searchParams = newParams;
        try {
            const faultLogApi = await this.proxy.$faultLogApi.faultLogAPI(newParams);
            if (faultLogApi?.data?.result?.detail?.length) {
                const {detail, totalCount} = methodsTotal.lineConverting(faultLogApi.data.result);
                this.reactiveArr.selectArr = deepCopy(detail);
                detail.forEach((item: CurrencyType, index: number) => {
                    FalutLogStaticData.faultLogTipData.forEach((t: string) => {
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
            this.errorCode.value = err.code as unknown as number;
            this.params.btnStatus.value = this.errorCode.value === 403;
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
            methodsTotal.initScrollLeft(this.tableRef.value);
            if (['module license disabled'].includes(err.message)) {
                return;
            }
            err.message && msgTip.error(err.message);
        }
        finally {
            this.isLoading.value = false;
            this.searchTip.value = Object.keys(newParams).length > 2;
        }
    };

    setStatus = (status: number) => {
        const newStatus = new Map([
            [0, language.t('surveillance.status.no')],
            [1, language.t('surveillance.status.yes')]
        ]);
        return newStatus.get(status);
    };

    setColor = (status: number | string) => {
        const val: string = String(status);
        const newColor = new Map([
            ['0', '#ff4d4f'],
            ['1', '#43b561']
        ]);
        return newColor.get(val)
    };

    mockData = () => {
        // mockFaultLogListAPI()
        // .then((res: any) => {
        //     const data =  methodsTotal.lineConverting(res.data);
        //     data.forEach((item: any, index: number) => {
        //         FalutLogStaticData.faultLogTipData.forEach((t: string) => {
        //             Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
        //         })
        //     })
        //     this.reactiveArr.tableData = data;
        //     paginationOperate.total.value = res.totalCount;
        // });
    }
};

export default FaultLogListOperate;
