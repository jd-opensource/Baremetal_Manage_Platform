<template>
    <div v-if="faultPushOperate.faultPushDiaLog.value">
        <fault-push
            :dia-log="faultPushOperate.faultPushDiaLog.value"
            :status="!faultPushOperate.pushItem.pushStatus"
            :alert-name="faultPushOperate.pushItem.alertName"
            :role-id="faultPushOperate.pushItem.ruleId"
            @dia-log-close="faultPushOperate.faultPushCancel"
            @determine-click="faultPushOperate.faultPushSure"
        />
    </div>
</template>
<script lang="ts" setup>

interface PropsType {
    faultPushOperate: any;
}

// defineEmits API 用来代替emit
const emitValue = defineEmits(['']);
// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>