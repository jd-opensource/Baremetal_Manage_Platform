<template>
    <article class="operate-management-count hardware-list">
        <ui-config-provider :locale="uiLocale.locale">
            <ui-table
                border
                style="width: 100%;"
                v-loading="hardware.isLoading.value"
                :max-height="tableScrollOperate.tableMaxHeight.value || 'auto'"
                :class="[
                    'currency-count-table',
                    tableClass(hardware.reactiveArr.tableData, hardware.isLoading.value)
                ]"
                :data="hardware.reactiveArr.tableData"
                @get-table-ref="hardware.getTableRef"
            >
                <!-- sn -->
                <ui-table-column
                    min-width="110"
                    align="center"
                    fixed
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['sn']?.selected"
                    :label="$t('hardwareStatusList.table.label.sn')"
                    :has-user-template="true"
                >
                    <!-- 跳转详情 -->
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
                                    @click="routerOperate1?.jumpRouter({
                                        sn: scope.row.sn,
                                        type: 'basicInfo',
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
                    min-width="110"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['idcName']?.selected"
                    :label="$t('hardwareStatusList.table.label.idc')"
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
                <!-- 品牌 -->
                <ui-table-column
                    min-width="110"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['brand']?.selected"
                    :label="$t('hardwareStatusList.table.label.brand')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`brandTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`brandTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.brand}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`brandTip${scope.$index}`], scope.row.brand)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.brand)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 型号 -->
                <ui-table-column
                    min-width="110"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['model']?.selected"
                    :label="$t('hardwareStatusList.table.label.model')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="scope.row[`modelTip${scope.$index}`].showTooltip"
                            :disabled="!scope.row[`modelTip${scope.$index}`].showTooltip"
                        >
                            <template #content>
                                <div class="set-tooltip-width">
                                    <span>
                                        {{scope.row.model}}
                                    </span>
                                </div>
                            </template>
                            <div
                                class="more-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, scope.row[`modelTip${scope.$index}`], scope.row.model)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.model)}}
                                </span>
                            </div>
                        </ui-tooltip>
                    </template>
                </ui-table-column>
                <!-- 管理状态 -->
                <ui-table-column
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['manageStatus']?.selected"
                    :min-width="setDiffClass('140', '130')"
                    :label="$t('hardwareStatusList.table.label.managementStatus')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <span
                            v-if="scope.row.manageStatus"
                            :class="['default-unify', managementStatusColor(scope.row.manageStatus)]"
                        >
                            <span style="marginRight: 3px">
                                {{hardware.setManagementStatus(scope.row.manageStatus)}}
                                <ui-icon
                                    style="top: 1.5px"
                                    class="is-loading"
                                    v-if="['putawaying'].includes(scope.row.manageStatus)"
                                >
                                    <Loading/>
                                </ui-icon>
                            </span>
                            <!-- <ui-tooltip
                                placement="right"
                                v-if="scope?.row?.reason && ['putawayfail'].includes(scope?.row?.manageStatus)"
                            >
                                <template #content>
                                    <div>
                                        {{scope?.row?.reason}}
                                    </div>
                                </template>
                                <img
                                    class="warning-tip"
                                    alt=""
                                    :src="($defInfo.imagePath('uiTooltip') as unknown as string)"
                                />
                            </ui-tooltip> -->
                        </span>
                        <span v-else>
                            {{$filter.emptyFilter()}}
                        </span>
                    </template>
                </ui-table-column>
                <!-- CPU -->
                <ui-table-column
                    min-width="90"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['cpuStatus']?.selected"
                    :label="$t('hardwareStatusList.table.label.cpu')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <img
                            alt=""
                            class="common-img"
                            :src="hardware.setStatus(scope.row.cpuStatus)"
                        />
                    </template>
                </ui-table-column>
                <!-- 内存 -->
                <ui-table-column
                    min-width="90"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['memStatus']?.selected"
                    :label="$t('hardwareStatusList.table.label.storage')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <img
                            alt=""
                            class="common-img"
                            :src="hardware.setStatus(scope.row.memStatus)"
                        />
                    </template>
                </ui-table-column>
                <!-- 硬盘 -->
                <ui-table-column
                    min-width="90"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['diskStatus']?.selected"
                    :label="$t('hardwareStatusList.table.label.hardDisk')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <img
                            alt=""
                            class="common-img"
                            :src="hardware.setStatus(scope.row.diskStatus)"
                        />
                    </template>
                </ui-table-column>
                <!-- 网卡 -->
                <ui-table-column
                    min-width="90"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['nicStatus']?.selected"
                    :label="$t('hardwareStatusList.table.label.networkCard')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <img
                            alt=""
                            class="common-img"
                            :src="hardware.setStatus(scope.row.nicStatus)"
                        />
                    </template>
                </ui-table-column>
                <!-- 电源 -->
                <ui-table-column
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['powerStatus']?.selected"
                    :min-width="setDiffClass('110', '90')"
                    :label="$t('hardwareStatusList.table.label.powerSupply')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <img
                            alt=""
                            class="common-img"
                            :src="hardware.setStatus(scope.row.powerStatus)"
                        />
                    </template>
                </ui-table-column>
                <!-- 其他 -->
                <ui-table-column
                    min-width="90"
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['otherStatus']?.selected"
                    :label="$t('hardwareStatusList.table.label.other')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <img
                            alt=""
                            class="common-img"
                            :src="hardware.setStatus(scope.row.otherStatus)"
                        />
                    </template>
                </ui-table-column>
                <!-- 操作 -->
                <ui-table-column
                    align="center"
                    fixed="right"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo['operate']?.selected"
                    :label="$t('hardwareStatusList.table.label.operate')"
                    :min-width="setDiffClass('120', '90')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <span
                            :class="setTextClass(true)"
                            @click="hardware.jumpRouter(scope.row.sn, path2)"
                        >
                            {{$t('hardwareStatusList.table.label.opt')}}
                        </span>
                    </template>
                </ui-table-column>
            </ui-table>
        </ui-config-provider>
    </article>
</template>

<script lang="ts" setup>
import {tableClass, setTextClass, hasShowTooltip, managementStatusColor, setDiffClass} from 'utils/index.ts'; // 表格class类
import {RouterOperate, CurrentInstance, uiLocale} from 'utils/publicClass.ts';
import customOperate from '../customList/custom';
import SetEmpty from './setEmpty';

interface PropsType {
    hardware: any;
    tableScrollOperate: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

const instanceProxy = new CurrentInstance().proxy;

const path: string = instanceProxy.$defInfo.routerPath('deviceDetail');

const path2: string = instanceProxy.$defInfo.routerPath('faultLog');

// 实例化-路由操作
const routerOperate1 = new RouterOperate(path);

new SetEmpty(props);

</script>