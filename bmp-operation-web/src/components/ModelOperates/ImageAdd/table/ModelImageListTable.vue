<template>
    <!-- element-plus 国际化 -->
    <ui-config-provider :locale="imageList.locale">
        <!-- 添加机型数据 -->
        <ui-table
            border
            style="width: 100%"
            max-height="260"
            v-loading="imageList.tableLoading.value"
            :data="imageList.reactiveArr.tableData"
            :class="tableClass(imageList.reactiveArr.tableData, imageList.tableLoading.value)"
            :row-key="imageList.getRowKeys"
            @selection-change="imageList.handleSelectionChange"
            @select-all="imageList.selectAllChange"
            @select="imageList.selectChange"
            @filter-change="filterOperate.filterChange"
            @get-table-ref="imageList.getTableRef"
            @row-click="imageList.rowClick"
        >
            <!-- 复选框 -->
            <ui-table-column
                type="selection"
                align="center"
            />
            <!-- 镜像名称 -->
            <ui-table-column
                align="center"
                min-width="120"
                :label="$t('addImage.content.label.imageName')"
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
            <!-- 镜像类型 -->
            <ui-table-column
                align="center"
                min-width="125"
                key="source"
                column-key="source"
                filter-placement="bottom-end"
                :class-name="filterStyleOperate.filterStatus['sources'] ? 'common-status' : 'def-type-status3'"
                :filters="filterOperate.ossStore?.source"
                :filter-method="filterOperate.imageTypeFilter"
                :label="$t('addImage.content.label.imageType')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span style="marginRight: 18px">
                        {{$filter.emptyFilter(scope?.row?.sourceName)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 体系架构 -->
            <ui-table-column
                align="center"
                min-width="110"
                :has-user-template="true"
                :label="$t('addImage.content.label.architecture')"
            >
                <template #default="{scope}">
                    <span>
                        {{$filter.emptyFilter(scope?.row?.architecture)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 操作系统平台 -->
            <ui-table-column
                align="center"
                key="osType"
                min-width="120"
                column-key="osType"
                filter-placement="bottom-end"
                :class-name="filterStyleOperate.filterStatus['osTypes'] ? 'common-status' : 'def-type-status4'"
                :filters="filterOperate.ossStore.osType"
                :filter-method="filterOperate.filterOsType"
                :has-user-template="true"
                :label="$t('addImage.content.label.operateSysPlatform')"
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
                            @mouseenter="hasShowTooltip($event, scope.row[`osTypeTip${scope.$index}`], scope.row.osType)"
                        >
                            <span style="marginRight: 22px">
                                {{$filter.emptyFilter(scope?.row?.osType)}}
                            </span> 
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 操作系统版本 -->
            <ui-table-column
                align="center"
                min-width="100"
                :label="$t('addImage.content.label.operateSysVersion')"
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
                            @mouseenter="hasShowTooltip($event, scope.row[`osVersionTip${scope.$index}`], scope.row.osVersion)"
                        >
                            <span>
                                {{$filter.emptyFilter(scope?.row?.osVersion)}}
                            </span>
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 系统盘分区模板 -->
            <ui-table-column
                prop="systemPartition"
                align="center"
                min-width="180"
                :has-user-template="true"
                :label="$t('addImage.content.label.sysPartitionModule')"
            >
                <template #default="{scope}">
                    <div v-if="scope.row?.newSystemPartition?.length">
                        <p
                            v-for="(item, index) in scope.row.newSystemPartition"
                            :key="index"
                        >
                            <span>{{$filter.emptyFilter(item.point)}}：</span>
                            <span>{{imageList.systemPartitionInfo?.setSize(item.size)}}，</span>
                            <span>{{$filter.emptyFilter(item.format)}}</span>
                        </p>
                    </div>
                    <div v-else>{{$filter.emptyFilter()}}</div>
                </template>
            </ui-table-column>
        </ui-table>
    </ui-config-provider>
</template>
<script lang="ts" setup>
import {tableClass, hasShowTooltip} from 'utils/index.ts';
import FilterStyleOperate from './filterStyle';

/**
 * props 类
*/
interface PropsType {
    imageList: any;
    filterOperate: any;
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props = withDefaults(defineProps<PropsType>(), {});

const filterStyleOperate = new FilterStyleOperate(props);
</script>