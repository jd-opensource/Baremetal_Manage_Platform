<template>
    <article class="operate-management-count fault-rules-list">
        <ui-config-provider :locale="uiLocale.locale">
            <ui-table
                border
                style="width: 100%;"
                v-loading="faultRules.isLoading.value"
                :max-height="tableScrollOperate.tableMaxHeight.value"
                :class="[
                    'currency-count-table',
                    tableClass(faultRules.reactiveArr.tableData, faultRules.isLoading.value)
                ]"
                :data="faultRules.reactiveArr.tableData"
                @get-table-ref="faultRules.getTableRef"
                @selection-change="faultRules.filterChange"
                @cell-click="faultRules.cellClick"
            >
                <ui-table-column
                    type="selection"
                    align="center"
                    fixed
                />
                <!-- 故障名称 -->
                <ui-table-column
                    prop="alertName"
                    min-width="110"
                    align="center"
                    fixed
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['alertName']?.selected"
                    :label="$t('faultRulesList.table.label.faultName')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`alertNameTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`alertNameTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.alertName}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`alertNameTip${scope.$index}`], scope.row.alertName)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.alertName)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 故障配件 -->
                <ui-table-column
                    prop="alertSpec"
                    min-width="120"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['alertSpec']?.selected"
                    :label="$t('faultRulesList.table.label.faultAccessories')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`alertSpecTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`alertSpecTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.alertSpec}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`alertSpecTip${scope.$index}`], scope.row.alertSpec)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.alertSpec)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 故障类型 -->
                <ui-table-column
                    prop="alertType"
                    min-width="110"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['alertType']?.selected"
                    :label="$t('faultRulesList.table.label.faultType')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`alertTypeTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`alertTypeTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.alertType}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`alertTypeTip${scope.$index}`], scope.row.alertType)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.alertType)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 判定条件 -->
                <ui-table-column
                    prop="alertCondition"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['alertCondition']?.selected"
                    :min-width="setDiffClass('140', '120')"
                    :label="$t('faultRulesList.table.label.judgmentConditions')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`alertConditionTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`alertConditionTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.alertCondition}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`alertConditionTip${scope.$index}`], scope.row.alertCondition)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.alertCondition)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 判定阈值 -->
                <ui-table-column
                    prop="alertThreshold"
                    align="center"
                    min-width="150"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['alertThreshold']?.selected"
                    :label="$t('faultRulesList.table.label.decisionThreshold')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`alertThresholdTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`alertThresholdTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.alertThreshold}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`alertThresholdTip${scope.$index}`], scope.row.alertThreshold)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.alertThreshold)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 故障等级 -->
                <ui-table-column
                    prop="alertLevel"
                    min-width="110"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['alertLevel']?.selected"
                    :label="$t('faultRulesList.table.label.faultLevel')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`alertLevelTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`alertLevelTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.alertLevel}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`alertLevelTip${scope.$index}`], scope.row.alertLevel)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.alertLevel)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 故障描述 -->
                <ui-table-column
                    prop="alertDesc"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['alertDesc']?.selected"
                    :min-width="setDiffClass('130', '110')"
                    :label="$t('faultRulesList.table.label.faultDesc')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`alertDescTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`alertDescTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.alertDesc}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`alertDescTip${scope.$index}`], scope.row.alertDesc)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.alertDesc)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 故障推送 -->
                <ui-table-column
                    prop="pushStatus"
                    min-width="110"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['pushStatus']?.selected"
                    :label="$t('faultRulesList.table.label.faultPush')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-switch
                            size="small"
                            style="--el-switch-on-color: #108ee9;"
                            v-model="scope.row.pushStatus"
                            :loading="scope.row.pushLoading"
                            :before-change="() => beforeChange(scope.row)"
                        />
                    </template>
                </ui-table-column>
                <!-- 启用状态 -->
                <ui-table-column
                    prop="useStatus"
                    min-width="110"
                    align="center"
                    fixed="right"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['useStatus']?.selected"
                    :label="$t('faultRulesList.table.label.enableStatus')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-switch
                            size="small"
                            style="--el-switch-on-color: #108ee9;"
                            v-model="scope.row.useStatus"
                            :loading="scope.row.useLoading"
                            :before-change="() => beforeChange2(scope.row)"
                        />
                    </template>
                </ui-table-column>
            </ui-table>
        </ui-config-provider>
        
    </article>
</template>

<script lang="ts" setup>
import {tableClass, hasShowTooltip, setDiffClass} from 'utils/index.ts'; // 表格class类
import SetEmpty from './setEmpty';
import {uiLocale} from 'utils/publicClass.ts';
import customOperate from '../customList/custom';

interface PropsType {
    faultRules: any;
    tableScrollOperate: any;
}

// defineEmits API 用来代替emit
const emitValue = defineEmits(['push-click', 'use-click']);
// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {
    faultRules: {},
    // customOperate: {}
});

new SetEmpty(props)

const beforeChange = (item: any) => {
    
    return new Promise((resolve) => {
        // nextTick(() => {
    emitValue('push-click', item)
        // })
    return resolve(false);
  })
}

const beforeChange2 = (item: any) => {
     
    return new Promise((resolve) => {
    // nextTick(() => {
        emitValue('use-click', item)
    // })
    return resolve(false);
  })
};
</script>
