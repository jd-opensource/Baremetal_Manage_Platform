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
            :label="$t('batchOperate.instanceName.label.instanceName')"
            :has-user-template="true"
        >
            <template #default="{scope}">
                <ui-tooltip
                    placement="bottom"
                    v-model="scope.row[`instanceNameTip${scope.$index}`].showTooltip"
                    :disabled="!scope.row[`instanceNameTip${scope.$index}`].showTooltip"
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
                        @mouseenter="hasShowTooltip($event, scope.row[`instanceNameTip${scope.$index}`], scope.row.instanceName)"
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
            :label="$t('batchOperate.instanceName.label.instanceId')"
            :has-user-template="true"
        >
            <template #default="{scope}">
                <ui-tooltip
                    placement="bottom"
                    v-model="scope.row[`instanceIdTip${scope.$index}`].showTooltip"
                    :disabled="!scope.row[`instanceIdTip${scope.$index}`].showTooltip"
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
                        @mouseenter="hasShowTooltip($event, scope.row[`instanceIdTip${scope.$index}`], scope.row.instanceId)"
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
            :label="$t('batchOperate.instanceName.label.outOfBandIP')"
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
            :label="$t('batchOperate.instanceName.label.intranceIPv4')"
            :has-user-template="true"
        >
            <template #default="{scope}">
                {{$filter.emptyFilter(scope.row.privateIpv4)}}
            </template>
        </ui-table-column>
    </ui-table>
</template>

<script lang="ts" setup>
import {CurrencyType} from '@utils/publicType';
import {hasShowTooltip, tableClass} from 'utils/index.ts';
import DeviceStaticData from 'staticData/device/index.ts';

interface PropsType {
    selectArr: any;
}

const props = withDefaults(defineProps<PropsType>(), {});
props.selectArr.forEach((item: CurrencyType, index: number) => {
    DeviceStaticData.batchEditInstanceNameTipData.forEach((t: string) => {
        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
    })
})

</script>
