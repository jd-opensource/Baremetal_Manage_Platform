<template>
    <article class="operate-management-count fault-log-list">  
        <ui-config-provider :locale="uiLocale.locale">
            <ui-table
                border
                style="width: 100%;"
                v-loading="faultLog.isLoading.value"
                :max-height="tableScrollOperate?.tableMaxHeight?.value"
                :class="tableClass(faultLog.reactiveArr.tableData, faultLog.isLoading.value)"
                :data="faultLog.reactiveArr.tableData"
                @get-table-ref="faultLog.getTableRef"
            >
                <!-- LogID -->
                <ui-table-column
                    prop="logID"
                    min-width="110"
                    align="center"
                    fixed
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['logid']?.selected"
                    :label="$t('faultLogList.table.label.logId')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`logidTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`logidTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.logid}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`logidTip${scope.$index}`], scope.row.logid)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.logid)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- sn -->
                <ui-table-column
                    prop="sn"
                    min-width="110"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['sn']?.selected"
                    :label="$t('faultLogList.table.label.sn')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`snTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`snTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.sn}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`snTip${scope.$index}`], scope.row.sn)"
                            >
                                <span
                                    :class="setTextClass(true)"
                                    @click="routerOperate?.jumpRouter({
                                        sn: scope.row.sn,
                                        type: 'hardwareMonitoring',
                                        deviceId: scope.row.deviceId
                                    })"
                                >
                                    {{$filter.emptyFilter(scope.row.sn)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 机房 -->
                <ui-table-column
                    prop="idcName"
                    min-width="110"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['idcName']?.selected"
                    :label="$t('faultLogList.table.label.idc')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`idcNameTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`idcNameTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.idcName}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`idcNameTip${scope.$index}`], scope.row.idcName)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.idcName)}}
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
                    :label="$t('faultLogList.table.label.level')"
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
                <!-- 故障类型 -->
                <ui-table-column
                    prop="alertType"
                    min-width="110"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['alertType']?.selected"
                    :label="$t('faultLogList.table.label.type')"
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
                <!-- 故障配件 -->
                <ui-table-column
                    prop="alertPart"
                    min-width="140"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['alertPart']?.selected"
                    :label="$t('faultLogList.table.label.accessory')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`alertPartTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`alertPartTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.alertPart}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`alertPartTip${scope.$index}`], scope.row.alertPart)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.alertPart)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- iloIP -->
                <ui-table-column
                    prop="iloIp"
                    min-width="110"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['iloIp']?.selected"
                    :label="$t('faultLogList.table.label.iloip')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`iloIpTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`iloIpTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.iloIp}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`iloIpTip${scope.$index}`], scope.row.iloIp)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.iloIp)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 故障描述 -->
                <ui-table-column
                    prop="alertDesc"
                    min-width="150"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['alertDesc']?.selected"
                    :label="$t('faultLogList.table.label.desc')"
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
                <!-- 故障原始日志 -->
                <ui-table-column
                    prop="alertContent"
                    min-width="130"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['alertContent']?.selected"
                    :label="$t('faultLogList.table.label.originalLog')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`alertContentTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`alertContentTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.alertContent}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`alertContentTip${scope.$index}`], scope.row.alertContent)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.alertContent)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 状态 -->
                <ui-table-column
                    prop="status"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['status']?.selected"
                    :min-width="setDiffClass('110', '80')"
                    :label="$t('faultLogList.table.label.status')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`statusTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`statusTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.status}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`statusTip${scope.$index}`], scope.row.status)"
                            >
                                <span :style="{color: faultLog.setColor(scope.row.status)}">
                                    {{faultLog.setStatus(scope.row.status)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 报警次数 -->
                <ui-table-column
                    prop="count"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['count']?.selected"
                    :min-width="setDiffClass('130', '80')"
                    :label="$t('faultLogList.table.label.numberOfAlarms')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`countTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`countTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.count}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`countTip${scope.$index}`], scope.row.count)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.count)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 带外故障发现时间 -->
                <ui-table-column
                    prop="collectTime"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['collectTime']?.selected"
                    :min-width="setDiffClass('220', '190')"
                    :label="$t('faultLogList.table.label.faultDetectionTime')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`collectTimeTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`collectTimeTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.collectTime}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`collectTimeTip${scope.$index}`], scope.row.collectTime)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.collectTime)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 首次故障报警时间 -->
                <ui-table-column
                    prop="eventTime"
                    min-width="170"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['eventTime']?.selected"
                    :label="$t('faultLogList.table.label.faultAlarmTime')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`eventTimeTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`eventTimeTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.eventTime}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`eventTimeTip${scope.$index}`], scope.row.eventTime)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.eventTime)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 更新时间 -->
                <ui-table-column
                    prop="updateTime"
                    min-width="170"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['updateTime']?.selected"
                    :label="$t('faultLogList.table.label.updateTime')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`updateTimeTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`updateTimeTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.updateTime}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`updateTimeTip${scope.$index}`], scope.row.updateTime)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.updateTime)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 操作 -->
                <ui-table-column
                    min-width="80"
                    align="center"
                    fixed="right"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['operate']?.selected"
                    :label="$t('faultLogList.table.label.operate')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="left"
                            :disabled="!scope.row.status"
                            :content="$t('faultLogList.table.operate.tip')"
                        >
                            <span
                                :class="[
                                    !scope.row.status
                                    ? setTextClass(true)
                                    : 'disabled-currency-style'
                                ]"
                                @click="emitValue('no-untreated', faultLog.reactiveArr.selectArr[scope.$index])"
                            >
                                {{$t('faultLogList.table.operate.title')}}
                            </span>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
            </ui-table>
        </ui-config-provider>
    </article>
</template>

<script lang="ts" setup>
import {tableClass, setTextClass, hasShowTooltip, setDiffClass} from 'utils/index.ts'; // 表格class类
import {RouterOperate, CurrentInstance, uiLocale} from 'utils/publicClass.ts';
import customOperate from '../customList/custom';
import SetEmpty from './setEmpty';

interface PropsType {
    faultLog: any;
    tableScrollOperate: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

// defineEmits API 来替代 emits
const emitValue = defineEmits(['no-untreated']);

const instanceProxy = new CurrentInstance().proxy;

const path: string = instanceProxy.$defInfo.routerPath('deviceDetail');

const routerOperate = new RouterOperate(path);

new SetEmpty(props)
</script>
