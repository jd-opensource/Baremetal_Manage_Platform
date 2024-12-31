<template>
    <div v-if="operate.resetPasswordDiaLog.value">
        <reset-password
            :dia-log="operate.resetPasswordDiaLog.value"
            @dia-log-close="operate.resetPasswordCancel"
            @determine-click="operate.resetPasswordSure"
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
