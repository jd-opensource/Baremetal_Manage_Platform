<template>
    <!-- 机型详情 -->
    <div class="operate-management-detail model-detail">
        <!-- 头部信息 -->
        <detail-header
            :name="baseInfo.reactiveArr?.detail?.name"
            :title="'modelDetail'"
            @header="router.push($defInfo.routerPath('modelList'))"
            @header-ref="tableS.getRefHeader"
        >
            <pluginComp.DropDownOperate
                :reactive-arr="baseInfo.reactiveArr"
                @edit="editModelOpt.editModelClick"
                @delete="delModel.deleteModelClick"
                @template-add="addModelTemplateOpt.addModelTemplateClick"
            />
        </detail-header>
        <!-- 标签页切换 -->
        <pluginComp.TabsData :tabs="tabs" @header-ref="tableS.getRef"/>
        <!-- 标签页为基本信息时显示 -->
        <section
            class="operate-management-detail-info model-detail-basic"
            v-if="tabs.activeName.value === 'essentialInformation'"
        >
            <pluginComp.BaseInfo
                :reactive-arr="baseInfo.reactiveArr"
                :base-info="baseInfo"
                :loading="baseInfo.basicInfoLoading.value"
            />
        </section>
        <!-- 标签页为关联镜像时显示 -->
        <main
            class="operate-management-detail-image-info"
            v-else
        >
            <pluginComp.AddImageBtn :table-detail="tableDetail" :add-image="addImage"/>
            <!-- 表格-关联镜像数据 -->
            <div class="operate-management-count model-detail-image">
                <no-data-tip
                    :status="tableDetail.searchTip.value"
                    :total="paginationOperate.total.value"
                    @back-click="reset.reset"
                />
                <pluginComp.TableData
                    :table-detail="tableDetail"
                    :table-s="tableS"
                    :filter-operate="filterOperate"
                    @add-image="addImage.addImageClick"
                    @delete="delImage.deleteImageClick"
                    @empty-click="addImage.addImageClick"
                />
            </div>
        </main>
        <!-- 分页组件 -->
        <pluginComp.Pagination :table-detail="tableDetail" :active-operate="tabs"/>
        <!-- 添加镜像组件 -->
        <pluginComp.AddImageC
            :addImageOperate="addImage"
            :basic-information="baseInfo"
            :architecture-name="dataInit.query.architecture"
        />
        <!-- 删除镜像组件 -->
        <pluginComp.ImageDelete :base-info="baseInfo" :del-image="delImage"/>
        <!-- 删除机型组件 -->
        <pluginComp.ModelDelete :del-model="delModel" :base-info="baseInfo"/>
    </div>
</template>
<script setup lang="ts">
import * as all from './all';
import pluginComp from './module';
import store from 'store/index.ts';
import {locationItem, RouterOperate, paginationOperate, CurrentInstance} from 'utils/publicClass.ts';

// 路由
const router = new RouterOperate().router;

const dataInit = all.initData(router);

const baseInfo = new all.BasicInfo(dataInit, router);

// const editModelOperate = new all.EditModelOperate(baseInfo);

const tableDetail = new all.TableDetail(dataInit, baseInfo);

const filterOperate = new all.FilterOperate(tableDetail);

const delImage = all.deleteImageOperate(tableDetail);

const delModel = all.deleteModelOperate(baseInfo);

const tabs = all.tabsOperate(tableDetail, dataInit, baseInfo);

const params = {
    active: tabs.activeName,
    filter: filterOperate.reactiveArr.filterParams,
    searchTip: tableDetail.searchTip
};

const tableS = all.tableScrollOperate(params);

new all.PaginationOperate(tableDetail, filterOperate);

const reset = all.resetOperate(filterOperate, tableDetail);

const addImage = all.addImageOperate(reset.reset);

const proxy = new CurrentInstance().proxy;

store.idcInfo.idcListNoOpt(proxy.$idcApi.idcListAPI);

const editModelOpt = new all.EditModelOpt(proxy, router);

const addModelTemplateOpt = new all.ModelTemplateOpt(proxy, router);

</script>
<style lang="scss">
@import 'assets/css/detail.scss';
@import './modelDetail.scss';
</style>
