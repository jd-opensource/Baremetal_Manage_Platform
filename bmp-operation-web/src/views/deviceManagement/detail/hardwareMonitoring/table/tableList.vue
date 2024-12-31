<template>
    <ui-config-provider :locale="uiLocale.locale">
        <ui-table
            ref="tableRef"
            style="width: 100%;"
            v-loading="faultLogTable.isTableLoading.value"
            border
            :max-height="tableScrollOperate?.tableMaxHeight?.value || 'auto'"
            :class="tableClass(faultLogTable.reactiveArr.tableData, faultLogTable.isTableLoading.value)"
            :data="faultLogTable.reactiveArr.tableData"
        >
            <!-- 序号 -->
            <ui-table-column
                prop="logid"
                min-width="80"
                align="center"
                fixed
                v-if="customOperate?.reactiveArr?.hasCustomInfo['logid']?.selected"
                :label="$t('deviceDetail.table.label.serialNumber')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span>
                        {{$filter.emptyFilter(scope.row.logid)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 故障类型 -->
            <ui-table-column
                prop="alertType"
                min-width="110"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['alertType']?.selected"
                :label="$t('deviceDetail.table.label.faultType')"
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
            <!-- 故障描述 -->
            <ui-table-column
                prop="alertDesc"
                min-width="120"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['alertDesc']?.selected"
                :label="$t('deviceDetail.table.label.faultDesc')"
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
            <!-- 故障报警时间 -->
            <ui-table-column
                prop="eventTime"
                min-width="140"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['eventTime']?.selected"
                :label="$t('deviceDetail.table.label.faultAlarmTime')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span>
                        {{$filter.emptyFilter(scope.row.eventTime)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 更新时间 -->
            <ui-table-column
                prop="updateTime"
                min-width="130"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['updateTime']?.selected"
                :label="$t('deviceDetail.table.label.updateTime')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span>
                        {{$filter.emptyFilter(scope.row.updateTime)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 报警次数 -->
            <ui-table-column
                prop="count"
                min-width="90"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['count']?.selected"
                :label="$t('deviceDetail.table.label.alarmNum')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span>
                        {{$filter.emptyFilter(scope.row.count)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 状态 -->
            <ui-table-column
                prop="status"
                min-width="90"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['status']?.selected"
                :label="$t('deviceDetail.table.label.status')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span :style="{color: faultLogTable.setColor(scope.row.status)}">
                        {{faultLogTable.setStatus(scope.row.status)}}
                    </span>
                </template>
            </ui-table-column>
        </ui-table>
    </ui-config-provider>
</template>

<script lang="ts" setup>
import {tableClass, hasShowTooltip} from 'utils/index.ts'; // 表格class类
import {uiLocale} from 'utils/publicClass.ts';
import SetEmpty from './setEmpty';

interface PropsType {
    faultLogTable: any;
    tableScrollOperate: any;
    customOperate: any
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {
    faultLogTable: {},
    // customOperate: {}
});

new SetEmpty(props);
// defineEmits API 来替代 emits
const emitValue = defineEmits(['edit']);
</script>
