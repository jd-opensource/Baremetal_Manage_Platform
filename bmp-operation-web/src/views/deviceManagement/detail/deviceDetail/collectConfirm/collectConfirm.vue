<template>
    <!-- 编辑备注 -->
    <div v-if="operate.resetCollectInfoDiaLog.value">
        <collect-info-confirm
            :dia-log="operate.resetCollectInfoDiaLog.value"
            :sn="deviceDetail.reactiveArr.detail?.sn"
            :has-clear-raid="deviceDetail.reactiveArr.detail?.collectStatus"
            @dia-log-close="operate.resetCollectInfoCancel"
            @determine-click="operate.resetCollectInfoSure"
            @opt-click="operate.resetCollectSuccessClick"
        >
        </collect-info-confirm>
    </div>
</template>
<script lang="ts" setup>
import DeviceDetail from '../detail';

// defineEmits API 来替代 emits
interface PropsType {
    operate: any;
    deviceDetail: DeviceDetail
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
