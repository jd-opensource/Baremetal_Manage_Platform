<template>
    <div class="operate-management">
        <!-- 头部标题内容 -->
        <header-info
            :type="'faultRulesList'"
            @header-ref="tableScroll.getHeaderRef"
        />
        <!-- 主体内容 -->
        <main class="operate-management-content">
            <!-- 主体内容-表单操作 -->
            <pluginComp.FormSearch
                :rule-form-operate="ruleFormOperate"
                @def-set-opt="defaultSetOperate.defaultSetClick"
                @search-ref="tableScroll.searchRef"
            />
            <div class="search-tip-operate">
                <no-data-tip
                    :status="faultRulesListOpt.searchTip.value"
                    :hasVisibility="true"
                    :total="paginationOperate.total.value"
                    @back-click="ruleFormOperate.clearClick"
                />
                <!-- 操作-刷新&设置&导出 -->
                <refresh-custom-export
                    :default-class="'operate-management-content-operate'"
                    :has-btn-operate="false"
                    :disabled="ruleFormOperate.btnStatus.value"
                    @operate-ref="tableScroll.getOperateRef"
                    @refresh="faultRulesListOpt.refresh"
                    @custom="all.customOperate.customClickOperate"
                    @export="exportDataOperate.diaLogClick"
                />
            </div>
            <!-- 表格内容 -->
            <pluginComp.TableList
                :fault-rules="faultRulesListOpt"
                :table-scroll-operate="tableScroll"
                @push-click="faultPushOperate.faultPushClick"
                @use-click="enableStatusOperate.enableStatusClick"
            />
        </main>
        <pluginComp.Pagination :fault-rules="faultRulesListOpt"/>
        <pluginComp.CustomList/>
        <pluginComp.ResetDefSet
            :default-set-operate="defaultSetOperate"
            :fault-rules-list-opt="faultRulesListOpt"
        />
        <pluginComp.FaultPushOpt :fault-push-operate="faultPushOperate"/>
        <pluginComp.EnableStatusOpt :enable-status-operate="enableStatusOperate"/>
        <pluginComp.ExportData
            :export-filter-data="exportFilterData"
            :export-data-operate="exportDataOperate"
        />
    </div>
</template>
<script setup lang="ts">
import * as all from './all';
import {paginationOperate} from 'utils/publicClass.ts';
import pluginComp from './module';

const ruleFormOperate = new all.RuleFormOperate((res: string) => {
    const requestData = [
        [
            (res: string) => !res.localeCompare('hasSearch'),
            () => faultRulesListOpt.searchTip.value && faultRulesListOpt.search()
        ],
        [
            (res: string) => !res.localeCompare('search'),
            () => faultRulesListOpt.search()
        ],
        [
            (res: string) => !res.localeCompare('clear'),
            () => faultRulesListOpt.init()
        ]
    ];
    for (const key of requestData) {
        if (key[0](res)) {
            key[1](res);
            break;
        }
    }
});
const tableScroll = new all.TableScrollOperate();

const faultRulesListOpt = new all.FaultRulesListOperate(ruleFormOperate);

new all.PaginationOperate(faultRulesListOpt);

const faultPushOperate = new all.FaultPushOperate(faultRulesListOpt);

const defaultSetOperate = new all.DefaultSetOperate(faultRulesListOpt);

const enableStatusOperate = new all.EnableStatusOperate(faultRulesListOpt);

const exportDataOperate = new all.ExportDataOperate(all.faultRulesExportAPI);
const exportFilterData = new all.ExportFilterData(ruleFormOperate, faultRulesListOpt.searchTip);

</script>
<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';
</style>