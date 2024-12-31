<template>
    <div v-if="operate.resetInstanceDiaLog.value">
        <reset-instance
            :dia-log="operate.resetInstanceDiaLog.value"
            :instance-name="operate.instanceName.value"
            :instance-id="operate.instanceId.value"
            @dia-log-close="operate.resetInstanceCancel"
            @determine-click="operate.resetInstanceSure"
        />
    </div>
</template>
<script lang="ts" setup>
// defineEmits API 来替代 emits
interface PropsType {
    operate: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>