<template>
    <!-- element-plus 国际化 -->
    <ui-config-provider :locale="uiLocale.locale">
        <ui-table
            ref="tableRef"
            v-loading="imageList?.isLoading?.value"
            border
            :class="tableClass(imageList.reactiveArr.tableData, imageList.isLoading.value)"
            :data="imageList.reactiveArr.tableData"
            :max-height="tableMaxHeight"
            @filter-change="filter.filterChange"
            @get-table-ref="filter.getTableRef"
        >
            <!-- 镜像名称 -->
            <ui-table-column
                v-if="customOperate?.reactiveArr?.hasCustomInfo['imageName']?.selected"
                fixed
                min-width="160"
                align="center"
                :label="$t('imageList.content.label.imageName')"
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
                            <span
                                :class="setTextClass(true)"
                                @click="routerOperate.jumpRouter({imageName: scope.row.imageName, imageId: scope.row.imageId, architecture: scope.row.architecture})"
                            >
                                {{$filter.emptyFilter(scope.row.imageName)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 镜像id -->
            <ui-table-column
                align="center"
                min-width="240"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['imageId']?.selected"
                :has-user-template="true"
                :label="$t('imageList.content.label.imageId')"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`imageIdTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`imageIdTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.imageId}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`imageIdTip${scope.$index}`], scope.row.imageId)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.imageId)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 镜像类型 -->
            <ui-table-column
                align="center"
                min-width="140"
                key="source"
                column-key="source"
                filter-placement="bottom-end"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['source']?.selected"
                :class-name="filterStyleOperate.filterStatus['source'] ? 'common-status' : 'def-type-status5'"
                :has-user-template="true"
                :filters="filter?.ossStore?.source"
                :filter-method="filter?.imageTypeFilter"
                :label="$t('imageList.content.label.imageType')"
            >
                <template #default="{scope}">
                    <span style="marginRight: 20px;">
                        {{$filter.emptyFilter(scope.row.sourceName)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 体系架构 -->
            <ui-table-column
                key="architecture"
                align="center"
                column-key="architecture"
                min-width="145"
                filter-placement="bottom-end"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['architecture']?.selected"
                :class-name="filterStyleOperate.filterStatus['architecture'] ? 'common-status' : 'def-type-status6'"
                :has-user-template="true"
                :filters="filter?.ossStore?.architecture"
                :filter-method="filter?.filterArchitecture"
                :label="$t('imageList.content.label.architecture')"
            >
                <template #default="{scope}">
                    <span style="marginRight: 22px;">
                        {{$filter.emptyFilter(scope.row.architecture)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 操作系统平台 -->
            <ui-table-column
                min-width="130"
                align="center"
                key="osType"
                column-key="osType"
                filter-placement="bottom-end"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['osType']?.selected"
                :class-name="filterStyleOperate.filterStatus['osType'] ? 'common-status' : 'def-type-status7'"
                :has-user-template="true"
                :filters="filter?.ossStore?.osType"
                :filter-method="filter?.operateSysFilter"
                :label="$t('imageList.content.label.operateSysPlatform')"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`osTypeTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`osTypeTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.osType}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`osTypeTip${scope.$index}`], scope.row.osType, 1.5)"
                        >
                            <span style="marginRight: 22px;">
                                {{$filter.emptyFilter(scope.row.osType)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 操作系统版本 -->
            <ui-table-column
                min-width="115"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['osVersion']?.selected"
                :label="$t('imageList.content.label.operateSysVersion')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <ui-tooltip
                        placement="bottom"
                        v-model="scope.row[`osVersionTip${scope.$index}`].showTooltip"
                        :disabled="!scope.row[`osVersionTip${scope.$index}`].showTooltip"
                    >
                        <template #content>
                            <div class="set-tooltip-width">
                                <span>
                                    {{scope.row.osVersion}}
                                </span>
                            </div>
                        </template>
                        <div
                            class="more-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, scope.row[`osVersionTip${scope.$index}`], scope.row.osVersion, 1.5)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.osVersion)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>

            <!-- 镜像格式 -->
            <ui-table-column
                min-width="100"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['format']?.selected"
                :label="$t('imageList.content.label.format')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    
                    <span>
                        {{$filter.emptyFilter(scope.row.format)}}
                    </span>  
                </template>
            </ui-table-column>
            <!-- 引导模式 -->
            <ui-table-column
                min-width="140"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['bootMode']?.selected"
                :label="$t('imageList.content.label.bootMode')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    
                    <span>
                        {{$filter.emptyFilter(scope.row.bootMode)}}
                    </span>  
                </template>
            </ui-table-column>
            <!-- 系统盘分区模板 -->
            <ui-table-column
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['systemPartition']?.selected"
                :min-width="setDiffClass('240', '160')"
                :label="$t('imageList.content.label.sysPartitionModule')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <div v-if="scope.row?.newSystemPartition?.length">
                        <p
                            v-for="(item, index) in scope.row.newSystemPartition"
                            :key="index"
                        >
                            <span>{{item.point}}：</span>
                            <span>{{imageList.systemPartitionInfo?.setSize(item.size)}}，</span>
                            <span>{{item.format}}</span>
                        </p>
                    </div>
                    <div v-else>{{$filter.emptyFilter()}}</div>
                </template>
            </ui-table-column>
            <!-- 镜像描述 -->
            <ui-table-column
                min-width="130"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['description']?.selected"
                :label="$t('imageList.content.label.imageDesc')"
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
                            @mouseenter="hasShowTooltip($event, scope.row[`descriptionTip${scope.$index}`], scope.row.description, 1.5)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope.row.description)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 已关联机型数 -->
            <ui-table-column
                min-width="110"
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['deviceTypeNum']?.selected"
                :label="$t('imageList.content.label.numberOfAssociatedModels')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span
                        v-if="scope.row.deviceTypeNum > 0"
                        :class=setTextClass(true)
                        @click="routerOperate.jumpRouter({imageName: scope.row.imageName, imageId: scope.row.imageId, architecture: scope.row.architecture, deviceTypeNum: scope.row.deviceTypeNum})"
                    >
                        {{scope.row.deviceTypeNum}}
                    </span>
                    <span v-else>--</span>
                </template>
            </ui-table-column>
            <!-- 创建时间 -->
            <ui-table-column
                align="center"
                min-width="150"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['createdTime']?.selected"
                :label="$t('imageList.content.label.createTime')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span>
                        {{$filter.emptyFilter(scope.row.createdTime)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 操作 -->
            <ui-table-column
                prop="operate"
                min-width="80"
                align="center"
                fixed="right"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['operate']?.selected"
                :has-user-template="true"
                :label="$t('imageList.content.label.operate.name')"
            >
                <template #default="{scope}">
                    <!-- 无法删除提示 -->
                    <ui-tooltip
                        placement="bottom"
                        v-if="scope.row.deviceTypeNum > 0"
                        :content="$t('imageList.tooltip.title')"
                    >
                        <span class="disabled-currency-style">
                            {{$t('imageList.content.label.operate.delete')}}
                        </span>
                    </ui-tooltip>
                    <!-- 预置镜像不支持删除 -->
                    <ui-tooltip
                        placement="bottom"
                        v-else-if="scope.row.source === 'common'"
                        :content="$t('imageList.tooltip.image')"
                    >
                        <span class="disabled-currency-style">
                            {{$t('imageList.content.label.operate.delete')}}
                        </span>
                    </ui-tooltip>
                    <!-- 删除 -->
                    <span
                        class="currency-style"
                        v-else
                        @click="deleteImage(scope.row)"
                    >
                        {{$t('imageList.content.label.operate.delete')}}
                    </span>
                </template>
            </ui-table-column>
        </ui-table>
    </ui-config-provider>
</template>
<script lang="ts" setup>
import {uiLocale, RouterOperate, CurrentInstance} from 'utils/publicClass.ts';
import {tableClass, setTextClass, hasShowTooltip, setDiffClass} from 'utils/index.ts';
import customOperate from '../custom/custom';
import FilterStyleOperate from './filterStyle';
import SetEmpty from './setEmpty';

interface PropsType {
    imageList: any;
    search: any;
    // tableScrollOperate: any;
    tableMaxHeight: any;
    importImage: any;
    deleteImage: any;
    filter: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});


const emitValue = defineEmits(['empty-click']);

const instanceProxy = new CurrentInstance().proxy;

const routerOperate = new RouterOperate(instanceProxy.$defInfo.routerPath('imageDetail'));

const filterStyleOperate = new FilterStyleOperate(props);

new SetEmpty(props, emitValue);

</script>