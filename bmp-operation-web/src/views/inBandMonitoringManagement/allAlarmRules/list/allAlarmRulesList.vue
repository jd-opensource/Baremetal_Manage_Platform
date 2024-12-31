<template>
    <div class="operate-management">
        <!-- 头部标题内容 -->
        <header-info
            :type="'allAlarmRulesList'"
            @header-ref="tableScroll.getHeaderRef"
        />
         <!-- 主体内容 -->
         <main class="operate-management-content">
            <form-search
                :rule-form-operate="ruleFormOperate"
                @search-ref="tableScroll.searchRef"
            />
            <div class="search-tip-operate">
                <!-- @back-click="ruleFormOperate.clearClick" -->
                <no-data-tip
                    :status="allAlarmRulesList.searchTip.value"
                    :has-visibility="true"
                    :total="paginationOperate.total.value"
                    @back-click="ruleFormOperate.clearClick"
                />
                <!-- 操作-刷新&设置&导出 -->
                <refresh-custom-export
                    :default-class="'operate-management-content-operate'"
                    :has-btn-operate="false"
                    @refresh="allAlarmRulesList.refresh"
                    @operate-ref="tableScroll.getOperateRef"
                    @custom="customOperate.customClickOperate"
                    @export="exportDataOperate.diaLogClick"
                />
            </div>
            <!-- 表格内容 -->
            <table-list
                :all-alarm-rules-list="allAlarmRulesList"
                :filter="filter"
                :table-max-height="tableScroll.tableMaxHeight.value"
                @alarm-history="alaramHistoryOpt.alarmHistoryClick"
            />
            <!-- 分页组件 -->
            <all-alarm-rule-list-pagination :pagination-opt="allAlarmRulesList"/>
        </main>
        <history-alarm :alaram-history-opt="alaramHistoryOpt"/>
        <rules-custom-list/>
        <export-data
            :export-filter-data="exportFilterData"
            :export-data-operate="exportDataOperate"
        />
    </div>
</template>
<script lang="ts" setup>
import {paginationOperate} from 'utils/publicClass.ts';
import AllAlarmRulesListOpt from './table/allAlarmRulesList';
import RuleFormOperate from './formSearch/rulesForm';
import formSearch from './formSearch/formSearch.vue';
import TableList from './table/tableList.vue';
import usePagination from 'hooks/pagination/usePagination.ts';
import AllAlarmRuleListPagination from 'hooks/pagination/paginationSelect.vue';
import TableScrollOperate from './table/tableScroll';
import customOperate from './custom/custom';
import ExportData from './export/exportData.vue';
import RulesCustomList from './custom/customList.vue'
import FilterOperate from './table/tableFilter';
import AlarmHistoryOpt from './historyAlarm/historyAlarm';
import HistoryAlarm from './historyAlarm/historyAlarm.vue';
import {ExportDataOperate, ExportFilterData, describeRulesExportAPI} from './export/export';
const filter = new FilterOperate((res: string) => {
    if (Object.is(res, 'search')) {
        allAlarmRulesList.search();
        return;
    }
    allAlarmRulesList.getAllAlarmRulesList()
});

const ruleFormOperate = new RuleFormOperate(filter, (res: string) => {
    const requestData = [
        [
            (res: string) => Object.is('hasSearch', res),
            () => allAlarmRulesList.searchTip.value && allAlarmRulesList.search()
        ],
        [
            (res: string) => Object.is('search', res),
            () => allAlarmRulesList.search()
        ],
        [
            (res: string) => !res.localeCompare('clear'),
            () => {
                if (allAlarmRulesList.searchTip.value) {
                    allAlarmRulesList.getAllAlarmRulesList()
                }
            }
        ]
    ];
    outer: for (const key of requestData) {
        if (key[0](res)) {
            key[1](res);
            break outer;
        }
    }
})

const allAlarmRulesList = new AllAlarmRulesListOpt(ruleFormOperate, filter);

usePagination(allAlarmRulesList.search);
const tableScroll = new TableScrollOperate();

const alaramHistoryOpt = new AlarmHistoryOpt();

const exportDataOperate = new ExportDataOperate(describeRulesExportAPI);
const exportFilterData = new ExportFilterData(ruleFormOperate, allAlarmRulesList.searchTip, filter);

</script>
<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';

</style>