<template>
    <div v-if="operate.batchRecoveryInstanceDiaLog.value">
        <batch-recovery-instance
            :dia-log="operate.batchRecoveryInstanceDiaLog.value"
            :select-arr="deviceList.reactiveArr.selectArr"
            @dia-log-close="operate.batchRecoveryInstanceCancel"
            @determine-click="operate.batchRecoveryInstanceSure"
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
