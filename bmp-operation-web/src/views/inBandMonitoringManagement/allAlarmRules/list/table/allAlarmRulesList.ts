/**
 * @file
 * @author
*/
import store from 'store/index.ts';
import {CurrencyType} from '@utils/publicType';
import {AxiosError} from 'axios';
import {msgTip} from 'utils/index.ts';
import {describeRulesAPI} from 'api/inBandMonitoring/allAlarmRules/request.ts'
import {CurrentInstance, RouterOperate, paginationOperate} from 'utils/publicClass.ts';
import AllAlarmRulesStatic from 'staticData/inBandMonitoring/allAlarmRules/index.ts'

class AllAlarmRulesListOpt {
    errorCode: Ref<number> = ref<number>(0);
    routerOperate = new RouterOperate();
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
    filter;

    constructor(params: any, filter: any) {

        this.filter = filter;
        this.params = params;
        this.getAllAlarmRulesList();
        // this.owerAuthorization();
        // onMounted(() => {
        //     nextTick(() => {
        // this.getAllAlarmRulesList();
        //     })
        // })
    }

    jumpDetail = (item: {ruleId: string}) => {
        const path = this.proxy!.$defInfo.routerPath('alarmRulesDetail');
        this.routerOperate.router.push(`${path}?ruleId=${item.ruleId}`)
    }

    getAllAlarmRulesList = async (otherParams?: any) => {
        this.isLoading.value = true;
        const params = {
            pageNumber: paginationOperate.pageNumber.value,
            pageSize: paginationOperate.pageSize.value,
            ...this.filter.reactiveArr.filterParams,
            ...otherParams
        }
        this.filterEmptyInfo.deleteEmtpyData(params);
        try {
            const res = await describeRulesAPI({...params});
            if (res?.data?.result?.rules?.length) {
                const {totalCount, rules} = res.data.result;
                rules.forEach((item: CurrencyType, index: number) => {
                    AllAlarmRulesStatic.allAlarmRulesListTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    });
                })
                this.reactiveArr.tableData = rules.map((item: any) => {
                    return {
                        ...item,
                        // showTooltip: false
                    }
                });
                
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
            this.mitt.instanceMitt?.proxy?.$Bus?.emit('all-alarm-rule-list', this.reactiveArr.tableData);
        }
    }

    setStatusColor = (status: number) => {
        const obj = new Map([
            [3, '#ff4d4f'],
            [2, '#999'],
            [1, '#43b561']
        ]);
        return obj.get(status);
    }

    refresh = () => {
        if (!this.reactiveArr.tableData.length) {
            this.params.clearClick();
            return;
        }
        this.searchTip.value ? this.search() : this.getAllAlarmRulesList();
    };

    search = () => {
        const params = {};
        for (let index in this.params.ruleForm) {
            if (!index.includes('alarm')) {
                // @ts-ignore
                params[index] = this.params.ruleForm[index];
            }
        }
        this.getAllAlarmRulesList(params);
    }
}

export default AllAlarmRulesListOpt;
