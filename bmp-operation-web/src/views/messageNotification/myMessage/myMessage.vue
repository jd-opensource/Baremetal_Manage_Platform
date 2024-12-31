<template>
    <div class="operate-management my-message-count">
        <!-- 头部标题内容 -->
        <header-info
            :type="'myMessage'"
            @header-ref="tableScroll.getHeaderRef"
        >
            {{$t('myMessage.header.title', [myMessageListOpt.total.value])}}
            <span
                class="currency-style no-read"
                @click="search.noReadClick"
            >
                {{myMessageListOpt.noReadTotal.value}}
            </span>
            {{$t('myMessage.header.title2')}}
            <!-- 条未读） -->
        </header-info>
        <!-- 主体内容 -->
        <main class="operate-management-content">
            <!-- 操作 -->
            <div class="search-tip-operate2">
                <div class="radio-active">
                    <ui-radio-group v-model="search.reactiveArr.radioGroup.hasReadName">
                        <ui-radio-button
                            value=""
                            :label="$t('myMessage.radioGroup.all')"
                        />
                        <ui-radio-button
                            value="1"
                            :label="$t('myMessage.radioGroup.read')"
                        />
                        <ui-radio-button
                            value="0"
                            :label="$t('myMessage.radioGroup.noRead')"
                        />
                    </ui-radio-group>
                </div>
                <!-- 操作-刷新&设置&导出 -->
                <refresh-custom-export
                    :has-btn-operate="false"
                    :other-class="'set-flex-end'"
                    :status="myMessageListOpt.searchTip.value"
                    :search-operate="search"
                    :place-holder="[
                        $t('myMessage.search.placeholder')
                    ]"
                    @refresh="myMessageListOpt.refresh"
                    @operate-ref="tableScroll.getOperateRef"
                    @custom="customOperate.customClickOperate"
                    @export="exportDataOperate.diaLogClick"
                />
                
            </div>
            <article class="operate-management-count no-top my-message-list">
                <no-data-tip
                    :status="myMessageListOpt.searchTip.value"
                    :total="paginationOperate.total.value"
                    @back-click="myMessageListOpt.reset"
                />
                <TableData
                    :my-message-list="myMessageListOpt"
                    :top-default-height="tableScroll.headerTitle.value.offsetHeight + tableScroll.operate.value.offsetHeight + 20"
                    :filter-operate="filterOperate"
                    :table-max-height="tableScroll.tableMaxHeight.value"
                />
            </article>
        </main>
        <custom-pagination
            :my-message-list="myMessageListOpt"
            :filter="filterOperate"
            @delete-message="deleteMsg.deleteMsgClick"
            @read-message="myMessageListOpt.readMessageClick"
        >
        </custom-pagination>
        <custom-list></custom-list>
        <message-delete :operate="deleteMsg" :data="myMessageListOpt.reactiveArr.selectArr"></message-delete>
        <!-- <message-notification-list-pagination :pagination-opt="myMessageListOpt"/> -->
        <ExportData
            :export-filter-data="exportFilterData"
            :export-data-operate="exportDataOperate"
        />
    </div>
</template>

<script setup lang="ts">
import SearchOperate from './search/search';
import TableData from './table/tableData.vue';
import MyMessageListOpt from './table/myMessageList';
import FilterOperate from './table/tableFilter';
import CustomPagination from './pagination/pagination.vue';
import tableScrollOperate from './table/tableScroll';
import MessageDelete from './messageDelete/messageDelete.vue';
import {paginationOperate, CurrentInstance} from 'utils/publicClass.ts';
import DeleteMsg from './messageDelete/deleteMsg';
import CustomList from './custom/customList.vue'
import customOperate from './custom/custom'
import ExportData from './export/exportData.vue';
import {ExportFilterData, ExportDataOperate} from './export/export';

const search = new SearchOperate(() => {
    clearTimeout(myMessageListOpt.listTimer.value);
    paginationOperate.pageNumber.value = 1;
    myMessageListOpt.init()
})
const proxy = new CurrentInstance().proxy;

const filterOperate = new FilterOperate(() => myMessageListOpt.init());
const myMessageListOpt = new MyMessageListOpt(filterOperate, search);

const tableScroll = tableScrollOperate(filterOperate.reactiveArr, search.reactiveArr, myMessageListOpt.searchTip, myMessageListOpt.routerOperate.router.currentRoute.value.path);

const deleteMsg = new DeleteMsg(myMessageListOpt);
const exportDataOperate = new ExportDataOperate(proxy.$messageApi.messageExportAPI);

const exportFilterData = new ExportFilterData(search.reactiveArr, filterOperate.reactiveArr);
</script>
<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';
@import './index.scss';
</style>