<template>
    <!-- 机型列表 -->
    <div class="operate-management">
        <!-- 标题头部信息 -->
        <header-info
            :type="'modelList'"
            @header-ref="tableScroll.getHeaderRef"
        />
        <!-- 主体内容 -->
        <main class="operate-management-content">
            <!-- 操作 -->
            <div class="search-tip-operate2">
                <!-- 操作-刷新&设置&导出 -->
                <refresh-custom-export
                    :type="'modelList.header.operate.addModel'"
                    :status="modelListOperate.searchTip.value"
                    :search-operate="search"
                    :place-holder="[
                        $t('modelList.header.placeholder.modelName'),
                        $t('modelList.header.placeholder.modelSpecifications')
                    ]"
                    @operate-ref="tableScroll.getOperateRef"
                    @refresh="refresh.refresh"
                    @custom="all.customOperate?.customClickOperate"
                    @export="exportDataOperate.diaLogClick"
                    @btn-ref="tableScroll.getBtnRef"
                    @btn-operate="addModelOpt.addModelClick"
                />
                
            </div>
            <!-- 表格内容 -->
            <article class="operate-management-count model-list">
                <no-data-tip
                    :status="modelListOperate.searchTip.value"
                    :total="paginationOperate.total.value"
                    @back-click="reset.reset"
                />
               <pluginComp.TableList
                    :table-max-height="tableScroll.tableMaxHeight.value"
                    :model-list="modelListOperate"
                    :filter="filter"
                    :search="search"
                    :delete-model-operate="deleteOperate"
                    @add-model="addModelOpt.addModelClick"
                    @edit-model="editModelOpt.editModelClick"
                    @empty-click="addModelOpt.addModelClick"
                    @add-model-template="addModelTemplateOpt.addModelTemplateClick"
                />
            </article>
        </main>
        <!-- 分页 -->
        <model-list-pagination :pagination-opt="modelListOperate"/>
        <!-- 自定义列表组件 -->
        <pluginComp.CustomList/>
        <!-- 删除机型组件 -->
        <pluginComp.DeleteModelOperate :delete-model-operate="deleteOperate"/>
        <!-- 导出数据组件 -->
        <pluginComp.ExportData
            :export-filter-data="exportFilterData"
            :export-data-operate="exportDataOperate"
        />
    </div>
</template>

<script setup lang="ts">
import * as all from './all';
import pluginComp from './module';
import store from 'store/index.ts';
import usePagination from 'hooks/pagination/usePagination.ts';
import ModelListPagination from 'hooks/pagination/paginationSelect.vue';
import {paginationOperate, CurrentInstance, RouterOperate} from 'utils/publicClass.ts';

const proxy = new CurrentInstance().proxy;

nextTick (() => {
    store.idcInfo.idcListNoOpt(proxy.$idcApi.idcListAPI);
})

// 实例化-筛选操作
const filter = new all.FilterOperate(() => modelListOperate.getModelList());

const search = new all.SearchOperate(() => modelListOperate.getModelList());

const modelListOperate = new all.ModelListOperate(search, filter);

const reset = all.resetOperate(filter, search);

const exportDataOperate = new all.ExportDataOperate(all.modelListExportAPI);

const exportFilterData = new all.ExportFilterData(search.reactiveArr, filter.reactiveArr);

usePagination(modelListOperate.getModelList);

// 删除机型操作-实例化
const deleteOperate = all.deleteModelOperate(modelListOperate, reset);

// 刷新操作
const refresh = all.refreshOperate(modelListOperate, reset);

const tableScroll = all.tableScrollOperate(filter.reactiveArr, search.reactiveArr, modelListOperate.searchTip);

const router = new RouterOperate().router;


const addModelOpt = new all.AddModelOperate(proxy, router);

const editModelOpt = new all.EditModelOpt(proxy, router);

const addModelTemplateOpt = new all.ModelTemplateOpt(proxy, router);

</script>

<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';
</style>
