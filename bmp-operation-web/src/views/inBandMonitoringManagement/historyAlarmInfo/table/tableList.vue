<template>
    <article class="operate-management-count history-alarm-list">  
        <ui-config-provider :locale="uiLocale.locale">
            <ui-table
                border
                style="width: 100%;"
                v-loading="historyAlarmInfo.isLoading.value"
                :max-height="tableMaxHeight || 'auto'"
                :class="tableClass(historyAlarmInfo.reactiveArr.tableData, historyAlarmInfo.isLoading.value)"
                :data="historyAlarmInfo.reactiveArr.tableData"
                @get-table-ref="historyAlarmInfo.getTableRef"
            >
                <!-- 报警时间 -->
                <!-- v-if="customOperate?.reactiveArr?.hasCustomInfo['sn']?.selected" -->

                <ui-table-column
                    prop="userName"
                    min-width="110"
                    align="center"
                    fixed
                    :label="$t('historyAlarmInfo.table.label.name')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                       <span>{{$filter.emptyFilter(scope.row.userName)}}</span>
                    </template>
                </ui-table-column>
                <ui-table-column
                    prop="alertTime"
                    min-width="160"
                    align="center"
                    :label="$t('historyAlarmInfo.table.label.alarmTime')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                       <span>{{$filter.emptyFilter(getDateMinutes(scope.row.alertTime))}}</span>
                    </template>
                </ui-table-column>
                <!-- 资源类型 -->
                <ui-table-column
                    prop="resourceName"
                    min-width="130"
                    align="center"
                    :label="$t('historyAlarmInfo.table.label.resourceName')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <span>{{$filter.emptyFilter(scope.row.resource)}}</span>
                    </template>
                </ui-table-column>
                <!-- 报警资源ID -->
                <ui-table-column
                    prop="resourceId"
                    min-width="205"
                    align="center"
                    :label="$t('historyAlarmInfo.table.label.resourceId')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <span>{{$filter.emptyFilter(scope.row.resourceId)}}</span>
                    </template>
                </ui-table-column>
                <!-- 触发条件 -->
                <ui-table-column
                    prop="triggerDescription"
                    min-width="140"
                    align="center"
                    :label="$t('historyAlarmInfo.table.label.triggerDescription')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`triggerDescriptionTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`triggerDescriptionTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.triggerDescription}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`triggerDescriptionTip${scope.$index}`], scope.row.triggerDescription, 2.1)"
                            >
                                <span
                                    style="marginRight: 22px"
                                    :class="setTextClass(false)"
                                >
                                    {{$filter.emptyFilter(scope.row.triggerDescription)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 报警值 -->
                <ui-table-column
                    prop="alertValue"
                    min-width="120"
                    align="center"
                    :label="$t('historyAlarmInfo.table.label.alertValue')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <div v-if="scope.row.alertValue">
                            <ui-tooltip
                                placement="bottom"
                                v-model="scope.row[`alertValueTip${scope.$index}`].showTooltip"
                                :disabled="!scope.row[`alertValueTip${scope.$index}`].showTooltip"
                            >
                                <template #content>
                                    <div class="set-tooltip-width">
                                        <span>
                                            {{scope.row.alertValue}}
                                        </span>
                                    </div>
                                </template>
                                <div
                                    class="more-text-ellipsis"
                                    @mouseenter="hasShowTooltip($event, scope.row[`alertValueTip${scope.$index}`], scope.row.alertValue, 2.1)"
                                >
                                    <span :class="setTextClass(false)">
                                        {{$filter.emptyFilter(scope.row.alertValue)}}
                                    </span>
                                </div>
                            </ui-tooltip>
                        </div>
                        <span v-else>{{$filter.emptyFilter()}}</span>
                    </template>
                </ui-table-column>
                <!-- 报警级别 -->
                <ui-table-column
                    prop="alertLevelDescription"
                    min-width="110"
                    align="center"
                    :label="$t('historyAlarmInfo.table.label.alertLevelDescription')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <span>{{$filter.emptyFilter(scope.row.alertLevelDescription)}}</span>
                    </template>
                </ui-table-column>
                <!-- 持续时间 -->
                <ui-table-column
                    prop="alertPeriod"
                    min-width="110"
                    align="center"
                    :label="$t('historyAlarmInfo.table.label.alertPeriod')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <span v-if="!['', void 0].includes(scope.row.alertPeriod)">
                            {{$filter.emptyFilter(scope.row.alertPeriod)}}{{$t('historyAlarmInfo.table.count.minute')}}
                        </span>
                        <span v-else>{{$filter.emptyFilter()}}</span>
                    </template>
                </ui-table-column>
                <!-- 通知对象 -->
                <ui-table-column
                    prop="userName"
                    :min-width="setDiffClass(130, 110)"
                    align="center"
                    fixed="right"
                    :label="$t('historyAlarmInfo.table.label.userName')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <!-- <router-link :to="$defInfo.routerPath('myProfile')"> -->
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.userName)}}
                            </span>
                        <!-- </router-link> -->
                    </template>
                </ui-table-column>
            </ui-table>
        </ui-config-provider>
    </article>
</template>

<script lang="ts" setup>
import {tableClass, setTextClass, hasShowTooltip, setDiffClass, getDateMinutes} from 'utils/index.ts'; // 表格class类
import {uiLocale} from 'utils/publicClass.ts';
// import customOperate from '../customList/custom';
import SetEmpty from './setEmpty';

interface PropsType {
    historyAlarmInfo: any;
    tableMaxHeight: number;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

new SetEmpty(props)
</script>
