<template>
    <div v-if="operate.securityVerificationDiaLog.value">
        <!-- 安全验证 -->
        <security-verification
            :dia-log="operate.securityVerificationDiaLog.value"
            :has-edit-operate="operate.hasEditOperate.value"
            @dia-log-close="operate.securityVerificationCancel"
            @determine-click="operate.securityVerificationSure"
            @has-edit-operate="editIdcOperate.hasEditOperateClick"
        />
    </div>
</template>

<script lang="ts" setup>
interface PropsType {
    operate: any;
    editIdcOperate: any;
}
withDefaults(defineProps<PropsType>(), {});


</script>
