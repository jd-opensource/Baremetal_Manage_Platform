<template>
    <!-- element-plus国际化 -->
    <ui-config-provider :locale="uiLocale.locale">
        <ui-table
            ref="tableRef"
            style="width: 100%"
            border
            v-loading="userList.isLoading.value"
            :max-height="tableMaxHeight"
            :class="tableClass(userList.reactiveArr.tableData, userList.isLoading.value)"
            :data="userList.reactiveArr.tableData"
            @filter-change="filterOperates.filterChange"
            @get-table-ref="filterOperates.getTableRef"
        >
            <template v-for="(item, ind) in userList.templateData">
                <ui-table-column
                    align="center"
                    v-if="customOperate?.reactiveArr?.hasCustomInfo[item.customInfo]?.selected && !item.columnKey"
                    :prop="item.prop"
                    :min-width="item.minWidth"
                    :fixed="item.fixed"
                    :label="item.label"
                    :filter-placement="item.filterPlacement"
                    :key="ind"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <div v-if="item.toolTip">
                            <ui-tooltip
                                placement="bottom"
                                v-model="scope.row[`${item.toolTip}${scope.$index}`].showTooltip"
                                :disabled="!scope.row[`${item.toolTip}${scope.$index}`].showTooltip"
                            >
                                <template #content>
                                    <div class="set-tooltip-width">
                                        {{scope.row[item.prop]}}
                                    </div>
                                </template>
                                <div
                                    class="more-text-ellipsis"
                                    @mouseenter="hasShowTooltip($event, scope.row[`${item.toolTip}${scope.$index}`], scope.row[item.prop])"
                                >
                                    <span :class="setTextClass(false)">
                                        {{$filter.emptyFilter(scope.row[item.prop])}}
                                    </span>
                                </div>
                            </ui-tooltip>
                        </div>
                        <span
                            v-else-if="!Object.is(item.prop, 'operate')"
                            :style="item.marginRight ? {marginRight: item.marginRight} : {}"
                            :class="setTextClass(item.setTextClass)"
                        >
                            {{$filter.emptyFilter(scope.row[item.prop])}}
                        </span>
                        <p v-else>
                            <!-- 编辑 -->
                            <span
                                :class="setTextClass(true, 'currency-style-right')"
                                @click="() => emitValue('edit-user', scope.row, scope.row.prefix)"
                            >
                                {{$t('userList.label.operate.edit')}}
                            </span>
                            <!-- 删除提示 -->
                            <ui-tooltip
                                placement="bottom"
                                v-if="!scope.row.instanceCount"
                                :disabled="!Object.is(scope.row.roleId, $defInfo.purview('admin'))"
                                :content="$t('userList.label.operate.tooltip.delete')"
                            >
                                <!-- 删除 -->
                                <span
                                    :class="[Object.is(scope.row.roleId, $defInfo.purview('admin')) ? 'disabled-currency-style' : 'currency-style']"
                                    @click="() => {
                                        if (Object.is(scope.row.roleId, $defInfo.purview('admin'))) return;
                                        emitValue('delete-user', scope.row)
                                    }"
                                >
                                    {{$t('userList.label.operate.delete')}}
                                </span>
                            </ui-tooltip>
                            <ui-tooltip
                                placement="bottom"
                                v-else
                                :content="$t('userList.label.operate.tooltip.delete2')"
                            >
                                <!-- 删除 -->
                                <span class="disabled-currency-style">
                                    {{$t('userList.label.operate.delete')}}
                                </span>
                            </ui-tooltip>
                        </p>
                    </template>
                </ui-table-column>
                <ui-table-column
                    align="center"
                    v-else
                    :prop="item.prop"
                    :min-width="item.minWidth"
                    :fixed="item.fixed"
                    :label="item.label"
                    :filter-placement="item.filterPlacement"
                    :key="item.key"
                    :column-key="item.columnKey"
                    :filters="filterOperates.ossStore[item.filters]"
                    :filter-method="filterOperates[item.filterMethod]"
                    :class-name="filterStyleOperate.filterStatus[item.filterStatus] ? 'common-status' : `def-type-status${item.filterClass}`"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <span
                            :style="item.marginRight ? {marginRight: item.marginRight} : {}"
                            :class="setTextClass(item.setTextClass)"
                        >
                            {{$filter.emptyFilter(scope.row[item.prop])}}
                        </span>
                    </template>
                </ui-table-column>
            </template>
        </ui-table>
    </ui-config-provider>
</template>

<script lang="ts" setup>
import {uiLocale} from 'utils/publicClass.ts';
import {tableClass, setTextClass, hasShowTooltip} from 'utils/index.ts';
import customOperate from '../custom/custom';
import SetEmpty from './setEmpty';
import FilterStyleOperate from './filterStyle';
// import UserStaticData from 'staticData/user/index.ts';

interface PropsType {
    userList: any;
    tableMaxHeight: number;
    filterOperates: {
        [x: string]: unknown;
        filterChange: any;
        getTableRef: any;
        ossStore: {
            [x: string]: unknown;
        };
    };
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});
// const toolTip = reactive(UserStaticData.userTooltip);

const filterStyleOperate = new FilterStyleOperate(props);
// defineEmits API 来替代 emits
const emitValue = defineEmits(['edit-user', 'delete-user', 'empty-click']);
new SetEmpty(props, emitValue);
</script>