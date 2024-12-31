<template>
    <!-- element-plus内置国际化 -->
    <ui-config-provider :locale="uiLocale.locale">
        <!-- 表格数据-设备列表 -->
        <ui-table
            ref="tableRef"
            v-loading="deviceList.isLoading.value"
            :class="[
                'currency-count-table',
                tableClass(deviceList.reactiveArr.tableData, deviceList.isLoading.value)
            ]"
            :border="true"
            :data="deviceList.reactiveArr.tableData"
            :max-height="tableMaxHeight"
            :row-key="deviceList.getRowKeys"
            @row-click="deviceList.rowClick"
            @filter-change="filterOperate.filterChange"
            @get-table-ref="filterOperate.getTableRef"
            @cell-click="deviceList.cellClick"
            @selection-change="deviceList.handleSelectionChange"
        >
        <!-- @cell-click="filterOperate.cellClick" -->

            <ui-table-column
                type="selection"
                align="center"
                :reserve-selection="true"
            />
            <!-- SN -->
            <ui-table-column
                prop="sn"
                min-width="115"
                align="center"
                fixed
                v-if="customOperate?.reactiveArr?.hasCustomInfo['sn']?.selected"
                :has-user-template="true"
                :label="$t('deviceList.content.label.sn')"
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
                            <router-link
                                v-if="!Object.is(deviceList.code.value, 403) && !Object.is(deviceList.monitorCode.value, 403)"
                                :to="$defInfo.routerPath('deviceDetail') + '?sn=' + scope.row.sn + '&deviceId=' + scope.row.deviceId + '&instanceId=' + scope.row.instanceId + '&type=basicInfo'">
                                <span class="currency-style">
                                    {{$filter.emptyFilter(scope.row.sn)}}
                                </span>
                            </router-link>
                            <span
                                class="currency-style"
                                v-else
                                @click="deviceList.jumpClick(routerOperate, scope.row)"
                            >
                                {{$filter.emptyFilter(scope.row.sn)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 机型名称 -->
            <ui-table-column
                prop="deviceTypeName"
                align="center"
                column-key="deviceTypeId"
                key="deviceTypeId"
                filter-placement="bottom-end"
                min-width="160"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['deviceSeries']?.selected"
                :class-name="filterStyleOt.filterStatus['name'] ? 'common-status' : 'def-type-status8'"
                :has-user-template="true"
                :label="$t('deviceList.content.label.name')"
                :filters="filterOperate.ossStore.deviceTypeId"
                :filter-method="filterOperate.nameFilter"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`deviceTypeNameTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`deviceTypeNameTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.deviceTypeName}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`deviceTypeNameTip${scope.$index}`], scope.row.deviceTypeName)"
                        >
                            <span style="marginRight: 18px">
                                {{$filter.emptyFilter(scope.row.deviceTypeName)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 机型类型 -->
            <ui-table-column
                align="center"
                prop="deviceSeriesName"
                column-key="deviceSeries"
                key="deviceSeries"
                filter-placement="bottom-end"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['deviceTypeName']?.selected"
                :class-name="filterStyleOt.filterStatus['deviceSeries'] ? 'common-status' : 'def-type-status9'"
                :min-width="setDiffClass('150', '126')"
                :has-user-template="true"
                :filters="filterOperate.ossStore.deviceSeries"
                :filter-method="filterOperate.deviceSeriesFilter"
                :label="$t('deviceList.content.label.modelName')"
            >
                <template #default="{scope}">
                    <span style="marginRight: 20px">
                        {{$filter.emptyFilter(scope.row.deviceSeriesName)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 管理状态 -->
            <ui-table-column
                prop="manageStatusName"
                align="center"
                column-key="manageStatus"
                key="manageStatus"
                filter-placement="bottom-end"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['manageStatusName']?.selected"
                :class-name="filterStyleOt.filterStatus['manageStatus'] ? 'common-status' : 'def-type-status10'"
                :min-width="setDiffClass('160', '136')"
                :has-user-template="true"
                :filters="filterOperate.ossStore.manageStatus"
                :filter-method="filterOperate.manageStatusFilter"
                :label="$t('deviceList.content.label.managementStatus')"
            >
                <template #default="{scope}">
                    <!-- 动态获取class -->
                    <div
                        style="marginRight: 16px; display: flex; alignItems: center;"
                        v-if="scope.row.manageStatusName"
                        :class="['default-unify', managementStatusColor(scope.row.manageStatus)]"
                    >
                        <span style="marginRight: 3px">
                            {{scope.row.manageStatusName}}
                            <ui-icon
                                style="top: 1.5px"
                                class="is-loading"
                                v-if="['putawaying'].includes(scope.row.manageStatus)"
                            >
                                <Loading/>
                            </ui-icon>
                        </span>
                        <ui-tooltip
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
                        </ui-tooltip>
                    </div>
                    <div
                        style="marginLeft: 10px"
                        v-else
                    >
                        --
                    </div>
                </template>
            </ui-table-column>
            <!-- 采集 -->
            <ui-table-column
                align="center"
                prop="collectStatus"
                column-key="collectStatus"
                key="collectStatus"
                filter-placement="bottom-end"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['collectStatus']?.selected"
                :min-width="setDiffClass('150', '110')"
                :class-name="filterStyleOt.filterStatus['collectStatus'] ? 'common-status' : 'def-type-status11'"
                :label="$t('deviceList.content.label.collectionStatus')"
                :filters="filterOperate.ossStore.collectStatus"
                :filter-method="filterOperate.statusFilter"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span style="marginRight: 20px">
                        {{$filter.emptyFilter(deviceList.setCollected(scope.row.collectStatus))}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 机房名称 -->
            <ui-table-column
                prop="idcName"
                align="center"
                min-width="110"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['idcName']?.selected"
                :label="$t('deviceList.content.label.computerRoomName')"
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
            <!-- 机柜编码 -->
            <ui-table-column
                prop="cabinet"
                min-width="95"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['cabinet']?.selected"
                :label="$t('deviceList.content.label.cabinetCode')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`cabinetTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`cabinetTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.cabinet}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`cabinetTip${scope.$index}`], scope.row.cabinet)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.cabinet)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 所在U位 -->
            <ui-table-column
                prop="uPosition"
                min-width="110"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['uPosition']?.selected"
                :label="$t('deviceList.content.label.uBit')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`uPositionTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`uPositionTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.uPosition}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`uPositionTip${scope.$index}`], scope.row.uPosition)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.uPosition)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 品牌 -->
            <ui-table-column
                prop="brand"
                min-width="90"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['brand']?.selected"
                :label="$t('deviceList.content.label.brand')"
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
                prop="model"
                min-width="130"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['model']?.selected"
                :label="$t('deviceList.content.label.model')"
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
            <!-- 体系架构 -->
            <ui-table-column
                prop="architecture"
                min-width="100"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['architecture']?.selected"
                :label="$t('deviceList.content.label.architecture')"
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
                            <span>
                                {{$filter.emptyFilter(scope.row.architecture)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 镜像名称 -->
            <ui-table-column
                prop="imageName"
                min-width="140"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['imageName']?.selected"
                :label="$t('deviceList.content.label.instanceImage')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`imageNameTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`imageNameTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.imageName}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`imageNameTip${scope.$index}`], scope.row.imageName)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.imageName)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 实例名称 -->
            <ui-table-column
                prop="instanceName"
                min-width="110"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['instanceName']?.selected"
                :label="$t('deviceList.content.label.instanceName')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`instanceNameTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`instanceNameTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.instanceName}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`instanceNameTip${scope.$index}`], scope.row.instanceName)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.instanceName)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 实例ID -->
            <ui-table-column
                prop="instanceId"
                min-width="150"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['instanceId']?.selected"
                :label="$t('deviceList.content.label.instanceID')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`instanceIdTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`instanceIdTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.instanceId}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`instanceIdTip${scope.$index}`], scope.row.instanceId)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.instanceId)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 实例状态 -->
            <ui-table-column
                prop="instanceStatusName"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['instanceStatus']?.selected"
                :min-width="setDiffClass('210', '150')"
                :label="$t('deviceList.content.label.instanceStatus')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span
                        style="margin: 0 auto 0"
                        :class="['default-unify', instanceStatusColor(scope.row.instanceStatus)]"
                    >
                        <span style="marginRight: 3px">
                            {{$filter.emptyFilter(scope.row.instanceStatusName)}}
                        </span>
                        <ui-icon
                            class="is-loading"
                            v-if="DeviceStaticData.loadingStatusData.includes(scope.row.instanceStatus)"
                        >
                            <Loading/>
                        </ui-icon>
                        <ui-tooltip
                            placement="right"
                            v-if="DeviceStaticData.FailedData.includes(scope.row.instanceStatus)"
                        >
                            <template #content>
                                <div>
                                    <div v-if="['Creation failed'].includes(scope.row.instanceStatus)">
                                        {{
                                            ['Creation failed'].includes(scope.row.instanceStatus)
                                            ? $filter.emptyFilter(scope.row.instanceReason)
                                            : $t('deviceList.operate.error.tip')
                                        }}
                                    </div>
                                    <div v-if="[
                                        'Startup failed',
                                        'Shutdown failed',
                                        'Reboot failed',
                                        'Delete failed',
                                        'Reinstall failed',
                                        'Resetpasswd failed'
                                    ].includes(scope.row.instanceStatus)">
                                        {{$t('deviceList.operate.error.tip')}}
                                    </div>
                                </div>
                            </template>
                            <img
                                class="warning-tip"
                                alt=""
                                :src="($defInfo.imagePath('uiTooltip') as unknown as string)"
                            />
                        </ui-tooltip>
                    </span>
                </template>
            </ui-table-column>
            <!-- 锁定状态 -->
            <ui-table-column
                prop="lockedname"
                min-width="105"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['instanceLockedStatus']?.selected"
                :label="$t('deviceList.content.label.lockStatus')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span :class="setTextClass(false)">
                        {{$filter.emptyFilter(scope.row.lockedname)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 实例属主 -->
            <ui-table-column
                prop="userName"
                min-width="120"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['instanceOwer']?.selected"
                :label="$t('deviceList.content.label.instanceOwner')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`userNameTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`userNameTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.userName}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`userNameTip${scope.$index}`], scope.row.userName)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.userName)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- CPU -->
            <ui-table-column
                prop="cpuInfo"
                min-width="150"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['cpuInfo']?.selected"
                :label="$t('deviceList.content.label.CPU')"
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
                            <span>
                                {{$filter.emptyFilter(scope.row.cpuInfo)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 内存 -->
            <ui-table-column
                prop="memInfo"
                min-width="130"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['memInfo']?.selected"
                :label="$t('deviceList.content.label.storage')"
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
                            <span>
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
                min-width="150"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['svInfo']?.selected"
                :label="$t('deviceList.content.label.sysDisc')"
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
                            <span>
                                {{$filter.emptyFilter(scope.row.svInfo)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column> -->
            <!-- 数据盘 -->
            <!-- <ui-table-column
                prop="dvInfo"
                min-width="220"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['dvInfo']?.selected"
                :label="$t('deviceList.content.label.dataDisc')"
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
                            <span>
                                {{$filter.emptyFilter(scope.row.dvInfo)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column> -->
            <!-- GPU -->
            <ui-table-column
                prop="gpuInfo"
                min-width="110"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['gpuInfo']?.selected"
                :label="$t('deviceList.content.label.GPU')"
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
                            <span>
                                {{$filter.emptyFilter(scope.row.gpuInfo)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 网口1交换机IP -->
            <ui-table-column
                prop="switchIp1"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['switchIp1']?.selected"
                :min-width="setDiffClass('146', '120')"
                :label="$t('deviceList.content.label.switchIPNetworkPortOne')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`switchIp1Tip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`switchIp1Tip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.switchIp1}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`switchIp1Tip${scope.$index}`], scope.row.switchIp1)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.switchIp1)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 网口2交换机IP -->
            <ui-table-column
                align="center"
                prop="switchIp2"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['switchIp2']?.selected"
                :min-width="setDiffClass('156', '120')"
                :label="$t('deviceList.content.label.switchIPNetworkPortTwo')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`switchIp2Tip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`switchIp2Tip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.switchIp2}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`switchIp2Tip${scope.$index}`], scope.row.switchIp2)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.switchIp2)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 网口1（eth0）上联端口 -->
            <ui-table-column
                prop="switchPort1"
                min-width="180"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['switchPort1']?.selected"
                :label="$t('deviceList.content.label.networkPortOneUplinkPort')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`switchPort1Tip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`switchPort1Tip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.switchPort1}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`switchPort1Tip${scope.$index}`], scope.row.switchPort1)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.switchPort1)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 网口2（eth1）上联端口 -->
            <ui-table-column
                prop="switchPort2"
                min-width="180"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['switchPort2']?.selected"
                :label="$t('deviceList.content.label.networkPortTwoUplinkPort')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`switchPort2Tip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`switchPort2Tip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.switchPort2}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`switchPort2Tip${scope.$index}`], scope.row.switchPort2)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.switchPort2)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 网络设置 -->
            <ui-table-column
                prop="interfaceModeName"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['nicInfo']?.selected"
                :min-width="setDiffClass('215', '110')"
                :label="$t('deviceList.content.label.networkSettings')"
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
                            <span>
                                {{$filter.emptyFilter(scope.row.interfaceModeName)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 带外IP -->
            <ui-table-column
                prop="iloIp"
                align="center"
                min-width="120"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['iloIp']?.selected"
                :label="$t('deviceList.content.label.outOfBandIP')"
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
            <!-- 内网IPv4(eth0) -->
            <ui-table-column
                prop="privateIpv4"
                align="center"
                min-width="120"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['privateIpv4']?.selected"
                :label="$t('deviceList.content.label.intranceIPv4')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`privateIpv4Tip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`privateIpv4Tip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.privateIpv4}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`privateIpv4Tip${scope.$index}`], scope.row.privateIpv4)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.privateIpv4)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 内网IPv4(eth1) -->
            <ui-table-column
                prop="privateEth1Ipv4"
                align="center"
                min-width="120"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['privateEth1Ipv4']?.selected"
                :label="$t('deviceList.content.label.intranceIPv4First')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`privateEth1Ipv4Tip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`privateEth1Ipv4Tip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.privateEth1Ipv4}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`privateEth1Ipv4Tip${scope.$index}`], scope.row.privateEth1Ipv4)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.privateEth1Ipv4)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 内网IPv6(eth0) -->
            <ui-table-column
                prop="privateIpv6"
                align="center"
                min-width="140"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['privateIpv6']?.selected"
                :label="$t('deviceList.content.label.intranetIPv6')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`privateIpv6Tip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`privateIpv6Tip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.privateIpv6}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`privateIpv6Tip${scope.$index}`], scope.row.privateIpv6)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.privateIpv6)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 内网IPv6(eth1) -->
            <ui-table-column
                prop="privateEth1Ipv6"
                align="center"
                min-width="140"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['privateEth1Ipv6']?.selected"
                :label="$t('deviceList.content.label.intranetIPv6First')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`privateEth1Ipv6Tip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`privateEth1Ipv6Tip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.privateEth1Ipv6}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`privateEth1Ipv6Tip${scope.$index}`], scope.row.privateEth1Ipv6)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.privateEth1Ipv6)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 创建时间 -->
            <ui-table-column
                prop="createdTime"
                align="center"
                min-width="160"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['createdTime']?.selected"
                :label="$t('deviceList.content.label.createTime')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    {{$filter.emptyFilter(scope.row.createdTime)}}
                </template>
            </ui-table-column>
            <!-- 备注 -->
            <ui-table-column
                prop="description"
                align="center"
                min-width="120"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['description']?.selected"
                :label="$t('deviceList.content.label.remark')"
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
                                <span>
                                    {{scope.row.description}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event,scope.row[`descriptionTip${scope.$index}`], scope.row.description)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.description)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 操作 -->
            <ui-table-column
                fixed="right"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['operate']?.selected"
                :min-width="setDiffClass('200', '160')"
                :has-user-template="true"
                :label="$t('deviceList.content.label.operate.name')"
            >
                <template #default="{scope}">
                    <!-- || !!(scope.row.deviceTypeId?.length) -->
                    <!-- 上架 -->
                    <ui-tooltip
                        placement="bottom"
                        v-if="!tooltipOpt.up(scope.row)"
                        :disabled="tooltipOpt.up(scope.row)"
                        :content="$t('deviceList.tooltip.up')"
                    >
                        <span
                            :class="[
                                'currency-style-right',
                                tooltipOpt.up(scope.row) ? 'currency-style' : 'no-drop-disabled'
                            ]"
                            @click="() => emitValue('up-down-delete', scope.row, 'up', tooltipOpt.up(scope.row))"
                        >
                            {{$t('deviceList.content.label.operate.up')}}
                        </span>
                    </ui-tooltip>
                    <ui-tooltip
                        placement="bottom"
                        v-else
                        :disabled="!!(scope.row.deviceTypeId?.length)"
                        :content="$t('deviceList.tooltip.model')"
                    >
                        <span
                            :class="[
                                'currency-style-right',
                                scope.row.deviceTypeId?.length ? 'currency-style' : 'no-drop-disabled'
                            ]"
                            @click="() => emitValue('up-down-delete', scope.row, 'up', scope.row.deviceTypeId?.length)"
                        >
                            {{$t('deviceList.content.label.operate.up')}}
                        </span>
                    </ui-tooltip>
                    <!-- 下架 -->
                    <ui-tooltip
                        placement="bottom"
                        :disabled="tooltipOpt.down(scope.row)"
                        :content="$t('deviceList.tooltip.down')"
                    >
                        <span
                            :class="[
                                'currency-style-right',
                                tooltipOpt.down(scope.row) ? 'currency-style' : 'no-drop-disabled'
                            ]"
                            @click="() => emitValue('up-down-delete', scope.row, 'down', tooltipOpt.down(scope.row))"
                        >
                            {{$t('deviceList.content.label.operate.down')}}
                        </span>
                    </ui-tooltip>
                    <!-- 下拉菜单 -->
                    <ui-dropdown>
                        <!-- 更多 -->
                        <template #dropdownTitle>
                            <p class="more-operate">
                                <span class="currency-style">
                                    {{$t('deviceList.content.label.operate.more')}}
                                </span>
                                <img
                                    class="arrow-bottom-img"
                                    src="@/assets/img/listImg/arrow-bottom.png"
                                    alt=""
                                />
                            </p>
                        </template>
                        <!-- 开机 -->
                        <ui-dropdown-item v-if="tooltipOpt.stoppedStatus(scope.row.instanceStatus)">
                            <span
                                class="drop-down-operate-content"
                                @click="() => emitValue('open-close-restart', scope.row, 'open', tooltipOpt.stoppedStatus(scope.row.instanceStatus))"
                            >
                                {{$t('deviceList.content.label.operate.open')}}
                            </span>
                        </ui-dropdown-item>
                        <!-- 关机 -->
                        <ui-dropdown-item v-else>
                            <!-- 关机提示 -->
                            <ui-tooltip
                                placement="left"
                                :disabled="tooltipOpt.runningStatus(scope.row.instanceStatus)"
                                :content="$t('deviceList.tooltip.turnOff')"
                            >
                                <!-- 关机 -->
                                <span
                                    :class="[
                                        tooltipOpt.runningStatus(scope.row.instanceStatus) ? 'drop-down-operate-content' : 'drop-down-disabled'
                                    ]"
                                    @click="() => emitValue('open-close-restart', scope.row, 'close', tooltipOpt.runningStatus(scope.row.instanceStatus))"
                                >
                                    {{$t('deviceList.content.label.operate.close')}}
                                </span>
                            </ui-tooltip>
                        </ui-dropdown-item>
                        <!-- 重启 -->
                        <ui-dropdown-item>
                            <!-- 重启提示 -->
                            <ui-tooltip
                                placement="left"
                                :disabled="tooltipOpt.runningStatus(scope.row.instanceStatus)"
                                :content="$t('deviceList.tooltip.restart')"
                            >
                                <!-- 重启 -->
                                <span
                                    :class="[
                                        tooltipOpt.runningStatus(scope.row.instanceStatus) ? 'drop-down-operate-content' : 'drop-down-disabled'
                                    ]"
                                    @click="() => emitValue('open-close-restart', scope.row, 'restart', tooltipOpt.runningStatus(scope.row.instanceStatus))"
                                >
                                    {{$t('deviceList.content.label.operate.restart')}}
                                </span>
                            </ui-tooltip>
                        </ui-dropdown-item>
                        <!-- 重置实例状态 -->
                        <ui-dropdown-item>
                            <!-- 重置实例状态提示 -->
                            <ui-tooltip
                                placement="left"
                                :disabled="tooltipOpt.resetInstance(scope.row.instanceStatus)"
                                :content="$t('deviceList.tooltip.reset')"
                            >
                                <!-- 重置实例状态 -->
                                <span
                                    :class="[
                                        tooltipOpt.resetInstance(scope.row.instanceStatus) ? 'drop-down-operate-content' : 'drop-down-disabled'
                                    ]"
                                    @click="emitValue('reset-instance', scope.row.instanceId, scope.row.instanceName, tooltipOpt.resetInstance(scope.row.instanceStatus))"
                                >
                                    {{$t('deviceList.content.label.operate.reset')}}
                                </span>
                            </ui-tooltip>
                        </ui-dropdown-item>
                        <!-- 设备移除 -->
                        <ui-dropdown-item>
                            <!-- 设备移除提示 -->
                            <ui-tooltip
                                placement="left"
                                :disabled="tooltipOpt.deviceRemove(scope.row.instanceStatus)"
                                :content="$t('deviceList.tooltip.remove')"
                            >
                                <!-- 设备移除 -->
                                <span
                                    :class="[
                                        tooltipOpt.deviceRemove(scope.row.instanceStatus) ? 'drop-down-operate-content' : 'drop-down-disabled'
                                    ]"
                                    @click="emitValue('remove-recovery', scope.row, 'remove', tooltipOpt.deviceRemove(scope.row.instanceStatus))"
                                >
                                    {{$t('deviceList.content.label.operate.remove')}}
                                </span>
                            </ui-tooltip>
                        </ui-dropdown-item>
                        <!-- 回收实例 -->
                        <ui-dropdown-item>
                            <!-- 回收实例提示 -->
                            <ui-tooltip
                                placement="left"
                                :disabled="tooltipOpt.recoveryInstance(scope.row)"
                                :content="tooltipOpt.recoveryInstanceContent(scope.row.locked)"
                            >
                                <!-- 回收实例 -->
                                <span
                                    :class="[
                                        tooltipOpt.recoveryInstance(scope.row)? 'drop-down-operate-content' : 'drop-down-disabled'
                                    ]"
                                    @click="emitValue('remove-recovery', scope.row, 'recovery', tooltipOpt.recoveryInstance(scope.row))"
                                >
                                    {{$t('deviceList.content.label.operate.recovery')}}
                                </span>
                            </ui-tooltip>
                        </ui-dropdown-item>
                        <!-- 锁定 -->
                        <ui-dropdown-item>
                            <span
                                class="drop-down-operate-content"
                                v-if="scope.row.instanceStatus !== 'reinstalling' && scope.row.locked === 'unlocked' && !DeviceStaticData.lockUnlocked.includes(scope.row.instanceStatus)"
                                @click="emitValue('lock-nolock', scope.row.instanceName, scope.row.instanceId, 'lock')"
                            >
                                {{$t('deviceList.content.label.operate.lock')}}
                            </span>
                            <!-- 锁定提示 -->
                            <ui-tooltip
                                placement="left"
                                v-else-if="scope.row.locked === 'unlocked'"
                                :content="$t('deviceList.tooltip.lock1')"
                            >
                                <span class="drop-down-disabled">
                                    {{$t('deviceList.content.label.operate.lock')}}
                                </span>
                            </ui-tooltip>
                        </ui-dropdown-item>
                        <!-- 解锁 -->
                        <ui-dropdown-item>
                            <span
                                class="drop-down-operate-content"
                                v-if="scope.row.locked === 'locked' && !DeviceStaticData.lockUnlocked.includes(scope.row.instanceStatus)"
                                @click="emitValue('lock-nolock', scope.row.instanceName, scope.row.instanceId, 'deblocking')"
                            >
                                {{$t('deviceList.content.label.operate.unlock')}}
                            </span>
                            <!-- 解锁提示 -->
                            <ui-tooltip
                                placement="left"
                                v-else-if="scope.row.locked === 'locked'"
                                :content="$t('deviceList.tooltip.lock2')"
                            >
                                <span class="drop-down-disabled">
                                    {{$t('deviceList.content.label.operate.unlock')}}
                                </span>
                            </ui-tooltip>
                        </ui-dropdown-item>
                        <!-- 重置密码 -->
                        <ui-dropdown-item>
                            <!-- 重置密码提示 -->
                            <ui-tooltip
                                placement="left"
                                :disabled="tooltipOpt.resetPasswordStatus(scope.row.instanceStatus)"
                                :content="$t('deviceList.tooltip.resetPassword')"
                            >
                                <!-- 重置密码 -->
                                <span
                                    :class="[
                                        tooltipOpt.resetPasswordStatus(scope.row.instanceStatus) ? 'drop-down-operate-content' : 'drop-down-disabled'
                                    ]"
                                    @click="emitValue('reset-password', scope.row, 'resetPassword', tooltipOpt.resetPasswordStatus(scope.row.instanceStatus))"
                                >
                                    {{$t('deviceList.content.label.operate.resetPassword')}}
                                </span>
                            </ui-tooltip>
                        </ui-dropdown-item>
                        <!-- 重装系统 -->
                        <ui-dropdown-item>
                            <!-- 重装系统提示 -->
                            <ui-tooltip
                                placement="left"
                                :disabled="tooltipOpt.resetSystemStatus(scope.row.instanceStatus)"
                                :content="$t('deviceList.tooltip.resetSystem')"
                            >
                                <!-- 重装系统 -->
                                <span
                                    :class="[
                                        tooltipOpt.resetSystemStatus(scope.row.instanceStatus) ? 'drop-down-operate-content' : 'drop-down-disabled'
                                    ]"
                                    @click="emitValue('reset-system', scope.row, 'resetSystem', tooltipOpt.resetSystemStatus(scope.row.instanceStatus))"
                                >
                                    {{$t('deviceList.content.label.operate.resetSystem')}}
                                </span>
                            </ui-tooltip>
                        </ui-dropdown-item>
                        <!-- 删除设备 -->
                        <ui-dropdown-item>
                            <!-- 删除提示 -->
                            <ui-tooltip
                                placement="left"
                                :disabled="tooltipOpt.deviceDelete(scope.row.manageStatus)"
                                :content="$t('deviceDetail.tooltip.delete')"
                            >
                                <!-- 删除 -->
                                <span
                                    :class="[
                                        tooltipOpt.deviceDelete(scope.row.manageStatus) ? 'drop-down-operate-content' : 'drop-down-disabled'
                                    ]"
                                    @click="emitValue('up-down-delete', scope.row, 'delete', tooltipOpt.deviceDelete(scope.row.manageStatus))"
                                >
                                    {{$t('deviceList.content.label.operate.delete')}}
                                </span>
                            </ui-tooltip>
                        </ui-dropdown-item>
                    </ui-dropdown>
                </template>
            </ui-table-column>
        </ui-table>
    </ui-config-provider>
</template>
<script lang="ts" setup>
import SetEmpty from './setEmpty';
import TooltipOpt from './tooltipOpt';
import customOperate from '../custom/custom';
import FilterStyleOperate from './filterStyle';
import DeviceStaticData from 'staticData/device/index.ts';
import {uiLocale, RouterOperate} from 'utils/publicClass.ts';
import {tableClass, hasShowTooltip, setDiffClass, setTextClass, instanceStatusColor, managementStatusColor} from  'utils/index.ts';


interface PropsType {
    deviceList: any;
    tableMaxHeight: any;
    filterOperate: any;
    search: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});


// const path: string = instanceProxy.$defInfo.routerPath('deviceDetail');

const routerOperate = new RouterOperate();

const emitValue = defineEmits([
    'up-down-delete',
    'open-close-restart',
    'reset-instance',
    'remove-recovery',
    'lock-nolock',
    'import-device',
    'empty-click',
    'reset-system',
    'reset-password'
]);

const tooltipOpt = new TooltipOpt();
const filterStyleOt = new FilterStyleOperate(props);
new SetEmpty(props, emitValue);

</script>