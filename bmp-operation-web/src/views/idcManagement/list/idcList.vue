<template>
    <!-- 机房列表 -->
    <div class="operate-management">
        <!-- 头部标题内容 -->
        <header-info
            :type="'idcList'"
            @header-ref="tableScroll.getHeaderRef"
        />
        <!-- 主体内容 -->
        <main class="operate-management-content">
            <div class="search-tip-operate2">
                <!-- 操作-刷新&设置&导出 -->
                <refresh-custom-export
                    :has-btn-operate="false"
                    :other-class="'set-flex-end'"
                    @operate-ref="tableScroll.getOperateRef"
                    @refresh="idcList.getDataList"
                    @custom="customOperate.default.customClickOperate"
                    @export="exportDataOperate.diaLogClick"
                />
            </div>
            <!-- 机房列表 -->
            <table-list
                :idc-list="idcList"
                :table-scroll-operate="tableScroll"
                @edit="verificationOperate.operateClick"
            />
        </main>
        <!-- 分页 -->
        <idc-list-pagination :pagination-opt="idcList"/>
        <!-- 编辑机房-安全验证 -->
        <verification :operate="verificationOperate"/>
        <!-- 编辑机房 -->
        <edit-idc-data :edit-operate="editIdc" :verification-operate="verificationOperate"/>
        <!-- 自定义列表 -->
        <custom-list/>
        <!-- 导出 -->
        <export-data :export-data-operate="exportDataOperate"/>
    </div>
</template>

<script setup lang="ts">
import pluginComp from './module';
import {CurrentInstance} from 'utils/publicClass.ts';
import usePagination from 'hooks/pagination/usePagination.ts';
import IdcListPagination from 'hooks/pagination/paginationSelect.vue';
import methods from './methods';

const [customOperate, editIdcOperate, exportInfo, IdcList, tableScrollOperate, securityVerificationOperate] = methods;

const [CustomList, EditIdcData, ExportData, TableList, Verification] = pluginComp.map((item: {default: any}) => item.default);

const idcList = new IdcList.default();

const editIdc = editIdcOperate.default(idcList);

const tableScroll = tableScrollOperate.default();

const verificationOperate = securityVerificationOperate.default(editIdc);

const proxy = new CurrentInstance().proxy;

usePagination(idcList.getDataList);

const exportDataOperate = new exportInfo.ExportDataOperate(proxy.$idcApi.idcListExportAPI);

</script>

<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';

</style>
