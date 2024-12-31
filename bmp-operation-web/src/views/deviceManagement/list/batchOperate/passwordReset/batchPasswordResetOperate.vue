<template>
    <div v-if="operate.batchResetPasswordDiaLog.value">
        <batch-reset-password
            :dia-log="operate.batchResetPasswordDiaLog.value"
            :check-arr="deviceList.reactiveArr.selectArr"
            @dia-log-close="operate.batchResetPasswordCancel"
            @determine-click="operate.batchResetPasswordSure"
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
