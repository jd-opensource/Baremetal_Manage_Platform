<template>
    <!-- element-plus内置国际化 -->
    <ui-config-provider :locale="uiLocale.locale">
        <!-- 表格数据-设备列表 -->
        <!-- 注：复选框勾选-prop必传【如果有限制哪一个不能点击】 -->
        <ui-table
            ref="tableRef"
            style="width: 100%"
            v-loading="myMessageList.isLoading.value"
            :class="[
                'currency-count-table',
                tableClass(myMessageList.reactiveArr.tableData, myMessageList.isLoading.value)
            ]"
            :border="true"
            :data="myMessageList.reactiveArr.tableData"
            :max-height="tableMaxHeight || 'auto'"
            :row-key="myMessageList.getRowKeys"
            @row-click="myMessageList.rowClick"
            @filter-change="filterOperate.filterChange"
            @get-table-ref="filterOperate.getTableRef"
            @cell-click="myMessageList.cellClick"
            @selection-change="myMessageList.handleSelectionChange"
        >
            <ui-table-column
                type="selection"
                align="center"
                width="90"
                fixed
                className="set-header-count-left"
                :reserve-selection="true"
            />
            <!-- 消息内容 -->
            <ui-table-column
                className="set-header-count-left"
                prop="content"
                fixed
                v-if="customOperate?.reactiveArr?.hasCustomInfo['detail']?.selected"
                :min-width="setDiffClass('360', '280')"
                :label="$t('myMessage.label.messageContent')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <p
                        :class="[[void 0, 0, ''].includes(scope.row?.hasRead) ? 'message-count-no' : 'message-count']"
                        @click="myMessageList.jump(scope.row)"
                    >
                        <p
                            class="radius-id"
                            v-if="[void 0, 0, ''].includes(scope.row?.hasRead)"
                        >
                        </p>
                        <span>{{myMessageList.setMessageCount(scope.row)}}</span>
                    </p>
                </template>
            </ui-table-column>
            <!-- 接收时间 -->
            <ui-table-column
                align="center"
                min-width="120"
                prop="finishTime"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['finishTime']?.selected"
                :label="$t('myMessage.label.finishTime')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    {{$filter.emptyFilter(getDateMinutes(scope.row.finishTime))}}
                </template>
            </ui-table-column>
            <!-- 消息类型 -->
            <ui-table-column
                align="center"
                prop="messageType"
                column-key="messageType"
                key="messageType"
                filter-placement="bottom-end"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['messageType']?.selected"
                :has-user-template="true"
                :min-width="setDiffClass('150', '100')"
                :label="$t('myMessage.label.messageType')"
                :filter-multiple="false"
                :class-name="filterStyleOperate.filterStatus['messageType'] ? 'common-status' : 'def-type-status'"
                :filters="filterOperate?.ossStore?.messageType"
                :filter-method="filterOperate.messageTypeFilter"
            >
                <template #default="{scope}">
                    <span style="margin-right: 16px;">
                        {{$filter.emptyFilter(scope.row.messageType)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 消息子类型 -->
            <!--  -->
            <ui-table-column
                align="center"
                prop="messageSubType"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['messageSubType']?.selected"
                fixed="right"
                :min-width="setDiffClass('190', '100')"
                :label="$t('myMessage.label.messageSubType')"
                :has-user-template="true"
                :className="myMessageList.setClass(filterOperate.reactiveArr.filterParams.messageSubType)"
            >
                <template #default="{scope}">
                   <span style="margin-right: 20px;">
                        {{$filter.emptyFilter(scope.row.messageSubType)}}
                   </span>
                </template>
            </ui-table-column>
           
        </ui-table>
        <custom-filter :filter-operate="filterOperate" :other-filter-style="filterStyleOperate" :operate="myMessageList" :type-class="'def-cls'" :top-default-height="topDefaultHeight"></custom-filter>
    </ui-config-provider>
   
</template>
<script lang="ts" setup>
import {tableClass, getDateMinutes, setDiffClass} from 'utils/index.ts';
import {uiLocale} from 'utils/publicClass.ts';
import SetEmpty from './setEmpty';
import FilterStyleOperate from './filterStyle';
import customOperate from '../custom/custom';
import MyMessageListOpt from './myMessageList';

interface PropsType {
    myMessageList: MyMessageListOpt;
    tableMaxHeight: number;
    filterOperate: any;
    topDefaultHeight: number;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

new SetEmpty();

const filterStyleOperate = new FilterStyleOperate(props);

</script>
