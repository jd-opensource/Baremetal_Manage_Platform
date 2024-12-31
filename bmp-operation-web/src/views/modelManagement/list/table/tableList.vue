<template>
    <!-- element-plus国际化 -->
    <ui-config-provider :locale="uiLocale.locale">
        <!-- 表格-机型列表数据 -->
        <ui-table
            ref="tableRef"
            style="width: 100%;"
            v-loading="modelList.isLoading.value"
            border
            :max-height="tableMaxHeight"
            :class="tableClass(modelList.reactiveArr.tableData, modelList.isLoading.value)"
            :data="modelList.reactiveArr.tableData"
            @filter-change="filter.filterChange"
            @get-table-ref="filter.getTableRef"
        >
            <!-- 机型名称 -->
            <ui-table-column
                prop="name"
                min-width="140"
                align="center"
                fixed
                v-if="customOperate?.reactiveArr?.hasCustomInfo['name']?.selected"
                :has-user-template="true"
                :label="$t('modelList.content.modelName')"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`nameTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`nameTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.name}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`nameTip${scope.$index}`], scope.row.name)"
                        >
                            <span
                                :class="setTextClass(true)"
                                @click="modelDetailRouter.jumpRouter({
                                    type: 'modeName',
                                    deviceTypeId: scope.row.deviceTypeId,
                                    architecture: scope.row.architecture
                                })"
                            >
                                {{$filter.emptyFilter(scope.row.name)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 机型类型 -->
            <!-- def-type-status0 -->
            <!-- def-model-type-status -->
            <ui-table-column
                prop="deviceSeriesName"
                column-key="deviceSeries"
                align="center"
                key="deviceSeries"
                filter-placement="bottom-end"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['deviceSeries']?.selected"
                :class-name="filterStyleOperate.filterStatus['deviceSeries'] ? 'common-status' : 'def-type-status0'"
                :min-width="setDiffClass('150', '120')"
                :has-user-template="true"
                :label="$t('modelList.content.modelType')"
                :filters="filter.ossStore.deviceSeries"
                :filter-method="filter.deviceSeriesFilter"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`deviceSeriesNameTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`deviceSeriesNameTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.deviceSeriesName}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`deviceSeriesNameTip${scope.$index}`], scope.row.deviceSeriesName)"
                        >
                            <span
                                style="marginRight: 22px"
                                :class="setTextClass(false)"
                            >
                                {{$filter.emptyFilter(scope.row.deviceSeriesName)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 机房名称 -->
            <ui-table-column
                prop="idcName"
                align="center"
                min-width="110"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['idcName']?.selected"
                :has-user-template="true"
                :label="$t('modelList.content.machineRoomName')"
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
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.idcName)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 机房英文名称 -->
            <ui-table-column
                prop="idcNameEn"
                align="center"
                min-width="110"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['idcRegion']?.selected"
                :has-user-template="true"
                :label="$t('modelList.content.machineRoomCode')"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`idcNameEnTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`idcNameEnTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.idcNameEn}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`idcNameEnTip${scope.$index}`], scope.row.idcNameEn)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.idcNameEn)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 体系架构 -->
            <ui-table-column
                prop="architecture"
                align="center"
                min-width="120"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['architecture']?.selected"
                :label="$t('modelList.content.architecture')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`architectureTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`architectureTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.architecture}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`architectureTip${scope.$index}`], scope.row.architecture)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.architecture)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 引导模式 -->
            <ui-table-column
                align="center"
                min-width="140"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['bootMode']?.selected"
                :label="$t('modelList.content.bootMode')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span :class="setTextClass(false)">
                        {{$filter.emptyFilter(scope.row.bootMode)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 机型规格 -->
            <ui-table-column
                prop="deviceType"
                align="center"
                min-width="145"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['deviceType']?.selected"
                :label="$t('modelList.content.modelSpecification')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`deviceTypeTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`deviceTypeTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.deviceType}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`deviceTypeTip${scope.$index}`], scope.row.deviceType)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.deviceType)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 镜像 -->
            <ui-table-column
                prop="imageCount"
                align="center"
                min-width="60"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['image']?.selected"
                :has-user-template="true"
                :label="$t('modelList.content.image')"
            >
                <template #default="{scope}">
                    <span
                        v-if="scope.row.imageCount > 0"
                        :class="setTextClass(true)"
                        @click="modelDetailRouter.jumpRouter({
                            type: 'image',
                            deviceTypeId: scope.row.deviceTypeId,
                            architecture: scope.row.architecture
                        })"
                    >
                        {{scope.row.imageCount}}
                    </span>
                    <span v-else>
                        {{$filter.emptyFilter()}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 设备 -->
            <ui-table-column
                prop="deviceCount"
                align="center"
                min-width="60"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['deviceCount']?.selected"
                :has-user-template="true"
                :label="$t('modelList.content.device')"
            >
                <template #default="{scope}">
                    <span
                        v-if="scope.row.deviceCount > 0"
                        :class="setTextClass(true)"
                        @click="jumpDeviceList(scope.row.deviceTypeId)"
                    >
                        {{scope.row.deviceCount}}
                    </span>
                    <span v-else>
                        {{$filter.emptyFilter()}}
                    </span>
                </template>
            </ui-table-column>
            <!-- CPU -->
            <ui-table-column
                prop="cpuInfo"
                align="center"
                min-width="240"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['cpu']?.selected"
                :label="$t('modelList.content.CPU')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`cpuInfoTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`cpuInfoTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.cpuInfo}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`cpuInfoTip${scope.$index}`], scope.row.cpuInfo)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.cpuInfo)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 内存 -->
            <ui-table-column
                prop="memInfo"
                align="center"
                min-width="215"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['memory']?.selected"
                :label="$t('modelList.content.storage')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`memInfoTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`memInfoTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.memInfo}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`memInfoTip${scope.$index}`], scope.row.memInfo)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.memInfo)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 系统盘 -->
            <!-- <ui-table-column
                prop="svInfo"
                align="center"
                min-width="230"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['systemVolume']?.selected"
                :label="$t('modelList.content.sysDisc')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`svInfoTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`svInfoTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.svInfo}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`svInfoTip${scope.$index}`], scope.row.svInfo)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.svInfo)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column> -->
            <!-- RAID配置 -->
            <!-- <ui-table-column
                prop="raidCan"
                align="center"
                min-width="130"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['raidCan']?.selected"
                :label="$t('modelList.content.raidConfig')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`raidCanTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`raidCanTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.raidCan}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`raidCanTip${scope.$index}`], scope.row.raidCan)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.raidCan)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column> -->
            <!-- 系统盘RAID -->
            <!-- <ui-table-column
                prop="raid"
                align="center"
                min-width="190"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['systemVolumeRaid']?.selected"
                :label="$t('modelList.content.sysRAID')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`raidTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`raidTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.raid}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`raidTip${scope.$index}`], scope.row.raid)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.raid)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column> -->
            <!-- 数据盘 -->
            <!-- <ui-table-column
                prop="dvInfo"
                align="center"
                min-width="220"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['dataVolume']?.selected"
                :label="$t('modelList.content.dataDisc')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`dvInfoTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`dvInfoTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.dvInfo}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`dvInfoTip${scope.$index}`], scope.row.dvInfo)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.dvInfo)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column> -->
            <!-- 网卡 -->
            <ui-table-column
                prop="nicInfo"
                align="center"
                min-width="100"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['nicInfo']?.selected"
                :label="$t('modelList.content.networkCard')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`nicInfoTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`nicInfoTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.nicInfo}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`nicInfoTip${scope.$index}`], scope.row.nicInfo)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.nicInfo)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 网络设置 -->
            <ui-table-column
                prop="interfaceModeName"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['nic']?.selected"
                :min-width="setDiffClass('215', '110')"
                :label="$t('modelList.content.networkSettings')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`interfaceModeNameTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`interfaceModeNameTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.interfaceModeName}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`interfaceModeNameTip${scope.$index}`], scope.row.interfaceModeName)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.interfaceModeName)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- GPU -->
            <ui-table-column
                prop="gpuInfo"
                align="center"
                min-width="110"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['gpu']?.selected"
                :label="$t('modelList.content.GPU')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`gpuInfoTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`gpuInfoTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.gpuInfo}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`gpuInfoTip${scope.$index}`], scope.row.gpuInfo)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.gpuInfo)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 描述 -->
            <ui-table-column
                prop="description"
                align="center"
                min-width="140"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['remark']?.selected"
                :label="$t('modelList.content.desc')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`descriptionTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`descriptionTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span :class="setTextClass(false)">
                                    {{scope.row.description}}
                                </span>
                            </div>
                        </template>
                        <div
                            :class="['more-text-ellipsis', setTextClass(false)]"
                            @mouseenter="hasShowTooltip($event, scope.row[`descriptionTip${scope.$index}`], scope.row.description)"
                        >
                            <span :class="setTextClass(false)">
                                {{$filter.emptyFilter(scope.row.description)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 操作 -->
            <ui-table-column
                prop="operate"
                align="center"
                min-width="160"
                fixed="right"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['operate']?.selected"
                :has-user-template="true"
                :label="$t('modelList.content.operate.name')"
            >
                <template #default="{scope}">
                    <!-- 编辑提示 -->
                    <ui-tooltip
                        placement="bottom"
                        :disabled="!(scope.row.deviceCount > 0)"
                        :content="$t('modelList.edit')"
                    >
                        <span
                            :class="[
                                'currency-style-right',
                                (scope.row.deviceCount > 0) ? 'disabled-currency-style ' : 'currency-style'
                            ]"
                            @click="emitValue('edit-model', scope.row)"
                        >
                            {{$t('modelList.content.operate.edit')}}
                        </span>
                    </ui-tooltip>
                    <!-- 删除提示 -->
                    <ui-tooltip
                        placement="bottom"
                        :disabled="!(scope.row.instanceCount > 0 || scope.row.deviceCount > 0)"
                        :content="$t('modelList.tooltip')"
                    >
                        <span
                            :class="[
                                'currency-style-right',
                                (scope.row.instanceCount > 0 || scope.row.deviceCount > 0) ? 'disabled-currency-style ' : 'currency-style'
                            
                            ]"
                            @click="deleteModelOperate.deleteModelClick(scope.row)"
                        >
                            {{$t('modelList.content.operate.delete')}}
                        </span>
                    </ui-tooltip>
                    <!-- 下拉菜单 -->
                    <ui-dropdown>
                        <!-- 更多 -->
                        <template #dropdownTitle>
                            <p class="more-operate">
                                <span class="currency-style">
                                    {{$t('modelList.content.operate.more')}}
                                </span>
                                <img
                                    class="arrow-bottom-img"
                                    src="@/assets/img/listImg/arrow-bottom.png"
                                    alt=""
                                />
                            </p>
                        </template>
                        <!-- 添加相同机型 -->
                        <ui-dropdown-item>
                            <span
                                class="drop-down-operate-content"
                                @click="emitValue('add-model-template', scope.row)"
                            >
                                {{$t('modelList.content.operate.addModel')}}
                            </span>
                        </ui-dropdown-item>
                    </ui-dropdown>
                </template>
            </ui-table-column>
        </ui-table>
    </ui-config-provider>
</template>

<script lang="ts" setup>
import {uiLocale, RouterOperate, CurrentInstance} from 'utils/publicClass.ts';
import {tableClass, setTextClass, hasShowTooltip, setDiffClass} from 'utils/index.ts'; // 表格class类
import customOperate from '../custom/custom';
import FilterStyleOperate from './filterStyle';
import SetEmpty from './setEmpty';

interface PropsType {
    modelList: any;
    tableMaxHeight: any;
    // editModelOperate: any;
    deleteModelOperate: any;
    // addIdenticalModelOperate: any;
    filter: any;
    search: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

// defineEmits API 来替代 emits
const emitValue = defineEmits(['add-model', 'empty-click', 'edit-model', 'add-model-template']);

const instanceProxy = new CurrentInstance().proxy;

const deviceListRouter = new RouterOperate(instanceProxy.$defInfo.routerPath('deviceList'));

const modelDetailRouter = new RouterOperate(instanceProxy.$defInfo.routerPath('modelDetail'));

const filterStyleOperate = new FilterStyleOperate(props);

const jumpDeviceList = (deviceTypeId: string) => {
    deviceListRouter.jumpRouter({params: deviceTypeId, status: true, type: 'deviceTypeId'})
};

new SetEmpty(props, emitValue);

</script>