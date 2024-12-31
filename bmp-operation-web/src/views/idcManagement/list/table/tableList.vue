<template>
    <!-- 表格内容 -->
    <article class="operate-management-count no-top idc-list">
        <ui-config-provider :locale="uiLocale.locale">
            <ui-table
                style="width: 100%;"
                v-loading="idcList.isLoading.value"
                border
                :max-height="tableScrollOperate.tableMaxHeight.value"
                :class="tableClass(idcList.reactiveArr.tableData, idcList.isLoading.value)"
                :data="idcList.reactiveArr.tableData"
            >
                <template v-for="(item, ind) in idcList.templateData">
                    <ui-table-column
                        v-if="customOperate?.reactiveArr?.hasCustomInfo[item.ifSelect]?.selected"
                        :key="ind"
                        :prop="item.prop"
                        :min-width="item.minWidth"
                        :align="item.align"
                        :fixed="item.fixed"
                        :label="item.label"
                        :has-user-template="item.hasUserTemplate"
                    >
                        <template #default="{scope}">
                            <ui-tooltip
                                placement="bottom"
                                v-if="item.hasShowTooltip"
                                v-model="scope.row[`${item.prop}Tip${scope.$index}`].showTooltip"
                                :disabled="!scope.row[`${item.prop}Tip${scope.$index}`].showTooltip"
                            >
                                <template #content>
                                    <div class="set-tooltip-width">
                                        <span>
                                            {{scope.row[item.prop]}}
                                        </span>
                                    </div>
                                </template>
                                <div
                                    class="more-text-ellipsis"
                                    @mouseenter="hasShowTooltip($event, scope.row[`${item.prop}Tip${scope.$index}`], scope.row[item.prop])"
                                >
                                    <!-- 是否跳转详情 -->
                                    <span
                                        :class="setTextClass(item.hasClick)"
                                        @click="item.hasClick ? routerOperate?.jumpRouter({idcId: scope.row.idcId}) : ''"
                                    >
                                        {{$filter.emptyFilter(scope.row[item.prop])}}
                                    </span>
                                </div>
                            </ui-tooltip>
                            <span
                                v-else
                                :class="setTextClass(item.hasClick)"
                                @click="() => item.hasClick ? emitValue('edit', scope.row) : null"
                            >
                                {{Object.is(item.prop, 'operate') ? $t('idcList.content.operate.edit') : $filter.emptyFilter(scope.row[item.prop])}}
                            </span>
                        </template>
                    </ui-table-column>
                </template>
            </ui-table>
        </ui-config-provider>
    </article>
</template>

<script lang="ts" setup>
import {RouterOperate, uiLocale} from 'utils/publicClass.ts';
import {tableClass, setTextClass, hasShowTooltip} from 'utils/index.ts'; // 表格class类
import SetEmpty from './setEmpty';
import customOperate from '../customList/custom';

interface PropsType {
    idcList: {
        isLoading: {
            value: boolean;
        },
        reactiveArr: {
            tableData: any;
        },
        templateData: {
            prop: string;
            align: string;
            minWidth: string;
            ifSelect: string;
            fixed: boolean | string;
            label: string;
            hasUserTemplate: boolean;
            hasClick: boolean;
            hasShowTooltip: boolean;
        }[];
    };
    tableScrollOperate: {
        tableMaxHeight: {
            value: number;
        }
    };
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

const routerOperate = new RouterOperate();

// defineEmits API 来替代 emits
const emitValue = defineEmits(['edit']);

new SetEmpty(props)
</script>
<style lang="scss">

</style>
