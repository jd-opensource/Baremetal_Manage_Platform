<template>
    <article class="operate-management-count all-alarm-rule-list">
        <ui-config-provider :locale="uiLocale.locale">
            <ui-table
                ref="tableRef"
                border
                style="width: 100%;"
                v-loading="allAlarmRulesList.isLoading.value"
                :max-height="tableMaxHeight || 'auto'"
                :class="tableClass(allAlarmRulesList.reactiveArr.tableData, allAlarmRulesList.isLoading.value)"
                :data="allAlarmRulesList.reactiveArr.tableData"
                @filter-change="filter.filterChange"
                @get-table-ref="filter.getTableRef"
            >
                <!-- 用户名称 -->
                <ui-table-column
                    prop="resourceName"
                    min-width="110"
                    align="center"
                    fixed
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['userName']?.selected"
                    :label="$t('allAlarmRulesList.table.label.userName')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <span>{{$filter.emptyFilter(scope.row.userName)}}</span>
                    </template>
                </ui-table-column>
                <!-- 规则名称 -->
                <ui-table-column
                    min-width="140"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['ruleName']?.selected"
                    :label="$t('allAlarmRulesList.table.label.ruleName')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                         <div v-if="![void 0, null, ''].includes(scope.row.ruleName)">
                            <ui-tooltip
                                placement="bottom"
                                v-model="scope.row.showTooltip"
                                :disabled="!scope.row.showTooltip"
                            >
                                <template #content>
                                    <div class="set-tooltip-width">
                                        <span>
                                            {{scope.row.ruleName}}
                                        </span>
                                    </div>
                                </template>
                                <div
                                    class="more-text-ellipsis"
                                    @mouseenter="hasShowTooltip($event, scope.row, scope.row.ruleName)"
                                >
                                    <span>
                                        {{$filter.emptyFilter(scope.row.ruleName)}}
                                    </span>
                                </div>
                            </ui-tooltip>
                        </div>
                       <div v-else> -- </div>
                    </template>
                </ui-table-column>
                <!-- 规则ID -->
                <ui-table-column
                    prop="ruleId"
                    min-width="210"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['ruleId']?.selected"
                    :label="$t('allAlarmRulesList.table.label.ruleId')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <div v-if="![void 0, null].includes(scope.row.ruleId)">
                            <ui-tooltip
                                placement="bottom"
                                v-model="scope.row[`ruleIdTip${scope.$index}`].showTooltip"
                                :disabled="!scope.row[`ruleIdTip${scope.$index}`].showTooltip"
                            >
                                <template #content>
                                    <div class="set-tooltip-width">
                                        <span>
                                            {{scope.row.ruleId}}
                                        </span>
                                    </div>
                                </template>
                                <div
                                    class="more-text-ellipsis"
                                    @mouseenter="hasShowTooltip($event, scope.row[`ruleIdTip${scope.$index}`], scope.row.ruleId)"
                                >
                                    <span>
                                        {{scope.row.ruleId}}
                                    </span>
                                </div>
                            </ui-tooltip>
                        </div>
                        <span v-else>
                            {{$filter.emptyFilter(scope.row.ruleId)}}
                        </span>
                    </template>
                </ui-table-column>
                <!-- 资源类型 -->
                <ui-table-column
                    prop="resourceName"
                    min-width="110"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['resourceName']?.selected"
                    :label="$t('allAlarmRulesList.table.label.resourceName')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <span>{{$filter.emptyFilter(scope.row.resourceName)}}</span>
                    </template>
                </ui-table-column>
                <!-- 实例关联数量 -->
                <ui-table-column
                    prop="instanceCount"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['instanceCount']?.selected"
                    :min-width="setDiffClass('150', '100')"
                    :label="$t('allAlarmRulesList.table.label.instanceCount')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <span>{{$filter.emptyFilter(scope.row.instanceCount)}}</span>
                    </template>
                </ui-table-column>
                <!-- 触发条件 -->
                <ui-table-column
                    prop="triggerDescription"
                    min-width="130"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['triggerDescription']?.selected"
                    :label="$t('allAlarmRulesList.table.label.triggerDescription')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <div v-if="scope.row?.triggerDescription?.length > 0">
                            <ui-tooltip
                                placement="bottom"
                                v-model="scope.row[`triggerDescriptionTip${scope.$index}`].showTooltip"
                                :disabled="!scope.row[`triggerDescriptionTip${scope.$index}`].showTooltip"
                            >
                                <template #content>
                                    <div class="set-tooltip-width">
                                        <span>
                                            {{scope.row.triggerDescription.join(',')}}
                                        </span>
                                    </div>
                                </template>
                                <div
                                    class="more-text-ellipsis"
                                    @mouseenter="hasShowTooltip($event, scope.row[`triggerDescriptionTip${scope.$index}`], scope.row.triggerDescription.join(','), 2.1)"
                                >
                                    <span
                                        style="marginRight: 22px"
                                        :class="setTextClass(false)"
                                    >
                                        {{scope.row.triggerDescription.join(',')}}
                                    </span>
                                </div>
                            </ui-tooltip>
                        </div>
                        <span v-else>{{$filter.emptyFilter()}}</span>
                    </template>
                </ui-table-column>
                <!-- 状态 -->
                <ui-table-column
                    prop="statusName"
                    min-width="100"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['statusName']?.selected"
                    key="status"
                    column-key="status"
                    filter-placement="bottom-end"
                    :class-name="filterStyleOperate.filterStatus['status'] ? 'common-status' : 'def-type-status15'"
                    :filters="filter?.ossStore?.status"
                    :filter-method="filter?.statusFilter"
                    :label="$t('allAlarmRulesList.table.label.statusName')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <span
                            :style="{
                                color: allAlarmRulesList.setStatusColor(scope.row.status),
                                marginRight: '16px'
                            }"
                        >
                            {{$filter.emptyFilter(scope.row.statusName)}}
                        </span>
                    </template>
                </ui-table-column>
                <!-- 操作 -->
                <ui-table-column
                    prop="operate"
                    min-width="160"
                    align="center"
                    fixed="right"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['operate']?.selected"
                    :label="$t('allAlarmRulesList.table.label.operate.title')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <div>
                            <span
                                :class="setTextClass(true, 'currency-style-right')"
                                @click="allAlarmRulesList.jumpDetail(scope.row)"
                            >
                                {{$t('allAlarmRulesList.table.label.operate.lookDetail')}}
                            </span>
                            <span
                                class="currency-style"
                                @click="emitValue('alarm-history', scope.row)"
                            >
                                {{$t('allAlarmRulesList.table.label.operate.alarmHistory')}}
                            </span>
                        </div>
                    </template>
                </ui-table-column>
            </ui-table>
        </ui-config-provider>
    </article>
</template>

<script lang="ts" setup>
import {tableClass, setTextClass, hasShowTooltip, setDiffClass} from 'utils/index.ts'; // 表格class类
import {RouterOperate, CurrentInstance, uiLocale} from 'utils/publicClass.ts';
import customOperate from '../custom/custom';
import FilterStyleOperate from './filterStyle';
import SetEmpty from './setEmpty';

interface PropsType {
    allAlarmRulesList: any;
    tableMaxHeight: number;
    filter: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

// defineEmits API 来替代 emits
const emitValue = defineEmits(['alarm-history']);

// const instanceProxy = new CurrentInstance().proxy;

// const path: string = instanceProxy.$defInfo.routerPath('deviceDetail');

// const routerOperate = new RouterOperate(path);
const filterStyleOperate = new FilterStyleOperate(props);
new SetEmpty(props)
</script>
