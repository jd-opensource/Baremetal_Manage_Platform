<template>
    <!-- 表格数据 -->
    <ui-table
        border
        style="width: 100%"
        :data="tableData"
        :class="[tableClass(tableData, false)]"
    >
        <!-- sn -->
        <ui-table-column
            prop="sn"
            align="center"
            min-width="120"
            :label="$t('publicTable.label.sn')"
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
                                {{scope.row.sn}}
                            </span>
                        </div>
                    </template>
                    <div
                        class="more-text-ellipsis"
                        @mouseenter="hasShowTooltip($event, scope.row, scope.row.sn)"
                    >
                        <span>
                            {{$filter.emptyFilter(scope.row.sn)}}
                        </span>    
                    </div>
                </ui-tooltip>
            </template>
        </ui-table-column>
        <!-- 实例名称 -->
        <ui-table-column
            prop="instanceName"
            align="center"
            min-width="120"
            :label="$t('publicTable.label.instanceName')"
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
            :label="$t('publicTable.label.instanceId')"
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
        <!-- 实例属主 -->
        <ui-table-column
            prop="userName"
            align="center"
            min-width="120"
            :label="$t('publicTable.label.instanceOwner')"
            :has-user-template="true"
        >
            <template #default="{scope}">
                {{$filter.emptyFilter(scope.row.userName)}}
            </template>
        </ui-table-column>
        <!-- 带外IP -->
        <ui-table-column
            prop="iloIp"
            align="center"
            min-width="100"
            :label="$t('publicTable.label.outOfBandIP')"
            :has-user-template="true"
        >
            <template #default="{scope}">
                {{$filter.emptyFilter(scope.row.iloIp)}}
            </template>
        </ui-table-column>
    </ui-table>
</template>

<script lang="ts" setup>
import {tableClass, hasShowTooltip} from 'utils/index.ts';

interface PropsType {
    tableData: any
}
// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>

