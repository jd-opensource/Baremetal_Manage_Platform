<template>
    <div class="operate-management-detail image-detail">
        <detail-header
            :name="$filter.emptyFilter(query.imageName as string)"
            :title="'imageDetail'"
            @header="router.push($defInfo.routerPath('imageList') as unknown as string)"
            @header-ref="tableS.getRefHeader"
        >
            <drop-down-operate
                :reactive-arr="baseInfo.reactiveArr"
                @delete="imageDeleteOpt.deleteImageClick"
            />
        </detail-header>
        <image-detail-tabs-data
            :tabs="activeOperate"
            @header-ref="tableS.getRef"
        />
        <image-detail-base-info
            :show="activeOperate.activeName.value"
            :reactive-arr="baseInfo.reactiveArr"
            :base-info-data="baseInfo.baseInfoData"
            :loading="baseInfo.infoLoading.value"
            @desc="descEditOpt.editDescClick"
        />
        <image-detail-table-data
            :height-value="tableS.tableMaxHeight.value"
            :table-data="tableDetail.reactiveArr.tableData"
            :show="activeOperate.activeName.value"
            :loading="tableDetail.isLoading.value"
            @add="modelAddOpt.addModelClick"
            @remove="removeOpt.removeClick"
            @empty-click="modelAddOpt.addModelClick"
        />
        <image-detail-pagination
            :active-operate="activeOperate"
            :table-detail="tableDetail"
        />
        <image-detail-model-add :query="query" :model-add="modelAddOpt"/>
        <image-detail-remove :query="query" :remove="removeOpt"/>
        <image-detail-image-delete :query="query" :image-delete="imageDeleteOpt"/>
        <image-detail-desc-edit
            :description="baseInfo.reactiveArr.detail.description"
            :query="query"
            :desc-edit="descEditOpt"
        />
    </div>
</template>

<script setup lang="ts">
import pluginComp from './module';
import {useRouter} from 'vue-router';
import methods from './methods';

const [ImageDetailBaseInfo, ImageDetailDescEdit, DropDownOperate, ImageDetailImageDelete, ImageDetailModelAdd, ImageDetailPagination, ImageDetailRemove, ImageDetailTableData, ImageDetailTabsData] = pluginComp;

const [BaseInfo, descEdit, imageDelete, modelAdd, PaginationOperate, remove, AssociatedModel, tableScrollOperate, tabsOperate] = methods;

const router = useRouter();

const query = router.currentRoute.value.query;

const baseInfo = new BaseInfo(query);

const tableDetail = new AssociatedModel(query, baseInfo);

const activeOperate = tabsOperate(tableDetail, query, baseInfo);

const descEditOpt = descEdit(baseInfo);

new PaginationOperate(tableDetail);

const modelAddOpt = modelAdd(tableDetail);

const removeOpt = remove(tableDetail);

const imageDeleteOpt = imageDelete(router, baseInfo);

const tableS = tableScrollOperate(activeOperate.activeName);

</script>

<style lang="scss">
@import 'assets/css/detail.scss';
@import './imageDetail.scss';
</style>