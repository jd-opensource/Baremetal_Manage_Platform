<template>
    <ui-config-provider :locale="uiLocale.locale">
        <ui-table
            border
            style="width: 100%;"
            v-loading="operateLog.isLoading.value"
            :max-height="tableMaxHeight || 'auto'"
            :class="tableClass(operateLog.reactiveArr.tableData, operateLog.isLoading.value)"
            :data="operateLog.reactiveArr.tableData"
            @filter-change="filter.filterChange"
            @get-table-ref="filter.getTableRef"
        >
            <!-- LogID -->
            <ui-table-column
                prop="logid"
                min-width="160"
                align="center"
                fixed
                v-if="customOperate?.reactiveArr?.hasCustomInfo['logid']?.selected"
                :label="$t('deviceDetail.operatLog.table.label.logId')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span>{{$filter.emptyFilter(scope.row.logid)}}</span>
                </template>
            </ui-table-column>
            <!-- 操作名称 -->
            <ui-table-column
                prop="operationName"
                min-width="110"
                align="center"
                key="operation"
                column-key="operation"
                filter-placement="bottom-end"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['operationName']?.selected"
                :class-name="filterStyleOperate.filterStatus['operation'] ? 'common-status' : 'def-type-status13'"
                :label="$t('deviceDetail.operatLog.table.label.operateName')"
                :filters="filter.ossStore.operation"
                :filter-method="filter?.operationFilter"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span style="margin-right: 18px;">
                        {{$filter.emptyFilter(scope.row.operationName)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 操作人 -->
            <ui-table-column
                prop="userName"
                min-width="110"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['userName']?.selected"
                :label="$t('deviceDetail.operatLog.table.label.operatePeople')"
            >
            </ui-table-column>
            <ui-table-column
                prop="operateTime"
                min-width="150"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['operateTime']?.selected"
                :label="$t('deviceDetail.operatLog.table.label.operateTime')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span>
                        {{getDateMinutes(scope.row.operateTime)}}
                    </span>
                </template>
            </ui-table-column>
            <ui-table-column
                prop="result"
                min-width="130"
                align="center"
                key="result"
                column-key="result"
                filter-placement="bottom-end"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['result']?.selected"
                :label="$t('deviceDetail.operatLog.table.label.operateFeedback')"
                :class-name="filterStyleOperate.filterStatus['result'] ? 'common-status' : 'def-type-status14'"
                :filters="filter.ossStore.result"
                :filter-method="filter?.resultFilter"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <div :style="scope.row.result === 'doing' ? {marginRight: '30px'} : {}" class="operate-feedback">
                        <img
                            alt=""
                            class="common-img"
                            v-if="![null, ''].includes(scope.row.result)"
                            :src="operateLog.resultSrc(scope.row.result, $defInfo.imagePath('optSuccess'), $defInfo.imagePath('optFail'), $defInfo.imagePath('optDoing'))"
                        />
                        <span>
                            {{$filter.emptyFilter(operateLog.setResult(scope.row.result))}}
                        </span>
                    </div>
                    
                </template>
            </ui-table-column>
            <ui-table-column
                prop="failReason"
                min-width="150"
                align="center"
                fixed="right"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['failReason']?.selected"
                :label="$t('deviceDetail.operatLog.table.label.failCode')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span>
                        {{$filter.emptyFilter(scope.row.failReason)}}
                    </span>
                </template>
            </ui-table-column>
        </ui-table>
    </ui-config-provider>
</template>
<script lang="ts" setup>
import {uiLocale} from 'utils/publicClass.ts';
import {tableClass, getDateMinutes} from 'utils/index.ts';
import customOperate from '../customList/custom';
import FilterStyleOperate from './filterStyle';

interface PropsType {
    operateLog: any;
    filter: any;
    tableMaxHeight: number;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

// const 
const filterStyleOperate = new FilterStyleOperate(props.operateLog, props.filter);
</script>