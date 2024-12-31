<template>
    <div v-if="operate.resetSystemDiaLog.value">
        <reset-system
            :dia-log="operate.resetSystemDiaLog.value"
            @dia-log-close="operate.resetSystemCancel"
            @determine-click="operate.resetSystemSure"
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
