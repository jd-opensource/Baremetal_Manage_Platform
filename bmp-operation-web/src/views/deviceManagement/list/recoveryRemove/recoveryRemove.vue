<template>
    <div v-if="operate.removeRecoveryDiaLog.value">
        <remove-recovery
            :dia-log="operate.removeRecoveryDiaLog.value"
            :title="operate.removeRecoveryTitle.value"
            @dia-log-close="operate.removeRecoveryCancel"
            @determine-click="operate.removeRecoverySure"
        />
    </div>
</template>
<script setup lang="ts">
interface PropsType {
    operate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
