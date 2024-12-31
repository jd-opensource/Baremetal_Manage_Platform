/**
 * @file
 * @author
*/
import {paginationOperate, language} from 'utils/publicClass.ts';
import {auditLogsAPI} from 'api/device/request.ts';
import {CurrencyType} from '@utils/publicType';
import {msgTip} from 'utils/index.ts';
import {AxiosError} from 'axios';
import store from 'store/index.ts';

class OperateLog {
    searchTip: Ref<boolean> = ref<boolean>(false);
    isLoading: Ref<boolean> = ref<boolean>(true); // loading态
    reactiveArr: {
        tableData: CurrencyType[]
    } = reactive<{
        tableData: CurrencyType[]
    }>({
        tableData: []
    });
    sn: string;
    ruleFormOperate;
    filter;
    deviceDetail;
    

    constructor(props: {sn: string; deviceDetail: {initData(): void;}}, ...args: any) {
        this.sn = props.sn;
        this.deviceDetail = props.deviceDetail;
        this.ruleFormOperate = args[0];
        this.filter = args[1];
        this.deviceDetail.initData();
        this.getOperateLogList();
    }

    searchClear = () => {
        for (let index in this.ruleFormOperate.ruleForm) {
            // @ts-ignore
            this.ruleFormOperate.ruleForm[index] = '';
        }
    }

    setResult = (str: string | null) => {
        if (str === null) return '';
        const obj = {
            'success': language.t('operate.success'),
            'fail': language.t('operate.fail'),
            'doing': language.t('operate.doing')
        };
        return obj[`${str}` as keyof typeof obj];
    }

    resultSrc = (item: string | null, ...args: string[]) => {
        if (item === null) return '';
        const obj = {
            'success': args[0],
            'fail': args[1],
            'doing': args[2]
        };
        return obj[`${item}` as keyof typeof obj];
    }

    getOperateLogList = async <T extends CurrencyType[], K extends number> () => {
        this.isLoading.value = true;
        const params = {
            sn: this.sn,
            pageNumber: paginationOperate.pageNumber.value,
            pageSize: paginationOperate.pageSize.value,
            ...this.filter.reactiveArr.filterParams,
            ...this.ruleFormOperate.ruleForm
        }
        this.filter.filterEmptyInfo.deleteEmtpyData(params);
        try {
            const res: {data: {result: {messages: T; totalCount: K}}} = await auditLogsAPI(params);
            if (res?.data?.result?.messages?.length) {
                this.reactiveArr.tableData = res.data.result.messages;
                paginationOperate.total.value = res.data.result.totalCount;
                return;
            }
            throw new Error('');
        }
        catch(e) {
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
            const err = e as AxiosError;
            if (!store.deviceDetailErrorTipInfo.errorStatus) {
                err?.message && msgTip.error(err.message);
            }
        }
        finally {
            this.isLoading.value = false;
            this.searchTip.value = Object.keys(params).length > 3;
        }
    }

    refresh = () => {
        if (!paginationOperate.total.value) {
            this.searchClear();
            this.reset();
            return;
        }
        this.reqList();
    }

    reqList = () => {
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.getOperateLogList();
    }

    reset = () => {
        this.searchClear();
        for (let index in this.filter.reactiveArr.filterParams) {
            // @ts-ignore
            this.filter.reactiveArr.filterParams[index] = '';
        }
        this.filter?.tableRef?.value?.clearFilter();
        this.reqList();
    }
}

export default OperateLog;
