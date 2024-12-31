<template>
    <!-- 操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'batch-recovery-instance'"
        :is-loading="batchRecoveryInstance.isLoading"
        :header-title="$t('batchOperate.recoveryInstance.title')"
        :src="($defInfo.imagePath('device') as unknown as string)"
        @dia-log-close="batchRecoveryInstance.cancelClick"
        @determine-click="batchRecoveryInstance.determineClick"
    >
        <div class="batch-recovery-instance-content">
            <p class="warnning-tip">
                <!-- 提示icon -->
                <img
                    alt=""
                    :src="($defInfo.imagePath('warningTip') as unknown as string)"
                />
                <!-- 提示 -->
                <span>{{$t('batchOperate.recoveryInstance.tip')}}</span>
            </p>
            <p class="tip-title">
                {{$t('batchOperate.recoveryInstance.recovery1')}}
                <span>{{selectArr.length}}</span>
                {{$t('batchOperate.recoveryInstance.recovery2')}}
            </p>
            <!-- 表格数据 -->
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
                    :label="$t('batchOperate.recoveryInstance.label.instanceName')"
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
                    :label="$t('batchOperate.recoveryInstance.label.instanceId')"
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
                    :label="$t('batchOperate.recoveryInstance.label.outOfBandIP')"
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
                    :label="$t('batchOperate.recoveryInstance.label.intranceIPv4')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        {{$filter.emptyFilter(scope.row.privateIpv4)}}
                    </template>
                </ui-table-column>
            </ui-table>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {BatchRecoveryInstanceAPI} from 'api/device/request.ts';
import {CurrencyType, SuccessType} from '@utils/publicType';
import {msgTip, hasShowTooltip, tableClass} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
    selectArr: CurrencyType[];
    status?: string;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

/**
 * defineEmits API 用来代替emit
*/
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

class BatchRecoveryInstance {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.requestRecoveryInstance();
    };

    requestRecoveryInstance = () => {
        BatchRecoveryInstanceAPI(
            {
                instanceIds: props.selectArr.map((item: CurrencyType) => item.instanceId)
            }
        )
        .then(({data} : {data: SuccessType}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        })
        .finally(() => {
            this.isLoading.value = false;
            this.cancelClick();
            emitValue('determine-click');
        })
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };
};

const batchRecoveryInstance = new BatchRecoveryInstance();
</script>
<style lang="scss">
@import './batchRecoveryInstance.scss';
</style>
