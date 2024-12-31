<template>
    <!-- 设备移除、实例回收操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'currency remove-recovery'"
        :header-title="setInfo.title(title)"
        :src="($defInfo.imagePath('device') as unknown as string)"
        :disabled-status="!removeRecovery.checkboxStatus.value"
        :is-loading="removeRecovery.isLoading"
        @dia-log-close="removeRecovery.cancelClick"
        @determine-click="removeRecovery.determineClick"
    >
        <div class="currency-count">
            <!-- headerTitle提示 -->
            <p class="header-title">
                <span>{{setInfo.headerTitle(title)}}</span>
            </p>
            <div class="remove-recovery-count-table currency-count-table">
                <!-- 表格数据 -->
                <remove-recovery-table :table-data="removeRecovery.tableData"/>
            </div>
            <div class="currency-count-title">
                <ui-icon class="icon">
                    <CircleCloseFilled />
                </ui-icon>
                <div class="text-count" v-html="setInfo.tooltipCount(title)"></div>
            </div>
            <ui-checkbox
                class="remove-recovery-checkbox"
                v-model="removeRecovery.checkboxStatus.value"
                size="small"
            >
                <span class="text-tip">
                    {{setInfo.sureTip(title)}}
                </span>
            </ui-checkbox>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import SetInfo from './utils/setInfo';
import RemoveRecovery from './utils/removeRecovery';

// props 类
interface PropsType {
    diaLog: boolean;
    title: string;
};

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});
const setInfo = new SetInfo();

const removeRecovery = new RemoveRecovery(props, emitValue);

</script>
<style lang="scss">
@import './removeRecovery.scss';
</style>
