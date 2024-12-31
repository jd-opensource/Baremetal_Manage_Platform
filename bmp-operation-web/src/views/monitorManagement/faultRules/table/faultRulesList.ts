import {faultRulesAPI} from 'api/surveillance/faultRule/request.ts';
import {paginationOperate} from 'utils/publicClass.ts';
import {msgTip, methodsTotal} from 'utils/index.ts';
import store from 'store/index.ts';
import { CurrencyType, CurrencyType3 } from '@utils/publicType';
import FaultRulesStaticeData from 'staticData/surveillance/faultRules/index.ts';

class FaultRulesListOperate {
    isLoading: Ref<boolean> = ref<boolean>(true); // loading态
    searchTip: Ref<boolean> = ref<boolean>(false);
    // store
    filterEmptyInfo: {
        deleteEmtpyData(arg0?: {}): Function;
    }  = store.filterEmpty;
    reactiveArr: {
        tableData: any[];
        selectArr: CurrencyType[];
    } = reactive<{
        tableData: any[];
        selectArr: CurrencyType[];
    }>({
        tableData: [], // 表格-机型列表数据
        selectArr: []
    });
    errorCode: Ref<number> = ref<number>(0);
    tableRef: {
        [x: string]: unknown;
        value?: {
            toggleRowSelection(arg1: CurrencyType, arg2: undefined): void;
            clearFilter(): unknown;
            clearSelection(): void;
            $refs: any;
        }
    } = {
        value: {
            clearFilter(): void{},
            toggleRowSelection(): void{},
            clearSelection(): void{},
            $refs: {}
        }
    };
    params: any;

    constructor (params: any) {
        this.params = params;
        this.init();
    };

    /**
     * 获取表格ref
     * @param {Object} tableEl 表格信息
    */
    getTableRef = (tableEl: CurrencyType3): void => {
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
        this.getFaultRulesList();
    };

    search = () => {
        this.getFaultRulesList(methodsTotal.toLine(this.params.ruleForm))
    };

    getFaultRulesList = (param?: any): void => {
        this.isLoading.value = true;
        const params = methodsTotal.toLine({
            pageNum: paginationOperate.pageNumber.value,
            pageSize: paginationOperate.pageSize.value
        });
        const newParamse = {...params, ...param}
        this.filterEmptyInfo.deleteEmtpyData(newParamse);
        this.params.searchParams = newParamse;
        faultRulesAPI(
            {
                ...newParamse,
            }
        ).then(({data}: {data: {result: {detail: CurrencyType[]}}}) => {
            if (data?.result?.detail?.length) {
                const {detail, totalCount} = methodsTotal.lineConverting(data.result);
                const newData = detail.map((item: CurrencyType, index: number) => {
                    FaultRulesStaticeData.faultRulesTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    });
                    return {
                        ...item,
                        pushLoading: false,
                        useLoading: false,
                        pushStatus: !!(item.pushStatus),
                        useStatus: !!(item.useStatus)
                    }
                })
                this.reactiveArr.tableData = newData;
                paginationOperate.total.value = totalCount;
                return;
            }
            return Promise.reject('');
        }).catch((err: {code: number; message: string}) => {
            this.errorCode.value = err.code as unknown as number;
            this.params.btnStatus.value = this.errorCode.value === 403;
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
            methodsTotal.initScrollLeft(this.tableRef.value);
            if (['module license disabled'].includes(err.message)) {
                return;
            }
            err.message && msgTip.error(err.message);
        }).finally(() => {
            this.isLoading.value = false;
            this.searchTip.value = Object.keys(newParamse).length > 2;
        });
    };

    filterChange = (val: CurrencyType[]) => {
        this.reactiveArr.selectArr = val;
        this.params.btnDisabled.value = !val.length;
    };

    cellClick = (row: CurrencyType, b: {property: string}) => {
        const type = ['pushStatus', 'useStatus'];
        if (type.indexOf(b.property) > -1) return;
        this.toggleSelection([row]);
    };

    toggleSelection = (rows?: CurrencyType[]) => {
        if (rows) {
            rows.forEach((row: CurrencyType) => {
                this.tableRef.value?.toggleRowSelection(row, undefined);
            });
            return;
        }
        this.tableRef.value?.clearSelection();
    };
};

export default FaultRulesListOperate;
