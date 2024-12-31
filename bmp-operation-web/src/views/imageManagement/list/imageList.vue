<template>
    <!-- 镜像列表 -->
    <div class="operate-management">
        <!-- 标题头部信息 -->
        <header-info
            :type="'imageList'"
            @header-ref="tableScroll.getHeaderRef"
        />
        <!-- 主体内容 -->
        <main class="operate-management-content">
            <div class="search-tip-operate2">
                <!-- 操作 -->
                <refresh-custom-export
                    :type="'imageList.header.exportImage'"
                    :status="imageList.searchTip.value"
                    :search-operate="search"
                    :place-holder="[
                        $t('imageList.search.placeholder.imageName'),
                        $t('imageList.search.placeholder.imageId')
                    ]"
                    @operate-ref="tableScroll.getOperateRef"
                    @refresh="refresh.refresh"
                    @custom="customOperate.default?.customClickOperate"
                    @btn-ref="tableScroll.getBtnRef"
                    @export="exportDataOperate.diaLogClick"
                    @btn-operate="importOperate?.importImageClick"
                />
            </div>
            <!-- 表格-镜像列表数据 -->
            <article class="operate-management-count image-list">
                <no-data-tip
                    :status="imageList.searchTip.value"
                    :total="paginationOperate.total.value"
                    @back-click="reset.reset"
                />
                <table-list
                    :filter="filter"
                    :search="search"
                    :image-list="imageList"
                    :table-max-height="tableScroll.tableMaxHeight.value"
                    :import-image="importOperate.importImageClick"
                    :delete-image="deleteOperate.deleteImageClick"
                    @empty-click="importOperate.importImageClick"
                />
            </article>
        </main>
        <!-- 分页组件 -->
        <image-list-pagination :pagination-opt="imageList"/>
        <!-- <pluginComp.Pagination :image-list="imageList"/> -->
        <!-- 自定义列表组件 -->
        <image-custom-list/>
        <!-- 导入镜像组件 -->
        <image-import :import-image-operate="importOperate"/>
        <!-- 删除镜像弹窗组件 -->
        <image-delete :delete-operate="deleteOperate"/>
        <!-- 导出数据组件 -->
        <export-data
            :export-filter-data="exportFilterData"
            :export-data-operate="exportDataOperate"
        />
    </div>
</template>

<script setup lang="ts">
import methods from './methods';
import pluginComp from './module';
import {paginationOperate} from 'utils/publicClass.ts';
import usePagination from 'hooks/pagination/usePagination.ts';
import ImageListPagination from 'hooks/pagination/paginationSelect.vue';

const [ImageCustomList, ImageDelete, ExportData, ImageImport, TableList] = pluginComp;

const [customOperate, deleteImageOperate, exportOpt, importImageOperate, SearchOperate, ImageList, refreshOperate, resetOperate, FilterOperate, tableScrollOperate] = methods;

const search = new SearchOperate.default(() => imageList.getImageList());

const filter = new FilterOperate.default(() => imageList.getImageList());

const imageList = new ImageList.default(search, filter);

// new all.PaginationOperate(imageList)
usePagination(imageList.getImageList);

const reset = resetOperate.default(filter, search);

const refresh = refreshOperate.default(imageList, reset);

const importOperate = importImageOperate.default(filter, reset);

const deleteOperate = deleteImageOperate.default(imageList, filter);

const tableScroll = tableScrollOperate.default(filter.reactiveArr, search.reactiveArr, imageList.searchTip);

const exportDataOperate = new exportOpt.ExportDataOperate(exportOpt.imagesExportAPI);

const exportFilterData = new exportOpt.ExportFilterData(search.reactiveArr, filter.reactiveArr);
</script>
<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';
</style>
