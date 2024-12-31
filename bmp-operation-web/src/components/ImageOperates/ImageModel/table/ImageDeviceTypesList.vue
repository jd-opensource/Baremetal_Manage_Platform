<template>
<!-- element-plus 国际化 -->
<ui-config-provider :locale="imageDeviceTypes.locale">
    <!-- 添加机型数据 -->
    <ui-table
        border
        style="width: 100%"
        v-loading="imageDeviceTypes.tableLoading.value"
        :height="imageDeviceTypes.reactiveArr.tableData.length ? '260' : 'auto'"
        :class="tableClass(imageDeviceTypes.reactiveArr.tableData, imageDeviceTypes.tableLoading.value)"
        :data="imageDeviceTypes.reactiveArr.tableData"
        :row-key="imageDeviceTypes.getRowKeys"
        @selection-change="imageDeviceTypes.handleSelectionChange"
        @select-all="imageDeviceTypes.selectAllChange"
        @select="imageDeviceTypes.selectChange"
        @get-table-ref="imageDeviceTypes.getTableRef"
        @row-click="imageDeviceTypes.rowClick"
    >
        <!-- 复选框 -->
        <ui-table-column
            type="selection"
            align="center"
        />
        <!-- 机型名称 -->
        <ui-table-column
            min-width="110"
            align="center"
            :label="$t('imageAddModel.label.modelName')"
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
            :label="$t('imageAddModel.label.modelType')"
            :has-user-template="true"
        >
            <template #default="{scope}">
                <span>
                    {{$filter.emptyFilter(scope.row.deviceSeriesName)}}
                </span>
            </template>
        </ui-table-column>
        <!-- 机房名称 -->
        <ui-table-column
            align="center"
            min-width="100"
            :label="$t('imageAddModel.label.idcName')"
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
            min-width="100"
            align="center"
            :label="$t('imageAddModel.label.nameEn')"
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
            min-width="110"
            align="center"
            :has-user-template="true"
            :label="$t('imageAddModel.label.architecture')"
        >
            <template #default="{scope}">
                <span>{{$filter.emptyFilter(scope.row.architecture)}}</span>
            </template>
        </ui-table-column>
        <!-- 机型规格 -->
        <ui-table-column
            align="center"
            :min-width="setDiffClass('125', '100')"
            :has-user-template="true"
            :label="$t('imageAddModel.label.modelSpecifications')"
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
            min-width="100"
            :label="$t('imageAddModel.label.desc')"
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
    </ui-table>
</ui-config-provider>
</template>
<script lang="ts" setup>
import {tableClass, setDiffClass, hasShowTooltip} from 'utils/index.ts';

/**
 * props 类
*/
interface PropsType {
    imageDeviceTypes: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>