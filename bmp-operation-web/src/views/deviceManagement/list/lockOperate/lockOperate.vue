<template>
    <div v-if="operate.lockDiaLog.value">
        <lock
            :dia-log="operate.lockDiaLog.value"
            :instance-name="operate.instanceName.value"
            :operate="operate.operate.value"
            :instance-id="operate.instanceId.value"
            @dia-log-close="operate.lockCancel"
            @determine-click="operate.lockSure"
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
