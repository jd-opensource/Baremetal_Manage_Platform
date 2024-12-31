<template>
    <div class="currency-count-table">
        <ui-config-provider :locale="uiLocale.locale">
            <ui-table
                ref="tableRef"
                style="width: 555px"
                max-height="200px"
                v-loading="formRulesOperate.isLoading.value"
                :class="[tableClass(formRulesOperate.reactiveArr.keyPariData, formRulesOperate.isLoading.value)]"
                :data="formRulesOperate.reactiveArr.keyPariData"
                @get-table-ref="tableOpt.getTableRef"
                @cell-click="tableOpt.cellClick"
                @selection-change="tableOpt.handleSelectionChange"
            >
                <ui-table-column
                    type="selection"
                    align="center"
                />
                <!-- 密钥名称 -->
                <ui-table-column
                    prop="name"
                    align="center"
                    min-width="120"
                    :label="$t('resetSystem.table.label.sshName')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <ui-tooltip
                            placement="bottom"
                            v-model="tooltip.name.showTooltip"
                            :disabled="!tooltip.name.showTooltip"
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
                                @mouseenter="hasShowTooltip($event, tooltip.name, scope.row.name)"
                            >
                                <span>
                                    {{$filter.emptyFilter(scope.row.name)}}
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
                    :label="$t('resetSystem.table.label.createdTime')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        {{$filter.emptyFilter(scope.row.createdTime)}}
                    </template>
                </ui-table-column>
            </ui-table>
            <p
                class="error-color-tip"
                v-if="rulesCheck.hasKeyFlag.value"
            >
                {{$t('resetSystem.emptyCheck.sshKey')}}
            </p>
        </ui-config-provider>
    </div>
</template>
<script lang="ts" setup>
import {uiLocale} from 'utils/publicClass.ts';
import {tableClass, hasShowTooltip} from 'utils/index.ts';
import DeviceStaticData from 'staticData/device/index.ts';


import tableOperate from './tableOperate';


type PropsType = {
    formRulesOperate: any;
    rulesCheck: any;
    // tableOpt: any;
}
/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props = withDefaults(defineProps<PropsType>(), {});
const tableOpt = tableOperate(props.formRulesOperate.ruleForm, props.rulesCheck);

const tooltip = reactive(DeviceStaticData.resetSysTemKeyTooltip)
</script>