<template>
    <div v-if="operate.batchOpenCloseRestartDiaLog.value">
        <batch-open-close-restart
            :dia-log="operate.batchOpenCloseRestartDiaLog.value"
            :select-arr="deviceList.reactiveArr.selectArr"
            :types="operate.types.value"
            @dia-log-close="operate.batchOpenCloseRestartCancel"
            @determine-click="operate.batchOpenCloseRestartSure"
        />
    </div>
</template>
<script setup lang="ts">
interface PropsType {
    operate: any;
    deviceList: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
