<template>
    <div v-if="operate.securityVerificationDiaLog.value">
        <!-- 安全验证 -->
        <security-verification
            :dia-log="operate.securityVerificationDiaLog.value"
            @dia-log-close="operate.securityVerificationCancel"
            @determine-click="operate.securityVerificationSure"
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
