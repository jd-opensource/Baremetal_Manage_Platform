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
        <!-- 机柜编码 -->
        <ui-table-column
            prop="cabinet"
            align="center"
            min-width="90"
            :label="$t('publicTable.label.cabinetCode')"
            :has-user-template="true"
        >
            <template #default="{scope}">
                {{$filter.emptyFilter(scope.row.cabinet)}}
            </template>
        </ui-table-column>
        <!-- 所在U位 -->
        <ui-table-column
            prop="uPosition"
            align="center"
            min-width="120"
            :label="$t('publicTable.label.uBit')"
            :has-user-template="true"
        >
            <template #default="{scope}">
                {{$filter.emptyFilter(scope.row.uPosition)}}
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
        <!-- 内网IPv4 -->
        <ui-table-column
            prop="privateIpv4"
            align="center"
            min-width="120"
            :label="$t('publicTable.label.intranceIPv4')"
            :has-user-template="true"
        >
            <template #default="{scope}">
                {{$filter.emptyFilter(scope.row.privateIpv4)}}
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
