<template>
    <div class="operate-management">
        <!-- 头部标题内容 -->
        <header-info
            :type="'faultLogList'"
            @header-ref="tableScroll.getHeaderRef"
        />
        <!-- 主体内容 -->
        <main class="operate-management-content">
            <pluginComp.FormSearch
                :rule-form-operate="ruleFormOperate"
                @search-ref="tableScroll.searchRef"
            />
            <div class="search-tip-operate">
                <no-data-tip
                    :status="faultLogListOperate.searchTip.value"
                    :has-visibility="true"
                    :total="paginationOperate.total.value"
                    @back-click="ruleFormOperate.clearClick"
                />
                <!-- 操作-刷新&设置&导出 -->
                <refresh-custom-export
                    :default-class="'operate-management-content-operate'"
                    :has-btn-operate="false"
                    :disabled="ruleFormOperate.btnStatus.value"
                    @operate-ref="tableScroll.getOperateRef"
                    @refresh="faultLogListOperate.refresh"
                    @custom="all.customOperate.customClickOperate"
                    @export="exportDataOperate.diaLogClick"
                />
            </div>
            <!-- 表格内容 -->
            <pluginComp.TableList
                :fault-log="faultLogListOperate"
                :table-scroll-operate="tableScroll"
                @no-untreated="faultHandingOpt.faultHandingConfirmClick"
            />
        </main>
        <!-- 分页 -->
        <pluginComp.Pagination :fault-log="faultLogListOperate"/>
        <!-- 自定义列表 -->
        <pluginComp.CustomList/>
        <pluginComp.FaultHanding :fault-handing-opt="faultHandingOpt"/>
        <pluginComp.ExportData
            :export-filter-data="exportFilterData"
            :export-data-operate="exportDataOperate"
        />
    </div>
</template>
<script setup lang="ts">
import * as all from './all';
import pluginComp from './module';
import RegularContent from 'utils/regular.ts';
import {paginationOperate, CurrentInstance} from 'utils/publicClass.ts';

all.unmountOperate();

const ruleFormOperate = new all.RuleFormOperate((res: string, params: any) => {

    const requestData = [
        [
            (res: string) => RegularContent.hasEqualReg('hasSearch').test(res),
            () => faultLogListOperate.searchTip.value && faultLogListOperate.search()
        ],
        [
            (res: string) => RegularContent.hasEqualReg('search').test(res),
            () => faultLogListOperate.search()
        ],
        [
            (res: string) => Object.is('clear', res),
            () => faultLogListOperate.init()
        ],
        [
            (res: string) => RegularContent.hasEqualReg('queryParams').test(res),
            () => faultLogListOperate.queryParamsClick(params)
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
const faultLogListOperate = new all.FaultLogListOperate(ruleFormOperate);
const faultHandingOpt = all.faultHandingConfirmOpt(faultLogListOperate);

new all.PaginationOperate(faultLogListOperate);

const proxy = new CurrentInstance().proxy;

const exportDataOperate = new all.ExportDataOperate(proxy.$faultLogApi.faultLogExportAPI);
const exportFilterData = new all.ExportFilterData(ruleFormOperate, faultLogListOperate.searchTip);

</script>
<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';

</style>