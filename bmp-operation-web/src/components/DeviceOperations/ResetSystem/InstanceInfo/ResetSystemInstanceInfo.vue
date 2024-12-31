<template>
    <div class="instance-count">
        <p class="title">
            {{$t('resetSystem.header.instanceInfo')}}
        </p>
        <div class="currency-count-table">
            <ui-config-provider :locale="uiLocale.locale">
                <ui-table
                    style="width: 100%"
                    border
                    :class="[tableClass(formRulesOperate.reactiveArr.tableData, false)]"
                    :data="formRulesOperate.reactiveArr.tableData"
                >
                    <!-- 实例名称 -->
                    <ui-table-column
                        prop="instanceName"
                        align="center"
                        min-width="120"
                        :label="$t('resetSystem.table.label.instanceName')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <ui-tooltip
                                placement="bottom"
                                v-model="scope.row.showTooltip"
                                :disabled="!scope.row.showTooltip"
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
                                    @mouseenter="hasShowTooltip($event, scope.row, scope.row.instanceName)"
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
                        min-width="120"
                        align="center"
                        :label="$t('resetSystem.table.label.instanceID')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <ui-tooltip
                                placement="bottom"
                                v-model="scope.row.showTooltip"
                                :disabled="!scope.row.showTooltip"
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
                                    @mouseenter="hasShowTooltip($event, scope.row, scope.row.instanceId)"
                                >
                                    <span>
                                        {{$filter.emptyFilter(scope.row.instanceId)}}
                                    </span>    
                                </div>
                            </ui-tooltip>
                        </template>
                    </ui-table-column>
                    <!-- 操作系统 -->
                    <ui-table-column
                        prop="imageName"
                        align="center"
                        min-width="120"
                        :label="$t('resetSystem.table.label.operatingSystem')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <ui-tooltip
                                placement="bottom"
                                v-model="scope.row.showTooltip"
                                :disabled="!scope.row.showTooltip"
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
                                    @mouseenter="hasShowTooltip($event, scope.row, scope.row.imageName)"
                                >
                                    <span>
                                        {{$filter.emptyFilter(scope.row.imageName)}}
                                    </span>    
                                </div>
                            </ui-tooltip>
                        </template>
                    </ui-table-column>
                    <!-- 创建时间 -->
                    <ui-table-column
                        prop="createdTime"
                        align="center"
                        min-width="100"
                        :label="$t('resetSystem.table.label.createdTime')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.createdTime)}}
                        </template>
                    </ui-table-column>
                </ui-table>
            </ui-config-provider>
        </div>
    </div>
</template>
<script lang="ts" setup>
import {uiLocale} from 'utils/publicClass.ts';
import {tableClass, hasShowTooltip} from 'utils/index.ts';

type PropsType = {
    formRulesOperate: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
