<template>
    <div v-if="faultHandingOpt.faultHandingConfirmDiaLog.value">
        <fault-Handling-confirm
            :dia-log="faultHandingOpt.faultHandingConfirmDiaLog.value"
            @dia-log-close="faultHandingOpt.faultHandingConfirmCancel"
            @determine-click="faultHandingOpt.faultHandingConfirmCancelSure"
        />
    </div>
</template>
<script lang="ts" setup>

interface PropsType {
    faultHandingOpt: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>
