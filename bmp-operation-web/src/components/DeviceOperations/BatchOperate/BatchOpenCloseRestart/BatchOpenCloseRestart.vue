<template>
    <!-- 开关机重启操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'open-close-restart'"
        :is-loading="batchOpenCloseRestart.isLoading"
        :header-title="setTip.titleTip(types)"
        :src="($defInfo.imagePath('device') as unknown as string)"
        @dia-log-close="batchOpenCloseRestart.cancelClick"
        @determine-click="batchOpenCloseRestart.determineClick"
    >
        <div class="open-close-restart-content">
            <p class="tip-title">
                {{setTip.tooltipLeft(types)}}
                <span>{{selectArr.length}}</span>
                {{setTip.tooltipRight(types)}}
            </p>
            <!-- 表格数据 -->
            <batch-open-close-restart-table :select-arr="selectArr"/>
        </div>
    </custom-dia-log>
</template>

<script lang="ts" setup>
import SetTip from './utils/setTip';
import BatchOpenCloseRestart from './utils/batchOpenCloseRestart';
import {CurrencyType} from '@utils/publicType';

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
    selectArr: CurrencyType[];
    status?: string;
    types: string;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});
const setTip = new SetTip();

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const batchOpenCloseRestart = new BatchOpenCloseRestart(props, emitValue);

</script>
<style lang="scss">
@import './batchOpenCloseRestart.scss';
</style>
