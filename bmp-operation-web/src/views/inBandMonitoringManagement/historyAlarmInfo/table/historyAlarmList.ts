/**
 * @file
 * @author
*/
import store from 'store/index.ts';
import {AxiosError} from 'axios';
import {msgTip} from 'utils/index.ts';
import {describeAlertsAPI} from 'api/inBandMonitoring/historyAlarmInfo/request.ts'
import {CurrentInstance, paginationOperate} from 'utils/publicClass.ts';
import HistoryAlarmStatic from 'staticData/inBandMonitoring/historyAlarm/index.ts';
import {CurrencyType} from '@utils/publicType';

class HistoryAlarmInfoOpt {
    errorCode: Ref<number> = ref<number>(0);
    searchTip: Ref<boolean> = ref<boolean>(false);
    isLoading: Ref<boolean> = ref<boolean>(true); // loading态
    // store
    filterEmptyInfo: {
        deleteEmtpyData(arg0?: {}): Function;
    }  = store.filterEmpty;
    tableRef: any;
    reactiveArr: {
        tableData: any[];
    } = reactive<{
        tableData: any[];
    }>({
        tableData: [], // 表格-机型列表数据
    });
    params;
    proxy = new CurrentInstance().proxy;
    mitt = new CurrentInstance();

    constructor(params: any) {
        this.params = params;
        this.getHistoryAlarmList();
    }

    getHistoryAlarmList = async (otherParams?: any) => {
        this.isLoading.value = true;
        const params = {
            pageNumber: paginationOperate.pageNumber.value,
            pageSize: paginationOperate.pageSize.value,
            ...otherParams
        }
        this.filterEmptyInfo.deleteEmtpyData(params);
        try {
            const res = await describeAlertsAPI({...params});
            if (res?.data?.result?.instances?.length) {
                const {totalCount, instances} = res.data.result;
                instances.forEach((item: CurrencyType, index: number) => {
                    HistoryAlarmStatic.historyAlarmListTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    });
                })
                this.reactiveArr.tableData = instances;
                paginationOperate.total.value = totalCount;
                return;
            }
            throw new Error('');
        }
        catch(e) {
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
            const err = e as AxiosError;
            // this.errorCode.value = err.code as unknown as number;
            // if (['module license disabled'].includes(err.message)) {
            //     return;
            // }
            err?.message && msgTip.error(err.message);
        }
        finally {
            this.isLoading.value = false;
            this.searchTip.value = Object.keys(params).length > 2;
            this.mitt.instanceMitt?.proxy?.$Bus?.emit('history-alarm-list', this.reactiveArr.tableData);
        }
    }

    refresh = () => {
        if (!this.reactiveArr.tableData.length) {
            this.params.clearClick();
            return;
        }
        this.searchTip.value ? this.search() : this.getHistoryAlarmList();
    };

    search = () => {
        const params = {};
        for (let index in this.params.ruleForm) {
            if (!index.includes('alarm')) {
                // @ts-ignore
                params[index] = this.params.ruleForm[index];
            }
        }
        this.getHistoryAlarmList(params);
    }
}

export default HistoryAlarmInfoOpt;
