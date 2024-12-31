<template>
    <!-- 操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'fault-handling-confirm'"
        :is-loading="faultHandingConfirm.isLoading"
        :header-title="$t('faultHandingConfirm.header.title')"
        :src="($defInfo.imagePath('surveillance') as unknown as string)"
        @dia-log-close="faultHandingConfirm.cancelClick"
        @determine-click="faultHandingConfirm.determineClick"
    >

        <div class="fault-handling-confirm-content">
            <p class="warnning-tip">
                <!-- 提示icon -->
                <img
                    alt=""
                    :src="($defInfo.imagePath('warningTip') as unknown as string)"
                />
                <!-- 提示 -->
                <span>{{$t('faultHandingConfirm.header.tip1')}}</span>
            </p>
            <p class="tip-title">
                {{$t('faultHandingConfirm.header.tip2')}}
            </p>
            <div class="currency-count-table">
                <ui-config-provider :locale="uiLocale.locale">
                    <!-- 表格数据 -->
                    <ui-table
                        style="width: 100%"
                        :max-height="180"
                        :class="[tableClass(faultHandingConfirm.reactiveArr.tableData, false)]"
                        :data="faultHandingConfirm.reactiveArr.tableData"
                    >
                        <!-- sn -->
                        <ui-table-column
                            prop="sn"
                            align="center"
                            min-width="100"
                            :label="$t('faultHandingConfirm.label.sn')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-tooltip
                                    placement="bottom"
                                    v-model="scope.row[`snTip${scope.$index}`].showTooltip"
                                    :disabled="!scope.row[`snTip${scope.$index}`].showTooltip"
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
                                        @mouseenter="hasShowTooltip($event, scope.row[`snTip${scope.$index}`], scope.row.sn)"
                                    >
                                        <span>
                                            {{$filter.emptyFilter(scope.row.sn)}}
                                        </span>    
                                    </div>
                                </ui-tooltip>
                            </template>
                        </ui-table-column>
                        <!-- 故障名称 -->
                        <ui-table-column
                            prop="faultName"
                            align="center"
                            min-width="120"
                            :label="$t('faultHandingConfirm.label.faultName')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-tooltip
                                    placement="bottom"
                                    v-model="scope.row[`alertNameTip${scope.$index}`].showTooltip"
                                    :disabled="!scope.row[`alertNameTip${scope.$index}`].showTooltip"
                                >
                                    <template #content>
                                        <div class="set-tooltip-width">
                                            <span>
                                                {{scope.row.alertName}}
                                            </span>
                                        </div>
                                    </template>
                                    <div
                                        class="more-text-ellipsis"
                                        @mouseenter="hasShowTooltip($event, scope.row[`alertNameTip${scope.$index}`], scope.row.alertName)"
                                    >
                                        <span>
                                            {{$filter.emptyFilter(scope.row.alertName)}}
                                        </span>    
                                    </div>
                                </ui-tooltip>
                            </template>
                        </ui-table-column>
                        <!-- 故障类型 -->
                        <ui-table-column
                            prop="alertType"
                            min-width="130"
                            align="center"
                            :label="$t('faultHandingConfirm.label.faultType')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-tooltip
                                    placement="bottom"
                                    v-model="scope.row[`alertTypeTip${scope.$index}`].showTooltip"
                                    :disabled="!scope.row[`alertTypeTip${scope.$index}`].showTooltip"
                                >
                                    <template #content>
                                        <div class="set-tooltip-width">
                                            <span>
                                                {{scope.row.alertType}}
                                            </span>
                                        </div>
                                    </template>
                                    <div
                                        class="more-text-ellipsis"
                                        @mouseenter="hasShowTooltip($event,scope.row[`alertTypeTip${scope.$index}`], scope.row.alertType)"
                                    >
                                        <span>
                                            {{$filter.emptyFilter(scope.row.alertType)}}
                                        </span>    
                                    </div>
                                </ui-tooltip>
                            </template>
                        </ui-table-column>
                        <!-- 故障等级 -->
                        <ui-table-column
                            prop="alertLevel"
                            align="center"
                            min-width="100"
                            :label="$t('faultHandingConfirm.label.faultLevel')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                {{$filter.emptyFilter(scope.row.alertLevel)}}
                            </template>
                        </ui-table-column>
                        <!-- 故障描述 -->
                        <ui-table-column
                            prop="alertDesc"
                            align="center"
                            min-width="120"
                            :label="$t('faultHandingConfirm.label.faultDesc')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-tooltip
                                    placement="bottom"
                                    v-model="scope.row[`alertDescTip${scope.$index}`].showTooltip"
                                    :disabled="!scope.row[`alertDescTip${scope.$index}`].showTooltip"
                                >
                                    <template #content>
                                        <div class="set-tooltip-width">
                                            <span>
                                                {{scope.row.alertDesc}}
                                            </span>
                                        </div>
                                    </template>
                                    <div
                                        class="more-text-ellipsis"
                                        @mouseenter="hasShowTooltip($event, scope.row[`alertDescTip${scope.$index}`], scope.row.alertDesc)"
                                    >
                                        <span>
                                            {{$filter.emptyFilter(scope.row.alertDesc)}}
                                        </span>    
                                    </div>
                                </ui-tooltip>
                            </template>
                        </ui-table-column>
                        <!-- 首次故障报警时间 -->
                        <ui-table-column
                            prop="collectTime"
                            align="center"
                            min-width="160"
                            :label="$t('faultHandingConfirm.label.time')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-tooltip
                                    placement="bottom"
                                    v-model="scope.row[`collectTimeTip${scope.$index}`].showTooltip"
                                    :disabled="!scope.row[`collectTimeTip${scope.$index}`].showTooltip"
                                >
                                    <template #content>
                                        <div class="set-tooltip-width">
                                            <span>
                                                {{scope.row.collectTime}}
                                            </span>
                                        </div>
                                    </template>
                                    <div
                                        class="more-text-ellipsis"
                                        @mouseenter="hasShowTooltip($event, scope.row[`collectTimeTip${scope.$index}`], scope.row.collectTime)"
                                    >
                                        <span>
                                            {{$filter.emptyFilter(scope.row.collectTime)}}
                                        </span>    
                                    </div>
                                </ui-tooltip>
                            </template>
                        </ui-table-column>
                    </ui-table>
                </ui-config-provider>
            </div>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {msgTip, hasShowTooltip, tableClass} from 'utils/index.ts';
import {uiLocale, CurrentInstance, language} from 'utils/publicClass.ts';
import FalutLogStaticData from 'staticData/surveillance/faultLog/index.ts';

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
};


// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {
    diaLog: false, // 蒙层
});

/**
 * defineEmits API 用来代替emit
*/
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

class FaultHandingConfirm {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);
    reactiveArr: any = reactive({
        tableData: []
    })
    instanceMitt: any = new CurrentInstance();

    constructor() {
        // 监听fault-handing-confirm，获取表格数据
        this.instanceMitt?.proxy?.$Bus?.on('fault-handing-confirm', this.getTableFn);
    };

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.#requestFaultLogDeal();
    };

    #requestFaultLogDeal = async () => {
        try {
            const logid = this.reactiveArr.tableData.find((item: {logid: number}) => item.logid).logid;
            const logDealApi = await this.instanceMitt.proxy.$faultLogApi.faultLogDealAPI({logid});
            if (logDealApi?.data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        }
        finally {
            this.isLoading.value = false;
            this.cancelClick();
            emitValue('determine-click');
        }
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
        this.instanceMitt.proxy.$Bus.off('fault-handing-confirm', this.getTableFn);
    };

    getTableFn = (item: any) => {
        const newData = [item].map((ite: any, index: number) => {
            FalutLogStaticData.faultDisposeOfConfirmTip.forEach((t: any) => {
                Object.assign(ite, {[`${t}${index}`]: {showTooltip: false}})
            });
            return item;
        });
        this.reactiveArr.tableData = newData;
    };
};

const faultHandingConfirm = new FaultHandingConfirm();

</script>
<style lang="scss">
@import './faultHandlingConfirm.scss';
</style>
