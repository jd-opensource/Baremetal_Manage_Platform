<template>
    <div class="operate-management">
        <!-- 头部标题内容 -->
        <header-info
            :type="'hardwareStatusList'"
            @header-ref="tableScroll.getHeaderRef"
        />
        <!-- 主体内容 -->
        <main class="operate-management-content">
            <!-- 主体内容-表单操作 -->
            <pluginComp.FormSearch
                :rule-form-operate="ruleFormOperate"
                @search-ref="tableScroll.searchRef"
            />
            <div class="search-tip-operate">
                <no-data-tip
                    :status="hardwareStatusListOpt.searchTip.value"
                    :hasVisibility="true"
                    :total="paginationOperate.total.value"
                    @back-click="ruleFormOperate.clearClick"
                />
                <!-- 操作-刷新&设置&导出 -->
                <refresh-custom-export
                    :default-class="'operate-management-content-operate'"
                    :has-btn-operate="false"
                    :other-class="'set-flex-end'"
                    :disabled="ruleFormOperate.btnStatus.value"
                    @refresh="hardwareStatusListOpt.refresh"
                    @custom="all.customOperate.customClickOperate"
                    @operate-ref="tableScroll.getOperateRef"
                    @export="exportDataOperate.diaLogClick"
                />
            </div>
            <!-- 表格内容 -->
            <pluginComp.TableList
                :hardware="hardwareStatusListOpt"
                :table-scroll-operate="tableScroll"
            />
        </main>
        <pluginComp.Pagination  :hardware="hardwareStatusListOpt"/>
        <pluginComp.CustomList/>
        <pluginComp.ExportData
            :export-filter-data="exportFilterData"
            :export-data-operate="exportDataOperate"
        />
    </div>
</template>
<script setup lang="ts">
import * as all from './all';
import pluginComp from './module';
import {methodsTotal} from 'utils/index.ts';
import {paginationOperate, CurrentInstance} from 'utils/publicClass.ts';

const ruleFormOperate = new all.RuleFormOperate((res: string, searchParams: any) => {
    const requestData = [
        [
            (res: string) => !res.localeCompare('hasSearch'),
            () => {
                const status = methodsTotal.lineConverting(searchParams).idcId?.length;
                status ? hardwareStatusListOpt.search() : hardwareStatusListOpt.init();
            }
        ],
        [
            (res: string) => !res.localeCompare('search'),
            () => hardwareStatusListOpt.search()
        ],
        [
            (res: string) => !res.localeCompare('clear'),
            () => hardwareStatusListOpt.init()
        ]
    ];
    for (const key of requestData) {
        if (key[0](res)) {
            key[1](res);
            break;
        }
    }
});

const proxy = new CurrentInstance().proxy;


const tableScroll = new all.TableScrollOperate();

const hardwareStatusListOpt = new all.HardwareStatusListOperate(ruleFormOperate);

new all.PaginationOperate(hardwareStatusListOpt);

const exportFilterData = new all.ExportFilterData(ruleFormOperate, hardwareStatusListOpt.searchTip);
const exportDataOperate = new all.ExportDataOperate(proxy.$hardwareStatusApi.hardwareStatusExportAPI);

</script>
<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';
</style>
