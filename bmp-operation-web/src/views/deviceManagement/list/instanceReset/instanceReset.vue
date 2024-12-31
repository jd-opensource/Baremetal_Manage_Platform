<template>
    <div v-if="instanceReset.resetInstanceDiaLog.value">
        <reset-instance
            :dia-log="instanceReset.resetInstanceDiaLog.value"
            :instance-name="instanceReset.instanceName.value"
            :instance-id="instanceReset.instanceId.value"
            @dia-log-close="instanceReset.resetInstanceCancel"
            @determine-click="instanceReset.resetInstanceSure"
        />
    </div>
</template>

<script setup lang="ts">
interface PropsType {
    instanceReset: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
