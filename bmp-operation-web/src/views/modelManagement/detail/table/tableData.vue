<template>
    <!-- element-plus国际化 -->
    <ui-config-provider :locale="locale">
        <ui-table
            ref="tableRef"
            style="width: 100%;"
            v-loading="tableDetail.isLoading.value"
            border
            :max-height="$filter.defaultWidth(tableS.tableMaxHeight.value)"
            :class="tableClass(tableDetail.reactiveArr.tableData, tableDetail.isLoading.value)"
            :data="tableDetail.reactiveArr.tableData"
            @filter-change="filterOperate.filterChange"
            @get-table-ref="filterOperate.getTableRef"
        >
            <!-- 镜像名称 -->
            <ui-table-column
                align="center"
                min-width="140"
                fixed
                :label="$t('modelDetail.tabs.relationImage.content.label.imageName')"
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
                min-width="130"
                key="source"
                column-key="source"
                filter-placement="bottom-end"
                :class-name="filterStyleOperate.filterStatus['source'] ? 'common-status' : 'def-type-status1'"
                :filters="filterOperate.ossStore?.source"
                :filter-method="filterOperate.imageTypeFilter"
                :label="$t('modelDetail.tabs.relationImage.content.label.imageType')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span style="marginRight: 19px">
                        {{$filter.emptyFilter(scope.row.sourceName)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 体系架构 -->
            <ui-table-column
                align="center"
                min-width="90"
                :has-user-template="true"
                :label="$t('modelDetail.tabs.relationImage.content.label.architecture')"
            >
                <template #default="{scope}">
                    {{$filter.emptyFilter(scope.row.architecture)}}
                </template>
            </ui-table-column>
            <!-- 操作系统平台 -->
            <ui-table-column
                align="center"
                filter-placement="bottom-end"
                column-key="osType"
                key="osType"
                min-width="130"
                :class-name="filterStyleOperate.filterStatus['osType'] ? 'common-status' : 'def-type-status2'"
                :has-user-template="true"
                :filters="filterOperate.ossStore.osType"
                :filter-method="filterOperate.operateSysFilter"
                :label="$t('modelDetail.tabs.relationImage.content.label.operateSysPlatform')"
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
                            <span style="marginRight: 23px">
                                {{$filter.emptyFilter(scope.row.osType)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 操作系统版本 -->
            <ui-table-column
                align="center"
                min-width="100"
                :label="$t('modelDetail.tabs.relationImage.content.label.operateSysVersion')"
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
                                {{$filter.emptyFilter(scope.row.osVersion)}}
                            </span>    
                        </div>
                    </ui-tooltip>
                </template>
            </ui-table-column>
            <!-- 分区模块 -->
            <ui-table-column
                min-width="180"
                align="center"
                prop="systemPartition"
                :has-user-template="true"
                :label="$t('modelDetail.tabs.relationImage.content.label.partitionModule')"
            >
                <template #default="{scope}">
                    <div v-if="scope.row?.newSystemPartition?.length">
                        <p
                            v-for="(item, index) in scope.row.newSystemPartition"
                            :key="index"
                        >
                            <span>{{$filter.emptyFilter(item.point)}}：</span>
                            <span>{{tableDetail.systemPartitionInfo?.setSize(item.size)}}，</span>
                            <span>{{$filter.emptyFilter(item.format)}}</span>
                        </p>
                    </div>
                    <div v-else>{{$filter.emptyFilter()}}</div>
                </template>
            </ui-table-column>
            <!-- 添加时间 -->
            <ui-table-column
                align="center"
                min-width="150"
                :label="$t('modelDetail.tabs.relationImage.content.label.addTime')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    {{$filter.emptyFilter(scope.row.createdTime)}}
                </template>
            </ui-table-column>
            <!-- 操作 -->
            <ui-table-column
                min-width="120"
                align="center"
                fixed="right"
                :has-user-template="true"
                :label="$t('modelDetail.tabs.relationImage.content.label.operate.name')"
            >
                <template #default="{scope}">
                    <!-- 删除 -->
                    <span
                        class="currency-style"
                        @click="emitValue('delete', scope.row)"
                    >
                        {{$t('modelDetail.tabs.relationImage.content.label.operate.delete')}}
                    </span>
                </template>
            </ui-table-column>
        </ui-table>
    </ui-config-provider>
</template>
<script lang="ts" setup>
import {tableClass, languageSwitch, hasShowTooltip} from 'utils/index.ts';
import FilterStyleOperate from './filterStyle';
import SetEmpty from './setEmpty';
// ui库内置的国际化
const locale = languageSwitch();

interface PropsType {
    tableDetail: any;
    tableS: any;
    filterOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});
// defineEmits API 来替代 emits
const emitValue = defineEmits(['add-image', 'delete', 'empty-click']);

const filterStyleOperate = new FilterStyleOperate(props);
new SetEmpty(props, emitValue)
</script>