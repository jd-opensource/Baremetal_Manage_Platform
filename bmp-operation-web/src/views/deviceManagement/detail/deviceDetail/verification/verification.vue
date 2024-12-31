<template>
    <div v-if="operate.securityVerificationDiaLog.value">
        <!-- 安全验证 -->
        <security-verification
            :dia-log="operate.securityVerificationDiaLog.value"
            :has-edit-operate="false"
            @dia-log-close="operate.verificationCancel"
            @has-edit-operate="operate.hasEditOperateClick"
        />
    </div>
</template>
<script lang="ts" setup>
// import OpenCloseRestart from './operate';
// defineEmits API 来替代 emits
interface PropsType {
    operate: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
// const operate = new OpenCloseRestart(props.deviceDetail);
</script>