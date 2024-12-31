<template>
    <ui-table
        border
        style="width: 100%"
        :max-height="180"
        :class="[tableClass(selectArr, false)]"
        :data="selectArr"
    >
        <!-- 实例名称 -->
        <ui-table-column
            prop="instanceName"
            align="center"
            min-width="120"
            :label="$t('batchOperate.openCloseRestart.label.instanceName')"
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
            min-width="130"
            align="center"
            :label="$t('batchOperate.openCloseRestart.label.instanceId')"
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
        <!-- 带外IP -->
        <ui-table-column
            prop="iloIp"
            align="center"
            min-width="100"
            :label="$t('batchOperate.openCloseRestart.label.outOfBandIP')"
            :has-user-template="true"
        >
            <template #default="{scope}">
                {{$filter.emptyFilter(scope.row.iloIp)}}
            </template>
        </ui-table-column>
        <!-- 内网IPv4 -->
        <ui-table-column
            prop="privateIpv4"
            align="center"
            min-width="120"
            :label="$t('batchOperate.openCloseRestart.label.intranceIPv4')"
            :has-user-template="true"
        >
            <template #default="{scope}">
                {{$filter.emptyFilter(scope.row.privateIpv4)}}
            </template>
        </ui-table-column>
    </ui-table>
</template>
<script lang="ts" setup>
import {hasShowTooltip, tableClass} from 'utils/index.ts';

/**
 * props 类
*/
interface PropsType {
    selectArr: any;
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropsType>(), {});
</script>