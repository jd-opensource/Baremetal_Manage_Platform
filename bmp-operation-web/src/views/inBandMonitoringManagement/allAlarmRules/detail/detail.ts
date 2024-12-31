/**
 * @file
 * @author
*/
import {CurrentInstance, language} from 'utils/publicClass.ts';
import {msgTip} from 'utils/index.ts';
import {describeRulesDetailAPI} from 'api/inBandMonitoring/allAlarmRules/request.ts'
import {AxiosError} from 'axios';
import AllStaticData from 'staticData/staticData.ts';

class AllAlarmRulesDetailOpt {
    isLoading: Ref<boolean> = ref<boolean>(true);
    instanceProxy = new CurrentInstance().proxy;
    reactiveArr: any = reactive({
        detail: {}
    });
    ruleId: Ref<string> = ref<string>('');
    fn;

    constructor(id: string, fn: Function) {
        this.fn = fn;
        this.ruleId.value = id;
        this.init();
    }

    init = async () => {
        this.isLoading.value = true;
        try {
            const res = await describeRulesDetailAPI({ruleId: this.ruleId.value})
            if (res?.data?.result?.statusName?.length) {
                this.reactiveArr.detail = res.data.result;
                return;
            }
            throw new Error('');
        }
        catch(e) {
            this.reactiveArr.detail = {};
            const err = e as AxiosError;
            const path: string = this.instanceProxy.$defInfo.routerPath('alarmRulesList');
            if (AllStaticData.noFoundData.includes(err?.message)) {
                this.fn(path);
                return;
            }
            err?.message && msgTip.error(err.message);
        }
        finally {
            this.isLoading.value = false;
        }
    }

    setName = () => {
        if (!this.reactiveArr?.detail?.statusName?.length) return '--';
        const {ruleName, statusName, status} = this.reactiveArr.detail;
        return ruleName + `<span style="margin-left: 10px; color: ${this.setStatusColor(status)}">${statusName}</span>`;
    }

    setStatusColor = (status: number) => {
        const obj = new Map([
            [3, '#ff4d4f'],
            [2, '#999'],
            [1, '#43b561']
        ]);
        return obj.get(status);
    }

    setNoticePeriod = () => {
        if (!this.reactiveArr?.detail?.noticeOption?.noticePeriod) return '--';
        const {noticePeriod} = this.reactiveArr.detail.noticeOption;
        const objData = new Map([
            [5, `${5}${language.t('allAlarmRulesDetail.count.minute')}`],
            [10, `${10}${language.t('allAlarmRulesDetail.count.minute')}`],
            [15, `${15}${language.t('allAlarmRulesDetail.count.minute')}`],
            [30, `${30}${language.t('allAlarmRulesDetail.count.minute')}`],
            [60, `${1}${language.t('allAlarmRulesDetail.count.hours')}`],
            [180, `${3}${language.t('allAlarmRulesDetail.count.hours')}`],
            [360, `${6}${language.t('allAlarmRulesDetail.count.hours')}`],
            [720, `${12}${language.t('allAlarmRulesDetail.count.hours')}`],
            [1440, `${24}${language.t('allAlarmRulesDetail.count.hours')}`]
        ]);
        return objData.get(noticePeriod)?? '--';
    }

    // setNoticePeriod = () => {
    //     if (!this.reactiveArr?.detail?.noticeOption?.noticePeriod) return '--';
    //     const {noticePeriod} = this.reactiveArr.detail.noticeOption;
    //     return `${noticePeriod}${language.t('allAlarmRulesDetail.count.minute')}`
    // }

    setNoticeWay = () => {
        if (!this.reactiveArr?.detail?.noticeOption?.noticeWay?.length) return '';
        const {noticeWay} = this.reactiveArr.detail.noticeOption;
        const value1 = language.t('allAlarmRulesDetail.notice.internalMessage');
        const value2 = language.t('allAlarmRulesDetail.notice.email');
        const newData = noticeWay.map((key: number) => Object.is(key, 1) ? value1 : value2);
        return newData.join('、');
    }
}

export default AllAlarmRulesDetailOpt;
