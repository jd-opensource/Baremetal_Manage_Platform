<template>
    <main
        class="operate-management-detail-image-info"
        v-if="show !== 'basicInfo'"
    >
        <ui-button
            class="image-add-model btn-bottom"
            type="primary"
            @click="emitValue('add')"
        >
            +  {{$t('imageDetail.add.model')}}
        </ui-button>
        <!-- 表格-关联机型数据 -->
        <div class="operate-management-count image-detail-model">
            <!-- element-plus国际化 -->
            <ui-config-provider :locale="locale">
                <ui-table
                    ref="tableRef"
                    style="width: 100%;"
                    v-loading="loading"
                    border
                    :max-height="$filter.defaultWidth(heightValue)"
                    :class="tableClass(tableData, loading)"
                    :data="tableData"
                >
                    <!-- 机型名称 -->
                    <ui-table-column
                        align="center"
                        min-width="120"
                        fixed
                        :label="$t('imageDetail.table.label.modelName')"
                        :has-user-template="true"
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
                                    <span>
                                        {{$filter.emptyFilter(scope.row.name)}}
                                    </span>    
                                </div>
                            </ui-tooltip>
                        </template>
                    </ui-table-column>
                    <!-- 机型类型 -->
                    <ui-table-column
                        align="center"
                        min-width="120"
                        :label="$t('imageDetail.table.label.modelType')"
                        :has-user-template="true"
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
                                    <span>
                                        {{$filter.emptyFilter(scope.row.deviceSeriesName)}}
                                    </span>    
                                </div>
                            </ui-tooltip>
                        </template>
                    </ui-table-column>
                    <!-- 机房名称 -->
                    <ui-table-column
                        align="center"
                        min-width="120"
                        :label="$t('imageDetail.table.label.idcName')"
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
                    <!-- 机房英文名称 -->
                    <ui-table-column
                        align="center"
                        min-width="120"
                        :label="$t('imageDetail.table.label.nameEn')"
                        :has-user-template="true"
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
                                    <span>
                                        {{$filter.emptyFilter(scope.row.idcNameEn)}}
                                    </span>    
                                </div>
                            </ui-tooltip>
                        </template>
                    </ui-table-column>
                    <!-- 体系架构 -->
                    <ui-table-column
                        align="center"
                        min-width="120"
                        :label="$t('imageDetail.table.label.architecture')"
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
                    <!-- 机型规格 -->
                    <ui-table-column
                        align="center"
                        min-width="130"
                        :label="$t('imageDetail.table.label.modelSpecifications')"
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
                                    <span>
                                        {{$filter.emptyFilter(scope.row.deviceType)}}
                                    </span>    
                                </div>
                            </ui-tooltip>
                        </template>
                    </ui-table-column>
                    <!-- 描述 -->
                    <ui-table-column
                        align="center"
                        min-width="120"
                        :label="$t('imageDetail.table.label.desc')"
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
                                    @mouseenter="hasShowTooltip($event, scope.row[`descriptionTip${scope.$index}`], scope.row.description)"
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
                        min-width="110"
                        align="center"
                        fixed="right"
                        :has-user-template="true"
                        :label="$t('imageDetail.table.label.operate.title')"
                    >
                        <template #default="{scope}">
                            <!-- 移除 -->
                            <ui-tooltip
                                placement="left"
                                v-if="scope.row.instanceStatus.includes('creating')"
                                :content="$t('imageDetail.tooltip.title2')"
                            >
                                <span class="disabled-currency-style">
                                    {{$t('imageDetail.table.label.operate.remove')}}
                                </span>
                            </ui-tooltip>
                            <span
                                class="currency-style"
                                v-else
                                @click="emitValue('remove', scope.row)"
                            >
                                {{$t('imageDetail.table.label.operate.remove')}}
                            </span>
                        </template>
                    </ui-table-column>
                </ui-table>
            </ui-config-provider>
        </div>
    </main>
</template>
<script lang="ts" setup>
import {languageSwitch, tableClass, hasShowTooltip} from 'utils/index.ts';
import SetEmpty from './setEmpty';

interface PropsType {
    loading: boolean;
    show: string;
    tableData: any;
    heightValue: any;
    // imgSrc?: string;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});
const locale = languageSwitch();

const emitValue = defineEmits(['add', 'remove', 'empty-click']);
new SetEmpty(props, emitValue)

</script>