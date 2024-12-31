<template>
    <div v-if="operate.importDeviceDiaLog.value">
        <import-device
            :dia-log="operate.importDeviceDiaLog.value"
            @dia-log-close="operate.importDeviceCancel"
            @determine-click="operate.importDeviceSure"
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
