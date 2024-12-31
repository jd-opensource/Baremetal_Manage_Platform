<template>
    <!-- 报警历史验证操作 -->
    <!-- customPagination.total.value -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="['alarm-history-count currency', !pa.total.value ? 'set-height' : '']"
        :is-loading="alarmHistoryOpt.isLoading"
        :header-title="$t('allAlarmRulesDetail.header.operate.alarmHistory')"
        :src="($defInfo.imagePath('defaultInBandMonitoringActive') as unknown as string)"
        @dia-log-close="alarmHistoryOpt.cancelClick"
        @determine-click="alarmHistoryOpt.determineClick"
    >
        <div class="currency-count">
            <ui-config-provider :locale="alarmHistoryOpt.locale">
                <div :class="[alarmHistoryOpt.searchTip.value ? 'reduce-search-opt' : 'search-opt']">
                    <ui-date-picker
                        class="custom-width custom-width4"
                        type="daterange"
                        v-model="alarmHistoryOpt.ruleForm.alarmTime"
                        :disabled-date="alarmHistoryOpt.disabledDate"
                        :default-time="[new Date(0, 0, 0, 0, 0, 0), new Date(0, 0, 0, 23, 59, 59)]"
                        :start-placeholder="$t('historyAlarmInfo.search.placeholder.startDate')"
                        :end-placeholder="$t('historyAlarmInfo.search.placeholder.endDate')"
                        @change="alarmHistoryOpt.dateChange"
                    />
                    <ui-button
                        type="primary"
                        @click="alarmHistoryOpt.search"
                    >
                        {{$t('historyAlarmInfo.search.btn.search')}}
                    </ui-button>
                </div>
                <div class="currency-count-table">
                    <no-data-tip
                        v-if="alarmHistoryOpt.searchTip.value"
                        :status="alarmHistoryOpt.searchTip.value"
                        :has-visibility="true"
                        :total="pa.total.value"
                        @back-click="alarmHistoryOpt.clearClick"
                    />
                        <ui-table
                            border
                            style="width: 100%"
                            max-height="350"
                            :class="[tableClass(alarmHistoryOpt.reactiveArr.tableData, alarmHistoryOpt.isLoading.value)]"
                            :data="alarmHistoryOpt.reactiveArr.tableData"
                        >
                            <ui-table-column
                                align="center"
                                min-width="150"
                                fixed
                                :label="$t('historyAlarmInfo.table.label.alarmTime')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    {{$filter.emptyFilter(getDateMinutes(scope.row.alertTime))}}
                                </template>
                            </ui-table-column>
                            <ui-table-column
                                align="center"
                                min-width="170"
                                :label="$t('historyAlarmInfo.table.label.resourceId')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <span>{{$filter.emptyFilter(scope.row.resourceId)}}</span>
                                </template>
                            </ui-table-column>
                            <ui-table-column
                                align="center"
                                min-width="160"
                                :label="$t('historyAlarmInfo.table.label.triggerDescription')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <ui-tooltip
                                        placement="bottom"
                                        v-model="scope.row[`triggerDescriptionTip${scope.$index}`].showTooltip"
                                        :disabled="!scope.row[`triggerDescriptionTip${scope.$index}`].showTooltip"
                                    >
                                        <template #content>
                                            <div class="set-tooltip-width">
                                                <span>
                                                    {{scope.row.triggerDescription}}
                                                </span>
                                            </div>
                                        </template>
                                        <div
                                            class="more-text-ellipsis"
                                            @mouseenter="hasShowTooltip($event, scope.row[`triggerDescriptionTip${scope.$index}`], scope.row.triggerDescription, 2.1)"
                                        >
                                            <span
                                                style="marginRight: 22px"
                                                :class="setTextClass(false)"
                                            >
                                                {{$filter.emptyFilter(scope.row.triggerDescription)}}
                                            </span>
                                        </div>
                                    </ui-tooltip>
                                </template>
                            </ui-table-column>
                            <!-- 报警值 -->
                            <ui-table-column
                                prop="alertValue"
                                min-width="110"
                                align="center"
                                :label="$t('historyAlarmInfo.table.label.alertValue')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <span v-if="scope.row.alertValue">{{$filter.emptyFilter(scope.row.alertValue)}}</span>
                                    <span v-else>{{$filter.emptyFilter()}}</span>
                                </template>
                            </ui-table-column>
                            <!-- 报警级别 -->
                            <ui-table-column
                                prop="alertLevelDescription"
                                min-width="110"
                                align="center"
                                :label="$t('historyAlarmInfo.table.label.alertLevelDescription')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <span>{{$filter.emptyFilter(scope.row.alertLevelDescription)}}</span>
                                </template>
                            </ui-table-column>
                            <!-- 持续时间 -->
                            <ui-table-column
                                prop="alertPeriod"
                                min-width="110"
                                align="center"
                                fixed="right"
                                :label="$t('historyAlarmInfo.table.label.alertPeriod')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <span v-if="!['', void 0].includes(scope.row.alertPeriod)">
                                        {{$filter.emptyFilter(scope.row.alertPeriod)}}{{$t('historyAlarmInfo.table.count.minute')}}
                                    </span>
                                    <span v-else>{{$filter.emptyFilter()}}</span>
                                </template>
                            </ui-table-column>
                        </ui-table>
                    <!-- 分页组件 -->
                    <rules-alarm-history-pagination :operate="alarmHistoryOpt" :customPagination="pa"/>
                </div>
            </ui-config-provider>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
// import * as all from './all';
import store from 'store/index.ts';
import {getDateMinutes, tableClass, languageSwitch, hasShowTooltip, setTextClass, msgTip} from 'utils/index.ts';
import {describeAlertsAPI} from 'api/inBandMonitoring/historyAlarmInfo/request.ts'
import HistoryAlarmStatic from 'staticData/inBandMonitoring/historyAlarm/index.ts';
import {CurrencyType} from '@utils/publicType';
import paginationOperate from './pagination/pagination';
import {AxiosError} from 'axios';

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
    ruleId: string;
    // imageId: string;
    // architecture: string;
    // imageName: string;
};

interface RuleFormType {
    alarmTime: number | string | null;
    startTime: number | null;
    endTime: number | null;
}

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    diaLog: false,
    // imageId: '',
    // architecture: '',
    // imageName: ''
});
// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const pa = paginationOperate(() => alarmHistoryOpt.init())


class AlarmHistoryOpt {
    searchTip: Ref<boolean> = ref<boolean>(false);
    locale = languageSwitch();
    filterEmptyInfo: {
        deleteEmtpyData(arg0?: {}): Function;
    }  = store.filterEmpty;
    isLoading: Ref<boolean> = ref<boolean>(false);
    // pageNumber: Ref<number> = ref<number>(1);
    // pageSize: Ref<number> = ref<number>(20);
    // total: Ref<number> = ref<number>(0);
    ruleForm: RuleFormType = reactive<RuleFormType>({
        alarmTime: null,
        startTime: null,
        endTime: null
    })
    reactiveArr = reactive({
        tableData: []
    })
    
    constructor() {
        this.init();
    }

    dateChange = (val: string[] | null) => {
        if (!val) {
            this.ruleForm.startTime = null;
            this.ruleForm.endTime = null;
            this.init();
            return;
        };
        this.ruleForm.startTime = new Date(val[0]).getTime() / 1000;
        this.ruleForm.endTime = new Date(val[1]).getTime() / 1000;
    }

    disabledDate = (time: Date) => {
        return time.getTime() > Date.now() || time.getTime() < Date.now() - 3600 * 1000 * 24 * 30;
    };

    search = () => {
        pa.pageNumber.value = 1;
        this.init();
    }

    clearClick = () => {
        const that = this;
        for (let index in this.ruleForm) {
            this.ruleForm[`${index}` as keyof typeof that.ruleForm] = null;
        }
        pa.pageNumber.value = 1;
        that.init();
    }

    init = async () => {
        this.isLoading.value = true;
        const params = {
            pageNumber: pa.pageNumber.value,
            pageSize: pa.pageSize.value,
            ruleId: props.ruleId,
            ...this.ruleForm
        }
        this.filterEmptyInfo.deleteEmtpyData(params)
        try {
            const res = await describeAlertsAPI({
                ...params
            });
            if (res?.data?.result?.instances?.length) {
                const {instances, totalCount} = res.data.result;
                instances.forEach((item: CurrencyType, index: number) => {
                    HistoryAlarmStatic.historyAlarmListTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    });
                })
                this.reactiveArr.tableData = instances;
                pa.total.value = totalCount;
                return;
            }
            throw new Error('');
        }
        catch(e) {
            this.reactiveArr.tableData = [];
            pa.total.value = 0;
            const err = e as AxiosError;
            err?.message && msgTip.error(err.message);
        }
        finally {
            this.isLoading.value = false;
            this.searchTip.value = Object.keys(params).length > 3;
        }
    }

    cancelClick = () => {
        emitValue('dia-log-close', false)
    }

    determineClick = () => {
        emitValue('determine-click');
        this.cancelClick();
    }
}
const alarmHistoryOpt = new AlarmHistoryOpt();

// usePagination(alarmHistoryOpt.init);
</script>
<style lang="scss">
@import './index.scss';
</style>
