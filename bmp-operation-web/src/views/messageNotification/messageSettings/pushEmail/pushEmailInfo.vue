<template>
    <div v-if="pushEmailOpt.emailPushDiaLog.value">
        <email-push
            :dia-log="pushEmailOpt.emailPushDiaLog.value"
            :operate="pushEmailOpt.operate.value"
            @dia-log-close="pushEmailOpt.cancel"
            @determine-click="pushEmailOpt.sure"
            @error-click="pushEmailOpt.errorClick"
        >
        </email-push>
    </div>
</template>
<script setup lang="ts">

interface PropsType {
    pushEmailOpt: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>