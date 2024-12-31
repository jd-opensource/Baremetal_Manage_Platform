<template>
    <div class="operate-log-count">
        <!-- header信息 操作日志-->
        <div class="no-table-img-url">
            <info-title @info-header-ref="tableScroll.infoHeaderRef"></info-title>
            <main class="operate-management-content">
                <form-rules
                    :rule-form-operate="ruleFormOperate"
                    @search-ref="tableScroll.searchRef"
                >
                </form-rules>
                <div class="search-tip-operate">
                    <no-data-tip
                        :status="operateLog.searchTip.value"
                        :has-visibility="true"
                        :total="paginationOperate.total.value"
                        @back-click="operateLog.reset"
                    />
                    <!-- 操作-刷新&设置&导出 -->
                    <refresh-custom-export
                        :default-class="'operate-management-content-operate'"
                        :has-btn-operate="false"
                        @operate-ref="tableScroll.getOperateRef"
                        @refresh="operateLog.refresh"
                        @custom="customOperate.customClickOperate"
                        @export="exportDataOperate.diaLogClick"
                    />
                </div>
                <article class="operate-management-count operate-log-list">
                    <table-list :operate-log="operateLog" :table-max-height="tableScroll.tableMaxHeight.value" :filter="filter"></table-list>

                </article>

            </main>
            <operate-log-list-pagination :pagination-opt="operateLog"/>

        </div>
        <ListCustom></ListCustom>
        <export-data
            :export-filter-data="exportFilterData"
            :export-data-operate="exportDataOperate"
        />
    </div>

</template>
<script lang="ts" setup>
import {paginationOperate} from 'utils/publicClass.ts';
import TableScrollOperate from './table/tableScroll'
import SetEmpty from './table/setEmpty';
import usePagination from 'hooks/pagination/usePagination.ts';
import OperateLogListPagination from 'hooks/pagination/paginationSelect.vue';
import ExportData from './export/exportData.vue';
import InfoTitle from './infoTitle.vue';
import RuleFormOperate from './formRules/ruleForm';
import FormRules from './formRules/formRules.vue';
import customOperate from './customList/custom';
import ListCustom from './customList/listCustom.vue';
import FilterOperate from './table/filter';
import OperateLog from './table/operateLog';
import TableList from './table/tableList.vue';
import {ExportDataOperate, auditLogsExportAPI, ExportFilterData} from './export/export';

interface PropsType {
    sn: string;
    deviceDetail: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});


const ruleFormOperate = new RuleFormOperate((type: string) => {
    if (type === 'clear') {
        operateLog.searchClear();
    }
    operateLog.getOperateLogList();
});

const filter = new FilterOperate(() => operateLog.getOperateLogList());

const operateLog = new OperateLog(props, ruleFormOperate, filter);

usePagination(operateLog.getOperateLogList);

new SetEmpty(operateLog);

const tableScroll = new TableScrollOperate();

const exportDataOperate = new ExportDataOperate(auditLogsExportAPI);

const exportFilterData = new ExportFilterData(props.sn, ruleFormOperate, filter.reactiveArr);

</script>
<style lang="scss" scoped>
@import './operateLog.scss';

</style>