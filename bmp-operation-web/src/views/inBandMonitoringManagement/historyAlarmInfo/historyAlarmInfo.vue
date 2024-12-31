<template>
    <div class="operate-management history-alarm-count">
        <!-- 头部标题内容 -->
        <header-info
            :type="'historyAlarmInfo'"
            @header-ref="tableScroll.getHeaderRef"
        />
        <!-- 主体内容 -->
        <main class="operate-management-content">
            <form-search
                :rule-form-operate="ruleFormOperate"
                @search-ref="tableScroll.searchRef"
            />
            <div class="search-tip-operate">
                <no-data-tip
                    :status="historyAlarmInfo.searchTip.value"
                    :has-visibility="true"
                    :total="paginationOperate.total.value"
                    @back-click="ruleFormOperate.clearClick"
                />
                <!-- 操作-刷新&设置&导出 -->
                <!--
                    @custom="all.customOperate.customClickOperate"
                     -->
                <refresh-custom-export
                    :default-class="'operate-management-content-operate'"
                    :has-btn-operate="false"
                    :custom-status="false"
                    @operate-ref="tableScroll.getOperateRef"
                    @refresh="historyAlarmInfo.refresh"
                    @export="exportDataOperate.diaLogClick"
                />
            </div>
            <!-- 表格内容 -->
            <table-list
                :history-alarm-info="historyAlarmInfo"
                :table-max-height="tableScroll.tableMaxHeight.value"
            />
            <!-- 分页组件 -->
            <history-alarm-list-pagination :pagination-opt="historyAlarmInfo"/>
        </main>
        <export-data
            :export-filter-data="exportFilterData"
            :export-data-operate="exportDataOperate"
        />
    </div>
</template>
<script lang="ts" setup>
import RuleFormOperate from './formSearch/rulesForm';
import {paginationOperate} from 'utils/publicClass.ts';
import FormSearch from './formSearch/formSearch.vue';
import TableList from './table/tableList.vue';
import HistoryAlarmInfoOpt from './table/historyAlarmList';
import usePagination from 'hooks/pagination/usePagination.ts';
import TableScrollOperate from './table/tableScroll';
import ExportData from './export/exportData.vue';
import HistoryAlarmListPagination from 'hooks/pagination/paginationSelect.vue';
import {ExportDataOperate, ExportFilterData, describeAlertsExportAPI} from './export/export';

const ruleFormOperate = new RuleFormOperate((res: string) => {
    const requestData = [
        [
            (res: string) => Object.is('hasSearch', res),
            () => historyAlarmInfo.searchTip.value && historyAlarmInfo.search()
        ],
        [
            (res: string) => !res.localeCompare('search'),
            () => historyAlarmInfo.search()
        ],
        [
            (res: string) => !res.localeCompare('clear'),
            () => {
                if (historyAlarmInfo.searchTip.value) {
                    historyAlarmInfo.getHistoryAlarmList()
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
});

const historyAlarmInfo = new HistoryAlarmInfoOpt(ruleFormOperate);
usePagination(historyAlarmInfo.search);
const tableScroll = new TableScrollOperate();


const exportDataOperate = new ExportDataOperate(describeAlertsExportAPI);
const exportFilterData = new ExportFilterData(ruleFormOperate, historyAlarmInfo.searchTip);

</script>
<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';

</style>